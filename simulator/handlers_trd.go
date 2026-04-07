package simulator

import (
	"google.golang.org/protobuf/proto"

	"gitee.com/shing1211/futuapi4go/pb/common"
	"gitee.com/shing1211/futuapi4go/pb/trdgetacclist"
	"gitee.com/shing1211/futuapi4go/pb/trdgetfunds"
	"gitee.com/shing1211/futuapi4go/pb/trdgethistoryorderfilllist"
	"gitee.com/shing1211/futuapi4go/pb/trdgethistoryorderlist"
	"gitee.com/shing1211/futuapi4go/pb/trdgetmarginratio"
	"gitee.com/shing1211/futuapi4go/pb/trdgetmaxtrdqtys"
	"gitee.com/shing1211/futuapi4go/pb/trdgetorderfee"
	"gitee.com/shing1211/futuapi4go/pb/trdgetorderfilllist"
	"gitee.com/shing1211/futuapi4go/pb/trdgetorderlist"
	"gitee.com/shing1211/futuapi4go/pb/trdgetpositionlist"
	"gitee.com/shing1211/futuapi4go/pb/trdmodifyorder"
	"gitee.com/shing1211/futuapi4go/pb/trdplaceorder"
	"gitee.com/shing1211/futuapi4go/pb/trdunlocktrade"
)

func (s *Server) RegisterTrdHandlers() {
	s.RegisterHandler(4001, s.handleGetAccList)
	s.RegisterHandler(4002, s.handleUnlockTrade)
	s.RegisterHandler(4003, s.handleGetFunds)
	s.RegisterHandler(4004, s.handleGetOrderFee)
	s.RegisterHandler(4005, s.handleGetMarginRatio)
	s.RegisterHandler(4006, s.handleGetMaxTrdQtys)
	s.RegisterHandler(5001, s.handlePlaceOrder)
	s.RegisterHandler(5002, s.handleModifyOrder)
	s.RegisterHandler(5003, s.handleGetOrderList)
	s.RegisterHandler(5004, s.handleGetHistoryOrderList)
	s.RegisterHandler(5005, s.handleGetOrderFillList)
	s.RegisterHandler(5006, s.handleGetHistoryOrderFillList)
	s.RegisterHandler(6001, s.handleGetPositionList)
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
