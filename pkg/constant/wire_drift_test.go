package constant

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

// This package hand-rolls enums that mirror the generated protobuf enums, and
// nothing links the two. A typo or a missed regeneration therefore makes them
// disagree with the wire silently. That is exactly how SubType_OrderBookOdd came
// to hold 18 - the value the wire assigns to SubType_KL_10Min - so asking to
// subscribe to the odd-lot order book requested 10-minute bars instead.
//
// The pairs below are compared by value, never by name: the generated enum spells
// quarter "Qurater" (an upstream Futu typo this package does not reproduce), so
// only the numbers are comparable.

func TestSubTypeMatchesWire(t *testing.T) {
	pairs := []struct {
		name string
		got  SubType
		want qotcommon.SubType
	}{
		{"SubType_None", SubType_None, qotcommon.SubType_SubType_None},
		{"SubType_Quote", SubType_Quote, qotcommon.SubType_SubType_Basic},
		{"SubType_OrderBook", SubType_OrderBook, qotcommon.SubType_SubType_OrderBook},
		{"SubType_Ticker", SubType_Ticker, qotcommon.SubType_SubType_Ticker},
		{"SubType_RT", SubType_RT, qotcommon.SubType_SubType_RT},
		{"SubType_Broker", SubType_Broker, qotcommon.SubType_SubType_Broker},
		{"SubType_K_Day", SubType_K_Day, qotcommon.SubType_SubType_KL_Day},
		{"SubType_K_5Min", SubType_K_5Min, qotcommon.SubType_SubType_KL_5Min},
		{"SubType_K_15Min", SubType_K_15Min, qotcommon.SubType_SubType_KL_15Min},
		{"SubType_K_30Min", SubType_K_30Min, qotcommon.SubType_SubType_KL_30Min},
		{"SubType_K_60Min", SubType_K_60Min, qotcommon.SubType_SubType_KL_60Min},
		{"SubType_K_1Min", SubType_K_1Min, qotcommon.SubType_SubType_KL_1Min},
		{"SubType_K_Week", SubType_K_Week, qotcommon.SubType_SubType_KL_Week},
		{"SubType_K_Month", SubType_K_Month, qotcommon.SubType_SubType_KL_Month},
		{"SubType_K_Quarter", SubType_K_Quarter, qotcommon.SubType_SubType_KL_Qurater},
		{"SubType_K_Year", SubType_K_Year, qotcommon.SubType_SubType_KL_Year},
		{"SubType_K_3Min", SubType_K_3Min, qotcommon.SubType_SubType_KL_3Min},
		{"SubType_K_10Min", SubType_K_10Min, qotcommon.SubType_SubType_KL_10Min},
		{"SubType_K_120Min", SubType_K_120Min, qotcommon.SubType_SubType_KL_120Min},
		{"SubType_K_180Min", SubType_K_180Min, qotcommon.SubType_SubType_KL_180Min},
		{"SubType_K_240Min", SubType_K_240Min, qotcommon.SubType_SubType_KL_240Min},
		{"SubType_OrderBookOdd", SubType_OrderBookOdd, qotcommon.SubType_SubType_OrderBook_Odd},
	}

	for _, p := range pairs {
		if int32(p.got) != int32(p.want) {
			t.Errorf("%s = %d, but the wire value is %d", p.name, int32(p.got), int32(p.want))
		}
	}
}

func TestKLTypeMatchesWire(t *testing.T) {
	pairs := []struct {
		name string
		got  KLType
		want qotcommon.KLType
	}{
		{"KLType_None", KLType_None, qotcommon.KLType_KLType_Unknown},
		{"KLType_K_1Min", KLType_K_1Min, qotcommon.KLType_KLType_1Min},
		{"KLType_K_Day", KLType_K_Day, qotcommon.KLType_KLType_Day},
		{"KLType_K_Week", KLType_K_Week, qotcommon.KLType_KLType_Week},
		{"KLType_K_Month", KLType_K_Month, qotcommon.KLType_KLType_Month},
		{"KLType_K_Year", KLType_K_Year, qotcommon.KLType_KLType_Year},
		{"KLType_K_5Min", KLType_K_5Min, qotcommon.KLType_KLType_5Min},
		{"KLType_K_15Min", KLType_K_15Min, qotcommon.KLType_KLType_15Min},
		{"KLType_K_30Min", KLType_K_30Min, qotcommon.KLType_KLType_30Min},
		{"KLType_K_60Min", KLType_K_60Min, qotcommon.KLType_KLType_60Min},
		{"KLType_K_3Min", KLType_K_3Min, qotcommon.KLType_KLType_3Min},
		{"KLType_K_Quarter", KLType_K_Quarter, qotcommon.KLType_KLType_Quarter},
		{"KLType_K_10Min", KLType_K_10Min, qotcommon.KLType_KLType_10Min},
		{"KLType_K_120Min", KLType_K_120Min, qotcommon.KLType_KLType_120Min},
		{"KLType_K_180Min", KLType_K_180Min, qotcommon.KLType_KLType_180Min},
		{"KLType_K_240Min", KLType_K_240Min, qotcommon.KLType_KLType_240Min},
	}

	for _, p := range pairs {
		if int32(p.got) != int32(p.want) {
			t.Errorf("%s = %d, but the wire value is %d", p.name, int32(p.got), int32(p.want))
		}
	}
}

// TestToSubTypeRoundTrip pins the KLType-to-SubType correspondence in both
// directions. The two enums number differently, so a round trip is the cheapest
// way to catch a pair that maps to the wrong value or is missing entirely.
func TestToSubTypeRoundTrip(t *testing.T) {
	types := []KLType{
		KLType_K_1Min, KLType_K_3Min, KLType_K_5Min, KLType_K_10Min,
		KLType_K_15Min, KLType_K_30Min, KLType_K_60Min, KLType_K_120Min,
		KLType_K_180Min, KLType_K_240Min, KLType_K_Day, KLType_K_Week,
		KLType_K_Month, KLType_K_Quarter, KLType_K_Year,
	}

	for _, kt := range types {
		st, err := kt.ToSubType()
		if err != nil {
			t.Errorf("%v.ToSubType() returned error: %v", kt, err)
			continue
		}
		if !st.IsValid() {
			t.Errorf("%v.ToSubType() = %v, which IsValid() rejects", kt, st)
		}
		if !st.IsKLType() {
			t.Errorf("%v.ToSubType() = %v, which IsKLType() rejects", kt, st)
		}
		back, err := st.ToKLType()
		if err != nil {
			t.Errorf("%v.ToKLType() returned error: %v", st, err)
			continue
		}
		if back != kt {
			t.Errorf("round trip %v -> %v -> %v lost the value", kt, st, back)
		}
	}
}

// TestToSubTypeRejectsNonKLine guards the non-K types: a quote or order-book
// subscription is not a bar interval.
func TestToSubTypeRejectsNonKLine(t *testing.T) {
	for _, st := range []SubType{SubType_Quote, SubType_OrderBook, SubType_Ticker, SubType_RT, SubType_Broker, SubType_OrderBookOdd} {
		if _, err := st.ToKLType(); err == nil {
			t.Errorf("%v.ToKLType() should fail, but succeeded", st)
		}
	}
	if _, err := KLType_None.ToSubType(); err == nil {
		t.Error("KLType_None.ToSubType() should fail, but succeeded")
	}
}

// TestIsValidAgreesWithIsKLType keeps the two validity checks from drifting,
// since IsValid delegates to IsKLType for the K types.
func TestIsValidAgreesWithIsKLType(t *testing.T) {
	for v := int32(0); v <= 22; v++ {
		st := SubType(v)
		if st.IsKLType() && !st.IsValid() {
			t.Errorf("%v.IsKLType() is true but IsValid() is false", st)
		}
	}
	for _, st := range []SubType{SubType_Quote, SubType_OrderBook, SubType_Ticker, SubType_RT, SubType_Broker, SubType_OrderBookOdd} {
		if !st.IsValid() {
			t.Errorf("%v should be valid", st)
		}
	}
}
