package client

import (
	"context"
	"fmt"
	"time"

	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetipolist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetreference"
	"github.com/shing1211/futuapi4go/pkg/pb/qotstockfilter"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/pkg/util"
)

// History KL pagination constants.
const (
	DefaultHistoryKLPageSize = 1000
	DefaultHistoryKLDelay    = 200 * time.Millisecond
)

// HistoryKLPaginationDelay controls the wait time between pages when
// fetching historical K-lines.
var HistoryKLPaginationDelay = DefaultHistoryKLDelay

// GetQuote retrieves the current quote for a security.
func GetQuote(ctx context.Context, c *Client, market constant.Market, code string) (*Quote, error) {
	if code == "" {
		return nil, fmt.Errorf("GetQuote: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	quotes, err := qot.GetBasicQot(ctx, c.inner, []*qotcommon.Security{sec})
	if err != nil {
		return nil, err
	}
	if len(quotes) == 0 {
		return nil, fmt.Errorf("no quote returned for %s", code)
	}

	q := quotes[0]
	return &Quote{
		Symbol:          code,
		Market:          int32(market),
		Price:           q.CurPrice,
		Open:            q.OpenPrice,
		High:            q.HighPrice,
		Low:             q.LowPrice,
		Volume:          q.Volume,
		Timestamp:       q.UpdateTime,
		Name:            q.Name,
		LastClose:       q.LastClosePrice,
		Turnover:        q.Turnover,
		TurnoverRate:    q.TurnoverRate,
		Amplitude:       q.Amplitude,
		IsSuspended:     q.IsSuspended,
		SecStatus:       q.SecStatus,
		ListTime:        q.ListTime,
		PriceSpread:     q.PriceSpread,
		DarkStatus:      q.DarkStatus,
		ListTimestamp:   q.ListTimestamp,
		UpdateTimestamp: q.UpdateTimestamp,
		PreMarket:       mapPreAfterMarketData(q.PreMarket),
		AfterMarket:     mapPreAfterMarketData(q.AfterMarket),
		Overnight:       mapPreAfterMarketData(q.Overnight),
		OptionExData:    q.OptionExData,
		FutureExData:    q.FutureExData,
		WarrantExData:   q.WarrantExData,
	}, nil
}

func mapPreAfterMarketData(d *qotcommon.PreAfterMarketData) *PreAfterMarketData {
	if d == nil {
		return nil
	}
	return &PreAfterMarketData{
		Price:      getFloat64(d.Price),
		HighPrice:  getFloat64(d.HighPrice),
		LowPrice:   getFloat64(d.LowPrice),
		Volume:     getInt64(d.Volume),
		Turnover:   getFloat64(d.Turnover),
		ChangeVal:  getFloat64(d.ChangeVal),
		ChangeRate: getFloat64(d.ChangeRate),
		Amplitude:  getFloat64(d.Amplitude),
	}
}

func mapRehabInfo(r *qotcommon.Rehab) *RehabInfo {
	if r == nil {
		return nil
	}
	return &RehabInfo{
		Time:           getStr(r.Time),
		CompanyActFlag: getInt64(r.CompanyActFlag),
		FwdFactorA:     getFloat64(r.FwdFactorA),
		FwdFactorB:     getFloat64(r.FwdFactorB),
		BwdFactorA:     getFloat64(r.BwdFactorA),
		BwdFactorB:     getFloat64(r.BwdFactorB),
		SplitBase:      getInt32(r.SplitBase),
		SplitErt:       getInt32(r.SplitErt),
		JoinBase:       getInt32(r.JoinBase),
		JoinErt:        getInt32(r.JoinErt),
		BonusBase:      getInt32(r.BonusBase),
		BonusErt:       getInt32(r.BonusErt),
		TransferBase:   getInt32(r.TransferBase),
		TransferErt:    getInt32(r.TransferErt),
		AllotBase:      getInt32(r.AllotBase),
		AllotErt:       getInt32(r.AllotErt),
		AllotPrice:     getFloat64(r.AllotPrice),
		AddBase:        getInt32(r.AddBase),
		AddErt:         getInt32(r.AddErt),
		AddPrice:       getFloat64(r.AddPrice),
		Dividend:       getFloat64(r.Dividend),
		SpDividend:     getFloat64(r.SpDividend),
		SpinOffBase:    getFloat64(r.SpinOffBase),
		SpinOffErt:     getFloat64(r.SpinOffErt),
		Timestamp:      getFloat64(r.Timestamp),
	}
}

// GetKLines retrieves K-line (candlestick) data.
func GetKLines(ctx context.Context, c *Client, market constant.Market, code string, klType constant.KLType, num int) (*KLinesResult, error) {
	if code == "" {
		return nil, fmt.Errorf("GetKLines: code is required")
	}
	if num <= 0 {
		return nil, fmt.Errorf("GetKLines: num must be greater than 0")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetKL(ctx, c.inner, &qot.GetKLRequest{
		Security:  sec,
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(klType),
		ReqNum:    int32(num),
	})
	if err != nil {
		return nil, err
	}

	klines := make([]KLine, len(resp.KLList))
	for i, kl := range resp.KLList {
		klines[i] = KLine{
			Time:         kl.Time,
			IsBlank:      kl.IsBlank,
			Open:         kl.OpenPrice,
			High:         kl.HighPrice,
			Low:          kl.LowPrice,
			Close:        kl.ClosePrice,
			Volume:       kl.Volume,
			LastClose:    kl.LastClosePrice,
			Turnover:     kl.Turnover,
			TurnoverRate: kl.TurnoverRate,
			Pe:           kl.Pe,
			ChangeRate:   kl.ChangeRate,
			Timestamp:    kl.Timestamp,
		}
	}

	result := &KLinesResult{
		Items: klines,
	}
	if resp.Security != nil {
		result.Security = resp.Security
	}
	if resp.Name != "" {
		result.Name = resp.Name
	}

	return result, nil
}

// Subscribe subscribes to real-time market data.
func Subscribe(ctx context.Context, c *Client, market constant.Market, code string, subTypes []constant.SubType) error {
	if code == "" {
		return fmt.Errorf("Subscribe: code is required")
	}
	if len(subTypes) == 0 {
		return fmt.Errorf("Subscribe: subTypes is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
		SecurityList:     []*qotcommon.Security{sec},
		SubTypeList:      subTypesConverted,
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
		IsFirstPush:      true,
	})
	return err
}

// Unsubscribe unsubscribes from real-time market data.
func Unsubscribe(ctx context.Context, c *Client, market constant.Market, code string, subTypes []constant.SubType) error {
	if code == "" {
		return fmt.Errorf("Unsubscribe: code is required")
	}
	if len(subTypes) == 0 {
		return fmt.Errorf("Unsubscribe: subTypes is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
		SecurityList:     []*qotcommon.Security{sec},
		SubTypeList:      subTypesConverted,
		IsSubOrUnSub:     false,
		IsRegOrUnRegPush: false,
	})
	return err
}

// UnsubscribeAll unsubscribes from all market data.
func UnsubscribeAll(ctx context.Context, c *Client) error {
	err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
		SubTypeList:  []qot.SubType{},
		IsSubOrUnSub: false,
		IsUnsubAll:   true,
	})
	return err
}

// SubscribeSymbols subscribes to real-time market data for multiple symbols in a single request.
func SubscribeSymbols(ctx context.Context, c *Client, market constant.Market, codes []string, subTypes []constant.SubType) error {
	if len(codes) == 0 {
		return fmt.Errorf("SubscribeSymbols: no codes provided")
	}
	if len(subTypes) == 0 {
		return fmt.Errorf("SubscribeSymbols: subTypes is required")
	}

	securities := make([]*qotcommon.Security, len(codes))
	marketPtr := int32(market)
	for i, code := range codes {
		securities[i] = &qotcommon.Security{Market: &marketPtr, Code: &code}
	}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
		SecurityList:     securities,
		SubTypeList:      subTypesConverted,
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
		IsFirstPush:      true,
	})
	return err
}

// UnsubscribeSymbols unsubscribes from real-time market data for multiple symbols.
func UnsubscribeSymbols(ctx context.Context, c *Client, market constant.Market, codes []string, subTypes []constant.SubType) error {
	if len(codes) == 0 {
		return fmt.Errorf("UnsubscribeSymbols: no codes provided")
	}
	if len(subTypes) == 0 {
		return fmt.Errorf("UnsubscribeSymbols: subTypes is required")
	}

	securities := make([]*qotcommon.Security, len(codes))
	marketPtr := int32(market)
	for i, code := range codes {
		securities[i] = &qotcommon.Security{Market: &marketPtr, Code: &code}
	}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
		SecurityList:     securities,
		SubTypeList:      subTypesConverted,
		IsSubOrUnSub:     false,
		IsRegOrUnRegPush: false,
	})
	return err
}

// QuerySubscription queries the current subscription status.
func QuerySubscription(ctx context.Context, c *Client) (*qot.GetSubInfoResponse, error) {
	return qot.GetSubInfo(ctx, c.inner)
}

// RegQotPush registers or unregisters real-time push notifications for a security.
func RegQotPush(ctx context.Context, c *Client, market constant.Market, code string, subTypes []constant.SubType, rehabTypes []constant.RehabType, isReg bool, isFirstPush bool) error {
	if code == "" {
		return fmt.Errorf("RegQotPush: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	subTypesConverted := make([]int32, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = int32(st)
	}
	rehabTypesConverted := make([]int32, len(rehabTypes))
	for i, rt := range rehabTypes {
		rehabTypesConverted[i] = int32(rt)
	}

	err := qot.RegQotPush(ctx, c.inner, &qot.RegQotPushRequest{
		SecurityList:   []*qotcommon.Security{sec},
		SubTypeList:    subTypesConverted,
		RehabTypeList:  rehabTypesConverted,
		IsRegOrUnReg:   isReg,
		IsFirstPush:    isFirstPush,
	})
	return err
}

// GetOrderBook retrieves order book data.
func GetOrderBook(ctx context.Context, c *Client, market constant.Market, code string, num int) (*OrderBookResult, error) {
	if code == "" {
		return nil, fmt.Errorf("GetOrderBook: code is required")
	}
	if num <= 0 {
		return nil, fmt.Errorf("GetOrderBook: num must be greater than 0")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetOrderBook(ctx, c.inner, &qot.GetOrderBookRequest{
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
	return &OrderBookResult{
		Items:    []OrderBook{*book},
		Security: resp.Security,
		Name:     resp.Name,
	}, nil
}

// GetTicker retrieves ticker data.
func GetTicker(ctx context.Context, c *Client, market constant.Market, code string, num int) (*TickerResult, error) {
	if code == "" {
		return nil, fmt.Errorf("GetTicker: code is required")
	}
	if num <= 0 {
		return nil, fmt.Errorf("GetTicker: num must be greater than 0")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetTicker(ctx, c.inner, &qot.GetTickerRequest{
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
			Time:         t.Time,
			Sequence:     t.Sequence,
			Price:        t.Price,
			Volume:       t.Volume,
			Direction:    dir,
			Turnover:     t.Turnover,
			RecvTime:     t.RecvTime,
			Type:         t.Type,
			TypeSign:     t.TypeSign,
			Timestamp:    t.Timestamp,
			PushDataType: t.PushDataType,
		}
	}
	return &TickerResult{
		Items:    tickers,
		Security: resp.Security,
		Name:     resp.Name,
	}, nil
}

// GetRT retrieves real-time data.
func GetRT(ctx context.Context, c *Client, market constant.Market, code string) (*RTResult, error) {
	if code == "" {
		return nil, fmt.Errorf("GetRT: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetRT(ctx, c.inner, &qot.GetRTRequest{Security: sec})
	if err != nil {
		return nil, err
	}

	rtData := make([]RT, len(resp.RTList))
	for i, r := range resp.RTList {
		rtData[i] = RT{
			Time:      r.Time,
			Minute:    r.Minute,
			IsBlank:   r.IsBlank,
			Price:     r.Price,
			Volume:    r.Volume,
			LastClose: r.LastClosePrice,
			AvgPrice:  r.AvgPrice,
			Turnover:  r.Turnover,
			Timestamp: r.Timestamp,
		}
	}
	return &RTResult{
		Items:    rtData,
		Security: resp.Security,
		Name:     resp.Name,
	}, nil
}

// GetBroker retrieves broker data.
func GetBroker(ctx context.Context, c *Client, market constant.Market, code string, num int) (*BrokerResult, error) {
	if code == "" {
		return nil, fmt.Errorf("GetBroker: code is required")
	}
	if num <= 0 {
		return nil, fmt.Errorf("GetBroker: num must be greater than 0")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetBroker(ctx, c.inner, &qot.GetBrokerRequest{
		Security: sec,
	})
	if err != nil {
		return nil, err
	}

	bidBrokers := make([]Broker, len(resp.BidBrokerList))
	for i, b := range resp.BidBrokerList {
		bidBrokers[i] = Broker{ID: b.ID, Name: b.Name, Pos: b.Pos, Volume: b.Volume, OrderID: b.OrderID}
	}
	askBrokers := make([]Broker, len(resp.AskBrokerList))
	for i, b := range resp.AskBrokerList {
		askBrokers[i] = Broker{ID: b.ID, Name: b.Name, Pos: b.Pos, Volume: b.Volume, OrderID: b.OrderID}
	}

	return &BrokerResult{
		Bids:     bidBrokers,
		Asks:     askBrokers,
		Security: resp.Security,
		Name:     resp.Name,
	}, nil
}

// GetStaticInfo retrieves static security info.
func GetStaticInfo(ctx context.Context, c *Client, market constant.Market, code string) ([]StaticInfo, error) {
	if code == "" {
		return nil, fmt.Errorf("GetStaticInfo: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetStaticInfo(ctx, c.inner, &qot.GetStaticInfoRequest{
		Market:       int32(market),
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
		var id int64
		var delisting bool
		var listTimestamp float64
		var exchType int32
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
			if s.Basic.Id != nil {
				id = *s.Basic.Id
			}
			if s.Basic.Delisting != nil {
				delisting = *s.Basic.Delisting
			}
			if s.Basic.ListTimestamp != nil {
				listTimestamp = *s.Basic.ListTimestamp
			}
			if s.Basic.ExchType != nil {
				exchType = *s.Basic.ExchType
			}
		}
		var security *qotcommon.Security
		if s.Basic != nil {
			security = s.Basic.Security
		}
		infos[i] = StaticInfo{
			Code:          code,
			Name:          name,
			Type:          secType,
			ListTime:      listTime,
			LotSize:       lotSize,
			Id:            id,
			Delisting:     delisting,
			ListTimestamp: listTimestamp,
			ExchType:      exchType,
			Security:      security,
			WarrantExData: s.WarrantExData,
			OptionExData:  s.OptionExData,
			FutureExData:  s.FutureExData,
		}
	}
	return infos, nil
}

// GetSecuritySnapshot returns snapshot data for the given securities.
func GetSecuritySnapshot(ctx context.Context, c *Client, securities []*qotcommon.Security) ([]*Snapshot, error) {
	if len(securities) == 0 {
		return nil, fmt.Errorf("GetSecuritySnapshot: securities is required")
	}
	resp, err := qot.GetSecuritySnapshot(ctx, c.inner, &qot.GetSecuritySnapshotRequest{
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
			PreMarket:               mapPreAfterMarketData(basic.PreMarket),
			AfterMarket:             mapPreAfterMarketData(basic.AfterMarket),
			Overnight:               mapPreAfterMarketData(basic.Overnight),
			EquityExData:            s.EquityExData,
			WarrantExData:           s.WarrantExData,
			OptionExData:            s.OptionExData,
			IndexExData:             s.IndexExData,
			PlateExData:             s.PlateExData,
			FutureExData:            s.FutureExData,
			TrustExData:             s.TrustExData,
		})
	}
	return result, nil
}

// GetTradeDate retrieves trade dates.
func GetTradeDate(ctx context.Context, c *Client, market constant.Market, startDate, endDate string) ([]string, error) {
	if startDate == "" {
		return nil, fmt.Errorf("GetTradeDate: startDate is required")
	}
	if endDate == "" {
		return nil, fmt.Errorf("GetTradeDate: endDate is required")
	}
	resp, err := qot.RequestTradeDate(ctx, c.inner, &qot.RequestTradeDateRequest{
		Market:    int32(market),
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

// RequestTradeDate requests trade dates for a specific security.
func RequestTradeDate(ctx context.Context, c *Client, market constant.Market, startDate, endDate string, code string) ([]string, error) {
	if startDate == "" {
		return nil, fmt.Errorf("RequestTradeDate: startDate is required")
	}
	if endDate == "" {
		return nil, fmt.Errorf("RequestTradeDate: endDate is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.RequestTradeDate(ctx, c.inner, &qot.RequestTradeDateRequest{
		Market:    marketPtr,
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

// GetFutureInfo retrieves futures information.
func GetFutureInfo(ctx context.Context, c *Client, code string) ([]FutureInfo, error) {
	if code == "" {
		return nil, fmt.Errorf("GetFutureInfo: code is required")
	}
	marketPtr := int32(2) // HK Future
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetFutureInfo(ctx, c.inner, &qot.GetFutureInfoRequest{
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
			Security:           f.Security,
			Origin:             f.Origin,
			TradeTimeList: func() []TradeTime {
				if len(f.TradeTimeList) == 0 {
					return nil
				}
				tt := make([]TradeTime, len(f.TradeTimeList))
				for j, t := range f.TradeTimeList {
					tt[j] = TradeTime{Begin: getFloat64(t.Begin), End: getFloat64(t.End)}
				}
				return tt
			}(),
		}
	}
	return infos, nil
}

// GetPlateSet retrieves plate set (板块) list.
func GetPlateSet(ctx context.Context, c *Client, market constant.Market) ([]Plate, error) {
	if !market.IsValid() {
		return nil, fmt.Errorf("GetPlateSet: invalid market")
	}
	marketPtr := int32(market)
	resp, err := qot.GetPlateSet(ctx, c.inner, &qot.GetPlateSetRequest{Market: marketPtr})
	if err != nil {
		return nil, err
	}

	plates := make([]Plate, len(resp.PlateSetList))
	for i, p := range resp.PlateSetList {
		code := ""
		if p.Plate != nil && p.Plate.Code != nil {
			code = *p.Plate.Code
		}
		plates[i] = Plate{Code: code, Name: p.Name, PlateType: p.PlateType}
	}
	return plates, nil
}

// GetIpoList retrieves IPO list.
func GetIpoList(ctx context.Context, c *Client, market constant.Market) ([]IpoData, error) {
	if !market.IsValid() {
		return nil, fmt.Errorf("GetIpoList: invalid market")
	}
	marketPtr := int32(market)
	resp, err := qot.GetIpoList(ctx, c.inner, &qot.GetIpoListRequest{Market: marketPtr})
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
		ipo := IpoData{
			Code:          code,
			Name:          ip.Basic.Name,
			ListDate:      ip.Basic.ListTime,
			ListTimestamp: ip.Basic.ListTimestamp,
		}
		if ip.CnExData != nil {
			d := ip.CnExData
			ipo.CnExData = &qotgetipolist.CNIpoExData{
				ApplyCode:              &d.ApplyCode,
				IssueSize:              &d.IssueSize,
				OnlineIssueSize:        &d.OnlineIssueSize,
				ApplyUpperLimit:        &d.ApplyUpperLimit,
				ApplyLimitMarketValue:  &d.ApplyLimitMarketValue,
				IsEstimateIpoPrice:     &d.IsEstimateIpoPrice,
				IpoPrice:               &d.IpoPrice,
				IndustryPeRate:         &d.IndustryPeRate,
				IsEstimateWinningRatio: &d.IsEstimateWinningRatio,
				WinningRatio:           &d.WinningRatio,
				IssuePeRate:            &d.IssuePeRate,
				IsHasWon:               &d.IsHasWon,
				WinningNumData:         d.WinningNumDataList,
			}
			if d.ApplyTime != "" {
				ipo.CnExData.ApplyTime = &d.ApplyTime
			}
			if d.ApplyTimestamp != 0 {
				ipo.CnExData.ApplyTimestamp = &d.ApplyTimestamp
			}
			if d.WinningTime != "" {
				ipo.CnExData.WinningTime = &d.WinningTime
			}
			if d.WinningTimestamp != 0 {
				ipo.CnExData.WinningTimestamp = &d.WinningTimestamp
			}
		}
		if ip.HkExData != nil {
			d := ip.HkExData
			ipo.HkExData = &qotgetipolist.HKIpoExData{
				IpoPriceMin:       &d.IpoPriceMin,
				IpoPriceMax:       &d.IpoPriceMax,
				ListPrice:         &d.ListPrice,
				LotSize:           &d.LotSize,
				EntrancePrice:     &d.EntrancePrice,
				IsSubscribeStatus: &d.IsSubscribeStatus,
			}
		}
		if ip.UsExData != nil {
			d := ip.UsExData
			ipo.UsExData = &qotgetipolist.USIpoExData{
				IpoPriceMin: &d.IpoPriceMin,
				IpoPriceMax: &d.IpoPriceMax,
				IssueSize:   &d.IssueSize,
			}
		}
		ipos = append(ipos, ipo)
	}
	return ipos, nil
}

// GetUserSecurityGroup retrieves user security group list.
func GetUserSecurityGroup(ctx context.Context, c *Client) ([]UserSecurityGroup, error) {
	resp, err := qot.GetUserSecurityGroup(ctx, c.inner, &qot.GetUserSecurityGroupRequest{})
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
func GetUserSecurity(ctx context.Context, c *Client, groupName string) ([]StaticInfo, error) {
	if groupName == "" {
		return nil, fmt.Errorf("GetUserSecurity: groupName is required")
	}
	resp, err := qot.GetUserSecurity(ctx, c.inner, groupName)
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
		listTime := ""
		if s.Basic.ListTime != nil {
			listTime = *s.Basic.ListTime
		}
		lotSize := int32(0)
		if s.Basic.LotSize != nil {
			lotSize = *s.Basic.LotSize
		}
		id := int64(0)
		if s.Basic.Id != nil {
			id = *s.Basic.Id
		}
		delisting := false
		if s.Basic.Delisting != nil {
			delisting = *s.Basic.Delisting
		}
		listTimestamp := float64(0)
		if s.Basic.ListTimestamp != nil {
			listTimestamp = *s.Basic.ListTimestamp
		}
		exchType := int32(0)
		if s.Basic.ExchType != nil {
			exchType = *s.Basic.ExchType
		}
		infos = append(infos, StaticInfo{
			Code:          code,
			Name:          name,
			Type:          secType,
			ListTime:      listTime,
			LotSize:       lotSize,
			Id:            id,
			Delisting:     delisting,
			ListTimestamp: listTimestamp,
			ExchType:      exchType,
			Security:      s.Basic.Security,
			WarrantExData: s.WarrantExData,
			OptionExData:  s.OptionExData,
			FutureExData:  s.FutureExData,
		})
	}
	return infos, nil
}

// GetMarketState retrieves market state (trading status).
func GetMarketState(ctx context.Context, c *Client, market constant.Market, code string) (*MarketStateResult, error) {
	if code == "" {
		return nil, fmt.Errorf("GetMarketState: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetMarketState(ctx, c.inner, &qot.GetMarketStateRequest{
		SecurityList: []*qotcommon.Security{sec},
	})
	if err != nil {
		return nil, err
	}

	if len(resp.MarketInfoList) == 0 {
		return &MarketStateResult{Code: code}, nil
	}

	info := resp.MarketInfoList[0]
	return &MarketStateResult{
		Code:  code,
		Name:  info.Name,
		State: info.MarketState,
	}, nil
}

// GetCapitalFlow retrieves capital flow data.
func GetCapitalFlow(ctx context.Context, c *Client, market constant.Market, code string, periodType ...constant.CapitalFlowPeriodType) (*CapitalFlowResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetCapitalFlow: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	period := int32(0)
	if len(periodType) > 0 {
		period = int32(periodType[0])
	}

	resp, err := qot.GetCapitalFlow(ctx, c.inner, &qot.GetCapitalFlowRequest{
		Security:   sec,
		PeriodType: period,
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
	return &CapitalFlowResponse{
		Items:              flows,
		LastValidTime:      resp.LastValidTime,
		LastValidTimestamp: resp.LastValidTimestamp,
	}, nil
}

// GetCapitalDistribution retrieves capital distribution.
func GetCapitalDistribution(ctx context.Context, c *Client, market constant.Market, code string) (*CapitalDistribution, error) {
	if code == "" {
		return nil, fmt.Errorf("GetCapitalDistribution: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetCapitalDistribution(ctx, c.inner, sec)
	if err != nil {
		return nil, err
	}

	if resp.CapitalDistribution == nil {
		return &CapitalDistribution{}, nil
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
func GetOwnerPlate(ctx context.Context, c *Client, market constant.Market, code string) (map[string]*OwnerPlateEntry, error) {
	if code == "" {
		return nil, fmt.Errorf("GetOwnerPlate: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetOwnerPlate(ctx, c.inner, &qot.GetOwnerPlateRequest{
		SecurityList: []*qotcommon.Security{sec},
	})
	if err != nil {
		return nil, err
	}

	result := make(map[string]*OwnerPlateEntry)
	for _, securityPlate := range resp.OwnerPlateList {
		sec := securityPlate.Security
		code := getStr(sec.Code)
		entry := &OwnerPlateEntry{Name: getStr(securityPlate.Name)}
		for _, plateInfo := range securityPlate.PlateInfoList {
			entry.Plates = append(entry.Plates, &OwnerPlateInfo{
				Code:      getStr(plateInfo.Plate.Code),
				Name:      getStr(plateInfo.Name),
				PlateType: getInt32(plateInfo.PlateType),
			})
		}
		result[code] = entry
	}
	return result, nil
}

// RequestHistoryKL requests historical K-line data with automatic pagination.
func RequestHistoryKL(ctx context.Context, c *Client, market constant.Market, code string, klType constant.KLType, startDate, endDate string) (*KLinesResult, error) {
	if code == "" {
		return nil, fmt.Errorf("RequestHistoryKL: code is required")
	}
	if startDate == "" {
		return nil, fmt.Errorf("RequestHistoryKL: startDate is required")
	}
	if endDate == "" {
		return nil, fmt.Errorf("RequestHistoryKL: endDate is required")
	}
	return RequestHistoryKLWithLimit(ctx, c, market, code, klType, startDate, endDate, DefaultHistoryKLPageSize)
}

// RequestHistoryKLWithLimit requests historical K-line data with a configurable page size.
func RequestHistoryKLWithLimit(ctx context.Context, c *Client, market constant.Market, code string, klType constant.KLType, startDate, endDate string, maxPerPage int32) (*KLinesResult, error) {
	if code == "" {
		return nil, fmt.Errorf("RequestHistoryKLWithLimit: code is required")
	}
	if startDate == "" {
		return nil, fmt.Errorf("RequestHistoryKLWithLimit: startDate is required")
	}
	if endDate == "" {
		return nil, fmt.Errorf("RequestHistoryKLWithLimit: endDate is required")
	}
	if maxPerPage <= 0 {
		return nil, fmt.Errorf("RequestHistoryKLWithLimit: maxPerPage must be greater than 0")
	}
	marketPtr := int32(market)
	klTypePtr := int32(klType)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	var allKLines []KLine
	var nextReqKey []byte
	var security *qotcommon.Security
	var name string

	for {
		resp, err := qot.RequestHistoryKL(ctx, c.inner, &qot.RequestHistoryKLRequest{
			Security:    sec,
			KlType:      klTypePtr,
			BeginTime:   startDate,
			EndTime:     endDate,
			MaxAckKLNum: maxPerPage,
			NextReqKey:  nextReqKey,
		})
		if err != nil {
			return nil, err
		}

		if security == nil && resp.Security != nil {
			security = resp.Security
		}
		if name == "" && resp.Name != "" {
			name = resp.Name
		}

		for _, kl := range resp.KLList {
			allKLines = append(allKLines, KLine{
				Time:         kl.Time,
				IsBlank:      kl.IsBlank,
				Open:         kl.OpenPrice,
				High:         kl.HighPrice,
				Low:          kl.LowPrice,
				Close:        kl.ClosePrice,
				Volume:       kl.Volume,
				LastClose:    kl.LastClosePrice,
				Turnover:     kl.Turnover,
				TurnoverRate: kl.TurnoverRate,
				Pe:           kl.Pe,
				ChangeRate:   kl.ChangeRate,
				Timestamp:    kl.Timestamp,
			})
		}

		if len(resp.NextReqKey) == 0 {
			break
		}
		nextReqKey = resp.NextReqKey

		time.Sleep(HistoryKLPaginationDelay)
	}

	return &KLinesResult{
		Items:    allKLines,
		Security: security,
		Name:     name,
	}, nil
}

// GetHistoryKL requests historical K-line data.
func GetHistoryKL(ctx context.Context, c *Client, market constant.Market, code string, klType constant.KLType, rehabType constant.RehabType, startDate, endDate string, maxNum int32) (*KLinesResult, error) {
	if code == "" {
		return nil, fmt.Errorf("GetHistoryKL: code is required")
	}
	if startDate == "" {
		return nil, fmt.Errorf("GetHistoryKL: startDate is required")
	}
	if endDate == "" {
		return nil, fmt.Errorf("GetHistoryKL: endDate is required")
	}
	if maxNum <= 0 {
		return nil, fmt.Errorf("GetHistoryKL: maxNum must be greater than 0")
	}
	marketPtr := int32(market)
	klTypePtr := int32(klType)
	rehabTypePtr := int32(rehabType)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetHistoryKL(ctx, c.inner, &qot.GetHistoryKLRequest{
		RehabType:   rehabTypePtr,
		KlType:      klTypePtr,
		Security:    sec,
		BeginTime:   startDate,
		EndTime:     endDate,
		MaxAckKLNum: maxNum,
	})
	if err != nil {
		return nil, err
	}

	klines := make([]KLine, len(resp.KLList))
	for i, kl := range resp.KLList {
		klines[i] = KLine{
			Time:         kl.Time,
			IsBlank:      kl.IsBlank,
			Open:         kl.OpenPrice,
			High:         kl.HighPrice,
			Low:          kl.LowPrice,
			Close:        kl.ClosePrice,
			Volume:       kl.Volume,
			LastClose:    kl.LastClosePrice,
			Turnover:     kl.Turnover,
			TurnoverRate: kl.TurnoverRate,
			Pe:           kl.Pe,
			ChangeRate:   kl.ChangeRate,
			Timestamp:    kl.Timestamp,
		}
	}

	return &KLinesResult{
		Items:           klines,
		NextKLTime:      resp.NextKLTime,
		NextKLTimestamp: resp.NextKLTimestamp,
	}, nil
}

// GetReference retrieves related/reference securities.
func GetReference(ctx context.Context, c *Client, market constant.Market, code string, refType qotgetreference.ReferenceType) ([]StaticInfo, error) {
	if code == "" {
		return nil, fmt.Errorf("GetReference: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetReference(ctx, c.inner, &qot.GetReferenceRequest{
		Security:      sec,
		ReferenceType: int32(refType),
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
			Code:          code,
			Name:          name,
			Type:          secType,
			Security:      s.Basic.Security,
			WarrantExData: s.WarrantExData,
			OptionExData:  s.OptionExData,
			FutureExData:  s.FutureExData,
		})
	}
	return infos, nil
}

// GetPlateSecurity retrieves securities in a plate.
func GetPlateSecurity(ctx context.Context, c *Client, market constant.Market, plateCode string) ([]StaticInfo, error) {
	if plateCode == "" {
		return nil, fmt.Errorf("GetPlateSecurity: plateCode is required")
	}
	marketPtr := int32(market)
	plate := &qotcommon.Security{Market: &marketPtr, Code: &plateCode}

	resp, err := qot.GetPlateSecurity(ctx, c.inner, &qot.GetPlateSecurityRequest{Plate: plate})
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
		listTime := ""
		if s.Basic.ListTime != nil {
			listTime = *s.Basic.ListTime
		}
		lotSize := int32(0)
		if s.Basic.LotSize != nil {
			lotSize = *s.Basic.LotSize
		}
		id := int64(0)
		if s.Basic.Id != nil {
			id = *s.Basic.Id
		}
		delisting := false
		if s.Basic.Delisting != nil {
			delisting = *s.Basic.Delisting
		}
		listTimestamp := float64(0)
		if s.Basic.ListTimestamp != nil {
			listTimestamp = *s.Basic.ListTimestamp
		}
		exchType := int32(0)
		if s.Basic.ExchType != nil {
			exchType = *s.Basic.ExchType
		}
		infos = append(infos, StaticInfo{
			Code:          code,
			Name:          name,
			Type:          secType,
			ListTime:      listTime,
			LotSize:       lotSize,
			Id:            id,
			Delisting:     delisting,
			ListTimestamp: listTimestamp,
			ExchType:      exchType,
			Security:      s.Basic.Security,
			WarrantExData: s.WarrantExData,
			OptionExData:  s.OptionExData,
			FutureExData:  s.FutureExData,
		})
	}
	return infos, nil
}

// GetOptionExpirationDate retrieves option expiration dates.
func GetOptionExpirationDate(ctx context.Context, c *Client, market constant.Market, code string) ([]OptionExpiration, error) {
	if code == "" {
		return nil, fmt.Errorf("GetOptionExpirationDate: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetOptionExpirationDate(ctx, c.inner, &qot.GetOptionExpirationDateRequest{
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
			Date:           e.StrikeTime,
			Days:           e.OptionExpiryDateDistance,
			Desc:           fmt.Sprintf("Cycle %d", e.Cycle),
			StrikeTimestamp: e.StrikeTimestamp,
			Cycle:          e.Cycle,
		})
	}
	return expirations, nil
}

// ModifyUserSecurity adds/removes securities from user group.
func ModifyUserSecurity(ctx context.Context, c *Client, groupName string, op int32, market constant.Market, codes []string) error {
	if groupName == "" {
		return fmt.Errorf("ModifyUserSecurity: groupName is required")
	}
	if len(codes) == 0 {
		return fmt.Errorf("ModifyUserSecurity: codes is required")
	}
	marketPtr := int32(market)
	securities := make([]*qotcommon.Security, len(codes))
	for i, code := range codes {
		securities[i] = &qotcommon.Security{Market: &marketPtr, Code: &code}
	}

	_, err := qot.ModifyUserSecurity(ctx, c.inner, &qot.ModifyUserSecurityRequest{
		GroupName:    groupName,
		Op:           op,
		SecurityList: securities,
	})
	return err
}

// GetSubInfo retrieves subscription info.
func GetSubInfo(ctx context.Context, c *Client) (*SubInfo, error) {
	resp, err := qot.GetSubInfo(ctx, c.inner)
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
		IsSub:          len(resp.ConnSubInfoList) > 0,
		SubTypes:       types,
		Security:       fmt.Sprintf("Used: %d, Remain: %d", resp.TotalUsedQuota, resp.RemainQuota),
		TotalUsedQuota: resp.TotalUsedQuota,
		RemainQuota:    resp.RemainQuota,
	}, nil
}

// SetPriceReminder creates/updates/deletes a price reminder.
func SetPriceReminder(ctx context.Context, c *Client, market constant.Market, code string, op constant.PriceReminderOp, reminderType constant.PriceReminderType, freq constant.PriceReminderFreq, value float64, note string) (int64, error) {
	if code == "" {
		return 0, fmt.Errorf("SetPriceReminder: code is required")
	}
	if op == constant.PriceReminderOp_Add && value <= 0 {
		return 0, fmt.Errorf("SetPriceReminder: value must be greater than 0 for Set operation")
	}
	marketPtr := int32(market)
	reminderTypePtr := int32(reminderType)
	freqPtr := int32(freq)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	resp, err := qot.SetPriceReminder(ctx, c.inner, &qot.SetPriceReminderRequest{
		Security: sec,
		Op:       int32(op),
		Type:     reminderTypePtr,
		Freq:     freqPtr,
		Value:    value,
		Note:     note,
	})
	if err != nil {
		return 0, err
	}
	return resp.Key, nil
}

// GetPriceReminder retrieves price reminders for a security.
func GetPriceReminder(ctx context.Context, c *Client, market constant.Market, code string) ([]*PriceReminderInfo, error) {
	if code == "" {
		return nil, fmt.Errorf("GetPriceReminder: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	resp, err := qot.GetPriceReminder(ctx, c.inner, sec, marketPtr)
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

// GetSuspend retrieves suspension information for securities.
func GetSuspend(ctx context.Context, c *Client, securities []*qotcommon.Security, beginTime, endTime string) ([]*SuspendInfo, error) {
	if len(securities) == 0 {
		return nil, fmt.Errorf("GetSuspend: securities is required")
	}
	if beginTime == "" {
		return nil, fmt.Errorf("GetSuspend: beginTime is required")
	}
	if endTime == "" {
		return nil, fmt.Errorf("GetSuspend: endTime is required")
	}
	resp, err := qot.GetSuspend(ctx, c.inner, &qot.GetSuspendRequest{
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

// GetCodeChange returns code change information for the given securities.
func GetCodeChange(ctx context.Context, c *Client, securities []*qotcommon.Security) ([]*CodeChangeInfo, error) {
	if len(securities) == 0 {
		return nil, fmt.Errorf("GetCodeChange: securities is required")
	}
	resp, err := qot.GetCodeChange(ctx, c.inner, &qot.GetCodeChangeRequest{
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
			Type:               cc.Type,
			Security:           cc.Security,
			RelatedSecurity:    cc.RelatedSecurity,
			PublicTime:         cc.PublicTime,
			PublicTimestamp:    cc.PublicTimestamp,
			EffectiveTime:      cc.EffectiveTime,
			EffectiveTimestamp: cc.EffectiveTimestamp,
			EndTime:            cc.EndTime,
			EndTimestamp:       cc.EndTimestamp,
		})
	}
	return result, nil
}

// GetHoldingChangeList retrieves holding change list.
func GetHoldingChangeList(ctx context.Context, c *Client, market constant.Market, code string, holderCategory constant.HolderCategory, beginTime, endTime string) ([]*HoldingChangeInfo, error) {
	if code == "" {
		return nil, fmt.Errorf("GetHoldingChangeList: code is required")
	}
	if beginTime == "" {
		return nil, fmt.Errorf("GetHoldingChangeList: beginTime is required")
	}
	if endTime == "" {
		return nil, fmt.Errorf("GetHoldingChangeList: endTime is required")
	}
	marketPtr := int32(market)
	holderCatPtr := int32(holderCategory)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	resp, err := qot.GetHoldingChangeList(ctx, c.inner, &qot.GetHoldingChangeListRequest{
		Security:       sec,
		HolderCategory: holderCatPtr,
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

// StockFilter filters stocks based on basic criteria.
func StockFilter(ctx context.Context, c *Client, market constant.Market, begin, num int32) ([]*StockFilterResult, error) {
	if begin < 0 {
		return nil, fmt.Errorf("StockFilter: begin must be >= 0")
	}
	if num <= 0 {
		return nil, fmt.Errorf("StockFilter: num must be greater than 0")
	}
	marketPtr := int32(market)
	resp, err := qot.StockFilter(ctx, c.inner, &qot.StockFilterRequest{
		Market: marketPtr,
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
			Security:                d.Security,
			Name:                    d.Name,
			BaseDataList:            d.BaseDataList,
			AccumulateDataList:      d.AccumulateDataList,
			FinancialDataList:       d.FinancialDataList,
			CustomIndicatorDataList: d.CustomIndicatorDataList,
		}
		for _, base := range d.BaseDataList {
			if base == nil {
				continue
			}
			fieldName := getInt32(base.FieldName)
			value := getFloat64(base.Value)
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
func GetOptionChain(ctx context.Context, c *Client, market constant.Market, code string, indexOptionType constant.IndexOptionType, optType constant.OptionType, condition int32, beginTime, endTime string) ([]*OptChain, error) {
	if code == "" {
		return nil, fmt.Errorf("GetOptionChain: code is required")
	}
	if beginTime == "" {
		beginTime = time.Now().Format("2006-01-02")
	}
	if endTime == "" {
		endTime = time.Now().AddDate(0, 1, 0).Format("2006-01-02")
	}
	marketPtr := int32(market)
	indexOptPtr := int32(indexOptionType)
	optTypePtr := int32(optType)
	owner := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetOptionChain(ctx, c.inner, &qot.GetOptionChainRequest{
		Owner:           owner,
		IndexOptionType: indexOptPtr,
		Type:            optTypePtr,
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

// GetWarrant returns the list of warrants for the given underlying security.
func GetWarrant(ctx context.Context, c *Client, market constant.Market, code string, begin, num int32, sortField constant.WarrantSortField, ascend bool, optType constant.WarrantType, issuer qotcommon.Issuer, status constant.WarrantStatus) (*WarrantResult, error) {
	if code == "" {
		return nil, fmt.Errorf("GetWarrant: code is required")
	}
	if begin < 0 {
		return nil, fmt.Errorf("GetWarrant: begin must be >= 0")
	}
	if num <= 0 {
		return nil, fmt.Errorf("GetWarrant: num must be greater than 0")
	}
	marketPtr := int32(market)
	owner := &qotcommon.Security{Market: &marketPtr, Code: &code}
	optTypePtr := int32(optType)
	sortFieldPtr := int32(sortField)
	issuerPtr := int32(issuer)
	statusPtr := int32(status)

	resp, err := qot.GetWarrant(ctx, c.inner, &qot.GetWarrantRequest{
		Begin:      begin,
		Num:        num,
		SortField:  sortFieldPtr,
		Ascend:     ascend,
		Owner:      owner,
		TypeList:   []int32{optTypePtr},
		IssuerList: []int32{issuerPtr},
		Status:     statusPtr,
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
	return &WarrantResult{
		Items:    result,
		LastPage: resp.LastPage,
		AllCount: resp.AllCount,
	}, nil
}

// RequestRehab requests rehabilitation (复权) data.
func RequestRehab(ctx context.Context, c *Client, market constant.Market, code string) ([]*RehabInfo, error) {
	if code == "" {
		return nil, fmt.Errorf("RequestRehab: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	resp, err := qot.RequestRehab(ctx, c.inner, &qot.RequestRehabRequest{
		Security: sec,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*RehabInfo, 0, len(resp.RehabList))
	for _, r := range resp.RehabList {
		result = append(result, mapRehabInfo(r))
	}
	return result, nil
}

// GetRehab returns rehabilitation (复权) data for the given security.
// Unlike RequestRehab (which may use cached data), GetRehab always fetches
// fresh rehabilitation data from the server.
// Deprecated: Removed in Futu v10.6 proto — proto package qotgetrehab no longer exists.
// Use RequestRehab instead.
func GetRehab(ctx context.Context, c *Client, market constant.Market, code string) ([]*RehabInfo, error) {
	_ = ctx
	_ = c
	_ = market
	_ = code
	return nil, fmt.Errorf("GetRehab: removed in Futu v10.6 — use RequestRehab instead")
}

// GetHistoryKLPoints retrieves K-line data at specific time points for a security.
func GetHistoryKLPoints(ctx context.Context, c *Client, market constant.Market, code string, times []string, klType constant.KLType, rehabType constant.RehabType, noDataMode qot.NoDataMode) (*qot.GetHistoryKLPointsResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetHistoryKLPoints: code is required")
	}
	if len(times) == 0 {
		return nil, fmt.Errorf("GetHistoryKLPoints: times is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetHistoryKLPoints(ctx, c.inner, &qot.GetHistoryKLPointsRequest{
		RehabType:  rehabType,
		KLType:    klType,
		NoDataMode: noDataMode,
		Securities: []*qotcommon.Security{sec},
		Times:     times,
	})
}

// GetTradeDates retrieves trade dates for a market within a date range.
func GetTradeDates(ctx context.Context, c *Client, market int32, beginTime, endTime string) ([]TradeDate, error) {
	if beginTime == "" {
		return nil, fmt.Errorf("GetTradeDates: beginTime is required")
	}
	if endTime == "" {
		return nil, fmt.Errorf("GetTradeDates: endTime is required")
	}
	resp, err := qot.RequestTradeDate(ctx, c.inner, &qot.RequestTradeDateRequest{
		Market:    market,
		BeginTime: beginTime,
		EndTime:   endTime,
	})
	if err != nil {
		return nil, err
	}
	dates := make([]TradeDate, 0, len(resp.TradeDateList))
	for _, d := range resp.TradeDateList {
		if d == nil {
			continue
		}
		dates = append(dates, TradeDate{
			Time:          util.ProtoStr(d.Time),
			Timestamp:     util.ProtoFloat64(d.Timestamp),
			TradeDateType: util.ProtoInt32(d.TradeDateType),
		})
	}
	return dates, nil
}

func GetFinancialsStatements(ctx context.Context, c *Client, market constant.Market, code string, statementType, financialType int32, currencyCode, nextKey string, num int32) (*qot.GetFinancialsStatementsResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetFinancialsStatements: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetFinancialsStatements(ctx, c.inner, &qot.GetFinancialsStatementsRequest{
		Security:       sec,
		StatementType:  statementType,
		FinancialType:  financialType,
		CurrencyCode:   currencyCode,
		NextKey:        nextKey,
		Num:            num,
	})
}

func GetFinancialsRevenueBreakdown(ctx context.Context, c *Client, market constant.Market, code string, date uint32, financialType int32, currencyCode string) (*qot.GetFinancialsRevenueBreakdownResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetFinancialsRevenueBreakdown: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetFinancialsRevenueBreakdown(ctx, c.inner, &qot.GetFinancialsRevenueBreakdownRequest{
		Security:      sec,
		Date:          date,
		FinancialType: financialType,
		CurrencyCode:  currencyCode,
	})
}

func GetResearchAnalystConsensus(ctx context.Context, c *Client, market constant.Market, code string) (*qot.GetResearchAnalystConsensusResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetResearchAnalystConsensus: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetResearchAnalystConsensus(ctx, c.inner, &qot.GetResearchAnalystConsensusRequest{Security: sec})
}

func GetResearchRatingSummary(ctx context.Context, c *Client, market constant.Market, code string, ratingDimensionType qotcommon.ResearchRatingDimensionType, uid string, nextKey string, num int32) (*qot.GetResearchRatingSummaryResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetResearchRatingSummary: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetResearchRatingSummary(ctx, c.inner, &qot.GetResearchRatingSummaryRequest{
		Security:            sec,
		RatingDimensionType: &ratingDimensionType,
		Uid:                 uid,
		NextKey:             nextKey,
		Num:                 num,
	})
}

func GetResearchMorningstarReport(ctx context.Context, c *Client, market constant.Market, code string) (*qot.GetResearchMorningstarReportResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetResearchMorningstarReport: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetResearchMorningstarReport(ctx, c.inner, &qot.GetResearchMorningstarReportRequest{Security: sec})
}

func GetValuationDetail(ctx context.Context, c *Client, market constant.Market, code string, valuationType, intervalType int32) (*qot.GetValuationDetailResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetValuationDetail: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetValuationDetail(ctx, c.inner, &qot.GetValuationDetailRequest{
		Security:      sec,
		ValuationType: valuationType,
		IntervalType:  intervalType,
	})
}

func GetValuationPlateStockList(ctx context.Context, c *Client, market constant.Market, code string, filterMarket constant.Market, filterCode string, valuationType int32, nextKey string, num, sortType, sortId int32) (*qot.GetValuationPlateStockListResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetValuationPlateStockList: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	req := &qot.GetValuationPlateStockListRequest{
		Security:      sec,
		ValuationType: valuationType,
		NextKey:       nextKey,
		Num:           num,
		SortType:      sortType,
		SortId:        sortId,
	}
	if filterCode != "" {
		filterMarketPtr := int32(filterMarket)
		req.FilterSecurity = &qotcommon.Security{Market: &filterMarketPtr, Code: &filterCode}
	}
	return qot.GetValuationPlateStockList(ctx, c.inner, req)
}

func GetCorporateActionsDividends(ctx context.Context, c *Client, market constant.Market, code string) (*qot.GetCorporateActionsDividendsResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetCorporateActionsDividends: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetCorporateActionsDividends(ctx, c.inner, &qot.GetCorporateActionsDividendsRequest{Security: sec})
}

func GetCorporateActionsBuybacks(ctx context.Context, c *Client, market constant.Market, code string, nextKey string, num int32) (*qot.GetCorporateActionsBuybacksResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetCorporateActionsBuybacks: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetCorporateActionsBuybacks(ctx, c.inner, &qot.GetCorporateActionsBuybacksRequest{
		Security: sec,
		NextKey:  nextKey,
		Num:      num,
	})
}

func GetCorporateActionsStockSplits(ctx context.Context, c *Client, market constant.Market, code string, nextKey string, num int32) (*qot.GetCorporateActionsStockSplitsResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetCorporateActionsStockSplits: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetCorporateActionsStockSplits(ctx, c.inner, &qot.GetCorporateActionsStockSplitsRequest{
		Security: sec,
		NextKey:  nextKey,
		Num:      num,
	})
}

func GetShareholdersOverview(ctx context.Context, c *Client, market constant.Market, code string, periodId int32) (*qot.GetShareholdersOverviewResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetShareholdersOverview: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetShareholdersOverview(ctx, c.inner, &qot.GetShareholdersOverviewRequest{Security: sec, PeriodId: periodId})
}

func GetShareholdersHoldingChanges(ctx context.Context, c *Client, market constant.Market, code string, nextKey string, num, sortType, sortColumn, filterType int32) (*qot.GetShareholdersHoldingChangesResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetShareholdersHoldingChanges: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetShareholdersHoldingChanges(ctx, c.inner, &qot.GetShareholdersHoldingChangesRequest{
		Security:   sec,
		NextKey:    nextKey,
		Num:        num,
		SortType:   sortType,
		SortColumn: sortColumn,
		FilterType: filterType,
	})
}

func GetShareholdersHolderDetail(ctx context.Context, c *Client, market constant.Market, code string, requestType int32, nextKey string, num, sortColumn, sortType, periodId, holderId int32) (*qot.GetShareholdersHolderDetailResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetShareholdersHolderDetail: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetShareholdersHolderDetail(ctx, c.inner, &qot.GetShareholdersHolderDetailRequest{
		Security:    sec,
		RequestType: requestType,
		NextKey:     nextKey,
		Num:         num,
		SortColumn:  sortColumn,
		SortType:    sortType,
		PeriodId:    periodId,
		HolderId:    holderId,
	})
}

func GetShareholdersInstitutional(ctx context.Context, c *Client, market constant.Market, code string, nextKey string, num int32) (*qot.GetShareholdersInstitutionalResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetShareholdersInstitutional: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetShareholdersInstitutional(ctx, c.inner, &qot.GetShareholdersInstitutionalRequest{
		Security: sec,
		NextKey:  nextKey,
		Num:      num,
	})
}

func GetInsiderHolderList(ctx context.Context, c *Client, market constant.Market, code string, nextKey string, num int32) (*qot.GetInsiderHolderListResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetInsiderHolderList: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetInsiderHolderList(ctx, c.inner, &qot.GetInsiderHolderListRequest{
		Security: sec,
		NextKey:  nextKey,
		Num:      num,
	})
}

func GetInsiderTradeList(ctx context.Context, c *Client, market constant.Market, code string, holderId int64, nextKey string, num int32) (*qot.GetInsiderTradeListResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetInsiderTradeList: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetInsiderTradeList(ctx, c.inner, &qot.GetInsiderTradeListRequest{
		Security: sec,
		HolderId: holderId,
		NextKey:  nextKey,
		Num:      num,
	})
}

func GetCompanyProfile(ctx context.Context, c *Client, market constant.Market, code string) (*qot.GetCompanyProfileResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetCompanyProfile: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetCompanyProfile(ctx, c.inner, &qot.GetCompanyProfileRequest{Security: sec})
}

func GetCompanyExecutives(ctx context.Context, c *Client, market constant.Market, code string) (*qot.GetCompanyExecutivesResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetCompanyExecutives: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetCompanyExecutives(ctx, c.inner, &qot.GetCompanyExecutivesRequest{Security: sec})
}

func GetCompanyExecutiveBackground(ctx context.Context, c *Client, market constant.Market, code string, leaderName string) (*qot.GetCompanyExecutiveBackgroundResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetCompanyExecutiveBackground: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetCompanyExecutiveBackground(ctx, c.inner, &qot.GetCompanyExecutiveBackgroundRequest{
		Security:   sec,
		LeaderName: leaderName,
	})
}

func GetCompanyOperationalEfficiency(ctx context.Context, c *Client, market constant.Market, code string, nextKey string, num int32, currencyCode string, financialType int32) (*qot.GetCompanyOperationalEfficiencyResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetCompanyOperationalEfficiency: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetCompanyOperationalEfficiency(ctx, c.inner, &qot.GetCompanyOperationalEfficiencyRequest{
		Security:      sec,
		NextKey:       nextKey,
		Num:           num,
		CurrencyCode:  currencyCode,
		FinancialType:  financialType,
	})
}

func GetTopTenBuySellBrokers(ctx context.Context, c *Client, market constant.Market, code string, daysBefore int32) (*qot.GetTopTenBuySellBrokersResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetTopTenBuySellBrokers: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetTopTenBuySellBrokers(ctx, c.inner, &qot.GetTopTenBuySellBrokersRequest{
		Security:   sec,
		DaysBefore: daysBefore,
	})
}

func GetDailyShortVolume(ctx context.Context, c *Client, market constant.Market, code string, nextKey string, num int32) (*qot.GetDailyShortVolumeResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetDailyShortVolume: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetDailyShortVolume(ctx, c.inner, &qot.GetDailyShortVolumeRequest{
		Security: sec,
		NextKey:  nextKey,
		Num:      num,
	})
}

func GetShortInterest(ctx context.Context, c *Client, market constant.Market, code string, nextKey string, num int32) (*qot.GetShortInterestResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetShortInterest: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetShortInterest(ctx, c.inner, &qot.GetShortInterestRequest{
		Security: sec,
		NextKey:  nextKey,
		Num:      num,
	})
}

func GetOptionVolatility(ctx context.Context, c *Client, market constant.Market, code string, queryTimePeriod qotcommon.OptionVolatilityTimePeriodType, hvTimePeriod int32) (*qot.GetOptionVolatilityResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetOptionVolatility: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetOptionVolatility(ctx, c.inner, &qot.GetOptionVolatilityRequest{
		Security:         sec,
		QueryTimePeriod:  queryTimePeriod,
		HvTimePeriod:     hvTimePeriod,
	})
}

func GetOptionExerciseProbability(ctx context.Context, c *Client, market constant.Market, code string) (*qot.GetOptionExerciseProbabilityResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetOptionExerciseProbability: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetOptionExerciseProbability(ctx, c.inner, &qot.GetOptionExerciseProbabilityRequest{Security: sec})
}

func GetHistoryKLQuota(ctx context.Context, c *Client) (*qot.RequestHistoryKLQuotaResponse, error) {
	return qot.RequestHistoryKLQuota(ctx, c.inner, &qot.RequestHistoryKLQuotaRequest{})
}

// GetFinancialsEarningsPriceMove retrieves earnings price move data.
func GetFinancialsEarningsPriceMove(ctx context.Context, c *Client, market constant.Market, code string, periodCount int32) (*qot.GetFinancialsEarningsPriceMoveResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetFinancialsEarningsPriceMove: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	req := &qot.GetFinancialsEarningsPriceMoveRequest{
		Security:    sec,
		PeriodCount: periodCount,
	}
	return qot.GetFinancialsEarningsPriceMove(ctx, c.inner, req)
}

// GetFinancialsEarningsPriceHistory retrieves earnings price history data.
func GetFinancialsEarningsPriceHistory(ctx context.Context, c *Client, market constant.Market, code string) (*qot.GetFinancialsEarningsPriceHistoryResponse, error) {
	if code == "" {
		return nil, fmt.Errorf("GetFinancialsEarningsPriceHistory: code is required")
	}
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	return qot.GetFinancialsEarningsPriceHistory(ctx, c.inner, &qot.GetFinancialsEarningsPriceHistoryRequest{Security: sec})
}
