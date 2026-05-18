package client

import (
	"context"
	"fmt"
	"time"

	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdflowsummary"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

// GetAccountList retrieves the list of trading accounts.
func GetAccountList(ctx context.Context, c *Client) ([]Account, error) {
	resp, err := trd.GetAccList(ctx, c.inner, constant.TrdCategory_Security, false)
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
func UnlockTrading(ctx context.Context, c *Client, pwdMD5 string) error {
	return trd.UnlockTrade(ctx, c.inner, &trd.UnlockTradeRequest{
		Unlock: true,
		PwdMD5: constant.SensitiveString(pwdMD5),
	})
}

// PlaceOrder places a trading order.
func PlaceOrder(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, code string, side constant.TrdSide, orderType constant.OrderType, price float64, qty float64, secMarket constant.TrdSecMarket) (*PlaceOrderResult, error) {
	if secMarket == 0 {
		secMarket = constant.TrdSecMarket(inferSecMarket(code))
	}
	resp, err := trd.PlaceOrder(ctx, c.inner, &trd.PlaceOrderRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    c.trdEnv,
		Code:      code,
		TrdSide:   side,
		OrderType: orderType,
		Price:     price,
		Qty:       qty,
		SecMarket: secMarket,
	})
	if err != nil {
		return nil, err
	}
	return &PlaceOrderResult{Header: resp.Header, OrderID: resp.OrderID, OrderIDEx: resp.OrderIDEx}, nil
}

// ModifyOrder modifies or cancels an existing order.
func ModifyOrder(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, orderID uint64, modifyOp constant.ModifyOrderOp, price float64, qty float64) (*trd.ModifyOrderResponse, error) {
	return trd.ModifyOrder(ctx, c.inner, &trd.ModifyOrderRequest{
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
func CancelAllOrder(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, trdEnv constant.TrdEnv) error {
	_, err := trd.ModifyOrder(ctx, c.inner, &trd.ModifyOrderRequest{
		AccID:         accID,
		TrdMarket:     market,
		TrdEnv:        trdEnv,
		OrderID:       0,
		ModifyOrderOp: constant.ModifyOrderOp_Cancel,
		Price:         0,
		Qty:           1,
		ForAll:        true,
	})
	return err
}

// GetPositionList retrieves the current positions.
func GetPositionList(ctx context.Context, c *Client, accID uint64) ([]Position, error) {
	resp, err := trd.GetPositionList(ctx, c.inner, &trd.GetPositionListRequest{
		AccID:     accID,
		TrdMarket: constant.TrdMarket_None,
		TrdEnv:    c.trdEnv,
	})
	if err != nil {
		return nil, err
	}

	positions := make([]Position, len(resp.PositionList))
	for i, p := range resp.PositionList {
		positions[i] = Position{
			PositionID:       p.PositionID,
			PositionSide:      p.PositionSide,
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
			TrdMarket:        p.SecMarket,
			DilutedCostPrice: p.DilutedCostPrice,
			AverageCostPrice: p.AverageCostPrice,
			AveragePnLRate:   p.AveragePlRatio,
			SecMarket:        p.SecMarket,
			TdTrdVal:         p.TdTrdVal,
		}
	}
	return positions, nil
}

// GetAccountInfo retrieves full account information including multi-currency cash and per-market assets.
func GetAccountInfo(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket) (*Funds, error) {
	resp, err := trd.GetFunds(ctx, c.inner, &trd.GetFundsRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    c.trdEnv,
	})
	if err != nil {
		return nil, err
	}
	f := resp.Funds
	cashList := make([]AccCashInfo, 0, len(f.CashInfoList))
	for _, ci := range f.CashInfoList {
		cashList = append(cashList, AccCashInfo{
			Currency:         ci.Currency,
			Cash:             ci.Cash,
			AvailableBalance: ci.AvailableBalance,
			NetCashPower:     ci.NetCashPower,
		})
	}
	marketList := make([]AccMarketInfo, 0, len(f.MarketInfoList))
	for _, m := range f.MarketInfoList {
		marketList = append(marketList, AccMarketInfo{
			TrdMarket: m.TrdMarket,
			Assets:    m.Assets,
		})
	}
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
		MarginCallMargin:  f.MarginCallMargin,
		IsPDT:             f.IsPDT,
		PDTSeq:            f.PDTSeq,
		BeginningDTBP:     f.BeginningDTBP,
		RemainingDTBP:     f.RemainingDTBP,
		DtCallAmount:      f.DtCallAmount,
		DtStatus:          f.DtStatus,
		CashInfoList:      cashList,
		MarketInfoList:    marketList,
		SecuritiesAssets:  f.SecuritiesAssets,
		FundAssets:       f.FundAssets,
		BondAssets:       f.BondAssets,
	}, nil
}

// GetFunds retrieves account funds for a specific account.
func GetFunds(ctx context.Context, c *Client, accID uint64) (*Funds, error) {
	accounts, err := GetAccountList(ctx, c)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, fmt.Errorf("no accounts available")
	}
	var acc Account
	found := false
	for _, a := range accounts {
		if a.AccID == accID {
			acc = a
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("account %d not found", accID)
	}
	if len(acc.TrdMarketAuthList) == 0 {
		return nil, fmt.Errorf("account %d has no market authorization", accID)
	}
	return GetAccountInfo(ctx, c, acc.AccID, constant.TrdMarket(acc.TrdMarketAuthList[0]))
}

// GetMaxTrdQtys retrieves maximum tradable quantities.
func GetMaxTrdQtys(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, code string, orderType constant.OrderType, price float64, secMarket constant.TrdSecMarket) (*MaxTrdQtysInfo, error) {
	if secMarket == 0 {
		secMarket = constant.TrdSecMarket(inferSecMarket(code))
	}
	resp, err := trd.GetMaxTrdQtys(ctx, c.inner, &trd.GetMaxTrdQtysRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    c.trdEnv,
		Code:      code,
		OrderType: orderType,
		Price:     price,
		SecMarket: secMarket,
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
		LongRequiredIM:      m.LongRequiredIM,
		ShortRequiredIM:     m.ShortRequiredIM,
		Session:             m.Session,
	}, nil
}

// GetOrderFee retrieves order fee information.
func GetOrderFee(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, orderIDExList []string) ([]*OrderFeeInfo, error) {
	resp, err := trd.GetOrderFee(ctx, c.inner, &trd.GetOrderFeeRequest{
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

// GetMarginRatio retrieves margin ratio for securities.
func GetMarginRatio(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, securities []*qotcommon.Security) ([]*MarginRatioInfo, error) {
	resp, err := trd.GetMarginRatio(ctx, c.inner, &trd.GetMarginRatioRequest{
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
			Security:        m.Security,
			IsLongPermit:    m.IsLongPermit,
			IsShortPermit:   m.IsShortPermit,
			ShortFeeRate:    m.ShortFeeRate,
			ImLongRatio:     m.ImLongRatio,
			ImShortRatio:    m.ImShortRatio,
			ShortPoolRemain: m.ShortPoolRemain,
			AlertLongRatio:  m.AlertLongRatio,
			AlertShortRatio: m.AlertShortRatio,
			McmLongRatio:    m.McmLongRatio,
			McmShortRatio:   m.McmShortRatio,
			MmLongRatio:     m.MmLongRatio,
			MmShortRatio:    m.MmShortRatio,
		})
	}
	return result, nil
}

// GetOrderList retrieves active orders.
func GetOrderList(ctx context.Context, c *Client, accID uint64) ([]Order, error) {
	resp, err := trd.GetOrderList(ctx, c.inner, &trd.GetOrderListRequest{
		AccID:     accID,
		TrdMarket: constant.TrdMarket_None,
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
			JpAccType:       o.JpAccType,
		}
	}
	return orders, nil
}

// GetHistoryOrderList retrieves historical orders.
func GetHistoryOrderList(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, startDate, endDate string) ([]Order, error) {
	var fc *trdcommon.TrdFilterConditions
	if startDate != "" || endDate != "" {
		fc = &trdcommon.TrdFilterConditions{
			BeginTime: &startDate,
			EndTime:   &endDate,
		}
	}
	resp, err := trd.GetHistoryOrderList(ctx, c.inner, &trd.GetHistoryOrderListRequest{
		AccID:            accID,
		TrdMarket:        market,
		TrdEnv:           c.trdEnv,
		FilterConditions: fc,
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
			JpAccType:       o.JpAccType,
		})
	}
	return orders, nil
}

// GetOrderFillList retrieves order fills (executions).
func GetOrderFillList(ctx context.Context, c *Client, accID uint64) ([]OrderFill, error) {
	resp, err := trd.GetOrderFillList(ctx, c.inner, &trd.GetOrderFillListRequest{
		AccID:     accID,
		TrdMarket: constant.TrdMarket_None,
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
func GetHistoryOrderFillList(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket) ([]OrderFill, error) {
	resp, err := trd.GetHistoryOrderFillList(ctx, c.inner, &trd.GetHistoryOrderFillListRequest{
		AccID:            accID,
		TrdMarket:        market,
		TrdEnv:           c.trdEnv,
		FilterConditions: &trdcommon.TrdFilterConditions{},
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

// GetFlowSummary retrieves account cash flow entries.
func GetFlowSummary(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, clearingDate string, direction trdflowsummary.TrdCashFlowDirection) ([]*FlowSummaryInfo, error) {
	if clearingDate == "" {
		clearingDate = time.Now().Format("2006-01-02")
	}
	resp, err := trd.GetFlowSummary(ctx, c.inner, &trd.GetFlowSummaryRequest{
		AccID:             accID,
		TrdMarket:         constant.TrdMarket(market),
		TrdEnv:            c.trdEnv,
		ClearingDate:      clearingDate,
		CashFlowDirection: int32(direction),
	})
	if err != nil {
		return nil, err
	}

	result := make([]*FlowSummaryInfo, 0, len(resp.FlowSummaryList))
	for _, f := range resp.FlowSummaryList {
		if f == nil {
			continue
		}
		result = append(result, &FlowSummaryInfo{
			CashFlowID:        f.CashFlowID,
			ClearingDate:      f.ClearingDate,
			SettlementDate:    f.SettlementDate,
			Currency:          f.Currency,
			CashFlowType:      f.CashFlowType,
			CashFlowDirection: f.CashFlowDirection,
			CashFlowAmount:    f.CashFlowAmount,
			CashFlowRemark:    f.CashFlowRemark,
		})
	}
	return result, nil
}

// GetAccTradingInfo retrieves maximum tradable quantities and margin info for a security.
func GetAccTradingInfo(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, code string, orderType constant.OrderType, price float64) (*AccTradingInfo, error) {
	secMarket := constant.MarketToTrdSecMarket[int32(market)]
	resp, err := trd.GetMaxTrdQtys(ctx, c.inner, &trd.GetMaxTrdQtysRequest{
		AccID:     accID,
		TrdMarket: market,
		TrdEnv:    c.trdEnv,
		OrderType: orderType,
		Code:      code,
		Price:     price,
		SecMarket: secMarket,
	})
	if err != nil {
		return nil, err
	}
	if resp.MaxTrdQtys == nil {
		return nil, fmt.Errorf("GetAccTradingInfo: MaxTrdQtys is nil")
	}
	m := resp.MaxTrdQtys
	return &AccTradingInfo{
		MaxCashBuy:          m.MaxCashBuy,
		MaxCashAndMarginBuy: m.MaxCashAndMarginBuy,
		MaxPositionSell:     m.MaxPositionSell,
		MaxSellShort:        m.MaxSellShort,
		MaxBuyBack:          m.MaxBuyBack,
		LongRequiredIM:      m.LongRequiredIM,
		ShortRequiredIM:     m.ShortRequiredIM,
		Session:             m.Session,
	}, nil
}

// SubAccPush subscribes to account push notifications.
func SubAccPush(ctx context.Context, c *Client, accIDList []uint64) error {
	return trd.SubAccPush(ctx, c.inner, &trd.SubAccPushRequest{
		AccIDList: accIDList,
	})
}

// ReconfirmOrder reconfirms an order requiring additional verification.
func ReconfirmOrder(ctx context.Context, c *Client, accID uint64, market constant.TrdMarket, orderID uint64, reason int32) (*ReconfirmOrderResult, error) {
	connID := c.inner.GetConnID()
	serialNo := c.inner.NextSerialNo()
	resp, err := trd.ReconfirmOrder(ctx, c.inner, &trd.ReconfirmOrderRequest{
		PacketID: &common.PacketID{
			ConnID:   &connID,
			SerialNo: &serialNo,
		},
		AccID:           accID,
		TrdMarket:       market,
		TrdEnv:          c.trdEnv,
		OrderID:         orderID,
		ReconfirmReason: reason,
	})
	if err != nil {
		return nil, err
	}
	return &ReconfirmOrderResult{
		AccID:     resp.AccID,
		TrdEnv:    resp.TrdEnv,
		TrdMarket: resp.TrdMarket,
		JpAccType: resp.JpAccType,
		OrderID:   resp.OrderID,
	}, nil
}
