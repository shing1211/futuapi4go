# futuapi4go - Go 语言实现的富途 OpenD API SDK

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  <img src="https://img.shields.io/badge/Platform-Futu%20OpenD-blue.svg" alt="Platform">
  <img src="https://img.shields.io/badge/Version-0.2.0-blue.svg" alt="Version">
</p>

> 🚀 Go 语言实现的富途 OpenD API SDK —— 为量化交易者打造的世界级 Golang 交易接口

## 实现进度总览

| 模块 | 状态 | APIs |
|------|------|------|
| 核心架构 | ✅ 完成 | 连接、心跳、用户信息 |
| 市场数据 (Qot) | ✅ 完成 | 42+ APIs |
| 交易接口 (Trd) | ✅ 完成 | 17 APIs |
| 系统与工具 | ✅ 完成 | 5 APIs |
| 推送通知 | ✅ 完成 | 11 handlers |
| 高级功能 | 🔄 进行中 | 自动重连、单元测试 |

📋 **完整实现状态**: [IMPLEMENTATION.md](IMPLEMENTATION.md)

---

## 特性

- ✅ 完整的市场数据接口 - 实时行情、K线、订单簿、逐笔成交、板块信息
- ✅ 交易功能支持 - 账户查询、订单管理、持仓查询、资金查询
- ✅ TCP/推送机制 - 实时行情推送、订单状态推送
- ✅ Protobuf 协议 - 高效的二进制序列化
- ✅ 连接管理 - 自动心跳保活
- ✅ 简洁的 API 设计 - 易于使用、类型安全

---

## 安装

```bash
go get gitee.com/shing1211/futuapi4go
```

## 环境要求

| 组件 | 版本要求 |
|------|----------|
| Golang | 1.21+ (推荐 1.26+) |
| Futu OpenD | 10.2.6208+ |

---

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
	// 1. 创建客户端
	cli := futuapi.New()

	// 2. 连接 OpenD
	err := cli.Connect("127.0.0.1:11111")
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 3. 查询行情
	market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	securities := []*qotcommon.Security{
		{Market: &market, Code: func() *string { s := "00700"; return &s }()},
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

## 项目结构

```
futuapi4go/
├── client/           # 核心客户端
│   ├── conn.go       # TCP连接与协议封装
│   ├── client.go     # 主客户端
│   └── errors.go     # 错误类型
├── qot/              # 市场数据API
│   ├── quote.go      # 行情查询
│   └── market.go     # 市场数据
├── trd/              # 交易API
│   └── trade.go      # 交易接口
├── sys/              # 系统API
│   └── system.go     # 系统接口
├── push/             # 推送处理
│   ├── qot_push.go   # 行情推送
│   └── trd_push.go   # 交易推送
├── pb/               # Protobuf生成代码
├── proto/            # Protobuf定义
└── examples/         # 使用示例
```

---

## 文档

| 文档 | 说明 |
|------|------|
| [IMPLEMENTATION.md](IMPLEMENTATION.md) | 详细实现计划与状态 |
| [USER_GUIDE.md](USER_GUIDE.md) | SDK 使用教程 |
| [DEVELOPER.md](DEVELOPER.md) | 开发者指南 |

---

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 致谢

- [富途](https://www.futunn.com/) 提供 OpenAPI
- [ftapi4go](https://github.com/futuopen/ftapi4go) 提供 Protobuf 定义
