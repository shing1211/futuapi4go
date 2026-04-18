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

// Package client provides a public Client type for the Futu OpenD SDK.
// This allows external projects to use the SDK.
package client

import (
	"context"
	"fmt"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotstockfilter"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/push"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/pkg/sys"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

// Client is the main client type for connecting to Futu OpenD.
// It wraps the internal client to provide a public API.
type Client struct {
	inner  *futuapi.Client
	trdEnv int32 // trading environment: 0=real, 1=simulate (default)
	trdMkt int32 // default trading market (0 = auto-detect per request)
}

// New creates a new client with optional configuration.
func New(opts ...Option) *Client {
	futuOpts := make([]futuapi.Option, len(opts))
	for i, o := range opts {
		futuOpts[i] = o
	}
	return &Client{
		inner:  futuapi.New(futuOpts...),
		trdEnv: 1, // default to simulate for safety
	}
}

// GetTradeEnv returns the current trading environment (0=real, 1=simulate).
func (c *Client) GetTradeEnv() int32 {
	return c.trdEnv
}

// WithTradeEnv returns a Client that uses the given trading environment.
// trdEnv: 0 = real trading, 1 = simulate trading.
// This is a client-scoped setting so all trading calls inherit it.
func (c *Client) WithTradeEnv(trdEnv int32) *Client {
	clone := *c
	clone.trdEnv = trdEnv
	return &clone
}

// WithTradeMarket returns a Client that uses the given default trading market.
// When 0, the market is auto-detected per request (current behavior).
func (c *Client) WithTradeMarket(trdMkt int32) *Client {
	clone := *c
	clone.trdMkt = trdMkt
	return &clone
}

// Connect connects to the Futu OpenD server at the given address.
func (c *Client) Connect(addr string) error {
	return c.inner.Connect(addr)
}

// ConnectAddr is an alias for Connect.
func (c *Client) ConnectAddr(addr string) error {
	return c.inner.Connect(addr)
}

// Close closes the connection to OpenD.
func (c *Client) Close() {
	c.inner.Close()
}

// GetConnID returns the connection ID assigned by OpenD.
func (c *Client) GetConnID() uint64 {
	return c.inner.GetConnID()
}

// GetServerVer returns the OpenD server version.
func (c *Client) GetServerVer() int32 {
	return c.inner.GetServerVer()
}

// EnsureConnected returns an error if the client is not connected.
func (c *Client) EnsureConnected() error {
	return c.inner.EnsureConnected()
}

// WithContext returns a client with the given context.
func (c *Client) WithContext(ctx context.Context) *Client {
	return &Client{inner: c.inner.WithContext(ctx)}
}

// Context returns the client's context.
func (c *Client) Context() context.Context {
	return c.inner.Context()
}

// RegisterHandler registers a handler for push notifications.
func (c *Client) RegisterHandler(protoID uint32, h func(protoID uint32, body []byte)) {
	c.inner.RegisterHandler(protoID, h)
}

// GetConn returns the underlying connection (for advanced use).
func (c *Client) GetConn() *futuapi.Conn {
	return c.inner.Conn()
}

// GetQuote retrieves the current quote for a security.
func GetQuote(c *Client, market int32, code string) (*Quote, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	quotes, err := qot.GetBasicQot(c.inner, []*qotcommon.Security{sec})
	if err != nil {
		return nil, err
	}
	if len(quotes) == 0 {
		return nil, fmt.Errorf("no quote returned for %s", code)
	}

	q := quotes[0]
	return &Quote{
		Symbol:       code,
		Market:       market,
		Price:        q.CurPrice,
		Open:         q.OpenPrice,
		High:         q.HighPrice,
		Low:          q.LowPrice,
		Volume:       q.Volume,
		Timestamp:    q.UpdateTime,
		Name:         q.Name,
		LastClose:    q.LastClosePrice,
		Turnover:     q.Turnover,
		TurnoverRate: q.TurnoverRate,
		Amplitude:    q.Amplitude,
	}, nil
}

// GetKLines retrieves K-line (candlestick) data.
func GetKLines(c *Client, market int32, code string, klType int32, num int) ([]KLine, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetKL(c.inner, &qot.GetKLRequest{
		Security:  sec,
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    klType,
		ReqNum:    int32(num),
	})
	if err != nil {
		return nil, err
	}

	klines := make([]KLine, len(resp.KLList))
	for i, kl := range resp.KLList {
		klines[i] = KLine{
			Time:       kl.Time,
			Open:       kl.OpenPrice,
			High:       kl.HighPrice,
			Low:        kl.LowPrice,
			Close:      kl.ClosePrice,
			Volume:     kl.Volume,
			LastClose:  kl.LastClosePrice,
			Turnover:   kl.Turnover,
			ChangeRate: kl.ChangeRate,
			Timestamp:  kl.Timestamp,
		}
	}
	return klines, nil
}

// Subscribe subscribes to real-time market data.
func Subscribe(c *Client, market int32, code string, subTypes []int32) error {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	_, err := qot.Subscribe(c.inner, &qot.SubscribeRequest{
		SecurityList:     []*qotcommon.Security{sec},
		SubTypeList:      subTypesConverted,
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
		IsFirstPush:      true,
	})
	return err
}

// Unsubscribe unsubscribes from real-time market data.
func Unsubscribe(c *Client, market int32, code string, subTypes []int32) error {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	_, err := qot.Subscribe(c.inner, &qot.SubscribeRequest{
		SecurityList:     []*qotcommon.Security{sec},
		SubTypeList:      subTypesConverted,
		IsSubOrUnSub:     false,
		IsRegOrUnRegPush: false,
	})
	return err
}

// UnsubscribeAll unsubscribes from all market data.
func UnsubscribeAll(c *Client) error {
	_, err := qot.Subscribe(c.inner, &qot.SubscribeRequest{
		SubTypeList:  []qot.SubType{},
		IsSubOrUnSub: false,
		IsUnsubAll:   true,
	})
	return err
}

// QuerySubscription queries the current subscription status.
func QuerySubscription(c *Client) (*qot.GetSubInfoResponse, error) {
	return qot.GetSubInfo(c.inner)
}

// RegQotPush registers or unregisters real-time push notifications for a security.
func RegQotPush(c *Client, market int32, code string, subTypes []int32, rehabTypes []int32, isReg bool, isFirstPush bool) error {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	_, err := qot.RegQotPush(c.inner, &qot.RegQotPushRequest{
		SecurityList:  []*qotcommon.Security{sec},
		SubTypeList:   subTypes,
		RehabTypeList: rehabTypes,
		IsRegOrUnReg:  isReg,
		IsFirstPush:   isFirstPush,
	})
	return err
}

// GetAccountList retrieves the list of trading accounts.
func GetAccountList(c *Client) ([]Account, error) {
	resp, err := trd.GetAccList(c.inner, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
	if err != nil {
		return nil, err
	}

	accounts := make([]Account, len(resp.AccList))
	for i, acc := range resp.AccList {
		accounts[i] = Account{
			AccID:             acc.AccID,
			AccType:           acc.AccType,
			TrdEnv:            acc.TrdEnv,
			CardNum:           acc.CardNum,
			AccStatus:         acc.AccStatus,
			TrdMarketAuthList: acc.TrdMarketAuthList,
			SecurityFirm:      acc.SecurityFirm,
			SimAccType:        acc.SimAccType,
			UniCardNum:        acc.UniCardNum,
			AccRole:           acc.AccRole,
			JpAccType:         acc.JpAccType,
		}
	}
	return accounts, nil
}

// UnlockTrading unlocks trading with the given password (MD5 hash).
func UnlockTrading(c *Client, pwdMD5 string) error {
	return trd.UnlockTrade(c.inner, &trd.UnlockTradeRequest{
		Unlock: true,
		PwdMD5: pwdMD5,
	})
}

// PlaceOrder places a trading order.
func PlaceOrder(c *Client, accID uint64, market int32, code string, side, orderType int32, price float64, qty float64) (*PlaceOrderResult, error) {
	resp, err := trd.PlaceOrder(c.inner, &trd.PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    c.trdEnv,
		Code:      code,
		TrdSide:   side,
		OrderType: orderType,
		Price:     price,
		Qty:       qty,
	})
	if err != nil {
		return nil, err
	}
	return &PlaceOrderResult{OrderID: resp.OrderID, OrderIDEx: resp.OrderIDEx}, nil
}

// ModifyOrder modifies or cancels an existing order.
func ModifyOrder(c *Client, accID uint64, market int32, orderID uint64, modifyOp int32, price float64, qty float64) (*trd.ModifyOrderResponse, error) {
	return trd.ModifyOrder(c.inner, &trd.ModifyOrderRequest{
		AccID:         accID,
		TrdMarket:     market,
		TrdEnv:        c.trdEnv,
		OrderID:       orderID,
		ModifyOrderOp: modifyOp,
		Price:         price,
		Qty:           qty,
	})
}

// CancelAllOrder cancels all pending orders for the specified account and market.
// Note: Simulate trading and HKCC accounts do not support CancelAllOrder.
func CancelAllOrder(c *Client, accID uint64, market int32, trdEnv int32) error {
	_, err := trd.ModifyOrder(c.inner, &trd.ModifyOrderRequest{
		AccID:         accID,
		TrdMarket:     market,
		TrdEnv:        trdEnv,
		OrderID:       0,
		ModifyOrderOp: 1, // Cancel
		Price:         0,
		Qty:           0,
		ForAll:        true,
	})
	return err
}

// GetPositionList retrieves the current positions.
func GetPositionList(c *Client, accID uint64) ([]Position, error) {
	resp, err := trd.GetPositionList(c.inner, &trd.GetPositionListRequest{
		AccID:     accID,
		TrdMarket: 0,
		TrdEnv:    c.trdEnv,
	})
	if err != nil {
		return nil, err
	}

	positions := make([]Position, len(resp.PositionList))
	for i, p := range resp.PositionList {
		positions[i] = Position{
			PositionID:       p.PositionID,
			Code:             p.Code,
			Name:             p.Name,
			Market:           p.TrdMarket,
			Quantity:         p.Qty,
			CanSellQty:       p.CanSellQty,
			CostPrice:        p.CostPrice,
			CurPrice:         p.Price,
			MarketVal:        p.Val,
			PnL:              p.PlVal,
			PnLRate:          p.PlRatio,
			TodayBuyQty:      p.TdBuyQty,
			TodayBuyVal:      p.TdBuyVal,
			TodaySellQty:     p.TdSellQty,
			TodaySellVal:     p.TdSellVal,
			TodayPnL:         p.TdPlVal,
			UnrealizedPL:     p.UnrealizedPL,
			RealizedPL:       p.RealizedPL,
			Currency:         p.Currency,
			TrdMarket:        p.TrdMarket,
			DilutedCostPrice: p.DilutedCostPrice,
			AverageCostPrice: p.AverageCostPrice,
			AveragePnLRate:   p.AveragePlRatio,
		}
	}
	return positions, nil
}

// GetFunds retrieves account funds.
func GetFunds(c *Client, accID uint64) (*Funds, error) {
	resp, err := trd.GetFunds(c.inner, &trd.GetFundsRequest{AccID: accID})
	if err != nil {
		return nil, err
	}
	f := resp.Funds
	return &Funds{
		Power:             f.Power,
		TotalAssets:       f.TotalAssets,
		Cash:              f.Cash,
		MarketVal:         f.MarketVal,
		FrozenCash:        f.FrozenCash,
		DebtCash:          f.DebtCash,
		AvlWithdrawalCash: f.AvlWithdrawalCash,
		Currency:          f.Currency,
		AvailableFunds:    f.AvailableFunds,
		UnrealizedPL:      f.UnrealizedPL,
		RealizedPL:        f.RealizedPL,
		RiskLevel:         f.RiskLevel,
		InitialMargin:     f.InitialMargin,
		MaintenanceMargin: f.MaintenanceMargin,
		MaxPowerShort:     f.MaxPowerShort,
		NetCashPower:      f.NetCashPower,
		LongMv:            f.LongMv,
		ShortMv:           f.ShortMv,
		PendingAsset:      f.PendingAsset,
		MaxWithdrawal:     f.MaxWithdrawal,
		RiskStatus:        f.RiskStatus,
	}, nil
}

// MaxTrdQtysInfo represents maximum tradable quantities.
type MaxTrdQtysInfo struct {
	MaxCashBuy          float64
	MaxCashAndMarginBuy float64
	MaxPositionSell     float64
	MaxSellShort        float64
	MaxBuyBack          float64
}

// GetMaxTrdQtys retrieves maximum tradable quantities.
func GetMaxTrdQtys(c *Client, accID uint64, market int32, code string, orderType int32, price float64) (*MaxTrdQtysInfo, error) {
	resp, err := trd.GetMaxTrdQtys(c.inner, &trd.GetMaxTrdQtysRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    c.trdEnv,
		Code:      code,
		OrderType: orderType,
		Price:     price,
	})
	if err != nil {
		return nil, err
	}
	m := resp.MaxTrdQtys
	return &MaxTrdQtysInfo{
		MaxCashBuy:          m.MaxCashBuy,
		MaxCashAndMarginBuy: m.MaxCashAndMarginBuy,
		MaxPositionSell:     m.MaxPositionSell,
		MaxSellShort:        m.MaxSellShort,
		MaxBuyBack:          m.MaxBuyBack,
	}, nil
}

// OrderFeeInfo represents fee information for an order.
type OrderFeeInfo struct {
	OrderIDEx string
	FeeAmount float64
	FeeList   []OrderFeeItemInfo
}

type OrderFeeItemInfo struct {
	Title string
	Value float64
}

// GetOrderFee retrieves order fee information.
func GetOrderFee(c *Client, accID uint64, market int32, orderIDExList []string) ([]*OrderFeeInfo, error) {
	resp, err := trd.GetOrderFee(c.inner, &trd.GetOrderFeeRequest{
		AccID:         accID,
		TrdMarket:     market,
		TrdEnv:        c.trdEnv,
		OrderIDExList: orderIDExList,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*OrderFeeInfo, 0, len(resp.OrderFeeList))
	for _, f := range resp.OrderFeeList {
		if f == nil {
			continue
		}
		feeList := make([]OrderFeeItemInfo, 0, len(f.FeeList))
		for _, item := range f.FeeList {
			feeList = append(feeList, OrderFeeItemInfo{Title: item.Title, Value: item.Value})
		}
		result = append(result, &OrderFeeInfo{
			OrderIDEx: f.OrderIDEx,
			FeeAmount: f.FeeAmount,
			FeeList:   feeList,
		})
	}
	return result, nil
}

// MarginRatioInfo represents margin ratio for a security.
type MarginRatioInfo struct {
	Security      *qotcommon.Security
	IsLongPermit  bool
	IsShortPermit bool
	ShortFeeRate  float64
	ImLongRatio   float64
	ImShortRatio  float64
}

// GetMarginRatio retrieves margin ratio for securities.
func GetMarginRatio(c *Client, accID uint64, market int32, securities []*qotcommon.Security) ([]*MarginRatioInfo, error) {
	resp, err := trd.GetMarginRatio(c.inner, &trd.GetMarginRatioRequest{
		AccID:        accID,
		TrdMarket:    market,
		TrdEnv:       c.trdEnv,
		SecurityList: securities,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*MarginRatioInfo, 0, len(resp.MarginRatioInfoList))
	for _, m := range resp.MarginRatioInfoList {
		if m == nil {
			continue
		}
		result = append(result, &MarginRatioInfo{
			Security:      m.Security,
			IsLongPermit:  m.IsLongPermit,
			IsShortPermit: m.IsShortPermit,
			ShortFeeRate:  m.ShortFeeRate,
			ImLongRatio:   m.ImLongRatio,
			ImShortRatio:  m.ImShortRatio,
		})
	}
	return result, nil
}

// GetOrderList retrieves active orders.
func GetOrderList(c *Client, accID uint64) ([]Order, error) {
	resp, err := trd.GetOrderList(c.inner, &trd.GetOrderListRequest{
		AccID:     accID,
		TrdMarket: 0,
		TrdEnv:    c.trdEnv,
	})
	if err != nil {
		return nil, err
	}

	orders := make([]Order, len(resp.OrderList))
	for i, o := range resp.OrderList {
		orders[i] = Order{
			OrderID:         o.OrderID,
			OrderIDEx:       o.OrderIDEx,
			Code:            o.Code,
			Name:            o.Name,
			TrdSide:         o.TrdSide,
			OrderType:       o.OrderType,
			OrderStatus:     o.OrderStatus,
			Price:           o.Price,
			Qty:             o.Qty,
			FillQty:         o.FillQty,
			FillAvgPrice:    o.FillAvgPrice,
			CreateTime:      o.CreateTime,
			UpdateTime:      o.UpdateTime,
			LastErrMsg:      o.LastErrMsg,
			SecMarket:       o.SecMarket,
			CreateTimestamp: o.CreateTimestamp,
			UpdateTimestamp: o.UpdateTimestamp,
			Remark:          o.Remark,
			TimeInForce:     o.TimeInForce,
			FillOutsideRTH:  o.FillOutsideRTH,
			AuxPrice:        o.AuxPrice,
			TrailType:       o.TrailType,
			TrailValue:      o.TrailValue,
			TrailSpread:     o.TrailSpread,
			Currency:        o.Currency,
			TrdMarket:       o.TrdMarket,
			Session:         o.Session,
		}
	}
	return orders, nil
}

// GetHistoryOrderList retrieves historical orders.
func GetHistoryOrderList(c *Client, accID uint64, market int32, startDate, endDate string) ([]Order, error) {
	resp, err := trd.GetHistoryOrderList(c.inner, &trd.GetHistoryOrderListRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    c.trdEnv,
	})
	if err != nil {
		return nil, err
	}

	orders := make([]Order, 0)
	for _, o := range resp.OrderList {
		if o == nil {
			continue
		}
		orders = append(orders, Order{
			OrderID:         getUint64(o.OrderID),
			OrderIDEx:       getStr(o.OrderIDEx),
			Code:            getStr(o.Code),
			Name:            getStr(o.Name),
			TrdSide:         getInt32(o.TrdSide),
			OrderType:       getInt32(o.OrderType),
			OrderStatus:     getInt32(o.OrderStatus),
			Price:           getFloat64(o.Price),
			Qty:             getFloat64(o.Qty),
			FillQty:         getFloat64(o.FillQty),
			FillAvgPrice:    getFloat64(o.FillAvgPrice),
			CreateTime:      getStr(o.CreateTime),
			UpdateTime:      getStr(o.UpdateTime),
			LastErrMsg:      getStr(o.LastErrMsg),
			SecMarket:       getInt32(o.SecMarket),
			CreateTimestamp: getFloat64(o.CreateTimestamp),
			UpdateTimestamp: getFloat64(o.UpdateTimestamp),
			Remark:          getStr(o.Remark),
			TimeInForce:     getInt32(o.TimeInForce),
			FillOutsideRTH:  getBool(o.FillOutsideRTH),
			AuxPrice:        getFloat64(o.AuxPrice),
			TrailType:       getInt32(o.TrailType),
			TrailValue:      getFloat64(o.TrailValue),
			TrailSpread:     getFloat64(o.TrailSpread),
			Currency:        getInt32(o.Currency),
			TrdMarket:       getInt32(o.TrdMarket),
			Session:         getInt32(o.Session),
		})
	}
	return orders, nil
}

// GetOrderFillList retrieves order fills (executions).
func GetOrderFillList(c *Client, accID uint64) ([]OrderFill, error) {
	resp, err := trd.GetOrderFillList(c.inner, &trd.GetOrderFillListRequest{
		AccID:     accID,
		TrdMarket: 0,
		TrdEnv:    c.trdEnv,
	})
	if err != nil {
		return nil, err
	}

	fills := make([]OrderFill, len(resp.OrderFillList))
	for i, f := range resp.OrderFillList {
		fills[i] = OrderFill{
			FillID:            f.FillID,
			FillIDEx:          f.FillIDEx,
			OrderID:           f.OrderID,
			OrderIDEx:         f.OrderIDEx,
			Code:              f.Code,
			Name:              f.Name,
			TrdSide:           f.TrdSide,
			Price:             f.Price,
			Qty:               f.Qty,
			CreateTime:        f.CreateTime,
			CounterBrokerID:   f.CounterBrokerID,
			CounterBrokerName: f.CounterBrokerName,
			SecMarket:         f.SecMarket,
			CreateTimestamp:   f.CreateTimestamp,
			UpdateTimestamp:   f.UpdateTimestamp,
			Status:            f.Status,
			TrdMarket:         f.TrdMarket,
			JpAccType:         f.JpAccType,
		}
	}
	return fills, nil
}

// GetHistoryOrderFillList retrieves historical order fills.
func GetHistoryOrderFillList(c *Client, accID uint64, market int32) ([]OrderFill, error) {
	resp, err := trd.GetHistoryOrderFillList(c.inner, &trd.GetHistoryOrderFillListRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    c.trdEnv,
	})
	if err != nil {
		return nil, err
	}

	fills := make([]OrderFill, 0, len(resp.OrderFillList))
	for _, f := range resp.OrderFillList {
		if f == nil {
			continue
		}
		fills = append(fills, OrderFill{
			FillID:            f.FillID,
			FillIDEx:          f.FillIDEx,
			OrderID:           f.OrderID,
			OrderIDEx:         f.OrderIDEx,
			Code:              f.Code,
			Name:              f.Name,
			TrdSide:           f.TrdSide,
			Price:             f.Price,
			Qty:               f.Qty,
			CreateTime:        f.CreateTime,
			CounterBrokerID:   f.CounterBrokerID,
			CounterBrokerName: f.CounterBrokerName,
			SecMarket:         f.SecMarket,
			CreateTimestamp:   f.CreateTimestamp,
			UpdateTimestamp:   f.UpdateTimestamp,
			Status:            f.Status,
			TrdMarket:         f.TrdMarket,
			JpAccType:         f.JpAccType,
		})
	}
	return fills, nil
}

// GetOrderBook retrieves order book data.
func GetOrderBook(c *Client, market int32, code string, num int) (*OrderBook, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetOrderBook(c.inner, &qot.GetOrderBookRequest{
		Security: sec,
		Num:      int32(num),
	})
	if err != nil {
		return nil, err
	}

	book := &OrderBook{
		Bids:                    make([]OrderBookItem, len(resp.OrderBookBidList)),
		Asks:                    make([]OrderBookItem, len(resp.OrderBookAskList)),
		SvrRecvTimeBid:          resp.SvrRecvTimeBid,
		SvrRecvTimeBidTimestamp: resp.SvrRecvTimeBidTimestamp,
		SvrRecvTimeAsk:          resp.SvrRecvTimeAsk,
		SvrRecvTimeAskTimestamp: resp.SvrRecvTimeAskTimestamp,
	}
	for i, b := range resp.OrderBookBidList {
		details := make([]OrderBookDetail, 0, len(b.DetailList))
		for _, d := range b.DetailList {
			details = append(details, OrderBookDetail{OrderID: d.OrderID, Volume: d.Volume})
		}
		book.Bids[i] = OrderBookItem{Price: b.Price, Volume: b.Volume, OrderCount: b.OrderCount, DetailList: details}
	}
	for i, a := range resp.OrderBookAskList {
		details := make([]OrderBookDetail, 0, len(a.DetailList))
		for _, d := range a.DetailList {
			details = append(details, OrderBookDetail{OrderID: d.OrderID, Volume: d.Volume})
		}
		book.Asks[i] = OrderBookItem{Price: a.Price, Volume: a.Volume, OrderCount: a.OrderCount, DetailList: details}
	}
	return book, nil
}

// GetTicker retrieves ticker data.
func GetTicker(c *Client, market int32, code string, num int) ([]Ticker, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetTicker(c.inner, &qot.GetTickerRequest{
		Security: sec,
		Num:      int32(num),
	})
	if err != nil {
		return nil, err
	}

	tickers := make([]Ticker, len(resp.TickerList))
	for i, t := range resp.TickerList {
		dir := "N/A"
		switch t.Dir {
		case 1:
			dir = "Buy"
		case 2:
			dir = "Sell"
		}
		tickers[i] = Ticker{
			Time:      t.Time,
			Sequence:  t.Sequence,
			Price:     t.Price,
			Volume:    t.Volume,
			Direction: dir,
			Turnover:  t.Turnover,
			RecvTime:  t.RecvTime,
			Type:      t.Type,
			TypeSign:  t.TypeSign,
			Timestamp: t.Timestamp,
		}
	}
	return tickers, nil
}

// GetRT retrieves real-time data.
func GetRT(c *Client, market int32, code string) ([]RT, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetRT(c.inner, &qot.GetRTRequest{Security: sec})
	if err != nil {
		return nil, err
	}

	rtData := make([]RT, len(resp.RTList))
	for i, r := range resp.RTList {
		rtData[i] = RT{
			Time:      r.Time,
			Price:     r.Price,
			Volume:    r.Volume,
			LastClose: r.LastClosePrice,
			AvgPrice:  r.AvgPrice,
			Turnover:  r.Turnover,
		}
	}
	return rtData, nil
}

// GetBroker retrieves broker data.
func GetBroker(c *Client, market int32, code string, num int) ([]Broker, []Broker, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetBroker(c.inner, &qot.GetBrokerRequest{
		Security: sec,
		Num:      int32(num),
	})
	if err != nil {
		return nil, nil, err
	}

	bidBrokers := make([]Broker, len(resp.BidBrokerList))
	for i, b := range resp.BidBrokerList {
		bidBrokers[i] = Broker{ID: b.ID, Name: b.Name, Pos: b.Pos, Volume: b.Volume}
	}
	askBrokers := make([]Broker, len(resp.AskBrokerList))
	for i, a := range resp.AskBrokerList {
		askBrokers[i] = Broker{ID: a.ID, Name: a.Name, Pos: a.Pos, Volume: a.Volume}
	}
	return bidBrokers, askBrokers, nil
}

// GetStaticInfo retrieves static security info.
func GetStaticInfo(c *Client, market int32, code string) ([]StaticInfo, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetStaticInfo(c.inner, &qot.GetStaticInfoRequest{
		Market:       market,
		SecurityList: []*qotcommon.Security{sec},
	})
	if err != nil {
		return nil, err
	}

	infos := make([]StaticInfo, len(resp.StaticInfoList))
	for i, s := range resp.StaticInfoList {
		var name string
		var secType int32
		var listTime string
		var lotSize int32
		if s.Basic != nil {
			if s.Basic.Name != nil {
				name = *s.Basic.Name
			}
			if s.Basic.SecType != nil {
				secType = *s.Basic.SecType
			}
			if s.Basic.ListTime != nil {
				listTime = *s.Basic.ListTime
			}
			if s.Basic.LotSize != nil {
				lotSize = *s.Basic.LotSize
			}
		}
		infos[i] = StaticInfo{Code: code, Name: name, Type: secType, ListTime: listTime, LotSize: lotSize}
	}
	return infos, nil
}

// GetTradeDate retrieves trade dates.
func GetTradeDate(c *Client, market int32, startDate, endDate string) ([]string, error) {
	resp, err := qot.GetTradeDate(c.inner, &qot.GetTradeDateRequest{
		Market:    market,
		BeginTime: startDate,
		EndTime:   endDate,
	})
	if err != nil {
		return nil, err
	}

	dates := make([]string, len(resp.TradeDateList))
	for i, td := range resp.TradeDateList {
		if td.Time != nil {
			dates[i] = *td.Time
		}
	}
	return dates, nil
}

// GetFutureInfo retrieves futures information.
func GetFutureInfo(c *Client, code string) ([]FutureInfo, error) {
	marketPtr := int32(2) // HK Future
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetFutureInfo(c.inner, &qot.GetFutureInfoRequest{
		SecurityList: []*qotcommon.Security{sec},
	})
	if err != nil {
		return nil, err
	}

	infos := make([]FutureInfo, len(resp.FutureInfoList))
	for i, f := range resp.FutureInfoList {
		secCode := ""
		if f.Security != nil && f.Security.Code != nil {
			secCode = *f.Security.Code
		}
		ownerCode := ""
		if f.Owner != nil && f.Owner.Code != nil {
			ownerCode = *f.Owner.Code
		}
		infos[i] = FutureInfo{
			Code:               secCode,
			Name:               f.Name,
			Expire:             f.LastTradeTime,
			LastTradeTimestamp: f.LastTradeTimestamp,
			Owner:              ownerCode,
			OwnerOther:         f.OwnerOther,
			Exchange:           f.Exchange,
			ContractType:       f.ContractType,
			ContractSize:       f.ContractSize,
			ContractSizeUnit:   f.ContractSizeUnit,
			QuoteCurrency:      f.QuoteCurrency,
			MinVar:             f.MinVar,
			MinVarUnit:         f.MinVarUnit,
			QuoteUnit:          f.QuoteUnit,
			TimeZone:           f.TimeZone,
			ExchangeFormatUrl:  f.ExchangeFormatUrl,
		}
	}
	return infos, nil
}

// GetPlateSet retrieves plate set (板块) list.
func GetPlateSet(c *Client, market int32) ([]Plate, error) {
	resp, err := qot.GetPlateSet(c.inner, &qot.GetPlateSetRequest{Market: market})
	if err != nil {
		return nil, err
	}

	plates := make([]Plate, len(resp.PlateSetList))
	for i, p := range resp.PlateSetList {
		code := ""
		if p.Plate != nil && p.Plate.Code != nil {
			code = *p.Plate.Code
		}
		plates[i] = Plate{Code: code, Name: p.Name}
	}
	return plates, nil
}

// GetIpoList retrieves IPO list.
func GetIpoList(c *Client, market int32) ([]IpoData, error) {
	resp, err := qot.GetIpoList(c.inner, &qot.GetIpoListRequest{Market: market})
	if err != nil {
		return nil, err
	}

	ipos := make([]IpoData, 0)
	for _, ip := range resp.IpoList {
		if ip.Basic == nil {
			continue
		}
		code := ""
		if ip.Basic.Security != nil && ip.Basic.Security.Code != nil {
			code = *ip.Basic.Security.Code
		}
		ipos = append(ipos, IpoData{
			Code:          code,
			Name:          ip.Basic.Name,
			ListDate:      ip.Basic.ListTime,
			ListTimestamp: ip.Basic.ListTimestamp,
		})
	}
	return ipos, nil
}

// GetUserSecurityGroup retrieves user security group list.
func GetUserSecurityGroup(c *Client) ([]UserSecurityGroup, error) {
	resp, err := qot.GetUserSecurityGroup(c.inner, &qot.GetUserSecurityGroupRequest{})
	if err != nil {
		return nil, err
	}

	groups := make([]UserSecurityGroup, 0)
	for _, g := range resp.GroupList {
		groups = append(groups, UserSecurityGroup{Name: g.GroupName, GroupType: g.GroupType})
	}
	return groups, nil
}

// GetUserSecurity retrieves user security list by group name.
func GetUserSecurity(c *Client, groupName string) ([]StaticInfo, error) {
	resp, err := qot.GetUserSecurity(c.inner, groupName)
	if err != nil {
		return nil, err
	}

	infos := make([]StaticInfo, 0)
	for _, s := range resp.StaticInfoList {
		if s == nil || s.Basic == nil {
			continue
		}
		code := ""
		if s.Basic.Security != nil && s.Basic.Security.Code != nil {
			code = *s.Basic.Security.Code
		}
		name := ""
		if s.Basic.Name != nil {
			name = *s.Basic.Name
		}
		secType := int32(0)
		if s.Basic.SecType != nil {
			secType = *s.Basic.SecType
		}
		infos = append(infos, StaticInfo{
			Code: code,
			Name: name,
			Type: secType,
		})
	}
	return infos, nil
}

// GetMarketState retrieves market state (trading status).
func GetMarketState(c *Client, market int32, code string) (int32, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetMarketState(c.inner, &qot.GetMarketStateRequest{
		SecurityList: []*qotcommon.Security{sec},
	})
	if err != nil {
		return 0, err
	}

	if len(resp.MarketInfoList) == 0 {
		return 0, nil
	}

	return resp.MarketInfoList[0].MarketState, nil
}

// CapitalFlow represents capital flow data.
type CapitalFlow struct {
	Time        string
	InFlow      float64
	MainInFlow  float64
	SuperInFlow float64
	BigInFlow   float64
	MidInFlow   float64
	SmlInFlow   float64
	Timestamp   float64
}

// GetCapitalFlow retrieves capital flow data.
func GetCapitalFlow(c *Client, market int32, code string) ([]CapitalFlow, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetCapitalFlow(c.inner, &qot.GetCapitalFlowRequest{
		Security: sec,
	})
	if err != nil {
		return nil, err
	}

	flows := make([]CapitalFlow, 0)
	for _, f := range resp.FlowItemList {
		flows = append(flows, CapitalFlow{
			Time:        f.Time,
			InFlow:      f.InFlow,
			MainInFlow:  f.MainInFlow,
			SuperInFlow: f.SuperInFlow,
			BigInFlow:   f.BigInFlow,
			MidInFlow:   f.MidInFlow,
			SmlInFlow:   f.SmlInFlow,
			Timestamp:   f.Timestamp,
		})
	}
	return flows, nil
}

// GetCapitalDistribution retrieves capital distribution.
func GetCapitalDistribution(c *Client, market int32, code string) (*CapitalDistribution, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetCapitalDistribution(c.inner, sec)
	if err != nil {
		return nil, err
	}

	if resp.CapitalDistribution == nil {
		return nil, nil
	}

	cd := resp.CapitalDistribution
	return &CapitalDistribution{
		MainInflow:      cd.CapitalInSuper,
		BigInflow:       cd.CapitalInBig,
		MidInflow:       cd.CapitalInMid,
		SmallInflow:     cd.CapitalInSmall,
		MainOutflow:     cd.CapitalOutSuper,
		BigOutflow:      cd.CapitalOutBig,
		MidOutflow:      cd.CapitalOutMid,
		SmallOutflow:    cd.CapitalOutSmall,
		UpdateTime:      cd.UpdateTime,
		UpdateTimestamp: cd.UpdateTimestamp,
	}, nil
}

// GetOwnerPlate retrieves owner plates.
func GetOwnerPlate(c *Client, market int32, code string) ([]string, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetOwnerPlate(c.inner, &qot.GetOwnerPlateRequest{
		SecurityList: []*qotcommon.Security{sec},
	})
	if err != nil {
		return nil, err
	}

	plates := make([]string, 0)
	for _, p := range resp.OwnerPlateList {
		for _, pi := range p.PlateInfoList {
			if pi.Name != nil {
				plates = append(plates, *pi.Name)
			}
		}
	}
	return plates, nil
}

// History KL pagination constants.
const (
	DefaultHistoryKLPageSize = 1000                   // max K-lines per API page (protocol limit)
	DefaultHistoryKLDelay    = 200 * time.Millisecond // delay between pagination pages
)

// HistoryKLPaginationDelay controls the wait time between pages when
// fetching historical K-lines. Adjust this to respect your OpenD rate limits.
var HistoryKLPaginationDelay = DefaultHistoryKLDelay

// RequestHistoryKL requests historical K-line data with automatic pagination.
// It fetches all available K-lines between startDate and endDate (inclusive),
// automatically handling page boundaries via NextReqKey.
func RequestHistoryKL(c *Client, market int32, code string, klType int32, startDate, endDate string) ([]KLine, error) {
	return RequestHistoryKLWithLimit(c, market, code, klType, startDate, endDate, DefaultHistoryKLPageSize)
}

// RequestHistoryKLWithLimit requests historical K-line data with a configurable
// page size. It automatically paginates until all data is retrieved.
// maxPerPage controls how many K-lines are requested per API call (max 1000).
// Uses HistoryKLPaginationDelay between pages.
func RequestHistoryKLWithLimit(c *Client, market int32, code string, klType int32, startDate, endDate string, maxPerPage int32) ([]KLine, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	var allKLines []KLine
	var nextReqKey []byte

	for {
		resp, err := qot.RequestHistoryKL(c.inner, &qot.RequestHistoryKLRequest{
			Security:    sec,
			KlType:      klType,
			BeginTime:   startDate,
			EndTime:     endDate,
			MaxAckKLNum: maxPerPage,
			NextReqKey:  nextReqKey,
		})
		if err != nil {
			return allKLines, err
		}

		for _, kl := range resp.KLList {
			allKLines = append(allKLines, KLine{
				Time:       kl.Time,
				Open:       kl.OpenPrice,
				High:       kl.HighPrice,
				Low:        kl.LowPrice,
				Close:      kl.ClosePrice,
				Volume:     kl.Volume,
				LastClose:  kl.LastClosePrice,
				Turnover:   kl.Turnover,
				ChangeRate: kl.ChangeRate,
				Timestamp:  kl.Timestamp,
			})
		}

		// Check if there are more pages
		if len(resp.NextReqKey) == 0 {
			break
		}
		nextReqKey = resp.NextReqKey

		// Rate limit: configurable delay between pages
		time.Sleep(HistoryKLPaginationDelay)
	}

	return allKLines, nil
}

// GetReference retrieves related/reference securities.
func GetReference(c *Client, market int32, code string, refType int32) ([]StaticInfo, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetReference(c.inner, &qot.GetReferenceRequest{
		Security:      sec,
		ReferenceType: refType,
	})
	if err != nil {
		return nil, err
	}

	infos := make([]StaticInfo, 0)
	for _, s := range resp.StaticInfoList {
		if s == nil || s.Basic == nil {
			continue
		}
		code := ""
		if s.Basic.Security != nil && s.Basic.Security.Code != nil {
			code = *s.Basic.Security.Code
		}
		name := ""
		if s.Basic.Name != nil {
			name = *s.Basic.Name
		}
		secType := int32(0)
		if s.Basic.SecType != nil {
			secType = *s.Basic.SecType
		}
		infos = append(infos, StaticInfo{
			Code: code,
			Name: name,
			Type: secType,
		})
	}
	return infos, nil
}

// GetPlateSecurity retrieves securities in a plate.
func GetPlateSecurity(c *Client, market int32, plateCode string) ([]StaticInfo, error) {
	marketPtr := market
	plate := &qotcommon.Security{Market: &marketPtr, Code: &plateCode}

	resp, err := qot.GetPlateSecurity(c.inner, &qot.GetPlateSecurityRequest{Plate: plate})
	if err != nil {
		return nil, err
	}

	infos := make([]StaticInfo, 0)
	for _, s := range resp.StaticInfoList {
		if s == nil || s.Basic == nil {
			continue
		}
		code := ""
		if s.Basic.Security != nil && s.Basic.Security.Code != nil {
			code = *s.Basic.Security.Code
		}
		name := ""
		if s.Basic.Name != nil {
			name = *s.Basic.Name
		}
		secType := int32(0)
		if s.Basic.SecType != nil {
			secType = *s.Basic.SecType
		}
		infos = append(infos, StaticInfo{
			Code: code,
			Name: name,
			Type: secType,
		})
	}
	return infos, nil
}

// GetOptionExpirationDate retrieves option expiration dates.
func GetOptionExpirationDate(c *Client, market int32, code string) ([]OptionExpiration, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetOptionExpirationDate(c.inner, &qot.GetOptionExpirationDateRequest{
		Owner: sec,
	})
	if err != nil {
		return nil, err
	}

	expirations := make([]OptionExpiration, 0)
	for _, e := range resp.DateList {
		if e == nil {
			continue
		}
		expirations = append(expirations, OptionExpiration{
			Date: e.StrikeTime,
			Days: e.OptionExpiryDateDistance,
			Desc: fmt.Sprintf("Cycle %d", e.Cycle),
		})
	}
	return expirations, nil
}

// ModifyUserSecurity adds/removes securities from user group.
func ModifyUserSecurity(c *Client, groupName string, op int32, market int32, codes []string) error {
	securities := make([]*qotcommon.Security, len(codes))
	for i, code := range codes {
		securities[i] = &qotcommon.Security{Market: &market, Code: &code}
	}

	_, err := qot.ModifyUserSecurity(c.inner, &qot.ModifyUserSecurityRequest{
		GroupName:    groupName,
		Op:           op,
		SecurityList: securities,
	})
	return err
}

// GetSubInfo retrieves subscription info.
func GetSubInfo(c *Client) (*SubInfo, error) {
	resp, err := qot.GetSubInfo(c.inner)
	if err != nil {
		return nil, err
	}

	quota := int32(0)
	subTypes := make(map[int32]bool)
	for _, si := range resp.ConnSubInfoList {
		if si != nil {
			if si.UsedQuota != nil {
				quota += *si.UsedQuota
			}
			for _, sub := range si.SubInfoList {
				if sub != nil && sub.SubType != nil {
					subTypes[*sub.SubType] = true
				}
			}
		}
	}

	types := make([]int32, 0, len(subTypes))
	for t := range subTypes {
		types = append(types, t)
	}

	return &SubInfo{
		IsSub:    len(resp.ConnSubInfoList) > 0,
		SubTypes: types,
		Security: fmt.Sprintf("Used: %d, Remain: %d", quota, resp.RemainQuota),
	}, nil
}

// RequestTradeDate requests trade dates for a specific security.
func RequestTradeDate(c *Client, market int32, startDate, endDate string, code string) ([]string, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.RequestTradeDate(c.inner, &qot.RequestTradeDateRequest{
		Market:    market,
		BeginTime: startDate,
		EndTime:   endDate,
		Security:  sec,
	})
	if err != nil {
		return nil, err
	}

	dates := make([]string, 0)
	for _, td := range resp.TradeDateList {
		if td == nil {
			continue
		}
		if td.Time != nil {
			dates = append(dates, *td.Time)
		}
	}
	return dates, nil
}

// OptChainItem represents a pair of call and put options at the same strike price.
type OptChainItem struct {
	Call *qotcommon.SecurityStaticInfo
	Put  *qotcommon.SecurityStaticInfo
}

// OptChain represents the option chain for a single expiration date.
type OptChain struct {
	StrikeTime      string
	StrikeTimestamp float64
	Option          []*OptChainItem
}

// StockFilterResult represents a single stock filter result.
type StockFilterResult struct {
	Security   *qotcommon.Security
	Name       string
	CurPrice   float64
	ChangeRate float64
	Volume     int64
	Turnover   float64
	HighPrice  float64
	LowPrice   float64
}

// StockFilter filters stocks based on basic criteria.
func StockFilter(c *Client, market int32, begin, num int32) ([]*StockFilterResult, error) {
	resp, err := qot.StockFilter(c.inner, &qot.StockFilterRequest{
		Market: market,
		Begin:  begin,
		Num:    num,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*StockFilterResult, 0, len(resp.DataList))
	for _, d := range resp.DataList {
		if d == nil {
			continue
		}
		r := &StockFilterResult{
			Security: d.Security,
			Name:     d.Name,
		}
		for _, base := range d.BaseDataList {
			if base == nil {
				continue
			}
			fieldName := base.GetFieldName()
			value := base.GetValue()
			switch qotstockfilter.StockField(fieldName) {
			case qotstockfilter.StockField_StockField_CurPrice:
				r.CurPrice = value
			case qotstockfilter.StockField_StockField_ChangeRate5min:
				r.ChangeRate = value
			case qotstockfilter.StockField_StockField_VolumeRatio:
				r.Volume = int64(value)
			}
		}
		result = append(result, r)
	}
	return result, nil
}

// GetOptionChain returns the option chain for the given underlying security.
func GetOptionChain(c *Client, market int32, code string, indexOptionType, optType, condition int32, beginTime, endTime string) ([]*OptChain, error) {
	marketPtr := market
	owner := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetOptionChain(c.inner, &qot.GetOptionChainRequest{
		Owner:           owner,
		IndexOptionType: indexOptionType,
		Type:            optType,
		Condition:       condition,
		BeginTime:       beginTime,
		EndTime:         endTime,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*OptChain, 0, len(resp.OptionChain))
	for _, chain := range resp.OptionChain {
		if chain == nil {
			continue
		}
		oc := &OptChain{
			StrikeTime:      chain.StrikeTime,
			StrikeTimestamp: chain.StrikeTimestamp,
			Option:          make([]*OptChainItem, 0, len(chain.Option)),
		}
		for _, opt := range chain.Option {
			if opt == nil {
				continue
			}
			item := &OptChainItem{
				Call: opt.Call,
				Put:  opt.Put,
			}
			oc.Option = append(oc.Option, item)
		}
		result = append(result, oc)
	}
	return result, nil
}

// WarrantData represents warrant data.
type WarrantData struct {
	Stock              *qotcommon.Security
	Owner              *qotcommon.Security
	Type               int32
	Issuer             int32
	MaturityTime       string
	MaturityTimestamp  float64
	ListTime           string
	ListTimestamp      float64
	LastTradeTime      string
	LastTradeTimestamp float64
	RecoveryPrice      float64
	ConversionRatio    float64
	LotSize            int32
	StrikePrice        float64
	LastClosePrice     float64
	Name               string
	CurPrice           float64
	PriceChangeVal     float64
	ChangeRate         float64
	Status             int32
	BidPrice           float64
	AskPrice           float64
	BidVol             int64
	AskVol             int64
	Volume             int64
	Turnover           float64
	Score              float64
	Premium            float64
	BreakEvenPoint     float64
	Leverage           float64
	Ipop               float64
	PriceRecoveryRatio float64
	ConversionPrice    float64
	StreetRate         float64
	StreetVol          int64
	Amplitude          float64
	IssueSize          int64
	HighPrice          float64
	LowPrice           float64
	ImpliedVolatility  float64
	Delta              float64
	EffectiveLeverage  float64
	UpperStrikePrice   float64
	LowerStrikePrice   float64
	InLinePriceStatus  int32
}

// GetWarrant returns the list of warrants for the given underlying security.
func GetWarrant(c *Client, market int32, code string, begin, num int32, sortField int32, ascend bool, optType, issuer, status int32) ([]*WarrantData, error) {
	marketPtr := market
	owner := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetWarrant(c.inner, &qot.GetWarrantRequest{
		Begin:     begin,
		Num:       num,
		SortField: sortField,
		Ascend:    ascend,
		Owner:     owner,
		TypeList:  []int32{optType},
		Status:    status,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*WarrantData, 0, len(resp.WarrantDataList))
	for _, w := range resp.WarrantDataList {
		if w == nil {
			continue
		}
		result = append(result, &WarrantData{
			Stock:              w.Stock,
			Owner:              w.Owner,
			Type:               w.Type,
			Issuer:             w.Issuer,
			MaturityTime:       w.MaturityTime,
			MaturityTimestamp:  w.MaturityTimestamp,
			ListTime:           w.ListTime,
			ListTimestamp:      w.ListTimestamp,
			LastTradeTime:      w.LastTradeTime,
			LastTradeTimestamp: w.LastTradeTimestamp,
			RecoveryPrice:      w.RecoveryPrice,
			ConversionRatio:    w.ConversionRatio,
			LotSize:            w.LotSize,
			StrikePrice:        w.StrikePrice,
			LastClosePrice:     w.LastClosePrice,
			Name:               w.Name,
			CurPrice:           w.CurPrice,
			PriceChangeVal:     w.PriceChangeVal,
			ChangeRate:         w.ChangeRate,
			Status:             w.Status,
			BidPrice:           w.BidPrice,
			AskPrice:           w.AskPrice,
			BidVol:             w.BidVol,
			AskVol:             w.AskVol,
			Volume:             w.Volume,
			Turnover:           w.Turnover,
			Score:              w.Score,
			Premium:            w.Premium,
			BreakEvenPoint:     w.BreakEvenPoint,
			Leverage:           w.Leverage,
			Ipop:               w.Ipop,
			PriceRecoveryRatio: w.PriceRecoveryRatio,
			ConversionPrice:    w.ConversionPrice,
			StreetRate:         w.StreetRate,
			StreetVol:          w.StreetVol,
			Amplitude:          w.Amplitude,
			IssueSize:          w.IssueSize,
			HighPrice:          w.HighPrice,
			LowPrice:           w.LowPrice,
			ImpliedVolatility:  w.ImpliedVolatility,
			Delta:              w.Delta,
			EffectiveLeverage:  w.EffectiveLeverage,
			UpperStrikePrice:   w.UpperStrikePrice,
			LowerStrikePrice:   w.LowerStrikePrice,
			InLinePriceStatus:  w.InLinePriceStatus,
		})
	}
	return result, nil
}

// Snapshot represents security snapshot data.
type Snapshot struct {
	Security                *qotcommon.Security
	Name                    string
	Type                    int32
	IsSuspend               bool
	LotSize                 int32
	CurPrice                float64
	ChangeVal               float64
	HighPrice               float64
	LowPrice                float64
	OpenPrice               float64
	LastClose               float64
	Volume                  int64
	Turnover                float64
	ListTime                string
	PriceSpread             float64
	UpdateTime              string
	TurnoverRate            float64
	ListTimestamp           float64
	UpdateTimestamp         float64
	AskPrice                float64
	BidPrice                float64
	AskVol                  int64
	BidVol                  int64
	EnableMargin            bool
	MortgageRatio           float64
	LongMarginInitialRatio  float64
	EnableShortSell         bool
	ShortSellRate           float64
	ShortAvailableVolume    int64
	ShortMarginInitialRatio float64
	Amplitude               float64
	AvgPrice                float64
	BidAskRatio             float64
	VolumeRatio             float64
	Highest52WeeksPrice     float64
	Lowest52WeeksPrice      float64
	HighestHistoryPrice     float64
	LowestHistoryPrice      float64
	SecStatus               int32
	ClosePrice5Minute       float64
}

// GetSecuritySnapshot returns snapshot data for the given securities.
func GetSecuritySnapshot(c *Client, securities []*qotcommon.Security) ([]*Snapshot, error) {
	resp, err := qot.GetSecuritySnapshot(c.inner, &qot.GetSecuritySnapshotRequest{
		SecurityList: securities,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*Snapshot, 0, len(resp.SnapshotList))
	for _, s := range resp.SnapshotList {
		if s == nil || s.Basic == nil {
			continue
		}
		basic := s.Basic
		result = append(result, &Snapshot{
			Security:                basic.Security,
			Name:                    getStr(basic.Name),
			Type:                    getInt32(basic.Type),
			IsSuspend:               getBool(basic.IsSuspend),
			LotSize:                 getInt32(basic.LotSize),
			CurPrice:                getFloat64(basic.CurPrice),
			ChangeVal:               getFloat64(basic.CurPrice) - getFloat64(basic.LastClosePrice),
			HighPrice:               getFloat64(basic.HighPrice),
			LowPrice:                getFloat64(basic.LowPrice),
			OpenPrice:               getFloat64(basic.OpenPrice),
			LastClose:               getFloat64(basic.LastClosePrice),
			Volume:                  getInt64(basic.Volume),
			Turnover:                getFloat64(basic.Turnover),
			ListTime:                getStr(basic.ListTime),
			PriceSpread:             getFloat64(basic.PriceSpread),
			UpdateTime:              getStr(basic.UpdateTime),
			TurnoverRate:            getFloat64(basic.TurnoverRate),
			ListTimestamp:           getFloat64(basic.ListTimestamp),
			UpdateTimestamp:         getFloat64(basic.UpdateTimestamp),
			AskPrice:                getFloat64(basic.AskPrice),
			BidPrice:                getFloat64(basic.BidPrice),
			AskVol:                  getInt64(basic.AskVol),
			BidVol:                  getInt64(basic.BidVol),
			EnableMargin:            getBool(basic.EnableMargin),
			MortgageRatio:           getFloat64(basic.MortgageRatio),
			LongMarginInitialRatio:  getFloat64(basic.LongMarginInitialRatio),
			EnableShortSell:         getBool(basic.EnableShortSell),
			ShortSellRate:           getFloat64(basic.ShortSellRate),
			ShortAvailableVolume:    getInt64(basic.ShortAvailableVolume),
			ShortMarginInitialRatio: getFloat64(basic.ShortMarginInitialRatio),
			Amplitude:               getFloat64(basic.Amplitude),
			AvgPrice:                getFloat64(basic.AvgPrice),
			BidAskRatio:             getFloat64(basic.BidAskRatio),
			VolumeRatio:             getFloat64(basic.VolumeRatio),
			Highest52WeeksPrice:     getFloat64(basic.Highest52WeeksPrice),
			Lowest52WeeksPrice:      getFloat64(basic.Lowest52WeeksPrice),
			HighestHistoryPrice:     getFloat64(basic.HighestHistoryPrice),
			LowestHistoryPrice:      getFloat64(basic.LowestHistoryPrice),
			SecStatus:               getInt32(basic.SecStatus),
			ClosePrice5Minute:       getFloat64(basic.ClosePrice5Minute),
		})
	}
	return result, nil
}

func getStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func getInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func getBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func getFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func getInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func getUint64(i *uint64) uint64 {
	if i == nil {
		return 0
	}
	return *i
}

// CodeChangeInfo represents information about a code change.
type CodeChangeInfo struct {
	Type            int32
	Security        *qotcommon.Security
	RelatedSecurity *qotcommon.Security
	PublicTime      string
	EffectiveTime   string
}

// GetCodeChange returns code change information for the given securities.
func GetCodeChange(c *Client, securities []*qotcommon.Security) ([]*CodeChangeInfo, error) {
	resp, err := qot.GetCodeChange(c.inner, &qot.GetCodeChangeRequest{
		SecurityList: securities,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*CodeChangeInfo, 0, len(resp.CodeChangeList))
	for _, cc := range resp.CodeChangeList {
		if cc == nil {
			continue
		}
		result = append(result, &CodeChangeInfo{
			Type:            cc.Type,
			Security:        cc.Security,
			RelatedSecurity: cc.RelatedSecurity,
			PublicTime:      cc.PublicTime,
			EffectiveTime:   cc.EffectiveTime,
		})
	}
	return result, nil
}

// GlobalState represents global connection state.
type GlobalState struct {
	ServerVer         int32
	ServerBuildNo     int32
	Time              int64
	LocalTime         float64
	QotLogined        bool
	TrdLogined        bool
	MarketHK          int32
	MarketUS          int32
	MarketSH          int32
	MarketSZ          int32
	MarketHKFuture    int32
	MarketUSFuture    int32
	MarketSGFuture    int32
	MarketJPFuture    int32
	ProgramStatus     int32
	ProgramStatusDesc string
}

// GetGlobalState retrieves global connection state.
func GetGlobalState(c *Client) (*GlobalState, error) {
	resp, err := sys.GetGlobalState(c.inner)
	if err != nil {
		return nil, err
	}

	return &GlobalState{
		ServerVer:      resp.ServerVer,
		ServerBuildNo:  resp.ServerBuildNo,
		Time:           resp.Time,
		LocalTime:      resp.LocalTime,
		QotLogined:     resp.QotLogined,
		TrdLogined:     resp.TrdLogined,
		MarketHK:       resp.MarketHK,
		MarketUS:       resp.MarketUS,
		MarketSH:       resp.MarketSH,
		MarketSZ:       resp.MarketSZ,
		MarketHKFuture: resp.MarketHKFuture,
		MarketUSFuture: resp.MarketUSFuture,
		MarketSGFuture: resp.MarketSGFuture,
		MarketJPFuture: resp.MarketJPFuture,
		ProgramStatus: func() int32 {
			ps := resp.ProgramStatus
			if ps != nil && ps.Type != nil {
				return int32(*ps.Type)
			}
			return 0
		}(),
		ProgramStatusDesc: func() string {
			ps := resp.ProgramStatus
			if ps != nil {
				return ps.GetStrExtDesc()
			}
			return ""
		}(),
	}, nil
}

// UserInfo represents user information.
type UserInfo struct {
	UserID    int64
	NickName  string
	AvatarUrl string
	ApiLevel  string
}

// GetUserInfo retrieves user information.
func GetUserInfo(c *Client) (*UserInfo, error) {
	resp, err := sys.GetUserInfo(c.inner)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		UserID:    resp.UserID,
		NickName:  resp.NickName,
		AvatarUrl: resp.AvatarUrl,
		ApiLevel:  resp.ApiLevel,
	}, nil
}

// DelayStatistics represents delay statistics for Qot push.
type DelayStatistics struct {
	QotPushType    int32
	DelayAvg       float64
	Count          int32
	ItemList       []DelayStatisticsItem
	ReqReplyList   []ReqReplyStatisticsItem
	PlaceOrderList []PlaceOrderStatisticsItem
}

// DelayStatisticsItem represents a single delay statistics item.
type DelayStatisticsItem struct {
	Begin           int32
	End             int32
	Count           int32
	Proportion      float64
	CumulativeRatio float64
}

// ReqReplyStatisticsItem represents request-reply statistics.
type ReqReplyStatisticsItem struct {
	ProtoID      int32
	Count        int32
	TotalCostAvg float64
	OpenDCostAvg float64
	NetDelayAvg  float64
	IsLocalReply bool
}

// PlaceOrderStatisticsItem represents order placement statistics.
type PlaceOrderStatisticsItem struct {
	OrderID    string
	TotalCost  float64
	OpenDCost  float64
	NetDelay   float64
	UpdateCost float64
}

// GetDelayStatistics retrieves delay statistics.
func GetDelayStatistics(c *Client) (*DelayStatistics, error) {
	resp, err := sys.GetDelayStatistics(c.inner)
	if err != nil {
		return nil, err
	}

	if len(resp.QotPushStatisticsList) == 0 {
		return &DelayStatistics{}, nil
	}

	stats := resp.QotPushStatisticsList[0]
	items := make([]DelayStatisticsItem, 0, len(stats.ItemList))
	for _, item := range stats.ItemList {
		items = append(items, DelayStatisticsItem{
			Begin:           item.GetBegin(),
			End:             item.GetEnd(),
			Count:           item.GetCount(),
			Proportion:      float64(item.GetProportion()),
			CumulativeRatio: float64(item.GetCumulativeRatio()),
		})
	}

	reqReplyList := make([]ReqReplyStatisticsItem, 0, len(resp.ReqReplyStatisticsList))
	for _, r := range resp.ReqReplyStatisticsList {
		reqReplyList = append(reqReplyList, ReqReplyStatisticsItem{
			ProtoID:      r.GetProtoID(),
			Count:        r.GetCount(),
			TotalCostAvg: float64(r.GetTotalCostAvg()),
			OpenDCostAvg: float64(r.GetOpenDCostAvg()),
			NetDelayAvg:  float64(r.GetNetDelayAvg()),
			IsLocalReply: r.GetIsLocalReply(),
		})
	}

	placeOrderList := make([]PlaceOrderStatisticsItem, 0, len(resp.PlaceOrderStatisticsList))
	for _, p := range resp.PlaceOrderStatisticsList {
		placeOrderList = append(placeOrderList, PlaceOrderStatisticsItem{
			OrderID:    p.GetOrderID(),
			TotalCost:  float64(p.GetTotalCost()),
			OpenDCost:  float64(p.GetOpenDCost()),
			NetDelay:   float64(p.GetNetDelay()),
			UpdateCost: float64(p.GetUpdateCost()),
		})
	}

	return &DelayStatistics{
		QotPushType:    stats.GetQotPushType(),
		DelayAvg:       float64(stats.GetDelayAvg()),
		Count:          stats.GetCount(),
		ItemList:       items,
		ReqReplyList:   reqReplyList,
		PlaceOrderList: placeOrderList,
	}, nil
}

// SuspendInfo represents suspension time for a security.
type SuspendInfo struct {
	Time      string
	Timestamp float64
}

// GetSuspend retrieves suspension information for securities.
func GetSuspend(c *Client, securities []*qotcommon.Security, beginTime, endTime string) ([]*SuspendInfo, error) {
	resp, err := qot.GetSuspend(c.inner, &qot.GetSuspendRequest{
		SecurityList: securities,
		BeginTime:    beginTime,
		EndTime:      endTime,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*SuspendInfo, 0)
	for _, s := range resp.SecuritySuspendList {
		if s == nil {
			continue
		}
		for _, su := range s.SuspendList {
			if su == nil {
				continue
			}
			result = append(result, &SuspendInfo{
				Time:      su.Time,
				Timestamp: su.Timestamp,
			})
		}
	}
	return result, nil
}

// PriceReminderOp constants for price reminder operations.
const (
	PriceReminderOpAdd    = 1
	PriceReminderOpUpdate = 2
	PriceReminderOpDelete = 3
)

// SetPriceReminder creates/updates/deletes a price reminder.
func SetPriceReminder(c *Client, market int32, code string, op, reminderType, freq int32, value float64, note string) (int64, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	resp, err := qot.SetPriceReminder(c.inner, &qot.SetPriceReminderRequest{
		Security: sec,
		Op:       op,
		Type:     reminderType,
		Freq:     freq,
		Value:    value,
		Note:     note,
	})
	if err != nil {
		return 0, err
	}
	return resp.Key, nil
}

// PriceReminderInfo represents a price reminder.
type PriceReminderInfo struct {
	Security *qotcommon.Security
	Name     string
	ItemList []PriceReminderItemInfo
}

// PriceReminderItemInfo represents a single price reminder item.
type PriceReminderItemInfo struct {
	Key                 int64
	Type                int32
	Freq                int32
	Value               float64
	Note                string
	IsEnable            bool
	ReminderSessionList []int32
}

// GetPriceReminder retrieves price reminders for a security.
func GetPriceReminder(c *Client, market int32, code string) ([]*PriceReminderInfo, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	resp, err := qot.GetPriceReminder(c.inner, sec, market)
	if err != nil {
		return nil, err
	}
	result := make([]*PriceReminderInfo, 0, len(resp.PriceReminderList))
	for _, pr := range resp.PriceReminderList {
		if pr == nil {
			continue
		}
		items := make([]PriceReminderItemInfo, 0, len(pr.ItemList))
		for _, item := range pr.ItemList {
			if item == nil {
				continue
			}
			items = append(items, PriceReminderItemInfo{
				Key:                 item.Key,
				Type:                item.Type,
				Freq:                item.Freq,
				Value:               item.Value,
				Note:                item.Note,
				IsEnable:            item.IsEnable,
				ReminderSessionList: item.ReminderSessionList,
			})
		}
		result = append(result, &PriceReminderInfo{
			Security: pr.Security,
			Name:     pr.Name,
			ItemList: items,
		})
	}
	return result, nil
}

// SubAccPush subscribes to account push notifications.
func SubAccPush(c *Client, accIDList []uint64) error {
	return trd.SubAccPush(c.inner, &trd.SubAccPushRequest{
		AccIDList: accIDList,
	})
}

// ReconfirmOrder reconfirms an order requiring additional verification.
func ReconfirmOrder(c *Client, accID uint64, market int32, orderID uint64, reason int32) (*ReconfirmOrderResult, error) {
	header := &trdcommon.TrdHeader{
		AccID:     &accID,
		TrdMarket: &market,
	}
	resp, err := trd.ReconfirmOrder(c.inner, &trd.ReconfirmOrderRequest{
		Header:          header,
		OrderID:         orderID,
		ReconfirmReason: reason,
	})
	if err != nil {
		return nil, err
	}
	return &ReconfirmOrderResult{
		AccID:     resp.Header.GetAccID(),
		TrdEnv:    resp.Header.GetTrdEnv(),
		TrdMarket: resp.Header.GetTrdMarket(),
		OrderID:   resp.OrderID,
	}, nil
}

type ReconfirmOrderResult struct {
	AccID     uint64
	TrdEnv    int32
	TrdMarket int32
	OrderID   uint64
}

// HoldingChangeInfo represents a holding change entry.
type HoldingChangeInfo struct {
	HolderName   string
	HoldingQty   float64
	HoldingRatio float64
	ChangeQty    float64
	ChangeRatio  float64
	Time         string
	Timestamp    float64
}

// GetHoldingChangeList retrieves holding change list.
func GetHoldingChangeList(c *Client, market int32, code string, holderCategory int32, beginTime, endTime string) ([]*HoldingChangeInfo, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	resp, err := qot.GetHoldingChangeList(c.inner, &qot.GetHoldingChangeListRequest{
		Security:       sec,
		HolderCategory: holderCategory,
		BeginTime:      beginTime,
		EndTime:        endTime,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*HoldingChangeInfo, 0, len(resp.HoldingChangeList))
	for _, h := range resp.HoldingChangeList {
		if h == nil {
			continue
		}
		result = append(result, &HoldingChangeInfo{
			HolderName:   getStr(h.HolderName),
			HoldingQty:   getFloat64(h.HoldingQty),
			HoldingRatio: getFloat64(h.HoldingRatio),
			ChangeQty:    getFloat64(h.ChangeQty),
			ChangeRatio:  getFloat64(h.ChangeRatio),
			Time:         getStr(h.Time),
			Timestamp:    getFloat64(h.Timestamp),
		})
	}
	return result, nil
}

// RehabInfo represents rehabilitation (复权) data.
type RehabInfo struct {
	Time       string
	FwdFactorA float64
	FwdFactorB float64
	BwdFactorA float64
	BwdFactorB float64
	SplitBase  int32
	SplitErt   int32
	JoinBase   int32
	JoinErt    int32
	BonusBase  int32
	BonusErt   int32
	AllotBase  int32
	AllotErt   int32
	AllotPrice float64
}

// RequestRehab requests rehabilitation (复权) data.
func RequestRehab(c *Client, market int32, code string) ([]*RehabInfo, error) {
	marketPtr := market
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	resp, err := qot.RequestRehab(c.inner, &qot.RequestRehabRequest{
		Security: sec,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*RehabInfo, 0, len(resp.RehabList))
	for _, r := range resp.RehabList {
		if r == nil {
			continue
		}
		result = append(result, &RehabInfo{
			Time:       getStr(r.Time),
			FwdFactorA: getFloat64(r.FwdFactorA),
			FwdFactorB: getFloat64(r.FwdFactorB),
			BwdFactorA: getFloat64(r.BwdFactorA),
			BwdFactorB: getFloat64(r.BwdFactorB),
			SplitBase:  getInt32(r.SplitBase),
			SplitErt:   getInt32(r.SplitErt),
			JoinBase:   getInt32(r.JoinBase),
			JoinErt:    getInt32(r.JoinErt),
			BonusBase:  getInt32(r.BonusBase),
			BonusErt:   getInt32(r.BonusErt),
			AllotBase:  getInt32(r.AllotBase),
			AllotErt:   getInt32(r.AllotErt),
			AllotPrice: getFloat64(r.AllotPrice),
		})
	}
	return result, nil
}

// HistoryKLQuotaInfo represents historical K-line quota info.
type HistoryKLQuotaInfo struct {
	UsedQuota   int32
	RemainQuota int32
	DetailList  []HistoryKLQuotaDetail
}

type HistoryKLQuotaDetail struct {
	Security         *qotcommon.Security
	Name             string
	RequestTime      string
	RequestTimestamp int64
}

// RequestHistoryKLQuota queries historical K-line quota.
func RequestHistoryKLQuota(c *Client) (*HistoryKLQuotaInfo, error) {
	resp, err := qot.RequestHistoryKLQuota(c.inner, &qot.RequestHistoryKLQuotaRequest{
		GetDetail: true,
	})
	if err != nil {
		return nil, err
	}
	details := make([]HistoryKLQuotaDetail, 0, len(resp.DetailList))
	for _, d := range resp.DetailList {
		if d == nil {
			continue
		}
		details = append(details, HistoryKLQuotaDetail{
			Security:         d.GetSecurity(),
			Name:             d.GetName(),
			RequestTime:      d.GetRequestTime(),
			RequestTimestamp: d.GetRequestTimeStamp(),
		})
	}
	return &HistoryKLQuotaInfo{
		UsedQuota:   resp.UsedQuota,
		RemainQuota: resp.RemainQuota,
		DetailList:  details,
	}, nil
}

// ============================================================================
// Types
// ============================================================================
// Types
// ============================================================================

// ============================================================================
// Types
// ============================================================================
// Types
// ============================================================================

// Quote represents a real-time quote.
type Quote struct {
	Symbol       string
	Market       int32
	Price        float64
	Open         float64
	High         float64
	Low          float64
	Volume       int64
	Timestamp    string
	Name         string
	LastClose    float64
	Turnover     float64
	TurnoverRate float64
	Amplitude    float64
}

// KLine represents a K-line (candlestick) data point.
type KLine struct {
	Time       string
	Open       float64
	High       float64
	Low        float64
	Close      float64
	Volume     int64
	LastClose  float64
	Turnover   float64
	ChangeRate float64
	Timestamp  float64
}

// Account represents a trading account.
type Account struct {
	AccID             uint64
	AccType           int32
	TrdEnv            int32
	CardNum           string
	AccStatus         int32
	TrdMarketAuthList []int32
	SecurityFirm      int32
	SimAccType        int32
	UniCardNum        string
	AccRole           int32
	JpAccType         []int32
}

// PlaceOrderResult represents a place order result.
type PlaceOrderResult struct {
	OrderID   uint64
	OrderIDEx string
}

// Position represents a position.
type Position struct {
	PositionID       uint64
	Code             string
	Name             string
	Market           int32
	Quantity         float64
	CanSellQty       float64
	CostPrice        float64
	CurPrice         float64
	MarketVal        float64
	PnL              float64
	PnLRate          float64
	TodayBuyQty      float64
	TodayBuyVal      float64
	TodaySellQty     float64
	TodaySellVal     float64
	TodayPnL         float64
	UnrealizedPL     float64
	RealizedPL       float64
	Currency         int32
	TrdMarket        int32
	DilutedCostPrice float64
	AverageCostPrice float64
	AveragePnLRate   float64
}

// Funds represents account funds.
type Funds struct {
	Power             float64
	TotalAssets       float64
	Cash              float64
	MarketVal         float64
	FrozenCash        float64
	DebtCash          float64
	AvlWithdrawalCash float64
	Currency          int32
	AvailableFunds    float64
	UnrealizedPL      float64
	RealizedPL        float64
	RiskLevel         int32
	InitialMargin     float64
	MaintenanceMargin float64
	MaxPowerShort     float64
	NetCashPower      float64
	LongMv            float64
	ShortMv           float64
	PendingAsset      float64
	MaxWithdrawal     float64
	RiskStatus        int32
}

// Order represents an order.
type Order struct {
	OrderID         uint64
	OrderIDEx       string
	Code            string
	Name            string
	TrdSide         int32
	OrderType       int32
	OrderStatus     int32
	Price           float64
	Qty             float64
	FillQty         float64
	FillAvgPrice    float64
	CreateTime      string
	UpdateTime      string
	LastErrMsg      string
	SecMarket       int32
	CreateTimestamp float64
	UpdateTimestamp float64
	Remark          string
	TimeInForce     int32
	FillOutsideRTH  bool
	AuxPrice        float64
	TrailType       int32
	TrailValue      float64
	TrailSpread     float64
	Currency        int32
	TrdMarket       int32
	Session         int32
}

// OrderFill represents an order fill.
type OrderFill struct {
	FillID            uint64
	FillIDEx          string
	OrderID           uint64
	OrderIDEx         string
	Code              string
	Name              string
	TrdSide           int32
	Price             float64
	Qty               float64
	CreateTime        string
	CounterBrokerID   int32
	CounterBrokerName string
	SecMarket         int32
	CreateTimestamp   float64
	UpdateTimestamp   float64
	Status            int32
	TrdMarket         int32
	JpAccType         int32
}

// OrderBook represents order book data.
type OrderBook struct {
	Bids                    []OrderBookItem
	Asks                    []OrderBookItem
	SvrRecvTimeBid          string
	SvrRecvTimeBidTimestamp float64
	SvrRecvTimeAsk          string
	SvrRecvTimeAskTimestamp float64
}

// OrderBookItem represents a single order book entry.
type OrderBookItem struct {
	Price      float64
	Volume     int64
	OrderCount int32
	DetailList []OrderBookDetail
}

// OrderBookDetail represents a detail entry in an order book item.
type OrderBookDetail struct {
	OrderID int64
	Volume  int64
}

// Ticker represents ticker data.
type Ticker struct {
	Time      string
	Sequence  int64
	Price     float64
	Volume    int64
	Direction string
	Turnover  float64
	RecvTime  float64
	Type      int32
	TypeSign  int32
	Timestamp float64
}

// RT represents real-time data.
type RT struct {
	Time      string
	Price     float64
	Volume    int64
	LastClose float64
	AvgPrice  float64
	Turnover  float64
}

// Broker represents broker data.
type Broker struct {
	ID     int64
	Name   string
	Pos    int32
	Volume int64
}

// StaticInfo represents static security info.
type StaticInfo struct {
	Code     string
	Name     string
	Type     int32
	ListTime string
	LotSize  int32
}

// FutureInfo represents futures info.
type FutureInfo struct {
	Code               string
	Name               string
	Expire             string
	LastTradeTimestamp float64
	Owner              string
	OwnerOther         string
	Exchange           string
	ContractType       string
	ContractSize       float64
	ContractSizeUnit   string
	QuoteCurrency      string
	MinVar             float64
	MinVarUnit         string
	QuoteUnit          string
	TimeZone           string
	ExchangeFormatUrl  string
}

// Plate represents a market plate (板块).
type Plate struct {
	Code string
	Name string
}

// IpoData represents IPO data.
type IpoData struct {
	Code          string
	Name          string
	ListDate      string
	ListTimestamp float64
}

// UserSecurityGroup represents user security group.
type UserSecurityGroup struct {
	Name      string
	GroupType int32
}

// SubInfo represents subscription info.
type SubInfo struct {
	IsSub    bool
	SubTypes []int32
	Security string
}

// CapitalDistribution represents capital distribution.
type CapitalDistribution struct {
	MainInflow      float64
	MainOutflow     float64
	MidInflow       float64
	MidOutflow      float64
	SmallInflow     float64
	SmallOutflow    float64
	BigInflow       float64
	BigOutflow      float64
	UpdateTime      string
	UpdateTimestamp float64
}

// OptionExpiration represents option expiration date.
type OptionExpiration struct {
	Date string
	Days int32
	Desc string
}

// OptionItem represents option chain item.
type OptionItem struct {
	Code         string
	Name         string
	CallPut      int32
	Strike       float64
	Expire       string
	Volume       int64
	OpenInterest int64
}

// WarrantItem represents warrant data.
type WarrantItem struct {
	Code    string
	Name    string
	CallPut int32
	Strike  float64
	Expire  string
	Volume  int64
	Price   float64
}

// StockFilterItem represents stock filter result.
type StockFilterItem struct {
	Code      string
	Name      string
	Price     float64
	ChangePct float64
	Volume    int64
	Amount    float64
}

// MarketStateInfo represents market state info.
type MarketStateInfo struct {
	Code        string
	Name        string
	MarketState int32
}

// Common market constants.
const (
	// QotMarket
	Market_HK_Security   = int32(qotcommon.QotMarket_QotMarket_HK_Security)
	Market_HK_Future     = int32(qotcommon.QotMarket_QotMarket_HK_Future)
	Market_US_Security   = int32(qotcommon.QotMarket_QotMarket_US_Security)
	Market_CNSH_Security = int32(qotcommon.QotMarket_QotMarket_CNSH_Security)
	Market_CNSZ_Security = int32(qotcommon.QotMarket_QotMarket_CNSZ_Security)

	// TrdSide
	Side_Buy  = int32(trdcommon.TrdSide_TrdSide_Buy)
	Side_Sell = int32(trdcommon.TrdSide_TrdSide_Sell)

	// OrderType
	OrderType_Normal = int32(trdcommon.OrderType_OrderType_Normal)
	OrderType_Market = int32(trdcommon.OrderType_OrderType_Market)
	OrderType_Stop   = int32(trdcommon.OrderType_OrderType_Stop)

	// KLType
	KLType_Day   = int32(qotcommon.KLType_KLType_Day)
	KLType_1Min  = int32(qotcommon.KLType_KLType_1Min)
	KLType_5Min  = int32(qotcommon.KLType_KLType_5Min)
	KLType_15Min = int32(qotcommon.KLType_KLType_15Min)
	KLType_30Min = int32(qotcommon.KLType_KLType_30Min)
	KLType_60Min = int32(qotcommon.KLType_KLType_60Min)
	KLType_Week  = int32(qotcommon.KLType_KLType_Week)
	KLType_Month = int32(qotcommon.KLType_KLType_Month)

	// SubType
	SubType_Basic     = int32(qot.SubType_Basic)
	SubType_OrderBook = int32(qot.SubType_OrderBook)
	SubType_Ticker    = int32(qot.SubType_Ticker)
	SubType_RT        = int32(qot.SubType_RT)
	SubType_KL        = int32(qot.SubType_KL)
	SubType_KL_1Min   = int32(qot.SubType_KL_1Min)
	SubType_KL_5Min   = int32(qot.SubType_KL_5Min)
	SubType_KL_15Min  = int32(qot.SubType_KL_15Min)
	SubType_KL_30Min  = int32(qot.SubType_KL_30Min)
	SubType_KL_60Min  = int32(qot.SubType_KL_60Min)
	SubType_KL_Day    = int32(qot.SubType_KL_Day)
	SubType_KL_Week   = int32(qot.SubType_KL_Week)
	SubType_KL_Month  = int32(qot.SubType_KL_Month)
	SubType_Broker    = int32(qot.SubType_Broker)
)

// ============================================================================
// Push Notification Handlers
// ============================================================================

// PushQuote represents a parsed real-time quote push notification.
type PushQuote struct {
	Market    int32
	Code      string
	Name      string
	CurPrice  float64
	OpenPrice float64
	HighPrice float64
	LowPrice  float64
	Volume    int64
	Turnover  float64
}

// PushKLine represents a parsed K-line push notification.
type PushKLine struct {
	Market int32
	Code   string
	Name   string
	KLType int32
	KLine
}

// PushOrderBook represents a parsed order book push notification.
type PushOrderBook struct {
	Market int32
	Code   string
	Name   string
	Bids   []OBItem
	Asks   []OBItem
}

// OBItem represents a single price level in the order book push data.
type OBItem struct {
	Price      float64
	Volume     int64
	OrderCount int64
}

// PushTicker represents a parsed tick-by-tick push notification.
type PushTicker struct {
	Market   int32
	Code     string
	Name     string
	Price    float64
	Volume   int64
	Turnover float64
	Side     int32 // TickerDirection: 1=Buy, 2=Sell, 0=Unknown
}

// SetPushHandler registers a handler that receives push notifications for
// specific protoIDs. The handler receives (protoID, rawBody) and should use
// the ParsePush* functions below to decode the body.
func (c *Client) SetPushHandler(protoID uint32, h func(protoID uint32, body []byte)) {
	c.inner.RegisterHandler(protoID, h)
}

// ParsePushQuote parses a raw push body (ProtoID 3005) into a PushQuote.
func ParsePushQuote(body []byte) (*PushQuote, error) {
	data, err := push.ParseUpdateBasicQot(body)
	if err != nil || data == nil {
		return nil, err
	}
	return &PushQuote{
		Market:    data.Security.GetMarket(),
		Code:      data.Security.GetCode(),
		Name:      data.Name,
		CurPrice:  data.CurPrice,
		OpenPrice: data.OpenPrice,
		HighPrice: data.HighPrice,
		LowPrice:  data.LowPrice,
		Volume:    data.Volume,
		Turnover:  data.Turnover,
	}, nil
}

// ParsePushKLine parses a raw push body (ProtoID 3007) into a PushKLine.
func ParsePushKLine(body []byte) (*PushKLine, error) {
	data, err := push.ParseUpdateKL(body)
	if err != nil || data == nil || len(data.KLList) == 0 {
		return nil, err
	}
	kl := data.KLList[0]
	return &PushKLine{
		Market: data.Security.GetMarket(),
		Code:   data.Security.GetCode(),
		Name:   data.Name,
		KLType: data.KlType,
		KLine: KLine{
			Time:       kl.GetTime(),
			Open:       kl.GetOpenPrice(),
			High:       kl.GetHighPrice(),
			Low:        kl.GetLowPrice(),
			Close:      kl.GetClosePrice(),
			Volume:     kl.GetVolume(),
			LastClose:  kl.GetLastClosePrice(),
			Turnover:   kl.GetTurnover(),
			ChangeRate: kl.GetChangeRate(),
			Timestamp:  kl.GetTimestamp(),
		},
	}, nil
}

// ParsePushOrderBook parses a raw push body (ProtoID 3013) into a PushOrderBook.
func ParsePushOrderBook(body []byte) (*PushOrderBook, error) {
	data, err := push.ParseUpdateOrderBook(body)
	if err != nil || data == nil {
		return nil, err
	}
	ob := &PushOrderBook{
		Market: data.Security.GetMarket(),
		Code:   data.Security.GetCode(),
		Name:   data.Name,
	}
	for _, b := range data.OrderBookBidList {
		ob.Bids = append(ob.Bids, OBItem{
			Price:      b.GetPrice(),
			Volume:     b.GetVolume(),
			OrderCount: int64(b.GetOrederCount()),
		})
	}
	for _, a := range data.OrderBookAskList {
		ob.Asks = append(ob.Asks, OBItem{
			Price:      a.GetPrice(),
			Volume:     a.GetVolume(),
			OrderCount: int64(a.GetOrederCount()),
		})
	}
	return ob, nil
}

// ParsePushTicker parses a raw push body (ProtoID 3011) into a PushTicker.
func ParsePushTicker(body []byte) (*PushTicker, error) {
	data, err := push.ParseUpdateTicker(body)
	if err != nil || data == nil || len(data.TickerList) == 0 {
		return nil, err
	}
	t := data.TickerList[0]
	return &PushTicker{
		Market:   data.Security.GetMarket(),
		Code:     data.Security.GetCode(),
		Name:     data.Name,
		Price:    t.GetPrice(),
		Volume:   t.GetVolume(),
		Turnover: t.GetTurnover(),
		Side:     t.GetDir(),
	}, nil
}

// PushRT represents a parsed real-time minute data push notification.
type PushRT struct {
	Market   int32
	Code     string
	Name     string
	Time     string
	Price    float64
	Volume   int64
	AvgPrice float64
	Turnover float64
}

// ParsePushRT parses a raw push body (ProtoID 3009) into PushRT data.
func ParsePushRT(body []byte) (*PushRT, error) {
	data, err := push.ParseUpdateRT(body)
	if err != nil || data == nil || len(data.RTList) == 0 {
		return nil, err
	}
	rt := data.RTList[0]
	return &PushRT{
		Market:   data.Security.GetMarket(),
		Code:     data.Security.GetCode(),
		Name:     data.Name,
		Time:     rt.GetTime(),
		Price:    rt.GetPrice(),
		Volume:   rt.GetVolume(),
		AvgPrice: rt.GetAvgPrice(),
	}, nil
}

// PushBroker represents a parsed broker queue push notification.
type PushBroker struct {
	Market int32
	Code   string
	Name   string
	Asks   []BrokerItem
	Bids   []BrokerItem
}

// BrokerItem represents a single broker queue entry.
type BrokerItem struct {
	Price    float64
	Volume   int64
	BrokerID int32
}

// ParsePushBroker parses a raw push body (ProtoID 3015) into PushBroker data.
func ParsePushBroker(body []byte) (*PushBroker, error) {
	data, err := push.ParseUpdateBroker(body)
	if err != nil || data == nil {
		return nil, err
	}
	ob := &PushBroker{
		Market: data.Security.GetMarket(),
		Code:   data.Security.GetCode(),
		Name:   data.Name,
	}
	for _, a := range data.AskBrokerList {
		ob.Asks = append(ob.Asks, BrokerItem{
			Volume:   a.GetVolume(),
			BrokerID: int32(a.GetId()),
		})
	}
	for _, b := range data.BidBrokerList {
		ob.Bids = append(ob.Bids, BrokerItem{
			Volume:   b.GetVolume(),
			BrokerID: int32(b.GetId()),
		})
	}
	return ob, nil
}

// Push ProtoID constants (re-exported from pkg/push for convenience).
const (
	ProtoID_Qot_UpdateBasicQot  = 3005
	ProtoID_Qot_UpdateKL        = 3007
	ProtoID_Qot_UpdateOrderBook = 3013
	ProtoID_Qot_UpdateTicker    = 3011
	ProtoID_Qot_UpdateRT        = 3009
	ProtoID_Qot_UpdateBroker    = 3015
)

// ============================================================================
// Options (aliases for internal client options)
type Option = futuapi.Option

// WithDialTimeout sets the connection dial timeout.
func WithDialTimeout(d time.Duration) Option {
	return futuapi.WithDialTimeout(d)
}

// WithAPISetTimeout sets the API request timeout.
func WithAPISetTimeout(d time.Duration) Option {
	return futuapi.WithAPITimeout(d)
}

// WithKeepAliveInterval sets the keep-alive interval.
func WithKeepAliveInterval(d time.Duration) Option {
	return futuapi.WithKeepAliveInterval(d)
}

// WithReconnectInterval sets the base reconnect interval when the connection drops.
func WithReconnectInterval(d time.Duration) Option {
	return futuapi.WithReconnectInterval(d)
}

// WithMaxRetries sets the maximum retry attempts.
func WithMaxRetries(n int) Option {
	return futuapi.WithMaxRetries(n)
}

// WithLogLevel sets the logging level (0=info, 1=warn, 2=error, 3=silent).
func WithLogLevel(level int) Option {
	return futuapi.WithLogLevel(level)
}

// Default timeouts and connection limits.
const (
	DefaultDialTimeout      = 10 * time.Second
	DefaultAPITimeout       = 30 * time.Second
	DefaultKeepAlive        = 30 * time.Second
	DefaultMaxRetries       = 3
	DefaultReconnectBackoff = 1.5
)

// Trading environment constants.
const (
	TrdEnv_Real     = int32(0) // Real trading
	TrdEnv_Simulate = int32(1) // Simulate/paper trading (default)
)
