# futuapi4go OpenD 模拟器实现计划

本文档详细规划了 OpenD 模拟器的完整实现，用于在没有真实 Futu OpenD 服务器的情况下测试 SDK。

## 实现目标

创建一个本地模拟服务器，能够响应 SDK 的所有 API 请求，返回合理的模拟数据。

## 架构设计

```
futuapi4go/
└── simulator/
    ├── server.go        # TCP 服务器核心
    ├── handlers.go      # API 处理器
    ├── handlers_qot.go  # 行情 API 处理器
    ├── handlers_trd.go  # 交易 API 处理器
    ├── handlers_sys.go  # 系统 API 处理器
    ├── handlers_push.go # 推送模拟处理器
    ├── data/            # 模拟数据存储
    │   └── mock_data.go # 配置化模拟数据
    └── main.go          # 独立运行入口
```

## 实现阶段

### 阶段一：核心服务器 ✅ 已完成

| 功能 | 状态 | 说明 |
|------|------|------|
| TCP 监听 | ✅ 完成 | 监听指定端口 |
| 协议解析 | ✅ 完成 | 解析 FT 协议头 |
| 处理器注册 | ✅ 完成 | 注册/注销处理器 |
| 连接管理 | ✅ 完成 | 并发连接处理 |

### 阶段二：系统 API 处理器 ✅ 部分完成

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| InitConnect | 1001 | ✅ 完成 | 初始化连接 |
| KeepAlive | 1002 | ✅ 完成 | 心跳响应 |
| GetGlobalState | 1004 | ✅ 完成 | 全局状态 |
| GetUserInfo | 1005 | ✅ 完成 | 用户信息 |
| GetDelayStatistics | 1006 | ⏳ 待实现 | 延迟统计 |
| Notify | 1003 | ⏳ 待实现 | 系统通知推送 |
| Verification | 8001 | ⏳ 待实现 | 验证接口 |

### 阶段三：Qot 行情 API 处理器

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetBasicQot | 2101 | ⏳ 待实现 | 获取实时行情 |
| GetKL | 2102 | ⏳ 待实现 | 获取实时K线 |
| GetHistoryKL | 2103 | ⏳ 待实现 | 获取历史K线 |
| RequestHistoryKL | 2104 | ⏳ 待实现 | 请求历史K线(异步) |
| GetOrderBook | 2106 | ⏳ 待实现 | 获取订单簿 |
| GetTicker | 2107 | ⏳ 待实现 | 获取逐笔成交 |
| GetRT | 2108 | ⏳ 待实现 | 获取实时分时数据 |
| GetSecuritySnapshot | 2110 | ⏳ 待实现 | 获取股票快照 |
| GetBroker | 2111 | ⏳ 待实现 | 获取买卖队列 |
| GetStaticInfo | 2201 | ⏳ 待实现 | 获取股票静态信息 |
| GetPlateSet | 2202 | ⏳ 待实现 | 获取板块集合 |
| GetPlateSecurity | 2203 | ⏳ 待实现 | 获取板块成分股 |
| GetOwnerPlate | 2204 | ⏳ 待实现 | 获取所属板块 |
| GetReference | 2205 | ⏳ 待实现 | 获取正股相关数据 |
| GetTradeDate | 2206 | ⏳ 待实现 | 获取交易日 |
| RequestTradeDate | 2207 | ⏳ 待实现 | 请求交易日 |
| GetMarketState | 2208 | ⏳ 待实现 | 获取市场状态 |
| GetSuspend | 2209 | ⏳ 待实现 | 获取停牌信息 |
| GetCodeChange | 2210 | ⏳ 待实现 | 获取代码变更信息 |
| GetFutureInfo | 2211 | ⏳ 待实现 | 获取期货信息 |
| GetIpoList | 2212 | ⏳ 待实现 | 获取IPO列表 |
| GetHoldingChangeList | 2213 | ⏳ 待实现 | 获取持仓变化列表 |
| RequestRehab | 2214 | ⏳ 待实现 | 请求复权数据 |
| GetCapitalFlow | 2301 | ⏳ 待实现 | 获取资金流向 |
| GetCapitalDistribution | 2302 | ⏳ 待实现 | 获取资金分布 |
| StockFilter | 2303 | ⏳ 待实现 | 股票筛选 |
| GetOptionChain | 2304 | ⏳ 待实现 | 获取期权链 |
| GetOptionExpirationDate | 2305 | ⏳ 待实现 | 获取期权到期日 |
| GetWarrant | 2306 | ⏳ 待实现 | 获取窝轮信息 |
| GetUserSecurity | 2401 | ⏳ 待实现 | 获取用户自选股 |
| GetUserSecurityGroup | 2402 | ⏳ 待实现 | 获取用户自选股分组 |
| ModifyUserSecurity | 2403 | ⏳ 待实现 | 修改用户自选股 |
| GetPriceReminder | 2404 | ⏳ 待实现 | 获取价格提醒 |
| SetPriceReminder | 2405 | ⏳ 待实现 | 设置价格提醒 |
| Subscribe | 3001 | ⏳ 待实现 | 订阅实时行情 |
| GetSubInfo | 3002 | ⏳ 待实现 | 获取订阅信息 |
| RegQotPush | 3003 | ⏳ 待实现 | 注册行情推送 |

### 阶段四：Trd 交易 API 处理器

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetAccList | 4001 | ⏳ 待实现 | 获取账户列表 |
| UnlockTrade | 4002 | ⏳ 待实现 | 解锁交易密码 |
| GetFunds | 4003 | ⏳ 待实现 | 获取资金信息 |
| GetOrderFee | 4004 | ⏳ 待实现 | 获取订单费用 |
| GetMarginRatio | 4005 | ⏳ 待实现 | 获取保证金比例 |
| GetMaxTrdQtys | 4006 | ⏳ 待实现 | 获取最大交易数量 |
| PlaceOrder | 5001 | ⏳ 待实现 | 下单 |
| ModifyOrder | 5002 | ⏳ 待实现 | 修改订单 |
| GetOrderList | 5003 | ⏳ 待实现 | 查询订单列表 |
| GetHistoryOrderList | 5004 | ⏳ 待实现 | 查询历史订单 |
| GetOrderFillList | 5005 | ⏳ 待实现 | 查询成交列表 |
| GetHistoryOrderFillList | 5006 | ⏳ 待实现 | 查询历史成交 |
| GetPositionList | 6001 | ⏳ 待实现 | 获取持仓列表 |
| GetFlowSummary | 2226 | ⏳ 待实现 | 获取账户资金流水 |
| ReconfirmOrder | 7004 | ⏳ 待实现 | 订单确认 |
| SubAccPush | 7005 | ⏳ 待实现 | 账户推送订阅 |

### 阶段五：推送模拟

| 推送类型 | ProtoID | 状态 | 说明 |
|---------|---------|------|------|
| Qot_UpdateBasicQot | 3101 | ⏳ 待实现 | 实时行情推送 |
| Qot_UpdateKL | 3102 | ⏳ 待实现 | K线推送 |
| Qot_UpdateOrderBook | 3103 | ⏳ 待实现 | 订单簿推送 |
| Qot_UpdateTicker | 3104 | ⏳ 待实现 | 逐笔成交推送 |
| Qot_UpdateRT | 3105 | ⏳ 待实现 | 分时数据推送 |
| Qot_UpdateBroker | 3106 | ⏳ 待实现 | 经纪商队列推送 |
| Qot_UpdatePriceReminder | 3107 | ⏳ 待实现 | 价格提醒推送 |
| Trd_UpdateOrder | 7001 | ⏳ 待实现 | 订单状态推送 |
| Trd_UpdateOrderFill | 7002 | ⏳ 待实现 | 成交推送 |
| Trd_Notify | 7003 | ⏳ 待实现 | 交易通知推送 |

### 阶段六：高级功能

| 功能 | 状态 | 说明 |
|------|------|------|
| 可配置模拟数据 | ⏳ 待实现 | JSON/YAML 配置 |
| 随机价格波动 | ⏳ 待实现 | 模拟真实市场 |
| 订单状态机 | ⏳ 待实现 | 模拟订单生命周期 |
| 错误场景模拟 | ⏳ 待实现 | 模拟各种错误情况 |
| 性能测试模式 | ⏳ 待实现 | 压力测试支持 |

## 使用方法

### 快速启动

```go
package main

import (
    "log"

    "gitee.com/shing1211/futuapi4go/simulator"
)

func main() {
    srv := simulator.New("127.0.0.1:11111")
    srv.RegisterDefaultHandlers()
    
    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }
    
    log.Println("Simulator started on 127.0.0.1:11111")
    log.Println("Press Ctrl+C to stop")
    
    // 阻塞直到收到退出信号
    <-make(chan struct{})
}
```

### 使用模拟器测试 SDK

```go
package main

import (
    "log"

    futuapi "gitee.com/shing1211/futuapi4go/client"
    "gitee.com/shing1211/futuapi4go/qot"
    "gitee.com/shing1211/futuapi4go/pb/qotcommon"
)

func main() {
    // 连接到模拟器
    cli := futuapi.New()
    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    // 查询行情 - 模拟器会返回预设数据
    market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
    securities := []*qotcommon.Security{
        {Market: &market, Code: func() *string { s := "00700"; return &s }()},
    }

    result, err := qot.GetBasicQot(cli, securities)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Received: %+v", result)
}
```

## 模拟数据设计

### 股票数据

```go
type MockStockData struct {
    Code     string
    Name     string
    Market   int32
    CurPrice float64
    OpenPrice float64
    HighPrice float64
    LowPrice  float64
    Volume   int64
}
```

### 账户数据

```go
type MockAccount struct {
    AccID      uint64
    AccName    string
    TrdEnv     int32
    TrdMarket  int32
    Cash       float64
    Positions  []MockPosition
    Orders     []MockOrder
}
```

## 测试计划

1. **单元测试** - 每个处理器独立测试
2. **集成测试** - 模拟器 + SDK 端到端测试
3. **性能测试** - 并发连接测试
4. **错误处理测试** - 网络异常、超时等场景

---

*最后更新: 2026-04-07*