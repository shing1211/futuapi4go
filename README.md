# futuapi4go

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat-square&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  <img src="https://img.shields.io/badge/Platform-Futu%20OpenD-blue.svg" alt="Platform">
</p>

> 🚀 Go 语言实现的富途 OpenD API SDK —— 让 Golang 开发者也能轻松使用富途量化交易接口

## 特性

- ✅ **完整的市场数据接口** - 实时行情、K线、订单簿、逐笔成交、板块信息
- ✅ **交易功能支持** - 账户查询、订单管理、持仓查询、资金查询
- ✅ **WebSocket 推送** - 实时行情推送、订单状态推送
- ✅ **Protobuf 协议** - 高效的二进制序列化
- ✅ **连接池管理** - 自动重连、心跳保活
- ✅ **简洁的 API 设计** - 易于使用、类型安全

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
    
    "gitee.com/shing1211/futuapi4go/client"
    "gitee.com/shing1211/futuapi4go/qot"
)

func main() {
    // 创建客户端
    cli := client.New()
    
    // 连接到 OpenD (默认地址 127.0.0.1:11111)
    err := cli.Connect("127.0.0.1:11111")
    if err != nil {
        log.Fatal(err)
    }
    defer cli.Close()
    
    // 查询港股腾讯控股实时行情
    req := &qot.GetBasicQotRequest{
        Security: []*qot.Security{
            {Market: qot.QotMarket_HK, Code: "00700"},
        },
    }
    
    resp, err := cli.GetBasicQot(req)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, stock := range resp.S2C.GetBasicQotList() {
        fmt.Printf("%s %s: 最新价=%.2f, 涨跌=%.2f%%\n",
            stock.GetSecurity().GetCode(),
            stock.GetSecurity().GetName(),
            stock.GetLast(),
            stock.GetChangeRate()*100,
        )
    }
}
```

## 配置 OpenD

在使用本 SDK 前，请确保已完成以下配置：

1. 下载并安装 [富途 OpenD](https://www.futunn.com/download/openAPI)
2. 登录牛牛账号
3. 启用「行情接口」和「交易接口」
4. 记录 TCP 连接地址（默认 `127.0.0.1:11111`）

## API 文档

### 市场数据 (Qot)

| 方法 | 说明 |
|------|------|
| `GetBasicQot` | 获取实时行情 |
| `GetKL` | 获取实时/历史K线 |
| `GetHistoryKL` | 获取历史K线 |
| `GetOrderBook` | 获取订单簿(档口) |
| `GetTicker` | 获取逐笔成交 |
| `GetMarketSnapshot` | 获取市场快照 |
| `Subscribe` | 订阅实时行情推送 |
| `Unsubscribe` | 取消订阅 |

### 交易接口 (Trd)

| 方法 | 说明 |
|------|------|
| `GetAccList` | 获取账户列表 |
| `UnlockTrade` | 解锁交易密码 |
| `PlaceOrder` | 下单 |
| `ModifyOrder` | 修改订单 |
| `CancelOrder` | 撤销订单 |
| `GetOrderList` | 查询订单列表 |
| `GetPositionList` | 查询持仓列表 |
| `GetFunds` | 查询资金 |

## 项目结构

```
futuapi4go/
├── client/           # 核心客户端实现
│   ├── conn.go       # TCP连接与协议封装
│   └── client.go     # 主客户端
├── qot/              # 市场数据API
│   └── quote.go      # 行情查询接口
├── trd/              # 交易API
│   └── trade.go      # 交易接口
├── pb/               # Protobuf生成的Go代码
├── proto/            # Protobuf定义文件
├── examples/         # 使用示例
├── go.mod            # Go模块定义
└── README.md         # 本文件
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