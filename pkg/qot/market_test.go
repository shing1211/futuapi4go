package qot

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetmarketstate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetownerplate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsubinfo"
)

func TestGetOwnerPlateRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetOwnerPlateRequest{
		SecurityList: []*qotcommon.Security{security},
	}

	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
	if req.SecurityList[0].GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", req.SecurityList[0].GetCode())
	}
}

func TestGetOwnerPlateResponseConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	plateCode := "BK001"
	plateName := "Technology"
	secCode := "00700"

	s2c := &qotgetownerplate.S2C{
		OwnerPlateList: []*qotgetownerplate.SecurityOwnerPlate{
			{
				Security: &qotcommon.Security{Market: &hkMarket, Code: &secCode},
				Name:     &secCode,
				PlateInfoList: []*qotcommon.PlateInfo{
					{
						Plate: &qotcommon.Security{Market: &hkMarket, Code: &plateCode},
						Name:  &plateName,
					},
				},
			},
		},
	}

	rsp := &GetOwnerPlateResponse{
		OwnerPlateList: s2c.GetOwnerPlateList(),
	}

	if len(rsp.OwnerPlateList) != 1 {
		t.Errorf("expected 1 owner plate entry, got %d", len(rsp.OwnerPlateList))
	}
	if len(rsp.OwnerPlateList[0].GetPlateInfoList()) != 1 {
		t.Errorf("expected 1 plate info, got %d", len(rsp.OwnerPlateList[0].GetPlateInfoList()))
	}
	if rsp.OwnerPlateList[0].GetPlateInfoList()[0].GetName() != "Technology" {
		t.Errorf("expected plate name Technology, got %s", rsp.OwnerPlateList[0].GetPlateInfoList()[0].GetName())
	}
}

func TestGetReferenceRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "HSI2405"; return &s }()}
	req := &GetReferenceRequest{
		Security:      security,
		ReferenceType: 1,
	}

	if req.Security.GetCode() != "HSI2405" {
		t.Errorf("expected code HSI2405, got %s", req.Security.GetCode())
	}
	if req.ReferenceType != 1 {
		t.Errorf("expected ReferenceType 1, got %d", req.ReferenceType)
	}
}

func TestGetReferenceResponseConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	secCode := "00700"
	secName := "Tencent"
	id := int64(1)
	lotSize := int32(100)
	secType := int32(1)
	listTime := "2026-01-01"

	rsp := &GetReferenceResponse{
		StaticInfoList: []*qotcommon.SecurityStaticInfo{
			{
				Basic: &qotcommon.SecurityStaticBasic{
					Security: &qotcommon.Security{Market: &hkMarket, Code: &secCode},
					Id:       &id,
					LotSize:  &lotSize,
					SecType:  &secType,
					Name:     &secName,
					ListTime: &listTime,
				},
			},
		},
	}

	if len(rsp.StaticInfoList) != 1 {
		t.Errorf("expected 1 static info, got %d", len(rsp.StaticInfoList))
	}
	if rsp.StaticInfoList[0].GetBasic().GetSecurity().GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", rsp.StaticInfoList[0].GetBasic().GetSecurity().GetCode())
	}
}

func TestMarketStateInfoFields(t *testing.T) {
	info := &MarketStateInfo{
		MarketState: 1,
	}

	if info.MarketState != 1 {
		t.Errorf("expected MarketState 1, got %d", info.MarketState)
	}
}

func TestGetMarketStateRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetMarketStateRequest{
		SecurityList: []*qotcommon.Security{security},
	}

	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
}

func TestGetMarketStateResponseConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	secCode := "00700"
	name := "Tencent"
	marketState := int32(1)

	s2c := &qotgetmarketstate.S2C{
		MarketInfoList: []*qotgetmarketstate.MarketInfo{
			{
				Security:    &qotcommon.Security{Market: &hkMarket, Code: &secCode},
				Name:        &name,
				MarketState: &marketState,
			},
		},
	}

	rsp := &GetMarketStateResponse{
		MarketInfoList: make([]*MarketStateInfo, 0, len(s2c.GetMarketInfoList())),
	}

	for _, mi := range s2c.GetMarketInfoList() {
		rsp.MarketInfoList = append(rsp.MarketInfoList, &MarketStateInfo{
			Security:    mi.GetSecurity(),
			Name:        mi.GetName(),
			MarketState: mi.GetMarketState(),
		})
	}

	if len(rsp.MarketInfoList) != 1 {
		t.Errorf("expected 1 market info, got %d", len(rsp.MarketInfoList))
	}
	if rsp.MarketInfoList[0].MarketState != 1 {
		t.Errorf("expected MarketState 1, got %d", rsp.MarketInfoList[0].MarketState)
	}
}

func TestGetSubInfoResponseConstruction(t *testing.T) {
	totalUsedQuota := int32(10)
	remainQuota := int32(90)
	subType := int32(1)
	usedQuota := int32(5)
	isOwnConnData := true

	s2c := &qotgetsubinfo.S2C{
		ConnSubInfoList: []*qotcommon.ConnSubInfo{
			{
				SubInfoList: []*qotcommon.SubInfo{
					{SubType: &subType},
				},
				UsedQuota:     &usedQuota,
				IsOwnConnData: &isOwnConnData,
			},
		},
		TotalUsedQuota: &totalUsedQuota,
		RemainQuota:    &remainQuota,
	}

	rsp := &GetSubInfoResponse{
		ConnSubInfoList: s2c.GetConnSubInfoList(),
		TotalUsedQuota:  s2c.GetTotalUsedQuota(),
		RemainQuota:     s2c.GetRemainQuota(),
	}

	if len(rsp.ConnSubInfoList) != 1 {
		t.Errorf("expected 1 conn sub info, got %d", len(rsp.ConnSubInfoList))
	}
	if rsp.TotalUsedQuota != 10 {
		t.Errorf("expected TotalUsedQuota 10, got %d", rsp.TotalUsedQuota)
	}
	if rsp.RemainQuota != 90 {
		t.Errorf("expected RemainQuota 90, got %d", rsp.RemainQuota)
	}
	if len(rsp.ConnSubInfoList[0].GetSubInfoList()) != 1 {
		t.Errorf("expected 1 sub info, got %d", len(rsp.ConnSubInfoList[0].GetSubInfoList()))
	}
}

func TestMarketProtoIDConstants(t *testing.T) {
	tests := []struct {
		name  string
		value int
	}{
		{"ProtoID_GetOwnerPlate", 3207},
		{"ProtoID_GetReference", 3206},
		{"ProtoID_GetMarketState", 3223},
		{"ProtoID_GetSubInfo", 3003},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got int
			switch tc.name {
			case "ProtoID_GetOwnerPlate":
				got = ProtoID_GetOwnerPlate
			case "ProtoID_GetReference":
				got = ProtoID_GetReference
			case "ProtoID_GetMarketState":
				got = ProtoID_GetMarketState
			case "ProtoID_GetSubInfo":
				got = ProtoID_GetSubInfo
			}
			if got != tc.value {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.value, got)
			}
		})
	}
}
