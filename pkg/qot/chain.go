package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialchainbyplate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialchaindetail"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialchainlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialplateinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialplatestock"
	"github.com/shing1211/futuapi4go/pkg/util"
)

func GetIndustrialChainList(ctx context.Context, c *futuapi.Client, req *qotgetindustrialchainlist.C2S) (*qotgetindustrialchainlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialChainList: req is nil")
	}
	var rsp qotgetindustrialchainlist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetIndustrialChainList, &qotgetindustrialchainlist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetIndustrialChainList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetIndustrialChainList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetIndustrialChainDetail(ctx context.Context, c *futuapi.Client, req *qotgetindustrialchaindetail.C2S) (*qotgetindustrialchaindetail.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialChainDetail: req is nil")
	}
	var rsp qotgetindustrialchaindetail.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetIndustrialChainDetail, &qotgetindustrialchaindetail.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetIndustrialChainDetail", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetIndustrialChainDetail", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetIndustrialChainByPlate(ctx context.Context, c *futuapi.Client, req *qotgetindustrialchainbyplate.C2S) (*qotgetindustrialchainbyplate.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialChainByPlate: req is nil")
	}
	var rsp qotgetindustrialchainbyplate.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetIndustrialChainByPlate, &qotgetindustrialchainbyplate.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetIndustrialChainByPlate", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetIndustrialChainByPlate", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetIndustrialPlateInfo(ctx context.Context, c *futuapi.Client, req *qotgetindustrialplateinfo.C2S) (*qotgetindustrialplateinfo.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialPlateInfo: req is nil")
	}
	var rsp qotgetindustrialplateinfo.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetIndustrialPlateInfo, &qotgetindustrialplateinfo.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetIndustrialPlateInfo", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetIndustrialPlateInfo", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetIndustrialPlateStock(ctx context.Context, c *futuapi.Client, req *qotgetindustrialplatestock.C2S) (*qotgetindustrialplatestock.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialPlateStock: req is nil")
	}
	var rsp qotgetindustrialplatestock.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetIndustrialPlateStock, &qotgetindustrialplatestock.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetIndustrialPlateStock", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetIndustrialPlateStock", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}
