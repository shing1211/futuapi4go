package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgettradedate"
)

const (
	ProtoID_Qot_GetTradeDate = 3225
)

func GetTradeDate(ctx context.Context, c *futuapi.Client, req *qotgettradedate.C2S) (*qotgettradedate.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetTradeDate: request is nil")
	}

	var rsp qotgettradedate.Response
	if err := c.RequestContext(ctx, ProtoID_Qot_GetTradeDate, &qotgettradedate.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetTradeDate", rsp.GetRetType(), rsp.GetRetMsg())
	}

	if rsp.GetS2C() == nil {
		return nil, wrapError("GetTradeDate", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return rsp.GetS2C(), nil
}
