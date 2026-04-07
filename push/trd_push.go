package push

import (
	"google.golang.org/protobuf/proto"

	"gitee.com/shing1211/futuapi4go/pb/trdcommon"
	"gitee.com/shing1211/futuapi4go/pb/trdupdateorder"
	"gitee.com/shing1211/futuapi4go/pb/trdupdateorderfill"
)

const (
	ProtoID_Trd_UpdateOrder     = 7001
	ProtoID_Trd_UpdateOrderFill = 7002
	ProtoID_Trd_Notify          = 7003
	ProtoID_Trd_ReconfirmOrder  = 7004
	ProtoID_Trd_SubAccPush      = 7005
)

type UpdateOrder struct {
	Header *trdcommon.TrdHeader
	Order  *trdcommon.Order
}

func ParseUpdateOrder(body []byte) (*UpdateOrder, error) {
	var rsp trdupdateorder.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdateOrder{
		Header: rsp.GetHeader(),
		Order:  rsp.GetOrder(),
	}, nil
}

type UpdateOrderFill struct {
	Header    *trdcommon.TrdHeader
	OrderFill *trdcommon.OrderFill
}

func ParseUpdateOrderFill(body []byte) (*UpdateOrderFill, error) {
	var rsp trdupdateorderfill.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdateOrderFill{
		Header:    rsp.GetHeader(),
		OrderFill: rsp.GetOrderFill(),
	}, nil
}
