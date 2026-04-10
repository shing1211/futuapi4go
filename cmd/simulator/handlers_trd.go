package simulator

import (
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetacclist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetfunds"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmarginratio"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetmaxtrdqtys"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfee"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetpositionlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplaceorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdunlocktrade"
)

func (s *Server) RegisterTrdHandlers() {
	s.RegisterHandler(2001, s.handleGetAccList)              // GetAccList
	s.RegisterHandler(2005, s.handleUnlockTrade)             // UnlockTrade
	s.RegisterHandler(2101, s.handleGetFunds)                // GetFunds
	s.RegisterHandler(2225, s.handleGetOrderFee)             // GetOrderFee
	s.RegisterHandler(2223, s.handleGetMarginRatio)          // GetMarginRatio
	s.RegisterHandler(2111, s.handleGetMaxTrdQtys)           // GetMaxTrdQtys
	s.RegisterHandler(2202, s.handlePlaceOrder)              // PlaceOrder
	s.RegisterHandler(2205, s.handleModifyOrder)             // ModifyOrder
	s.RegisterHandler(2201, s.handleGetOrderList)            // GetOrderList
	s.RegisterHandler(2221, s.handleGetHistoryOrderList)     // GetHistoryOrderList
	s.RegisterHandler(2211, s.handleGetOrderFillList)        // GetOrderFillList
	s.RegisterHandler(2222, s.handleGetHistoryOrderFillList) // GetHistoryOrderFillList
	s.RegisterHandler(2102, s.handleGetPositionList)         // GetPositionList
}

func (s *Server) handleGetAccList(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &trdgetacclist.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleUnlockTrade(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &trdunlocktrade.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetFunds(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)

	totalAssets := 500000.00
	cash := 250000.00
	marketVal := 74093.80
	frozenCash := 37046.90
	power := 425906.10

	s2c := &trdgetfunds.S2C{
		Funds: &trdcommon.Funds{
			TotalAssets: &totalAssets,
			Cash:        &cash,
			MarketVal:   &marketVal,
			FrozenCash:  &frozenCash,
			Power:       &power,
		},
	}
	
	resp := &trdgetfunds.Response{
		RetType: &retType,
		S2C:     s2c,
	}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetOrderFee(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &trdgetorderfee.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetMarginRatio(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &trdgetmarginratio.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetMaxTrdQtys(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &trdgetmaxtrdqtys.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handlePlaceOrder(pkt *Packet) (*Packet, error) {
	var req trdplaceorder.C2S
	if err := proto.Unmarshal(pkt.Body, &req); err != nil {
		return s.errorResponse(pkt, err)
	}

	retType := int32(common.RetType_RetType_Succeed)
	orderID := uint64(9876543210)

	// Store order for later retrieval
	if s.Orders != nil {
		code := req.GetCode()
		trdSide := req.GetTrdSide()
		orderType := req.GetOrderType()
		price := req.GetPrice()
		qty := req.GetQty()
		submittedAt := time.Now().Format("2006-01-02 15:04:05")
		orderStatus := int32(trdcommon.OrderStatus_OrderStatus_Submitted)

		order := &trdcommon.Order{
			OrderID:     &orderID,
			Code:        &code,
			TrdSide:     &trdSide,
			OrderType:   &orderType,
			OrderStatus: &orderStatus,
			Price:       &price,
			Qty:         &qty,
			CreateTime:  &submittedAt,
		}

		s.Orders[orderID] = order
	}

	resp := &trdplaceorder.Response{
		RetType: &retType,
		S2C:     &trdplaceorder.S2C{OrderID: &orderID},
	}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleModifyOrder(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &trdmodifyorder.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetOrderList(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	
	orderList := make([]*trdcommon.Order, 0)
	
	// Return stored orders if any
	if s.Orders != nil {
		for _, order := range s.Orders {
			orderList = append(orderList, order)
		}
	}
	
	// If no orders, return a sample HSI order
	if len(orderList) == 0 {
		orderID := uint64(1001)
		code := "HSImain"
		name := "HANG SENG INDEX FUTURES"
		trdSide := int32(trdcommon.TrdSide_TrdSide_Buy)
		orderType := int32(trdcommon.OrderType_OrderType_Normal)
		orderStatus := int32(trdcommon.OrderStatus_OrderStatus_Submitted)
		price := 18520.00
		qty := 1.0
		fillQty := 0.0
		createTime := time.Now().Format("2006-01-02 15:04:05")
		
		orderList = append(orderList, &trdcommon.Order{
			OrderID:     &orderID,
			Code:        &code,
			Name:        &name,
			TrdSide:     &trdSide,
			OrderType:   &orderType,
			OrderStatus: &orderStatus,
			Price:       &price,
			Qty:         &qty,
			FillQty:     &fillQty,
			CreateTime:  &createTime,
		})
	}
	
	resp := &trdgetorderlist.Response{
		RetType: &retType,
		S2C: &trdgetorderlist.S2C{
			OrderList: orderList,
		},
	}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetHistoryOrderList(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &trdgethistoryorderlist.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetOrderFillList(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &trdgetorderfilllist.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetHistoryOrderFillList(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &trdgethistoryorderfilllist.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetPositionList(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	
	positionList := make([]*trdcommon.Position, 0)
	
	// Return stored positions if any
	if s.Positions != nil {
		for _, pos := range s.Positions {
			positionList = append(positionList, pos)
		}
	}
	
	// If no positions, return a sample HSI futures position
	if len(positionList) == 0 {
		code := "HSImain"
		name := "HANG SENG INDEX FUTURES"
		qty := 2.0
		canSellQty := 2.0
		price := 18523.45
		costPrice := 18480.00
		plVal := 86.90 // (18523.45 - 18480.00) * 2
		plRatio := 0.235
		
		positionList = append(positionList, &trdcommon.Position{
			Code:       &code,
			Name:       &name,
			Qty:        &qty,
			CanSellQty: &canSellQty,
			Price:      &price,
			CostPrice:  &costPrice,
			PlVal:      &plVal,
			PlRatio:    &plRatio,
		})
	}
	
	resp := &trdgetpositionlist.Response{
		RetType: &retType,
		S2C: &trdgetpositionlist.S2C{
			PositionList: positionList,
		},
	}
	return s.successResponse(pkt, resp)
}

