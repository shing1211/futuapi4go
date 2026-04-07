# 更新日志

All notable changes to this project will be documented in this file.

## [0.3.0] - 2026-04-07

### Added

#### OpenD 模拟器 (Simulator)
- TCP 服务器核心 (46字节协议头, LittleEndian)
- 系统 API 处理器 (4): InitConnect, KeepAlive, GetGlobalState, GetUserInfo
- Qot 行情处理器 (42): 覆盖所有 Qot API
- Trd 交易处理器 (13): 覆盖所有交易 API
- 模拟器示例程序 (examples/simulator/main.go)

### Fixed

#### SDK Bug Fixes
- qot/quote.go: Subscribe - 添加缺失的 retType 错误检查
- qot/quote.go: ModifyUserSecurity - 添加缺失的 retType 错误检查
- qot/quote.go: RegQotPush - 添加缺失的 retType 错误检查

### Documentation
- 更新 IMPLEMENTATION.md 添加模拟器统计
- 更新 SIMULATOR.md 完整实现状态
- 更新 README.md 项目状态

## [0.2.0] - 2026-04-07

### Added

#### Qot - 市场数据 API (29 APIs)
- GetBasicQot (2101) - 获取实时行情
- GetKL (2102) - 获取实时K线
- GetOrderBook (2106) - 获取订单簿
- GetTicker (2107) - 获取逐笔成交
- GetRT (2108) - 获取实时分时数据
- GetSecuritySnapshot (2110) - 获取股票快照
- GetBroker (2111) - 获取买卖队列
- GetStaticInfo (2201) - 获取股票静态信息
- GetPlateSet (2202) - 获取板块集合
- GetPlateSecurity (2203) - 获取板块成分股
- GetOwnerPlate (2204) - 获取所属板块
- GetReference (2205) - 获取正股相关数据
- GetTradeDate (2206) - 获取交易日
- RequestTradeDate (2207) - 请求交易日
- GetMarketState (2208) - 获取市场状态
- GetSuspend (2209) - 获取停牌信息
- GetCodeChange (2210) - 获取代码变更信息
- GetFutureInfo (2211) - 获取期货信息
- GetIpoList (2212) - 获取IPO列表
- GetHoldingChangeList (2213) - 获取持仓变化列表
- RequestRehab (2214) - 请求复权数据
- GetCapitalFlow (2301) - 获取资金流向
- GetCapitalDistribution (2302) - 获取资金分布
- StockFilter (2303) - 股票筛选
- GetOptionChain (2304) - 获取期权链
- GetOptionExpirationDate (2305) - 获取期权到期日
- GetWarrant (2306) - 获取窝轮信息
- GetUserSecurity (2401) - 获取用户自选股
- GetUserSecurityGroup (2402) - 获取用户自选股分组
- ModifyUserSecurity (2403) - 修改用户自选股
- GetPriceReminder (2404) - 获取价格提醒
- SetPriceReminder (2405) - 设置价格提醒
- Subscribe (3001) - 订阅实时行情
- GetSubInfo (3002) - 获取订阅信息
- RegQotPush (3003) - 注册行情推送
- RequestHistoryKLQuota (3104) - 获取历史K线额度使用明细
- RequestHistoryKL (2104) - 请求历史K线(异步)

#### Qot - 推送通知 (7 handlers)
- Qot_UpdateBasicQot (3101) - 实时行情推送
- Qot_UpdateKL (3102) - K线推送
- Qot_UpdateOrderBook (3103) - 订单簿推送
- Qot_UpdateTicker (3104) - 逐笔成交推送
- Qot_UpdateRT (3105) - 分时数据推送
- Qot_UpdateBroker (3106) - 经纪商队列推送
- Qot_UpdatePriceReminder (3107) - 价格提醒推送

#### Trd - 交易 API (14 APIs)
- GetAccList (4001) - 获取账户列表
- UnlockTrade (4002) - 解锁交易密码
- GetFunds (4003) - 获取资金信息
- GetOrderFee (4004) - 获取订单费用
- GetMarginRatio (4005) - 获取保证金比例
- GetMaxTrdQtys (4006) - 获取最大交易数量
- PlaceOrder (5001) - 下单
- ModifyOrder (5002) - 修改订单
- GetOrderList (5003) - 查询订单列表
- GetHistoryOrderList (5004) - 查询历史订单
- GetOrderFillList (5005) - 查询成交列表
- GetHistoryOrderFillList (5006) - 查询历史成交
- GetPositionList (6001) - 获取持仓列表
- SubAccPush (7005) - 账户推送订阅
- ReconfirmOrder (7004) - 订单确认

#### Trd - 推送通知 (3 handlers)
- Trd_UpdateOrder (7001) - 订单状态推送
- Trd_UpdateOrderFill (7002) - 成交推送
- Trd_Notify (7003) - 交易通知推送

#### System - 系统 API (4 APIs)
- GetGlobalState (1004) - 获取全局状态
- GetUserInfo (1005) - 获取用户信息
- GetDelayStatistics (1006) - 获取延迟统计
- Verification (8001) - 验证接口

#### System - 推送通知 (1 handler)
- Notify (1003) - 系统通知推送

### Updated
- Protobuf 定义更新至 v10.2.6208 (74 proto files)
- README.md 添加详细的 API 实现状态表格

## [0.1.0] - 2026-04-07

### Added
- 初始版本发布
- 核心客户端实现 (TCP连接、协议封装)
- InitConnect 连接初始化
- 基本的 Protobuf 消息定义
- README、许可证等基础文件

### Planned Features
- 市场数据 API (Qot) - 实时行情、K线、订单簿
- 交易 API (Trd) - 账户、下单、持仓
- WebSocket 推送支持
- 完整的错误处理和重连机制
- 更多使用示例