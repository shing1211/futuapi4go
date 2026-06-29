package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetarkactivetransaction"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetarkfundholding"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetarkstockdynamic"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutiondistribution"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionholdingchange"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionholdinglist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionprofile"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetratingchange"
	"github.com/shing1211/futuapi4go/pkg/util"
)

func GetInstitutionList(ctx context.Context, c *futuapi.Client, req *qotgetinstitutionlist.C2S) (*qotgetinstitutionlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionList: req is nil")
	}
	var rsp qotgetinstitutionlist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetInstitutionList, &qotgetinstitutionlist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetInstitutionList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetInstitutionList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetInstitutionProfile(ctx context.Context, c *futuapi.Client, req *qotgetinstitutionprofile.C2S) (*qotgetinstitutionprofile.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionProfile: req is nil")
	}
	var rsp qotgetinstitutionprofile.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetInstitutionProfile, &qotgetinstitutionprofile.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetInstitutionProfile", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetInstitutionProfile", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetInstitutionDistribution(ctx context.Context, c *futuapi.Client, req *qotgetinstitutiondistribution.C2S) (*qotgetinstitutiondistribution.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionDistribution: req is nil")
	}
	var rsp qotgetinstitutiondistribution.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetInstitutionDistribution, &qotgetinstitutiondistribution.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetInstitutionDistribution", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetInstitutionDistribution", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetInstitutionHoldingChange(ctx context.Context, c *futuapi.Client, req *qotgetinstitutionholdingchange.C2S) (*qotgetinstitutionholdingchange.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionHoldingChange: req is nil")
	}
	var rsp qotgetinstitutionholdingchange.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetInstitutionHoldingChange, &qotgetinstitutionholdingchange.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetInstitutionHoldingChange", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetInstitutionHoldingChange", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetInstitutionHoldingList(ctx context.Context, c *futuapi.Client, req *qotgetinstitutionholdinglist.C2S) (*qotgetinstitutionholdinglist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionHoldingList: req is nil")
	}
	var rsp qotgetinstitutionholdinglist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetInstitutionHoldingList, &qotgetinstitutionholdinglist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetInstitutionHoldingList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetInstitutionHoldingList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetArkFundHolding(ctx context.Context, c *futuapi.Client, req *qotgetarkfundholding.C2S) (*qotgetarkfundholding.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetArkFundHolding: req is nil")
	}
	var rsp qotgetarkfundholding.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetArkFundHolding, &qotgetarkfundholding.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetArkFundHolding", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetArkFundHolding", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetArkStockDynamic(ctx context.Context, c *futuapi.Client, req *qotgetarkstockdynamic.C2S) (*qotgetarkstockdynamic.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetArkStockDynamic: req is nil")
	}
	var rsp qotgetarkstockdynamic.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetArkStockDynamic, &qotgetarkstockdynamic.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetArkStockDynamic", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetArkStockDynamic", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetArkActiveTransaction(ctx context.Context, c *futuapi.Client, req *qotgetarkactivetransaction.C2S) (*qotgetarkactivetransaction.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetArkActiveTransaction: req is nil")
	}
	var rsp qotgetarkactivetransaction.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetArkActiveTransaction, &qotgetarkactivetransaction.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetArkActiveTransaction", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetArkActiveTransaction", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

func GetRatingChange(ctx context.Context, c *futuapi.Client, req *qotgetratingchange.C2S) (*qotgetratingchange.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetRatingChange: req is nil")
	}
	var rsp qotgetratingchange.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetRatingChange, &qotgetratingchange.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetRatingChange", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetRatingChange", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}
