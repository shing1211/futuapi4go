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
	ProtoID_TestCmd            = 1008 // 测试命令
	ProtoID_UsedQuota          = 1010 // 已使用额度

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
	ProtoID_Qot_GetRehab               = 3102 // 获取复权信息
	ProtoID_Qot_RequestHistoryKL      = 3103 // 拉取历史K线
	ProtoID_Qot_RequestHistoryKLQuota = 3104 // 拉取历史K线已经用掉的额度
	ProtoID_Qot_RequestRehab          = 3105 // 获取除权信息
	ProtoID_Qot_GetHistoryKLPoints     = 3106 // 获取历史K线数据点

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
	ProtoID_Qot_GetOptionExpirationDate         = 3224 // 获取期权到期日
	ProtoID_Qot_GetFinancialsEarningsPriceMove  = 3225 // 获取财报价格变动
	ProtoID_Qot_GetFinancialsEarningsPriceHistory = 3226 // 获取财报价格历史
	ProtoID_Qot_GetFinancialsStatements         = 3227 // 获取财务报表
	ProtoID_Qot_GetFinancialsRevenueBreakdown = 3228 // 获取营收 breakdown
	ProtoID_Qot_GetResearchAnalystConsensus   = 3229 // 获取研究分析师共识
	ProtoID_Qot_GetResearchRatingSummary      = 3230 // 获取研究评级摘要
	ProtoID_Qot_GetResearchMorningstarReport  = 3231 // 获取晨星研究报告
	ProtoID_Qot_GetValuationDetail           = 3232 // 获取估值详情
	ProtoID_Qot_GetValuationPlateStockList    = 3233 // 获取估值板块股票列表
	ProtoID_Qot_GetCorporateActionsDividends  = 3234 // 获取股息红利
	ProtoID_Qot_GetCorporateActionsBuybacks   = 3235 // 获取回购
	ProtoID_Qot_GetCorporateActionsStockSplits = 3236 // 获取拆股
	ProtoID_Qot_GetShareholdersOverview        = 3237 // 获取股东概况
	ProtoID_Qot_GetShareholdersHoldingChanges = 3238 // 获取股东持股变动
	ProtoID_Qot_GetShareholdersHolderDetail   = 3239 // 获取股东详情
	ProtoID_Qot_GetShareholdersInstitutional   = 3240 // 获取机构股东
	ProtoID_Qot_GetInsiderHolderList          = 3241 // 获取内部人列表
	ProtoID_Qot_GetInsiderTradeList           = 3242 // 获取内部人交易列表
	ProtoID_Qot_GetCompanyProfile             = 3243 // 获取公司概况
	ProtoID_Qot_GetCompanyExecutives           = 3244 // 获取公司高管
	ProtoID_Qot_GetCompanyExecutiveBackground  = 3245 // 获取高管背景
	ProtoID_Qot_GetCompanyOperationalEfficiency = 3246 // 获取经营效率
	ProtoID_Qot_GetTopTenBuySellBrokers        = 3247 // 获取十大买卖券商
	ProtoID_Qot_GetDailyShortVolume           = 3248 // 获取每日short volume
	ProtoID_Qot_GetShortInterest              = 3249 // 获取做空利息
	ProtoID_Qot_GetOptionVolatility           = 3250 // 获取期权波动率
	ProtoID_Qot_GetOptionExerciseProbability  = 3251 // 获取期权行权概率

	// v10.6+ Screen APIs (C++-only OpenD)
	ProtoID_Qot_StockScreen      = 3252 // 条件选股(新版股票筛选)
	ProtoID_Qot_OptionScreen     = 3253 // 期权筛选
	ProtoID_Qot_WarrantScreen    = 3254 // 窝轮筛选
	ProtoID_Qot_GetOptionQuote   = 3255 // 获取期权实时行情
	ProtoID_Qot_GetOptionStrategy = 3256 // 获取期权策略组合列表
	ProtoID_Qot_GetOptionStrategyAnalysis = 3257 // 获取期权策略组合分析
	ProtoID_Qot_GetOptionStrategySpread  = 3258 // 获取期权策略价差列表

	// v10.7+ Combo Trade APIs
	ProtoID_Trd_GetComboMaxTrdQtys = 2112 // 查询组合最大可交易数量
	ProtoID_Trd_PlaceComboOrder    = 2227 // 组合下单

	// v10.8+ Indicator & Search APIs
	ProtoID_Qot_GetIndicatorList  = 3259 // 获取指标列表
	ProtoID_Qot_RequestIndicatorCalc = 3260 // 请求异步指标计算
	ProtoID_Qot_PushIndicatorCalc = 3261 // 指标计算推送
	ProtoID_Qot_GetSearchQuote    = 3262 // 搜索报价
	ProtoID_Qot_GetSearchNews     = 3263 // 搜索新闻

	// v10.8+ Option Analytics APIs (3301-3314)
	ProtoID_Qot_GetOptionMarketStatistic = 3301 // 获取期权市场统计
	ProtoID_Qot_GetOptionUnderlyingHisStatistic = 3302 // 获取期权标的的历史统计
	ProtoID_Qot_GetOptionUnderlyingOverview = 3303 // 批量获取期权标的概览
	ProtoID_Qot_GetOptionUnderlyingHisVolatility = 3304 // 获取期权标的历史波动率
	ProtoID_Qot_GetOptionUnderlyingRank = 3305 // 获取期权标的排名
	ProtoID_Qot_GetOptionRank      = 3306 // 获取期权合约排名
	ProtoID_Qot_GetOptionEvent     = 3307 // 获取期权大事件
	ProtoID_Qot_GetOptionEventAlert = 3308 // 获取期权大事件提醒
	ProtoID_Qot_SetOptionEventAlert = 3309 // 设置期权大事件提醒
	ProtoID_Qot_UpdateOptionEvent  = 3310 // 推送期权大事件
	ProtoID_Qot_GetOptionZeroDteScreener = 3311 // 获取0DTE期权筛选
	ProtoID_Qot_GetOptionZeroDteContract = 3312 // 获取0DTE期权合约
	ProtoID_Qot_GetOptionEarningsScreener = 3313 // 获取期权财报筛选
	ProtoID_Qot_GetOptionSellerScreener  = 3314 // 获取期权卖方筛选

	// v10.8+ Market Fundamentals APIs (3401-3417)
	ProtoID_Qot_GetEarningsCalendar = 3401 // 获取财报日历
	ProtoID_Qot_GetMacroIndicatorList = 3402 // 获取宏观指标列表
	ProtoID_Qot_GetMacroIndicatorHistory = 3403 // 获取宏观指标历史
	ProtoID_Qot_GetFedWatchTargetRate = 3404 // 获取FedWatch目标利率概率
	ProtoID_Qot_GetFedWatchDotPlot = 3405 // 获取FedWatch点阵图
	ProtoID_Qot_GetEarningsBeatRank = 3406 // 获取财报超预期排名
	ProtoID_Qot_GetDividendRank    = 3407 // 获取股息排名
	ProtoID_Qot_GetDividendCalendar = 3408 // 获取股息日历
	ProtoID_Qot_GetEconomicCalendar = 3409 // 获取经济日历
	ProtoID_Qot_GetUSPreMarketRank = 3410 // 获取美股盘前排名
	ProtoID_Qot_GetUSAfterHoursRank = 3411 // 获取美股盘后排名
	ProtoID_Qot_GetUSOvernightRank = 3412 // 获取美股隔夜排名
	ProtoID_Qot_GetTopMoversRank   = 3413 // 获取热门异动排名
	ProtoID_Qot_GetHotList         = 3414 // 获取热门榜单
	ProtoID_Qot_GetShortSellingRank = 3415 // 获取做空排名
	ProtoID_Qot_GetPeriodChangeRank = 3416 // 获取区间涨跌排名
	ProtoID_Qot_GetHighDividendSOERank = 3417 // 获取高股息国企排名

	// v10.8+ Institutional APIs (3418-3425)
	ProtoID_Qot_GetInstitutionList = 3418 // 获取机构列表
	ProtoID_Qot_GetInstitutionProfile = 3419 // 获取机构简介
	ProtoID_Qot_GetInstitutionDistribution = 3420 // 获取机构配置分布
	ProtoID_Qot_GetInstitutionHoldingChange = 3421 // 获取机构持仓变动
	ProtoID_Qot_GetInstitutionHoldingList = 3422 // 获取机构持仓列表
	ProtoID_Qot_GetArkFundHolding = 3423 // 获取ARK基金持仓
	ProtoID_Qot_GetArkStockDynamic = 3424 // 获取ARK个股动态
	ProtoID_Qot_GetArkActiveTransaction = 3425 // 获取ARK主动交易

	// v10.8+ Other APIs (3426-3433)
	ProtoID_Qot_GetRatingChange   = 3426 // 获取评级变动
	ProtoID_Qot_GetIndustrialChainList = 3427 // 获取产业链列表
	ProtoID_Qot_GetIndustrialChainDetail = 3428 // 获取产业链详情
	ProtoID_Qot_GetIndustrialChainByPlate = 3429 // 根据板块获取产业链
	ProtoID_Qot_GetIndustrialPlateInfo = 3430 // 获取产业板块信息
	ProtoID_Qot_GetIndustrialPlateStock = 3431 // 获取产业板块成分股
	ProtoID_Qot_GetHeatMapData    = 3432 // 获取热力图数据
	ProtoID_Qot_GetRiseFallDistribution = 3433 // 获取涨跌分布

	// v10.8+ SkillWrap APIs (3801-3803)
	ProtoID_SkillWrap_TechnicalUnusual  = 3801 // 技术指标异动
	ProtoID_SkillWrap_FinancialUnusual  = 3802 // 财务数据异动
	ProtoID_SkillWrap_DerivativeUnusual = 3803 // 衍生品数据异动

	// SkillWrap language codes — used by *UnusualReq.LanguageId to control the
	// language of the AI-generated `Content` field in the response.
	SkillWrapLang_ZH_CN int32 = 0 // 简体中文
	SkillWrapLang_ZH_HK int32 = 1 // 繁體中文
	SkillWrapLang_EN    int32 = 2 // English
	SkillWrapLang_TH    int32 = 4 // ภาษาไทย
	SkillWrapLang_JA    int32 = 5 // 日本語
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
	ProtoID_Qot_UpdateOptionEvent,
	ProtoID_Qot_PushIndicatorCalc,
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
	Market_CC   = 101 // 加密货币市场
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
	SecurityType_Crypto   SecurityType = 12 // 加密货币
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
	SubType_K_Year     SubType = 16 // 年K线
	SubType_OrderBookOdd SubType = 18 // 碎股买卖摆盘
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
	MarketState_StibAfterHoursWait MarketState = 25 // 盘后竞价等待
	MarketState_StibAfterHoursBegin MarketState = 26 // 盘后竞价开始
	MarketState_StibAfterHoursEnd MarketState = 27 // 盘后竞价结束
	MarketState_Night             MarketState = 28 // 夜间交易
	MarketState_TradeAtLast       MarketState = 29 // 收盘集合竞价
	MarketState_Overnight         MarketState = 30 // 隔夜交易
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
	TrdMarket_Crypto            TrdMarket = 7   // 加密货币市场
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
	TrdCategory_Crypto   TrdCategory = 3 // 加密货币
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
	SecurityFirm_None           SecurityFirm = 0 // 未知
	SecurityFirm_FutuSecurities SecurityFirm = 1 // 富途证券(香港)
	SecurityFirm_FutuInc        SecurityFirm = 2 // 富途证券(美国)
	SecurityFirm_FutuSG         SecurityFirm = 3 // 富途证券(新加坡)
	SecurityFirm_FutuAU         SecurityFirm = 4 // 富途证券(澳洲)
	SecurityFirm_FutuCA         SecurityFirm = 5 // 富途证券(加拿大)
	SecurityFirm_FutuMY         SecurityFirm = 6 // 富途证券(马来西亚)
	SecurityFirm_FutuJP         SecurityFirm = 7 // 富途证券(日本)
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

// =============================================================================
// IndicatorLangType (指标语言类型)
// =============================================================================

// IndicatorLangType represents the language type for indicators
type IndicatorLangType int32

const (
	IndicatorLangType_Unknown IndicatorLangType = 0 // 未知
	IndicatorLangType_MyLang  IndicatorLangType = 1 // 麦语言
	IndicatorLangType_Python  IndicatorLangType = 2 // Python
)

// =============================================================================
// IndicatorSearchMode (指标搜索模式)
// =============================================================================

// IndicatorSearchMode represents the search mode for indicators
type IndicatorSearchMode int32

const (
	IndicatorSearchMode_Partial IndicatorSearchMode = 0 // 模糊匹配
	IndicatorSearchMode_Exact   IndicatorSearchMode = 1 // 精确匹配
)

// =============================================================================
// IndicatorParamValueType (指标参数值类型)
// =============================================================================

// IndicatorParamValueType represents the type of indicator parameter value
type IndicatorParamValueType int32

const (
	IndicatorParamValueType_Unknown IndicatorParamValueType = 0 // 未知
	IndicatorParamValueType_INT     IndicatorParamValueType = 1 // 整数
	IndicatorParamValueType_FLOAT   IndicatorParamValueType = 2 // 浮点数
	IndicatorParamValueType_STRING  IndicatorParamValueType = 3 // 字符串
	IndicatorParamValueType_COLOR   IndicatorParamValueType = 4 // 颜色
	IndicatorParamValueType_SHAPE   IndicatorParamValueType = 5 // 形状
	IndicatorParamValueType_LINE    IndicatorParamValueType = 6 // 线型
	IndicatorParamValueType_BOOL    IndicatorParamValueType = 7 // 布尔
)

// =============================================================================
// IndicatorShape (指标形状)
// =============================================================================

// IndicatorShape represents the shape of an indicator point
type IndicatorShape int32

const (
	IndicatorShape_Unknown        IndicatorShape = 0  // 未知
	IndicatorShape_XCross         IndicatorShape = 1  // X型交叉
	IndicatorShape_Cross          IndicatorShape = 2  // 十字
	IndicatorShape_Circle         IndicatorShape = 3  // 圆形
	IndicatorShape_TriangleUp     IndicatorShape = 4  // 三角向上
	IndicatorShape_TriangleDown   IndicatorShape = 5  // 三角向下
	IndicatorShape_Flag           IndicatorShape = 6  // 旗帜
	IndicatorShape_ArrowUp        IndicatorShape = 7  // 箭头向上
	IndicatorShape_ArrowDown      IndicatorShape = 8  // 箭头向下
	IndicatorShape_Square         IndicatorShape = 9  // 方形
	IndicatorShape_Diamond        IndicatorShape = 10 // 菱形
	IndicatorShape_LabelUp        IndicatorShape = 11 // 标签向上
	IndicatorShape_LabelDown      IndicatorShape = 12 // 标签向下
)

// =============================================================================
// IndicatorLineType (指标线型)
// =============================================================================

// IndicatorLineType represents the type of indicator line
type IndicatorLineType int32

const (
	IndicatorLineType_Unknown        IndicatorLineType = 0 // 未知
	IndicatorLineType_Solid          IndicatorLineType = 1 // 实线
	IndicatorLineType_Dashed         IndicatorLineType = 2 // 虚线
	IndicatorLineType_Dot            IndicatorLineType = 3 // 点线
	IndicatorLineType_Cross          IndicatorLineType = 4 // 十字
	IndicatorLineType_Circle         IndicatorLineType = 5 // 圆形
	IndicatorLineType_Histogram      IndicatorLineType = 6 // 柱状图
	IndicatorLineType_HistogramLine  IndicatorLineType = 7 // 柱状线
	IndicatorLineType_Step           IndicatorLineType = 8 // 阶梯
	IndicatorLineType_StepDiamonds   IndicatorLineType = 9 // 阶梯菱形
)

// =============================================================================
// NewsSubType (新闻子类型)
// =============================================================================

// NewsSubType represents the sub-type of news
type NewsSubType int32

const (
	NewsSubType_All    NewsSubType = 0 // 所有类型
	NewsSubType_News   NewsSubType = 1 // 新闻
	NewsSubType_Notice NewsSubType = 2 // 公告
	NewsSubType_Rating NewsSubType = 3 // 评级
)

// =============================================================================
// OptionMarket (期权市场)
// =============================================================================

// OptionMarket represents the option market
type OptionMarket int32

const (
	OptionMarket_Unknown    OptionMarket = 0 // 未知
	OptionMarket_US_Security OptionMarket = 1 // 美股期权
	OptionMarket_US_Index   OptionMarket = 2 // 美股指数期权
	OptionMarket_HK_Security OptionMarket = 3 // 港股期权
	OptionMarket_HK_Index   OptionMarket = 4 // 港股指数期权
)

// =============================================================================
// OptionStatisticDataType (期权统计数据类型)
// =============================================================================

// OptionStatisticDataType represents the type of option statistic data
type OptionStatisticDataType int32

const (
	OptionStatisticDataType_Volume       OptionStatisticDataType = 0 // 成交量
	OptionStatisticDataType_OpenInterest OptionStatisticDataType = 1 // 持仓量
)

// =============================================================================
// OptionHVTimeRange (期权历史波动率时间范围)
// =============================================================================

// OptionHVTimeRange represents the time range for option historical volatility
type OptionHVTimeRange int32

const (
	OptionHVTimeRange_30  OptionHVTimeRange = 0 // 30天
	OptionHVTimeRange_60  OptionHVTimeRange = 1 // 60天
	OptionHVTimeRange_90  OptionHVTimeRange = 2 // 90天
	OptionHVTimeRange_120 OptionHVTimeRange = 3 // 120天
	OptionHVTimeRange_365 OptionHVTimeRange = 4 // 365天
)

// =============================================================================
// UnderlyingRankSortType (标的排名排序类型)
// =============================================================================

// UnderlyingRankSortType represents the sort type for underlying ranking
type UnderlyingRankSortType int32

const (
	UnderlyingRankSortType_Unknown                  UnderlyingRankSortType = 0  // 未知
	UnderlyingRankSortType_Volume                   UnderlyingRankSortType = 1  // 成交量
	UnderlyingRankSortType_OpenInterest             UnderlyingRankSortType = 2  // 持仓量
	UnderlyingRankSortType_IV                       UnderlyingRankSortType = 3  // 隐含波动率
	UnderlyingRankSortType_HV                       UnderlyingRankSortType = 4  // 历史波动率
	UnderlyingRankSortType_IVRank                   UnderlyingRankSortType = 5  // IV分位
	UnderlyingRankSortType_IVPercentile             UnderlyingRankSortType = 6  // IV百分位
	UnderlyingRankSortType_IVChange                 UnderlyingRankSortType = 7  // IV变化
	UnderlyingRankSortType_HVChange                 UnderlyingRankSortType = 8  // HV变化
	UnderlyingRankSortType_VolumeRatio              UnderlyingRankSortType = 9  // 量比
	UnderlyingRankSortType_OIRatio                  UnderlyingRankSortType = 10 // 持仓量比
	UnderlyingRankSortType_MarketCap                UnderlyingRankSortType = 11 // 市值
	UnderlyingRankSortType_ChangeRatio              UnderlyingRankSortType = 12 // 涨跌幅
	UnderlyingRankSortType_Price                    UnderlyingRankSortType = 13 // 价格
)

// =============================================================================
// OptionRankType (期权排名类型)
// =============================================================================

// OptionRankType represents the type of option contract ranking
type OptionRankType int32

const (
	OptionRankType_Unknown         OptionRankType = 0  // 未知
	OptionRankType_Volume          OptionRankType = 1  // 成交量
	OptionRankType_OpenInterest    OptionRankType = 2  // 持仓量
	OptionRankType_VolumeChange    OptionRankType = 3  // 成交量变化
	OptionRankType_OIChange        OptionRankType = 4  // 持仓量变化
	OptionRankType_IV              OptionRankType = 5  // 隐含波动率
	OptionRankType_HV              OptionRankType = 6  // 历史波动率
	OptionRankType_IVRank          OptionRankType = 7  // IV分位
	OptionRankType_IVPercentile    OptionRankType = 8  // IV百分位
	OptionRankType_PriceChange     OptionRankType = 9  // 价格变化
	OptionRankType_Premium         OptionRankType = 10 // 权利金
)

// =============================================================================
// ZeroDteSortType (0DTE排序类型)
// =============================================================================

// ZeroDteSortType represents the sort type for zero DTE options
type ZeroDteSortType int32

const (
	ZeroDteSortType_Unknown     ZeroDteSortType = 0 // 未知
	ZeroDteSortType_Volume      ZeroDteSortType = 1 // 成交量
	ZeroDteSortType_IV          ZeroDteSortType = 2 // 隐含波动率
	ZeroDteSortType_ChangeRatio ZeroDteSortType = 3 // 涨跌幅
	ZeroDteSortType_OpenInterest ZeroDteSortType = 4 // 持仓量
	ZeroDteSortType_MarketCap   ZeroDteSortType = 5 // 市值
)

// =============================================================================
// ZeroDteIndicatorType (0DTE筛选指标类型)
// =============================================================================

// ZeroDteIndicatorType represents the indicator type for zero DTE screening
type ZeroDteIndicatorType int32

const (
	ZeroDteIndicatorType_Unknown          ZeroDteIndicatorType = 0  // 未知
	ZeroDteIndicatorType_OwnerList        ZeroDteIndicatorType = 1  // 持有列表
	ZeroDteIndicatorType_HasEarningsWeek  ZeroDteIndicatorType = 2  // 本周有财报
	ZeroDteIndicatorType_Volume           ZeroDteIndicatorType = 3  // 成交量
	ZeroDteIndicatorType_OpenInterest     ZeroDteIndicatorType = 4  // 持仓量
	ZeroDteIndicatorType_IV               ZeroDteIndicatorType = 5  // 隐含波动率
	ZeroDteIndicatorType_HV               ZeroDteIndicatorType = 6  // 历史波动率
	ZeroDteIndicatorType_IVRank           ZeroDteIndicatorType = 7  // IV分位
	ZeroDteIndicatorType_IVPercentile     ZeroDteIndicatorType = 8  // IV百分位
	ZeroDteIndicatorType_Price            ZeroDteIndicatorType = 9  // 价格
	ZeroDteIndicatorType_ChangeRatio      ZeroDteIndicatorType = 10 // 涨跌幅
)

// =============================================================================
// ZeroDteContractSortType (0DTE合约排序类型)
// =============================================================================

// ZeroDteContractSortType represents the sort type for zero DTE contracts
type ZeroDteContractSortType int32

const (
	ZeroDteContractSortType_Unknown     ZeroDteContractSortType = 0 // 未知
	ZeroDteContractSortType_Volume      ZeroDteContractSortType = 1 // 成交量
	ZeroDteContractSortType_OpenInterest ZeroDteContractSortType = 2 // 持仓量
	ZeroDteContractSortType_IV          ZeroDteContractSortType = 3 // 隐含波动率
	ZeroDteContractSortType_Delta       ZeroDteContractSortType = 4 // Delta
)

// =============================================================================
// ZeroDteContractIndicatorType (0DTE合约筛选指标类型)
// =============================================================================

// ZeroDteContractIndicatorType represents the indicator type for zero DTE contracts
type ZeroDteContractIndicatorType int32

const (
	ZeroDteContractIndicatorType_Unknown             ZeroDteContractIndicatorType = 0  // 未知
	ZeroDteContractIndicatorType_OptionType          ZeroDteContractIndicatorType = 1  // 期权类型
	ZeroDteContractIndicatorType_Volume              ZeroDteContractIndicatorType = 2  // 成交量
	ZeroDteContractIndicatorType_OpenInterest        ZeroDteContractIndicatorType = 3  // 持仓量
	ZeroDteContractIndicatorType_IV                  ZeroDteContractIndicatorType = 4  // 隐含波动率
	ZeroDteContractIndicatorType_Delta               ZeroDteContractIndicatorType = 5  // Delta
	ZeroDteContractIndicatorType_Gamma               ZeroDteContractIndicatorType = 6  // Gamma
	ZeroDteContractIndicatorType_Theta               ZeroDteContractIndicatorType = 7  // Theta
	ZeroDteContractIndicatorType_Vega                ZeroDteContractIndicatorType = 8  // Vega
	ZeroDteContractIndicatorType_Rho                 ZeroDteContractIndicatorType = 9  // Rho
	ZeroDteContractIndicatorType_Price               ZeroDteContractIndicatorType = 10 // 价格
	ZeroDteContractIndicatorType_ChangeRatio         ZeroDteContractIndicatorType = 11 // 涨跌幅
	ZeroDteContractIndicatorType_BreakEvenPoint      ZeroDteContractIndicatorType = 12 // 盈亏平衡点
	ZeroDteContractIndicatorType_ToBEP               ZeroDteContractIndicatorType = 13 // 距离盈亏平衡点
	ZeroDteContractIndicatorType_BuyProfitProb       ZeroDteContractIndicatorType = 14 // 买入盈利概率
	ZeroDteContractIndicatorType_SellProfitProb      ZeroDteContractIndicatorType = 15 // 卖出盈利概率
)

// =============================================================================
// EarningsSortType (财报扫描排序类型)
// =============================================================================

// EarningsSortType represents the sort type for earnings screening
type EarningsSortType int32

const (
	EarningsSortType_Unknown              EarningsSortType = 0  // 未知
	EarningsSortType_EarningsDate         EarningsSortType = 1  // 财报日期
	EarningsSortType_Volume               EarningsSortType = 2  // 成交量
	EarningsSortType_IV                   EarningsSortType = 3  // 隐含波动率
	EarningsSortType_MarketCap            EarningsSortType = 4  // 市值
	EarningsSortType_ChangeRatio          EarningsSortType = 5  // 涨跌幅
	EarningsSortType_Price                EarningsSortType = 6  // 价格
	EarningsSortType_IVRank               EarningsSortType = 7  // IV分位
	EarningsSortType_IVPercentile         EarningsSortType = 8  // IV百分位
	EarningsSortType_HV                   EarningsSortType = 9  // 历史波动率
	EarningsSortType_OpenInterest         EarningsSortType = 10 // 持仓量
	EarningsSortType_LastReportIVCrush    EarningsSortType = 11 // 上次财报IV平仓
	EarningsSortType_HistoryReportIVCrush EarningsSortType = 12 // 历史财报IV平仓
	EarningsSortType_LastReportChgRatio   EarningsSortType = 13 // 上次财报涨跌幅
	EarningsSortType_HistoryReportChgRatio EarningsSortType = 14 // 历史财报涨跌幅
	EarningsSortType_EstimateEPSYoY       EarningsSortType = 15 // 预估EPS同比
	EarningsSortType_EstimateRevenueYoY   EarningsSortType = 16 // 预估营收同比
	EarningsSortType_ExpectedMoveRatio    EarningsSortType = 17 // 预期波动幅度
)

// =============================================================================
// StockCategory (股票分类)
// =============================================================================

// StockCategory represents the category of stock for screening
type StockCategory int32

const (
	StockCategory_All    StockCategory = 0 // 所有
	StockCategory_Equity StockCategory = 1 // 正股
	StockCategory_ETF    StockCategory = 2 // ETF
)

// =============================================================================
// IndexComponent (指数成分股)
// =============================================================================

// IndexComponent represents a major stock index component filter
type IndexComponent int32

const (
	IndexComponent_Unknown IndexComponent = 0 // 未知
	IndexComponent_DJI     IndexComponent = 1 // 道琼斯
	IndexComponent_IXIC    IndexComponent = 2 // 纳斯达克综合
	IndexComponent_NDX     IndexComponent = 3 // 纳斯达克100
	IndexComponent_SPX     IndexComponent = 4 // 标普500
)

// =============================================================================
// ExpirationType (到期类型)
// =============================================================================

// ExpirationType represents the expiration type for options
type ExpirationType int32

const (
	ExpirationType_Unknown     ExpirationType = 0 // 未知
	ExpirationType_Monthly     ExpirationType = 1 // 月度
	ExpirationType_Weekly      ExpirationType = 2 // 周度
	ExpirationType_EndOfMonth  ExpirationType = 3 // 月末
	ExpirationType_Quarterly   ExpirationType = 4 // 季度
)

// =============================================================================
// EarningsPubType (财报发布时间)
// =============================================================================

// EarningsPubType represents the publication time type for earnings
type EarningsPubType int32

const (
	EarningsPubType_Unknown EarningsPubType = 0 // 未知
	EarningsPubType_Before  EarningsPubType = 1 // 盘前
	EarningsPubType_After   EarningsPubType = 2 // 盘后
)

// =============================================================================
// SellerType (期权卖方策略类型)
// =============================================================================

// SellerType represents the option seller strategy type
type SellerType int32

const (
	SellerType_Unknown      SellerType = 0 // 未知
	SellerType_CoveredCall   SellerType = 1 // 备兑开仓
	SellerType_CashSecuredPut SellerType = 2 // 现金担保看跌
)

// =============================================================================
// SellerSortType (期权卖方排序类型)
// =============================================================================

// SellerSortType represents the sort type for option seller screening
type SellerSortType int32

const (
	SellerSortType_Unknown          SellerSortType = 0 // 未知
	SellerSortType_AnnualizedReturn SellerSortType = 1 // 年化收益率
	SellerSortType_IntervalReturn   SellerSortType = 2 // 区间收益率
	SellerSortType_ITMProbability   SellerSortType = 3 // ITM概率
	SellerSortType_Premium          SellerSortType = 4 // 权利金
)

// =============================================================================
// EarningsCalendarSortType (财报日历排序类型)
// =============================================================================

// EarningsCalendarSortType represents the sort type for earnings calendar
type EarningsCalendarSortType int32

const (
	EarningsCalendarSortType_Unknown      EarningsCalendarSortType = 0 // 未知
	EarningsCalendarSortType_Hot          EarningsCalendarSortType = 1 // 热度
	EarningsCalendarSortType_MarketCap    EarningsCalendarSortType = 2 // 市值
	EarningsCalendarSortType_OptionVolume EarningsCalendarSortType = 3 // 期权成交量
	EarningsCalendarSortType_IV           EarningsCalendarSortType = 4 // 隐含波动率
	EarningsCalendarSortType_IVRank       EarningsCalendarSortType = 5 // IV分位
	EarningsCalendarSortType_IVPercentile EarningsCalendarSortType = 6 // IV百分位
	EarningsCalendarSortType_RTMarketCap  EarningsCalendarSortType = 7 // 实时市值
)

// =============================================================================
// EarningsCalendarPubType (财报日历发布时间类型)
// =============================================================================

// EarningsCalendarPubType represents the publication type for earnings calendar
type EarningsCalendarPubType int32

const (
	EarningsCalendarPubType_Unknown EarningsCalendarPubType = 0 // 未知
	EarningsCalendarPubType_Regular EarningsCalendarPubType = 1 // 常规
	EarningsCalendarPubType_Before  EarningsCalendarPubType = 2 // 盘前
	EarningsCalendarPubType_After   EarningsCalendarPubType = 3 // 盘后
)

// =============================================================================
// EarningsCalendarEstimateType (财报日历预估类型)
// =============================================================================

// EarningsCalendarEstimateType represents the estimate type for earnings calendar
type EarningsCalendarEstimateType int32

const (
	EarningsCalendarEstimateType_Unknown  EarningsCalendarEstimateType = 0 // 未知
	EarningsCalendarEstimateType_EPS      EarningsCalendarEstimateType = 1 // 每股收益
	EarningsCalendarEstimateType_Revenue  EarningsCalendarEstimateType = 2 // 营收
	EarningsCalendarEstimateType_EBIT     EarningsCalendarEstimateType = 3 // 息税前利润
)

// =============================================================================
// EarningsCalendarPeriodType (财报日历周期类型)
// =============================================================================

// EarningsCalendarPeriodType represents the period type for earnings calendar
type EarningsCalendarPeriodType int32

const (
	EarningsCalendarPeriodType_Unknown     EarningsCalendarPeriodType = 0 // 未知
	EarningsCalendarPeriodType_Quarterly   EarningsCalendarPeriodType = 1 // 季度
	EarningsCalendarPeriodType_SemiAnnual  EarningsCalendarPeriodType = 2 // 半年度
	EarningsCalendarPeriodType_Annual      EarningsCalendarPeriodType = 3 // 年度
)

// =============================================================================
// EarningsCalendarStockListType (财报日历股票列表类型)
// =============================================================================

// EarningsCalendarStockListType represents the stock list type for earnings calendar
type EarningsCalendarStockListType int32

const (
	EarningsCalendarStockListType_Unknown  EarningsCalendarStockListType = 0 // 未知
	EarningsCalendarStockListType_Watchlist EarningsCalendarStockListType = 1 // 自选股
	EarningsCalendarStockListType_Position EarningsCalendarStockListType = 2 // 持仓
	EarningsCalendarStockListType_Special  EarningsCalendarStockListType = 3 // 特别关注
)

// =============================================================================
// MacroRegion (宏观区域)
// =============================================================================

// MacroRegion represents a region for macroeconomic data
type MacroRegion int32

const (
	MacroRegion_Unknown MacroRegion = 0 // 未知
	MacroRegion_HK      MacroRegion = 1 // 香港
	MacroRegion_US      MacroRegion = 2 // 美国
	MacroRegion_JP      MacroRegion = 3 // 日本
	MacroRegion_SG      MacroRegion = 4 // 新加坡
	MacroRegion_AU      MacroRegion = 5 // 澳大利亚
	MacroRegion_CA      MacroRegion = 6 // 加拿大
	MacroRegion_MY      MacroRegion = 7 // 马来西亚
	MacroRegion_CN      MacroRegion = 8 // 中国
)

// =============================================================================
// MacroDataUnitType (宏观数据单位类型)
// =============================================================================

// MacroDataUnitType represents the unit type for macroeconomic data
type MacroDataUnitType int32

const (
	MacroDataUnitType_Unknown MacroDataUnitType = 0 // 未知
	MacroDataUnitType_Percent MacroDataUnitType = 1 // 百分比
	MacroDataUnitType_Value   MacroDataUnitType = 2 // 值
	MacroDataUnitType_Index   MacroDataUnitType = 3 // 指数
)

// =============================================================================
// BeatType (超预期类型)
// =============================================================================

// BeatType represents the type of earnings beat
type BeatType int32

const (
	BeatType_Unknown BeatType = 0 // 未知
	BeatType_EPS     BeatType = 1 // 每股收益
	BeatType_Revenue BeatType = 2 // 营收
	BeatType_EBIT    BeatType = 3 // 息税前利润
)

// =============================================================================
// BeatTerm (财报超预期周期)
// =============================================================================

// BeatTerm represents the term for earnings beat
type BeatTerm int32

const (
	BeatTerm_Latest         BeatTerm = 0 // 最新
	BeatTerm_LatestQuarter  BeatTerm = 1 // 最新季度
	BeatTerm_LatestHalf     BeatTerm = 2 // 最新半年度
	BeatTerm_LatestAnnual   BeatTerm = 3 // 最新年度
	BeatTerm_All            BeatTerm = 4 // 全部
)

// =============================================================================
// PostPeriodType (财报后周期类型)
// =============================================================================

// PostPeriodType represents the post-earnings period type
type PostPeriodType int32

const (
	PostPeriodType_Unknown        PostPeriodType = 0 // 未知
	PostPeriodType_Regular        PostPeriodType = 1 // 常规
	PostPeriodType_Before         PostPeriodType = 2 // 盘前
	PostPeriodType_After          PostPeriodType = 3 // 盘后
	PostPeriodType_IntradayTrading PostPeriodType = 4 // 盘中交易
)

// =============================================================================
// DistributionFrequency (分红频率)
// =============================================================================

// DistributionFrequency represents the frequency of dividend distribution
type DistributionFrequency int32

const (
	DistributionFrequency_Unknown    DistributionFrequency = 0 // 未知
	DistributionFrequency_Annual     DistributionFrequency = 1 // 年度
	DistributionFrequency_SemiAnnual DistributionFrequency = 2 // 半年度
	DistributionFrequency_Quarterly  DistributionFrequency = 3 // 季度
	DistributionFrequency_Monthly    DistributionFrequency = 4 // 月度
)

// =============================================================================
// DividendRankType (股息排名类型)
// =============================================================================

// DividendRankType represents the type of dividend ranking
type DividendRankType int32

const (
	DividendRankType_Unknown      DividendRankType = 0 // 未知
	DividendRankType_HighYield    DividendRankType = 1 // 高股息
	DividendRankType_DividendGrowth DividendRankType = 2 // 股息增长
)

// =============================================================================
// EconomicImportance (经济数据重要性)
// =============================================================================

// EconomicImportance represents the importance level of economic data
type EconomicImportance int32

const (
	EconomicImportance_All    EconomicImportance = 0 // 全部
	EconomicImportance_Low    EconomicImportance = 1 // 低
	EconomicImportance_Medium EconomicImportance = 2 // 中
	EconomicImportance_High   EconomicImportance = 3 // 高
)

// =============================================================================
// RankSortDir (排名排序方向)
// =============================================================================

// RankSortDir represents the sort direction for rankings
type RankSortDir int32

const (
	RankSortDir_Descending RankSortDir = 0 // 降序
	RankSortDir_Ascending  RankSortDir = 1 // 升序
)

// =============================================================================
// SimpleRankIndicatorType (简单排名指标类型)
// =============================================================================

// SimpleRankIndicatorType represents the indicator type for simple rankings
type SimpleRankIndicatorType int32

const (
	SimpleRankIndicatorType_Unknown   SimpleRankIndicatorType = 0 // 未知
	SimpleRankIndicatorType_Price     SimpleRankIndicatorType = 1 // 价格
	SimpleRankIndicatorType_MarketCap SimpleRankIndicatorType = 2 // 市值
	SimpleRankIndicatorType_PE        SimpleRankIndicatorType = 3 // 市盈率
)

// =============================================================================
// PriceFilter (价格筛选)
// =============================================================================

// PriceFilter represents a price range filter for rankings
type PriceFilter int32

const (
	PriceFilter_All               PriceFilter = 0 // 全部
	PriceFilter_LessThan1         PriceFilter = 1 // 小于1
	PriceFilter_Between1And10     PriceFilter = 2 // 1-10
	PriceFilter_Between10And100   PriceFilter = 3 // 10-100
	PriceFilter_GreaterThan100    PriceFilter = 4 // 大于100
	PriceFilter_Near52WeekHigh    PriceFilter = 5 // 接近52周高点
	PriceFilter_Near52WeekLow     PriceFilter = 6 // 接近52周低点
)

// =============================================================================
// HotListSortField (热门榜单排序字段)
// =============================================================================

// HotListSortField represents the sort field for hot list
type HotListSortField int32

const (
	HotListSortField_Unknown    HotListSortField = 0 // 未知
	HotListSortField_TradeHeat  HotListSortField = 1 // 交易热度
	HotListSortField_SearchHeat HotListSortField = 2 // 搜索热度
	HotListSortField_NewsHeat   HotListSortField = 3 // 新闻热度
	HotListSortField_AverageHeat HotListSortField = 4 // 平均热度
)

// =============================================================================
// ShortSellingSortField (做空排名排序字段)
// =============================================================================

// ShortSellingSortField represents the sort field for short selling ranking
type ShortSellingSortField int32

const (
	ShortSellingSortField_Unknown            ShortSellingSortField = 0  // 未知
	ShortSellingSortField_ShortNumberChange  ShortSellingSortField = 1  // 做空数量变化
	ShortSellingSortField_ShortRatioChange   ShortSellingSortField = 2  // 做空比例变化
	ShortSellingSortField_ShortNumber        ShortSellingSortField = 3  // 做空数量
	ShortSellingSortField_ShortRatio         ShortSellingSortField = 4  // 做空比例
	ShortSellingSortField_Volume             ShortSellingSortField = 5  // 成交量
	ShortSellingSortField_PositionVolume     ShortSellingSortField = 6  // 持仓量
	ShortSellingSortField_PositionRatio      ShortSellingSortField = 7  // 持仓占比
	ShortSellingSortField_DaysToCover        ShortSellingSortField = 8  // 回补天数
	ShortSellingSortField_WeekAvgVolume      ShortSellingSortField = 9  // 周均成交量
	ShortSellingSortField_WeekAvgShortNumber ShortSellingSortField = 10 // 周均做空量
	ShortSellingSortField_WeekAvgShortRatio  ShortSellingSortField = 11 // 周均做空比
	ShortSellingSortField_MonthAvgVolume     ShortSellingSortField = 12 // 月均成交量
	ShortSellingSortField_MonthAvgShortNumber ShortSellingSortField = 13 // 月均做空量
	ShortSellingSortField_MonthAvgShortRatio  ShortSellingSortField = 14 // 月均做空比
)

// =============================================================================
// RankPeriodType (排名周期类型)
// =============================================================================

// RankPeriodType represents the period type for period change ranking
type RankPeriodType int32

const (
	RankPeriodType_Unknown  RankPeriodType = 0 // 未知
	RankPeriodType_5Min    RankPeriodType = 1 // 5分钟
	RankPeriodType_1Day    RankPeriodType = 2 // 1日
	RankPeriodType_5Day    RankPeriodType = 3 // 5日
	RankPeriodType_20Day   RankPeriodType = 4 // 20日
	RankPeriodType_60Day   RankPeriodType = 5 // 60日
	RankPeriodType_120Day  RankPeriodType = 6 // 120日
	RankPeriodType_250Day  RankPeriodType = 7 // 250日
	RankPeriodType_YTD     RankPeriodType = 8 // 年初至今
)

// =============================================================================
// InstitutionListSortField (机构列表排序字段)
// =============================================================================

// InstitutionListSortField represents the sort field for institution list
type InstitutionListSortField int32

const (
	InstitutionListSortField_Unknown            InstitutionListSortField = 0 // 未知
	InstitutionListSortField_PositionValue      InstitutionListSortField = 1 // 持仓市值
	InstitutionListSortField_PositionValueChange InstitutionListSortField = 2 // 持仓市值变化
	InstitutionListSortField_PositionCount      InstitutionListSortField = 3 // 持仓数量
	InstitutionListSortField_PositionCountChange InstitutionListSortField = 4 // 持仓数量变化
)

// =============================================================================
// InstitutionHoldingChangeType (机构持仓变动类型)
// =============================================================================

// InstitutionHoldingChangeType represents the type of institution holding change
type InstitutionHoldingChangeType int32

const (
	InstitutionHoldingChangeType_Unknown  InstitutionHoldingChangeType = 0 // 未知
	InstitutionHoldingChangeType_New      InstitutionHoldingChangeType = 1 // 新增
	InstitutionHoldingChangeType_SoldOut  InstitutionHoldingChangeType = 2 // 清仓
	InstitutionHoldingChangeType_Increase InstitutionHoldingChangeType = 3 // 增持
	InstitutionHoldingChangeType_Decrease InstitutionHoldingChangeType = 4 // 减持
)

// =============================================================================
// ArkHoldingType (ARK持仓类型)
// =============================================================================

// ArkHoldingType represents the type of ARK fund holding
type ArkHoldingType int32

const (
	ArkHoldingType_Position ArkHoldingType = 0 // 持仓
	ArkHoldingType_Increase ArkHoldingType = 1 // 增持
	ArkHoldingType_Decrease ArkHoldingType = 2 // 减持
	ArkHoldingType_New      ArkHoldingType = 3 // 新增
	ArkHoldingType_SoldOut  ArkHoldingType = 4 // 清仓
)

// =============================================================================
// ArkCycleType (ARK周期类型)
// =============================================================================

// ArkCycleType represents the cycle type for ARK data
type ArkCycleType int32

const (
	ArkCycleType_1Day   ArkCycleType = 0 // 1日
	ArkCycleType_5Day   ArkCycleType = 1 // 5日
	ArkCycleType_10Day  ArkCycleType = 2 // 10日
	ArkCycleType_30Day  ArkCycleType = 3 // 30日
	ArkCycleType_60Day  ArkCycleType = 4 // 60日
)

// =============================================================================
// ArkDynamicType (ARK动态类型)
// =============================================================================

// ArkDynamicType represents the type of ARK stock dynamic
type ArkDynamicType int32

const (
	ArkDynamicType_Unknown               ArkDynamicType = 0 // 未知
	ArkDynamicType_ConsecutiveSameDir    ArkDynamicType = 1 // 连续同向
	ArkDynamicType_RecentTransaction     ArkDynamicType = 2 // 近期交易
	ArkDynamicType_LastTransaction       ArkDynamicType = 3 // 最后交易
	ArkDynamicType_NoDynamic             ArkDynamicType = 4 // 无动态
)

// =============================================================================
// RatingChangeType (评级变动类型)
// =============================================================================

// RatingChangeType represents the type of rating change
type RatingChangeType int32

const (
	RatingChangeType_Unknown   RatingChangeType = 0 // 未知
	RatingChangeType_Upgrade   RatingChangeType = 1 // 上调
	RatingChangeType_Downgrade RatingChangeType = 2 // 下调
	RatingChangeType_NewRating RatingChangeType = 3 // 新增评级
)

// =============================================================================
// RatingLevel (评级等级)
// =============================================================================

// RatingLevel represents the level of a rating
type RatingLevel int32

const (
	RatingLevel_Unknown RatingLevel = 0 // 未知
	RatingLevel_Sell    RatingLevel = 1 // 卖出
	RatingLevel_Hold    RatingLevel = 2 // 持有
	RatingLevel_Buy     RatingLevel = 3 // 买入
)

// =============================================================================
// IndustrialChainType (产业链类型)
// =============================================================================

// IndustrialChainType represents the type of industrial chain
type IndustrialChainType int32

const (
	IndustrialChainType_Unknown   IndustrialChainType = 0 // 未知
	IndustrialChainType_Chain     IndustrialChainType = 1 // 产业链
	IndustrialChainType_Parallel  IndustrialChainType = 2 // 平行
	IndustrialChainType_UpMidDown IndustrialChainType = 3 // 上中下游
)

// =============================================================================
// PlateStockSortField (板块股票排序字段)
// =============================================================================

// PlateStockSortField represents the sort field for plate stocks
type PlateStockSortField int32

const (
	PlateStockSortField_Unknown    PlateStockSortField = 0 // 未知
	PlateStockSortField_Code       PlateStockSortField = 1 // 代码
	PlateStockSortField_ChangeRate PlateStockSortField = 2 // 涨跌幅
	PlateStockSortField_Turnover   PlateStockSortField = 3 // 成交额
	PlateStockSortField_Volume     PlateStockSortField = 4 // 成交量
	PlateStockSortField_MarketVal  PlateStockSortField = 5 // 市值
)

// =============================================================================
// HeatMapSortField (热力图排序字段)
// =============================================================================

// HeatMapSortField represents the sort field for heat map data
type HeatMapSortField int32

const (
	HeatMapSortField_Unknown    HeatMapSortField = 0 // 未知
	HeatMapSortField_ChangeRate HeatMapSortField = 1 // 涨跌幅
	HeatMapSortField_MarketVal  HeatMapSortField = 2 // 市值
	HeatMapSortField_Turnover   HeatMapSortField = 3 // 成交额
	HeatMapSortField_Hot        HeatMapSortField = 4 // 热度
)

// =============================================================================
// HeatMapPlateType (热力图板块类型)
// =============================================================================

// HeatMapPlateType represents the plate type for heat map
type HeatMapPlateType int32

const (
	HeatMapPlateType_Industry HeatMapPlateType = 0 // 行业
	HeatMapPlateType_Concept  HeatMapPlateType = 1 // 概念
	HeatMapPlateType_Theme    HeatMapPlateType = 2 // 主题
)

// =============================================================================
// RiseFallDistributionType (涨跌分布类型)
// =============================================================================

// RiseFallDistributionType represents the type of rise/fall distribution
type RiseFallDistributionType int32

const (
	RiseFallDistributionType_Unknown           RiseFallDistributionType = 0 // 未知
	RiseFallDistributionType_RiseLimit         RiseFallDistributionType = 1 // 涨停
	RiseFallDistributionType_PositiveInfinity  RiseFallDistributionType = 2 // 涨幅7%以上
	RiseFallDistributionType_NormalRange       RiseFallDistributionType = 3 // 正常区间
	RiseFallDistributionType_NegativeInfinity  RiseFallDistributionType = 4 // 跌幅7%以上
	RiseFallDistributionType_FallLimit         RiseFallDistributionType = 5 // 跌停
)
