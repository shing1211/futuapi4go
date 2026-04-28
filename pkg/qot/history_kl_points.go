// Package qot provides market data APIs for the Futu OpenD SDK.
//
// This package covers real-time quotes, K-lines, order book, tick data,
// broker queue, capital flow, stock screening, options, warrants, and
// historical data requests. All functions require a connected client.
//
// For Python SDK migration, use the constant package for Python-style constants:
//
//	import (
//	    "github.com/shing1211/futuapi4go/pkg/constant"
//	    "github.com/shing1211/futuapi4go/pkg/qot"
//	    "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
//	)
//
//	// Create security with Python-style constants
//	market := constant.Market_HK
//	code := "00700"
//	securities := []*qotcommon.Security{
//	    {Market: &market, Code: &code},
//	}
//
//	// Use constant package for K-line types, rehab types, etc.
//	// KLType: constant.KLType_K_Day, constant.KLType_K_1Min, etc.
//	// RehabType: constant.RehabType_Forward (QFQ), constant.RehabType_Backward (HFQ)
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

//go:generate bash scripts/generate-enums.sh qotgethistoryklpoints NoDataMode DataStatus

package qot

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgethistoryklpoints"
)

const (
	ProtoID_GetHistoryKLPoints = 3106
)

// NoDataMode specifies how to return data when the requested time point is empty.
type NoDataMode = qotgethistoryklpoints.NoDataMode

const (
	NoDataMode_Null     NoDataMode = 0 // 直接返回空数据
	NoDataMode_Forward  NoDataMode = 1 // 往前取值，返回前一个时间点数据
	NoDataMode_Backward NoDataMode = 2 // 向后取值，返回后一个时间点数据
)

// DataStatus indicates the status and source of the data returned for a time point.
type DataStatus = qotgethistoryklpoints.DataStatus

const (
	DataStatus_Null     DataStatus = 0 // 空数据
	DataStatus_Current  DataStatus = 1 // 当前时间点数据
	DataStatus_Previous DataStatus = 2 // 前一个时间点数据
	DataStatus_Back     DataStatus = 3 // 后一个时间点数据
)

// GetHistoryKLPointsRequest represents the request for historical K-line points.
type GetHistoryKLPointsRequest struct {
	RehabType    constant.RehabType // 复权类型
	KLType      constant.KLType   // K线类型
	NoDataMode  NoDataMode       // 当请求时间点数据为空时，如何返回数据
	Securities  []*qotcommon.Security
	Times       []string // 时间字符串 (e.g., "2024-01-01 09:30:00")
	MaxReqSecuritiesNum int32   // 最多返回多少只股票的数据，如果未指定表示不限制
	NeedKLFieldsFlag   int64   // 指定返回K线结构体特定某几项数据
}

// HistoryPointsKL represents a single K-line data point at a specific time.
type HistoryPointsKL struct {
	Status   DataStatus     // 数据状态
	ReqTime  string       // 请求的时间
	KL      *qotcommon.KLine
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
// This is useful for backtesting and point-in-time analysis.
//
// Parameters:
//   - ctx: Context for the request
//   - c: Futu API client
//   - req: Request containing securities and time points
//
// Returns historical K-line data for each security at each time point,
// or an error if the request fails.
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
		return nil, fmt.Errorf("GetHistoryKLPoints: s2c is nil")
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