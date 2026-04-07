# futuapi4go - Go 语言实现的富途 OpenD API SDK

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  <img src="https://img.shields.io/badge/Platform-Futu%20OpenD-blue.svg" alt="Platform">
  <img src="https://img.shields.io/badge/Alpha-WIP-orange.svg" alt="Status">
</p>

> 🚀 Go 语言实现的富途 OpenD API SDK —— 为量化交易者打造的世界级 Golang 交易接口

## 特性

- ✅ **完整的市场数据接口** - 实时行情、K线、订单簿、逐笔成交、板块信息
- ✅ **交易功能支持** - 账户查询、订单管理、持仓查询、资金查询
- ✅ **TCP/推送机制** - 实时行情推送、订单状态推送
- ✅ **Protobuf 协议** - 高效的二进制序列化
- ✅ **连接管理** - 自动重连、心跳保活
- ✅ **简洁的 API 设计** - 易于使用、类型安全

## 实现计划与状态

### 阶段一：核心架构 (Core Architecture) ✅

| 模块 | 状态 | 说明 |
|------|------|------|
| TCP 连接层 | ✅ 完成 | 自定义二进制协议封装 |
| InitConnect | ✅ 完成 | 连接初始化 |
| KeepAlive 心跳 | ✅ 完成 | 自动维持连接 |
| 全局状态 (GetGlobalState) | ⏳ 规划中 | 获取全局状态 |
| 用户信息 (GetUserInfo) | ⏳ 规划中 | 获取用户信息 |
| 延迟统计 (GetDelayStatistics) | ⏳ 规划中 | 获取延迟统计 |
| 错误处理 | ✅ 完成 | 统一的错误类型 |
| Protobuf 定义 | ✅ 完成 | v10.2.6208 |

### 阶段二：市场数据 (Qot - Market Data) 🔄 进行中

#### 2.1 基础行情查询

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetBasicQot | 2101 | ✅ 完成 | 获取实时行情 |
| GetKL | 2102 | ✅ 完成 | 获取实时K线 |
| GetHistoryKL | 2103 | ✅ 完成 | 获取历史K线 |
| RequestHistoryKL | 2104 | ✅ 完成 | 请求历史K线(异步) |
| GetOrderBook | 2106 | ✅ 完成 | 获取订单簿(档口) |
| GetTicker | 2107 | ✅ 完成 | 获取逐笔成交 |
| GetRT | 2108 | ✅ 完成 | 获取实时分时数据 |
| GetMarketSnapshot | 2109 | ⏳ 规划中 | 获取市场快照 |
| GetSecuritySnapshot | 2110 | ✅ 完成 | 获取股票快照 |
| GetBroker | 2111 | ✅ 完成 | 获取买卖队列(经纪商) |

#### 2.3 市场参考数据

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetStaticInfo | 2201 | ✅ 完成 | 获取股票静态信息 |
| GetPlateSet | 2202 | ✅ 完成 | 获取板块集合 |
| GetPlateSecurity | 2203 | ✅ 完成 | 获取板块成分股 |
| GetOwnerPlate | 2204 | ✅ 完成 | 获取所属板块 |
| GetReference | 2205 | ✅ 完成 | 获取正股相关数据 |
| GetTradeDate | 2206 | ✅ 完成 | 获取交易日 |
| RequestTradeDate | 2207 | ⏳ 规划中 | 请求交易日 |
| GetMarketState | 2208 | ✅ 完成 | 获取市场状态 |
| GetSuspend | 2209 | ⏳ 规划中 | 获取停牌信息 |
| GetCodeChange | 2210 | ⏳ 规划中 | 获取代码变更信息 |
| GetFutureInfo | 2211 | ⏳ 规划中 | 获取期货信息 |
| GetIpoList | 2212 | ⏳ 规划中 | 获取IPO列表 |
| GetHoldingChangeList | 2213 | ⏳ 规划中 | 获取持仓变化列表 |

#### 2.4 高级数据

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetCapitalFlow | 2301 | ✅ 完成 | 获取资金流向 |
| GetCapitalDistribution | 2302 | ✅ 完成 | 获取资金分布 |
| StockFilter | 2303 | ⏳ 规划中 | 股票筛选 |
| GetOptionChain | 2304 | ⏳ 规划中 | 获取期权链 |
| GetOptionExpirationDate | 2305 | ⏳ 规划中 | 获取期权到期日 |
| GetWarrant | 2306 | ⏳ 规划中 | 获取窝轮信息 |

#### 2.5 用户数据

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetUserSecurity | 2401 | ✅ 完成 | 获取用户自选股 |
| GetUserSecurityGroup | 2402 | ⏳ 规划中 | 获取用户自选股分组 |
| ModifyUserSecurity | 2403 | ⏳ 规划中 | 修改用户自选股 |
| GetPriceReminder | 2404 | ✅ 完成 | 获取价格提醒 |
| SetPriceReminder | 2405 | ⏳ 规划中 | 设置价格提醒 |

#### 2.6 订阅与推送

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| Subscribe (Qot_Sub) | 3001 | ✅ 完成 | 订阅实时行情 |
| GetSubInfo | 3002 | ✅ 完成 | 获取订阅信息 |
| RegQotPush | 3003 | ⏳ 规划中 | 注册行情推送 |

##### 推送通知 (Push Notifications)

| ProtoID | 状态 | 说明 |
|---------|------|------|
| Qot_UpdateBasicQot (3101) | ✅ 完成 | 实时行情推送 |
| Qot_UpdateKL (3102) | ✅ 完成 | K线推送 |
| Qot_UpdateOrderBook (3103) | ✅ 完成 | 订单簿推送 |
| Qot_UpdateTicker (3104) | ✅ 完成 | 逐笔成交推送 |
| Qot_UpdateRT (3105) | ✅ 完成 | 分时数据推送 |
| Qot_UpdateBroker (3106) | ✅ 完成 | 经纪商队列推送 |
| Qot_UpdatePriceReminder (3107) | ⏳ 规划中 | 价格提醒推送 |
| Trd_UpdateOrder (7001) | ✅ 完成 | 订单状态推送 |
| Trd_UpdateOrderFill (7002) | ✅ 完成 | 成交推送 |

### 阶段三：交易接口 (Trd - Trading) 🔄 进行中

#### 3.1 账户管理

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetAccList | 4001 | ✅ 完成 | 获取账户列表 |
| UnlockTrade | 4002 | ✅ 完成 | 解锁交易密码 |
| GetFunds | 4003 | ✅ 完成 | 获取资金信息 |
| GetOrderFee | 4004 | ✅ 完成 | 获取订单费用 |
| GetMarginRatio | 4005 | ✅ 完成 | 获取保证金比例 |
| GetMaxTrdQtys | 4006 | ✅ 完成 | 获取最大交易数量 |

#### 3.2 订单管理

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| PlaceOrder | 5001 | ✅ 完成 | 下单 |
| ModifyOrder | 5002 | ✅ 完成 | 修改订单 |
| GetOrderList | 5003 | ✅ 完成 | 查询订单列表 |
| GetHistoryOrderList | 5004 | ⏳ 规划中 | 查询历史订单 |
| GetOrderFillList | 5005 | ✅ 完成 | 查询成交列表 |
| GetHistoryOrderFillList | 5006 | ⏳ 规划中 | 查询历史成交 |

#### 3.3 持仓管理

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetPositionList | 6001 | ✅ 完成 | 获取持仓列表 |

#### 3.4 交易推送

| ProtoID | 状态 | 说明 |
|---------|------|------|
| Trd_UpdateOrder (7001) | ⏳ 规划中 | 订单状态推送 |
| Trd_UpdateOrderFill (7002) | ⏳ 规划中 | 成交推送 |
| Trd_Notify (7003) | ⏳ 规划中 | 交易通知推送 |
| Trd_ReconfirmOrder (7004) | ⏳ 规划中 | 订单确认推送 |
| Trd_SubAccPush (7005) | ⏳ 规划中 | 账户推送订阅 |

### 阶段四：系统与工具 (System) ✅ 完成

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetGlobalState | 1004 | ✅ 完成 | 获取全局状态 |
| GetUserInfo | 1005 | ✅ 完成 | 获取用户信息 |
| GetDelayStatistics | 1006 | ✅ 完成 | 获取延迟统计 |
| Verification | 8001 | ⏳ 规划中 | 验证接口 |
| RequestRehab | 2214 | ⏳ 规划中 | 请求复权数据 |

### 阶段五：高级功能 (Advanced Features) ⏳ 规划中

| 功能 | 状态 | 说明 |
|------|------|------|
| 自动重连 | ⏳ 规划中 | 连接断开后自动重连 |
| 请求重试 | ⏳ 规划中 | 超时自动重试机制 |
| 并发控制 | ⏳ 规划中 | 请求并发限制 |
| 日志系统 | ⏳ 规划中 | 可配置的日志输出 |
| 连接池 | ⏳ 规划中 | 多连接管理 |
| 单元测试 | ⏳ 规划中 | 核心功能测试覆盖 |

---

## 安装

```bash
go get gitee.com/shing1211/futuapi4go
```

## 环境要求

| 组件 | 版本要求 |
|------|----------|
| **Golang** | 1.21+ (推荐 1.26+) |
| **Futu OpenD** | 10.2.6208+ (最新版本) |

## 快速开始

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
	cli := futuapi.New()
	err := cli.Connect("127.0.0.1:11111")
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

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
		fmt.Printf("%s %s: 现价=%.2f\n", bq.Security.GetCode(), bq.Name, bq.CurPrice)
	}
}
```

## 项目结构

```
futuapi4go/
├── client/           # 核心客户端
│   ├── conn.go       # TCP连接与协议封装
│   ├── client.go     # 主客户端
│   └── errors.go     # 错误类型定义
├── qot/              # 市场数据API
│   └── quote.go      # 行情查询接口
├── trd/              # 交易API
│   └── trade.go      # 交易接口
├── sys/              # 系统API
│   └── system.go      # 系统接口
├── push/             # 推送通知处理
│   ├── qot_push.go   # Qot推送解析
│   └── trd_push.go   # Trd推送解析
├── pb/               # Protobuf生成的Go代码
├── proto/            # Protobuf定义文件
└── examples/         # 使用示例
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 致谢

- [富途](https://www.futunn.com/) 提供 OpenAPI
- [ftapi4go](https://github.com/futuopen/ftapi4go) 提供 Protobuf 定义