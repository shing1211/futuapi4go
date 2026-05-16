// Package qot provides market data APIs for the Futu OpenD SDK.
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

package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetholdingchangelist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequestrehab"
)

// GetHoldingChangeListRequest defines parameters for GetHoldingChangeList.
type GetHoldingChangeListRequest struct {
	Security       *qotcommon.Security
	HolderCategory int32
	BeginTime      string
	EndTime        string
}

// GetHoldingChangeListResponse is the response type for GetHoldingChangeList.
type GetHoldingChangeListResponse struct {
	Security          *qotcommon.Security
	HoldingChangeList []*qotcommon.ShareHoldingChange
}

// GetHoldingChangeList returns the holding change list for the given security.
func GetHoldingChangeList(ctx context.Context, c *futuapi.Client, req *GetHoldingChangeListRequest) (*GetHoldingChangeListResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHoldingChangeList: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("security is required")
	}

	c2s := &qotgetholdingchangelist.C2S{
		Security:       req.Security,
		HolderCategory: &req.HolderCategory,
		BeginTime:      &req.BeginTime,
		EndTime:        &req.EndTime,
	}

	pkt := &qotgetholdingchangelist.Request{C2S: c2s}
	var rsp qotgetholdingchangelist.Response

	if err := c.RequestContext(ctx, ProtoID_GetHoldingChangeList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHoldingChangeList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetHoldingChangeList: s2c is nil")
	}

	return &GetHoldingChangeListResponse{
		Security:          s2c.GetSecurity(),
		HoldingChangeList: s2c.GetHoldingChangeList(),
	}, nil
}

// RequestRehabRequest defines parameters for RequestRehab.
type RequestRehabRequest struct {
	Security *qotcommon.Security
}

// RequestRehabResponse is the response type for RequestRehab.
type RequestRehabResponse struct {
	RehabList []*qotcommon.Rehab
}

// RequestRehab requests rehabilitation (复权) data for the given security.
func RequestRehab(ctx context.Context, c *futuapi.Client, req *RequestRehabRequest) (*RequestRehabResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("RequestRehab: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("security is required")
	}

	c2s := &qotrequestrehab.C2S{
		Security: req.Security,
	}

	pkt := &qotrequestrehab.Request{C2S: c2s}
	var rsp qotrequestrehab.Response

	if err := c.RequestContext(ctx, ProtoID_RequestRehab, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("RequestRehab", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("RequestRehab: s2c is nil")
	}

	return &RequestRehabResponse{
		RehabList: s2c.GetRehabList(),
	}, nil
}

// GetRehabRequest defines parameters for GetRehab.
type GetRehabRequest struct {
	Security *qotcommon.Security
}

// GetRehabResponse is the response type for GetRehab.
type GetRehabResponse struct {
	RehabList []*qotcommon.Rehab
}

// GetRehab returns rehabilitation (复权) data for the given security.
func GetRehab(ctx context.Context, c *futuapi.Client, req *GetRehabRequest) (*GetRehabResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetRehab: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("security is required")
	}

	c2s := &qotrequestrehab.C2S{
		Security: req.Security,
	}

	pkt := &qotrequestrehab.Request{C2S: c2s}
	var rsp qotrequestrehab.Response

	if err := c.RequestContext(ctx, ProtoID_RequestRehab, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetRehab", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetRehab: s2c is nil")
	}

	return &GetRehabResponse{
		RehabList: s2c.GetRehabList(),
	}, nil
}
