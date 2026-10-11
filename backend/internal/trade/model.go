// Package trade 交割单业务模块。
package trade

// Trade 一笔合约交割单。
type Trade struct {
	ID     int64   `json:"id"`
	Date   string  `json:"date"`
	Symbol string  `json:"symbol"`
	Side   string  `json:"side"` // long / short
	Lots   float64 `json:"lots"`
	Entry  float64 `json:"entry"`
	Exit   float64 `json:"exit"`
	Stop   float64 `json:"stop"`
	PnL    float64 `json:"pnl"`
	Review string  `json:"review"`
}

// Stats 交割单分析指标。
type Stats struct {
	Total       int     `json:"total"`
	Wins        int     `json:"wins"`
	WinRate     float64 `json:"winRate"`
	TotalPnL    float64 `json:"totalPnl"`
	Expectancy  float64 `json:"expectancy"`
	AvgR        float64 `json:"avgR"`
	MaxDrawdown float64 `json:"maxDrawdown"`
}
