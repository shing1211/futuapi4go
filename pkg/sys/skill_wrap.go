package sys

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/skillwrapapi"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_SkillWrapAPI = 8001
)

// GetTechnicalUnusual queries technical unusual stocks.
func GetTechnicalUnusual(ctx context.Context, c *futuapi.Client, req *skillwrapapi.TechnicalUnusualReq) (*skillwrapapi.TechnicalUnusualRsp, error) {
	if req == nil {
		return nil, fmt.Errorf("GetTechnicalUnusual: request is nil")
	}
	if req.StockSymbol == nil || *req.StockSymbol == "" {
		return nil, fmt.Errorf("GetTechnicalUnusual: stock_symbol is required")
	}

	var rsp skillwrapapi.TechnicalUnusualRsp
	if err := c.RequestContext(ctx, ProtoID_SkillWrapAPI, req, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetTechnicalUnusual", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	return &rsp, nil
}

// GetFinancialUnusual queries financial unusual stocks.
func GetFinancialUnusual(ctx context.Context, c *futuapi.Client, req *skillwrapapi.FinancialUnusualReq) (*skillwrapapi.FinancialUnusualRsp, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFinancialUnusual: request is nil")
	}
	if req.StockSymbol == nil || *req.StockSymbol == "" {
		return nil, fmt.Errorf("GetFinancialUnusual: stock_symbol is required")
	}

	var rsp skillwrapapi.FinancialUnusualRsp
	if err := c.RequestContext(ctx, ProtoID_SkillWrapAPI, req, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFinancialUnusual", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	return &rsp, nil
}

// GetDerivativeUnusual queries derivative unusual stocks.
func GetDerivativeUnusual(ctx context.Context, c *futuapi.Client, req *skillwrapapi.DerivativeUnusualReq) (*skillwrapapi.DerivativeUnusualRsp, error) {
	if req == nil {
		return nil, fmt.Errorf("GetDerivativeUnusual: request is nil")
	}
	if req.StockSymbol == nil || *req.StockSymbol == "" {
		return nil, fmt.Errorf("GetDerivativeUnusual: stock_symbol is required")
	}

	var rsp skillwrapapi.DerivativeUnusualRsp
	if err := c.RequestContext(ctx, ProtoID_SkillWrapAPI, req, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetDerivativeUnusual", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	return &rsp, nil
}
