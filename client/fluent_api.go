package client

import (
	"context"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/pkg/sys"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

// QuoteAPI provides a fluent API for market data operations.
// Use client.Quote() to get an instance.
type QuoteAPI struct {
	client *futuapi.Client
}

// GetBasicQot retrieves basic quote for securities.
func (api *QuoteAPI) GetBasicQot(ctx context.Context, securities []*qotcommon.Security) ([]*qot.BasicQot, error) {
	return qot.GetBasicQot(ctx, api.client, securities)
}

// GetKL retrieves K-line data.
func (api *QuoteAPI) GetKL(ctx context.Context, req *qot.GetKLRequest) (*qot.GetKLResponse, error) {
	return qot.GetKL(ctx, api.client, req)
}

// GetOrderBook retrieves order book (买卖盘) data.
func (api *QuoteAPI) GetOrderBook(ctx context.Context, req *qot.GetOrderBookRequest) (*qot.GetOrderBookResponse, error) {
	return qot.GetOrderBook(ctx, api.client, req)
}

// GetTicker retrieves ticker (逐笔) data.
func (api *QuoteAPI) GetTicker(ctx context.Context, req *qot.GetTickerRequest) (*qot.GetTickerResponse, error) {
	return qot.GetTicker(ctx, api.client, req)
}

// GetRT retrieves real-time (分时) data.
func (api *QuoteAPI) GetRT(ctx context.Context, req *qot.GetRTRequest) (*qot.GetRTResponse, error) {
	return qot.GetRT(ctx, api.client, req)
}

// GetStaticInfo retrieves static stock info.
func (api *QuoteAPI) GetStaticInfo(ctx context.Context, req *qot.GetStaticInfoRequest) (*qot.GetStaticInfoResponse, error) {
	return qot.GetStaticInfo(ctx, api.client, req)
}

// GetSecuritySnapshot retrieves security snapshot.
func (api *QuoteAPI) GetSecuritySnapshot(ctx context.Context, req *qot.GetSecuritySnapshotRequest) (*qot.GetSecuritySnapshotResponse, error) {
	return qot.GetSecuritySnapshot(ctx, api.client, req)
}

// Subscribe subscribes to market data.
func (api *QuoteAPI) Subscribe(ctx context.Context, req *qot.SubscribeRequest) error {
	return qot.Subscribe(ctx, api.client, req)
}

// RegQotPush registers for push notifications.
func (api *QuoteAPI) RegQotPush(ctx context.Context, req *qot.RegQotPushRequest) error {
	return qot.RegQotPush(ctx, api.client, req)
}

// GetSubInfo retrieves subscription info.
func (api *QuoteAPI) GetSubInfo(ctx context.Context) (*qot.GetSubInfoResponse, error) {
	return qot.GetSubInfo(ctx, api.client)
}

// RequestHistoryKL requests historical K-line data.
func (api *QuoteAPI) RequestHistoryKL(ctx context.Context, req *qot.RequestHistoryKLRequest) (*qot.RequestHistoryKLResponse, error) {
	return qot.RequestHistoryKL(ctx, api.client, req)
}

// StockFilter filters stocks by criteria.
func (api *QuoteAPI) StockFilter(ctx context.Context, req *qot.StockFilterRequest) (*qot.StockFilterResponse, error) {
	return qot.StockFilter(ctx, api.client, req)
}

// GetHistoryKLPoints retrieves historical K-line at specific times.
func (api *QuoteAPI) GetHistoryKLPoints(ctx context.Context, req *qot.GetHistoryKLPointsRequest) (*qot.GetHistoryKLPointsResponse, error) {
	return qot.GetHistoryKLPoints(ctx, api.client, req)
}

// GetCapitalFlow retrieves capital flow data.
func (api *QuoteAPI) GetCapitalFlow(ctx context.Context, req *qot.GetCapitalFlowRequest) (*qot.GetCapitalFlowResponse, error) {
	return qot.GetCapitalFlow(ctx, api.client, req)
}

// GetOptionChain retrieves option chain.
func (api *QuoteAPI) GetOptionChain(ctx context.Context, req *qot.GetOptionChainRequest) (*qot.GetOptionChainResponse, error) {
	return qot.GetOptionChain(ctx, api.client, req)
}

// GetUserSecurity retrieves user security list.
func (api *QuoteAPI) GetUserSecurity(ctx context.Context, groupName string) (*qot.GetUserSecurityResponse, error) {
	return qot.GetUserSecurity(ctx, api.client, groupName)
}

// ModifyUserSecurity modifies user security.
func (api *QuoteAPI) ModifyUserSecurity(ctx context.Context, req *qot.ModifyUserSecurityRequest) (*qot.ModifyUserSecurityResponse, error) {
	return qot.ModifyUserSecurity(ctx, api.client, req)
}

// SetPriceReminder sets price reminder.
func (api *QuoteAPI) SetPriceReminder(ctx context.Context, req *qot.SetPriceReminderRequest) (*qot.SetPriceReminderResponse, error) {
	return qot.SetPriceReminder(ctx, api.client, req)
}

// GetPriceReminder retrieves price reminders.
func (api *QuoteAPI) GetPriceReminder(ctx context.Context, security *qotcommon.Security, market int32) (*qot.GetPriceReminderResponse, error) {
	return qot.GetPriceReminder(ctx, api.client, security, market)
}

// TradeAPI provides a fluent API for trading operations.
// Use client.Trade() to get an instance.
type TradeAPI struct {
	client *futuapi.Client
	trdEnv constant.TrdEnv
	trdMkt constant.TrdMarket
}

// GetAccList retrieves account list.
func (api *TradeAPI) GetAccList(ctx context.Context, trdCat constant.TrdCategory, isAll bool) (*trd.GetAccListResponse, error) {
	return trd.GetAccList(ctx, api.client, trdCat, isAll)
}

// Unlock unlocks trading with the given password.
// Pass empty string to lock trading.
func (api *TradeAPI) Unlock(ctx context.Context, pwdMD5 string) error {
	req := &trd.UnlockTradeRequest{Unlock: pwdMD5 != ""}
	if pwdMD5 != "" {
		req.PwdMD5 = constant.SensitiveString(pwdMD5)
	}
	return trd.UnlockTrade(ctx, api.client, req)
}

// GetFunds retrieves account funds.
func (api *TradeAPI) GetFunds(ctx context.Context, req *trd.GetFundsRequest) (*trd.GetFundsResponse, error) {
	return trd.GetFunds(ctx, api.client, req)
}

// GetPositionList retrieves positions.
func (api *TradeAPI) GetPositionList(ctx context.Context, req *trd.GetPositionListRequest) (*trd.GetPositionListResponse, error) {
	return trd.GetPositionList(ctx, api.client, req)
}

// GetOrderList retrieves orders.
func (api *TradeAPI) GetOrderList(ctx context.Context, req *trd.GetOrderListRequest) (*trd.GetOrderListResponse, error) {
	return trd.GetOrderList(ctx, api.client, req)
}

// PlaceOrder places an order.
func (api *TradeAPI) PlaceOrder(ctx context.Context, req *trd.PlaceOrderRequest) (*trd.PlaceOrderResponse, error) {
	return trd.PlaceOrder(ctx, api.client, req)
}

// ModifyOrder modifies an order.
func (api *TradeAPI) ModifyOrder(ctx context.Context, req *trd.ModifyOrderRequest) (*trd.ModifyOrderResponse, error) {
	return trd.ModifyOrder(ctx, api.client, req)
}

// GetOrderFillList retrieves order fills.
func (api *TradeAPI) GetOrderFillList(ctx context.Context, req *trd.GetOrderFillListRequest) (*trd.GetOrderFillListResponse, error) {
	return trd.GetOrderFillList(ctx, api.client, req)
}

// GetHistoryOrderList retrieves history orders.
func (api *TradeAPI) GetHistoryOrderList(ctx context.Context, req *trd.GetHistoryOrderListRequest) (*trd.GetHistoryOrderListResponse, error) {
	return trd.GetHistoryOrderList(ctx, api.client, req)
}

// GetHistoryOrderFillList retrieves history fills.
func (api *TradeAPI) GetHistoryOrderFillList(ctx context.Context, req *trd.GetHistoryOrderFillListRequest) (*trd.GetHistoryOrderFillListResponse, error) {
	return trd.GetHistoryOrderFillList(ctx, api.client, req)
}

// GetMaxTrdQtys retrieves max trade quantities.
func (api *TradeAPI) GetMaxTrdQtys(ctx context.Context, req *trd.GetMaxTrdQtysRequest) (*trd.GetMaxTrdQtysResponse, error) {
	return trd.GetMaxTrdQtys(ctx, api.client, req)
}

// GetOrderFee retrieves order fee.
func (api *TradeAPI) GetOrderFee(ctx context.Context, req *trd.GetOrderFeeRequest) (*trd.GetOrderFeeResponse, error) {
	return trd.GetOrderFee(ctx, api.client, req)
}

// GetMarginRatio retrieves margin ratio.
func (api *TradeAPI) GetMarginRatio(ctx context.Context, req *trd.GetMarginRatioRequest) (*trd.GetMarginRatioResponse, error) {
	return trd.GetMarginRatio(ctx, api.client, req)
}

// GetFlowSummary retrieves flow summary.
func (api *TradeAPI) GetFlowSummary(ctx context.Context, req *trd.GetFlowSummaryRequest) (*trd.GetFlowSummaryResponse, error) {
	return trd.GetFlowSummary(ctx, api.client, req)
}

// NewOrderBuilder creates a new order for building with the client's trading environment.
func (api *TradeAPI) NewOrder(accID uint64, market constant.TrdMarket) *trd.OrderBuilder {
	return trd.NewOrder(accID, market, api.trdEnv)
}

// SystemAPI provides a fluent API for system operations.
// Use client.System() to get an instance.
type SystemAPI struct {
	client *futuapi.Client
}

// GetGlobalState retrieves global state.
func (api *SystemAPI) GetGlobalState(ctx context.Context) (*sys.GetGlobalStateResponse, error) {
	return sys.GetGlobalState(ctx, api.client)
}

// GetUserInfo retrieves user info.
func (api *SystemAPI) GetUserInfo(ctx context.Context, req *sys.GetUserInfoRequest) (*sys.GetUserInfoResponse, error) {
	return sys.GetUserInfo(ctx, api.client, req)
}

// GetDelayStatistics retrieves delay statistics.
func (api *SystemAPI) GetDelayStatistics(ctx context.Context) (*sys.GetDelayStatisticsResponse, error) {
	return sys.GetDelayStatistics(ctx, api.client, nil)
}

// GetUsedQuota retrieves quota usage.
func (api *SystemAPI) GetUsedQuota(ctx context.Context) (*sys.GetUsedQuotaResponse, error) {
	return sys.GetUsedQuota(ctx, api.client)
}
