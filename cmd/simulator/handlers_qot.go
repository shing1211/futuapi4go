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

package simulator

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetstaticinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsubinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsub"
)

func (s *Server) RegisterQotHandlers() {
	s.RegisterHandler(3004, s.handleGetBasicQot)         // GetBasicQot
	s.RegisterHandler(3006, s.handleGetKL)               // GetKL
	s.RegisterHandler(3012, s.handleGetOrderBook)        // GetOrderBook
	s.RegisterHandler(3010, s.handleGetTicker)           // GetTicker
	s.RegisterHandler(3008, s.handleGetRT)               // GetRT
	s.RegisterHandler(3203, s.handleGetSecuritySnapshot) // GetSecuritySnapshot
	s.RegisterHandler(3014, s.handleGetBroker)           // GetBroker
	s.RegisterHandler(2201, s.handleGetStaticInfo)
	s.RegisterHandler(2202, s.handleGetPlateSet)
	s.RegisterHandler(2203, s.handleGetPlateSecurity)
	s.RegisterHandler(3207, s.handleGetOwnerPlate) // GetOwnerPlate
	s.RegisterHandler(3206, s.handleGetReference)  // GetReference
	s.RegisterHandler(2205, s.handleGetTradeDate)
	s.RegisterHandler(3219, s.handleRequestTradeDate) // RequestTradeDate
	s.RegisterHandler(3223, s.handleGetMarketState)   // GetMarketState
	s.RegisterHandler(2209, s.handleGetSuspend)
	s.RegisterHandler(2210, s.handleGetCodeChange)
	s.RegisterHandler(2211, s.handleGetFutureInfo)
	s.RegisterHandler(2212, s.handleGetIpoList)
	s.RegisterHandler(2213, s.handleGetHoldingChangeList)
	s.RegisterHandler(2214, s.handleRequestRehab)
	s.RegisterHandler(3211, s.handleGetCapitalFlow)          // GetCapitalFlow
	s.RegisterHandler(3212, s.handleGetCapitalDistribution)  // GetCapitalDistribution
	s.RegisterHandler(3215, s.handleStockFilter)             // StockFilter
	s.RegisterHandler(3209, s.handleGetOptionChain)          // GetOptionChain
	s.RegisterHandler(3224, s.handleGetOptionExpirationDate) // GetOptionExpirationDate
	s.RegisterHandler(3210, s.handleGetWarrant)              // GetWarrant
	s.RegisterHandler(2401, s.handleGetUserSecurity)
	s.RegisterHandler(2402, s.handleGetUserSecurityGroup)
	s.RegisterHandler(2403, s.handleModifyUserSecurity)
	s.RegisterHandler(2404, s.handleGetPriceReminder)
	s.RegisterHandler(2405, s.handleSetPriceReminder)
	s.RegisterHandler(3001, s.handleSubscribe)
	s.RegisterHandler(3003, s.handleGetSubInfo) // GetSubInfo
	s.RegisterHandler(3003, s.handleRegQotPush) // RegQotPush
}

func (s *Server) handleGetBasicQot(pkt *Packet) (*Packet, error) {
	var req qotgetbasicqot.Request
	if err := proto.Unmarshal(pkt.Body, &req); err != nil {
		return s.errorResponse(pkt, err)
	}
	c2s := req.GetC2S()
	_ = c2s // suppress unused warning

	retType := int32(common.RetType_RetType_Succeed)

	bqList := make([]*qotcommon.BasicQot, 0)
	for _, sec := range c2s.GetSecurityList() {
		market := sec.GetMarket()
		code := sec.GetCode()
		key := fmt.Sprintf("%d.%s", market, code)

		if quote, ok := s.Quotes[key]; ok {
			bqList = append(bqList, quote)
		} else {
			// Detect HSI
			isHSI := market == 1 && code == "800100"

			name := fmt.Sprintf("Mock_%s", code)
			price := 100.0
			high := 105.0
			low := 95.0
			open := 99.0
			lastClose := 98.5
			vol := int64(1000000)
			turn := 100000000.0
			turnoverRate := 0.01
			amplitude := 1.5

			if isHSI {
				name = "HANG SENG INDEX"
				price = 18523.45
				high = 18590.12
				low = 18420.50
				open = 18480.00
				lastClose = 18498.23
				vol = 2345678900
				turn = 98765432100.50
				turnoverRate = 0.0234
				amplitude = 0.92
			}

			bq := &qotcommon.BasicQot{
				Security:       sec,
				Name:           &name,
				CurPrice:       &price,
				HighPrice:      &high,
				LowPrice:       &low,
				OpenPrice:      &open,
				LastClosePrice: &lastClose,
				Volume:         &vol,
				Turnover:       &turn,
				TurnoverRate:   &turnoverRate,
				Amplitude:      &amplitude,
			}
			bqList = append(bqList, bq)
		}
	}

	resp := &qotgetbasicqot.Response{
		RetType: &retType,
		S2C:     &qotgetbasicqot.S2C{BasicQotList: bqList},
	}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetKL(pkt *Packet) (*Packet, error) {
	var req qotgetkl.Request
	if err := proto.Unmarshal(pkt.Body, &req); err != nil {
		return s.errorResponse(pkt, err)
	}
	c2s := req.GetC2S()

	retType := int32(common.RetType_RetType_Succeed)

	// Detect if requesting HSI
	sec := c2s.GetSecurity()
	isHSI := sec.GetMarket() == 1 && sec.GetCode() == "800100"

	name := "Mock"
	if isHSI {
		name = "HANG SENG INDEX"
	}

	klList := make([]*qotcommon.KLine, 0, c2s.GetReqNum())
	basePrice := 18400.0
	if isHSI {
		basePrice = 18400.0
	} else {
		basePrice = 100.0
	}

	for i := int32(0); i < c2s.GetReqNum(); i++ {
		timeStr := fmt.Sprintf("2026-01-%02d 15:00:00", i+1)
		ts := float64(1735689600 + int64(i)*86400)

		open := basePrice + float64(i)*2.5
		close := open + float64(i%5)*5.0 - 10.0
		high := close + float64(i%3)*3.0
		low := open - float64(i%3)*3.0
		vol := int64(1000000 + i*10000)
		turn := float64(vol) * (open + close) / 2.0
		lastClose := open

		kl := &qotcommon.KLine{
			Time:           &timeStr,
			Timestamp:      &ts,
			OpenPrice:      &open,
			ClosePrice:     &close,
			HighPrice:      &high,
			LowPrice:       &low,
			Volume:         &vol,
			Turnover:       &turn,
			LastClosePrice: &lastClose,
		}
		klList = append(klList, kl)
	}

	resp := &qotgetkl.Response{
		RetType: &retType,
		S2C:     &qotgetkl.S2C{Security: c2s.GetSecurity(), Name: &name, KlList: klList},
	}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetOrderBook(pkt *Packet) (*Packet, error) {
	var req qotgetorderbook.C2S
	if err := proto.Unmarshal(pkt.Body, &req); err != nil {
		return s.errorResponse(pkt, err)
	}

	retType := int32(common.RetType_RetType_Succeed)

	// Detect HSI
	sec := req.GetSecurity()
	isHSI := sec.GetMarket() == 1 && sec.GetCode() == "800100"

	name := "Mock"
	if isHSI {
		name = "HANG SENG INDEX"
	}

	askList := make([]*qotcommon.OrderBook, 0, req.GetNum())
	bidList := make([]*qotcommon.OrderBook, 0, req.GetNum())

	basePrice := 100.0
	if isHSI {
		basePrice = 18523.45
	}

	for i := int32(0); i < req.GetNum(); i++ {
		spread := 0.01
		if isHSI {
			spread = 5.0
		}

		askPrice := basePrice + float64(i+1)*spread
		bidPrice := basePrice - float64(i+1)*spread
		askVol := int64(1000 + i*100)
		bidVol := int64(1000 + i*100)
		orderCount := int32(10 + i*2)

		askList = append(askList, &qotcommon.OrderBook{
			Price:      &askPrice,
			Volume:     &askVol,
			OrderCount: &orderCount,
		})
		bidList = append(bidList, &qotcommon.OrderBook{
			Price:      &bidPrice,
			Volume:     &bidVol,
			OrderCount: &orderCount,
		})
	}

	resp := &qotgetorderbook.Response{
		RetType: &retType,
		S2C: &qotgetorderbook.S2C{
			Security:         req.Security,
			Name:             &name,
			OrderBookAskList: askList,
			OrderBookBidList: bidList,
		},
	}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetTicker(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetRT(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetBroker(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetStaticInfo(pkt *Packet) (*Packet, error) {
	var req qotgetstaticinfo.C2S
	if err := proto.Unmarshal(pkt.Body, &req); err != nil {
		return s.errorResponse(pkt, err)
	}

	retType := int32(common.RetType_RetType_Succeed)

	infoList := make([]*qotcommon.SecurityStaticInfo, 0, len(req.SecurityList))
	for _, sec := range req.SecurityList {
		id := int64(123)
		lotSize := int32(100)
		secType := int32(1)
		listTime := "2020-01-01"
		secName := sec.GetCode()
		basic := &qotcommon.SecurityStaticBasic{
			Security: sec,
			Id:       &id,
			LotSize:  &lotSize,
			SecType:  &secType,
			Name:     &secName,
			ListTime: &listTime,
		}
		infoList = append(infoList, &qotcommon.SecurityStaticInfo{Basic: basic})
	}

	resp := &qotgetstaticinfo.Response{
		RetType: &retType,
		S2C:     &qotgetstaticinfo.S2C{StaticInfoList: infoList},
	}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetPlateSet(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetPlateSecurity(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetOwnerPlate(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetReference(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetTradeDate(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleRequestTradeDate(pkt *Packet) (*Packet, error) {
	return s.handleGetTradeDate(pkt)
}

func (s *Server) handleGetMarketState(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetSuspend(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetCodeChange(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetFutureInfo(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetIpoList(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetHoldingChangeList(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleRequestRehab(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetCapitalFlow(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetCapitalDistribution(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleStockFilter(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetOptionChain(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetOptionExpirationDate(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetWarrant(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetUserSecurity(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetUserSecurityGroup(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleModifyUserSecurity(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetPriceReminder(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleSetPriceReminder(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleSubscribe(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	msg := "success"
	resp := &qotsub.Response{RetType: &retType, RetMsg: &msg}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetSubInfo(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	used := int32(10)
	remain := int32(90)
	resp := &qotgetsubinfo.Response{
		RetType: &retType,
		S2C:     &qotgetsubinfo.S2C{TotalUsedQuota: &used, RemainQuota: &remain},
	}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleRegQotPush(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	msg := "success"
	resp := &qotsub.Response{RetType: &retType, RetMsg: &msg}
	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetSecuritySnapshot(pkt *Packet) (*Packet, error) {
	retType := int32(common.RetType_RetType_Succeed)
	resp := &qotgetbasicqot.Response{RetType: &retType}
	return s.successResponse(pkt, resp)
}
