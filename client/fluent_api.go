package client

import (
	"context"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetarkactivetransaction"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetarkfundholding"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetarkstockdynamic"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetdividendcalendar"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetdividendrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetearningsbeatrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetearningscalendar"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteconomiccalendar"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfedwatchdotplot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfedwatchtargetrate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetheatmapdata"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgethighdividendsoerank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgethotlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindicatorlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialchainbyplate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialchaindetail"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialchainlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialplateinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialplatestock"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutiondistribution"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionholdingchange"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionholdinglist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionprofile"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetmacroindicatorhistory"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetmacroindicatorlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionevent"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptioneventalert"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionmarketstatistic"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyinghisstatistic"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyinghisvolatility"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyingoverview"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyingrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionzerodtecontract"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionzerodtescreener"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionearningsscreener"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionsellerscreener"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetperiodchangerank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetratingchange"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrisefalldistr"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsearchnews"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsearchquote"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshortsellingrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgettopmoverrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusafterhoursrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusovernightrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetuspremarketrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequestindicatorcalc"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsetoptioneventalert"
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

func (api *QuoteAPI) GetFinancialsStatements(ctx context.Context, req *qot.GetFinancialsStatementsRequest) (*qot.GetFinancialsStatementsResponse, error) {
	return qot.GetFinancialsStatements(ctx, api.client, req)
}

func (api *QuoteAPI) GetFinancialsRevenueBreakdown(ctx context.Context, req *qot.GetFinancialsRevenueBreakdownRequest) (*qot.GetFinancialsRevenueBreakdownResponse, error) {
	return qot.GetFinancialsRevenueBreakdown(ctx, api.client, req)
}

func (api *QuoteAPI) GetResearchAnalystConsensus(ctx context.Context, req *qot.GetResearchAnalystConsensusRequest) (*qot.GetResearchAnalystConsensusResponse, error) {
	return qot.GetResearchAnalystConsensus(ctx, api.client, req)
}

func (api *QuoteAPI) GetResearchRatingSummary(ctx context.Context, req *qot.GetResearchRatingSummaryRequest) (*qot.GetResearchRatingSummaryResponse, error) {
	return qot.GetResearchRatingSummary(ctx, api.client, req)
}

func (api *QuoteAPI) GetResearchMorningstarReport(ctx context.Context, req *qot.GetResearchMorningstarReportRequest) (*qot.GetResearchMorningstarReportResponse, error) {
	return qot.GetResearchMorningstarReport(ctx, api.client, req)
}

func (api *QuoteAPI) GetValuationDetail(ctx context.Context, req *qot.GetValuationDetailRequest) (*qot.GetValuationDetailResponse, error) {
	return qot.GetValuationDetail(ctx, api.client, req)
}

func (api *QuoteAPI) GetValuationPlateStockList(ctx context.Context, req *qot.GetValuationPlateStockListRequest) (*qot.GetValuationPlateStockListResponse, error) {
	return qot.GetValuationPlateStockList(ctx, api.client, req)
}

func (api *QuoteAPI) GetCorporateActionsDividends(ctx context.Context, req *qot.GetCorporateActionsDividendsRequest) (*qot.GetCorporateActionsDividendsResponse, error) {
	return qot.GetCorporateActionsDividends(ctx, api.client, req)
}

func (api *QuoteAPI) GetCorporateActionsBuybacks(ctx context.Context, req *qot.GetCorporateActionsBuybacksRequest) (*qot.GetCorporateActionsBuybacksResponse, error) {
	return qot.GetCorporateActionsBuybacks(ctx, api.client, req)
}

func (api *QuoteAPI) GetCorporateActionsStockSplits(ctx context.Context, req *qot.GetCorporateActionsStockSplitsRequest) (*qot.GetCorporateActionsStockSplitsResponse, error) {
	return qot.GetCorporateActionsStockSplits(ctx, api.client, req)
}

func (api *QuoteAPI) GetShareholdersOverview(ctx context.Context, req *qot.GetShareholdersOverviewRequest) (*qot.GetShareholdersOverviewResponse, error) {
	return qot.GetShareholdersOverview(ctx, api.client, req)
}

func (api *QuoteAPI) GetShareholdersHoldingChanges(ctx context.Context, req *qot.GetShareholdersHoldingChangesRequest) (*qot.GetShareholdersHoldingChangesResponse, error) {
	return qot.GetShareholdersHoldingChanges(ctx, api.client, req)
}

func (api *QuoteAPI) GetShareholdersHolderDetail(ctx context.Context, req *qot.GetShareholdersHolderDetailRequest) (*qot.GetShareholdersHolderDetailResponse, error) {
	return qot.GetShareholdersHolderDetail(ctx, api.client, req)
}

func (api *QuoteAPI) GetShareholdersInstitutional(ctx context.Context, req *qot.GetShareholdersInstitutionalRequest) (*qot.GetShareholdersInstitutionalResponse, error) {
	return qot.GetShareholdersInstitutional(ctx, api.client, req)
}

func (api *QuoteAPI) GetInsiderHolderList(ctx context.Context, req *qot.GetInsiderHolderListRequest) (*qot.GetInsiderHolderListResponse, error) {
	return qot.GetInsiderHolderList(ctx, api.client, req)
}

func (api *QuoteAPI) GetInsiderTradeList(ctx context.Context, req *qot.GetInsiderTradeListRequest) (*qot.GetInsiderTradeListResponse, error) {
	return qot.GetInsiderTradeList(ctx, api.client, req)
}

func (api *QuoteAPI) GetCompanyProfile(ctx context.Context, req *qot.GetCompanyProfileRequest) (*qot.GetCompanyProfileResponse, error) {
	return qot.GetCompanyProfile(ctx, api.client, req)
}

func (api *QuoteAPI) GetCompanyExecutives(ctx context.Context, req *qot.GetCompanyExecutivesRequest) (*qot.GetCompanyExecutivesResponse, error) {
	return qot.GetCompanyExecutives(ctx, api.client, req)
}

func (api *QuoteAPI) GetCompanyExecutiveBackground(ctx context.Context, req *qot.GetCompanyExecutiveBackgroundRequest) (*qot.GetCompanyExecutiveBackgroundResponse, error) {
	return qot.GetCompanyExecutiveBackground(ctx, api.client, req)
}

func (api *QuoteAPI) GetCompanyOperationalEfficiency(ctx context.Context, req *qot.GetCompanyOperationalEfficiencyRequest) (*qot.GetCompanyOperationalEfficiencyResponse, error) {
	return qot.GetCompanyOperationalEfficiency(ctx, api.client, req)
}

func (api *QuoteAPI) GetTopTenBuySellBrokers(ctx context.Context, req *qot.GetTopTenBuySellBrokersRequest) (*qot.GetTopTenBuySellBrokersResponse, error) {
	return qot.GetTopTenBuySellBrokers(ctx, api.client, req)
}

func (api *QuoteAPI) GetDailyShortVolume(ctx context.Context, req *qot.GetDailyShortVolumeRequest) (*qot.GetDailyShortVolumeResponse, error) {
	return qot.GetDailyShortVolume(ctx, api.client, req)
}

func (api *QuoteAPI) GetShortInterest(ctx context.Context, req *qot.GetShortInterestRequest) (*qot.GetShortInterestResponse, error) {
	return qot.GetShortInterest(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionVolatility(ctx context.Context, req *qot.GetOptionVolatilityRequest) (*qot.GetOptionVolatilityResponse, error) {
	return qot.GetOptionVolatility(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionExerciseProbability(ctx context.Context, req *qot.GetOptionExerciseProbabilityRequest) (*qot.GetOptionExerciseProbabilityResponse, error) {
	return qot.GetOptionExerciseProbability(ctx, api.client, req)
}

func (api *QuoteAPI) GetHistoryKLQuota(ctx context.Context, req *qot.RequestHistoryKLQuotaRequest) (*qot.RequestHistoryKLQuotaResponse, error) {
	return qot.RequestHistoryKLQuota(ctx, api.client, req)
}

func (api *QuoteAPI) GetFinancialsEarningsPriceMove(ctx context.Context, req *qot.GetFinancialsEarningsPriceMoveRequest) (*qot.GetFinancialsEarningsPriceMoveResponse, error) {
	return qot.GetFinancialsEarningsPriceMove(ctx, api.client, req)
}

func (api *QuoteAPI) GetFinancialsEarningsPriceHistory(ctx context.Context, req *qot.GetFinancialsEarningsPriceHistoryRequest) (*qot.GetFinancialsEarningsPriceHistoryResponse, error) {
	return qot.GetFinancialsEarningsPriceHistory(ctx, api.client, req)
}

// GetOptionQuote retrieves real-time quotes for option combo legs.
func (api *QuoteAPI) GetOptionQuote(ctx context.Context, req *qot.GetOptionQuoteRequest) (*qot.GetOptionQuoteResponse, error) {
	return qot.GetOptionQuote(ctx, api.client, req)
}

// GetOptionStrategy retrieves option strategy combo lists.
func (api *QuoteAPI) GetOptionStrategy(ctx context.Context, req *qot.GetOptionStrategyRequest) (*qot.GetOptionStrategyResponse, error) {
	return qot.GetOptionStrategy(ctx, api.client, req)
}

// GetOptionStrategyAnalysis returns P&L analysis for an option strategy combination.
func (api *QuoteAPI) GetOptionStrategyAnalysis(ctx context.Context, req *qot.GetOptionStrategyAnalysisRequest) (*qot.GetOptionStrategyAnalysisResponse, error) {
	return qot.GetOptionStrategyAnalysis(ctx, api.client, req)
}

// GetOptionStrategySpread returns available spread values for an option strategy.
func (api *QuoteAPI) GetOptionStrategySpread(ctx context.Context, req *qot.GetOptionStrategySpreadRequest) (*qot.GetOptionStrategySpreadResponse, error) {
	return qot.GetOptionStrategySpread(ctx, api.client, req)
}

// v10.8+ QuoteAPI extensions

func (api *QuoteAPI) GetSearchQuote(ctx context.Context, req *qotgetsearchquote.C2S) (*qotgetsearchquote.S2C, error) {
	return qot.GetSearchQuote(ctx, api.client, req)
}

func (api *QuoteAPI) GetSearchNews(ctx context.Context, req *qotgetsearchnews.C2S) (*qotgetsearchnews.S2C, error) {
	return qot.GetSearchNews(ctx, api.client, req)
}

func (api *QuoteAPI) GetIndicatorList(ctx context.Context, req *qotgetindicatorlist.C2S) (*qotgetindicatorlist.S2C, error) {
	return qot.GetIndicatorList(ctx, api.client, req)
}

func (api *QuoteAPI) RequestIndicatorCalc(ctx context.Context, req *qotrequestindicatorcalc.C2S) (*qotrequestindicatorcalc.S2C, error) {
	return qot.RequestIndicatorCalc(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionMarketStatistic(ctx context.Context, req *qotgetoptionmarketstatistic.C2S) (*qotgetoptionmarketstatistic.S2C, error) {
	return qot.GetOptionMarketStatistic(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionUnderlyingHisStatistic(ctx context.Context, req *qotgetoptionunderlyinghisstatistic.C2S) (*qotgetoptionunderlyinghisstatistic.S2C, error) {
	return qot.GetOptionUnderlyingHisStatistic(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionUnderlyingOverview(ctx context.Context, req *qotgetoptionunderlyingoverview.C2S) (*qotgetoptionunderlyingoverview.S2C, error) {
	return qot.GetOptionUnderlyingOverview(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionUnderlyingHisVolatility(ctx context.Context, req *qotgetoptionunderlyinghisvolatility.C2S) (*qotgetoptionunderlyinghisvolatility.S2C, error) {
	return qot.GetOptionUnderlyingHisVolatility(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionUnderlyingRank(ctx context.Context, req *qotgetoptionunderlyingrank.C2S) (*qotgetoptionunderlyingrank.S2C, error) {
	return qot.GetOptionUnderlyingRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionRank(ctx context.Context, req *qotgetoptionrank.C2S) (*qotgetoptionrank.S2C, error) {
	return qot.GetOptionRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionEvent(ctx context.Context, req *qotgetoptionevent.C2S) (*qotgetoptionevent.S2C, error) {
	return qot.GetOptionEvent(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionEventAlert(ctx context.Context, req *qotgetoptioneventalert.C2S) (*qotgetoptioneventalert.S2C, error) {
	return qot.GetOptionEventAlert(ctx, api.client, req)
}

func (api *QuoteAPI) SetOptionEventAlert(ctx context.Context, req *qotsetoptioneventalert.C2S) (*qotsetoptioneventalert.S2C, error) {
	return qot.SetOptionEventAlert(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionZeroDteScreener(ctx context.Context, req *qotgetoptionzerodtescreener.C2S) (*qotgetoptionzerodtescreener.S2C, error) {
	return qot.GetOptionZeroDteScreener(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionZeroDteContract(ctx context.Context, req *qotgetoptionzerodtecontract.C2S) (*qotgetoptionzerodtecontract.S2C, error) {
	return qot.GetOptionZeroDteContract(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionEarningsScreener(ctx context.Context, req *qotgetoptionearningsscreener.C2S) (*qotgetoptionearningsscreener.S2C, error) {
	return qot.GetOptionEarningsScreener(ctx, api.client, req)
}

func (api *QuoteAPI) GetOptionSellerScreener(ctx context.Context, req *qotgetoptionsellerscreener.C2S) (*qotgetoptionsellerscreener.S2C, error) {
	return qot.GetOptionSellerScreener(ctx, api.client, req)
}

func (api *QuoteAPI) GetUSPreMarketRank(ctx context.Context, req *qotgetuspremarketrank.C2S) (*qotgetuspremarketrank.S2C, error) {
	return qot.GetUSPreMarketRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetUSAfterHoursRank(ctx context.Context, req *qotgetusafterhoursrank.C2S) (*qotgetusafterhoursrank.S2C, error) {
	return qot.GetUSAfterHoursRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetUSOvernightRank(ctx context.Context, req *qotgetusovernightrank.C2S) (*qotgetusovernightrank.S2C, error) {
	return qot.GetUSOvernightRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetTopMoversRank(ctx context.Context, req *qotgettopmoverrank.C2S) (*qotgettopmoverrank.S2C, error) {
	return qot.GetTopMoversRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetHotList(ctx context.Context, req *qotgethotlist.C2S) (*qotgethotlist.S2C, error) {
	return qot.GetHotList(ctx, api.client, req)
}

func (api *QuoteAPI) GetShortSellingRank(ctx context.Context, req *qotgetshortsellingrank.C2S) (*qotgetshortsellingrank.S2C, error) {
	return qot.GetShortSellingRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetPeriodChangeRank(ctx context.Context, req *qotgetperiodchangerank.C2S) (*qotgetperiodchangerank.S2C, error) {
	return qot.GetPeriodChangeRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetHighDividendSOERank(ctx context.Context, req *qotgethighdividendsoerank.C2S) (*qotgethighdividendsoerank.S2C, error) {
	return qot.GetHighDividendSOERank(ctx, api.client, req)
}

func (api *QuoteAPI) GetInstitutionList(ctx context.Context, req *qotgetinstitutionlist.C2S) (*qotgetinstitutionlist.S2C, error) {
	return qot.GetInstitutionList(ctx, api.client, req)
}

func (api *QuoteAPI) GetInstitutionProfile(ctx context.Context, req *qotgetinstitutionprofile.C2S) (*qotgetinstitutionprofile.S2C, error) {
	return qot.GetInstitutionProfile(ctx, api.client, req)
}

func (api *QuoteAPI) GetInstitutionDistribution(ctx context.Context, req *qotgetinstitutiondistribution.C2S) (*qotgetinstitutiondistribution.S2C, error) {
	return qot.GetInstitutionDistribution(ctx, api.client, req)
}

func (api *QuoteAPI) GetInstitutionHoldingChange(ctx context.Context, req *qotgetinstitutionholdingchange.C2S) (*qotgetinstitutionholdingchange.S2C, error) {
	return qot.GetInstitutionHoldingChange(ctx, api.client, req)
}

func (api *QuoteAPI) GetInstitutionHoldingList(ctx context.Context, req *qotgetinstitutionholdinglist.C2S) (*qotgetinstitutionholdinglist.S2C, error) {
	return qot.GetInstitutionHoldingList(ctx, api.client, req)
}

func (api *QuoteAPI) GetArkFundHolding(ctx context.Context, req *qotgetarkfundholding.C2S) (*qotgetarkfundholding.S2C, error) {
	return qot.GetArkFundHolding(ctx, api.client, req)
}

func (api *QuoteAPI) GetArkStockDynamic(ctx context.Context, req *qotgetarkstockdynamic.C2S) (*qotgetarkstockdynamic.S2C, error) {
	return qot.GetArkStockDynamic(ctx, api.client, req)
}

func (api *QuoteAPI) GetArkActiveTransaction(ctx context.Context, req *qotgetarkactivetransaction.C2S) (*qotgetarkactivetransaction.S2C, error) {
	return qot.GetArkActiveTransaction(ctx, api.client, req)
}

func (api *QuoteAPI) GetRatingChange(ctx context.Context, req *qotgetratingchange.C2S) (*qotgetratingchange.S2C, error) {
	return qot.GetRatingChange(ctx, api.client, req)
}

func (api *QuoteAPI) GetIndustrialChainList(ctx context.Context, req *qotgetindustrialchainlist.C2S) (*qotgetindustrialchainlist.S2C, error) {
	return qot.GetIndustrialChainList(ctx, api.client, req)
}

func (api *QuoteAPI) GetIndustrialChainDetail(ctx context.Context, req *qotgetindustrialchaindetail.C2S) (*qotgetindustrialchaindetail.S2C, error) {
	return qot.GetIndustrialChainDetail(ctx, api.client, req)
}

func (api *QuoteAPI) GetIndustrialChainByPlate(ctx context.Context, req *qotgetindustrialchainbyplate.C2S) (*qotgetindustrialchainbyplate.S2C, error) {
	return qot.GetIndustrialChainByPlate(ctx, api.client, req)
}

func (api *QuoteAPI) GetIndustrialPlateInfo(ctx context.Context, req *qotgetindustrialplateinfo.C2S) (*qotgetindustrialplateinfo.S2C, error) {
	return qot.GetIndustrialPlateInfo(ctx, api.client, req)
}

func (api *QuoteAPI) GetIndustrialPlateStock(ctx context.Context, req *qotgetindustrialplatestock.C2S) (*qotgetindustrialplatestock.S2C, error) {
	return qot.GetIndustrialPlateStock(ctx, api.client, req)
}

func (api *QuoteAPI) GetHeatMapData(ctx context.Context, req *qotgetheatmapdata.C2S) (*qotgetheatmapdata.S2C, error) {
	return qot.GetHeatMapData(ctx, api.client, req)
}

func (api *QuoteAPI) GetRiseFallDistribution(ctx context.Context, req *qotgetrisefalldistr.C2S) (*qotgetrisefalldistr.S2C, error) {
	return qot.GetRiseFallDistribution(ctx, api.client, req)
}

func (api *QuoteAPI) GetEarningsCalendar(ctx context.Context, req *qotgetearningscalendar.C2S) (*qotgetearningscalendar.S2C, error) {
	return qot.GetEarningsCalendar(ctx, api.client, req)
}

func (api *QuoteAPI) GetMacroIndicatorList(ctx context.Context, req *qotgetmacroindicatorlist.C2S) (*qotgetmacroindicatorlist.S2C, error) {
	return qot.GetMacroIndicatorList(ctx, api.client, req)
}

func (api *QuoteAPI) GetMacroIndicatorHistory(ctx context.Context, req *qotgetmacroindicatorhistory.C2S) (*qotgetmacroindicatorhistory.S2C, error) {
	return qot.GetMacroIndicatorHistory(ctx, api.client, req)
}

func (api *QuoteAPI) GetFedWatchTargetRate(ctx context.Context, req *qotgetfedwatchtargetrate.C2S) (*qotgetfedwatchtargetrate.S2C, error) {
	return qot.GetFedWatchTargetRate(ctx, api.client, req)
}

func (api *QuoteAPI) GetFedWatchDotPlot(ctx context.Context, req *qotgetfedwatchdotplot.C2S) (*qotgetfedwatchdotplot.S2C, error) {
	return qot.GetFedWatchDotPlot(ctx, api.client, req)
}

func (api *QuoteAPI) GetEarningsBeatRank(ctx context.Context, req *qotgetearningsbeatrank.C2S) (*qotgetearningsbeatrank.S2C, error) {
	return qot.GetEarningsBeatRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetDividendRank(ctx context.Context, req *qotgetdividendrank.C2S) (*qotgetdividendrank.S2C, error) {
	return qot.GetDividendRank(ctx, api.client, req)
}

func (api *QuoteAPI) GetDividendCalendar(ctx context.Context, req *qotgetdividendcalendar.C2S) (*qotgetdividendcalendar.S2C, error) {
	return qot.GetDividendCalendar(ctx, api.client, req)
}

func (api *QuoteAPI) GetEconomicCalendar(ctx context.Context, req *qotgeteconomiccalendar.C2S) (*qotgeteconomiccalendar.S2C, error) {
	return qot.GetEconomicCalendar(ctx, api.client, req)
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

func (api *TradeAPI) GetFlowSummary(ctx context.Context, req *trd.GetFlowSummaryRequest) (*trd.GetFlowSummaryResponse, error) {
	return trd.GetFlowSummary(ctx, api.client, req)
}

// GetComboMaxTrdQtys retrieves maximum tradable quantities for combo orders.
func (api *TradeAPI) GetComboMaxTrdQtys(ctx context.Context, req *trd.GetComboMaxTrdQtysRequest) (*trd.GetComboMaxTrdQtysResponse, error) {
	return trd.GetComboMaxTrdQtys(ctx, api.client, req)
}

// PlaceComboOrder places a combo order for option strategies.
func (api *TradeAPI) PlaceComboOrder(ctx context.Context, req *trd.PlaceComboOrderRequest) (*trd.PlaceComboOrderResponse, error) {
	return trd.PlaceComboOrder(ctx, api.client, req)
}

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
