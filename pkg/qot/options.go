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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfutureinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionchain"
	qotgetoptionexpirationdate "github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionexpirationdate"
)

// GetOptionExpirationDateRequest defines parameters for GetOptionExpirationDate.
type GetOptionExpirationDateRequest struct {
	Owner           *qotcommon.Security
	IndexOptionType int32
}

// OptionExpirationDateInfo represents information about an option expiration date.
type OptionExpirationDateInfo struct {
	StrikeTime               string
	StrikeTimestamp          float64
	OptionExpiryDateDistance int32
	Cycle                    int32
}

// GetOptionExpirationDateResponse is the response type for GetOptionExpirationDate.
type GetOptionExpirationDateResponse struct {
	DateList []*OptionExpirationDateInfo
}

// GetOptionExpirationDate returns the list of option expiration dates for the given underlying.
func GetOptionExpirationDate(ctx context.Context, c *futuapi.Client, req *GetOptionExpirationDateRequest) (*GetOptionExpirationDateResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionExpirationDate: request is nil")
	}
	if req.Owner == nil {
		return nil, fmt.Errorf("owner security is required")
	}

	c2s := &qotgetoptionexpirationdate.C2S{
		Owner:           req.Owner,
		IndexOptionType: &req.IndexOptionType,
	}

	pkt := &qotgetoptionexpirationdate.Request{C2S: c2s}
	var rsp qotgetoptionexpirationdate.Response

	if err := c.RequestContext(ctx, ProtoID_GetOptionExpirationDate, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionExpirationDate", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetOptionExpirationDate", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetOptionExpirationDateResponse{
		DateList: make([]*OptionExpirationDateInfo, 0, len(s2c.GetDateList())),
	}

	for _, d := range s2c.GetDateList() {
		if d == nil {
			continue
		}
		result.DateList = append(result.DateList, &OptionExpirationDateInfo{
			StrikeTime:               d.GetStrikeTime(),
			StrikeTimestamp:          d.GetStrikeTimestamp(),
			OptionExpiryDateDistance: d.GetOptionExpiryDateDistance(),
			Cycle:                    d.GetCycle(),
		})
	}

	return result, nil
}

// GetOptionChainRequest defines parameters for GetOptionChain.
type GetOptionChainRequest struct {
	Owner           *qotcommon.Security
	IndexOptionType int32
	Type            int32
	Condition       int32
	BeginTime       string
	EndTime         string
	DataFilter      *qotgetoptionchain.DataFilter
}

// OptionItem represents a pair of call and put options at the same strike price.
type OptionItem struct {
	Call *qotcommon.SecurityStaticInfo
	Put  *qotcommon.SecurityStaticInfo
}

// OptionChain represents the option chain for a single expiration date.
type OptionChain struct {
	StrikeTime      string
	StrikeTimestamp float64
	Option          []*OptionItem
}

// GetOptionChainResponse is the response type for GetOptionChain.
type GetOptionChainResponse struct {
	OptionChain []*OptionChain
}

// GetOptionChain returns the option chain (期权链) for the given underlying security.
func GetOptionChain(ctx context.Context, c *futuapi.Client, req *GetOptionChainRequest) (*GetOptionChainResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionChain: request is nil")
	}
	if req.Owner == nil {
		return nil, fmt.Errorf("owner security is required")
	}

	c2s := &qotgetoptionchain.C2S{
		Owner:           req.Owner,
		IndexOptionType: &req.IndexOptionType,
		Type:            &req.Type,
		Condition:       &req.Condition,
		BeginTime:       &req.BeginTime,
		EndTime:         &req.EndTime,
		DataFilter:      req.DataFilter,
	}

	pkt := &qotgetoptionchain.Request{C2S: c2s}
	var rsp qotgetoptionchain.Response

	if err := c.RequestContext(ctx, ProtoID_GetOptionChain, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetOptionChain", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetOptionChain", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetOptionChainResponse{
		OptionChain: make([]*OptionChain, 0, len(s2c.GetOptionChain())),
	}

	for _, chain := range s2c.GetOptionChain() {
		if chain == nil {
			continue
		}
		oc := &OptionChain{
			StrikeTime:      chain.GetStrikeTime(),
			StrikeTimestamp: chain.GetStrikeTimestamp(),
			Option:          make([]*OptionItem, 0, len(chain.GetOption())),
		}

		for _, opt := range chain.GetOption() {
			if opt == nil {
				continue
			}
			item := &OptionItem{}
			if opt.GetCall() != nil {
				item.Call = opt.GetCall()
			}
			if opt.GetPut() != nil {
				item.Put = opt.GetPut()
			}
			oc.Option = append(oc.Option, item)
		}

		result.OptionChain = append(result.OptionChain, oc)
	}

	return result, nil
}

// GetFutureInfoRequest defines parameters for GetFutureInfo.
type GetFutureInfoRequest struct {
	SecurityList []*qotcommon.Security
}

// FutureInfo represents detailed information about a futures contract.
type FutureInfo struct {
	Name               string
	Security           *qotcommon.Security
	LastTradeTime      string
	LastTradeTimestamp float64
	Owner              *qotcommon.Security
	OwnerOther         string
	Exchange           string
	ContractType       string
	ContractSize       float64
	ContractSizeUnit   string
	QuoteCurrency      string
	MinVar             float64
	MinVarUnit         string
	QuoteUnit          string
	TradeTimeList      []*qotgetfutureinfo.TradeTime
	TimeZone           string
	ExchangeFormatUrl  string
	Origin             *qotcommon.Security
}

// GetFutureInfoResponse is the response type for GetFutureInfo.
type GetFutureInfoResponse struct {
	FutureInfoList []*FutureInfo
}

// GetFutureInfo returns futures contract information for the given securities.
func GetFutureInfo(ctx context.Context, c *futuapi.Client, req *GetFutureInfoRequest) (*GetFutureInfoResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFutureInfo: request is nil")
	}
	if len(req.SecurityList) == 0 {
		return nil, fmt.Errorf("security list is empty")
	}

	c2s := &qotgetfutureinfo.C2S{
		SecurityList: req.SecurityList,
	}

	pkt := &qotgetfutureinfo.Request{C2S: c2s}
	var rsp qotgetfutureinfo.Response

	if err := c.RequestContext(ctx, ProtoID_GetFutureInfo, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetFutureInfo", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("GetFutureInfo", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &GetFutureInfoResponse{
		FutureInfoList: make([]*FutureInfo, 0, len(s2c.GetFutureInfoList())),
	}

	for _, fi := range s2c.GetFutureInfoList() {
		if fi == nil {
			continue
		}
		result.FutureInfoList = append(result.FutureInfoList, &FutureInfo{
			Name:               fi.GetName(),
			Security:           fi.GetSecurity(),
			LastTradeTime:      fi.GetLastTradeTime(),
			LastTradeTimestamp: fi.GetLastTradeTimestamp(),
			Owner:              fi.GetOwner(),
			OwnerOther:         fi.GetOwnerOther(),
			Exchange:           fi.GetExchange(),
			ContractType:       fi.GetContractType(),
			ContractSize:       fi.GetContractSize(),
			ContractSizeUnit:   fi.GetContractSizeUnit(),
			QuoteCurrency:      fi.GetQuoteCurrency(),
			MinVar:             fi.GetMinVar(),
			MinVarUnit:         fi.GetMinVarUnit(),
			QuoteUnit:          fi.GetQuoteUnit(),
			TradeTimeList:      fi.GetTradeTime(),
			TimeZone:           fi.GetTimeZone(),
			ExchangeFormatUrl:  fi.GetExchangeFormatUrl(),
			Origin:             fi.GetOrigin(),
		})
	}

	return result, nil
}
