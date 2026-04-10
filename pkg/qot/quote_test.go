package qot

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
)

func TestGetKLRequestValidation(t *testing.T) {
	// Test valid request
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	security := &qotcommon.Security{Market: &hkMarket, Code: &code}

	req := &GetKLRequest{
		Security:  security,
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    10,
	}

	if req.Security == nil {
		t.Error("Security should not be nil")
	}
	if req.Security.GetCode() != code {
		t.Errorf("expected code %s, got %s", code, req.Security.GetCode())
	}
	if req.ReqNum != 10 {
		t.Errorf("expected ReqNum 10, got %d", req.ReqNum)
	}
}

func TestGetOrderBookRequestValidation(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	security := &qotcommon.Security{Market: &hkMarket, Code: &code}

	req := &GetOrderBookRequest{
		Security: security,
		Num:      10,
	}

	if req.Num != 10 {
		t.Errorf("expected Num 10, got %d", req.Num)
	}
}

func TestSubscribeRequestValidation(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	security := &qotcommon.Security{Market: &hkMarket, Code: &code}

	req := &SubscribeRequest{
		SecurityList:     []*qotcommon.Security{security},
		SubTypeList:      []SubType{SubType_Basic, SubType_KL},
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
	}

	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
	if len(req.SubTypeList) != 2 {
		t.Errorf("expected 2 sub types, got %d", len(req.SubTypeList))
	}
	if !req.IsSubOrUnSub {
		t.Error("expected IsSubOrUnSub to be true")
	}
}

func TestBasicQotStructFields(t *testing.T) {
	bq := &BasicQot{
		Security:  &qotcommon.Security{Market: func() *int32 { v := int32(1); return &v }(), Code: func() *string { s := "00700"; return &s }()},
		Name:      "Tencent",
		CurPrice:  350.50,
		OpenPrice: 348.00,
		HighPrice: 352.00,
		LowPrice:  347.00,
		Volume:    12345678,
		Turnover:  4321098765.00,
	}

	if bq.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", bq.Security.GetCode())
	}
	if bq.Name != "Tencent" {
		t.Errorf("expected name Tencent, got %s", bq.Name)
	}
	if bq.CurPrice != 350.50 {
		t.Errorf("expected CurPrice 350.50, got %f", bq.CurPrice)
	}
}

func TestKLineStructFields(t *testing.T) {
	kl := &KLine{
		Time:           "2026-04-08 15:00:00",
		IsBlank:        false,
		HighPrice:      352.00,
		OpenPrice:      348.00,
		LowPrice:       347.00,
		ClosePrice:     350.50,
		LastClosePrice: 349.00,
		Volume:         12345678,
		Turnover:       4321098765.00,
		ChangeRate:     0.43,
		Timestamp:      1775635200.0,
	}

	if kl.Time != "2026-04-08 15:00:00" {
		t.Errorf("unexpected Time: %s", kl.Time)
	}
	if kl.ClosePrice != 350.50 {
		t.Errorf("expected ClosePrice 350.50, got %f", kl.ClosePrice)
	}
}

func TestRTFields(t *testing.T) {
	rt := &RT{
		Time:           "2026-04-08 15:30:00",
		Price:          350.50,
		LastClosePrice: 349.00,
		AvgPrice:       349.80,
		Volume:         12345678,
		Turnover:       4321098765.00,
	}

	if rt.Time != "2026-04-08 15:30:00" {
		t.Errorf("expected Time, got %s", rt.Time)
	}
	if rt.Price != 350.50 {
		t.Errorf("expected Price 350.50, got %f", rt.Price)
	}
	if rt.LastClosePrice != 349.00 {
		t.Errorf("expected LastClosePrice 349.00, got %f", rt.LastClosePrice)
	}
	if rt.AvgPrice != 349.80 {
		t.Errorf("expected AvgPrice 349.80, got %f", rt.AvgPrice)
	}
	if rt.Volume != 12345678 {
		t.Errorf("expected Volume 12345678, got %d", rt.Volume)
	}
	if rt.Turnover != 4321098765.00 {
		t.Errorf("expected Turnover 4321098765.00, got %f", rt.Turnover)
	}
}

func TestGetRTRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetRTRequest{
		Security: security,
	}

	if req.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", req.Security.GetCode())
	}
}

func TestGetRTResponseConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	rsp := &GetRTResponse{
		Security: security,
		Name:     "Tencent",
		RTList: []*RT{
			{Time: "10:00:00", Price: 350.0, Volume: 1000, Turnover: 350000.0},
			{Time: "10:01:00", Price: 350.5, Volume: 2000, Turnover: 701000.0},
		},
	}

	if rsp.Name != "Tencent" {
		t.Errorf("expected Name Tencent, got %s", rsp.Name)
	}
	if len(rsp.RTList) != 2 {
		t.Errorf("expected 2 RT entries, got %d", len(rsp.RTList))
	}
	if rsp.RTList[1].Price != 350.5 {
		t.Errorf("expected price 350.5, got %f", rsp.RTList[1].Price)
	}
}

func TestOrderBookFields(t *testing.T) {
	ob := &OrderBook{
		Price:      350.50,
		Volume:     10000,
		OrderCount: 5,
		DetailList: []*OrderBookDetail{
			{Volume: 500, OrderID: 12345},
		},
	}

	if ob.Price != 350.50 {
		t.Errorf("expected Price 350.50, got %f", ob.Price)
	}
	if ob.Volume != 10000 {
		t.Errorf("expected Volume 10000, got %d", ob.Volume)
	}
	if ob.OrderCount != 5 {
		t.Errorf("expected OrderCount 5, got %d", ob.OrderCount)
	}
	if len(ob.DetailList) != 1 {
		t.Errorf("expected 1 detail, got %d", len(ob.DetailList))
	}
	if ob.DetailList[0].OrderID != 12345 {
		t.Errorf("expected OrderID 12345, got %d", ob.DetailList[0].OrderID)
	}
}

func TestOrderBookDetailFields(t *testing.T) {
	obd := &OrderBookDetail{
		Volume:  500,
		OrderID: 12345,
	}

	if obd.Volume != 500 {
		t.Errorf("expected Volume 500, got %d", obd.Volume)
	}
	if obd.OrderID != 12345 {
		t.Errorf("expected OrderID 12345, got %d", obd.OrderID)
	}
}

func TestGetOrderBookResponseConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	rsp := &GetOrderBookResponse{
		Security:                security,
		Name:                    "Tencent",
		OrderBookAskList:        []*OrderBook{{Price: 351.00, Volume: 5000, OrderCount: 3}},
		OrderBookBidList:        []*OrderBook{{Price: 350.00, Volume: 5000, OrderCount: 3}},
		SvrRecvTimeBid:          "10:00:00",
		SvrRecvTimeBidTimestamp: 1775635200.0,
		SvrRecvTimeAsk:          "10:00:00",
		SvrRecvTimeAskTimestamp: 1775635200.0,
	}

	if len(rsp.OrderBookAskList) != 1 {
		t.Errorf("expected 1 ask level, got %d", len(rsp.OrderBookAskList))
	}
	if len(rsp.OrderBookBidList) != 1 {
		t.Errorf("expected 1 bid level, got %d", len(rsp.OrderBookBidList))
	}
	if rsp.OrderBookAskList[0].Price != 351.00 {
		t.Errorf("expected ask price 351.00, got %f", rsp.OrderBookAskList[0].Price)
	}
}

func TestTickerFields(t *testing.T) {
	ticker := &Ticker{
		Time:      "2026-04-08 15:00:00",
		Sequence:  123456,
		Dir:       1,
		Price:     350.50,
		Volume:    1000,
		Turnover:  350500.00,
		RecvTime:  1775635200.0,
		Type:      0,
		TypeSign:  1,
		Timestamp: 1775635200.0,
	}

	if ticker.Price != 350.50 {
		t.Errorf("expected Price 350.50, got %f", ticker.Price)
	}
	if ticker.Volume != 1000 {
		t.Errorf("expected Volume 1000, got %d", ticker.Volume)
	}
	if ticker.Dir != 1 {
		t.Errorf("expected Dir 1, got %d", ticker.Dir)
	}
	if ticker.Turnover != 350500.00 {
		t.Errorf("expected Turnover 350500.00, got %f", ticker.Turnover)
	}
}

func TestGetTickerRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetTickerRequest{
		Security: security,
		Num:      100,
	}

	if req.Num != 100 {
		t.Errorf("expected Num 100, got %d", req.Num)
	}
}

func TestGetTickerResponseConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	rsp := &GetTickerResponse{
		Security: security,
		Name:     "Tencent",
		TickerList: []*Ticker{
			{Time: "10:00:00", Price: 350.0, Volume: 500},
		},
	}

	if rsp.Name != "Tencent" {
		t.Errorf("expected Name Tencent, got %s", rsp.Name)
	}
	if len(rsp.TickerList) != 1 {
		t.Errorf("expected 1 ticker, got %d", len(rsp.TickerList))
	}
}

func TestBrokerFields(t *testing.T) {
	broker := &Broker{
		ID:     12345,
		Name:   "Citi",
		Pos:    1,
		Volume: 5000,
	}

	if broker.ID != 12345 {
		t.Errorf("expected ID 12345, got %d", broker.ID)
	}
	if broker.Name != "Citi" {
		t.Errorf("expected Name Citi, got %s", broker.Name)
	}
	if broker.Pos != 1 {
		t.Errorf("expected Pos 1, got %d", broker.Pos)
	}
}

func TestGetBrokerRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetBrokerRequest{
		Security: security,
		Num:      10,
	}

	if req.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", req.Security.GetCode())
	}
}

func TestGetBrokerResponseConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	rsp := &GetBrokerResponse{
		Security: security,
		Name:     "Tencent",
		AskBrokerList: []*Broker{
			{ID: 1, Name: "Citi", Pos: 1, Volume: 5000},
		},
		BidBrokerList: []*Broker{
			{ID: 2, Name: "HSBC", Pos: 1, Volume: 6000},
		},
	}

	if len(rsp.AskBrokerList) != 1 {
		t.Errorf("expected 1 ask broker, got %d", len(rsp.AskBrokerList))
	}
	if len(rsp.BidBrokerList) != 1 {
		t.Errorf("expected 1 bid broker, got %d", len(rsp.BidBrokerList))
	}
	if rsp.AskBrokerList[0].Name != "Citi" {
		t.Errorf("expected Citi, got %s", rsp.AskBrokerList[0].Name)
	}
}

func TestSubTypeConstants(t *testing.T) {
	tests := []struct {
		subType SubType
		want    int32
	}{
		{SubType_Basic, 1},
		{SubType_OrderBook, 2},
		{SubType_Ticker, 3},
		{SubType_KL, 4},
		{SubType_RT, 5},
		{SubType_Broker, 6},
	}

	for _, tc := range tests {
		if int32(tc.subType) != tc.want {
			t.Errorf("expected SubType %v = %d, got %d", tc.subType, tc.want, int32(tc.subType))
		}
	}
}

func TestProtoIDConstants(t *testing.T) {
	tests := []struct {
		name  string
		value int
	}{
		{"ProtoID_GetBasicQot", 3004},
		{"ProtoID_GetKL", 3006},
		{"ProtoID_RequestHistoryKL", 3103},
		{"ProtoID_GetOrderBook", 3012},
		{"ProtoID_GetTicker", 3010},
		{"ProtoID_GetRT", 3008},
		{"ProtoID_GetMarketSnapshot", 3203},
		{"ProtoID_GetSecuritySnapshot", 3203},
		{"ProtoID_GetBroker", 3014},
		{"ProtoID_GetStaticInfo", 2201},
		{"ProtoID_GetPlateSet", 2202},
		{"ProtoID_GetPlateSecurity", 2203},
		{"ProtoID_GetOwnerPlate", 3207},
		{"ProtoID_GetReference", 3206},
		{"ProtoID_GetTradeDate", 2205},
		{"ProtoID_RequestTradeDate", 3219},
		{"ProtoID_GetMarketState", 3223},
		{"ProtoID_GetSuspend", 2209},
		{"ProtoID_GetCodeChange", 2210},
		{"ProtoID_GetFutureInfo", 2211},
		{"ProtoID_GetIpoList", 2212},
		{"ProtoID_GetHoldingChangeList", 2213},
		{"ProtoID_RequestRehab", 2214},
		{"ProtoID_GetCapitalFlow", 3211},
		{"ProtoID_GetCapitalDistribution", 3212},
		{"ProtoID_StockFilter", 3215},
		{"ProtoID_GetOptionChain", 3209},
		{"ProtoID_GetOptionExpirationDate", 3224},
		{"ProtoID_GetWarrant", 3210},
		{"ProtoID_GetUserSecurity", 2401},
		{"ProtoID_GetUserSecurityGroup", 2402},
		{"ProtoID_ModifyUserSecurity", 2403},
		{"ProtoID_GetPriceReminder", 2404},
		{"ProtoID_SetPriceReminder", 2405},
		{"ProtoID_Subscribe", 3001},
		{"ProtoID_GetSubInfo", 3003},
		{"ProtoID_RegQotPush", 3003},
		{"ProtoID_RequestHistoryKLQuota", 3104},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got int
			switch tc.name {
			case "ProtoID_GetBasicQot":
				got = ProtoID_GetBasicQot
			case "ProtoID_GetKL":
				got = ProtoID_GetKL
			case "ProtoID_RequestHistoryKL":
				got = ProtoID_RequestHistoryKL
			case "ProtoID_GetOrderBook":
				got = ProtoID_GetOrderBook
			case "ProtoID_GetTicker":
				got = ProtoID_GetTicker
			case "ProtoID_GetRT":
				got = ProtoID_GetRT
			case "ProtoID_GetMarketSnapshot":
				got = ProtoID_GetMarketSnapshot
			case "ProtoID_GetSecuritySnapshot":
				got = ProtoID_GetSecuritySnapshot
			case "ProtoID_GetBroker":
				got = ProtoID_GetBroker
			case "ProtoID_GetStaticInfo":
				got = ProtoID_GetStaticInfo
			case "ProtoID_GetPlateSet":
				got = ProtoID_GetPlateSet
			case "ProtoID_GetPlateSecurity":
				got = ProtoID_GetPlateSecurity
			case "ProtoID_GetOwnerPlate":
				got = ProtoID_GetOwnerPlate
			case "ProtoID_GetReference":
				got = ProtoID_GetReference
			case "ProtoID_GetTradeDate":
				got = ProtoID_GetTradeDate
			case "ProtoID_RequestTradeDate":
				got = ProtoID_RequestTradeDate
			case "ProtoID_GetMarketState":
				got = ProtoID_GetMarketState
			case "ProtoID_GetSuspend":
				got = ProtoID_GetSuspend
			case "ProtoID_GetCodeChange":
				got = ProtoID_GetCodeChange
			case "ProtoID_GetFutureInfo":
				got = ProtoID_GetFutureInfo
			case "ProtoID_GetIpoList":
				got = ProtoID_GetIpoList
			case "ProtoID_GetHoldingChangeList":
				got = ProtoID_GetHoldingChangeList
			case "ProtoID_RequestRehab":
				got = ProtoID_RequestRehab
			case "ProtoID_GetCapitalFlow":
				got = ProtoID_GetCapitalFlow
			case "ProtoID_GetCapitalDistribution":
				got = ProtoID_GetCapitalDistribution
			case "ProtoID_StockFilter":
				got = ProtoID_StockFilter
			case "ProtoID_GetOptionChain":
				got = ProtoID_GetOptionChain
			case "ProtoID_GetOptionExpirationDate":
				got = ProtoID_GetOptionExpirationDate
			case "ProtoID_GetWarrant":
				got = ProtoID_GetWarrant
			case "ProtoID_GetUserSecurity":
				got = ProtoID_GetUserSecurity
			case "ProtoID_GetUserSecurityGroup":
				got = ProtoID_GetUserSecurityGroup
			case "ProtoID_ModifyUserSecurity":
				got = ProtoID_ModifyUserSecurity
			case "ProtoID_GetPriceReminder":
				got = ProtoID_GetPriceReminder
			case "ProtoID_SetPriceReminder":
				got = ProtoID_SetPriceReminder
			case "ProtoID_Subscribe":
				got = ProtoID_Subscribe
			case "ProtoID_GetSubInfo":
				got = ProtoID_GetSubInfo
			case "ProtoID_RegQotPush":
				got = ProtoID_RegQotPush
			case "ProtoID_RequestHistoryKLQuota":
				got = ProtoID_RequestHistoryKLQuota
			}
			if got != tc.value {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.value, got)
			}
		})
	}
}

func TestGetStaticInfoRequestConstruction(t *testing.T) {
	req := &GetStaticInfoRequest{
		Market:  1,
		SecType: 1,
		SecurityList: []*qotcommon.Security{
			{Market: func() *int32 { v := int32(1); return &v }(), Code: func() *string { s := "00700"; return &s }()},
		},
	}

	if req.Market != 1 {
		t.Errorf("expected Market 1, got %d", req.Market)
	}
	if req.SecType != 1 {
		t.Errorf("expected SecType 1, got %d", req.SecType)
	}
	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
}

func TestGetPlateSetRequestConstruction(t *testing.T) {
	req := &GetPlateSetRequest{
		Market: int32(qotcommon.QotMarket_QotMarket_HK_Security),
	}

	if req.Market != int32(qotcommon.QotMarket_QotMarket_HK_Security) {
		t.Errorf("expected HK market, got %d", req.Market)
	}
}

func TestGetPlateSecurityRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	plate := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "BK1001"; return &s }()}
	req := &GetPlateSecurityRequest{
		Plate: plate,
	}

	if req.Plate.GetCode() != "BK1001" {
		t.Errorf("expected plate code BK1001, got %s", req.Plate.GetCode())
	}
}

func TestPlateFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	plate := &Plate{
		Plate: &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "BK1001"; return &s }()},
		Name:  "Technology",
	}

	if plate.Name != "Technology" {
		t.Errorf("expected Name Technology, got %s", plate.Name)
	}
	if plate.Plate.GetCode() != "BK1001" {
		t.Errorf("expected code BK1001, got %s", plate.Plate.GetCode())
	}
}

func TestGetTradeDateRequestConstruction(t *testing.T) {
	req := &GetTradeDateRequest{
		Market:    int32(qotcommon.QotMarket_QotMarket_HK_Security),
		BeginTime: "2026-01-01",
		EndTime:   "2026-12-31",
	}

	if req.BeginTime != "2026-01-01" {
		t.Errorf("expected BeginTime 2026-01-01, got %s", req.BeginTime)
	}
	if req.EndTime != "2026-12-31" {
		t.Errorf("expected EndTime 2026-12-31, got %s", req.EndTime)
	}
}

func TestRequestTradeDateRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &RequestTradeDateRequest{
		Market:    int32(qotcommon.QotMarket_QotMarket_HK_Security),
		BeginTime: "2026-01-01",
		EndTime:   "2026-12-31",
		Security:  security,
	}

	if req.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", req.Security.GetCode())
	}
}

func TestRequestHistoryKLRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &RequestHistoryKLRequest{
		RehabType:    0,
		KlType:       4,
		Security:     security,
		BeginTime:    "2026-01-01",
		EndTime:      "2026-04-08",
		MaxAckKLNum:  100,
		NextReqKey:   []byte{},
		ExtendedTime: false,
	}

	if req.KlType != 4 {
		t.Errorf("expected KlType 4, got %d", req.KlType)
	}
	if req.MaxAckKLNum != 100 {
		t.Errorf("expected MaxAckKLNum 100, got %d", req.MaxAckKLNum)
	}
	if req.ExtendedTime != false {
		t.Errorf("expected ExtendedTime false, got %v", req.ExtendedTime)
	}
}

func TestRequestHistoryKLResponseConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	rsp := &RequestHistoryKLResponse{
		Security:   security,
		Name:       "Tencent",
		KLList:     []*qotcommon.KLine{},
		NextReqKey: []byte("nextkey"),
	}

	if rsp.Name != "Tencent" {
		t.Errorf("expected Name Tencent, got %s", rsp.Name)
	}
	if string(rsp.NextReqKey) != "nextkey" {
		t.Errorf("expected next key 'nextkey', got %s", rsp.NextReqKey)
	}
}

func TestGetSecuritySnapshotRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetSecuritySnapshotRequest{
		SecurityList: []*qotcommon.Security{security},
	}

	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
}

func TestSubscribeRequestFullConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &SubscribeRequest{
		SecurityList:         []*qotcommon.Security{security},
		SubTypeList:          []SubType{SubType_Basic, SubType_KL, SubType_OrderBook, SubType_Ticker, SubType_RT, SubType_Broker},
		IsSubOrUnSub:         true,
		IsRegOrUnRegPush:     true,
		RegPushRehabTypeList: []int32{0},
		IsFirstPush:          true,
		IsUnsubAll:           false,
	}

	if len(req.SubTypeList) != 6 {
		t.Errorf("expected 6 sub types, got %d", len(req.SubTypeList))
	}
	if !req.IsFirstPush {
		t.Error("expected IsFirstPush true")
	}
	if req.IsUnsubAll {
		t.Error("expected IsUnsubAll false")
	}
}

func TestSubscribeResponseConstruction(t *testing.T) {
	rsp := &SubscribeResponse{
		RetType: 0,
		RetMsg:  "success",
	}

	if rsp.RetType != 0 {
		t.Errorf("expected RetType 0, got %d", rsp.RetType)
	}
	if rsp.RetMsg != "success" {
		t.Errorf("expected RetMsg success, got %s", rsp.RetMsg)
	}
}

func TestGetCapitalFlowRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetCapitalFlowRequest{
		Security:   security,
		PeriodType: 1,
		BeginTime:  "2026-01-01",
		EndTime:    "2026-04-08",
	}

	if req.PeriodType != 1 {
		t.Errorf("expected PeriodType 1, got %d", req.PeriodType)
	}
}

func TestCapitalFlowItemFields(t *testing.T) {
	item := &CapitalFlowItem{
		InFlow:      1000000.0,
		Time:        "2026-04-08",
		Timestamp:   1775635200.0,
		MainInFlow:  800000.0,
		SuperInFlow: 200000.0,
		BigInFlow:   300000.0,
		MidInFlow:   250000.0,
		SmlInFlow:   250000.0,
	}

	if item.InFlow != 1000000.0 {
		t.Errorf("expected InFlow 1000000.0, got %f", item.InFlow)
	}
	if item.MainInFlow != 800000.0 {
		t.Errorf("expected MainInFlow 800000.0, got %f", item.MainInFlow)
	}
}

func TestGetCapitalFlowResponseConstruction(t *testing.T) {
	rsp := &GetCapitalFlowResponse{
		FlowItemList: []*CapitalFlowItem{
			{InFlow: 1000000.0, Time: "2026-04-08"},
		},
		LastValidTime:      "2026-04-08",
		LastValidTimestamp: 1775635200.0,
	}

	if len(rsp.FlowItemList) != 1 {
		t.Errorf("expected 1 flow item, got %d", len(rsp.FlowItemList))
	}
	if rsp.LastValidTime != "2026-04-08" {
		t.Errorf("expected LastValidTime 2026-04-08, got %s", rsp.LastValidTime)
	}
}

func TestCapitalDistributionFields(t *testing.T) {
	cd := &CapitalDistribution{
		CapitalInSuper:  1000000.0,
		CapitalInBig:    500000.0,
		CapitalInMid:    300000.0,
		CapitalInSmall:  200000.0,
		CapitalOutSuper: 800000.0,
		CapitalOutBig:   400000.0,
		CapitalOutMid:   200000.0,
		CapitalOutSmall: 100000.0,
		UpdateTime:      "2026-04-08 15:00:00",
		UpdateTimestamp: 1775635200.0,
	}

	if cd.CapitalInSuper != 1000000.0 {
		t.Errorf("expected CapitalInSuper 1000000.0, got %f", cd.CapitalInSuper)
	}
	if cd.CapitalOutSmall != 100000.0 {
		t.Errorf("expected CapitalOutSmall 100000.0, got %f", cd.CapitalOutSmall)
	}
}

func TestGetPriceReminderResponseConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	rsp := &GetPriceReminderResponse{
		PriceReminderList: []*PriceReminderInfo{
			{
				Security: security,
				Name:     "Tencent",
				ItemList: []*PriceReminderItemInfo{
					{Key: 1, Type: 1, Value: 350.0, Note: "Alert", Freq: 0, IsEnable: true},
				},
			},
		},
	}

	if len(rsp.PriceReminderList) != 1 {
		t.Errorf("expected 1 price reminder, got %d", len(rsp.PriceReminderList))
	}
	if len(rsp.PriceReminderList[0].ItemList) != 1 {
		t.Errorf("expected 1 item, got %d", len(rsp.PriceReminderList[0].ItemList))
	}
	if rsp.PriceReminderList[0].ItemList[0].Value != 350.0 {
		t.Errorf("expected Value 350.0, got %f", rsp.PriceReminderList[0].ItemList[0].Value)
	}
}

func TestPriceReminderItemInfoFields(t *testing.T) {
	item := &PriceReminderItemInfo{
		Key:      123,
		Type:     1,
		Value:    350.50,
		Note:     "Test alert",
		Freq:     0,
		IsEnable: true,
	}

	if item.Key != 123 {
		t.Errorf("expected Key 123, got %d", item.Key)
	}
	if item.Type != 1 {
		t.Errorf("expected Type 1, got %d", item.Type)
	}
	if item.Value != 350.50 {
		t.Errorf("expected Value 350.50, got %f", item.Value)
	}
	if !item.IsEnable {
		t.Error("expected IsEnable true")
	}
}

func TestGetOptionExpirationDateRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	owner := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "HSI"; return &s }()}
	req := &GetOptionExpirationDateRequest{
		Owner:           owner,
		IndexOptionType: 0,
	}

	if req.Owner.GetCode() != "HSI" {
		t.Errorf("expected owner HSI, got %s", req.Owner.GetCode())
	}
}

func TestOptionExpirationDateInfoFields(t *testing.T) {
	info := &OptionExpirationDateInfo{
		StrikeTime:               "2026-04-30",
		StrikeTimestamp:          1746057600.0,
		OptionExpiryDateDistance: 22,
		Cycle:                    1,
	}

	if info.StrikeTime != "2026-04-30" {
		t.Errorf("expected StrikeTime 2026-04-30, got %s", info.StrikeTime)
	}
	if info.OptionExpiryDateDistance != 22 {
		t.Errorf("expected OptionExpiryDateDistance 22, got %d", info.OptionExpiryDateDistance)
	}
}

func TestGetOptionExpirationDateResponseConstruction(t *testing.T) {
	rsp := &GetOptionExpirationDateResponse{
		DateList: []*OptionExpirationDateInfo{
			{StrikeTime: "2026-04-30", StrikeTimestamp: 1746057600.0},
		},
	}

	if len(rsp.DateList) != 1 {
		t.Errorf("expected 1 date, got %d", len(rsp.DateList))
	}
}

func TestGetOptionChainRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	owner := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "HSI"; return &s }()}
	req := &GetOptionChainRequest{
		Owner:           owner,
		IndexOptionType: 0,
		Type:            1,
		Condition:       1,
		BeginTime:       "2026-04-01",
		EndTime:         "2026-04-30",
	}

	if req.Type != 1 {
		t.Errorf("expected Type 1, got %d", req.Type)
	}
	if req.Condition != 1 {
		t.Errorf("expected Condition 1, got %d", req.Condition)
	}
}

func TestOptionItemFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	callCode := "HSI2405C35000"
	putCode := "HSI2405P35000"
	callBasic := &qotcommon.SecurityStaticBasic{
		Security: &qotcommon.Security{Market: &hkMarket, Code: &callCode},
	}
	putBasic := &qotcommon.SecurityStaticBasic{
		Security: &qotcommon.Security{Market: &hkMarket, Code: &putCode},
	}
	call := &qotcommon.SecurityStaticInfo{Basic: callBasic}
	put := &qotcommon.SecurityStaticInfo{Basic: putBasic}
	item := &OptionItem{
		Call: call,
		Put:  put,
	}

	if item.Call.GetBasic().GetSecurity().GetCode() != "HSI2405C35000" {
		t.Errorf("expected call code HSI2405C35000, got %s", item.Call.GetBasic().GetSecurity().GetCode())
	}
	if item.Put.GetBasic().GetSecurity().GetCode() != "HSI2405P35000" {
		t.Errorf("expected put code HSI2405P35000, got %s", item.Put.GetBasic().GetSecurity().GetCode())
	}
}

func TestOptionChainFields(t *testing.T) {
	chain := &OptionChain{
		StrikeTime:      "2026-04-30",
		StrikeTimestamp: 1746057600.0,
		Option:          []*OptionItem{},
	}

	if chain.StrikeTime != "2026-04-30" {
		t.Errorf("expected StrikeTime 2026-04-30, got %s", chain.StrikeTime)
	}
}

func TestGetOptionChainResponseConstruction(t *testing.T) {
	rsp := &GetOptionChainResponse{
		OptionChain: []*OptionChain{
			{StrikeTime: "2026-04-30", StrikeTimestamp: 1746057600.0},
		},
	}

	if len(rsp.OptionChain) != 1 {
		t.Errorf("expected 1 chain, got %d", len(rsp.OptionChain))
	}
}

func TestStockFilterRequestConstruction(t *testing.T) {
	req := &StockFilterRequest{
		Begin:  0,
		Num:    100,
		Market: int32(qotcommon.QotMarket_QotMarket_HK_Security),
	}

	if req.Begin != 0 {
		t.Errorf("expected Begin 0, got %d", req.Begin)
	}
	if req.Num != 100 {
		t.Errorf("expected Num 100, got %d", req.Num)
	}
}

func TestStockFilterDataFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	data := &StockFilterData{
		Security: &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()},
		Name:     "Tencent",
	}

	if data.Name != "Tencent" {
		t.Errorf("expected Name Tencent, got %s", data.Name)
	}
}

func TestStockFilterResponseConstruction(t *testing.T) {
	rsp := &StockFilterResponse{
		LastPage: true,
		AllCount: 50,
		DataList: []*StockFilterData{
			{Name: "Tencent"},
		},
	}

	if !rsp.LastPage {
		t.Error("expected LastPage true")
	}
	if rsp.AllCount != 50 {
		t.Errorf("expected AllCount 50, got %d", rsp.AllCount)
	}
}

func TestGetWarrantRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	owner := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetWarrantRequest{
		Begin:      0,
		Num:        100,
		SortField:  1,
		Ascend:     true,
		Owner:      owner,
		TypeList:   []int32{1, 2},
		IssuerList: []int32{1},
	}

	if req.SortField != 1 {
		t.Errorf("expected SortField 1, got %d", req.SortField)
	}
	if !req.Ascend {
		t.Error("expected Ascend true")
	}
}

func TestWarrantDataFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	wd := &WarrantData{
		Stock:             &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "WT00700"; return &s }()},
		Owner:             &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()},
		Type:              1,
		Issuer:            1,
		MaturityTime:      "2026-04-30",
		MaturityTimestamp: 1746057600.0,
		CurPrice:          0.5,
		ChangeRate:        5.0,
		Volume:            10000,
		Turnover:          5000.0,
		Premium:           10.0,
		Leverage:          5.0,
		Delta:             0.5,
	}

	if wd.Type != 1 {
		t.Errorf("expected Type 1, got %d", wd.Type)
	}
	if wd.Premium != 10.0 {
		t.Errorf("expected Premium 10.0, got %f", wd.Premium)
	}
	if wd.Leverage != 5.0 {
		t.Errorf("expected Leverage 5.0, got %f", wd.Leverage)
	}
}

func TestGetSuspendRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetSuspendRequest{
		SecurityList: []*qotcommon.Security{security},
		BeginTime:    "2026-01-01",
		EndTime:      "2026-04-08",
	}

	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
}

func TestSuspendInfoFields(t *testing.T) {
	si := &SuspendInfo{
		Time:      "2026-03-01",
		Timestamp: 1773043200.0,
	}

	if si.Time != "2026-03-01" {
		t.Errorf("expected Time 2026-03-01, got %s", si.Time)
	}
}

func TestSecuritySuspendInfoFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	ssi := &SecuritySuspendInfo{
		Security: &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()},
		SuspendList: []*SuspendInfo{
			{Time: "2026-03-01", Timestamp: 1773043200.0},
		},
	}

	if ssi.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", ssi.Security.GetCode())
	}
	if len(ssi.SuspendList) != 1 {
		t.Errorf("expected 1 suspend entry, got %d", len(ssi.SuspendList))
	}
}

func TestGetFutureInfoRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "HSI2405"; return &s }()}
	req := &GetFutureInfoRequest{
		SecurityList: []*qotcommon.Security{security},
	}

	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
}

func TestFutureInfoFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	fi := &FutureInfo{
		Name:               "HSI2405",
		Security:           &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "HSI2405"; return &s }()},
		LastTradeTime:      "2026-04-30",
		LastTradeTimestamp: 1746057600.0,
		ContractSize:       50.0,
		MinVar:             1.0,
		TimeZone:           "UTC+8",
	}

	if fi.Name != "HSI2405" {
		t.Errorf("expected Name HSI2405, got %s", fi.Name)
	}
	if fi.ContractSize != 50.0 {
		t.Errorf("expected ContractSize 50.0, got %f", fi.ContractSize)
	}
}

func TestGetCodeChangeRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetCodeChangeRequest{
		SecurityList: []*qotcommon.Security{security},
		TypeList:     []int32{1, 2},
	}

	if len(req.TypeList) != 2 {
		t.Errorf("expected 2 types, got %d", len(req.TypeList))
	}
}

func TestCodeChangeInfoFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	cc := &CodeChangeInfo{
		Type:               1,
		Security:           &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()},
		RelatedSecurity:    &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700_H"; return &s }()},
		PublicTime:         "2026-03-01",
		PublicTimestamp:    1773043200.0,
		EffectiveTime:      "2026-04-01",
		EffectiveTimestamp: 1777776000.0,
	}

	if cc.Type != 1 {
		t.Errorf("expected Type 1, got %d", cc.Type)
	}
	if cc.RelatedSecurity.GetCode() != "00700_H" {
		t.Errorf("expected related code 00700_H, got %s", cc.RelatedSecurity.GetCode())
	}
}

func TestGetIpoListRequestConstruction(t *testing.T) {
	req := &GetIpoListRequest{
		Market: int32(qotcommon.QotMarket_QotMarket_HK_Security),
	}

	if req.Market != int32(qotcommon.QotMarket_QotMarket_HK_Security) {
		t.Errorf("expected HK market, got %d", req.Market)
	}
}

func TestBasicIpoDataFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	basic := &BasicIpoData{
		Security:      &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "NEWCO"; return &s }()},
		Name:          "New Company",
		ListTime:      "2026-04-15",
		ListTimestamp: 1744675200.0,
	}

	if basic.Name != "New Company" {
		t.Errorf("expected Name New Company, got %s", basic.Name)
	}
}

func TestCNIpoExDataFields(t *testing.T) {
	cn := &CNIpoExData{
		ApplyCode:              "786543",
		IssueSize:              100000000,
		OnlineIssueSize:        20000000,
		ApplyUpperLimit:        10000,
		IsEstimateIpoPrice:     true,
		IpoPrice:               20.0,
		IsEstimateWinningRatio: true,
		WinningRatio:           0.5,
		IsHasWon:               false,
	}

	if cn.ApplyCode != "786543" {
		t.Errorf("expected ApplyCode 786543, got %s", cn.ApplyCode)
	}
	if cn.IsEstimateIpoPrice != true {
		t.Error("expected IsEstimateIpoPrice true")
	}
}

func TestHKIpoExDataFields(t *testing.T) {
	hk := &HKIpoExData{
		IpoPriceMin:       50.0,
		IpoPriceMax:       60.0,
		ListPrice:         55.0,
		LotSize:           100,
		EntrancePrice:     55.5,
		IsSubscribeStatus: true,
	}

	if hk.IpoPriceMax != 60.0 {
		t.Errorf("expected IpoPriceMax 60.0, got %f", hk.IpoPriceMax)
	}
	if hk.LotSize != 100 {
		t.Errorf("expected LotSize 100, got %d", hk.LotSize)
	}
}

func TestUSIpoExDataFields(t *testing.T) {
	us := &USIpoExData{
		IpoPriceMin: 20.0,
		IpoPriceMax: 25.0,
		IssueSize:   50000000,
	}

	if us.IssueSize != 50000000 {
		t.Errorf("expected IssueSize 50000000, got %d", us.IssueSize)
	}
}

func TestIpoDataFields(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	ipo := &IpoData{
		Basic: &BasicIpoData{
			Security: &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "NEWCO"; return &s }()},
			Name:     "New Company",
		},
		HkExData: &HKIpoExData{
			IpoPriceMin: 50.0,
			IpoPriceMax: 60.0,
			LotSize:     100,
		},
	}

	if ipo.Basic.Name != "New Company" {
		t.Errorf("expected Name New Company, got %s", ipo.Basic.Name)
	}
	if ipo.HkExData.IpoPriceMax != 60.0 {
		t.Errorf("expected IpoPriceMax 60.0, got %f", ipo.HkExData.IpoPriceMax)
	}
}

func TestGetHoldingChangeListRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &GetHoldingChangeListRequest{
		Security:       security,
		HolderCategory: 1,
		BeginTime:      "2026-01-01",
		EndTime:        "2026-04-08",
	}

	if req.HolderCategory != 1 {
		t.Errorf("expected HolderCategory 1, got %d", req.HolderCategory)
	}
}

func TestGetUserSecurityGroupRequestConstruction(t *testing.T) {
	req := &GetUserSecurityGroupRequest{
		GroupType: 1,
	}

	if req.GroupType != 1 {
		t.Errorf("expected GroupType 1, got %d", req.GroupType)
	}
}

func TestUserSecurityGroupDataFields(t *testing.T) {
	group := &UserSecurityGroupData{
		GroupName: "My Watchlist",
		GroupType: 1,
	}

	if group.GroupName != "My Watchlist" {
		t.Errorf("expected GroupName My Watchlist, got %s", group.GroupName)
	}
}

func TestModifyUserSecurityRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &ModifyUserSecurityRequest{
		GroupName:    "My Watchlist",
		Op:           1,
		SecurityList: []*qotcommon.Security{security},
	}

	if req.Op != 1 {
		t.Errorf("expected Op 1, got %d", req.Op)
	}
	if len(req.SecurityList) != 1 {
		t.Errorf("expected 1 security, got %d", len(req.SecurityList))
	}
}

func TestSetPriceReminderRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &SetPriceReminderRequest{
		Security: security,
		Op:       1,
		Key:      123,
		Type:     1,
		Freq:     0,
		Value:    350.0,
		Note:     "Alert when price reaches 350",
	}

	if req.Value != 350.0 {
		t.Errorf("expected Value 350.0, got %f", req.Value)
	}
}

func TestSetPriceReminderResponseConstruction(t *testing.T) {
	rsp := &SetPriceReminderResponse{
		Key: 123,
	}

	if rsp.Key != 123 {
		t.Errorf("expected Key 123, got %d", rsp.Key)
	}
}

func TestRegQotPushRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &RegQotPushRequest{
		SecurityList:  []*qotcommon.Security{security},
		SubTypeList:   []int32{1, 2},
		RehabTypeList: []int32{0},
		IsRegOrUnReg:  true,
		IsFirstPush:   true,
	}

	if !req.IsRegOrUnReg {
		t.Error("expected IsRegOrUnReg true")
	}
	if !req.IsFirstPush {
		t.Error("expected IsFirstPush true")
	}
}

func TestRegQotPushResponseConstruction(t *testing.T) {
	rsp := &RegQotPushResponse{
		RetType: 0,
		RetMsg:  "success",
	}

	if rsp.RetType != 0 {
		t.Errorf("expected RetType 0, got %d", rsp.RetType)
	}
}

func TestRequestRehabRequestConstruction(t *testing.T) {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{Market: &hkMarket, Code: func() *string { s := "00700"; return &s }()}
	req := &RequestRehabRequest{
		Security: security,
	}

	if req.Security.GetCode() != "00700" {
		t.Errorf("expected code 00700, got %s", req.Security.GetCode())
	}
}

func TestRequestHistoryKLQuotaRequestConstruction(t *testing.T) {
	req := &RequestHistoryKLQuotaRequest{
		GetDetail: true,
	}

	if !req.GetDetail {
		t.Error("expected GetDetail true")
	}
}

func TestRequestHistoryKLQuotaResponseConstruction(t *testing.T) {
	rsp := &RequestHistoryKLQuotaResponse{
		UsedQuota:   50,
		RemainQuota: 450,
	}

	if rsp.UsedQuota != 50 {
		t.Errorf("expected UsedQuota 50, got %d", rsp.UsedQuota)
	}
	if rsp.RemainQuota != 450 {
		t.Errorf("expected RemainQuota 450, got %d", rsp.RemainQuota)
	}
}

