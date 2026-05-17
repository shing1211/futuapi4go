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

	"google.golang.org/protobuf/proto"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgethistorykl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgethistoryklpoints"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistorykl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistoryklquota"
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

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("RequestHistoryKL", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("RequestHistoryKL", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &RequestHistoryKLResponse{
		Security:   s2c.GetSecurity(),
		Name:       s2c.GetName(),
		NextReqKey: s2c.GetNextReqKey(),
		KLList:     make([]*KLine, 0, len(s2c.GetKlList())),
	}

	for _, kl := range s2c.GetKlList() {
		if kl == nil {
			continue
		}
		result.KLList = append(result.KLList, &KLine{
			Time:           kl.GetTime(),
			IsBlank:        kl.GetIsBlank(),
			HighPrice:      kl.GetHighPrice(),
			OpenPrice:      kl.GetOpenPrice(),
			LowPrice:       kl.GetLowPrice(),
			ClosePrice:     kl.GetClosePrice(),
			LastClosePrice: kl.GetLastClosePrice(),
			Volume:         kl.GetVolume(),
			Turnover:       kl.GetTurnover(),
			TurnoverRate:   kl.GetTurnoverRate(),
			Pe:             kl.GetPe(),
			ChangeRate:     kl.GetChangeRate(),
			Timestamp:      kl.GetTimestamp(),
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
// Deprecated: Use RequestHistoryKL for paginated requests via NextReqKey.
func GetHistoryKL(ctx context.Context, c *futuapi.Client, req *GetHistoryKLRequest) (*GetHistoryKLResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHistoryKL: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetHistoryKL: security is required")
	}

	c2s := &qotgethistorykl.C2S{
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

	pkt := &qotgethistorykl.Request{C2S: c2s}
	var rsp qotgethistorykl.Response

	if err := c.RequestContext(ctx, ProtoID_GetHistoryKL, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHistoryKL", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetHistoryKL", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetHistoryKLResponse{
		Security:        s2c.GetSecurity(),
		NextKLTime:      s2c.GetNextKLTime(),
		NextKLTimestamp: s2c.GetNextKLTimestamp(),
		KLList:          make([]*KLine, 0, len(s2c.GetKlList())),
	}

	for _, kl := range s2c.GetKlList() {
		if kl == nil {
			continue
		}
		result.KLList = append(result.KLList, &KLine{
			Time:           kl.GetTime(),
			IsBlank:        kl.GetIsBlank(),
			HighPrice:      kl.GetHighPrice(),
			OpenPrice:      kl.GetOpenPrice(),
			LowPrice:       kl.GetLowPrice(),
			ClosePrice:     kl.GetClosePrice(),
			LastClosePrice: kl.GetLastClosePrice(),
			Volume:         kl.GetVolume(),
			Turnover:       kl.GetTurnover(),
			TurnoverRate:   kl.GetTurnoverRate(),
			Pe:             kl.GetPe(),
			ChangeRate:     kl.GetChangeRate(),
			Timestamp:      kl.GetTimestamp(),
		})
	}

	return result, nil
}


// NoDataMode specifies how to return data when the requested time point is empty.
type NoDataMode = qotgethistoryklpoints.NoDataMode

const (
	NoDataMode_Null     NoDataMode = 0
	NoDataMode_Forward  NoDataMode = 1
	NoDataMode_Backward NoDataMode = 2
)

// DataStatus indicates the status and source of the data returned for a time point.
type DataStatus = qotgethistoryklpoints.DataStatus

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
func GetHistoryKLPoints(ctx context.Context, c *futuapi.Client, req *GetHistoryKLPointsRequest) (*GetHistoryKLPointsResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHistoryKLPoints: request is nil")
	}
	if len(req.Securities) == 0 {
		return nil, fmt.Errorf("GetHistoryKLPoints: securities is required")
	}
	if len(req.Times) == 0 {
		return nil, fmt.Errorf("GetHistoryKLPoints: times is required")
	}

	c2s := &qotgethistoryklpoints.C2S{
		RehabType: func() *int32 { v := int32(req.RehabType); return &v }(),
		KlType:    func() *int32 { v := int32(req.KLType); return &v }(),
		NoDataMode: func() *int32 { v := int32(req.NoDataMode); return &v }(),
		SecurityList: req.Securities,
		TimeList:    req.Times,
	}
	if req.MaxReqSecuritiesNum > 0 {
		c2s.MaxReqSecurityNum = &req.MaxReqSecuritiesNum
	}
	if req.NeedKLFieldsFlag > 0 {
		c2s.NeedKLFieldsFlag = &req.NeedKLFieldsFlag
	}

	pkt := &qotgethistoryklpoints.Request{C2S: c2s}
	var rsp qotgethistoryklpoints.Response

	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetHistoryKLPoints, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetHistoryKLPoints", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetHistoryKLPoints", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	klPointList := s2c.GetKlPointList()
	if klPointList == nil {
		return &GetHistoryKLPointsResponse{}, nil
	}

	result := make([]*SecurityHistoryKLPoints, 0, len(klPointList))
	for _, shkp := range klPointList {
		if shkp == nil {
			continue
		}

		klList := shkp.GetKlList()
		parsedKLList := make([]*HistoryPointsKL, 0, len(klList))
		for _, kl := range klList {
			if kl == nil {
				continue
			}
			parsedKLList = append(parsedKLList, &HistoryPointsKL{
				Status:  DataStatus(kl.GetStatus()),
				ReqTime: kl.GetReqTime(),
				KL:     kl.GetKl(),
			})
		}

		result = append(result, &SecurityHistoryKLPoints{
			Security: shkp.GetSecurity(),
			KLList:   parsedKLList,
		})
	}

	return &GetHistoryKLPointsResponse{
		KLPointList: result,
		HasNext:   s2c.GetHasNext(),
	}, nil
}

var _ proto.Message = (*qotgethistoryklpoints.Request)(nil)
var _ proto.Message = (*qotgethistoryklpoints.Response)(nil)

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

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("RequestHistoryKLQuota", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("RequestHistoryKLQuota", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &RequestHistoryKLQuotaResponse{
		UsedQuota:   s2c.GetUsedQuota(),
		RemainQuota: s2c.GetRemainQuota(),
		DetailList:  make([]*qotrequesthistoryklquota.DetailItem, 0, len(s2c.GetDetailList())),
	}

	for _, d := range s2c.GetDetailList() {
		if d == nil {
			continue
		}
		result.DetailList = append(result.DetailList, d)
	}

	return result, nil
}
