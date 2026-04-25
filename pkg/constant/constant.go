// Package constant provides enums and constants compatible with the Python futu-api SDK.
//
// This package mirrors the Python SDK's constant module for easy migration from Python to Go.
// The naming conventions follow the Python SDK closely.
//
// For Python developers, this replaces:
//
//	import futu as ft
//	ft.Market.HK          -> constant.Market_HK
//	ft.SecurityType.STOCK  -> constant.SecurityType_STOCK
//	ft.KLType.K_DAY      -> constant.KLType_K_DAY
//	ft.TrdEnv.SIMULATE    -> constant.TrdEnv_Simulate
//	ft.TrdSide.BUY        -> constant.TrdSide_Buy
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/constant"
//
//	// Market (int32 value)
//	market := constant.Market_HK  // 1
//
//	// Security Type
//	secType := constant.SecurityType_STOCK  // 3
//
//	// K-line Type
//	klType := constant.KLType_K_Day  // 6
//
//	// Rehab Type (AuType in Python)
//	rehabType := constant.RehabType_Forward  // 1 (QFQ)
//
//	// Subscription Type
//	subType := constant.SubType_Quote  // 1
//
//	// Trading Environment
//	trdEnv := constant.TrdEnv_Real  // 1
//
// # Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package constant

// =============================================================================
// Protocol IDs (ProtoId)
// =============================================================================

const (
	ProtoID_InitConnect        = 1001 // 初始化连接
	ProtoID_GetGlobalState     = 1002 // 获取全局状态
	ProtoID_Notify             = 1003 // 通知推送
	ProtoID_KeepAlive          = 1004 // 心跳保活
	ProtoID_GetUserInfo        = 1005 // 获取用户信息
	ProtoID_Verification       = 1006 // 请求或输入验证码
	ProtoID_GetDelayStatistics = 1007 // 获取延迟统计
	ProtoID_TestCmd            = 1008
	ProtoID_InitQuantMode      = 1009

	// Trading APIs
	ProtoID_Trd_GetAccList              = 2001 // 获取业务账户列表
	ProtoID_Trd_UnlockTrade             = 2005 // 解锁或锁定交易
	ProtoID_Trd_SubAccPush              = 2008 // 订阅业务账户的交易推送数据
	ProtoID_Trd_GetFunds                = 2101 // 获取账户资金
	ProtoID_Trd_GetPositionList         = 2102 // 获取账户持仓
	ProtoID_Trd_GetMaxTrdQtys           = 2111 // 查询最大买卖数量
	ProtoID_Trd_GetOrderList            = 2201 // 获取订单列表
	ProtoID_Trd_PlaceOrder              = 2202 // 下单
	ProtoID_Trd_ModifyOrder             = 2205 // 修改订单
	ProtoID_Trd_UpdateOrder             = 2208 // 订单状态变动通知(推送)
	ProtoID_Trd_GetOrderFillList        = 2211 // 获取成交列表
	ProtoID_Trd_UpdateOrderFill         = 2218 // 成交通知(推送)
	ProtoID_Trd_GetHistoryOrderList     = 2221 // 获取历史订单列表
	ProtoID_Trd_GetHistoryOrderFillList = 2222 // 获取历史成交列表
	ProtoID_Trd_GetMarginRatio          = 2223 // 获取融资融券数据
	ProtoID_Trd_GetOrderFee             = 2225 // 获取订单费用
	ProtoID_Trd_FlowSummary             = 2226 // 获取现金流水

	// Qot (Quote) APIs
	ProtoID_Qot_Sub                 = 3001 // 订阅或者反订阅
	ProtoID_Qot_RegQotPush          = 3002 // 注册推送
	ProtoID_Qot_GetSubInfo          = 3003 // 获取订阅信息
	ProtoID_Qot_GetBasicQot         = 3004 // 获取股票基本行情
	ProtoID_Qot_UpdateBasicQot      = 3005 // 推送股票基本行情
	ProtoID_Qot_GetKL               = 3006 // 获取K线
	ProtoID_Qot_UpdateKL            = 3007 // 推送K线
	ProtoID_Qot_GetRT               = 3008 // 获取分时
	ProtoID_Qot_UpdateRT            = 3009 // 推送分时
	ProtoID_Qot_GetTicker           = 3010 // 获取逐笔
	ProtoID_Qot_UpdateTicker        = 3011 // 推送逐笔
	ProtoID_Qot_GetOrderBook        = 3012 // 获取买卖盘
	ProtoID_Qot_UpdateOrderBook     = 3013 // 推送买卖盘
	ProtoID_Qot_GetBroker           = 3014 // 获取经纪队列
	ProtoID_Qot_UpdateBroker        = 3015 // 推送经纪队列
	ProtoID_Qot_UpdatePriceReminder = 3019 // 到价提醒通知

	// Historical Data
	ProtoID_Qot_RequestHistoryKL      = 3103 // 拉取历史K线
	ProtoID_Qot_RequestHistoryKLQuota = 3104 // 拉取历史K线已经用掉的额度
	ProtoID_Qot_RequestRehab          = 3105 // 获取除权信息

	// Other Qot APIs
	ProtoID_Qot_GetSuspend              = 3201 // 获取股票停牌信息
	ProtoID_Qot_GetStaticInfo           = 3202 // 获取股票列表
	ProtoID_Qot_GetSecuritySnapshot     = 3203 // 获取股票快照
	ProtoID_Qot_GetPlateSet             = 3204 // 获取板块集合下的板块
	ProtoID_Qot_GetPlateSecurity        = 3205 // 获取板块下的股票
	ProtoID_Qot_GetReference            = 3206 // 获取正股相关股票，暂时只有窝轮
	ProtoID_Qot_GetOwnerPlate           = 3207 // 获取股票所属板块
	ProtoID_Qot_GetHoldingChangeList    = 3208 // 获取高管持股变动
	ProtoID_Qot_GetOptionChain          = 3209 // 获取期权链
	ProtoID_Qot_GetWarrant              = 3210 // 拉取窝轮信息
	ProtoID_Qot_GetCapitalFlow          = 3211 // 获取资金流向
	ProtoID_Qot_GetCapitalDistribution  = 3212 // 获取资金分布
	ProtoID_Qot_GetUserSecurity         = 3213 // 获取自选股分组下的股票
	ProtoID_Qot_ModifyUserSecurity      = 3214 // 修改自选股分组下的股票
	ProtoID_Qot_StockFilter             = 3215 // 条件选股
	ProtoID_Qot_GetCodeChange           = 3216 // 代码变换
	ProtoID_Qot_GetIpoList              = 3217 // 获取新股Ipo
	ProtoID_Qot_GetFutureInfo           = 3218 // 获取期货资料
	ProtoID_Qot_RequestTradeDate        = 3219 // 在线拉取交易日
	ProtoID_Qot_SetPriceReminder        = 3220 // 设置到价提醒
	ProtoID_Qot_GetPriceReminder        = 3221 // 获取到价提醒
	ProtoID_Qot_GetUserSecurityGroup    = 3222 // 获取自选股分组
	ProtoID_Qot_GetMarketState          = 3223 // 获取指定品种的市场状态
	ProtoID_Qot_GetOptionExpirationDate = 3224 // 获取期权到期日
)

// AllPushIDs returns all push notification ProtoIDs
var AllPushIDs = []int32{
	ProtoID_Notify,
	ProtoID_Trd_UpdateOrder,
	ProtoID_Trd_UpdateOrderFill,
	ProtoID_Qot_UpdateBroker,
	ProtoID_Qot_UpdateOrderBook,
	ProtoID_Qot_UpdateKL,
	ProtoID_Qot_UpdateRT,
	ProtoID_Qot_UpdateBasicQot,
	ProtoID_Qot_UpdateTicker,
	ProtoID_Qot_UpdatePriceReminder,
}

// IsPushID returns true if the ProtoID is a push notification
func IsPushID(p int32) bool {
	for _, pushID := range AllPushIDs {
		if p == pushID {
			return true
		}
	}
	return false
}

// =============================================================================
// Market (行情市场)
// =============================================================================

// Market represents the market for quotes (行情市场)
type Market int32

const (
	Market_None = 0  // 未知市场
	Market_HK   = 1  // 香港市场
	Market_US   = 11 // 美国市场
	Market_SH   = 21 // 沪市
	Market_SZ   = 22 // 深市
	Market_SG   = 31 // 新加坡市场
	Market_JP   = 41 // 日本市场
	Market_AU   = 51 // 澳大利亚市场
	Market_MY   = 61 // 马来西亚市场
	Market_CA   = 71 // 加拿大市场
	Market_FX   = 81 // 外汇市场
)

// MarketToTrdSecMarket maps QotMarket (int32) to TrdSecMarket.
var MarketToTrdSecMarket = map[int32]TrdSecMarket{
	Market_None: TrdSecMarket_Unknown,
	Market_SH:   TrdSecMarket_CN_SH,
	Market_SZ:   TrdSecMarket_CN_SZ,
	Market_HK:   TrdSecMarket_HK,
	Market_US:   TrdSecMarket_US,
	Market_SG:   TrdSecMarket_SG,
	Market_JP:   TrdSecMarket_JP,
	Market_AU:   TrdSecMarket_AU,
	Market_MY:   TrdSecMarket_MY,
	Market_CA:   TrdSecMarket_CA,
	Market_FX:   TrdSecMarket_FX,
}

// =============================================================================
// SecurityType (证券类型)
// =============================================================================

// SecurityType represents the type of security
type SecurityType int32

const (
	SecurityType_None     SecurityType = 0  // 未知
	SecurityType_Bond     SecurityType = 1  // 场内债券
	SecurityType_Bwrt     SecurityType = 2  // 一揽子权证
	SecurityType_Stock    SecurityType = 3  // 正股
	SecurityType_ETF      SecurityType = 4  // 信托/ETF
	SecurityType_Warrant  SecurityType = 5  // 窝轮
	SecurityType_Index    SecurityType = 6  // 指数
	SecurityType_Plate    SecurityType = 7  // 板块
	SecurityType_Drvt     SecurityType = 8  // 期权
	SecurityType_PlateSet SecurityType = 9  // 板块集
	SecurityType_Future   SecurityType = 10 // 期货
	SecurityType_Forex    SecurityType = 11 // 外汇
)

// =============================================================================
// SubType (订阅类型)
// =============================================================================

// SubType represents the type of subscription (real-time data)
type SubType int32

const (
	SubType_None      SubType = 0  // 无
	SubType_Quote     SubType = 1  // 报价
	SubType_OrderBook SubType = 2  // 买卖摆盘
	SubType_Ticker    SubType = 4  // 逐笔
	SubType_Broker    SubType = 14 // 买卖经纪
	SubType_RT        SubType = 5  // 分时
	SubType_K_1Min    SubType = 11 // 1分钟K线
	SubType_K_3Min    SubType = 17 // 3分钟K线
	SubType_K_5Min    SubType = 7  // 5分钟K线
	SubType_K_15Min   SubType = 8  // 15分钟K线
	SubType_K_30Min   SubType = 9  // 30分钟K线
	SubType_K_60Min   SubType = 10 // 60分钟K线
	SubType_K_Day     SubType = 6  // 日K线
	SubType_K_Week    SubType = 12 // 周K线
	SubType_K_Month   SubType = 13 // 月K线
	SubType_K_Quarter SubType = 15 // 季度K线
	SubType_K_Year    SubType = 16 // 年K线
)

// IsKLType returns true if the SubType is a K-line type
func (s SubType) IsKLType() bool {
	return s == SubType_K_1Min || s == SubType_K_3Min || s == SubType_K_5Min ||
		s == SubType_K_15Min || s == SubType_K_30Min || s == SubType_K_60Min ||
		s == SubType_K_Day || s == SubType_K_Week || s == SubType_K_Month ||
		s == SubType_K_Quarter || s == SubType_K_Year
}

// =============================================================================
// KLType (K线类型)
// =============================================================================

// KLType represents the type of K-line (candlestick)
type KLType int32

const (
	KLType_None      KLType = 0  // 未知
	KLType_K_1Min    KLType = 1  // 1分钟
	KLType_K_Day     KLType = 2  // 日K
	KLType_K_Week    KLType = 3  // 周K
	KLType_K_Month   KLType = 4  // 月K
	KLType_K_Year    KLType = 5  // 年K
	KLType_K_5Min    KLType = 6  // 5分钟
	KLType_K_15Min   KLType = 7  // 15分钟
	KLType_K_30Min   KLType = 8  // 30分钟
	KLType_K_60Min   KLType = 9  // 60分钟
	KLType_K_3Min    KLType = 10 // 3分钟
	KLType_K_Quarter KLType = 11 // 季度K
)

// =============================================================================
// RehabType (复权类型)
// =============================================================================

// RehabType represents the type of price rehabilitation (复权)
type RehabType int32

const (
	RehabType_None     RehabType = 0 // 不复权
	RehabType_Forward  RehabType = 1 // 前复权
	RehabType_Backward RehabType = 2 // 后复权
)

// =============================================================================
// PlateSetType (板块集合类型)
// =============================================================================

// PlateSetType represents the type of plate set
type PlateSetType int32

const (
	PlateSetType_All      PlateSetType = 0 // 所有板块
	PlateSetType_Industry PlateSetType = 1 // 行业板块
	PlateSetType_Region   PlateSetType = 2 // 地域板块
	PlateSetType_Concept  PlateSetType = 3 // 概念板块
	PlateSetType_Other    PlateSetType = 4 // 其他板块
)

// =============================================================================
// TickerDirection (逐笔方向)
// =============================================================================

// TickerDirection represents the direction of a ticker (trade)
type TickerDirection int32

const (
	TickerDirection_None    TickerDirection = 0 // 未知
	TickerDirection_Buy     TickerDirection = 1 // 买
	TickerDirection_Sell    TickerDirection = 2 // 卖
	TickerDirection_Neutral TickerDirection = 3 // 中性
)

// =============================================================================
// HolderCategory (持有者类别)
// =============================================================================

// HolderCategory represents the category of stock holder
type HolderCategory int32

const (
	HolderCategory_None      HolderCategory = 0 // 未知
	HolderCategory_Agency    HolderCategory = 1 // 机构
	HolderCategory_Fund      HolderCategory = 2 // 基金
	HolderCategory_SeniorMgr HolderCategory = 3 // 高管
)

// =============================================================================
// OptionType (期权类型)
// =============================================================================

// OptionType represents the type of option
type OptionType int32

const (
	OptionType_None OptionType = 0 // 未知
	OptionType_Call OptionType = 1 // 涨/认购
	OptionType_Put  OptionType = 2 // 跌/认沽
)

// =============================================================================
// OptionCondType (价内价外)
// =============================================================================

// OptionCondType represents the condition type for options
type OptionCondType int32

const (
	OptionCondType_None    OptionCondType = 0 // 全部
	OptionCondType_WithIn  OptionCondType = 1 // 价内
	OptionCondType_Outside OptionCondType = 2 // 价外
)

// =============================================================================
// MarketState (市场状态)
// =============================================================================

// MarketState represents the state of the market
type MarketState int32

const (
	MarketState_None             MarketState = 0  // 无交易,美股未开盘
	MarketState_Auction          MarketState = 1  // 竞价
	MarketState_WaitingOpen      MarketState = 2  // 早盘前等待开盘
	MarketState_Morning          MarketState = 3  // 早盘
	MarketState_Rest             MarketState = 4  // 午间休市
	MarketState_Afternoon        MarketState = 5  // 午盘
	MarketState_Closed           MarketState = 6  // 收盘
	MarketState_PreMarketBegin   MarketState = 7  // 盘前开始
	MarketState_PreMarketEnd     MarketState = 8  // 盘前结束
	MarketState_AfterHoursBegin  MarketState = 9  // 盘后开始
	MarketState_AfterHoursEnd    MarketState = 10 // 盘后结束
	MarketState_NightOpen        MarketState = 11 // 夜市开盘
	MarketState_NightEnd         MarketState = 12 // 夜市收盘
	MarketState_FutureDayOpen    MarketState = 13 // 期指日市开盘
	MarketState_FutureDayBreak   MarketState = 14 // 期指日市休市
	MarketState_FutureDayClose   MarketState = 15 // 期指日市收盘
	MarketState_FutureDayWait    MarketState = 16 // 期指日市等待开盘
	MarketState_HK_CAS           MarketState = 17 // 港股盘后竞价
	MarketState_FutureNightWait  MarketState = 18 // 夜市等待开盘
	MarketState_FutureAfternoon  MarketState = 19 // 期货下午开盘
	MarketState_FutureSwitchDate MarketState = 20 // 期货切交易日
	MarketState_FutureOpen       MarketState = 21 // 期货开盘
	MarketState_FutureBreak      MarketState = 22 // 期货中盘休息
	MarketState_FutureBreakOver  MarketState = 23 // 期货休息后开盘
	MarketState_FutureClose      MarketState = 24 // 期货收盘
)

// =============================================================================
// TrdEnv (交易环境)
// =============================================================================

// TrdEnv represents the trading environment
type TrdEnv int32

const (
	TrdEnv_Simulate TrdEnv = 0 // 仿真环境(模拟环境)
	TrdEnv_Real     TrdEnv = 1 // 真实环境
)

// =============================================================================
// TrdMarket (交易市场)
// =============================================================================

// TrdMarket represents the trading market
type TrdMarket int32

const (
	TrdMarket_None              TrdMarket = 0   // 未知市场
	TrdMarket_HK                TrdMarket = 1   // 香港市场
	TrdMarket_US                TrdMarket = 2   // 美国市场
	TrdMarket_CN                TrdMarket = 3   // 大陆市场
	TrdMarket_HKCC              TrdMarket = 4   // 香港A股通市场
	TrdMarket_Futures           TrdMarket = 5   // 期货市场
	TrdMarket_SG                TrdMarket = 6   // 新加坡市场
	TrdMarket_AU                TrdMarket = 8   // 澳洲市场
	TrdMarket_JP                TrdMarket = 15  // 日本市场
	TrdMarket_MY                TrdMarket = 111 // 马来西亚市场
	TrdMarket_CA                TrdMarket = 112 // 加拿大市场
	TrdMarket_FuturesSimulateHK TrdMarket = 10  // 模拟交易期货市场
	TrdMarket_FuturesSimulateUS TrdMarket = 11  // 模拟交易期货市场
	TrdMarket_FuturesSimulateSG TrdMarket = 12  // 模拟交易期货市场
	TrdMarket_FuturesSimulateJP TrdMarket = 13  // 模拟交易期货市场
	TrdMarket_HKFund            TrdMarket = 113 // 香港基金市场
	TrdMarket_USFund            TrdMarket = 123 // 美国基金市场
	TrdMarket_SGFund            TrdMarket = 124 // 新加坡基金市场
	TrdMarket_MYFund            TrdMarket = 125 // 马来西亚基金市场
	TrdMarket_JPFund            TrdMarket = 126 // 日本基金市场
)

// =============================================================================
// TrdSecMarket (可交易证券所属市场)
// =============================================================================

// TrdSecMarket represents the market for tradable securities
type TrdSecMarket int32

const (
	TrdSecMarket_Unknown TrdSecMarket = 0  // 未知市场
	TrdSecMarket_HK      TrdSecMarket = 1  // 香港市场(股票、窝轮、牛熊、期权、期货等)
	TrdSecMarket_US      TrdSecMarket = 2  // 美国市场(股票、期权、期货等)
	TrdSecMarket_CN_SH   TrdSecMarket = 31 // 沪股市场(股票)
	TrdSecMarket_CN_SZ   TrdSecMarket = 32 // 深股市场(股票)
	TrdSecMarket_SG      TrdSecMarket = 41 // 新加坡市场(期货)
	TrdSecMarket_JP      TrdSecMarket = 51 // 日本市场(期货)
	TrdSecMarket_AU      TrdSecMarket = 61 // 澳大利亚
	TrdSecMarket_MY      TrdSecMarket = 71 // 马来西亚
	TrdSecMarket_CA      TrdSecMarket = 81 // 加拿大
	TrdSecMarket_FX      TrdSecMarket = 91 // 外汇
)

// =============================================================================
// PositionSide (持仓方向)
// =============================================================================

// PositionSide represents the side of a position
type PositionSide int32

const (
	PositionSide_None  PositionSide = 0 // 未知
	PositionSide_Long  PositionSide = 1 // 多仓
	PositionSide_Short PositionSide = 2 // 空仓
)

// =============================================================================
// OrderType (订单类型)
// =============================================================================

// OrderType represents the type of order
type OrderType int32

const (
	OrderType_None              OrderType = 0  // 未知
	OrderType_Normal            OrderType = 1  // 普通订单(港股的增强限价单、A股限价委托、美股的限价单)
	OrderType_Market            OrderType = 2  // 市价
	OrderType_AbsoluteLimit     OrderType = 3  // 港股_限价(只有价格完全匹配才成交)
	OrderType_Auction           OrderType = 4  // 港股_竞价
	OrderType_AuctionLimit      OrderType = 5  // 港股_竞价限价
	OrderType_SpecialLimit      OrderType = 6  // 港股_特别限价(即市价IOC)
	OrderType_SpecialLimitAll   OrderType = 7  // 港股_特别限价(要么全部成交，要么自动撤单)
	OrderType_Stop              OrderType = 10 // 止损市价单
	OrderType_StopLimit         OrderType = 11 // 止损限价单
	OrderType_MarketIfTouched   OrderType = 12 // 触及市价单(止盈)
	OrderType_LimitIfTouched    OrderType = 13 // 触及限价单(止盈)
	OrderType_TrailingStop      OrderType = 14 // 跟踪止损市价单
	OrderType_TrailingStopLimit OrderType = 15 // 跟踪止损限价单
	OrderType_TWAP              OrderType = 20 // 算法订单TWAP市价单(仅展示)
	OrderType_TWAPLimit         OrderType = 21 // 算法订单TWAP限价单(仅展示)
	OrderType_VWAP              OrderType = 22 // 算法订单VWAP市价单(仅展示)
	OrderType_VWAPLimit         OrderType = 23 // 算法订单VWAP限价单(仅展示)
)

// =============================================================================
// OrderStatus (订单状态)
// =============================================================================

// OrderStatus represents the status of an order
type OrderStatus int32

const (
	OrderStatus_None           OrderStatus = 0  // 未知状态
	OrderStatus_Unsubmitted    OrderStatus = 1  // 未提交
	OrderStatus_WaitingSubmit  OrderStatus = 2  // 等待提交
	OrderStatus_Submitting     OrderStatus = 3  // 提交中
	OrderStatus_SubmitFailed   OrderStatus = 4  // 提交失败，下单失败
	OrderStatus_TimeOut        OrderStatus = 5  // 处理超时，结果未知
	OrderStatus_Submitted      OrderStatus = 6  // 已提交，等待成交
	OrderStatus_FilledPart     OrderStatus = 7  // 部分成交
	OrderStatus_FilledAll      OrderStatus = 8  // 全部已成
	OrderStatus_CancellingPart OrderStatus = 9  // 正在撤单_部分(部分已成交，正在撤销剩余部分)
	OrderStatus_CancellingAll  OrderStatus = 10 // 正在撤单_全部
	OrderStatus_CancelledPart  OrderStatus = 11 // 部分成交，剩余部分已撤单
	OrderStatus_CancelledAll   OrderStatus = 12 // 全部已撤单，无成交
	OrderStatus_Failed         OrderStatus = 13 // 下单失败，服务拒绝
	OrderStatus_Disabled       OrderStatus = 14 // 已失效
	OrderStatus_Deleted        OrderStatus = 15 // 已删除，无成交的订单才能删除
	OrderStatus_FillCancelled  OrderStatus = 16 // 成交被撤销，一般遇不到，意思是已经成交的订单被回滚撤销，成交无效变为废单
)

// =============================================================================
// ModifyOrderOp (修改订单操作)
// =============================================================================

// ModifyOrderOp represents the operation to modify an order
type ModifyOrderOp int32

const (
	ModifyOrderOp_None    ModifyOrderOp = 0 // 未知
	ModifyOrderOp_Normal  ModifyOrderOp = 1 // 修改订单的数量、价格
	ModifyOrderOp_Cancel  ModifyOrderOp = 2 // 取消订单
	ModifyOrderOp_Disable ModifyOrderOp = 3 // 使订单失效
	ModifyOrderOp_Enable  ModifyOrderOp = 4 // 使订单生效
	ModifyOrderOp_Delete  ModifyOrderOp = 5 // 删除订单
)

// =============================================================================
// TrdSide (交易方向)
// =============================================================================

// TrdSide represents the side of a trade (buy or sell)
type TrdSide int32

const (
	TrdSide_None      TrdSide = 0 // 未知
	TrdSide_Buy       TrdSide = 1 // 买
	TrdSide_Sell      TrdSide = 2 // 卖
	TrdSide_SellShort TrdSide = 3 // 卖空
	TrdSide_BuyBack   TrdSide = 4 // 买回
)

// =============================================================================
// TrdCategory (交易品类)
// =============================================================================

// TrdCategory represents the category of trade
type TrdCategory int32

const (
	TrdCategory_None     TrdCategory = 0 // 未知
	TrdCategory_Security TrdCategory = 1 // 证券
	TrdCategory_Future   TrdCategory = 2 // 期货
)

// =============================================================================
// TrailType (跟踪止损类型)
// =============================================================================

// TrailType represents the type of trailing stop
type TrailType int32

const (
	TrailType_None   TrailType = 0 // 未知
	TrailType_Ratio  TrailType = 1 // 跟踪百分比
	TrailType_Amount TrailType = 2 // 跟踪额
)

// =============================================================================
// TimeInForce (订单有效期)
// =============================================================================

// TimeInForce represents the time in force for an order
type TimeInForce int32

const (
	TimeInForce_None TimeInForce = 0 // 未知
	TimeInForce_Day  TimeInForce = 1 // 当日有效
	TimeInForce_GTC  TimeInForce = 2 // 取消前有效
	TimeInForce_IOC  TimeInForce = 3 // 即时或取消
	TimeInForce_FOK  TimeInForce = 4 // 全部成交或取消
)

// =============================================================================
// DealStatus (成交状态)
// =============================================================================

// DealStatus represents the status of a deal
type DealStatus int32

const (
	DealStatus_OK        DealStatus = 0 // 正常
	DealStatus_Cancelled DealStatus = 1 // 成交被取消
	DealStatus_Changed   DealStatus = 2 // 成交被更改
)

// =============================================================================
// WarrantType (窝轮类型)
// =============================================================================

// WarrantType represents the type of warrant
type WarrantType int32

const (
	WarrantType_None   WarrantType = 0 // 未知
	WarrantType_Buy    WarrantType = 1 // 认购
	WarrantType_Sell   WarrantType = 2 // 认沽
	WarrantType_Bull   WarrantType = 3 // 牛
	WarrantType_Bear   WarrantType = 4 // 熊
	WarrantType_InLine WarrantType = 5 // 界内证
)

// =============================================================================
// PriceReminderType (到价提醒类型)
// =============================================================================

// PriceReminderType represents the type of price reminder
type PriceReminderType int32

const (
	PriceReminderType_None   PriceReminderType = 0 // 未知
	PriceReminderType_Above  PriceReminderType = 1 // 高于
	PriceReminderType_Below  PriceReminderType = 2 // 低于
	PriceReminderType_Remind PriceReminderType = 3 // 提醒
)

// =============================================================================
// PriceReminderOp (到价提醒操作)
// =============================================================================

// PriceReminderOp represents the operation for price reminder
type PriceReminderOp int32

const (
	PriceReminderOp_None PriceReminderOp = 0 // 未知
	PriceReminderOp_Add  PriceReminderOp = 1 // 添加
	PriceReminderOp_Del  PriceReminderOp = 2 // 删除
	PriceReminderOp_Edit PriceReminderOp = 3 // 修改
)

// =============================================================================
// AccounterType (户类型)
// =============================================================================

// AccounterType represents the type of account
type AccounterType int32

const (
	AccounterType_None    AccounterType = 0 // 未知
	AccounterType_Cash    AccounterType = 1 // 现金账户
	AccounterType_Margin  AccounterType = 2 // 保证金账户
	AccounterType_Short   AccounterType = 3 // 沽空账户
	AccounterType_Futures AccounterType = 4 // 期货账户
	AccounterType_Option  AccounterType = 5 // 期权账户
	AccounterType_Fund    AccounterType = 6 // 基金账户
)

// =============================================================================
// AccStatus (账户状态)
// =============================================================================

// AccStatus represents the status of an account
type AccStatus int32

const (
	AccStatus_None     AccStatus = 0 // 未知
	AccStatus_Normal   AccStatus = 1 // 正常
	AccStatus_Disabled AccStatus = 2 // 禁用
	AccStatus_Deleted  AccStatus = 3 // 已删除
	AccStatus_Locked   AccStatus = 4 // 锁定
)

// =============================================================================
// Currency (货币类型)
// =============================================================================

// Currency represents the type of currency
type Currency int32

const (
	Currency_None  Currency = 0  // 未知
	Currency_HKD   Currency = 1  // 港币
	Currency_USD   Currency = 2  // 美元
	Currency_CNY   Currency = 3  // 人民币
	Currency_HKD_C Currency = 4  // 港币(柜台)
	Currency_USD_C Currency = 5  // 美元(柜台)
	Currency_SGD   Currency = 6  // 新加坡元
	Currency_AUD   Currency = 7  // 澳元
	Currency_JPY   Currency = 8  // 日元
	Currency_MYR   Currency = 9  // 马来西亚林吉特
	Currency_CAD   Currency = 10 // 加拿大元
	Currency_EUR   Currency = 11 // 欧元
	Currency_GBP   Currency = 12 // 英镑
	Currency_CHF   Currency = 13 // 瑞士法郎
	Currency_THB   Currency = 14 // 泰铢
)

// =============================================================================
// PushDataType (推送数据类型)
// =============================================================================

// PushDataType represents the type of pushed data
type PushDataType int32

const (
	PushDataType_None      PushDataType = 0 // 未知
	PushDataType_Realtime  PushDataType = 1 // 实时
	PushDataType_ByDisConn PushDataType = 2 // 断线后补
	PushDataType_Cache     PushDataType = 3 // 缓存
)

// =============================================================================
// SecurityFirm (券商)
// =============================================================================

// SecurityFirm represents the security firm
type SecurityFirm int32

const (
	SecurityFirm_None                    SecurityFirm = 0 // 未知
	SecurityFirm_FutuSecurities          SecurityFirm = 1 // 富途证券
	SecurityFirm_FutuFuturesHK           SecurityFirm = 2 // 富途期货(香港)
	SecurityFirm_FutuSecuritiesFuturesHK SecurityFirm = 3 // 富途证券(期货香港)
	SecurityFirm_Virtu                   SecurityFirm = 4 // Virtu
	SecurityFirm_FutuWealth              SecurityFirm = 5 // 富途财富
	SecurityFirm_FutuTrust               SecurityFirm = 6 // 富途信托
)

// =============================================================================
// NotifyType (通知类型)
// =============================================================================

// NotifyType represents the type of notification
type NotifyType int32

const (
	NotifyType_None          NotifyType = 0 // 未知
	NotifyType_GtwEvent      NotifyType = 1 // 网关事件
	NotifyType_ProgramStatus NotifyType = 2 // 程序状态
	NotifyType_ConnStatus    NotifyType = 3 // 连接状态
	NotifyType_QotRight      NotifyType = 4 // 行情权限
	NotifyType_APILevel      NotifyType = 5 // API级别
	NotifyType_APIQuota      NotifyType = 6 // API配额
	NotifyType_UsedQuota     NotifyType = 7 // 已用配额
)

// =============================================================================
// VerificationType (验证码类型)
// =============================================================================

// VerificationType represents the type of verification
type VerificationType int32

const (
	VerificationType_None            VerificationType = 0 // 未知
	VerificationType_PicVerifyCode   VerificationType = 1 // 图形验证码
	VerificationType_SMSVerifyCode   VerificationType = 2 // 短信验证码
	VerificationType_EmailVerifyCode VerificationType = 3 // 邮箱验证码
)

// =============================================================================
// VerificationOp (验证码操作)
// =============================================================================

// VerificationOp represents the operation for verification
type VerificationOp int32

const (
	VerificationOp_None   VerificationOp = 0 // 未知
	VerificationOp_Get    VerificationOp = 1 // 获取验证码
	VerificationOp_Verify VerificationOp = 2 // 验证验证码
)

// =============================================================================
// IndexOptionType (指数期权类型)
// =============================================================================

// IndexOptionType represents the type of index option
type IndexOptionType int32

const (
	IndexOptionType_None        IndexOptionType = 0 // 所有
	IndexOptionType_Standard    IndexOptionType = 1 // 标准期权
	IndexOptionType_NonStandard IndexOptionType = 2 // 非标准期权
)

// =============================================================================
// CapitalFlowPeriodType (资金流向周期类型)
// =============================================================================

// CapitalFlowPeriodType represents the period type for capital flow
type CapitalFlowPeriodType int32

const (
	CapitalFlowPeriodType_None     CapitalFlowPeriodType = 0 // 未知
	CapitalFlowPeriodType_Intraday CapitalFlowPeriodType = 1 // 当日
	CapitalFlowPeriodType_Day5     CapitalFlowPeriodType = 2 // 5日
	CapitalFlowPeriodType_Day10    CapitalFlowPeriodType = 3 // 10日
	CapitalFlowPeriodType_Day20    CapitalFlowPeriodType = 4 // 20日
	CapitalFlowPeriodType_Day30    CapitalFlowPeriodType = 5 // 30日
	CapitalFlowPeriodType_Day60    CapitalFlowPeriodType = 6 // 60日
	CapitalFlowPeriodType_Day90    CapitalFlowPeriodType = 7 // 90日
)

// =============================================================================
// StockOwnerType (窝轮持有者类型)
// =============================================================================

// StockOwnerType represents the type of stock owner for warrants
type StockOwnerType int32

const (
	StockOwnerType_None  StockOwnerType = 0 // 全部
	StockOwnerType_Stock StockOwnerType = 1 // 正股
	StockOwnerType_Index StockOwnerType = 2 // 指数
)

// =============================================================================
// WarrantSortField (窝轮排序字段)
// =============================================================================

// WarrantSortField represents the field to sort warrants by
type WarrantSortField int32

const (
	WarrantSortField_None              WarrantSortField = 0  // 无排序
	WarrantSortField_Code              WarrantSortField = 1  // 代码
	WarrantSortField_LotSize           WarrantSortField = 2  // 每手
	WarrantSortField_Name              WarrantSortField = 3  // 名称
	WarrantSortField_Price             WarrantSortField = 4  // 当前价
	WarrantSortField_PriceRatio        WarrantSortField = 5  // 溢价率
	WarrantSortField_EffectiveLeverage WarrantSortField = 6  // 有效杠杆
	WarrantSortField_UpperStrikePrice  WarrantSortField = 7  // 上限价
	WarrantSortField_LowerStrikePrice  WarrantSortField = 8  // 下限价
	WarrantSortField_CurPrice          WarrantSortField = 9  // 街码量比
	WarrantSortField_VolRatio          WarrantSortField = 10 // 成交量比
	WarrantSortField_ImpliedVolatility WarrantSortField = 11 // 引伸波幅
	WarrantSortField_Delta             WarrantSortField = 12 // Delta
	WarrantSortField_ImplDelta         WarrantSortField = 13 // 引伸delta
	WarrantSortField_Vega              WarrantSortField = 14 // Vega
	WarrantSortField_Gamma             WarrantSortField = 15 // Gamma
	WarrantSortField_Theta             WarrantSortField = 16 // Theta
	WarrantSortField_Rho               WarrantSortField = 17 // Rho
)

// =============================================================================
// WarrantStatus (窝轮状态)
// =============================================================================

// WarrantStatus represents the status of a warrant
type WarrantStatus int32

const (
	WarrantStatus_None      WarrantStatus = 0 // 全部
	WarrantStatus_Normal    WarrantStatus = 1 // 正常
	WarrantStatus_Suspend   WarrantStatus = 2 // 停牌
	WarrantStatus_StopTrade WarrantStatus = 3 // 停止交易
)

// =============================================================================
// SecurityListStatus (证券列表状态)
// =============================================================================

// SecurityListStatus represents the status of a security list
type SecurityListStatus int32

const (
	SecurityListStatus_None     SecurityListStatus = 0 // 空
	SecurityListStatus_Normal   SecurityListStatus = 1 // 正常
	SecurityListStatus_Stop     SecurityListStatus = 2 // 停牌
	SecurityListStatus_Delisted SecurityListStatus = 3 // 已退市
	SecurityListStatus_PreStart SecurityListStatus = 4 // 预上市
	SecurityListStatus_Suspend  SecurityListStatus = 5 // 停牌
	SecurityListStatus_Cash     SecurityListStatus = 6 // 现金
	SecurityListStatus_Invalid  SecurityListStatus = 7 // 失效
)

// =============================================================================
// AcGrantRights (账户开通权限)
// =============================================================================

// AcGrantRights represents the granted rights for an account
type AcGrantRights int32

const (
	AcGrantRights_None          AcGrantRights = 0  // 无
	AcGrantRights_HKStock       AcGrantRights = 1  // 港股
	AcGrantRights_USStock       AcGrantRights = 2  // 美股
	AcGrantRights_CNHK          AcGrantRights = 3  // 沪股通
	AcGrantRights_SNHK          AcGrantRights = 4  // 深股通
	AcGrantRights_HKFuture      AcGrantRights = 5  // 港股期货
	AcGrantRights_HKOption      AcGrantRights = 6  // 港股期权
	AcGrantRights_SGDFuture     AcGrantRights = 7  // 新加坡期货
	AcGrantRights_USOption      AcGrantRights = 8  // 美股期权
	AcGrantRights_JPFuture      AcGrantRights = 9  // 日本期货
	AcGrantRights_MYFuture      AcGrantRights = 10 // 马来西亚期货
	AcGrantRights_AUFuture      AcGrantRights = 11 // 澳大利亚期货
	AcGrantRights_CNFuture      AcGrantRights = 12 // A股期货
	AcGrantRights_SGDFutureMain AcGrantRights = 13 // 新加坡期货主连
	AcGrantRights_JPFutureMain  AcGrantRights = 14 // 日本期货主连
	AcGrantRights_MYFutureMain  AcGrantRights = 15 // 马来西亚期货主连
	AcGrantRights_AUFutureMain  AcGrantRights = 16 // 澳大利亚期货主连
)

// =============================================================================
// AccRight (账户权限)
// =============================================================================

// AccRight represents the right for an account
type AccRight int32

const (
	AccRight_None      AccRight = 0  // 无
	AccRight_HKStock   AccRight = 1  // 港股证券
	AccRight_USStock   AccRight = 2  // 美股证券
	AccRight_HKFuture  AccRight = 3  // 港股期货
	AccRight_HKOption  AccRight = 4  // 港股期权
	AccRight_USOption  AccRight = 5  // 美股期权
	AccRight_SGDFuture AccRight = 6  // 新加坡期货
	AccRight_JPFuture  AccRight = 7  // 日本期货
	AccRight_MYFuture  AccRight = 8  // 马来西亚期货
	AccRight_AUFuture  AccRight = 9  // 澳大利亚期货
	AccRight_CNFuture  AccRight = 10 // A股期货
	AccRight_CNHK      AccRight = 11 // 沪股通
	AccRight_SNHK      AccRight = 12 // 深股通
	AccRight_HKFund    AccRight = 13 // 香港基金
	AccRight_USFund    AccRight = 14 // 美国基金
	AccRight_SGDFund   AccRight = 15 // 新加坡基金
	AccRight_MYFund    AccRight = 16 // 马来西亚基金
	AccRight_JPFund    AccRight = 17 // 日本基金
)

// =============================================================================
// AccMarket (账户市场)
// =============================================================================

// AccMarket represents the market for an account
type AccMarket int32

const (
	AccMarket_None      AccMarket = 0  // 无
	AccMarket_HK        AccMarket = 1  // 港股
	AccMarket_US        AccMarket = 2  // 美股
	AccMarket_CN        AccMarket = 3  // A股
	AccMarket_HKFuture  AccMarket = 4  // 港股期货
	AccMarket_Future    AccMarket = 5  // 期货
	AccMarket_SGDFuture AccMarket = 6  // 新加坡期货
	AccMarket_JPFuture  AccMarket = 7  // 日本期货
	AccMarket_MYFuture  AccMarket = 8  // 马来西亚期货
	AccMarket_AUFuture  AccMarket = 9  // 澳大利亚期货
	AccMarket_CNFuture  AccMarket = 10 // A股期货
	AccMarket_HKCC      AccMarket = 11 // 港股通
	AccMarket_HKFund    AccMarket = 12 // 香港基金
	AccMarket_USFund    AccMarket = 13 // 美国基金
	AccMarket_SGDFund   AccMarket = 14 // 新加坡基金
	AccMarket_MYFund    AccMarket = 15 // 马来西亚基金
	AccMarket_JPFund    AccMarket = 16 // 日本基金
)

// =============================================================================
// AccTradingMarket (账户交易市场)
// =============================================================================

// AccTradingMarket represents the trading market for an account
type AccTradingMarket int32

const (
	AccTradingMarket_None      AccTradingMarket = 0  // 无
	AccTradingMarket_HK        AccTradingMarket = 1  // 香港
	AccTradingMarket_US        AccTradingMarket = 2  // 美国
	AccTradingMarket_CN        AccTradingMarket = 3  // 大陆
	AccTradingMarket_SGDFuture AccTradingMarket = 4  // 新加坡期货
	AccTradingMarket_JPFuture  AccTradingMarket = 5  // 日本期货
	AccTradingMarket_MYFuture  AccTradingMarket = 6  // 马来西亚期货
	AccTradingMarket_AUFuture  AccTradingMarket = 7  // 澳大利亚期货
	AccTradingMarket_CNHK      AccTradingMarket = 8  // 沪股通
	AccTradingMarket_SNHK      AccTradingMarket = 9  // 深股通
	AccTradingMarket_HKFund    AccTradingMarket = 10 // 香港基金
	AccTradingMarket_USFund    AccTradingMarket = 11 // 美国基金
	AccTradingMarket_SGDFund   AccTradingMarket = 12 // 新加坡基金
	AccTradingMarket_MYFund    AccTradingMarket = 13 // 马来西亚基金
	AccTradingMarket_JPFund    AccTradingMarket = 14 // 日本基金
)

// =============================================================================
// AccAuthenStatus (账户认证状态)
// =============================================================================

// AccAuthenStatus represents the authentication status for an account
type AccAuthenStatus int32

const (
	AccAuthenStatus_None          AccAuthenStatus = 0 // 无
	AccAuthenStatus_Normal        AccAuthenStatus = 1 // 正常
	AccAuthenStatus_OnlyQuotation AccAuthenStatus = 2 // 仅有行情权限
	AccAuthenStatus_Pending       AccAuthenStatus = 3 // 认证中
	AccAuthenStatus_Locked        AccAuthenStatus = 4 // 账户被锁定
	AccAuthenStatus_Frozen        AccAuthenStatus = 5 // 账户被冻结
)

// =============================================================================
// RiskLevel (风险等级)
// =============================================================================

// RiskLevel represents the risk level
type RiskLevel int32

const (
	RiskLevel_None   RiskLevel = 0 // 未知
	RiskLevel_Low    RiskLevel = 1 // 低风险
	RiskLevel_Medium RiskLevel = 2 // 中风险
	RiskLevel_High   RiskLevel = 3 // 高风险
)

// =============================================================================
// PDTStatus ( PDT状态)
// =============================================================================

// PDTStatus represents the PDT (Pattern Day Trader) status
type PDTStatus int32

const (
	PDTStatus_None       PDTStatus = 0 // 无
	PDTStatus_Warning    PDTStatus = 1 // 警告
	PDTStatus_Restricted PDTStatus = 2 // 限制
	PDTStatus_Call       PDTStatus = 3 // 追缴
)

// =============================================================================
// DelayStatisticsType (延迟统计类型)
// =============================================================================

// DelayStatisticsType represents the type of delay statistics
type DelayStatisticsType int32

const (
	DelayStatisticsType_None       DelayStatisticsType = 0 // 未知
	DelayStatisticsType_QotPush    DelayStatisticsType = 1 // 行情推送
	DelayStatisticsType_ReqReply   DelayStatisticsType = 2 // 请求应答
	DelayStatisticsType_PlaceOrder DelayStatisticsType = 3 // 下单
)

// =============================================================================
// SortDirection (排序方向)
// =============================================================================

// SortDirection represents the direction of sorting
type SortDirection int32

const (
	SortDirection_None    SortDirection = 0 // 无
	SortDirection_Ascend  SortDirection = 1 // 升序
	SortDirection_Descend SortDirection = 2 // 降序
)

// =============================================================================
// StockFilterField (选股字段)
// =============================================================================

// StockFilterField represents the field for stock filtering
type StockFilterField int32

const (
	StockFilterField_None         StockFilterField = 0  // 无
	StockFilterField_ChangeRate   StockFilterField = 1  // 涨跌幅
	StockFilterField_ChangeVal    StockFilterField = 2  // 涨跌额
	StockFilterField_Volume       StockFilterField = 3  // 成交量
	StockFilterField_Turnover     StockFilterField = 4  // 成交额
	StockFilterField_TurnoverRate StockFilterField = 5  // 换手率
	StockFilterField_VolumeRatio  StockFilterField = 6  // 量比
	StockFilterField_BidAskRatio  StockFilterField = 7  // 委比
	StockFilterField_DayAmplitude StockFilterField = 8  // 日振幅
	StockFilterField_MarketVal    StockFilterField = 9  // 总市值
	StockFilterField_CirculateVal StockFilterField = 10 // 流通市值
	StockFilterField_More         StockFilterField = 11 // 更多
)

// =============================================================================
// UserSecurityGroupType (自选股分组类型)
// =============================================================================

// UserSecurityGroupType represents the type of user security group
type UserSecurityGroupType int32

const (
	UserSecurityGroupType_None     UserSecurityGroupType = 0 // 未知
	UserSecurityGroupType_Optional UserSecurityGroupType = 1 // 自选
)

// =============================================================================
// AcUpdateFields (账户更新字段)
// =============================================================================

// AcUpdateFields represents the fields for account update
type AcUpdateFields int32

const (
	AcUpdateFields_None      AcUpdateFields = 0  // 无
	AcUpdateFields_EnvStatus AcUpdateFields = 1  // 环境状态
	AcUpdateFields_AccInfo   AcUpdateFields = 2  // 账户信息
	AcUpdateFields_RiskInfo  AcUpdateFields = 3  // 风控信息
	AcUpdateFields_BP        AcUpdateFields = 4  // 购买力
	AcUpdateFields_Funds     AcUpdateFields = 5  // 资金
	AcUpdateFields_Position  AcUpdateFields = 6  // 持仓
	AcUpdateFields_Order     AcUpdateFields = 7  // 订单
	AcUpdateFields_Trade     AcUpdateFields = 8  // 成交
	AcUpdateFields_Options   AcUpdateFields = 9  // 期权
	AcUpdateFields_Report    AcUpdateFields = 10 // 对账单
)

// =============================================================================
// OrderFillNotifyType (成交通知类型)
// =============================================================================

// OrderFillNotifyType represents the type of order fill notification
type OrderFillNotifyType int32

const (
	OrderFillNotifyType_None    OrderFillNotifyType = 0 // 未知
	OrderFillNotifyType_Fill    OrderFillNotifyType = 1 // 订单成交
	OrderFillNotifyType_Cancel  OrderFillNotifyType = 2 // 订单取消
	OrderFillNotifyType_Changed OrderFillNotifyType = 3 // 订单修改
)

// =============================================================================
// TrdAccCertType (交易账户证书类型)
// =============================================================================

// TrdAccCertType represents the certificate type for trading account
type TrdAccCertType int32

const (
	TrdAccCertType_None       TrdAccCertType = 0 // 未知
	TrdAccCertType_HKID       TrdAccCertType = 1 // 香港身份证
	TrdAccCertType_Passport   TrdAccCertType = 2 // 护照
	TrdAccCertType_License    TrdAccCertType = 3 // 商业执照
	TrdAccCertType_CreditCard TrdAccCertType = 4 // 信用卡
	TrdAccCertType_Other      TrdAccCertType = 5 // 其他
)

// =============================================================================
// GtwEventType (网关事件类型)
// =============================================================================

// GtwEventType represents the type of gateway event
type GtwEventType int32

const (
	GtwEventType_None                GtwEventType = 0  // 未知
	GtwEventType_LocalCfgLoadFailed  GtwEventType = 1  // 本地配置文件加载失败
	GtwEventType_APISvrRunFailed     GtwEventType = 2  // 网关监听服务运行失败
	GtwEventType_ForceUpdate         GtwEventType = 3  // 强制升级网关
	GtwEventType_LoginFailed         GtwEventType = 4  // 登录牛牛服务器失败
	GtwEventType_UnAgreeDisclaimer   GtwEventType = 5  // 未同意免责声明，无法运行
	GtwEventType_NetCfgMissing       GtwEventType = 6  // 缺少网络连接配置
	GtwEventType_KickedOut           GtwEventType = 7  // 登录被踢下线
	GtwEventType_LoginPwdChanged     GtwEventType = 8  // 登陆密码变更
	GtwEventType_BanLogin            GtwEventType = 9  // 牛牛后台不允许该账号登陆
	GtwEventType_NeedPicVerifyCode   GtwEventType = 10 // 登录需要输入图形验证码
	GtwEventType_NeedPhoneVerifyCode GtwEventType = 11 // 登录需要输入手机验证码
	GtwEventType_AppDataNotExist     GtwEventType = 12 // 程序打包数据丢失
	GtwEventType_NessaryDataMissing  GtwEventType = 13 // 必要的数据没同步成功
	GtwEventType_TradePwdChanged     GtwEventType = 14 // 交易密码变更通知
	GtwEventType_EnableDeviceLock    GtwEventType = 15 // 需启用设备锁
)

// =============================================================================
// PriceReminderFreq (到价提醒频率)
// =============================================================================

// PriceReminderFreq represents the frequency of price reminder
type PriceReminderFreq int32

const (
	PriceReminderFreq_None      PriceReminderFreq = 0 // 未知
	PriceReminderFreq_Once      PriceReminderFreq = 1 // 只提醒一次
	PriceReminderFreq_OnceDaily PriceReminderFreq = 2 // 每天提醒一次
	PriceReminderFreq_Always    PriceReminderFreq = 3 // 持续提醒
)

// =============================================================================
// RetType (返回结果)
// =============================================================================

// RetType represents the return type of an API call
type RetType int32

const (
	RetType_Succeed    RetType = 0    // 成功
	RetType_Failed     RetType = -1   // 失败
	RetType_TimeOut    RetType = -100 // 超时
	RetType_DisConnect RetType = -200 // 连接断开
	RetType_Unknown    RetType = -400 // 未知结果
	RetType_Invalid    RetType = -500 // 包内容非法
)

// =============================================================================
// PacketEncAlgo (包加密算法)
// =============================================================================

// PacketEncAlgo represents the packet encryption algorithm
type PacketEncAlgo int32

const (
	PacketEncAlgo_FTAES_ECB PacketEncAlgo = 0  // 富途修改过的AES的ECB加密模式
	PacketEncAlgo_None      PacketEncAlgo = -1 // 不加密
	PacketEncAlgo_AES_ECB   PacketEncAlgo = 1  // 标准的AES的ECB加密模式
	PacketEncAlgo_AES_CBC   PacketEncAlgo = 2  // 标准的AES的CBC加密模式
)

// =============================================================================
// ProtoFmt (协议格式)
// =============================================================================

// ProtoFmt represents the protocol format
type ProtoFmt int32

const (
	ProtoFmt_Protobuf ProtoFmt = 0 // Google Protobuf格式
	ProtoFmt_Json     ProtoFmt = 1 // Json格式
)

// =============================================================================
// UserAttribution (用户注册归属地)
// =============================================================================

// UserAttribution represents the user attribution region
type UserAttribution int32

const (
	UserAttribution_Unknown UserAttribution = 0 // 未知
	UserAttribution_NN      UserAttribution = 1 // 大陆
	UserAttribution_MM      UserAttribution = 2 // MooMoo
	UserAttribution_SG      UserAttribution = 3 // 新加坡
	UserAttribution_AU      UserAttribution = 4 // 澳洲
	UserAttribution_JP      UserAttribution = 5 // 日本
	UserAttribution_HK      UserAttribution = 6 // 香港
)

// =============================================================================
// ProgramStatusType (程序状态)
// =============================================================================

// ProgramStatusType represents the program status
type ProgramStatusType int32

const (
	ProgramStatusType_None                 ProgramStatusType = 0  // 未知
	ProgramStatusType_Loaded               ProgramStatusType = 1  // 已完成类似加载配置,启动服务器等操作,服务器启动之前的状态无需返回
	ProgramStatusType_Loging               ProgramStatusType = 2  // 登录中
	ProgramStatusType_NeedPicVerifyCode    ProgramStatusType = 3  // 需要图形验证码
	ProgramStatusType_NeedPhoneVerifyCode  ProgramStatusType = 4  // 需要手机验证码
	ProgramStatusType_LoginFailed          ProgramStatusType = 5  // 登录失败,详细原因在描述返回
	ProgramStatusType_ForceUpdate          ProgramStatusType = 6  // 客户端版本过低
	ProgramStatusType_NessaryDataPreparing ProgramStatusType = 7  // 正在拉取类似免责声明等一些必要信息
	ProgramStatusType_NessaryDataMissing   ProgramStatusType = 8  // 缺少必要信息
	ProgramStatusType_UnAgreeDisclaimer    ProgramStatusType = 9  // 未同意免责声明
	ProgramStatusType_Ready                ProgramStatusType = 10 // 可以接收业务协议收发,正常可用状态
	ProgramStatusType_ForceLogout          ProgramStatusType = 11 // 被强制退出登录,例如修改了登录密码,中途打开设备锁等
	ProgramStatusType_DisclaimerPullFailed ProgramStatusType = 12 // 拉取免责声明标志失败
)

// =============================================================================
// Session (交易时段)
// =============================================================================

// Session represents the trading session
type Session int32

const (
	Session_None      Session = 0 // 未知
	Session_RTH       Session = 1 // 常规交易时段 (Regular Trading Hours)
	Session_ETH       Session = 2 // 盘后交易时段 (Extended Trading Hours)
	Session_All       Session = 3 // 所有时段
	Session_Overnight Session = 4 // 隔夜交易
)

// =============================================================================
// TrdAccType (交易账户类型)
// =============================================================================

// TrdAccType represents the type of trading account
type TrdAccType int32

const (
	TrdAccType_Unknown     TrdAccType = 0 // 未知类型
	TrdAccType_Cash        TrdAccType = 1 // 现金账户
	TrdAccType_Margin      TrdAccType = 2 // 保证金账户
	TrdAccType_TFSA        TrdAccType = 3 // 加拿大免税账户
	TrdAccType_RRSP        TrdAccType = 4 // 加拿大注册退休账户
	TrdAccType_SRRSP       TrdAccType = 5 // 加拿大配偶退休账户
	TrdAccType_Derivatives TrdAccType = 6 // 日本衍生品账户
)

// =============================================================================
// TrdAccStatus (交易账户状态)
// =============================================================================

// TrdAccStatus represents the status of a trading account
type TrdAccStatus int32

const (
	TrdAccStatus_Active   TrdAccStatus = 0 // 正常
	TrdAccStatus_Disabled TrdAccStatus = 1 // 停用
)

// =============================================================================
// TrdAccRole (账户类型)
// =============================================================================

// TrdAccRole represents the role of a trading account
type TrdAccRole int32

const (
	TrdAccRole_Unknown TrdAccRole = 0 // 未知
	TrdAccRole_Normal  TrdAccRole = 1 // 普通账户
	TrdAccRole_Master  TrdAccRole = 2 // 主账户
	TrdAccRole_IPO     TrdAccRole = 3 // IPO账户，仅MY券商
)

// =============================================================================
// CltRiskLevel (账户风险控制等级)
// =============================================================================

// CltRiskLevel represents the client risk control level
type CltRiskLevel int32

const (
	CltRiskLevel_Unknown      CltRiskLevel = -1 // 未知
	CltRiskLevel_Safe         CltRiskLevel = 0  // 安全
	CltRiskLevel_Warning      CltRiskLevel = 1  // 预警
	CltRiskLevel_Danger       CltRiskLevel = 2  // 危险
	CltRiskLevel_AbsoluteSafe CltRiskLevel = 3  // 绝对安全
	CltRiskLevel_OptDanger    CltRiskLevel = 4  // 危险，期权相关
)

// =============================================================================
// CltRiskStatus (风险状态，共分9个等级)
// =============================================================================

// CltRiskStatus represents the client risk status (9 levels, LEVEL1 safest)
type CltRiskStatus int32

const (
	CltRiskStatus_Unknown CltRiskStatus = 0 // 未知
	CltRiskStatus_Level1  CltRiskStatus = 1 // 非常安全
	CltRiskStatus_Level2  CltRiskStatus = 2 // 安全
	CltRiskStatus_Level3  CltRiskStatus = 3 // 较安全
	CltRiskStatus_Level4  CltRiskStatus = 4 // 较低风险
	CltRiskStatus_Level5  CltRiskStatus = 5 // 中等风险
	CltRiskStatus_Level6  CltRiskStatus = 6 // 较高风险
	CltRiskStatus_Level7  CltRiskStatus = 7 // 预警
	CltRiskStatus_Level8  CltRiskStatus = 8 // 预警
	CltRiskStatus_Level9  CltRiskStatus = 9 // 预警
)

// =============================================================================
// DTStatus (日内交易限制情况)
// =============================================================================

// DTStatus represents the day trading status
type DTStatus int32

const (
	DTStatus_Unknown   DTStatus = 0 // 未知
	DTStatus_Unlimited DTStatus = 1 // 无限次
	DTStatus_EMCall    DTStatus = 2 // EM Call
	DTStatus_DTCall    DTStatus = 3 // DT Call
)

// =============================================================================
// TrdSubAccType (JP子账户类型)
// =============================================================================

// TrdSubAccType represents the type of trading sub-account (mainly for JP)
type TrdSubAccType int32

const (
	TrdSubAccType_None                          TrdSubAccType = 0  // 未知
	TrdSubAccType_JP_GENERAL                    TrdSubAccType = 1  // 日本-一般口座-long
	TrdSubAccType_JP_TOKUTEI                    TrdSubAccType = 2  // 日本-特定口座-long
	TrdSubAccType_JP_NISA_GENERAL               TrdSubAccType = 3  // 日本-一般NISA
	TrdSubAccType_JP_NISA_TSUMITATE             TrdSubAccType = 4  // 日本-累计NISA
	TrdSubAccType_JP_GENERAL_SHORT              TrdSubAccType = 5  // 日本-一般口座-Short
	TrdSubAccType_JP_TOKUTEI_SHORT              TrdSubAccType = 6  // 日本-特定口座-Short
	TrdSubAccType_JP_HONPO_GENERAL              TrdSubAccType = 7  // 日本-本国信用交易抵押品-一般
	TrdSubAccType_JP_GAIKOKU_GENERAL            TrdSubAccType = 8  // 日本-外国信用交易抵押品-一般
	TrdSubAccType_JP_HONPO_TOKUTEI              TrdSubAccType = 9  // 日本-本国信用交易抵押品-特定
	TrdSubAccType_JP_GAIKOKU_TOKUTEI            TrdSubAccType = 10 // 日本-外国信用交易抵押品-特定
	TrdSubAccType_JP_DERIVATIVE_LONG            TrdSubAccType = 11 // 日本-衍生品-Long
	TrdSubAccType_JP_DERIVATIVE_SHORT           TrdSubAccType = 12 // 日本-衍生品-Short
	TrdSubAccType_JP_HONPO_DERIVATIVE_GENERAL   TrdSubAccType = 13 // 日本-本国衍生品证据金-一般
	TrdSubAccType_JP_GAIKOKU_DERIVATIVE_GENERAL TrdSubAccType = 14 // 日本-外国衍生品证据金-一般
	TrdSubAccType_JP_HONPO_DERIVATIVE_TOKUTEI   TrdSubAccType = 15 // 日本-本国衍生品证据金-特定
	TrdSubAccType_JP_GAIKOKU_DERIVATIVE_TOKUTEI TrdSubAccType = 16 // 日本-外国衍生品证据金-特定
)

// =============================================================================
// TrdAssetCategory (资产类别)
// =============================================================================

// TrdAssetCategory represents the trading asset category
type TrdAssetCategory int32

const (
	TrdAssetCategory_Unknown TrdAssetCategory = 0 // 未知
	TrdAssetCategory_JP      TrdAssetCategory = 1 // 本国
	TrdAssetCategory_US      TrdAssetCategory = 2 // 外国
)
