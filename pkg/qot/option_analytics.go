package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionevent"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptioneventalert"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionmarketstatistic"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyinghisstatistic"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyinghisvolatility"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyingoverview"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyingrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionzerodtecontract"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionzerodtescreener"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionearningsscreener"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionsellerscreener"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsetoptioneventalert"
	"github.com/shing1211/futuapi4go/pkg/util"
)

func GetOptionMarketStatistic(ctx context.Context, c *futuapi.Client, req *qotgetoptionmarketstatistic.C2S) (*qotgetoptionmarketstatistic.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionMarketStatistic: req is nil")
	}
	var rsp qotgetoptionmarketstatistic.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionMarketStatistic, &qotgetoptionmarketstatistic.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionMarketStatistic", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionMarketStatistic", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionUnderlyingHisStatistic(ctx context.Context, c *futuapi.Client, req *qotgetoptionunderlyinghisstatistic.C2S) (*qotgetoptionunderlyinghisstatistic.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionUnderlyingHisStatistic: req is nil")
	}
	var rsp qotgetoptionunderlyinghisstatistic.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionUnderlyingHisStatistic, &qotgetoptionunderlyinghisstatistic.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionUnderlyingHisStatistic", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionUnderlyingHisStatistic", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionUnderlyingOverview(ctx context.Context, c *futuapi.Client, req *qotgetoptionunderlyingoverview.C2S) (*qotgetoptionunderlyingoverview.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionUnderlyingOverview: req is nil")
	}
	var rsp qotgetoptionunderlyingoverview.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionUnderlyingOverview, &qotgetoptionunderlyingoverview.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionUnderlyingOverview", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionUnderlyingOverview", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionUnderlyingHisVolatility(ctx context.Context, c *futuapi.Client, req *qotgetoptionunderlyinghisvolatility.C2S) (*qotgetoptionunderlyinghisvolatility.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionUnderlyingHisVolatility: req is nil")
	}
	var rsp qotgetoptionunderlyinghisvolatility.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionUnderlyingHisVolatility, &qotgetoptionunderlyinghisvolatility.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionUnderlyingHisVolatility", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionUnderlyingHisVolatility", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionUnderlyingRank(ctx context.Context, c *futuapi.Client, req *qotgetoptionunderlyingrank.C2S) (*qotgetoptionunderlyingrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionUnderlyingRank: req is nil")
	}
	var rsp qotgetoptionunderlyingrank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionUnderlyingRank, &qotgetoptionunderlyingrank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionUnderlyingRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionUnderlyingRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionRank(ctx context.Context, c *futuapi.Client, req *qotgetoptionrank.C2S) (*qotgetoptionrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionRank: req is nil")
	}
	var rsp qotgetoptionrank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionRank, &qotgetoptionrank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionEvent(ctx context.Context, c *futuapi.Client, req *qotgetoptionevent.C2S) (*qotgetoptionevent.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionEvent: req is nil")
	}
	var rsp qotgetoptionevent.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionEvent, &qotgetoptionevent.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionEvent", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionEvent", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionEventAlert(ctx context.Context, c *futuapi.Client, req *qotgetoptioneventalert.C2S) (*qotgetoptioneventalert.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionEventAlert: req is nil")
	}
	var rsp qotgetoptioneventalert.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionEventAlert, &qotgetoptioneventalert.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionEventAlert", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionEventAlert", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func SetOptionEventAlert(ctx context.Context, c *futuapi.Client, req *qotsetoptioneventalert.C2S) (*qotsetoptioneventalert.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("SetOptionEventAlert: req is nil")
	}
	var rsp qotsetoptioneventalert.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_SetOptionEventAlert, &qotsetoptioneventalert.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("SetOptionEventAlert", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("SetOptionEventAlert", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionZeroDteScreener(ctx context.Context, c *futuapi.Client, req *qotgetoptionzerodtescreener.C2S) (*qotgetoptionzerodtescreener.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionZeroDteScreener: req is nil")
	}
	var rsp qotgetoptionzerodtescreener.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionZeroDteScreener, &qotgetoptionzerodtescreener.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionZeroDteScreener", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionZeroDteScreener", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionZeroDteContract(ctx context.Context, c *futuapi.Client, req *qotgetoptionzerodtecontract.C2S) (*qotgetoptionzerodtecontract.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionZeroDteContract: req is nil")
	}
	var rsp qotgetoptionzerodtecontract.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionZeroDteContract, &qotgetoptionzerodtecontract.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionZeroDteContract", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionZeroDteContract", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionEarningsScreener(ctx context.Context, c *futuapi.Client, req *qotgetoptionearningsscreener.C2S) (*qotgetoptionearningsscreener.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionEarningsScreener: req is nil")
	}
	var rsp qotgetoptionearningsscreener.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionEarningsScreener, &qotgetoptionearningsscreener.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionEarningsScreener", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionEarningsScreener", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetOptionSellerScreener(ctx context.Context, c *futuapi.Client, req *qotgetoptionsellerscreener.C2S) (*qotgetoptionsellerscreener.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionSellerScreener: req is nil")
	}
	var rsp qotgetoptionsellerscreener.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetOptionSellerScreener, &qotgetoptionsellerscreener.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionSellerScreener", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetOptionSellerScreener", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}
