package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetheatmapdata"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrisefalldistr"
	"github.com/shing1211/futuapi4go/pkg/util"
)

func GetHeatMapData(ctx context.Context, c *futuapi.Client, req *qotgetheatmapdata.C2S) (*qotgetheatmapdata.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHeatMapData: req is nil")
	}
	var rsp qotgetheatmapdata.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetHeatMapData, &qotgetheatmapdata.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHeatMapData", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetHeatMapData", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetRiseFallDistribution(ctx context.Context, c *futuapi.Client, req *qotgetrisefalldistr.C2S) (*qotgetrisefalldistr.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetRiseFallDistribution: req is nil")
	}
	var rsp qotgetrisefalldistr.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetRiseFallDistribution, &qotgetrisefalldistr.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetRiseFallDistribution", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetRiseFallDistribution", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}
