# futuapi4go 用户指南

本指南面向量化交易者，介绍如何使用 futuapi4go SDK 进行市场数据查询和交易操作。

## 目录

1. [快速开始](#快速开始)
2. [连接管理](#连接管理)
3. [市场数据查询](#市场数据查询)
4. [交易操作](#交易操作)
5. [实时推送](#实时推送)
6. [常见问题](#常见问题)

---

## 快速开始

### 安装

```bash
go get gitee.com/shing1211/futuapi4go
```

### 基本使用流程

```go
package main

import (
    "fmt"
    "log"
    futuapi "gitee.com/shing1211/futuapi4go/client"
    "gitee.com/shing1211/futuapi4go/qot"
    "gitee.com/shing1211/futuapi4go/pb/qotcommon"
)

func main() {
    // 1. 创建客户端
    cli := futuapi.New()
    
    // 2. 连接 OpenD
    err := cli.Connect("127.0.0.1:11111")
    if err != nil {
        log.Fatal(err)
    }
    defer cli.Close()
    
    // 3. 调用 API
    market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
    code := "00700"
    securities := []*qotcommon.Security{
        {Market: &market, Code: &code},
    }
    
    result, err := qot.GetBasicQot(cli, securities)
    if err != nil {
        log.Fatal(err)
    }
    
    // 4. 处理结果
    for _, bq := range result {
        fmt.Printf("%s %s: 现价=%.2f\n", 
            bq.Security.GetCode(), bq.Name, bq.CurPrice)
    }
}
```

---

## 连接管理

### 创建连接

```go
cli := futuapi.New()
err := cli.Connect("127.0.0.1:11111")
if err != nil {
    log.Fatal(err)
}
defer cli.Close()
```

### 初始化连接（获取连接ID）

```go
// 初始化连接，获取用户信息
userInfo, err := sys.InitConnect(cli, "your_app_id", "your_hash")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("用户: %s, 连接ID: %d\n", userInfo.GetNickName(), userInfo.GetConnectionID())
```

### 心跳保活

SDK 自动发送心跳包维持连接，无需手动操作。

---

## 市场数据查询

### 获取实时行情 (GetBasicQot)

```go
market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
code := "00700"
securities := []*qotcommon.Security{
    {Market: &market, Code: &code},
}

result, err := qot.GetBasicQot(cli, securities)
if err != nil {
    log.Fatal(err)
}

for _, bq := range result {
    fmt.Printf("%s: 现价=%.2f, 涨跌=%.2f%%\n",
        bq.Name, bq.CurPrice, bq.ChangeRate)
}
```

### 获取 K 线 (GetKL)

```go
req := &qot.GetKLRequest{
    Security:  &qotcommon.Security{Market: &market, Code: &code},
    RehabType: int32(qotcommon.RehabType_RehabType_None),
    KLType:    int32(qotcommon.KLType_KLType_Day),
    ReqNum:    100,
}

result, err := qot.GetKL(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, kl := range result.KLList {
    fmt.Printf("%s: 开=%.2f, 高=%.2f, 低=%.2f, 收=%.2f\n",
        kl.Time, kl.OpenPrice, kl.HighPrice, kl.LowPrice, kl.ClosePrice)
}
```

### 获取订单簿 (GetOrderBook)

```go
req := &qot.GetOrderBookRequest{
    Security: &qotcommon.Security{Market: &market, Code: &code},
    Num:      10,
}

result, err := qot.GetOrderBook(cli, req)
if err != nil {
    log.Fatal(err)
}

fmt.Println("买方:")
for _, bid := range result.OrderBookBidList {
    fmt.Printf("  价格=%.2f, 成交量=%d\n", bid.Price, bid.Volume)
}

fmt.Println("卖方:")
for _, ask := range result.OrderBookAskList {
    fmt.Printf("  价格=%.2f, 成交量=%d\n", ask.Price, ask.Volume)
}
```

### 获取分时数据 (GetRT)

```go
req := &qot.GetRTRequest{
    Security: &qotcommon.Security{Market: &market, Code: &code},
}

result, err := qot.GetRT(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, rt := range result.RTList {
    fmt.Printf("%s: 价格=%.2f, 成交量=%d\n",
        rt.Time, rt.Price, rt.Volume)
}
```

### 获取资金流向 (GetCapitalFlow)

```go
req := &qot.GetCapitalFlowRequest{
    Security:   &qotcommon.Security{Market: &market, Code: &code},
    PeriodType: 1, // 日线
}

result, err := qot.GetCapitalFlow(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, f := range result.FlowItemList {
    fmt.Printf("%s: 主力流入=%.2f\n", f.Time, f.MainInFlow)
}
```

### 股票筛选 (StockFilter)

```go
req := &qot.StockFilterRequest{
    Begin:  0,
    Num:    10,
    Market: int32(qotcommon.QotMarket_QotMarket_HK_Security),
    BaseFilterList: []*qotstockfilter.BaseFilter{
        {
            FieldName:  int32(qotstockfilter.StockField_StockField_CurPrice),
            FilterMin:  proto.Float64(10.0),
            FilterMax:  proto.Float64(100.0),
            IsNoFilter: proto.Bool(false),
        },
    },
}

result, err := qot.StockFilter(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, d := range result.DataList {
    fmt.Printf("%s: %s\n", d.Security.GetCode(), d.Name)
}
```

### 获取期权链 (GetOptionChain)

```go
req := &qot.GetOptionChainRequest{
    Owner:      &qotcommon.Security{Market: &market, Code: &code},
    BeginTime:  "2024-01-01",
    EndTime:    "2024-12-31",
    DataFilter: nil,
}

result, err := qot.GetOptionChain(cli, req)
if err != nil {
    log.Fatal(err)
}

for _, chain := range result.OptionChain {
    fmt.Printf("行权日: %s\n", chain.StrikeTime)
    for _, opt := range chain.Option {
        if opt.Call != nil {
            fmt.Printf("  认购: %s\n", opt.Call.GetCode())
        }
        if opt.Put != nil {
            fmt.Printf("  认沽: %s\n", opt.Put.GetCode())
        }
    }
}
```

---

## 交易操作

### 解锁交易

```go
// 必须先解锁才能进行交易
err = trd.UnlockTrade(cli, "your_trade_password")
if err != nil {
    log.Fatal(err)
}
```

### 查询账户资金

```go
// 获取账户列表
accList, err := trd.GetAccList(cli)
if err != nil {
    log.Fatal(err)
}

// 使用第一个账户
acc := accList[0]

// 查询资金
funds, err := trd.GetFunds(cli, acc.AccID, int32(trdcommon.TrdMarket_TrdMarket_HK))
if err != nil {
    log.Fatal(err)
}

fmt.Printf("现金: %.2f, 冻结: %.2f\n", 
    funds.GetCash(), funds.GetFrozenCash())
```

### 查询持仓

```go
positions, err := trd.GetPositionList(cli, acc.AccID, 0, nil)
if err != nil {
    log.Fatal(err)
}

for _, pos := range positions.PositionList {
    fmt.Printf("%s: 数量=%d, 成本=%.2f, 当前=%.2f\n",
        pos.Security.GetCode(),
        pos.GetQty(),
        pos.GetCostPrice(),
        pos.GetMarketVal())
}
```

### 下单

```go
// 买入 100 股腾讯
orderID, err := trd.PlaceOrder(cli, &trd.PlaceOrderRequest{
    AccID:        acc.AccID,
    TrdSide:      int32(trdcommon.TrdSide_TrdSide_Buy),
    OrderType:    int32(trdcommon.OrderType_OrderType_Normal),
    Market:       int32(trdcommon.TrdMarket_TrdMarket_HK),
    Security:     &trdcommon.Security{Market: &market, Code: &code},
    Qty:          100,
    Price:        350.00,
    PriceType:    int32(trdcommon.PriceType_PriceType_Normal),
})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("订单号: %s\n", orderID)
```

### 修改订单

```go
err = trd.ModifyOrder(cli, &trd.ModifyOrderRequest{
    AccID:     acc.AccID,
    OrderID:   orderID,
    Market:    int32(trdcommon.TrdMarket_TrdMarket_HK),
    ModifyType: int32(trdcommon.ModifyOrderType_ModifyOrderType_Normal),
    Qty:       200, // 修改数量
    Price:     360.00, // 修改价格
})
if err != nil {
    log.Fatal(err)
}
```

### 查询订单列表

```go
orders, err := trd.GetOrderList(cli, acc.AccID, 0, nil)
if err != nil {
    log.Fatal(err)
}

for _, o := range orders.OrderList {
    fmt.Printf("订单 %s: 状态=%d, 数量=%d, 价格=%.2f\n",
        o.GetOrderID(), o.GetState(), o.GetQty(), o.GetPrice())
}
```

---

## 实时推送

### 订阅行情

```go
// 设置推送回调
cli.SetQotPushHandler(func(packet *conn.Packet) {
    switch packet.ProtoID {
    case qot.ProtoID_GetBasicQot:
        // 处理行情推送
    case qot.ProtoID_GetKL:
        // 处理 K 线推送
    }
})

// 订阅实时行情
security := &qotcommon.Security{Market: &market, Code: &code}
_, err = qot.Subscribe(cli, &qot.SubscribeRequest{
    SecurityList:     []*qotcommon.Security{security},
    SubTypeList:      []qot.SubType{qot.SubType_Basic, qot.SubType_KL},
    IsSubOrUnSub:     true,
    IsRegOrUnRegPush: true,
})
```

### 订单状态推送

```go
// 设置交易推送回调
cli.SetTrdPushHandler(func(packet *conn.Packet) {
    switch packet.ProtoID {
    case trd.ProtoID_UpdateOrder:
        // 处理订单更新
    case trd.ProtoID_UpdateOrderFill:
        // 处理成交更新
    }
})
```

---

## 常见问题

### Q: 连接失败怎么办？

1. 确认 Futu OpenD 已启动并正常运行
2. 确认端口号正确（默认 11111）
3. 确认网络连接正常

```go
err := cli.Connect("127.0.0.1:11111")
if err != nil {
    log.Fatal("连接失败:", err)
}
```

### Q: 如何处理错误？

所有 API 调用都可能返回错误，建议统一处理：

```go
result, err := qot.GetBasicQot(cli, securities)
if err != nil {
    // 区分错误类型
    if strings.Contains(err.Error(), "timeout") {
        // 处理超时
    } else if strings.Contains(err.Error(), "not connected") {
        // 处理断连
    } else {
        log.Fatal(err)
    }
}
```

### Q: 如何获取多个股票行情？

```go
securities := []*qotcommon.Security{
    {Market: &market, Code: &code1},
    {Market: &market, Code: &code2},
    {Market: &market, Code: &code3},
}

result, err := qot.GetBasicQot(cli, securities)
```

### Q: 如何设置价格提醒？

```go
// 获取价格提醒
result, err := qot.GetPriceReminder(cli, security, market)

// 设置提醒需要在 Futu OpenD 客户端中操作
```

### Q: 交易前需要什么准备？

1. 解锁交易密码：`trd.UnlockTrade()`
2. 获取交易账户：`trd.GetAccList()`
3. 确保账户有足够资金

---

## 市场常量参考

### 股票市场 (QotMarket)

| 市场 | 值 | 说明 |
|------|-----|------|
| HK_Security | 1 | 港股 |
| US_Security | 11 | 美股 |
| SH_Security | 31 | 沪股 |
| SZ_Security | 32 | 深股 |

### K 线类型 (KLType)

| 类型 | 值 | 说明 |
|------|-----|------|
| KLType_Min1 | 1 | 1 分钟 |
| KLType_Min5 | 2 | 5 分钟 |
| KLType_Min15 | 3 | 15 分钟 |
| KLType_Min30 | 4 | 30 分钟 |
| KLType_Min60 | 5 | 60 分钟 |
| KLType_Day | 4 | 日线 |
| KLType_Week | 5 | 周线 |
| KLType_Month | 6 | 月线 |

### 交易方向 (TrdSide)

| 方向 | 值 | 说明 |
|------|-----|------|
| Buy | 1 | 买入 |
| Sell | 2 | 卖出 |

### 订单状态 (OrderState)

| 状态 | 值 | 说明 |
|------|-----|------|
| Unknown | 0 | 未知 |
| Submitting | 1 | 提交中 |
| Submitted | 2 | 已提交 |
| Filled | 3 | 已成交 |
| PartiallyFilled | 4 | 部分成交 |
| Cancelled | 5 | 已取消 |
| Rejected | 6 | 已拒绝 |
