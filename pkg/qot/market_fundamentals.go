package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetdividendcalendar"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetdividendrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetearningsbeatrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetearningscalendar"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteconomiccalendar"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfedwatchdotplot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfedwatchtargetrate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetmacroindicatorhistory"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetmacroindicatorlist"
	"github.com/shing1211/futuapi4go/pkg/util"
)

func GetEarningsCalendar(ctx context.Context, c *futuapi.Client, req *qotgetearningscalendar.C2S) (*qotgetearningscalendar.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEarningsCalendar: req is nil")
	}
	var rsp qotgetearningscalendar.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEarningsCalendar, &qotgetearningscalendar.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEarningsCalendar", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEarningsCalendar", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetMacroIndicatorList(ctx context.Context, c *futuapi.Client, req *qotgetmacroindicatorlist.C2S) (*qotgetmacroindicatorlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetMacroIndicatorList: req is nil")
	}
	var rsp qotgetmacroindicatorlist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetMacroIndicatorList, &qotgetmacroindicatorlist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetMacroIndicatorList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetMacroIndicatorList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetMacroIndicatorHistory(ctx context.Context, c *futuapi.Client, req *qotgetmacroindicatorhistory.C2S) (*qotgetmacroindicatorhistory.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetMacroIndicatorHistory: req is nil")
	}
	var rsp qotgetmacroindicatorhistory.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetMacroIndicatorHistory, &qotgetmacroindicatorhistory.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetMacroIndicatorHistory", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetMacroIndicatorHistory", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetFedWatchTargetRate(ctx context.Context, c *futuapi.Client, req *qotgetfedwatchtargetrate.C2S) (*qotgetfedwatchtargetrate.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFedWatchTargetRate: req is nil")
	}
	var rsp qotgetfedwatchtargetrate.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetFedWatchTargetRate, &qotgetfedwatchtargetrate.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFedWatchTargetRate", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetFedWatchTargetRate", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetFedWatchDotPlot(ctx context.Context, c *futuapi.Client, req *qotgetfedwatchdotplot.C2S) (*qotgetfedwatchdotplot.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFedWatchDotPlot: req is nil")
	}
	var rsp qotgetfedwatchdotplot.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetFedWatchDotPlot, &qotgetfedwatchdotplot.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFedWatchDotPlot", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetFedWatchDotPlot", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetEarningsBeatRank(ctx context.Context, c *futuapi.Client, req *qotgetearningsbeatrank.C2S) (*qotgetearningsbeatrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEarningsBeatRank: req is nil")
	}
	var rsp qotgetearningsbeatrank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEarningsBeatRank, &qotgetearningsbeatrank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEarningsBeatRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEarningsBeatRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetDividendRank(ctx context.Context, c *futuapi.Client, req *qotgetdividendrank.C2S) (*qotgetdividendrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetDividendRank: req is nil")
	}
	var rsp qotgetdividendrank.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetDividendRank, &qotgetdividendrank.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetDividendRank", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetDividendRank", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetDividendCalendar(ctx context.Context, c *futuapi.Client, req *qotgetdividendcalendar.C2S) (*qotgetdividendcalendar.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetDividendCalendar: req is nil")
	}
	var rsp qotgetdividendcalendar.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetDividendCalendar, &qotgetdividendcalendar.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetDividendCalendar", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetDividendCalendar", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetEconomicCalendar(ctx context.Context, c *futuapi.Client, req *qotgeteconomiccalendar.C2S) (*qotgeteconomiccalendar.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEconomicCalendar: req is nil")
	}
	var rsp qotgeteconomiccalendar.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEconomicCalendar, &qotgeteconomiccalendar.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEconomicCalendar", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEconomicCalendar", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}
