# Personal OS

前后端分离的个人平台：记录资产、收支、合约交割单、力扣，辅助成为数字游民与极客。

## 技术栈

- 后端：Go + Gin + SQLite（`backend/`，模块化 `internal/<模块>/`）
- 前端：React + TypeScript + Vite + Ant Design + ECharts（`frontend/`）

## 快速开始

```bash
# 后端（终端 1）
cd backend
go mod tidy
go run .

# 前端（终端 2）
cd frontend
pnpm install
pnpm run dev
```

打开 http://localhost:5173

## 模块

- [x] 交割单：入库 / 列表 / 分析（胜率·期望·R 倍数·最大回撤）+ 累计盈亏曲线
- [ ] 资产、收支、力扣（待需求文档）

## 文档

- 每日进度：`docs/进度.md`
- 个人发展基准：`docs/01-总纲.md`
- Web 需求基线：`docs/03-需求.md`
- 编码规范：`docs/04-编码规范.md`
