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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetownerplate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetplatesecurity"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetplateset"
)

const (
	ProtoID_GetOwnerPlate = 3207
)

// Plate represents a market plate (板块).
type Plate struct {
	Plate *qotcommon.Security
	Name  string
}

// GetPlateSetRequest defines parameters for GetPlateSet.
type GetPlateSetRequest struct {
	Market       int32
	PlateSetType int32
}

// GetPlateSetResponse is the response type for GetPlateSet.
type GetPlateSetResponse struct {
	PlateSetList []*Plate
}

// GetPlateSet returns the set of plates for the given market.
func GetPlateSet(ctx context.Context, c *futuapi.Client, req *GetPlateSetRequest) (*GetPlateSetResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetPlateSet: request is nil")
	}
	if req.Market == 0 {
		return nil, fmt.Errorf("invalid market: must be non-zero")
	}

	c2s := &qotgetplateset.C2S{
		Market:       &req.Market,
		PlateSetType: &req.PlateSetType,
	}

	pkt := &qotgetplateset.Request{C2S: c2s}
	var rsp qotgetplateset.Response

	if err := c.RequestContext(ctx, ProtoID_GetPlateSet, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetPlateSet", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetPlateSet: s2c is nil")
	}

	result := &GetPlateSetResponse{
		PlateSetList: make([]*Plate, 0, len(s2c.GetPlateInfoList())),
	}

	for _, p := range s2c.GetPlateInfoList() {
		if p == nil {
			continue
		}
		result.PlateSetList = append(result.PlateSetList, &Plate{
			Plate: p.GetPlate(),
			Name:  p.GetName(),
		})
	}

	return result, nil
}

// GetPlateSecurityRequest defines parameters for GetPlateSecurity.
type GetPlateSecurityRequest struct {
	Plate *qotcommon.Security
}

// GetPlateSecurityResponse is the response type for GetPlateSecurity.
type GetPlateSecurityResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

// GetPlateSecurity returns securities belonging to the given plate.
func GetPlateSecurity(ctx context.Context, c *futuapi.Client, req *GetPlateSecurityRequest) (*GetPlateSecurityResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetPlateSecurity: request is nil")
	}
	if req.Plate == nil {
		return nil, fmt.Errorf("plate is required")
	}

	c2s := &qotgetplatesecurity.C2S{
		Plate: req.Plate,
	}

	pkt := &qotgetplatesecurity.Request{C2S: c2s}
	var rsp qotgetplatesecurity.Response

	if err := c.RequestContext(ctx, ProtoID_GetPlateSecurity, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetPlateSecurity", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetPlateSecurity: s2c is nil")
	}

	return &GetPlateSecurityResponse{
		StaticInfoList: s2c.GetStaticInfoList(),
	}, nil
}

// GetOwnerPlateRequest defines parameters for GetOwnerPlate.
type GetOwnerPlateRequest struct {
	SecurityList []*qotcommon.Security
}

// GetOwnerPlateResponse is the response type for GetOwnerPlate.
type GetOwnerPlateResponse struct {
	OwnerPlateList []*qotgetownerplate.SecurityOwnerPlate
}

// GetOwnerPlate returns the owner plates for the given securities.
func GetOwnerPlate(ctx context.Context, c *futuapi.Client, req *GetOwnerPlateRequest) (*GetOwnerPlateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOwnerPlate: request is nil")
	}
	if len(req.SecurityList) == 0 {
		return nil, fmt.Errorf("security list is empty")
	}

	c2s := &qotgetownerplate.C2S{
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetownerplate.Request{C2S: c2s}
	var rsp qotgetownerplate.Response

	if err := c.RequestContext(ctx, ProtoID_GetOwnerPlate, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOwnerPlate", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetOwnerPlate: s2c is nil")
	}

	return &GetOwnerPlateResponse{
		OwnerPlateList: s2c.GetOwnerPlateList(),
	}, nil
}
