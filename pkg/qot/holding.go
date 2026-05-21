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
	"github.com/shing1211/futuapi4go/pkg/util"
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

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHoldingChangeList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("GetHoldingChangeList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &GetHoldingChangeListResponse{
		Security:          s2c.Security,
		HoldingChangeList: s2c.HoldingChangeList,
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

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("RequestRehab", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("RequestRehab", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &RequestRehabResponse{
		RehabList: s2c.RehabList,
	}, nil
}

// GetRehabRequest defines parameters for GetRehab.
// Deprecated: Removed in Futu v10.6 proto — proto package qotgetrehab no longer exists.
// Use RequestRehab instead.
type GetRehabRequest struct {
	Security *qotcommon.Security
}

// GetRehabResponse is the response type for GetRehab.
// Deprecated: Removed in Futu v10.6 proto.
type GetRehabResponse struct {
	SecurityRehabList any //nolint:revive // deprecated type
}

// GetRehab returns rehabilitation (复权) data for the given security.
// Deprecated: Removed in Futu v10.6 proto — proto package qotgetrehab no longer exists.
// Use RequestRehab instead.
func GetRehab(ctx context.Context, c *futuapi.Client, req *GetRehabRequest) (*GetRehabResponse, error) {
	_ = req
	_ = c
	return nil, fmt.Errorf("GetRehab: removed in Futu v10.6 — use RequestRehab instead")
}
