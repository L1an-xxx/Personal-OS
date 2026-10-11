# frontend

React + TypeScript + Vite + Ant Design + ECharts。

## 运行

先启动后端（`cd backend && go run .`，监听 8080），再：

```bash
cd frontend
pnpm install   # 或 npm install
pnpm run dev   # 或 npm run dev
```

打开 `http://localhost:5173`（已配置 `/api` 代理到后端 8080，无跨域问题）。

## 页面

- 新增交割单表单 → `POST /api/trades`（入库）
- 统计卡片 → `GET /api/trades/stats`（胜率 / 期望 / R 倍数 / 最大回撤）
- 累计盈亏曲线（ECharts）
- 交割单列表 → `GET /api/trades`
