# futuapi4go

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
- ✅ **WebSocket 推送** - 实时行情推送、订单状态推送
- ✅ **Protobuf 协议** - 高效的二进制序列化
- ✅ **连接池管理** - 自动重连、心跳保活
- ✅ **简洁的 API 设计** - 易于使用、类型安全

## 实现计划与状态

### 阶段一：核心架构 (Core Architecture) ✅

| 模块 | 状态 | 说明 |
|------|------|------|
| TCP 连接层 | ✅ 完成 | 自定义二进制协议封装 |
| InitConnect | ✅ 完成 | 连接初始化 |
| 心跳保活 (KeepAlive) | ✅ 完成 | 自动维持连接 |
| 错误处理 | ✅ 完成 | 统一的错误类型 |
| Protobuf 定义 | ✅ 完成 | v10.2.6208 |

### 阶段二：市场数据 (Qot - Market Data) 🔄 进行中

| API | 状态 | 说明 |
|-----|------|------|
| GetBasicQot | ✅ 完成 | 获取实时行情 |
| GetKL | ⏳ 规划中 | 获取实时/历史K线 |
| GetHistoryKL | ⏳ 规划中 | 获取历史K线 |
| GetOrderBook | ⏳ 规划中 | 获取订单簿(档口) |
| GetTicker | ⏳ 规划中 | 获取逐笔成交 |
| GetRT | ⏳ 规划中 | 获取实时分时数据 |
| GetMarketSnapshot | ⏳ 规划中 | 获取市场快照 |
| GetBroker | ⏳ 规划中 | 获取买卖队列 |
| Subscribe/Unsubscribe | ⏳ 规划中 | 订阅/取消订阅实时行情 |
| Subscription Push | ⏳ 规划中 | 实时行情推送处理 |

### 阶段三：交易接口 (Trd - Trading) ⏳ 规划中

| API | 状态 | 说明 |
|-----|------|------|
| GetAccList | ⏳ 规划中 | 获取账户列表 |
| UnlockTrade | ⏳ 规划中 | 解锁交易密码 |
| GetFunds | ⏳ 规划中 | 获取资金信息 |
| GetPositionList | ⏳ 规划中 | 获取持仓列表 |
| GetOrderList | ⏳ 规划中 | 获取订单列表 |
| GetOrderFillList | ⏳ 规划中 | 获取成交列表 |
| PlaceOrder | ⏳ 规划中 | 下单 |
| ModifyOrder | ⏳ 规划中 | 修改订单 |
| CancelOrder | ⏳ 规划中 | 撤销订单 |

### 阶段四：高级功能 (Advanced Features) ⏳ 规划中

| 功能 | 状态 | 说明 |
|------|------|------|
| 自动重连 | ⏳ 规划中 | 连接断开后自动重连 |
| 请求重试 | ⏳ 规划中 | 超时自动重试机制 |
| 并发控制 | ⏳ 规划中 | 请求并发限制 |
| 日志系统 | ⏳ 规划中 | 可配置的日志输出 |
| 性能优化 | ⏳ 规划中 | 连接池、批处理优化 |
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

> **注意**: Protobuf 定义文件基于 Futu OpenD v10.2.6208，请确保使用对应版本或更高版本的 OpenD 以获得最佳兼容性。

### 安装 OpenD

下载并安装 [富途 OpenD](https://www.futunn.com/download/openAPI)：

1. 登录牛牛账号
2. 启用「行情接口」和「交易接口」
3. 记录 TCP 连接地址（默认 `127.0.0.1:11111`）

## 快速开始

```go
package main

import (
	"fmt"
	"log"

	futuapi "gitee.com/shing1211/futuapi4go/client"
	"gitee.com/shing1211/futuapi4go/qot"
	"github.com/futuopen/ftapi4go/pb/qotcommon"
)

func main() {
	// 创建客户端
	cli := futuapi.New()

	// 连接到 OpenD (默认地址 127.0.0.1:11111)
	err := cli.Connect("127.0.0.1:11111")
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 查询港股腾讯控股实时行情
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
		fmt.Printf("%s %s: 现价=%.2f 开=%.2f 高=%.2f 低=%.2f\n",
			bq.Security.GetCode(),
			bq.Name,
			bq.CurPrice,
			bq.OpenPrice,
			bq.HighPrice,
			bq.LowPrice,
		)
	}
}
```

## 项目结构

```
futuapi4go/
├── client/           # 核心客户端实现
│   ├── conn.go      # TCP连接与协议封装
│   ├── client.go    # 主客户端(含心跳)
│   └── errors.go    # 错误类型定义
├── qot/             # 市场数据API
│   └── quote.go     # 行情查询接口
├── trd/             # 交易API
│   └── trade.go     # 交易接口
├── pb/              # Protobuf生成的Go代码
├── proto/           # Protobuf定义文件
├── examples/        # 使用示例
├── go.mod           # Go模块定义
└── README.md        # 本文件
```

## 测试

```bash
# 运行所有测试
go test ./...

# 运行特定包测试
go test ./client/...
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License - 查看 [LICENSE](LICENSE) 文件

## 致谢

- [富途](https://www.futunn.com/) 提供 OpenAPI
- [ftapi4go](https://github.com/futuopen/ftapi4go) 提供 Protobuf 定义