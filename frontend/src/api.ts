import type { Stats, Trade } from './types'

const BASE = '/api'

async function req<T>(url: string, init?: RequestInit): Promise<T> {
  const res = await fetch(BASE + url, {
    headers: { 'Content-Type': 'application/json' },
    ...init,
  })
  if (!res.ok) {
    throw new Error(await res.text())
  }
  return res.json() as Promise<T>
}

export const listTrades = () => req<Trade[]>('/trades')

export const createTrade = (t: Omit<Trade, 'id'>) =>
  req<Trade>('/trades', { method: 'POST', body: JSON.stringify(t) })

export const getStats = () => req<Stats>('/trades/stats')
