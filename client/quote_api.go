package client

import (
	"context"
	"fmt"
	"time"

	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgettradedate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetreference"
	"github.com/shing1211/futuapi4go/pkg/pb/qotstockfilter"
	"github.com/shing1211/futuapi4go/pkg/qot"
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
		Symbol:       code,
		Market:       int32(market),
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
		IsSuspended:  q.IsSuspended,
		SecStatus:   q.SecStatus,
	}, nil
}

// GetKLines retrieves K-line (candlestick) data.
func GetKLines(ctx context.Context, c *Client, market constant.Market, code string, klType constant.KLType, num int) ([]KLine, error) {
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
			Time:       kl.Time,
			IsBlank:    kl.IsBlank,
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
func Subscribe(ctx context.Context, c *Client, market constant.Market, code string, subTypes []constant.SubType) error {
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	_, err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
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
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	_, err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
		SecurityList:     []*qotcommon.Security{sec},
		SubTypeList:      subTypesConverted,
		IsSubOrUnSub:     false,
		IsRegOrUnRegPush: false,
	})
	return err
}

// UnsubscribeAll unsubscribes from all market data.
func UnsubscribeAll(ctx context.Context, c *Client) error {
	_, err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
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

	securities := make([]*qotcommon.Security, len(codes))
	marketPtr := int32(market)
	for i, code := range codes {
		securities[i] = &qotcommon.Security{Market: &marketPtr, Code: &code}
	}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	_, err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
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

	securities := make([]*qotcommon.Security, len(codes))
	marketPtr := int32(market)
	for i, code := range codes {
		securities[i] = &qotcommon.Security{Market: &marketPtr, Code: &code}
	}

	subTypesConverted := make([]qot.SubType, len(subTypes))
	for i, st := range subTypes {
		subTypesConverted[i] = qot.SubType(st)
	}

	_, err := qot.Subscribe(ctx, c.inner, &qot.SubscribeRequest{
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

	_, err := qot.RegQotPush(ctx, c.inner, &qot.RegQotPushRequest{
		SecurityList:   []*qotcommon.Security{sec},
		SubTypeList:    subTypesConverted,
		RehabTypeList:  rehabTypesConverted,
		IsRegOrUnReg:   isReg,
		IsFirstPush:    isFirstPush,
	})
	return err
}

// GetOrderBook retrieves order book data.
func GetOrderBook(ctx context.Context, c *Client, market constant.Market, code string, num int) (*OrderBook, error) {
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
	return book, nil
}

// GetTicker retrieves ticker data.
func GetTicker(ctx context.Context, c *Client, market constant.Market, code string, num int) ([]Ticker, error) {
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
func GetRT(ctx context.Context, c *Client, market constant.Market, code string) ([]RT, error) {
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
	return rtData, nil
}

// GetBroker retrieves broker data.
func GetBroker(ctx context.Context, c *Client, market constant.Market, code string, num int) ([]Broker, []Broker, error) {
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetBroker(ctx, c.inner, &qot.GetBrokerRequest{
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
func GetStaticInfo(ctx context.Context, c *Client, market constant.Market, code string) ([]StaticInfo, error) {
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

// GetSecuritySnapshot returns snapshot data for the given securities.
func GetSecuritySnapshot(ctx context.Context, c *Client, securities []*qotcommon.Security) ([]*Snapshot, error) {
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
		})
	}
	return result, nil
}

// GetTradeDate retrieves trade dates.
func GetTradeDate(ctx context.Context, c *Client, market constant.Market, startDate, endDate string) ([]string, error) {
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
		}
	}
	return infos, nil
}

// GetPlateSet retrieves plate set (板块) list.
func GetPlateSet(ctx context.Context, c *Client, market constant.Market) ([]Plate, error) {
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
		plates[i] = Plate{Code: code, Name: p.Name}
	}
	return plates, nil
}

// GetIpoList retrieves IPO list.
func GetIpoList(ctx context.Context, c *Client, market constant.Market) ([]IpoData, error) {
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
		infos = append(infos, StaticInfo{
			Code: code,
			Name: name,
			Type: secType,
		})
	}
	return infos, nil
}

// GetMarketState retrieves market state (trading status).
func GetMarketState(ctx context.Context, c *Client, market constant.Market, code string) (int32, error) {
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetMarketState(ctx, c.inner, &qot.GetMarketStateRequest{
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

// GetCapitalFlow retrieves capital flow data.
func GetCapitalFlow(ctx context.Context, c *Client, market constant.Market, code string, periodType ...constant.CapitalFlowPeriodType) ([]CapitalFlow, error) {
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
	return flows, nil
}

// GetCapitalDistribution retrieves capital distribution.
func GetCapitalDistribution(ctx context.Context, c *Client, market constant.Market, code string) (*CapitalDistribution, error) {
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetCapitalDistribution(ctx, c.inner, sec)
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
func GetOwnerPlate(ctx context.Context, c *Client, market constant.Market, code string) ([]string, error) {
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	resp, err := qot.GetOwnerPlate(ctx, c.inner, &qot.GetOwnerPlateRequest{
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

// RequestHistoryKL requests historical K-line data with automatic pagination.
func RequestHistoryKL(ctx context.Context, c *Client, market constant.Market, code string, klType constant.KLType, startDate, endDate string) ([]KLine, error) {
	return RequestHistoryKLWithLimit(ctx, c, market, code, klType, startDate, endDate, DefaultHistoryKLPageSize)
}

// RequestHistoryKLWithLimit requests historical K-line data with a configurable page size.
func RequestHistoryKLWithLimit(ctx context.Context, c *Client, market constant.Market, code string, klType constant.KLType, startDate, endDate string, maxPerPage int32) ([]KLine, error) {
	marketPtr := int32(market)
	klTypePtr := int32(klType)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}

	var allKLines []KLine
	var nextReqKey []byte

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

		if len(resp.NextReqKey) == 0 {
			break
		}
		nextReqKey = resp.NextReqKey

		time.Sleep(HistoryKLPaginationDelay)
	}

	return allKLines, nil
}

// GetHistoryKL requests historical K-line data.
func GetHistoryKL(ctx context.Context, c *Client, market constant.Market, code string, klType constant.KLType, rehabType constant.RehabType, startDate, endDate string, maxNum int32) ([]KLine, error) {
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

// GetReference retrieves related/reference securities.
func GetReference(ctx context.Context, c *Client, market constant.Market, code string, refType qotgetreference.ReferenceType) ([]StaticInfo, error) {
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
			Code: code,
			Name: name,
			Type: secType,
		})
	}
	return infos, nil
}

// GetPlateSecurity retrieves securities in a plate.
func GetPlateSecurity(ctx context.Context, c *Client, market constant.Market, plateCode string) ([]StaticInfo, error) {
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
		infos = append(infos, StaticInfo{
			Code: code,
			Name: name,
			Type: secType,
		})
	}
	return infos, nil
}

// GetOptionExpirationDate retrieves option expiration dates.
func GetOptionExpirationDate(ctx context.Context, c *Client, market constant.Market, code string) ([]OptionExpiration, error) {
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
			Date: e.StrikeTime,
			Days: e.OptionExpiryDateDistance,
			Desc: fmt.Sprintf("Cycle %d", e.Cycle),
		})
	}
	return expirations, nil
}

// ModifyUserSecurity adds/removes securities from user group.
func ModifyUserSecurity(ctx context.Context, c *Client, groupName string, op int32, market constant.Market, codes []string) error {
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
		IsSub:    len(resp.ConnSubInfoList) > 0,
		SubTypes: types,
		Security: fmt.Sprintf("Used: %d, Remain: %d", quota, resp.RemainQuota),
	}, nil
}

// SetPriceReminder creates/updates/deletes a price reminder.
func SetPriceReminder(ctx context.Context, c *Client, market constant.Market, code string, op constant.PriceReminderOp, reminderType constant.PriceReminderType, freq constant.PriceReminderFreq, value float64, note string) (int64, error) {
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
			Type:            cc.Type,
			Security:        cc.Security,
			RelatedSecurity: cc.RelatedSecurity,
			PublicTime:      cc.PublicTime,
			EffectiveTime:   cc.EffectiveTime,
		})
	}
	return result, nil
}

// GetHoldingChangeList retrieves holding change list.
func GetHoldingChangeList(ctx context.Context, c *Client, market constant.Market, code string, holderCategory constant.HolderCategory, beginTime, endTime string) ([]*HoldingChangeInfo, error) {
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
func GetOptionChain(ctx context.Context, c *Client, market constant.Market, code string, indexOptionType constant.IndexOptionType, optType constant.OptionType, condition int32, beginTime, endTime string) ([]*OptChain, error) {
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
func GetWarrant(ctx context.Context, c *Client, market constant.Market, code string, begin, num int32, sortField constant.WarrantSortField, ascend bool, optType constant.WarrantType, issuer qotcommon.Issuer, status constant.WarrantStatus) ([]*WarrantData, error) {
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
	return result, nil
}

// RequestRehab requests rehabilitation (复权) data.
func RequestRehab(ctx context.Context, c *Client, market constant.Market, code string) ([]*RehabInfo, error) {
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

// GetRehab returns rehabilitation (复权) data for the given security.
// Unlike RequestRehab (which may use cached data), GetRehab always fetches
// fresh rehabilitation data from the server.
func GetRehab(ctx context.Context, c *Client, market constant.Market, code string) ([]*RehabInfo, error) {
	marketPtr := int32(market)
	sec := &qotcommon.Security{Market: &marketPtr, Code: &code}
	resp, err := qot.GetRehab(ctx, c.inner, &qot.GetRehabRequest{
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

// RequestHistoryKLQuota queries historical K-line quota.
func RequestHistoryKLQuota(ctx context.Context, c *Client) (*HistoryKLQuotaInfo, error) {
	resp, err := qot.RequestHistoryKLQuota(ctx, c.inner, &qot.RequestHistoryKLQuotaRequest{
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

// GetHistoryKLPoints retrieves K-line data at specific time points for a security.
func GetHistoryKLPoints(ctx context.Context, c *Client, market constant.Market, code string, times []string, klType constant.KLType, rehabType constant.RehabType, noDataMode qot.NoDataMode) (*qot.GetHistoryKLPointsResponse, error) {
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
	req := &qotgettradedate.C2S{
		Market:    &market,
		BeginTime: &beginTime,
		EndTime:   &endTime,
	}
	resp, err := qot.GetTradeDate(ctx, c.inner, req)
	if err != nil {
		return nil, err
	}
	dates := make([]TradeDate, 0, len(resp.TradeDateList))
	for _, d := range resp.TradeDateList {
		if d == nil {
			continue
		}
		dates = append(dates, TradeDate{
			Time:          d.GetTime(),
			Timestamp:     d.GetTimestamp(),
			TradeDateType: d.GetTradeDateType(),
		})
	}
	return dates, nil
}
