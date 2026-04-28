package constant

func (m Market) String() string {
	switch m {
	case Market_None:
		return "Market_None"
	case Market_HK:
		return "Market_HK"
	case Market_US:
		return "Market_US"
	case Market_SH:
		return "Market_SH"
	case Market_SZ:
		return "Market_SZ"
	case Market_SG:
		return "Market_SG"
	case Market_JP:
		return "Market_JP"
	case Market_AU:
		return "Market_AU"
	case Market_MY:
		return "Market_MY"
	case Market_CA:
		return "Market_CA"
	case Market_FX:
		return "Market_FX"
	default:
		return "Market_Unknown"
	}
}

// Prefix returns the short market prefix string (e.g. "HK", "US", "SH", "SZ").
// Returns "UNKNOWN" for unrecognized values.
func (m Market) Prefix() string {
	switch m {
	case Market_HK:
		return "HK"
	case Market_US:
		return "US"
	case Market_SH:
		return "SH"
	case Market_SZ:
		return "SZ"
	case Market_SG:
		return "SG"
	case Market_JP:
		return "JP"
	case Market_AU:
		return "AU"
	case Market_MY:
		return "MY"
	case Market_CA:
		return "CA"
	case Market_FX:
		return "FX"
	default:
		return "UNKNOWN"
	}
}

func (s SecurityType) String() string {
	switch s {
	case SecurityType_None:
		return "SecurityType_None"
	case SecurityType_Bond:
		return "SecurityType_Bond"
	case SecurityType_Bwrt:
		return "SecurityType_Bwrt"
	case SecurityType_Stock:
		return "SecurityType_Stock"
	case SecurityType_ETF:
		return "SecurityType_ETF"
	case SecurityType_Warrant:
		return "SecurityType_Warrant"
	case SecurityType_Index:
		return "SecurityType_Index"
	case SecurityType_Plate:
		return "SecurityType_Plate"
	case SecurityType_Drvt:
		return "SecurityType_Drvt"
	case SecurityType_PlateSet:
		return "SecurityType_PlateSet"
	case SecurityType_Future:
		return "SecurityType_Future"
	case SecurityType_Forex:
		return "SecurityType_Forex"
	default:
		return "SecurityType_Unknown"
	}
}

func (s SubType) String() string {
	switch s {
	case SubType_None:
		return "SubType_None"
	case SubType_Quote:
		return "SubType_Quote"
	case SubType_OrderBook:
		return "SubType_OrderBook"
	case SubType_Ticker:
		return "SubType_Ticker"
	case SubType_Broker:
		return "SubType_Broker"
	case SubType_RT:
		return "SubType_RT"
	case SubType_K_1Min:
		return "SubType_K_1Min"
	case SubType_K_3Min:
		return "SubType_K_3Min"
	case SubType_K_5Min:
		return "SubType_K_5Min"
	case SubType_K_15Min:
		return "SubType_K_15Min"
	case SubType_K_30Min:
		return "SubType_K_30Min"
	case SubType_K_60Min:
		return "SubType_K_60Min"
	case SubType_K_Day:
		return "SubType_K_Day"
	case SubType_K_Week:
		return "SubType_K_Week"
	case SubType_K_Month:
		return "SubType_K_Month"
	case SubType_K_Quarter:
		return "SubType_K_Quarter"
	case SubType_K_Year:
		return "SubType_K_Year"
	default:
		return "SubType_Unknown"
	}
}

func (k KLType) String() string {
	switch k {
	case KLType_None:
		return "KLType_None"
	case KLType_K_1Min:
		return "KLType_K_1Min"
	case KLType_K_3Min:
		return "KLType_K_3Min"
	case KLType_K_5Min:
		return "KLType_K_5Min"
	case KLType_K_15Min:
		return "KLType_K_15Min"
	case KLType_K_30Min:
		return "KLType_K_30Min"
	case KLType_K_60Min:
		return "KLType_K_60Min"
	case KLType_K_Day:
		return "KLType_K_Day"
	case KLType_K_Week:
		return "KLType_K_Week"
	case KLType_K_Month:
		return "KLType_K_Month"
	case KLType_K_Quarter:
		return "KLType_K_Quarter"
	case KLType_K_Year:
		return "KLType_K_Year"
	default:
		return "KLType_Unknown"
	}
}

func (r RehabType) String() string {
	switch r {
	case RehabType_None:
		return "RehabType_None"
	case RehabType_Forward:
		return "RehabType_Forward"
	case RehabType_Backward:
		return "RehabType_Backward"
	default:
		return "RehabType_Unknown"
	}
}

func (t TrdEnv) String() string {
	switch t {
	case TrdEnv_Simulate:
		return "TrdEnv_Simulate"
	case TrdEnv_Real:
		return "TrdEnv_Real"
	default:
		return "TrdEnv_Unknown"
	}
}

func (t TrdMarket) String() string {
	switch t {
	case TrdMarket_None:
		return "TrdMarket_None"
	case TrdMarket_HK:
		return "TrdMarket_HK"
	case TrdMarket_US:
		return "TrdMarket_US"
	case TrdMarket_CN:
		return "TrdMarket_CN"
	case TrdMarket_HKCC:
		return "TrdMarket_HKCC"
	case TrdMarket_Futures:
		return "TrdMarket_Futures"
	case TrdMarket_SG:
		return "TrdMarket_SG"
	case TrdMarket_AU:
		return "TrdMarket_AU"
	case TrdMarket_JP:
		return "TrdMarket_JP"
	case TrdMarket_MY:
		return "TrdMarket_MY"
	case TrdMarket_CA:
		return "TrdMarket_CA"
	case TrdMarket_FuturesSimulateHK:
		return "TrdMarket_FuturesSimulateHK"
	case TrdMarket_FuturesSimulateUS:
		return "TrdMarket_FuturesSimulateUS"
	case TrdMarket_FuturesSimulateSG:
		return "TrdMarket_FuturesSimulateSG"
	case TrdMarket_FuturesSimulateJP:
		return "TrdMarket_FuturesSimulateJP"
	case TrdMarket_HKFund:
		return "TrdMarket_HKFund"
	case TrdMarket_USFund:
		return "TrdMarket_USFund"
	case TrdMarket_SGFund:
		return "TrdMarket_SGFund"
	case TrdMarket_MYFund:
		return "TrdMarket_MYFund"
	case TrdMarket_JPFund:
		return "TrdMarket_JPFund"
	default:
		return "TrdMarket_Unknown"
	}
}

// Prefix returns the short trading-market prefix string (e.g. "HK", "US", "CN").
// Returns "UNKNOWN" for unrecognized values.
func (t TrdMarket) Prefix() string {
	switch t {
	case TrdMarket_HK, TrdMarket_HKCC, TrdMarket_HKFund:
		return "HK"
	case TrdMarket_US, TrdMarket_USFund:
		return "US"
	case TrdMarket_CN:
		return "CN"
	case TrdMarket_SG, TrdMarket_SGFund:
		return "SG"
	case TrdMarket_AU:
		return "AU"
	case TrdMarket_JP, TrdMarket_JPFund:
		return "JP"
	case TrdMarket_MY, TrdMarket_MYFund:
		return "MY"
	case TrdMarket_CA:
		return "CA"
	case TrdMarket_Futures, TrdMarket_FuturesSimulateHK, TrdMarket_FuturesSimulateUS, TrdMarket_FuturesSimulateSG, TrdMarket_FuturesSimulateJP:
		return "FUT"
	default:
		return "UNKNOWN"
	}
}

// Int32 converts the enum value to int32 for use in protobuf request fields.
func (m Market) Int32() int32  { return int32(m) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (t TrdMarket) Int32() int32 { return int32(t) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (t TrdEnv) Int32() int32    { return int32(t) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (s TrdSide) Int32() int32   { return int32(s) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (o OrderType) Int32() int32 { return int32(o) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (k KLType) Int32() int32    { return int32(k) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (s SubType) Int32() int32   { return int32(s) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (r RehabType) Int32() int32 { return int32(r) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (m ModifyOrderOp) Int32() int32 { return int32(m) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (s TrdSecMarket) Int32() int32  { return int32(s) }
// Int32 converts the enum value to int32 for use in protobuf request fields.
func (t TimeInForce) Int32() int32   { return int32(t) }

func (m Market) IsValid() bool {
	switch m {
	case Market_HK, Market_US, Market_SH, Market_SZ, Market_SG, Market_JP, Market_AU, Market_MY, Market_CA, Market_FX:
		return true
	default:
		return false
	}
}

func (t TrdMarket) IsValid() bool {
	switch t {
	case TrdMarket_HK, TrdMarket_US, TrdMarket_CN, TrdMarket_HKCC, TrdMarket_Futures, TrdMarket_SG, TrdMarket_AU, TrdMarket_JP:
		return true
	default:
		return false
	}
}

func (t TrdEnv) IsValid() bool {
	return t == TrdEnv_Simulate || t == TrdEnv_Real
}

func (s TrdSide) IsValid() bool {
	switch s {
	case TrdSide_Buy, TrdSide_Sell, TrdSide_SellShort, TrdSide_BuyBack:
		return true
	default:
		return false
	}
}

func (o OrderType) IsValid() bool {
	switch o {
	case OrderType_Normal, OrderType_Market, OrderType_AbsoluteLimit, OrderType_Auction, OrderType_AuctionLimit, OrderType_SpecialLimit, OrderType_Stop, OrderType_StopLimit, OrderType_TrailingStop, OrderType_TrailingStopLimit:
		return true
	default:
		return false
	}
}

func (k KLType) IsValid() bool {
	switch k {
	case KLType_K_1Min, KLType_K_3Min, KLType_K_5Min, KLType_K_15Min, KLType_K_30Min, KLType_K_60Min, KLType_K_Day, KLType_K_Week, KLType_K_Month, KLType_K_Quarter, KLType_K_Year:
		return true
	default:
		return false
	}
}

func (s SubType) IsValid() bool {
	switch s {
	case SubType_Quote, SubType_OrderBook, SubType_Ticker, SubType_Broker, SubType_RT, SubType_K_1Min, SubType_K_3Min, SubType_K_5Min, SubType_K_15Min, SubType_K_30Min, SubType_K_60Min, SubType_K_Day, SubType_K_Week, SubType_K_Month, SubType_K_Quarter, SubType_K_Year:
		return true
	default:
		return false
	}
}

func (r RehabType) IsValid() bool {
	switch r {
	case RehabType_None, RehabType_Forward, RehabType_Backward:
		return true
	default:
		return false
	}
}

func (m ModifyOrderOp) IsValid() bool {
	switch m {
	case ModifyOrderOp_Normal, ModifyOrderOp_Cancel, ModifyOrderOp_Disable, ModifyOrderOp_Enable, ModifyOrderOp_Delete:
		return true
	default:
		return false
	}
}

func (s TrdSecMarket) IsValid() bool {
	switch s {
	case TrdSecMarket_HK, TrdSecMarket_US, TrdSecMarket_CN_SH, TrdSecMarket_CN_SZ, TrdSecMarket_SG, TrdSecMarket_JP, TrdSecMarket_AU, TrdSecMarket_MY, TrdSecMarket_CA, TrdSecMarket_FX:
		return true
	default:
		return false
	}
}

func (t TimeInForce) IsValid() bool {
	switch t {
	case TimeInForce_None, TimeInForce_Day, TimeInForce_GTC, TimeInForce_IOC, TimeInForce_FOK:
		return true
	default:
		return false
	}
}

func (s TrdSecMarket) String() string {
	switch s {
	case TrdSecMarket_Unknown:
		return "TrdSecMarket_Unknown"
	case TrdSecMarket_HK:
		return "TrdSecMarket_HK"
	case TrdSecMarket_US:
		return "TrdSecMarket_US"
	case TrdSecMarket_CN_SH:
		return "TrdSecMarket_CN_SH"
	case TrdSecMarket_CN_SZ:
		return "TrdSecMarket_CN_SZ"
	case TrdSecMarket_SG:
		return "TrdSecMarket_SG"
	case TrdSecMarket_JP:
		return "TrdSecMarket_JP"
	case TrdSecMarket_AU:
		return "TrdSecMarket_AU"
	case TrdSecMarket_MY:
		return "TrdSecMarket_MY"
	case TrdSecMarket_CA:
		return "TrdSecMarket_CA"
	case TrdSecMarket_FX:
		return "TrdSecMarket_FX"
	default:
		return "TrdSecMarket_Unknown"
	}
}

func (t TrdSide) String() string {
	switch t {
	case TrdSide_None:
		return "TrdSide_None"
	case TrdSide_Buy:
		return "TrdSide_Buy"
	case TrdSide_Sell:
		return "TrdSide_Sell"
	case TrdSide_SellShort:
		return "TrdSide_SellShort"
	case TrdSide_BuyBack:
		return "TrdSide_BuyBack"
	default:
		return "TrdSide_Unknown"
	}
}

func (o OrderType) String() string {
	switch o {
	case OrderType_None:
		return "OrderType_None"
	case OrderType_Normal:
		return "OrderType_Normal"
	case OrderType_Market:
		return "OrderType_Market"
	case OrderType_AbsoluteLimit:
		return "OrderType_AbsoluteLimit"
	case OrderType_Auction:
		return "OrderType_Auction"
	case OrderType_AuctionLimit:
		return "OrderType_AuctionLimit"
	case OrderType_SpecialLimit:
		return "OrderType_SpecialLimit"
	case OrderType_SpecialLimitAll:
		return "OrderType_SpecialLimitAll"
	case OrderType_Stop:
		return "OrderType_Stop"
	case OrderType_StopLimit:
		return "OrderType_StopLimit"
	case OrderType_MarketIfTouched:
		return "OrderType_MarketIfTouched"
	case OrderType_LimitIfTouched:
		return "OrderType_LimitIfTouched"
	case OrderType_TrailingStop:
		return "OrderType_TrailingStop"
	case OrderType_TrailingStopLimit:
		return "OrderType_TrailingStopLimit"
	case OrderType_TWAP:
		return "OrderType_TWAP"
	case OrderType_TWAPLimit:
		return "OrderType_TWAPLimit"
	case OrderType_VWAP:
		return "OrderType_VWAP"
	case OrderType_VWAPLimit:
		return "OrderType_VWAPLimit"
	default:
		return "OrderType_Unknown"
	}
}

func (o OrderStatus) String() string {
	switch o {
	case OrderStatus_None:
		return "OrderStatus_None"
	case OrderStatus_Unsubmitted:
		return "OrderStatus_Unsubmitted"
	case OrderStatus_WaitingSubmit:
		return "OrderStatus_WaitingSubmit"
	case OrderStatus_Submitting:
		return "OrderStatus_Submitting"
	case OrderStatus_SubmitFailed:
		return "OrderStatus_SubmitFailed"
	case OrderStatus_TimeOut:
		return "OrderStatus_TimeOut"
	case OrderStatus_Submitted:
		return "OrderStatus_Submitted"
	case OrderStatus_FilledPart:
		return "OrderStatus_FilledPart"
	case OrderStatus_FilledAll:
		return "OrderStatus_FilledAll"
	case OrderStatus_CancellingPart:
		return "OrderStatus_CancellingPart"
	case OrderStatus_CancellingAll:
		return "OrderStatus_CancellingAll"
	case OrderStatus_CancelledPart:
		return "OrderStatus_CancelledPart"
	case OrderStatus_CancelledAll:
		return "OrderStatus_CancelledAll"
	case OrderStatus_Failed:
		return "OrderStatus_Failed"
	case OrderStatus_Disabled:
		return "OrderStatus_Disabled"
	case OrderStatus_Deleted:
		return "OrderStatus_Deleted"
	case OrderStatus_FillCancelled:
		return "OrderStatus_FillCancelled"
	default:
		return "OrderStatus_Unknown"
	}
}

func (m ModifyOrderOp) String() string {
	switch m {
	case ModifyOrderOp_None:
		return "ModifyOrderOp_None"
	case ModifyOrderOp_Normal:
		return "ModifyOrderOp_Normal"
	case ModifyOrderOp_Cancel:
		return "ModifyOrderOp_Cancel"
	case ModifyOrderOp_Disable:
		return "ModifyOrderOp_Disable"
	case ModifyOrderOp_Enable:
		return "ModifyOrderOp_Enable"
	case ModifyOrderOp_Delete:
		return "ModifyOrderOp_Delete"
	default:
		return "ModifyOrderOp_Unknown"
	}
}

func (c Currency) String() string {
	switch c {
	case Currency_None:
		return "Currency_None"
	case Currency_HKD:
		return "Currency_HKD"
	case Currency_USD:
		return "Currency_USD"
	case Currency_CNY:
		return "Currency_CNY"
	case Currency_HKD_C:
		return "Currency_HKD_C"
	case Currency_USD_C:
		return "Currency_USD_C"
	case Currency_SGD:
		return "Currency_SGD"
	case Currency_AUD:
		return "Currency_AUD"
	case Currency_JPY:
		return "Currency_JPY"
	case Currency_MYR:
		return "Currency_MYR"
	case Currency_CAD:
		return "Currency_CAD"
	case Currency_EUR:
		return "Currency_EUR"
	case Currency_GBP:
		return "Currency_GBP"
	case Currency_CHF:
		return "Currency_CHF"
	case Currency_THB:
		return "Currency_THB"
	default:
		return "Currency_Unknown"
	}
}

func (m MarketState) String() string {
	switch m {
	case MarketState_None:
		return "MarketState_None"
	case MarketState_Auction:
		return "MarketState_Auction"
	case MarketState_WaitingOpen:
		return "MarketState_WaitingOpen"
	case MarketState_Morning:
		return "MarketState_Morning"
	case MarketState_Rest:
		return "MarketState_Rest"
	case MarketState_Afternoon:
		return "MarketState_Afternoon"
	case MarketState_Closed:
		return "MarketState_Closed"
	case MarketState_PreMarketBegin:
		return "MarketState_PreMarketBegin"
	case MarketState_PreMarketEnd:
		return "MarketState_PreMarketEnd"
	case MarketState_AfterHoursBegin:
		return "MarketState_AfterHoursBegin"
	case MarketState_AfterHoursEnd:
		return "MarketState_AfterHoursEnd"
	case MarketState_NightOpen:
		return "MarketState_NightOpen"
	case MarketState_NightEnd:
		return "MarketState_NightEnd"
	default:
		return "MarketState_Unknown"
	}
}

func (o OptionType) String() string {
	switch o {
	case OptionType_None:
		return "OptionType_None"
	case OptionType_Call:
		return "OptionType_Call"
	case OptionType_Put:
		return "OptionType_Put"
	default:
		return "OptionType_Unknown"
	}
}

func (r RetType) String() string {
	switch r {
	case RetType_Succeed:
		return "RetType_Succeed"
	case RetType_Failed:
		return "RetType_Failed"
	case RetType_TimeOut:
		return "RetType_TimeOut"
	case RetType_DisConnect:
		return "RetType_DisConnect"
	case RetType_Unknown:
		return "RetType_Unknown"
	case RetType_Invalid:
		return "RetType_Invalid"
	default:
		return "RetType_Unknown"
	}
}

func (r RiskLevel) String() string {
	switch r {
	case RiskLevel_None:
		return "RiskLevel_None"
	case RiskLevel_Low:
		return "RiskLevel_Low"
	case RiskLevel_Medium:
		return "RiskLevel_Medium"
	case RiskLevel_High:
		return "RiskLevel_High"
	default:
		return "RiskLevel_Unknown"
	}
}

func (r CltRiskStatus) String() string {
	switch r {
	case CltRiskStatus_Unknown:
		return "CltRiskStatus_Unknown"
	case CltRiskStatus_Level1:
		return "CltRiskStatus_Level1"
	case CltRiskStatus_Level2:
		return "CltRiskStatus_Level2"
	case CltRiskStatus_Level3:
		return "CltRiskStatus_Level3"
	case CltRiskStatus_Level4:
		return "CltRiskStatus_Level4"
	case CltRiskStatus_Level5:
		return "CltRiskStatus_Level5"
	case CltRiskStatus_Level6:
		return "CltRiskStatus_Level6"
	case CltRiskStatus_Level7:
		return "CltRiskStatus_Level7"
	case CltRiskStatus_Level8:
		return "CltRiskStatus_Level8"
	case CltRiskStatus_Level9:
		return "CltRiskStatus_Level9"
	default:
		return "CltRiskStatus_Unknown"
	}
}

func (t TimeInForce) String() string {
	switch t {
	case TimeInForce_None:
		return "TimeInForce_None"
	case TimeInForce_Day:
		return "TimeInForce_Day"
	case TimeInForce_GTC:
		return "TimeInForce_GTC"
	case TimeInForce_IOC:
		return "TimeInForce_IOC"
	case TimeInForce_FOK:
		return "TimeInForce_FOK"
	default:
		return "TimeInForce_Unknown"
	}
}

func (t TrailType) String() string {
	switch t {
	case TrailType_None:
		return "TrailType_None"
	case TrailType_Ratio:
		return "TrailType_Ratio"
	case TrailType_Amount:
		return "TrailType_Amount"
	default:
		return "TrailType_Unknown"
	}
}

func (d DealStatus) String() string {
	switch d {
	case DealStatus_OK:
		return "DealStatus_OK"
	case DealStatus_Cancelled:
		return "DealStatus_Cancelled"
	case DealStatus_Changed:
		return "DealStatus_Changed"
	default:
		return "DealStatus_Unknown"
	}
}

func (a AccStatus) String() string {
	switch a {
	case AccStatus_None:
		return "AccStatus_None"
	case AccStatus_Normal:
		return "AccStatus_Normal"
	case AccStatus_Disabled:
		return "AccStatus_Disabled"
	case AccStatus_Deleted:
		return "AccStatus_Deleted"
	case AccStatus_Locked:
		return "AccStatus_Locked"
	default:
		return "AccStatus_Unknown"
	}
}
