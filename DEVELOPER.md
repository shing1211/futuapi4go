# futuapi4go 开发者指南

本指南面向 SDK 开发者，介绍项目架构、代码规范和贡献流程。

## 目录

1. [项目架构](#项目架构)
2. [协议层实现](#协议层实现)
3. [API 实现模式](#api-实现模式)
4. [代码规范](#代码规范)
5. [protobuf 定义](#protobuf-定义)
6. [测试](#测试)
7. [调试](#调试)

---

## 项目架构

```
futuapi4go/
├── client/           # 核心客户端
│   ├── conn.go       # TCP 连接与二进制协议封装
│   ├── client.go     # 主客户端、连接管理
│   └── errors.go     # 错误类型定义
├── qot/              # 市场数据 API
│   └── quote.go      # 所有 Qot API 实现
├── trd/              # 交易 API
│   └── trade.go      # 所有 Trd API 实现
├── sys/              # 系统 API
│   └── system.go     # 系统级 API
├── push/             # 推送通知处理
│   ├── qot_push.go   # Qot 推送解析器
│   └── trd_push.go   # Trd 推送解析器
├── pb/               # Protobuf 生成的 Go 代码（本地模块）
├── proto/            # 原始 Protobuf 定义文件
└── examples/         # 使用示例
```

### 核心组件

#### 1. 连接层 (client/conn.go)

TCP 连接与自定义二进制协议封装：

- **协议头** (46 字节)：Magic "FT" + ProtoID + SerialNo + BodyLen
- **心跳机制**：自动发送 KeepAlive 保活
- **数据包读写**：WritePacket / ReadPacket

```go
// 连接结构
type Conn struct {
    conn   net.Conn
    protoID int32
    serialNo int32
    mu       sync.Mutex
}

// 写入数据包
func (c *Conn) WritePacket(protoID int32, serialNo int32, body []byte) error

// 读取数据包
func (c *Conn) ReadPacket() (*Packet, error)
```

#### 2. 客户端 (client/client.go)

高层 API，封装连接和序列化：

```go
type Client struct {
    conn *Conn
    serialNo int32
}

// 创建新客户端
func New() *Client

// 连接 OpenD
func (c *Client) Connect(addr string) error

// 关闭连接
func (c *Close() error

// 获取下一个序列号
func (c *Client) NextSerialNo() int32
```

---

## 协议层实现

### Futu OpenD 协议格式

每个请求/响应包包含：

| 字段 | 长度 | 说明 |
|------|------|------|
| Magic | 2 字节 | "FT" 固定值 |
| ProtoID | 4 字节 | 协议编号 (BigEndian) |
| SerialNo | 4 字节 | 序列号 (BigEndian) |
| BodyLen | 4 字节 | 包体长度 (BigEndian) |
| Body | 变长 | Protobuf 序列化数据 |

### 协议编码实现

```go
// conn.go 中的 WritePacket
func (c *Conn) WritePacket(protoID int32, serialNo int32, body []byte) error {
    // 1. 组装协议头
    header := make([]byte, 14)
    binary.BigEndian.PutUint16(header[0:2], 0x4654)  // "FT"
    binary.BigEndian.PutUint32(header[2:6], uint32(protoID))
    binary.BigEndian.PutUint32(header[6:10], uint32(serialNo))
    binary.BigEndian.PutUint32(header[10:14], uint32(len(body)))
    
    // 2. 发送头 + 包体
    _, err := c.conn.Write(append(header, body...))
    return err
}
```

---

## API 实现模式

所有 API 遵循统一模式，以 `GetBasicQot` 为例：

### 1. 定义请求/响应结构体

```go
// 请求结构体 - 对应 Protobuf C2S
type GetKLRequest struct {
    Security  *qotcommon.Security
    RehabType int32
    KLType    int32
    ReqNum    int32
}

// 响应结构体 - 自定义，方便使用
type GetKLResponse struct {
    Security *qotcommon.Security
    Name     string
    KLList   []*KLine
}
```

### 2. 实现 API 函数

```go
func GetKL(c *futuapi.Client, req *GetKLRequest) (*GetKLResponse, error) {
    // 1. 构建 Protobuf 请求
    c2s := &qotgetkl.C2S{
        Security:  req.Security,
        RehabType: &req.RehabType,
        KlType:    &req.KLType,
        ReqNum:    &req.ReqNum,
    }
    pkt := &qotgetkl.Request{C2S: c2s}

    // 2. 序列化
    body, err := proto.Marshal(pkt)
    if err != nil {
        return nil, err
    }

    // 3. 发送请求
    serialNo := c.NextSerialNo()
    if err := c.Conn().WritePacket(ProtoID_GetKL, serialNo, body); err != nil {
        return nil, err
    }

    // 4. 读取响应
    pktResp, err := c.Conn().ReadPacket()
    if err != nil {
        return nil, err
    }

    // 5. 反序列化
    var rsp qotgetkl.Response
    if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
        return nil, err
    }

    // 6. 检查返回码
    if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
        return nil, fmt.Errorf("GetKL failed: retType=%d, retMsg=%s", 
            rsp.GetRetType(), rsp.GetRetMsg())
    }

    // 7. 转换结果
    s2c := rsp.GetS2C()
    if s2c == nil {
        return nil, fmt.Errorf("GetKL: s2c is nil")
    }

    result := &GetKLResponse{
        Security: s2c.GetSecurity(),
        Name:     s2c.GetName(),
        KLList:   make([]*KLine, 0, len(s2c.GetKlList())),
    }

    for _, kl := range s2c.GetKlList() {
        result.KLList = append(result.KLList, &KLine{
            Time:       kl.GetTime(),
            ClosePrice: kl.GetClosePrice(),
            // ... 其他字段
        })
    }

    return result, nil
}
```

### 3. 关键注意事项

- **指针字段**：Protobuf 生成的结构体字段多为指针，必须赋值地址
- **命名差异**：有时 Protobuf getter 方法名与字段名不同（如 `GetOrederCount` 而非 `GetOrderCount`）
- **必填 vs 可选**：必填字段必须非 nil，可选字段可传 nil
- **错误处理**：始终检查 RetType 是否为成功值

---

## 代码规范

### Protobuf 字段处理

```go
// ✅ 正确：使用指针赋值
c2s := &qotgetkl.C2S{
    Security:  req.Security,
    RehabType: &req.RehabType,  // int32 需要取地址
    KlType:    &req.KLType,
    ReqNum:    &req.ReqNum,
}

// ❌ 错误：直接传值
c2s := &qotgetkl.C2S{
    Security:  req.Security,
    RehabType: req.RehabType,  // 编译错误
}
```

### 字符串字段

```go
// ✅ 正确
beginTime := "2024-01-01"
c2s := &xxx.C2S{
    BeginTime: &beginTime,
}

// ❌ 错误
c2s := &xxx.C2S{
    BeginTime: "2024-01-01",  // 编译错误
}
```

### 切片字段

```go
// ✅ 正确
typeList := []int32{1, 2, 3}
c2s := &xxx.C2S{
    TypeList: typeList,
}

// ❌ 错误
c2s := &xxx.C2S{
    TypeList: []int32{1, 2, 3},  // 编译错误
}
```

### Protobuf 导入

所有 protobuf 包使用本地路径：

```go
import (
    "google.golang.org/protobuf/proto"
    futuapi "gitee.com/shing1211/futuapi4go/client"
    "gitee.com/shing1211/futuapi4go/pb/common"
    "gitee.com/shing1211/futuapi4go/pb/qotcommon"
    "gitee.com/shing1211/futuapi4go/pb/qotgetkl"
)
```

---

## Protobuf 定义

### 文件位置

- **源定义**：`proto/` 目录
- **生成代码**：`pb/` 目录（本地 Go 模块）

### 添加新 Protobuf

1. 如果使用新的 proto 文件，先用 protoc 生成 Go 代码：

```bash
cd proto
protoc --go_out=../pb --go_opt=paths=source_relative \
    -I. \
    Qot_GetKL.proto Qot_Common.proto Common.proto
```

2. 更新 `pb/go.mod` 确保模块名正确：

```go
module gitee.com/shing1211/futuapi4go/pb
```

### ProtoID 常量定义

每个 API 在 quote.go 中定义常量：

```go
const (
    ProtoID_GetBasicQot = 2101
    ProtoID_GetKL       = 2102
    // ...
)
```

协议号参考 Futu OpenD API 文档。

---

## 测试

### 编译测试

```bash
# 编译所有包
go build ./...

# 运行测试
go test ./...

# 静态分析
go vet ./...
```

### 集成测试

需要 Futu OpenD 运行在本地：

```bash
# 运行示例
go run examples/main.go
```

### 常见编译错误

#### 1. protobuf 导入错误

```
could not import google.golang.org/protobuf/reflect/protoreflect
```

解决：确保 `go.mod` 包含正确的依赖：

```go
require (
    google.golang.org/protobuf v1.32.0
)
```

#### 2. LSP 报错但编译通过

某些 LSP（如 gopls）可能无法解析本地 `replace` 指令。只要 `go build` 通过即可。

---

## 调试

### 启用连接调试

```go
cli := futuapi.New()
cli.SetDebug(true)
```

### 查看协议包

在 `conn.go` 的 WritePacket/ReadPacket 中添加日志：

```go
func (c *Conn) WritePacket(protoID int32, serialNo int32, body []byte) error {
    log.Printf(">>> WritePacket: protoID=%d, serialNo=%d, len=%d", 
        protoID, serialNo, len(body))
    // ...
}
```

### 常见问题排查

#### 1. 连接被拒绝

- 确认 Futu OpenD 已启动
- 确认端口号正确（默认 11111）
- 确认防火墙允许连接

#### 2. 超时错误

- 检查网络连接
- 增加超时时间设置

#### 3. Protobuf 反序列化失败

- 检查 proto 版本是否与 OpenD 一致
- 确认包体完整（检查 BodyLen）

#### 4. 返回错误码

```go
// 常见 RetType 值
RetType_Succeed      = 0   // 成功
RetType_CommonFail   = -1   // 普通失败
RetType_SystemFail   = -2   // 系统错误
RetType_NoAuth       = -3   // 未授权
```

---

## 贡献流程

### 1. 实现新 API

1. 在 `proto/` 确认 proto 定义存在
2. 在 `pb/` 确认 Go 代码已生成
3. 在 `qot/quote.go` 添加：
   - import 语句
   - ProtoID 常量
   - 请求/响应结构体
   - API 函数实现
4. 编译验证：`go build ./...`
5. 提交代码

### 2. 更新文档

- API 状态更新到 `README.md`
- 新 API 添加使用示例到 `USER_GUIDE.md`
- 重大架构变更更新 `DEVELOPER.md`

### 3. Commit 规范

```
<模块>: <简短描述>

详细说明（可选）

解决的问题: Fixes #xxx
```

示例：

```
qot: implement GetWarrant (2306)

Add support for querying warrant data with comprehensive
filter options including maturity, price, premium, etc.
```

---

## 模块维护

### pb 模块管理

`pb/` 是独立的 Go 模块：

```bash
cd pb
go mod tidy
go build ./...
```

主项目通过 `replace` 指令使用本地 pb：

```go
// go.mod
require (
    gitee.com/shing1211/futuapi4go/pb v0.0.0
)

replace gitee.com/shing1211/futuapi4go/pb => ./pb
```

### 依赖更新

```bash
# 主项目
go get -u gitee.com/shing1211/futuapi4go/pb

# pb 模块
cd pb && go get -u google.golang.org/protobuf
```
