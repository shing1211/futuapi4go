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
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistorykl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistoryklquota"
	"github.com/shing1211/futuapi4go/pkg/util"
)

// RequestHistoryKLRequest defines parameters for RequestHistoryKL.
type RequestHistoryKLRequest struct {
	RehabType        int32
	KlType           int32
	Security         *qotcommon.Security
	BeginTime        string
	EndTime          string
	MaxAckKLNum      int32
	NeedKLFieldsFlag int64
	NextReqKey       []byte
	ExtendedTime     bool
	Session          int32
}

// RequestHistoryKLResponse is the response type for RequestHistoryKL.
type RequestHistoryKLResponse struct {
	Security   *qotcommon.Security
	Name       string
	KLList     []*KLine
	NextReqKey []byte
}

// RequestHistoryKL requests historical K-line (candlestick) data for the given security.
func RequestHistoryKL(ctx context.Context, c *futuapi.Client, req *RequestHistoryKLRequest) (*RequestHistoryKLResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("RequestHistoryKL: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("security is required")
	}

	c2s := &qotrequesthistorykl.C2S{
		RehabType: &req.RehabType,
		KlType:    &req.KlType,
		Security:  req.Security,
		BeginTime: &req.BeginTime,
		EndTime:   &req.EndTime,
	}
	if req.MaxAckKLNum != 0 {
		c2s.MaxAckKLNum = &req.MaxAckKLNum
	}
	if req.NeedKLFieldsFlag != 0 {
		c2s.NeedKLFieldsFlag = &req.NeedKLFieldsFlag
	}
	if len(req.NextReqKey) > 0 {
		c2s.NextReqKey = req.NextReqKey
	}
	if req.ExtendedTime {
		c2s.ExtendedTime = &req.ExtendedTime
	}
	if req.Session != 0 {
		c2s.Session = &req.Session
	}

	pkt := &qotrequesthistorykl.Request{C2S: c2s}
	var rsp qotrequesthistorykl.Response

	if err := c.RequestContext(ctx, ProtoID_RequestHistoryKL, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("RequestHistoryKL", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("RequestHistoryKL", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &RequestHistoryKLResponse{
		Security:   s2c.Security,
		Name:       util.ProtoStr(s2c.Name),
		NextReqKey: s2c.NextReqKey,
		KLList:     make([]*KLine, 0, len(s2c.KlList)),
	}

	for _, kl := range s2c.KlList {
		if kl == nil {
			continue
		}
		result.KLList = append(result.KLList, &KLine{
			Time:           util.ProtoStr(kl.Time),
			IsBlank:        util.ProtoBool(kl.IsBlank),
			HighPrice:      util.ProtoFloat64(kl.HighPrice),
			OpenPrice:      util.ProtoFloat64(kl.OpenPrice),
			LowPrice:       util.ProtoFloat64(kl.LowPrice),
			ClosePrice:     util.ProtoFloat64(kl.ClosePrice),
			LastClosePrice: util.ProtoFloat64(kl.LastClosePrice),
			Volume:         util.ProtoInt64(kl.Volume),
			Turnover:       util.ProtoFloat64(kl.Turnover),
			TurnoverRate:   util.ProtoFloat64(kl.TurnoverRate),
			Pe:             util.ProtoFloat64(kl.Pe),
			ChangeRate:     util.ProtoFloat64(kl.ChangeRate),
			Timestamp:      util.ProtoFloat64(kl.Timestamp),
		})
	}

	return result, nil
}

// GetHistoryKLRequest defines parameters for GetHistoryKL (deprecated, use RequestHistoryKL for pagination).
type GetHistoryKLRequest struct {
	RehabType        int32
	KlType           int32
	Security         *qotcommon.Security
	BeginTime        string
	EndTime          string
	MaxAckKLNum      int32
	NeedKLFieldsFlag int64
}

// GetHistoryKLResponse is the response type for GetHistoryKL.
type GetHistoryKLResponse struct {
	Security        *qotcommon.Security
	KLList          []*KLine
	NextKLTime      string
	NextKLTimestamp float64
}

// GetHistoryKL returns historical K-line data for the given security.
// Deprecated: Removed in Futu v10.6 proto — proto package qotgethistorykl no longer exists.
// Use RequestHistoryKL for paginated requests via NextReqKey.
func GetHistoryKL(ctx context.Context, c *futuapi.Client, req *GetHistoryKLRequest) (*GetHistoryKLResponse, error) {
	// The underlying Qot_GetHistoryKL proto was removed in Futu v10.6.
	// This function is kept as a stub to avoid breaking existing callers.
	return nil, fmt.Errorf("GetHistoryKL: removed in Futu v10.6 — use RequestHistoryKL instead")
}

type noopGetHistoryKLResponse struct{}

func (noopGetHistoryKLResponse) GetCachedSchema() any { return nil }


// NoDataMode specifies how to return data when the requested time point is empty.
// Deprecated: Removed in Futu v10.6 proto — proto package qotgethistoryklpoints no longer exists.
type NoDataMode = int32

const (
	NoDataMode_Null     NoDataMode = 0
	NoDataMode_Forward  NoDataMode = 1
	NoDataMode_Backward NoDataMode = 2
)

// DataStatus indicates the status and source of the data returned for a time point.
// Deprecated: Removed in Futu v10.6 proto — proto package qotgethistoryklpoints no longer exists.
type DataStatus = int32

const (
	DataStatus_Null     DataStatus = 0
	DataStatus_Current  DataStatus = 1
	DataStatus_Previous DataStatus = 2
	DataStatus_Back     DataStatus = 3
)

// GetHistoryKLPointsRequest represents the request for historical K-line points.
type GetHistoryKLPointsRequest struct {
	RehabType           constant.RehabType
	KLType             constant.KLType
	NoDataMode         NoDataMode
	Securities         []*qotcommon.Security
	Times              []string
	MaxReqSecuritiesNum int32
	NeedKLFieldsFlag   int64
}

// HistoryPointsKL represents a single K-line data point at a specific time.
type HistoryPointsKL struct {
	Status  DataStatus
	ReqTime string
	KL     *qotcommon.KLine
}

// SecurityHistoryKLPoints represents K-line points for a single security.
type SecurityHistoryKLPoints struct {
	Security *qotcommon.Security
	KLList   []*HistoryPointsKL
}

// GetHistoryKLPointsResponse represents the response with historical K-line points.
type GetHistoryKLPointsResponse struct {
	KLPointList []*SecurityHistoryKLPoints
	HasNext     bool
}

// GetHistoryKLPoints retrieves historical K-line data at specific time points.
// Deprecated: Removed in Futu v10.6 proto — proto package qotgethistoryklpoints no longer exists.
func GetHistoryKLPoints(ctx context.Context, c *futuapi.Client, req *GetHistoryKLPointsRequest) (*GetHistoryKLPointsResponse, error) {
	// The underlying Qot_GetHistoryKLPoints proto was removed in Futu v10.6.
	return nil, fmt.Errorf("GetHistoryKLPoints: removed in Futu v10.6")
}

type noopGetHistoryKLPointsResponse struct{}

func (noopGetHistoryKLPointsResponse) GetCachedSchema() any { return nil }

// RequestHistoryKLQuotaRequest defines parameters for RequestHistoryKLQuota.
type RequestHistoryKLQuotaRequest struct {
	GetDetail bool
}

// RequestHistoryKLQuotaResponse is the response type for RequestHistoryKLQuota.
type RequestHistoryKLQuotaResponse struct {
	UsedQuota   int32
	RemainQuota int32
	DetailList  []*qotrequesthistoryklquota.DetailItem
}

// RequestHistoryKLQuota returns the quota usage for historical K-line requests.
func RequestHistoryKLQuota(ctx context.Context, c *futuapi.Client, req *RequestHistoryKLQuotaRequest) (*RequestHistoryKLQuotaResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("RequestHistoryKLQuota: request is nil")
	}
	c2s := &qotrequesthistoryklquota.C2S{
		BGetDetail: &req.GetDetail,
	}

	pkt := &qotrequesthistoryklquota.Request{C2S: c2s}
	var rsp qotrequesthistoryklquota.Response

	if err := c.RequestContext(ctx, ProtoID_RequestHistoryKLQuota, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("RequestHistoryKLQuota", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("RequestHistoryKLQuota", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &RequestHistoryKLQuotaResponse{
		UsedQuota:   util.ProtoInt32(s2c.UsedQuota),
		RemainQuota: util.ProtoInt32(s2c.RemainQuota),
		DetailList:  make([]*qotrequesthistoryklquota.DetailItem, 0, len(s2c.DetailList)),
	}

	for _, d := range s2c.DetailList {
		if d == nil {
			continue
		}
		result.DetailList = append(result.DetailList, d)
	}

	return result, nil
}
