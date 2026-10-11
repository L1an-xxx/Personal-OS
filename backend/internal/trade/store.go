package trade

import (
	"database/sql"
	"math"
)

const schema = `
CREATE TABLE IF NOT EXISTS trades (
	id     INTEGER PRIMARY KEY AUTOINCREMENT,
	date   TEXT NOT NULL,
	symbol TEXT NOT NULL,
	side   TEXT NOT NULL,
	lots   REAL NOT NULL,
	entry  REAL NOT NULL,
	exit   REAL NOT NULL,
	stop   REAL NOT NULL,
	pnl    REAL NOT NULL,
	review TEXT NOT NULL DEFAULT ''
);
`

// Store 封装交割单的数据库访问。
type Store struct {
	db *sql.DB
}

// NewStore 返回绑定 db 的 Store。
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Migrate 建表（幂等）。
func (s *Store) Migrate() error {
	_, err := s.db.Exec(schema)
	return err
}

// Create 新增一笔交割单，并把自增 ID 写回 t。
func (s *Store) Create(t *Trade) error {
	res, err := s.db.Exec(
		`INSERT INTO trades (date, symbol, side, lots, entry, exit, stop, pnl, review)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		t.Date, t.Symbol, t.Side, t.Lots, t.Entry, t.Exit, t.Stop, t.PnL, t.Review,
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = id
	return nil
}

// List 按日期升序返回全部交割单。
func (s *Store) List() ([]Trade, error) {
	rows, err := s.db.Query(
		`SELECT id, date, symbol, side, lots, entry, exit, stop, pnl, review
		 FROM trades ORDER BY date ASC, id ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trades := make([]Trade, 0)
	for rows.Next() {
		var t Trade
		if err := rows.Scan(&t.ID, &t.Date, &t.Symbol, &t.Side, &t.Lots, &t.Entry, &t.Exit, &t.Stop, &t.PnL, &t.Review); err != nil {
			return nil, err
		}
		trades = append(trades, t)
	}
	return trades, rows.Err()
}

// Stats 计算胜率、期望、R 倍数、最大回撤等指标。
func (s *Store) Stats() (Stats, error) {
	rows, err := s.db.Query(`SELECT pnl, entry, stop, lots FROM trades ORDER BY date ASC, id ASC`)
	if err != nil {
		return Stats{}, err
	}
	defer rows.Close()

	var st Stats
	var cum, peak, rSum float64
	var rCount int
	for rows.Next() {
		var pnl, entry, stop, lots float64
		if err := rows.Scan(&pnl, &entry, &stop, &lots); err != nil {
			return Stats{}, err
		}
		st.Total++
		if pnl > 0 {
			st.Wins++
		}
		st.TotalPnL += pnl

		if risk := math.Abs(entry-stop) * lots; risk > 0 {
			rSum += pnl / risk
			rCount++
		}

		cum += pnl
		if cum > peak {
			peak = cum
		}
		if dd := peak - cum; dd > st.MaxDrawdown {
			st.MaxDrawdown = dd
		}
	}
	if err := rows.Err(); err != nil {
		return Stats{}, err
	}

	if st.Total > 0 {
		st.WinRate = float64(st.Wins) / float64(st.Total)
		st.Expectancy = st.TotalPnL / float64(st.Total)
	}
	if rCount > 0 {
		st.AvgR = rSum / float64(rCount)
	}
	return st, nil
}
