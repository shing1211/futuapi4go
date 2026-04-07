# futuapi4go

<p align="center">
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Version-0.2.0-blue.svg" alt="Version">
  </a>
  <a href="https://gitee.com/shing1211/futuapi4go">
    <img src="https://img.shields.io/badge/Platform-Futu%20OpenD-blue.svg" alt="Platform">
  </a>
</p>

<p align="center">
  <strong>Go 语言实现的富途 OpenD API SDK</strong><br>
  为量化交易者打造的世界级 Golang 交易接口
</p>

---

## 项目状态

| 模块 | 状态 | 说明 |
|------|------|------|
| 核心架构 | ✅ 已完成 | TCP 连接、自动心跳、用户信息 |
| 市场数据 (Qot) | ✅ 已完成 | 42+ APIs 完整实现 |
| 交易接口 (Trd) | ✅ 已完成 | 账户、订单、持仓管理 |
| 推送通知 | ✅ 已完成 | 实时行情与交易推送 |
| 系统 API | ✅ 已完成 | 全局状态、验证接口 |
| 高级功能 | ✅ 已完成 | 自动重连、日志系统 |
| 测试工具 | 🔄 开发中 | OpenD 模拟器 |

📋 **完整实现清单**: [IMPLEMENTATION.md](IMPLEMENTATION.md)
🧪 **OpenD 模拟器**: [SIMULATOR.md](SIMULATOR.md)

---

## 功能特性

### 市场数据
- 实时行情 (GetBasicQot)
- K线数据 (GetKL, RequestHistoryKL)
- 订单簿 (GetOrderBook)
- 逐笔成交 (GetTicker)
- 分时数据 (GetRT)
- 板块信息 (GetPlateSet, GetPlateSecurity)
- 资金流向 (GetCapitalFlow, GetCapitalDistribution)
- 期权数据 (GetOptionChain, GetOptionExpirationDate)
- 自选股管理 (GetUserSecurity, ModifyUserSecurity)
- 价格提醒 (SetPriceReminder, GetPriceReminder)

### 交易功能
- 账户管理 (GetAccList, UnlockTrade)
- 资金查询 (GetFunds)
- 下单与改单 (PlaceOrder, ModifyOrder)
- 订单管理 (GetOrderList, GetHistoryOrderList)
- 成交记录 (GetOrderFillList, GetHistoryOrderFillList)
- 持仓查询 (GetPositionList)

### 推送服务
- 实时行情推送 (Qot_UpdateBasicQot)
- K线推送 (Qot_UpdateKL)
- 订单簿推送 (Qot_UpdateOrderBook)
- 逐笔成交推送 (Qot_UpdateTicker)
- 订单状态推送 (Trd_UpdateOrder)
- 成交推送 (Trd_UpdateOrderFill)
- 系统通知推送 (Notify)

---

## 快速开始

### 安装

```bash
go get gitee.com/shing1211/futuapi4go
```

### 环境要求

| 组件 | 版本 |
|------|------|
| Golang | 1.21+ |
| Futu OpenD | 10.2.6208+ |

### 基础示例

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
	// 创建客户端
	cli := futuapi.New()

	// 连接 OpenD
	err := cli.Connect("127.0.0.1:11111")
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 查询腾讯控股行情
	market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	securities := []*qotcommon.Security{
		{Market: &market, Code: func() *string { s := "00700"; return &s }()},
	}

	result, err := qot.GetBasicQot(cli, securities)
	if err != nil {
		log.Fatal(err)
	}

	for _, bq := range result {
		fmt.Printf("%s %s: 现价=%.2f 开盘=%.2f 最高=%.2f 最低=%.2f 成交量=%d\n",
			bq.Security.GetCode(), bq.Name, bq.CurPrice, bq.OpenPrice,
			bq.HighPrice, bq.LowPrice, bq.Volume)
	}
}
```

---

## 项目结构

```
futuapi4go/
├── client/           # 核心客户端
│   ├── conn.go       # TCP 连接与二进制协议
│   ├── client.go     # 主客户端与连接管理
│   └── errors.go     # 错误类型定义
├── qot/              # 市场数据 API
│   ├── quote.go      # 行情查询接口
│   └── market.go     # 市场数据接口
├── trd/              # 交易 API
│   └── trade.go      # 交易接口
├── sys/              # 系统 API
│   └── system.go     # 系统级接口
├── simulator/        # OpenD 模拟器 (开发中)
│   ├── server.go     # TCP 服务器
│   └── handlers.go   # API 处理器
├── push/             # 推送通知处理
├── pb/               # Protobuf 生成代码
├── proto/            # Protobuf 定义
└── examples/         # 使用示例
├── push/             # 推送通知处理
│   ├── qot_push.go   # 行情推送解析
│   └── trd_push.go   # 交易推送解析
├── pb/               # Protobuf 生成代码
├── proto/            # Protobuf 定义文件
└── examples/          # 使用示例
```

---

## 文档

| 文档 | 说明 |
|------|------|
| [IMPLEMENTATION.md](IMPLEMENTATION.md) | 详细实现清单 |
| [USER_GUIDE.md](USER_GUIDE.md) | 用户使用指南 |
| [DEVELOPER.md](DEVELOPER.md) | 开发者指南 |
| [CHANGELOG.md](CHANGELOG.md) | 更新日志 |

---

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 致谢

- [富途](https://www.futunn.com/) - 提供 OpenAPI
