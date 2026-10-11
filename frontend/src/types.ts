export interface Trade {
  id: number
  date: string
  symbol: string
  side: string
  lots: number
  entry: number
  exit: number
  stop: number
  pnl: number
  review: string
}

export interface Stats {
  total: number
  wins: number
  winRate: number
  totalPnl: number
  expectancy: number
  avgR: number
  maxDrawdown: number
}
