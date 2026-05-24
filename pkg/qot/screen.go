package qot

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotoptionscreen"
	"github.com/shing1211/futuapi4go/pkg/pb/qotstockscreen"
	"github.com/shing1211/futuapi4go/pkg/pb/qotwarrantscreen"
	"github.com/shing1211/futuapi4go/pkg/util"
)

// ============================================================================
// StockScreen — 条件选股(新版股票筛选)
// ============================================================================

// StockScreenRequest defines parameters for StockScreen.
type StockScreenRequest struct {
	FilterList        []*qotstockscreen.ScreenQuery
	RetrieveList      []*qotstockscreen.RetrieveQuery
	Sort              *qotstockscreen.Sort
	WatchlistStockIds []uint64
	SortList          []*qotstockscreen.Sort
	PageFrom          int32
	PageCount         int32
}

// StockScreenItem represents a single stock screening result.
type StockScreenItem struct {
	StockId uint64
	Results []*qotstockscreen.RspItemResult
}

// StockScreenResponse is the response type for StockScreen.
type StockScreenResponse struct {
	LastPage bool
	AllCount int32
	DataList []*StockScreenItem
}

// StockScreen filters stocks using the new v10.6 stock screening engine.
func StockScreen(ctx context.Context, c *futuapi.Client, req *StockScreenRequest) (*StockScreenResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("StockScreen: request is nil")
	}
	if req.PageCount <= 0 || req.PageCount > 300 {
		return nil, fmt.Errorf("StockScreen: PageCount must be between 1 and 300")
	}

	c2s := &qotstockscreen.C2S{
		FilterList:        req.FilterList,
		RetrieveList:      req.RetrieveList,
		Sort:              req.Sort,
		WatchlistStockIds: req.WatchlistStockIds,
		SortList:          req.SortList,
		PageFrom:          &req.PageFrom,
		PageCount:         &req.PageCount,
	}
	pkt := &qotstockscreen.Request{C2S: c2s}
	var rsp qotstockscreen.Response

	if err := c.RequestContext(ctx, ProtoID_StockScreen, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("StockScreen", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("StockScreen", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &StockScreenResponse{
		LastPage: util.ProtoInt32(s2c.LastPage) == 1,
		AllCount: util.ProtoInt32(s2c.AllCount),
		DataList: make([]*StockScreenItem, 0, len(s2c.DataList)),
	}

	for _, d := range s2c.DataList {
		if d == nil {
			continue
		}
		result.DataList = append(result.DataList, &StockScreenItem{
			StockId: util.ProtoUint64(d.StockId),
			Results: d.Results,
		})
	}

	return result, nil
}

// ============================================================================
// WarrantScreen — 窝轮筛选
// ============================================================================

// WarrantScreenRequest defines parameters for WarrantScreen.
type WarrantScreenRequest struct {
	MarketType int32
	IsDelay    bool
	FilterList []*qotwarrantscreen.ScreenGroup
	SortList   []*qotwarrantscreen.Sort
	OnlyCount  bool
	PageFrom   int32
	PageCount  int32
}

// WarrantScreenResponse is the response type for WarrantScreen.
type WarrantScreenResponse struct {
	LastPage  bool
	AllCount  int32
	Warrants  []*qotwarrantscreen.WarrantItem
}

// WarrantScreen filters warrants (窝轮) using the v10.6 screening engine.
func WarrantScreen(ctx context.Context, c *futuapi.Client, req *WarrantScreenRequest) (*WarrantScreenResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("WarrantScreen: request is nil")
	}
	if req.PageCount <= 0 {
		return nil, fmt.Errorf("WarrantScreen: PageCount must be positive")
	}

	c2s := &qotwarrantscreen.C2S{
		MarketType: &req.MarketType,
		IsDelay:    &req.IsDelay,
		FilterList: req.FilterList,
		SortList:   req.SortList,
		OnlyCount:  &req.OnlyCount,
		PageFrom:   &req.PageFrom,
		PageCount:  &req.PageCount,
	}
	pkt := &qotwarrantscreen.Request{C2S: c2s}
	var rsp qotwarrantscreen.Response

	if err := c.RequestContext(ctx, ProtoID_WarrantScreen, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("WarrantScreen", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("WarrantScreen", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &WarrantScreenResponse{
		LastPage: util.ProtoBool(s2c.LastPage),
		AllCount: util.ProtoInt32(s2c.AllCount),
		Warrants: make([]*qotwarrantscreen.WarrantItem, 0, len(s2c.Warrants)),
	}

	for _, w := range s2c.Warrants {
		if w == nil {
			continue
		}
		result.Warrants = append(result.Warrants, w)
	}

	return result, nil
}

// ============================================================================
// OptionScreen — 期权筛选
// ============================================================================

// OptionScreenRequest defines parameters for OptionScreen.
type OptionScreenRequest struct {
	MarketCategoryList     []int32
	FilterList             []*qotoptionscreen.FilterGroup
	SortList               []*qotoptionscreen.Sort
	OptionRetrieveList     []int32
	UnderlyingRetrieveList []int32
	PageFrom               int32
	PageCount              int32
}

// OptionScreenResponse is the response type for OptionScreen.
type OptionScreenResponse struct {
	LastPage bool
	AllCount int32
	DataList []*qotoptionscreen.OptionScreenItem
}

// OptionScreen filters options using the v10.6 screening engine.
func OptionScreen(ctx context.Context, c *futuapi.Client, req *OptionScreenRequest) (*OptionScreenResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("OptionScreen: request is nil")
	}
	if len(req.MarketCategoryList) == 0 {
		return nil, fmt.Errorf("OptionScreen: MarketCategoryList must have at least one category")
	}
	if req.PageCount <= 0 || req.PageCount > 1000 {
		return nil, fmt.Errorf("OptionScreen: PageCount must be between 1 and 1000")
	}

	c2s := &qotoptionscreen.C2S{
		MarketCategoryList:     req.MarketCategoryList,
		FilterList:             req.FilterList,
		SortList:               req.SortList,
		OptionRetrieveList:     req.OptionRetrieveList,
		UnderlyingRetrieveList: req.UnderlyingRetrieveList,
		PageFrom:               &req.PageFrom,
		PageCount:              &req.PageCount,
	}
	pkt := &qotoptionscreen.Request{C2S: c2s}
	var rsp qotoptionscreen.Response

	if err := c.RequestContext(ctx, ProtoID_OptionScreen, pkt, &rsp); err != nil {
		return nil, err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("OptionScreen", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	s2c := rsp.S2C
	if s2c == nil {
		return nil, wrapError("OptionScreen", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	result := &OptionScreenResponse{
		LastPage: util.ProtoBool(s2c.LastPage),
		AllCount: util.ProtoInt32(s2c.AllCount),
		DataList: make([]*qotoptionscreen.OptionScreenItem, 0, len(s2c.DataList)),
	}

	for _, d := range s2c.DataList {
		if d == nil {
			continue
		}
		result.DataList = append(result.DataList, d)
	}

	return result, nil
}
