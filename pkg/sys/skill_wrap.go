// Package sys provides system APIs for the Futu OpenD SDK.
//
// This package covers connection state, user info, verification, quota,
// delay statistics, and the AI-driven SkillWrap "unusual activity" queries
// (technical / financial / derivative).
//
// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package sys

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/skillwrapapi"
)

// GetTechnicalUnusual queries AI-driven technical indicator unusual activity
// for a stock symbol (code or name) over the last `TimeRange` days.
//
// Available since Futu Protocol v10.8 (re-introduced after removal in v10.6).
// Returns an AI-generated analysis blob in `Content` — language controlled
// by `LanguageId` in the request (see constant.SkillWrapLang_*).
func GetTechnicalUnusual(ctx context.Context, c *futuapi.Client, req *skillwrapapi.TechnicalUnusualReq) (*skillwrapapi.TechnicalUnusualRsp, error) {
	if c == nil {
		return nil, fmt.Errorf("GetTechnicalUnusual: client is nil")
	}
	if req == nil {
		return nil, fmt.Errorf("GetTechnicalUnusual: req is nil")
	}
	if req.GetStockSymbol() == "" {
		return nil, fmt.Errorf("GetTechnicalUnusual: StockSymbol is required")
	}
	var rsp skillwrapapi.TechnicalUnusualRsp
	if err := c.RequestContext(ctx, uint32(constant.ProtoID_SkillWrap_TechnicalUnusual), req, &rsp); err != nil {
		return nil, err
	}
	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetTechnicalUnusual: %s (retType=%d)", rsp.GetRetMsg(), rsp.GetRetType())
	}
	return &rsp, nil
}

// GetFinancialUnusual queries AI-driven financial unusual activity
// for a stock symbol over the last `TimeRange` days.
//
// Available since Futu Protocol v10.8.
func GetFinancialUnusual(ctx context.Context, c *futuapi.Client, req *skillwrapapi.FinancialUnusualReq) (*skillwrapapi.FinancialUnusualRsp, error) {
	if c == nil {
		return nil, fmt.Errorf("GetFinancialUnusual: client is nil")
	}
	if req == nil {
		return nil, fmt.Errorf("GetFinancialUnusual: req is nil")
	}
	if req.GetStockSymbol() == "" {
		return nil, fmt.Errorf("GetFinancialUnusual: StockSymbol is required")
	}
	var rsp skillwrapapi.FinancialUnusualRsp
	if err := c.RequestContext(ctx, uint32(constant.ProtoID_SkillWrap_FinancialUnusual), req, &rsp); err != nil {
		return nil, err
	}
	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetFinancialUnusual: %s (retType=%d)", rsp.GetRetMsg(), rsp.GetRetType())
	}
	return &rsp, nil
}

// GetDerivativeUnusual queries AI-driven derivative unusual activity
// for a stock symbol over the last `TimeRange` days.
//
// Available since Futu Protocol v10.8.
func GetDerivativeUnusual(ctx context.Context, c *futuapi.Client, req *skillwrapapi.DerivativeUnusualReq) (*skillwrapapi.DerivativeUnusualRsp, error) {
	if c == nil {
		return nil, fmt.Errorf("GetDerivativeUnusual: client is nil")
	}
	if req == nil {
		return nil, fmt.Errorf("GetDerivativeUnusual: req is nil")
	}
	if req.GetStockSymbol() == "" {
		return nil, fmt.Errorf("GetDerivativeUnusual: StockSymbol is required")
	}
	var rsp skillwrapapi.DerivativeUnusualRsp
	if err := c.RequestContext(ctx, uint32(constant.ProtoID_SkillWrap_DerivativeUnusual), req, &rsp); err != nil {
		return nil, err
	}
	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, fmt.Errorf("GetDerivativeUnusual: %s (retType=%d)", rsp.GetRetMsg(), rsp.GetRetType())
	}
	return &rsp, nil
}
