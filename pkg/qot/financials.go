// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialsstatements"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialrevenuebreakdown"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetFinancialsStatements       = 3227
	ProtoID_GetFinancialsRevenueBreakdown = 3228
)

type GetFinancialsStatementsRequest struct {
	Security       *qotcommon.Security
	StatementType  int32
	FinancialType  int32
	CurrencyCode   string
	NextKey        string
	Num            int32
}

type GetFinancialsStatementsResponse struct {
	StructureList []*qotgetfinancialsstatements.FinancialFieldInfo
	ReportList    []*qotgetfinancialsstatements.FinancialReport
	NextKey       string
}

func GetFinancialsStatements(ctx context.Context, c *futuapi.Client, req *GetFinancialsStatementsRequest) (*GetFinancialsStatementsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFinancialsStatements: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetFinancialsStatements: Security is nil")
	}
	c2s := &qotgetfinancialsstatements.C2S{
		Security: req.Security,
	}
	if req.StatementType != 0 {
		v := qotcommon.FinancialStatementsType(req.StatementType)
		c2s.StatementType = &v
	}
	if req.FinancialType != 0 {
		v := qotcommon.F10Type(req.FinancialType)
		c2s.FinancialType = &v
	}
	if req.CurrencyCode != "" {
		c2s.CurrencyCode = &req.CurrencyCode
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num != 0 {
		c2s.Num = &req.Num
	}
	pkt := &qotgetfinancialsstatements.Request{C2S: c2s}
	var rsp qotgetfinancialsstatements.Response

	if err := c.RequestContext(ctx, ProtoID_GetFinancialsStatements, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFinancialsStatements", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetFinancialsStatements", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetFinancialsStatementsResponse{
		StructureList: s2c.StructureList,
		ReportList:    s2c.ReportList,
		NextKey:      util.ProtoStr(s2c.NextKey),
	}, nil
}

type GetFinancialsRevenueBreakdownRequest struct {
	Security      *qotcommon.Security
	Date          uint32
	FinancialType int32
	CurrencyCode  string
}

type GetFinancialsRevenueBreakdownResponse struct {
	Period         string
	BreakdownList  []*qotgetfinancialrevenuebreakdown.RevenueBreakdownGroup
	CurrencyCode   string
	ScreenDateList []*qotgetfinancialrevenuebreakdown.ScreenDate
}

func GetFinancialsRevenueBreakdown(ctx context.Context, c *futuapi.Client, req *GetFinancialsRevenueBreakdownRequest) (*GetFinancialsRevenueBreakdownResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFinancialsRevenueBreakdown: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetFinancialsRevenueBreakdown: Security is nil")
	}
	c2s := &qotgetfinancialrevenuebreakdown.C2S{
		Security: req.Security,
	}
	if req.Date != 0 {
		c2s.Date = &req.Date
	}
	if req.FinancialType != 0 {
		v := qotcommon.F10Type(req.FinancialType)
		c2s.FinancialType = &v
	}
	if req.CurrencyCode != "" {
		c2s.CurrencyCode = &req.CurrencyCode
	}
	pkt := &qotgetfinancialrevenuebreakdown.Request{C2S: c2s}
	var rsp qotgetfinancialrevenuebreakdown.Response

	if err := c.RequestContext(ctx, ProtoID_GetFinancialsRevenueBreakdown, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFinancialsRevenueBreakdown", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetFinancialsRevenueBreakdown", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetFinancialsRevenueBreakdownResponse{
		Period:        util.ProtoStr(s2c.Period),
		BreakdownList: s2c.BreakdownList,
		CurrencyCode:  util.ProtoStr(s2c.CurrencyCode),
		ScreenDateList: s2c.ScreenDateList,
	}, nil
}