package simulator

import (
	"google.golang.org/protobuf/proto"

	"gitee.com/shing1211/futuapi4go/pkg/pb/common"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgetacclist"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgetfunds"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderfilllist"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgethistoryorderlist"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgetmarginratio"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgetmaxtrdqtys"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgetorderfee"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgetorderfilllist"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgetorderlist"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdgetpositionlist"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdplaceorder"
	"gitee.com/shing1211/futuapi4go/pkg/pb/trdunlocktrade"
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
	resp := &trdgetfunds.Response{RetType: &retType}
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
	orderID := uint64(1234567890)

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
	resp := &trdgetorderlist.Response{RetType: &retType}
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
	resp := &trdgetpositionlist.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}
