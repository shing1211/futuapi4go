package constant

import "testing"

func TestProtoIDName(t *testing.T) {
	cases := map[int32]string{
		1001: "InitConnect",
		1002: "GetGlobalState",
		1004: "KeepAlive",
		2101: "Trd_GetFunds",
		2102: "Trd_GetPositionList",
		2201: "Trd_GetOrderList",
		2202: "Trd_PlaceOrder",
		3401: "Qot_GetEarningsCalendar",
	}
	for id, want := range cases {
		if got := ProtoIDName(id); got != want {
			t.Errorf("ProtoIDName(%d) = %q, want %q", id, got, want)
		}
	}
}

func TestProtoIDNameUnknownFallsBackToNumber(t *testing.T) {
	if got := ProtoIDName(999999); got != "protoID_999999" {
		t.Errorf("unknown id = %q, want protoID_999999", got)
	}
	if got := ProtoIDName(0); got != "protoID_0" {
		t.Errorf("zero id = %q, want protoID_0", got)
	}
}
