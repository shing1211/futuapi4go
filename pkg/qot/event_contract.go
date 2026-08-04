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
	"github.com/shing1211/futuapi4go/pkg/pb/qotfiltercompetition"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontract"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractcategory"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractcombolist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractcomborfq"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontracteventlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractkline"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractmilestonelist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractserieslist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractsnapshot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistoryeventcontractkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsubeventcontract"
	"github.com/shing1211/futuapi4go/pkg/util"
)

// ECSubType represents the type of Event Contract real-time data subscription.
// Values map to SubType used for EC subscribe.
type ECSubType int32

const (
	ECSubType_OrderBook ECSubType = 1 // 摆盘
	ECSubType_Kline     ECSubType = 2 // K线
	ECSubType_Ticker    ECSubType = 3 // 逐笔
)

// FilterCompetition returns the list of available competition filters for the
// event contract platform. C2S.Category / C2S.Tag are optional; leave empty to
// fetch the full tree.
func FilterCompetition(ctx context.Context, c *futuapi.Client, req *qotfiltercompetition.C2S) (*qotfiltercompetition.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("FilterCompetition: req is nil")
	}
	var rsp qotfiltercompetition.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_FilterCompetition, &qotfiltercompetition.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("FilterCompetition", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("FilterCompetition", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractCategory returns the top-level event contract category list.
func GetEventContractCategory(ctx context.Context, c *futuapi.Client, req *qotgeteventcontractcategory.C2S) (*qotgeteventcontractcategory.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractCategory: req is nil")
	}
	var rsp qotgeteventcontractcategory.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractCategory, &qotgeteventcontractcategory.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractCategory", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractCategory", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractSeriesList returns the list of Event Contract Series under
// the given category/tag filter.
func GetEventContractSeriesList(ctx context.Context, c *futuapi.Client, req *qotgeteventcontractserieslist.C2S) (*qotgeteventcontractserieslist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractSeriesList: req is nil")
	}
	var rsp qotgeteventcontractserieslist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractSeriesList, &qotgeteventcontractserieslist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractSeriesList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractSeriesList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractEventList returns the list of Event Contract Events under the
// given Series. Supports pagination via NextPage.
func GetEventContractEventList(ctx context.Context, c *futuapi.Client, req *qotgeteventcontracteventlist.C2S) (*qotgeteventcontracteventlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractEventList: req is nil")
	}
	if req.Series == nil {
		return nil, fmt.Errorf("GetEventContractEventList: series is required")
	}
	var rsp qotgeteventcontracteventlist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractEventList, &qotgeteventcontracteventlist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractEventList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractEventList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContract returns the list of Event Contract contracts under the
// given Event. Supports pagination via NextPage.
func GetEventContract(ctx context.Context, c *futuapi.Client, req *qotgeteventcontract.C2S) (*qotgeteventcontract.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContract: req is nil")
	}
	if req.Event == nil {
		return nil, fmt.Errorf("GetEventContract: event is required")
	}
	var rsp qotgeteventcontract.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContract, &qotgeteventcontract.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContract", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContract", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractMilestoneList returns the list of Event Contract milestones.
func GetEventContractMilestoneList(ctx context.Context, c *futuapi.Client, req *qotgeteventcontractmilestonelist.C2S) (*qotgeteventcontractmilestonelist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractMilestoneList: req is nil")
	}
	var rsp qotgeteventcontractmilestonelist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractMilestoneList, &qotgeteventcontractmilestonelist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractMilestoneList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractMilestoneList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractSnapshot returns a batch of Event Contract snapshots.
// C2S.SecurityList must be non-empty.
func GetEventContractSnapshot(ctx context.Context, c *futuapi.Client, req *qotgeteventcontractsnapshot.C2S) (*qotgeteventcontractsnapshot.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractSnapshot: req is nil")
	}
	if len(req.SecurityList) == 0 {
		return nil, fmt.Errorf("GetEventContractSnapshot: security list is empty")
	}
	var rsp qotgeteventcontractsnapshot.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractSnapshot, &qotgeteventcontractsnapshot.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractSnapshot", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractSnapshot", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractOrderBook returns the Event Contract order book snapshot.
// Requires an active EC_ORDER_BOOK subscription for the contract.
func GetEventContractOrderBook(ctx context.Context, c *futuapi.Client, req *qotgeteventcontractorderbook.C2S) (*qotgeteventcontractorderbook.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractOrderBook: req is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetEventContractOrderBook: security is required")
	}
	var rsp qotgeteventcontractorderbook.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractOrderBook, &qotgeteventcontractorderbook.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractOrderBook", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractOrderBook", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractKline returns Event Contract K-line snapshot.
// Requires an active EC_KLINE subscription for the contract.
func GetEventContractKline(ctx context.Context, c *futuapi.Client, req *qotgeteventcontractkline.C2S) (*qotgeteventcontractkline.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractKline: req is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetEventContractKline: security is required")
	}
	var rsp qotgeteventcontractkline.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractKline, &qotgeteventcontractkline.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractKline", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractKline", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractTicker returns Event Contract tick-by-tick snapshot.
// Requires an active EC_TICKER subscription for the contract.
func GetEventContractTicker(ctx context.Context, c *futuapi.Client, req *qotgeteventcontractticker.C2S) (*qotgeteventcontractticker.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractTicker: req is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("GetEventContractTicker: security is required")
	}
	var rsp qotgeteventcontractticker.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractTicker, &qotgeteventcontractticker.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractTicker", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractTicker", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// RequestHistoryEventContractKL pulls historical Event Contract K-line data
// using paginated requests via NextReqKey.
func RequestHistoryEventContractKL(ctx context.Context, c *futuapi.Client, req *qotrequesthistoryeventcontractkl.C2S) (*qotrequesthistoryeventcontractkl.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("RequestHistoryEventContractKL: req is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("RequestHistoryEventContractKL: security is required")
	}
	var rsp qotrequesthistoryeventcontractkl.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_RequestHistoryEventContractKL, &qotrequesthistoryeventcontractkl.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("RequestHistoryEventContractKL", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("RequestHistoryEventContractKL", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractComboList returns the list of Events that can be combined
// into Combo positions. The returned MVC security must be passed through to
// GetEventContractComboRfq.
func GetEventContractComboList(ctx context.Context, c *futuapi.Client, req *qotgeteventcontractcombolist.C2S) (*qotgeteventcontractcombolist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractComboList: req is nil")
	}
	var rsp qotgeteventcontractcombolist.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractComboList, &qotgeteventcontractcombolist.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractComboList", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractComboList", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// GetEventContractComboRfq submits a Combo leg list (with the MVC security
// returned by GetEventContractComboList) and returns the bid/ask price and
// the quote ID used for combo order placement.
func GetEventContractComboRfq(ctx context.Context, c *futuapi.Client, req *qotgeteventcontractcomborfq.C2S) (*qotgeteventcontractcomborfq.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractComboRfq: req is nil")
	}
	if len(req.ComboLegList) == 0 {
		return nil, fmt.Errorf("GetEventContractComboRfq: combo leg list is empty")
	}
	if req.Mvc == nil || *req.Mvc == "" {
		return nil, fmt.Errorf("GetEventContractComboRfq: mvc is required")
	}
	var rsp qotgeteventcontractcomborfq.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_GetEventContractComboRfq, &qotgeteventcontractcomborfq.Request{C2S: req}, &rsp); err != nil {
		return nil, err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetEventContractComboRfq", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	if rsp.S2C == nil {
		return nil, wrapError("GetEventContractComboRfq", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}
	return rsp.S2C, nil
}

// SubEventContract subscribes to / unsubscribes from Event Contract real-time
// data. Set IsSubOrUnSub=false to unsubscribe matching items; set
// IsUnsubAll=true to cancel every EC subscription on the connection.
func SubEventContract(ctx context.Context, c *futuapi.Client, req *qotsubeventcontract.C2S) error {
	if req == nil {
		return fmt.Errorf("SubEventContract: req is nil")
	}
	if req.IsUnsubAll == nil || !*req.IsUnsubAll {
		if len(req.SecurityList) == 0 {
			return fmt.Errorf("SubEventContract: security list is empty")
		}
		if len(req.SubTypeList) == 0 {
			return fmt.Errorf("SubEventContract: subtype list is empty")
		}
	}
	pkt := &qotsubeventcontract.Request{C2S: req}
	var rsp qotsubeventcontract.Response
	if err := c.RequestContext(ctx, constant.ProtoID_Qot_SubEventContract, pkt, &rsp); err != nil {
		return err
	}
	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return wrapError("SubEventContract", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}
	return nil
}

// BuildECSecurity constructs a Security message for an Event Contract
// instrument (market = QotMarket_EventContract, code = e.g. "EC.xxx").
func BuildECSecurity(code string) *qotcommon.Security {
	market := int32(qotcommon.QotMarket_QotMarket_EventContract)
	return &qotcommon.Security{
		Market: &market,
		Code:   &code,
	}
}