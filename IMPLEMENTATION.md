# futuapi4go 实现计划与状态

本文档详细记录了 futuapi4go SDK 的 API 实现进度。

## 实现进度总览

| 阶段 | 状态 | APIs 数量 |
|------|------|----------|
| 阶段一：核心架构 | ✅ 完成 | 8 |
| 阶段二：市场数据 (Qot) | ✅ 完成 | 42 |
| 阶段三：交易接口 (Trd) | ✅ 完成 | 17 |
| 阶段四：系统与工具 | ✅ 完成 | 5 |
| 阶段五：高级功能 | 🔄 进行中 | - |

---

## 阶段一：核心架构 (Core Architecture) ✅

| 模块 | 状态 | 说明 |
|------|------|------|
| TCP 连接层 | ✅ 完成 | 自定义二进制协议封装 |
| InitConnect | ✅ 完成 | 连接初始化 (1001) |
| KeepAlive 心跳 | ✅ 完成 | 自动维持连接 (1002) |
| 全局状态 (GetGlobalState) | ✅ 完成 | 获取全局状态 (1004) |
| 用户信息 (GetUserInfo) | ✅ 完成 | 获取用户信息 (1005) |
| 延迟统计 (GetDelayStatistics) | ✅ 完成 | 获取延迟统计 (1006) |
| 错误处理 | ✅ 完成 | 统一的错误类型 |
| Protobuf 定义 | ✅ 完成 | v10.2.6208 (74 files) |

---

## 阶段二：市场数据 (Qot - Market Data) ✅

### 2.1 基础行情查询

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetBasicQot | 2101 | ✅ 完成 | 获取实时行情 |
| GetKL | 2102 | ✅ 完成 | 获取实时K线 |
| GetOrderBook | 2106 | ✅ 完成 | 获取订单簿(档口) |
| GetTicker | 2107 | ✅ 完成 | 获取逐笔成交 |
| GetRT | 2108 | ✅ 完成 | 获取实时分时数据 |
| GetSecuritySnapshot | 2110 | ✅ 完成 | 获取股票快照 |
| GetBroker | 2111 | ✅ 完成 | 获取买卖队列(经纪商) |

### 2.2 市场参考数据

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetStaticInfo | 2201 | ✅ 完成 | 获取股票静态信息 |
| GetPlateSet | 2202 | ✅ 完成 | 获取板块集合 |
| GetPlateSecurity | 2203 | ✅ 完成 | 获取板块成分股 |
| GetOwnerPlate | 2204 | ✅ 完成 | 获取所属板块 |
| GetReference | 2205 | ✅ 完成 | 获取正股相关数据 |
| GetTradeDate | 2206 | ✅ 完成 | 获取交易日 |
| RequestTradeDate | 2207 | ✅ 完成 | 请求交易日 |
| GetMarketState | 2208 | ✅ 完成 | 获取市场状态 |
| GetSuspend | 2209 | ✅ 完成 | 获取停牌信息 |
| GetCodeChange | 2210 | ✅ 完成 | 获取代码变更信息 |
| GetFutureInfo | 2211 | ✅ 完成 | 获取期货信息 |
| GetIpoList | 2212 | ✅ 完成 | 获取IPO列表 |
| GetHoldingChangeList | 2213 | ✅ 完成 | 获取持仓变化列表 |
| RequestRehab | 2214 | ✅ 完成 | 请求复权数据 |

### 2.3 高级数据

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetCapitalFlow | 2301 | ✅ 完成 | 获取资金流向 |
| GetCapitalDistribution | 2302 | ✅ 完成 | 获取资金分布 |
| StockFilter | 2303 | ✅ 完成 | 股票筛选 |
| GetOptionChain | 2304 | ✅ 完成 | 获取期权链 |
| GetOptionExpirationDate | 2305 | ✅ 完成 | 获取期权到期日 |
| GetWarrant | 2306 | ✅ 完成 | 获取窝轮信息 |

### 2.4 用户数据

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetUserSecurity | 2401 | ✅ 完成 | 获取用户自选股 |
| GetUserSecurityGroup | 2402 | ✅ 完成 | 获取用户自选股分组 |
| ModifyUserSecurity | 2403 | ✅ 完成 | 修改用户自选股 |
| GetPriceReminder | 2404 | ✅ 完成 | 获取价格提醒 |
| SetPriceReminder | 2405 | ✅ 完成 | 设置价格提醒 |

### 2.5 订阅与推送

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| Subscribe (Qot_Sub) | 3001 | ✅ 完成 | 订阅实时行情 |
| GetSubInfo | 3002 | ✅ 完成 | 获取订阅信息 |
| RegQotPush | 3003 | ✅ 完成 | 注册行情推送 |
| RequestHistoryKLQuota | 3104 | ✅ 完成 | 获取历史K线额度使用明细 |
| RequestHistoryKL | 2104 | ✅ 完成 | 请求历史K线(异步) |

### 2.6 推送通知 (Push Notifications)

| ProtoID | 状态 | 说明 |
|---------|------|------|
| Notify (1003) | ✅ 完成 | 系统通知推送 |
| Qot_UpdateBasicQot (3101) | ✅ 完成 | 实时行情推送 |
| Qot_UpdateKL (3102) | ✅ 完成 | K线推送 |
| Qot_UpdateOrderBook (3103) | ✅ 完成 | 订单簿推送 |
| Qot_UpdateTicker (3104) | ✅ 完成 | 逐笔成交推送 |
| Qot_UpdateRT (3105) | ✅ 完成 | 分时数据推送 |
| Qot_UpdateBroker (3106) | ✅ 完成 | 经纪商队列推送 |
| Qot_UpdatePriceReminder (3107) | ✅ 完成 | 价格提醒推送 |

---

## 阶段三：交易接口 (Trd - Trading) ✅

### 3.1 账户管理

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetAccList | 4001 | ✅ 完成 | 获取账户列表 |
| UnlockTrade | 4002 | ✅ 完成 | 解锁交易密码 |
| GetFunds | 4003 | ✅ 完成 | 获取资金信息 |
| GetOrderFee | 4004 | ✅ 完成 | 获取订单费用 |
| GetMarginRatio | 4005 | ✅ 完成 | 获取保证金比例 |
| GetMaxTrdQtys | 4006 | ✅ 完成 | 获取最大交易数量 |
| GetFlowSummary | 2226 | ✅ 完成 | 获取账户资金流水 |

### 3.2 订单管理

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| PlaceOrder | 5001 | ✅ 完成 | 下单 |
| ModifyOrder | 5002 | ✅ 完成 | 修改订单 |
| GetOrderList | 5003 | ✅ 完成 | 查询订单列表 |
| GetHistoryOrderList | 5004 | ✅ 完成 | 查询历史订单 |
| GetOrderFillList | 5005 | ✅ 完成 | 查询成交列表 |
| GetHistoryOrderFillList | 5006 | ✅ 完成 | 查询历史成交 |

### 3.3 持仓管理

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetPositionList | 6001 | ✅ 完成 | 获取持仓列表 |

### 3.4 交易推送

| ProtoID | 状态 | 说明 |
|---------|------|------|
| Trd_UpdateOrder (7001) | ✅ 完成 | 订单状态推送 |
| Trd_UpdateOrderFill (7002) | ✅ 完成 | 成交推送 |
| Trd_Notify (7003) | ✅ 完成 | 交易通知推送 |
| Trd_ReconfirmOrder (7004) | ✅ 完成 | 订单确认推送 |
| Trd_SubAccPush (7005) | ✅ 完成 | 账户推送订阅 |

---

## 阶段四：系统与工具 (System) ✅

| API | ProtoID | 状态 | 说明 |
|-----|---------|------|------|
| GetGlobalState | 1004 | ✅ 完成 | 获取全局状态 |
| GetUserInfo | 1005 | ✅ 完成 | 获取用户信息 |
| GetDelayStatistics | 1006 | ✅ 完成 | 获取延迟统计 |
| Verification | 8001 | ✅ 完成 | 验证接口 |

---

## 阶段五：高级功能 (Advanced Features) 🔄

| 功能 | 状态 | 说明 |
|------|------|------|
| 连接保活 (KeepAlive) | ✅ 完成 | 自动心跳维持连接 |
| 自动重连 | ⏳ 规划中 | 连接断开后自动重连 |
| 请求重试 | ⏳ 规划中 | 超时自动重试机制 |
| 并发控制 | ⏳ 规划中 | 请求并发限制 |
| 日志系统 | ⏳ 规划中 | 可配置的日志输出 |
| 连接池 | ⏳ 规划中 | 多连接管理 |
| 单元测试 | ⏳ 规划中 | 核心功能测试覆盖 |

---

## 已废弃 API

| API | ProtoID | 说明 |
|-----|---------|------|
| GetMarketSnapshot | 2109 | 被 GetSecuritySnapshot (2110) 替代 |

---

## 实现统计

| 类别 | 数量 |
|------|------|
| 总 API 数 | 74 proto files |
| 已实现 APIs | ~48 |
| 已实现 Push Handlers | 11 |
| 文档完整度 | 100% |

---

*最后更新: 2026-04-07*
