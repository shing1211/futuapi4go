// Package client provides a public Client type for the Futu OpenD SDK.
// This allows external projects to use the SDK.
package client

import (
	"context"
	"fmt"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

// Client is the main client type for connecting to Futu OpenD.
// It wraps the internal client to provide a public API.
type Client struct {
	inner *futuapi.Client
}

// New creates a new client with optional configuration.
func New(opts ...Option) *Client {
	futuOpts := make([]futuapi.Option, len(opts))
	for i, o := range opts {
		futuOpts[i] = o
	}
	return &Client{
		inner: futuapi.New(futuOpts...),
	}
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
		Symbol:    code,
		Market:    market,
		Price:     q.CurPrice,
		Open:      q.OpenPrice,
		High:      q.HighPrice,
		Low:       q.LowPrice,
		Volume:    q.Volume,
		Timestamp: q.UpdateTime,
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
			Time:   kl.Time,
			Open:   kl.OpenPrice,
			High:   kl.HighPrice,
			Low:    kl.LowPrice,
			Close:  kl.ClosePrice,
			Volume: kl.Volume,
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

// GetAccountList retrieves the list of trading accounts.
func GetAccountList(c *Client) ([]Account, error) {
	resp, err := trd.GetAccList(c.inner, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
	if err != nil {
		return nil, err
	}

	accounts := make([]Account, len(resp.AccList))
	for i, acc := range resp.AccList {
		accounts[i] = Account{
			AccID:     acc.AccID,
			AccType:   acc.AccType,
			TrdEnv:    acc.TrdEnv,
			CardNum:   acc.CardNum,
			AccStatus: acc.AccStatus,
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
		TrdEnv:    1,
		Code:      code,
		TrdSide:   side,
		OrderType: orderType,
		Price:     price,
		Qty:       qty,
	})
	if err != nil {
		return nil, err
	}
	return &PlaceOrderResult{OrderID: resp.OrderID}, nil
}

// GetPositionList retrieves the current positions.
func GetPositionList(c *Client, accID uint64) ([]Position, error) {
	resp, err := trd.GetPositionList(c.inner, &trd.GetPositionListRequest{
		AccID:     accID,
		TrdMarket: 0,
		TrdEnv:    1,
	})
	if err != nil {
		return nil, err
	}

	positions := make([]Position, len(resp.PositionList))
	for i, p := range resp.PositionList {
		positions[i] = Position{
			Symbol:    p.Code,
			Market:    0,
			Quantity:  p.Qty,
			CostPrice: p.CostPrice,
			CurPrice:  p.Price,
			PnL:       p.PlVal,
			PnLRate:   p.PlRatio,
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
		Cash:        f.Cash,
		BuyingPower: f.AvailableFunds,
		MarketValue: f.MarketVal,
		TotalAsset:  f.TotalAssets,
	}, nil
}

// GetOrderList retrieves active orders.
func GetOrderList(c *Client, accID uint64) ([]Order, error) {
	resp, err := trd.GetOrderList(c.inner, &trd.GetOrderListRequest{
		AccID:     accID,
		TrdMarket: 0,
		TrdEnv:    1,
	})
	if err != nil {
		return nil, err
	}

	orders := make([]Order, len(resp.OrderList))
	for i, o := range resp.OrderList {
		orders[i] = Order{
			OrderID:    o.OrderID,
			Code:       o.Code,
			Name:       o.Name,
			TrdSide:    o.TrdSide,
			OrderType:  o.OrderType,
			Price:      o.Price,
			Qty:        o.Qty,
			OrderState: o.OrderStatus,
		}
	}
	return orders, nil
}

// GetOrderFillList retrieves order fills (executions).
func GetOrderFillList(c *Client, accID uint64) ([]OrderFill, error) {
	resp, err := trd.GetOrderFillList(c.inner, &trd.GetOrderFillListRequest{
		AccID:     accID,
		TrdMarket: 0,
		TrdEnv:    1,
	})
	if err != nil {
		return nil, err
	}

	fills := make([]OrderFill, len(resp.OrderFillList))
	for i, f := range resp.OrderFillList {
		fills[i] = OrderFill{
			OrderID: f.OrderID,
			Code:    f.Code,
			Name:    f.Name,
			TrdSide: f.TrdSide,
			Price:   f.Price,
			Qty:     f.Qty,
		}
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
		Bids: make([]OrderBookItem, len(resp.OrderBookBidList)),
		Asks: make([]OrderBookItem, len(resp.OrderBookAskList)),
	}
	for i, b := range resp.OrderBookBidList {
		book.Bids[i] = OrderBookItem{Price: b.Price, Volume: b.Volume}
	}
	for i, a := range resp.OrderBookAskList {
		book.Asks[i] = OrderBookItem{Price: a.Price, Volume: a.Volume}
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
			Price:     t.Price,
			Volume:    t.Volume,
			Direction: dir,
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
			Time:   r.Time,
			Price:  r.Price,
			Volume: r.Volume,
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
		bidBrokers[i] = Broker{ID: b.ID, Name: b.Name}
	}
	askBrokers := make([]Broker, len(resp.AskBrokerList))
	for i, a := range resp.AskBrokerList {
		askBrokers[i] = Broker{ID: a.ID, Name: a.Name}
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
		if s.Basic != nil {
			if s.Basic.Name != nil {
				name = *s.Basic.Name
			}
			if s.Basic.SecType != nil {
				secType = *s.Basic.SecType
			}
		}
		infos[i] = StaticInfo{Code: code, Name: name, Type: secType}
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
		infos[i] = FutureInfo{
			Code:     secCode,
			Name:     f.Name,
			Expire:   f.LastTradeTime,
			InstType: 0,
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

// ============================================================================
// Types
// ============================================================================

// Quote represents a real-time quote.
type Quote struct {
	Symbol    string
	Market    int32
	Price     float64
	Open      float64
	High      float64
	Low       float64
	Volume    int64
	Timestamp string
}

// KLine represents a K-line (candlestick) data point.
type KLine struct {
	Time   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int64
}

// Account represents a trading account.
type Account struct {
	AccID     uint64
	AccType   int32
	TrdEnv    int32
	CardNum   string
	AccStatus int32
}

// PlaceOrderResult represents a place order result.
type PlaceOrderResult struct {
	OrderID uint64
}

// Position represents a position.
type Position struct {
	Symbol    string
	Market    int32
	Quantity  float64
	CostPrice float64
	CurPrice  float64
	PnL       float64
	PnLRate   float64
}

// Funds represents account funds.
type Funds struct {
	Cash        float64
	BuyingPower float64
	MarketValue float64
	TotalAsset  float64
}

// Order represents an order.
type Order struct {
	OrderID    uint64
	Code       string
	Name       string
	TrdSide    int32
	OrderType  int32
	Price      float64
	Qty        float64
	OrderState int32
}

// OrderFill represents an order fill.
type OrderFill struct {
	OrderID uint64
	Code    string
	Name    string
	TrdSide int32
	Price   float64
	Qty     float64
}

// OrderBook represents order book data.
type OrderBook struct {
	Bids []OrderBookItem
	Asks []OrderBookItem
}

// OrderBookItem represents a single order book entry.
type OrderBookItem struct {
	Price  float64
	Volume int64
}

// Ticker represents ticker data.
type Ticker struct {
	Time      string
	Price     float64
	Volume    int64
	Direction string
}

// RT represents real-time data.
type RT struct {
	Time   string
	Price  float64
	Volume int64
}

// Broker represents broker data.
type Broker struct {
	ID   int64
	Name string
}

// StaticInfo represents static security info.
type StaticInfo struct {
	Code string
	Name string
	Type int32
}

// FutureInfo represents futures info.
type FutureInfo struct {
	Code     string
	Name     string
	Expire   string
	InstType int32
}

// Plate represents a market plate (板块).
type Plate struct {
	Code string
	Name string
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

// Option configures the client (alias for backward compatibility).
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

// WithMaxRetries sets the maximum retry attempts.
func WithMaxRetries(n int) Option {
	return futuapi.WithMaxRetries(n)
}

// WithLogLevel sets the logging level (0=info, 1=warn, 2=error, 3=silent).
func WithLogLevel(level int) Option {
	return futuapi.WithLogLevel(level)
}

// Default timeouts.
const (
	DefaultDialTimeout      = 10 * time.Second
	DefaultAPITimeout       = 30 * time.Second
	DefaultKeepAlive        = 30 * time.Second
	DefaultMaxRetries       = 3
	DefaultReconnectBackoff = 1.5
)
