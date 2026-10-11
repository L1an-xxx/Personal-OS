import { useEffect, useRef, useState } from 'react'
import {
  Button,
  Card,
  Col,
  Form,
  Input,
  InputNumber,
  message,
  Row,
  Select,
  Statistic,
  Table,
  Typography,
} from 'antd'
import * as echarts from 'echarts'
import { createTrade, getStats, listTrades } from './api'
import type { Stats, Trade } from './types'

const sideOptions = [
  { label: '做多 long', value: 'long' },
  { label: '做空 short', value: 'short' },
]

function PnlChart({ trades }: { trades: Trade[] }) {
  const ref = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (!ref.current) return
    const chart = echarts.init(ref.current)
    const sorted = [...trades].sort((a, b) => a.date.localeCompare(b.date))
    let cum = 0
    const data = sorted.map((t) => {
      cum += t.pnl
      return cum
    })
    chart.setOption({
      title: { text: '累计盈亏曲线' },
      tooltip: { trigger: 'axis' },
      xAxis: { type: 'category', data: sorted.map((t) => t.date) },
      yAxis: { type: 'value' },
      series: [{ name: '累计盈亏', type: 'line', smooth: true, data }],
    })
    return () => chart.dispose()
  }, [trades])

  return <div ref={ref} style={{ height: 320 }} />
}

function App() {
  const [trades, setTrades] = useState<Trade[]>([])
  const [stats, setStats] = useState<Stats | null>(null)
  const [loading, setLoading] = useState(false)
  const [form] = Form.useForm()

  const load = async () => {
    setLoading(true)
    try {
      const [t, s] = await Promise.all([listTrades(), getStats()])
      setTrades(t)
      setStats(s)
    } catch (e) {
      message.error('后端连接失败：' + String(e))
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    load()
  }, [])

  const onSubmit = async (values: any) => {
    try {
      await createTrade(values)
      message.success('已入库')
      form.resetFields()
      await load()
    } catch (e) {
      message.error(String(e))
    }
  }

  const columns = [
    { title: '日期', dataIndex: 'date' },
    { title: '品种', dataIndex: 'symbol' },
    { title: '方向', dataIndex: 'side' },
    { title: '手数', dataIndex: 'lots' },
    { title: '入场', dataIndex: 'entry' },
    { title: '出场', dataIndex: 'exit' },
    { title: '止损', dataIndex: 'stop' },
    { title: '盈亏', dataIndex: 'pnl' },
    { title: '复盘', dataIndex: 'review' },
  ]

  return (
    <div style={{ padding: 24, maxWidth: 1200, margin: '0 auto' }}>
      <Typography.Title level={3}>Personal OS · 交割单</Typography.Title>

      <Row gutter={16}>
        <Col span={10}>
          <Card title="新增交割单">
            <Form form={form} layout="vertical" onFinish={onSubmit}>
              <Form.Item name="date" label="日期" rules={[{ required: true }]}>
                <Input placeholder="2026-10-11" />
              </Form.Item>
              <Form.Item name="symbol" label="品种" rules={[{ required: true }]}>
                <Input placeholder="XAUUSD" />
              </Form.Item>
              <Form.Item name="side" label="方向" rules={[{ required: true }]}>
                <Select options={sideOptions} />
              </Form.Item>
              <Form.Item name="lots" label="手数" rules={[{ required: true }]}>
                <InputNumber style={{ width: '100%' }} />
              </Form.Item>
              <Form.Item name="entry" label="入场价" rules={[{ required: true }]}>
                <InputNumber style={{ width: '100%' }} />
              </Form.Item>
              <Form.Item name="exit" label="出场价" rules={[{ required: true }]}>
                <InputNumber style={{ width: '100%' }} />
              </Form.Item>
              <Form.Item name="stop" label="止损价" rules={[{ required: true }]}>
                <InputNumber style={{ width: '100%' }} />
              </Form.Item>
              <Form.Item name="pnl" label="盈亏（金额）" rules={[{ required: true }]}>
                <InputNumber style={{ width: '100%' }} />
              </Form.Item>
              <Form.Item name="review" label="复盘">
                <Input.TextArea rows={2} />
              </Form.Item>
              <Button type="primary" htmlType="submit" block>
                入库
              </Button>
            </Form>
          </Card>
        </Col>

        <Col span={14}>
          <Card title="统计" loading={loading}>
            <Row gutter={16}>
              <Col span={8}>
                <Statistic title="总笔数" value={stats?.total ?? 0} />
              </Col>
              <Col span={8}>
                <Statistic
                  title="胜率"
                  value={stats ? (stats.winRate * 100).toFixed(1) + '%' : '—'}
                />
              </Col>
              <Col span={8}>
                <Statistic title="总盈亏" value={stats?.totalPnl ?? 0} precision={2} />
              </Col>
              <Col span={8}>
                <Statistic title="期望/笔" value={stats?.expectancy ?? 0} precision={2} />
              </Col>
              <Col span={8}>
                <Statistic title="R 倍数" value={stats?.avgR ?? 0} precision={2} />
              </Col>
              <Col span={8}>
                <Statistic title="最大回撤" value={stats?.maxDrawdown ?? 0} precision={2} />
              </Col>
            </Row>
          </Card>
          <Card title="累计盈亏曲线" style={{ marginTop: 16 }}>
            <PnlChart trades={trades} />
          </Card>
        </Col>
      </Row>

      <Card title="交割单列表" style={{ marginTop: 16 }}>
        <Table
          rowKey="id"
          size="small"
          loading={loading}
          dataSource={trades}
          columns={columns}
          pagination={{ pageSize: 10 }}
        />
      </Card>
    </div>
  )
}

export default App
