// Package qot provides market data APIs for the Futu OpenD SDK.
//
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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcompanyprofile"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcompanyexecutives"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcompanyexecutivebackground"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcompanyoperationalefficiency"
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_GetCompanyProfile             = 3243
	ProtoID_GetCompanyExecutives          = 3244
	ProtoID_GetCompanyExecutiveBackground = 3245
	ProtoID_GetCompanyOperationalEfficiency = 3246
)

type GetCompanyProfileRequest struct {
	Security *qotcommon.Security
}

type GetCompanyProfileResponse struct {
	ItemList []*qotgetcompanyprofile.CompanyLabItem
}

func GetCompanyProfile(ctx context.Context, c *futuapi.Client, req *GetCompanyProfileRequest) (*GetCompanyProfileResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetCompanyProfile: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetCompanyProfile: Security is nil")
	}
	c2s := &qotgetcompanyprofile.C2S{
		Security: req.Security,
	}
	pkt := &qotgetcompanyprofile.Request{C2S: c2s}
	var rsp qotgetcompanyprofile.Response

	if err := c.RequestContext(ctx, ProtoID_GetCompanyProfile, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCompanyProfile", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetCompanyProfile", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetCompanyProfileResponse{
		ItemList: s2c.ItemList,
	}, nil
}

type GetCompanyExecutivesRequest struct {
	Security *qotcommon.Security
}

type GetCompanyExecutivesResponse struct {
	DirectorList []*qotgetcompanyexecutives.DirectorInfo
}

func GetCompanyExecutives(ctx context.Context, c *futuapi.Client, req *GetCompanyExecutivesRequest) (*GetCompanyExecutivesResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetCompanyExecutives: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetCompanyExecutives: Security is nil")
	}
	c2s := &qotgetcompanyexecutives.C2S{
		Security: req.Security,
	}
	pkt := &qotgetcompanyexecutives.Request{C2S: c2s}
	var rsp qotgetcompanyexecutives.Response

	if err := c.RequestContext(ctx, ProtoID_GetCompanyExecutives, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCompanyExecutives", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetCompanyExecutives", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetCompanyExecutivesResponse{
		DirectorList: s2c.DirectorList,
	}, nil
}

type GetCompanyExecutiveBackgroundRequest struct {
	Security   *qotcommon.Security
	LeaderName string
}

type GetCompanyExecutiveBackgroundResponse struct {
	BriefBackground string
}

func GetCompanyExecutiveBackground(ctx context.Context, c *futuapi.Client, req *GetCompanyExecutiveBackgroundRequest) (*GetCompanyExecutiveBackgroundResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetCompanyExecutiveBackground: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetCompanyExecutiveBackground: Security is nil")
	}
	c2s := &qotgetcompanyexecutivebackground.C2S{
		Security: req.Security,
	}
	if req.LeaderName != "" {
		c2s.LeaderName = &req.LeaderName
	}
	pkt := &qotgetcompanyexecutivebackground.Request{C2S: c2s}
	var rsp qotgetcompanyexecutivebackground.Response

	if err := c.RequestContext(ctx, ProtoID_GetCompanyExecutiveBackground, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCompanyExecutiveBackground", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetCompanyExecutiveBackground", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetCompanyExecutiveBackgroundResponse{
		BriefBackground: util.ProtoStr(s2c.BriefBackground),
	}, nil
}

type GetCompanyOperationalEfficiencyRequest struct {
	Security      *qotcommon.Security
	NextKey       string
	Num           int32
	CurrencyCode  string
	FinancialType int32
}

type GetCompanyOperationalEfficiencyResponse struct {
	ItemList     []*qotgetcompanyoperationalefficiency.OperationalEfficiencyItem
	NextKey      string
	CurrencyCode string
}

func GetCompanyOperationalEfficiency(ctx context.Context, c *futuapi.Client, req *GetCompanyOperationalEfficiencyRequest) (*GetCompanyOperationalEfficiencyResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetCompanyOperationalEfficiency: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetCompanyOperationalEfficiency: Security is nil")
	}
	c2s := &qotgetcompanyoperationalefficiency.C2S{
		Security: req.Security,
	}
	if req.NextKey != "" {
		c2s.NextKey = &req.NextKey
	}
	if req.Num != 0 {
		c2s.Num = &req.Num
	}
	if req.CurrencyCode != "" {
		c2s.CurrencyCode = &req.CurrencyCode
	}
	if req.FinancialType != 0 {
		ft := qotcommon.F10Type(req.FinancialType)
		c2s.FinancialType = &ft
	}
	pkt := &qotgetcompanyoperationalefficiency.Request{C2S: c2s}
	var rsp qotgetcompanyoperationalefficiency.Response

	if err := c.RequestContext(ctx, ProtoID_GetCompanyOperationalEfficiency, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetCompanyOperationalEfficiency", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetCompanyOperationalEfficiency", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetCompanyOperationalEfficiencyResponse{
		ItemList:     s2c.ItemList,
		NextKey:     util.ProtoStr(s2c.NextKey),
		CurrencyCode: util.ProtoStr(s2c.CurrencyCode),
	}, nil
}