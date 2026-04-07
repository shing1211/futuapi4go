package trd

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "gitee.com/shing1211/futuapi4go/client"
	"github.com/futuopen/ftapi4go/pb/common"
	"github.com/futuopen/ftapi4go/pb/trdcommon"
	"github.com/futuopen/ftapi4go/pb/trdgetacclist"
	"github.com/futuopen/ftapi4go/pb/trdgetfunds"
	"github.com/futuopen/ftapi4go/pb/trdgetorderfilllist"
	"github.com/futuopen/ftapi4go/pb/trdgetorderlist"
	"github.com/futuopen/ftapi4go/pb/trdgetpositionlist"
	"github.com/futuopen/ftapi4go/pb/trdmodifyorder"
	"github.com/futuopen/ftapi4go/pb/trdplaceorder"
	"github.com/futuopen/ftapi4go/pb/trdunlocktrade"
)

const (
	ProtoID_GetAccList       = 4001
	ProtoID_UnlockTrade      = 4002
	ProtoID_GetFunds         = 4003
	ProtoID_GetPositionList  = 6001
	ProtoID_GetOrderList     = 5003
	ProtoID_GetOrderFillList = 5005
	ProtoID_PlaceOrder       = 5001
	ProtoID_ModifyOrder      = 5002
)

type Acc struct {
	TrdEnv    int32
	AccID     uint64
	AccType   int32
	CardNum   string
	AccStatus int32
}

type GetAccListResponse struct {
	AccList []*Acc
}

func GetAccList(c *futuapi.Client, trdCategory int32, needGeneralSecAccount bool) (*GetAccListResponse, error) {
	c2s := &trdgetacclist.C2S{
		TrdCategory:           &trdCategory,
		NeedGeneralSecAccount: &needGeneralSecAccount,
	}

	pkt := &trdgetacclist.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetAccList, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp trdgetacclist.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetAccList failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetAccList: s2c is nil")
	}

	result := &GetAccListResponse{
		AccList: make([]*Acc, 0, len(s2c.GetAccList())),
	}

	for _, acc := range s2c.GetAccList() {
		result.AccList = append(result.AccList, &Acc{
			TrdEnv:    acc.GetTrdEnv(),
			AccID:     acc.GetAccID(),
			AccType:   acc.GetAccType(),
			CardNum:   acc.GetCardNum(),
			AccStatus: acc.GetAccStatus(),
		})
	}

	return result, nil
}

type Funds struct {
	Power          float64
	TotalAssets    float64
	Cash           float64
	MarketVal      float64
	FrozenCash     float64
	DebtCash       float64
	Currency       int32
	AvailableFunds float64
}

type GetFundsRequest struct {
	AccID     uint64
	TrdMarket int32
}

type GetFundsResponse struct {
	Funds *Funds
}

func GetFunds(c *futuapi.Client, req *GetFundsRequest) (*GetFundsResponse, error) {
	header := &trdcommon.TrdHeader{
		AccID:     &req.AccID,
		TrdMarket: &req.TrdMarket,
	}

	c2s := &trdgetfunds.C2S{
		Header: header,
	}

	pkt := &trdgetfunds.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetFunds, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp trdgetfunds.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetFunds failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetFunds: s2c is nil")
	}

	f := s2c.GetFunds()
	return &GetFundsResponse{
		Funds: &Funds{
			Power:          f.GetPower(),
			TotalAssets:    f.GetTotalAssets(),
			Cash:           f.GetCash(),
			MarketVal:      f.GetMarketVal(),
			FrozenCash:     f.GetFrozenCash(),
			DebtCash:       f.GetDebtCash(),
			Currency:       f.GetCurrency(),
			AvailableFunds: f.GetAvailableFunds(),
		},
	}, nil
}

type Position struct {
	Code       string
	Name       string
	Qty        float64
	CanSellQty float64
	Price      float64
	CostPrice  float64
	Val        float64
	PlVal      float64
	PlRatio    float64
}

type GetPositionListRequest struct {
	AccID     uint64
	TrdMarket int32
}

type GetPositionListResponse struct {
	PositionList []*Position
}

func GetPositionList(c *futuapi.Client, req *GetPositionListRequest) (*GetPositionListResponse, error) {
	header := &trdcommon.TrdHeader{
		AccID:     &req.AccID,
		TrdMarket: &req.TrdMarket,
	}

	c2s := &trdgetpositionlist.C2S{
		Header: header,
	}

	pkt := &trdgetpositionlist.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetPositionList, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp trdgetpositionlist.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetPositionList failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetPositionList: s2c is nil")
	}

	result := &GetPositionListResponse{
		PositionList: make([]*Position, 0, len(s2c.GetPositionList())),
	}

	for _, p := range s2c.GetPositionList() {
		result.PositionList = append(result.PositionList, &Position{
			Code:       p.GetCode(),
			Name:       p.GetName(),
			Qty:        p.GetQty(),
			CanSellQty: p.GetCanSellQty(),
			Price:      p.GetPrice(),
			CostPrice:  p.GetCostPrice(),
			Val:        p.GetVal(),
			PlVal:      p.GetPlVal(),
			PlRatio:    p.GetPlRatio(),
		})
	}

	return result, nil
}

type Order struct {
	OrderID      uint64
	Code         string
	Name         string
	TrdSide      int32
	OrderType    int32
	OrderStatus  int32
	Price        float64
	Qty          float64
	FillQty      float64
	CreateTime   string
	UpdateTime   string
	FillAvgPrice float64
}

type GetOrderListRequest struct {
	AccID     uint64
	TrdMarket int32
}

type GetOrderListResponse struct {
	OrderList []*Order
}

func GetOrderList(c *futuapi.Client, req *GetOrderListRequest) (*GetOrderListResponse, error) {
	header := &trdcommon.TrdHeader{
		AccID:     &req.AccID,
		TrdMarket: &req.TrdMarket,
	}

	c2s := &trdgetorderlist.C2S{
		Header: header,
	}

	pkt := &trdgetorderlist.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetOrderList, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp trdgetorderlist.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetOrderList failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOrderList: s2c is nil")
	}

	result := &GetOrderListResponse{
		OrderList: make([]*Order, 0, len(s2c.GetOrderList())),
	}

	for _, o := range s2c.GetOrderList() {
		result.OrderList = append(result.OrderList, &Order{
			OrderID:      o.GetOrderID(),
			Code:         o.GetCode(),
			Name:         o.GetName(),
			TrdSide:      o.GetTrdSide(),
			OrderType:    o.GetOrderType(),
			OrderStatus:  o.GetOrderStatus(),
			Price:        o.GetPrice(),
			Qty:          o.GetQty(),
			FillQty:      o.GetFillQty(),
			CreateTime:   o.GetCreateTime(),
			UpdateTime:   o.GetUpdateTime(),
			FillAvgPrice: o.GetFillAvgPrice(),
		})
	}

	return result, nil
}

type OrderFill struct {
	OrderID    uint64
	FillID     uint64
	Code       string
	Name       string
	TrdSide    int32
	Price      float64
	Qty        float64
	CreateTime string
}

type GetOrderFillListRequest struct {
	AccID     uint64
	TrdMarket int32
}

type GetOrderFillListResponse struct {
	OrderFillList []*OrderFill
}

func GetOrderFillList(c *futuapi.Client, req *GetOrderFillListRequest) (*GetOrderFillListResponse, error) {
	header := &trdcommon.TrdHeader{
		AccID:     &req.AccID,
		TrdMarket: &req.TrdMarket,
	}

	c2s := &trdgetorderfilllist.C2S{
		Header: header,
	}

	pkt := &trdgetorderfilllist.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_GetOrderFillList, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp trdgetorderfilllist.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetOrderFillList failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOrderFillList: s2c is nil")
	}

	result := &GetOrderFillListResponse{
		OrderFillList: make([]*OrderFill, 0, len(s2c.GetOrderFillList())),
	}

	for _, f := range s2c.GetOrderFillList() {
		result.OrderFillList = append(result.OrderFillList, &OrderFill{
			OrderID:    f.GetOrderID(),
			FillID:     f.GetFillID(),
			Code:       f.GetCode(),
			Name:       f.GetName(),
			TrdSide:    f.GetTrdSide(),
			Price:      f.GetPrice(),
			Qty:        f.GetQty(),
			CreateTime: f.GetCreateTime(),
		})
	}

	return result, nil
}

type PlaceOrderRequest struct {
	AccID     uint64
	TrdMarket int32
	Code      string
	TrdSide   int32
	OrderType int32
	Price     float64
	Qty       float64
}

type PlaceOrderResponse struct {
	OrderID uint64
}

func PlaceOrder(c *futuapi.Client, req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	header := &trdcommon.TrdHeader{
		AccID:     &req.AccID,
		TrdMarket: &req.TrdMarket,
	}

	c2s := &trdplaceorder.C2S{
		Header:    header,
		TrdSide:   &req.TrdSide,
		OrderType: &req.OrderType,
		Code:      &req.Code,
		Price:     &req.Price,
		Qty:       &req.Qty,
	}

	pkt := &trdplaceorder.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_PlaceOrder, serialNo, body); err != nil {
		return nil, err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return nil, err
	}

	var rsp trdplaceorder.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("PlaceOrder failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("PlaceOrder: s2c is nil")
	}

	return &PlaceOrderResponse{
		OrderID: s2c.GetOrderID(),
	}, nil
}

type ModifyOrderRequest struct {
	AccID     uint64
	TrdMarket int32
	OrderID   uint64
	Price     float64
	Qty       float64
}

func ModifyOrder(c *futuapi.Client, req *ModifyOrderRequest) error {
	header := &trdcommon.TrdHeader{
		AccID:     &req.AccID,
		TrdMarket: &req.TrdMarket,
	}

	orderID := req.OrderID
	c2s := &trdmodifyorder.C2S{
		Header:  header,
		OrderID: &orderID,
		Price:   &req.Price,
		Qty:     &req.Qty,
	}

	pkt := &trdmodifyorder.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_ModifyOrder, serialNo, body); err != nil {
		return err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return err
	}

	var rsp trdmodifyorder.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return fmt.Errorf("ModifyOrder failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return nil
}

type UnlockTradeRequest struct {
	Unlock       bool
	PwdMD5       string
	SecurityFirm int32
}

func UnlockTrade(c *futuapi.Client, req *UnlockTradeRequest) error {
	c2s := &trdunlocktrade.C2S{
		Unlock:       &req.Unlock,
		PwdMD5:       &req.PwdMD5,
		SecurityFirm: &req.SecurityFirm,
	}

	pkt := &trdunlocktrade.Request{C2S: c2s}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return err
	}

	serialNo := c.NextSerialNo()
	if err := c.Conn().WritePacket(ProtoID_UnlockTrade, serialNo, body); err != nil {
		return err
	}

	pktResp, err := c.Conn().ReadPacket()
	if err != nil {
		return err
	}

	var rsp trdunlocktrade.Response
	if err := proto.Unmarshal(pktResp.Body, &rsp); err != nil {
		return err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return fmt.Errorf("UnlockTrade failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	return nil
}
