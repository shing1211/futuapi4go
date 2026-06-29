package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsearchquote"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsearchnews"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindicatorlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequestindicatorcalc"
	"github.com/shing1211/futuapi4go/pkg/util"
)

func GetSearchQuote(ctx context.Context, c *futuapi.Client, req *qotgetsearchquote.C2S) (*qotgetsearchquote.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetSearchQuote: req is nil")
	}
	var rsp qotgetsearchquote.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetSearchQuote, &qotgetsearchquote.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetSearchQuote", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetSearchQuote", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetSearchNews(ctx context.Context, c *futuapi.Client, req *qotgetsearchnews.C2S) (*qotgetsearchnews.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetSearchNews: req is nil")
	}
	var rsp qotgetsearchnews.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetSearchNews, &qotgetsearchnews.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetSearchNews", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetSearchNews", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetIndicatorList(ctx context.Context, c *futuapi.Client, req *qotgetindicatorlist.C2S) (*qotgetindicatorlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndicatorList: req is nil")
	}
	var rsp qotgetindicatorlist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetIndicatorList, &qotgetindicatorlist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetIndicatorList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetIndicatorList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func RequestIndicatorCalc(ctx context.Context, c *futuapi.Client, req *qotrequestindicatorcalc.C2S) (*qotrequestindicatorcalc.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("RequestIndicatorCalc: req is nil")
	}
	var rsp qotrequestindicatorcalc.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_RequestIndicatorCalc, &qotrequestindicatorcalc.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("RequestIndicatorCalc", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("RequestIndicatorCalc", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}
