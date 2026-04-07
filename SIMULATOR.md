# futuapi4go OpenD 模拟器

本地模拟服务器，用于在没有真实 Futu OpenD 服务器的情况下测试 SDK。

## 项目状态

| 模块 | 状态 | 说明 |
|------|------|------|
| 核心服务器 | ✅ 已完成 | TCP 监听、协议解析(46字节头)、处理器注册 |
| 系统 API | ✅ 已完成 | InitConnect, KeepAlive, GetGlobalState, GetUserInfo |
| Qot 行情 API | ✅ 已完成 | 42 个 Handler 已注册 |
| Trd 交易 API | ✅ 已完成 | 13 个 Handler 已注册 |
| 推送模拟 | ✅ 已完成 | 11 个 Push Handler 已注册 |

## 架构设计

```
simulator/
├── server.go         # TCP 服务器核心 (46字节协议头, LittleEndian)
├── handlers.go       # 系统 API 处理器 (4 handlers)
├── handlers_qot.go    # 行情 API 处理器 (42 handlers)
├── handlers_trd.go    # 交易 API 处理器 (13 handlers)
├── handlers_push.go   # 推送处理器 (11 handlers)
└── simulator_test.go  # 测试
```

## 运行模拟器

```bash
cd examples/simulator
go run main.go
```

## 协议兼容性

模拟器使用与真实 Futu OpenD 完全相同的协议格式：

| 字段 | 大小 | 字节序 | 说明 |
|------|------|--------|------|
| Magic | 2 | - | "FT" |
| ProtoID | 4 | Little | 协议 ID |
| ProtoFmt | 4 | Little | 协议格式 (Protobuf) |
| ProtoVer | 2 | Little | 协议版本 |
| SerialNo | 4 | Little | 序列号 |
| BodyLen | 4 | Little | 包体长度 |
| BodySHA1 | 20 | - | SHA1 (暂未使用) |
| Reserved | 8 | - | 保留字段 |

**总 Header 长度: 46 字节**

## 已实现 Qot Handler (2101-2405)

| API | ProtoID | 状态 |
|-----|---------|------|
| GetBasicQot | 2101 | ✅ 完整实现 |
| GetKL | 2102 | ✅ 完整实现 |
| GetOrderBook | 2106 | ✅ 完整实现 |
| GetTicker | 2107 | ✅ 存根 |
| GetRT | 2108 | ✅ 存根 |
| GetSecuritySnapshot | 2110 | ✅ 存根 |
| GetBroker | 2111 | ✅ 存根 |
| GetStaticInfo | 2201 | ✅ 完整实现 |
| GetPlateSet | 2202 | ✅ 存根 |
| GetPlateSecurity | 2203 | ✅ 存根 |
| GetOwnerPlate | 2204 | ✅ 存根 |
| GetReference | 2205 | ✅ 存根 |
| GetTradeDate | 2206 | ✅ 完整实现 |
| RequestTradeDate | 2207 | ✅ 存根 |
| GetMarketState | 2208 | ✅ 存根 |
| GetSuspend | 2209 | ✅ 存根 |
| GetCodeChange | 2210 | ✅ 存根 |
| GetFutureInfo | 2211 | ✅ 存根 |
| GetIpoList | 2212 | ✅ 存根 |
| GetHoldingChangeList | 2213 | ✅ 存根 |
| RequestRehab | 2214 | ✅ 存根 |
| GetCapitalFlow | 2301 | ✅ 存根 |
| GetCapitalDistribution | 2302 | ✅ 存根 |
| StockFilter | 2303 | ✅ 存根 |
| GetOptionChain | 2304 | ✅ 存根 |
| GetOptionExpirationDate | 2305 | ✅ 存根 |
| GetWarrant | 2306 | ✅ 存根 |
| GetUserSecurity | 2401 | ✅ 存根 |
| GetUserSecurityGroup | 2402 | ✅ 存根 |
| ModifyUserSecurity | 2403 | ✅ 存根 |
| GetPriceReminder | 2404 | ✅ 存根 |
| SetPriceReminder | 2405 | ✅ 存根 |
| Subscribe | 3001 | ✅ 完整实现 |
| GetSubInfo | 3002 | ✅ 完整实现 |
| RegQotPush | 3003 | ✅ 完整实现 |

## 使用方法

### 启动模拟器

```go
package main

import (
    "log"

    "gitee.com/shing1211/futuapi4go/simulator"
)

func main() {
    srv := simulator.New("127.0.0.1:11111")
    srv.RegisterDefaultHandlers()  // 系统 handlers
    srv.RegisterQotHandlers()     // 行情 handlers
    
    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }
    
    log.Println("Simulator started on 127.0.0.1:11111")
    <-make(chan struct{}) // 阻塞直到收到信号
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
    cli := futuapi.New()
    if err := cli.Connect("127.0.0.1:11111"); err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

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

## 下一步计划

1. ✅ 实现 Qot 行情 API Handler 框架
2. ✅ 完善存根 Handler，添加真实模拟数据
3. ✅ 实现 Trd 交易 API Handler
4. ✅ 添加推送模拟支持
5. ✅ 添加可配置模拟数据

## 高级模拟器规划

### 1. 价格/订单模拟引擎
- [ ] 实时价格变动模拟（tick-by-tick）
- [ ] 随机游走/趋势模拟
- [ ] 订单簿模拟（买卖盘深度）
- [ ] 订单撮合引擎（模拟成交）

### 2. 错误/边界测试注入
- [ ] 网络延迟模拟（固定/随机）
- [ ] 网络故障注入（断连、超时）
- [ ] 返回错误响应（各种 RetType）
- [ ] 边界值测试数据

### 3. 场景录制/回放
- [ ] 记录真实市场数据
- [ ] 回放历史数据
- [ ] 批量测试场景

---

*最后更新: 2026-04-07*
