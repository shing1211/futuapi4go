package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgethighdividendsoerank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgethotlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetperiodchangerank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshortsellingrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgettopmoverrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusafterhoursrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusovernightrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetuspremarketrank"
	"github.com/shing1211/futuapi4go/pkg/util"
)

func GetUSPreMarketRank(ctx context.Context, c *futuapi.Client, req *qotgetuspremarketrank.C2S) (*qotgetuspremarketrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetUSPreMarketRank: req is nil")
	}
	var rsp qotgetuspremarketrank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetUSPreMarketRank, &qotgetuspremarketrank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetUSPreMarketRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetUSPreMarketRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetUSAfterHoursRank(ctx context.Context, c *futuapi.Client, req *qotgetusafterhoursrank.C2S) (*qotgetusafterhoursrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetUSAfterHoursRank: req is nil")
	}
	var rsp qotgetusafterhoursrank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetUSAfterHoursRank, &qotgetusafterhoursrank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetUSAfterHoursRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetUSAfterHoursRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetUSOvernightRank(ctx context.Context, c *futuapi.Client, req *qotgetusovernightrank.C2S) (*qotgetusovernightrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetUSOvernightRank: req is nil")
	}
	var rsp qotgetusovernightrank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetUSOvernightRank, &qotgetusovernightrank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetUSOvernightRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetUSOvernightRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetTopMoversRank(ctx context.Context, c *futuapi.Client, req *qotgettopmoverrank.C2S) (*qotgettopmoverrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetTopMoversRank: req is nil")
	}
	var rsp qotgettopmoverrank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetTopMoversRank, &qotgettopmoverrank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetTopMoversRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetTopMoversRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetHotList(ctx context.Context, c *futuapi.Client, req *qotgethotlist.C2S) (*qotgethotlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHotList: req is nil")
	}
	var rsp qotgethotlist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetHotList, &qotgethotlist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHotList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetHotList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetShortSellingRank(ctx context.Context, c *futuapi.Client, req *qotgetshortsellingrank.C2S) (*qotgetshortsellingrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetShortSellingRank: req is nil")
	}
	var rsp qotgetshortsellingrank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetShortSellingRank, &qotgetshortsellingrank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetShortSellingRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetShortSellingRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetPeriodChangeRank(ctx context.Context, c *futuapi.Client, req *qotgetperiodchangerank.C2S) (*qotgetperiodchangerank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetPeriodChangeRank: req is nil")
	}
	var rsp qotgetperiodchangerank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetPeriodChangeRank, &qotgetperiodchangerank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetPeriodChangeRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetPeriodChangeRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetHighDividendSOERank(ctx context.Context, c *futuapi.Client, req *qotgethighdividendsoerank.C2S) (*qotgethighdividendsoerank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHighDividendSOERank: req is nil")
	}
	var rsp qotgethighdividendsoerank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetHighDividendSOERank, &qotgethighdividendsoerank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHighDividendSOERank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetHighDividendSOERank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}
