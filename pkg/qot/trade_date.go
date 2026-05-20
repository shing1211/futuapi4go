package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgettradedate"
	"github.com/shing1211/futuapi4go/pkg/util"
)

type GetTradeDateRequest struct {
	Market    int32
	BeginTime string
	EndTime   string
}

type TradeDateInfo struct {
	Time          string
	Timestamp     float64
	TradeDateType int32
}

type GetTradeDateResponse struct {
	TradeDateList []*TradeDateInfo
}

func GetTradeDate(ctx context.Context, c *futuapi.Client, req *GetTradeDateRequest) (*GetTradeDateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetTradeDate: request is nil")
	}

	c2s := &qotgettradedate.C2S{
		Market:    &req.Market,
		BeginTime: &req.BeginTime,
		EndTime:   &req.EndTime,
	}

	var rsp qotgettradedate.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetTradeDate, &qotgettradedate.Request{C2S: c2s}, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetTradeDate", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetTradeDate", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetTradeDateResponse{
		TradeDateList: make([]*TradeDateInfo, 0, len(s2c.TradeDateList)),
	}
	for _, td := range s2c.TradeDateList {
		if td == nil {
			continue
		}
		result.TradeDateList = append(result.TradeDateList, &TradeDateInfo{
			Time:          util.ProtoStr(td.Time),
			Timestamp:     util.ProtoFloat64(td.Timestamp),
			TradeDateType: td.GetTradeDateType(),
		})
	}
	return result, nil
}
