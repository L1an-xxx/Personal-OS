# backend

Go + Gin + SQLite（`modernc.org/sqlite`，纯 Go，无需 CGO）。

## 结构

```
backend/
├── main.go                 # 启动 + 装配（无业务代码）
└── internal/
    └── trade/
        ├── model.go        # Trade / Stats 结构体
        ├── store.go        # 数据库访问
        └── handler.go      # HTTP 处理
```

## 运行

```bash
cd backend
go mod tidy
go run .
```

服务启动在 `http://localhost:8080`，数据库文件 `personalos.db` 会自动创建。

## 接口

- `GET  /api/health` — 健康检查
- `POST /api/trades` — 新增一笔交割单（JSON）
- `GET  /api/trades` — 列出全部交割单
- `GET  /api/trades/stats` — 分析：总笔数 / 胜率 / 总盈亏 / 期望 / R 倍数 / 最大回撤

## 交割单字段

`date, symbol, side(long|short), lots, entry, exit, stop, pnl, review`

示例（POST /api/trades）：

```json
{
  "date": "2026-10-11",
  "symbol": "XAUUSD",
  "side": "long",
  "lots": 0.1,
  "entry": 2650.0,
  "exit": 2660.0,
  "stop": 2640.0,
  "pnl": 100.0,
  "review": "回踩做多"
}
```
