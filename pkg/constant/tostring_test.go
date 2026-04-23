package constant

import "testing"

func TestMarketString(t *testing.T) {
	tests := []struct {
		m   Market
		want string
	}{
		{Market_None, "Market_None"},
		{Market_HK, "Market_HK"},
		{Market_US, "Market_US"},
		{Market_SH, "Market_SH"},
		{Market_SZ, "Market_SZ"},
		{Market_SG, "Market_SG"},
		{Market_JP, "Market_JP"},
		{Market_AU, "Market_AU"},
		{Market_MY, "Market_MY"},
		{Market_CA, "Market_CA"},
		{Market_FX, "Market_FX"},
		{Market(999), "Market_Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("Market(%d).String() = %q, want %q", tt.m, got, tt.want)
			}
		})
	}
}

func TestMarketPrefix(t *testing.T) {
	tests := []struct {
		m   Market
		want string
	}{
		{Market_HK, "HK"},
		{Market_US, "US"},
		{Market_SH, "SH"},
		{Market_SZ, "SZ"},
		{Market_SG, "SG"},
		{Market_JP, "JP"},
		{Market_AU, "AU"},
		{Market_MY, "MY"},
		{Market_CA, "CA"},
		{Market_FX, "FX"},
		{Market_None, "UNKNOWN"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.m.Prefix(); got != tt.want {
				t.Errorf("Market(%d).Prefix() = %q, want %q", tt.m, got, tt.want)
			}
		})
	}
}

func TestTrdEnvString(t *testing.T) {
	tests := []struct {
		t   TrdEnv
		want string
	}{
		{TrdEnv_Simulate, "TrdEnv_Simulate"},
		{TrdEnv_Real, "TrdEnv_Real"},
		{TrdEnv(99), "TrdEnv_Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("TrdEnv(%d).String() = %q, want %q", tt.t, got, tt.want)
			}
		})
	}
}

func TestTrdSideString(t *testing.T) {
	tests := []struct {
		s   TrdSide
		want string
	}{
		{TrdSide_Buy, "TrdSide_Buy"},
		{TrdSide_Sell, "TrdSide_Sell"},
		{TrdSide_SellShort, "TrdSide_SellShort"},
		{TrdSide_BuyBack, "TrdSide_BuyBack"},
		{TrdSide_None, "TrdSide_None"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("TrdSide(%d).String() = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestOrderStatusString(t *testing.T) {
	tests := []struct {
		s   OrderStatus
		want string
	}{
		{OrderStatus_Submitted, "OrderStatus_Submitted"},
		{OrderStatus_FilledAll, "OrderStatus_FilledAll"},
		{OrderStatus_CancelledAll, "OrderStatus_CancelledAll"},
		{OrderStatus_Failed, "OrderStatus_Failed"},
		{OrderStatus_FillCancelled, "OrderStatus_FillCancelled"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("OrderStatus(%d).String() = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestOrderTypeString(t *testing.T) {
	tests := []struct {
		o   OrderType
		want string
	}{
		{OrderType_Normal, "OrderType_Normal"},
		{OrderType_Market, "OrderType_Market"},
		{OrderType_TrailingStop, "OrderType_TrailingStop"},
		{OrderType_TWAP, "OrderType_TWAP"},
		{OrderType_VWAPLimit, "OrderType_VWAPLimit"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("OrderType(%d).String() = %q, want %q", tt.o, got, tt.want)
			}
		})
	}
}

func TestCurrencyString(t *testing.T) {
	tests := []struct {
		c   Currency
		want string
	}{
		{Currency_HKD, "Currency_HKD"},
		{Currency_USD, "Currency_USD"},
		{Currency_CNY, "Currency_CNY"},
		{Currency_SGD, "Currency_SGD"},
		{Currency_EUR, "Currency_EUR"},
		{Currency_GBP, "Currency_GBP"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Currency(%d).String() = %q, want %q", tt.c, got, tt.want)
			}
		})
	}
}

func TestKLTypeString(t *testing.T) {
	tests := []struct {
		k   KLType
		want string
	}{
		{KLType_K_Day, "KLType_K_Day"},
		{KLType_K_1Min, "KLType_K_1Min"},
		{KLType_K_Week, "KLType_K_Week"},
		{KLType_K_Month, "KLType_K_Month"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("KLType(%d).String() = %q, want %q", tt.k, got, tt.want)
			}
		})
	}
}

func TestOptionTypeString(t *testing.T) {
	tests := []struct {
		o   OptionType
		want string
	}{
		{OptionType_Call, "OptionType_Call"},
		{OptionType_Put, "OptionType_Put"},
		{OptionType_None, "OptionType_None"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("OptionType(%d).String() = %q, want %q", tt.o, got, tt.want)
			}
		})
	}
}

func TestSubTypeString(t *testing.T) {
	tests := []struct {
		s   SubType
		want string
	}{
		{SubType_Quote, "SubType_Quote"},
		{SubType_K_Day, "SubType_K_Day"},
		{SubType_K_1Min, "SubType_K_1Min"},
		{SubType_Ticker, "SubType_Ticker"},
		{SubType_RT, "SubType_RT"},
		{SubType_OrderBook, "SubType_OrderBook"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("SubType(%d).String() = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestRetTypeString(t *testing.T) {
	tests := []struct {
		r   RetType
		want string
	}{
		{RetType_Succeed, "RetType_Succeed"},
		{RetType_Failed, "RetType_Failed"},
		{RetType_TimeOut, "RetType_TimeOut"},
		{RetType_DisConnect, "RetType_DisConnect"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("RetType(%d).String() = %q, want %q", tt.r, got, tt.want)
			}
		})
	}
}

func TestSecurityTypeString(t *testing.T) {
	tests := []struct {
		s   SecurityType
		want string
	}{
		{SecurityType_Stock, "SecurityType_Stock"},
		{SecurityType_Warrant, "SecurityType_Warrant"},
		{SecurityType_ETF, "SecurityType_ETF"},
		{SecurityType_Future, "SecurityType_Future"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("SecurityType(%d).String() = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestModifyOrderOpString(t *testing.T) {
	tests := []struct {
		m   ModifyOrderOp
		want string
	}{
		{ModifyOrderOp_Normal, "ModifyOrderOp_Normal"},
		{ModifyOrderOp_Cancel, "ModifyOrderOp_Cancel"},
		{ModifyOrderOp_Disable, "ModifyOrderOp_Disable"},
		{ModifyOrderOp_Enable, "ModifyOrderOp_Enable"},
		{ModifyOrderOp_Delete, "ModifyOrderOp_Delete"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("ModifyOrderOp(%d).String() = %q, want %q", tt.m, got, tt.want)
			}
		})
	}
}

func TestTimeInForceString(t *testing.T) {
	tests := []struct {
		t   TimeInForce
		want string
	}{
		{TimeInForce_Day, "TimeInForce_Day"},
		{TimeInForce_GTC, "TimeInForce_GTC"},
		{TimeInForce_IOC, "TimeInForce_IOC"},
		{TimeInForce_FOK, "TimeInForce_FOK"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("TimeInForce(%d).String() = %q, want %q", tt.t, got, tt.want)
			}
		})
	}
}

func TestTrdMarketString(t *testing.T) {
	tests := []struct {
		m   TrdMarket
		want string
	}{
		{TrdMarket_HK, "TrdMarket_HK"},
		{TrdMarket_US, "TrdMarket_US"},
		{TrdMarket_CN, "TrdMarket_CN"},
		{TrdMarket_HKCC, "TrdMarket_HKCC"},
		{TrdMarket_Futures, "TrdMarket_Futures"},
		{TrdMarket_SG, "TrdMarket_SG"},
		{TrdMarket_AU, "TrdMarket_AU"},
		{TrdMarket_JP, "TrdMarket_JP"},
		{TrdMarket_MY, "TrdMarket_MY"},
		{TrdMarket_CA, "TrdMarket_CA"},
		{TrdMarket_None, "TrdMarket_None"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("TrdMarket(%d).String() = %q, want %q", tt.m, got, tt.want)
			}
		})
	}
}

func TestTrdSecMarketString(t *testing.T) {
	tests := []struct {
		m   TrdSecMarket
		want string
	}{
		{TrdSecMarket_HK, "TrdSecMarket_HK"},
		{TrdSecMarket_US, "TrdSecMarket_US"},
		{TrdSecMarket_CN_SH, "TrdSecMarket_CN_SH"},
		{TrdSecMarket_CN_SZ, "TrdSecMarket_CN_SZ"},
		{TrdSecMarket_FX, "TrdSecMarket_FX"},
		{TrdSecMarket_Unknown, "TrdSecMarket_Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("TrdSecMarket(%d).String() = %q, want %q", tt.m, got, tt.want)
			}
		})
	}
}

func TestRehabTypeString(t *testing.T) {
	tests := []struct {
		r   RehabType
		want string
	}{
		{RehabType_Forward, "RehabType_Forward"},
		{RehabType_Backward, "RehabType_Backward"},
		{RehabType_None, "RehabType_None"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("RehabType(%d).String() = %q, want %q", tt.r, got, tt.want)
			}
		})
	}
}
