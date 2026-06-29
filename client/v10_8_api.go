package client

import (
	"context"
	"fmt"

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
)

// v10.8+ API wrappers — search

func GetSearchQuote(ctx context.Context, c *Client, req *qotgetsearchquote.C2S) (*qotgetsearchquote.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetSearchQuote: req is nil")
	}
	return qot.GetSearchQuote(ctx, c.inner, req)
}

func GetSearchNews(ctx context.Context, c *Client, req *qotgetsearchnews.C2S) (*qotgetsearchnews.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetSearchNews: req is nil")
	}
	return qot.GetSearchNews(ctx, c.inner, req)
}

func GetIndicatorList(ctx context.Context, c *Client, req *qotgetindicatorlist.C2S) (*qotgetindicatorlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndicatorList: req is nil")
	}
	return qot.GetIndicatorList(ctx, c.inner, req)
}

func RequestIndicatorCalc(ctx context.Context, c *Client, req *qotrequestindicatorcalc.C2S) (*qotrequestindicatorcalc.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("RequestIndicatorCalc: req is nil")
	}
	return qot.RequestIndicatorCalc(ctx, c.inner, req)
}

// v10.8+ Option Analytics API wrappers

func GetOptionMarketStatistic(ctx context.Context, c *Client, req *qotgetoptionmarketstatistic.C2S) (*qotgetoptionmarketstatistic.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionMarketStatistic: req is nil")
	}
	return qot.GetOptionMarketStatistic(ctx, c.inner, req)
}

func GetOptionUnderlyingHisStatistic(ctx context.Context, c *Client, req *qotgetoptionunderlyinghisstatistic.C2S) (*qotgetoptionunderlyinghisstatistic.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionUnderlyingHisStatistic: req is nil")
	}
	return qot.GetOptionUnderlyingHisStatistic(ctx, c.inner, req)
}

func GetOptionUnderlyingOverview(ctx context.Context, c *Client, req *qotgetoptionunderlyingoverview.C2S) (*qotgetoptionunderlyingoverview.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionUnderlyingOverview: req is nil")
	}
	return qot.GetOptionUnderlyingOverview(ctx, c.inner, req)
}

func GetOptionUnderlyingHisVolatility(ctx context.Context, c *Client, req *qotgetoptionunderlyinghisvolatility.C2S) (*qotgetoptionunderlyinghisvolatility.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionUnderlyingHisVolatility: req is nil")
	}
	return qot.GetOptionUnderlyingHisVolatility(ctx, c.inner, req)
}

func GetOptionUnderlyingRank(ctx context.Context, c *Client, req *qotgetoptionunderlyingrank.C2S) (*qotgetoptionunderlyingrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionUnderlyingRank: req is nil")
	}
	return qot.GetOptionUnderlyingRank(ctx, c.inner, req)
}

func GetOptionRank(ctx context.Context, c *Client, req *qotgetoptionrank.C2S) (*qotgetoptionrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionRank: req is nil")
	}
	return qot.GetOptionRank(ctx, c.inner, req)
}

func GetOptionEvent(ctx context.Context, c *Client, req *qotgetoptionevent.C2S) (*qotgetoptionevent.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionEvent: req is nil")
	}
	return qot.GetOptionEvent(ctx, c.inner, req)
}

func GetOptionEventAlert(ctx context.Context, c *Client, req *qotgetoptioneventalert.C2S) (*qotgetoptioneventalert.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionEventAlert: req is nil")
	}
	return qot.GetOptionEventAlert(ctx, c.inner, req)
}

func SetOptionEventAlert(ctx context.Context, c *Client, req *qotsetoptioneventalert.C2S) (*qotsetoptioneventalert.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("SetOptionEventAlert: req is nil")
	}
	return qot.SetOptionEventAlert(ctx, c.inner, req)
}

func GetOptionZeroDteScreener(ctx context.Context, c *Client, req *qotgetoptionzerodtescreener.C2S) (*qotgetoptionzerodtescreener.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionZeroDteScreener: req is nil")
	}
	return qot.GetOptionZeroDteScreener(ctx, c.inner, req)
}

func GetOptionZeroDteContract(ctx context.Context, c *Client, req *qotgetoptionzerodtecontract.C2S) (*qotgetoptionzerodtecontract.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionZeroDteContract: req is nil")
	}
	return qot.GetOptionZeroDteContract(ctx, c.inner, req)
}

func GetOptionEarningsScreener(ctx context.Context, c *Client, req *qotgetoptionearningsscreener.C2S) (*qotgetoptionearningsscreener.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionEarningsScreener: req is nil")
	}
	return qot.GetOptionEarningsScreener(ctx, c.inner, req)
}

func GetOptionSellerScreener(ctx context.Context, c *Client, req *qotgetoptionsellerscreener.C2S) (*qotgetoptionsellerscreener.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetOptionSellerScreener: req is nil")
	}
	return qot.GetOptionSellerScreener(ctx, c.inner, req)
}

// v10.8+ Ranking API wrappers

func GetUSPreMarketRank(ctx context.Context, c *Client, req *qotgetuspremarketrank.C2S) (*qotgetuspremarketrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetUSPreMarketRank: req is nil")
	}
	return qot.GetUSPreMarketRank(ctx, c.inner, req)
}

func GetUSAfterHoursRank(ctx context.Context, c *Client, req *qotgetusafterhoursrank.C2S) (*qotgetusafterhoursrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetUSAfterHoursRank: req is nil")
	}
	return qot.GetUSAfterHoursRank(ctx, c.inner, req)
}

func GetUSOvernightRank(ctx context.Context, c *Client, req *qotgetusovernightrank.C2S) (*qotgetusovernightrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetUSOvernightRank: req is nil")
	}
	return qot.GetUSOvernightRank(ctx, c.inner, req)
}

func GetTopMoversRank(ctx context.Context, c *Client, req *qotgettopmoverrank.C2S) (*qotgettopmoverrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetTopMoversRank: req is nil")
	}
	return qot.GetTopMoversRank(ctx, c.inner, req)
}

func GetHotList(ctx context.Context, c *Client, req *qotgethotlist.C2S) (*qotgethotlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHotList: req is nil")
	}
	return qot.GetHotList(ctx, c.inner, req)
}

func GetShortSellingRank(ctx context.Context, c *Client, req *qotgetshortsellingrank.C2S) (*qotgetshortsellingrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetShortSellingRank: req is nil")
	}
	return qot.GetShortSellingRank(ctx, c.inner, req)
}

func GetPeriodChangeRank(ctx context.Context, c *Client, req *qotgetperiodchangerank.C2S) (*qotgetperiodchangerank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetPeriodChangeRank: req is nil")
	}
	return qot.GetPeriodChangeRank(ctx, c.inner, req)
}

func GetHighDividendSOERank(ctx context.Context, c *Client, req *qotgethighdividendsoerank.C2S) (*qotgethighdividendsoerank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHighDividendSOERank: req is nil")
	}
	return qot.GetHighDividendSOERank(ctx, c.inner, req)
}

// v10.8+ Institutional Data API wrappers

func GetInstitutionList(ctx context.Context, c *Client, req *qotgetinstitutionlist.C2S) (*qotgetinstitutionlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionList: req is nil")
	}
	return qot.GetInstitutionList(ctx, c.inner, req)
}

func GetInstitutionProfile(ctx context.Context, c *Client, req *qotgetinstitutionprofile.C2S) (*qotgetinstitutionprofile.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionProfile: req is nil")
	}
	return qot.GetInstitutionProfile(ctx, c.inner, req)
}

func GetInstitutionDistribution(ctx context.Context, c *Client, req *qotgetinstitutiondistribution.C2S) (*qotgetinstitutiondistribution.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionDistribution: req is nil")
	}
	return qot.GetInstitutionDistribution(ctx, c.inner, req)
}

func GetInstitutionHoldingChange(ctx context.Context, c *Client, req *qotgetinstitutionholdingchange.C2S) (*qotgetinstitutionholdingchange.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionHoldingChange: req is nil")
	}
	return qot.GetInstitutionHoldingChange(ctx, c.inner, req)
}

func GetInstitutionHoldingList(ctx context.Context, c *Client, req *qotgetinstitutionholdinglist.C2S) (*qotgetinstitutionholdinglist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetInstitutionHoldingList: req is nil")
	}
	return qot.GetInstitutionHoldingList(ctx, c.inner, req)
}

func GetArkFundHolding(ctx context.Context, c *Client, req *qotgetarkfundholding.C2S) (*qotgetarkfundholding.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetArkFundHolding: req is nil")
	}
	return qot.GetArkFundHolding(ctx, c.inner, req)
}

func GetArkStockDynamic(ctx context.Context, c *Client, req *qotgetarkstockdynamic.C2S) (*qotgetarkstockdynamic.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetArkStockDynamic: req is nil")
	}
	return qot.GetArkStockDynamic(ctx, c.inner, req)
}

func GetArkActiveTransaction(ctx context.Context, c *Client, req *qotgetarkactivetransaction.C2S) (*qotgetarkactivetransaction.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetArkActiveTransaction: req is nil")
	}
	return qot.GetArkActiveTransaction(ctx, c.inner, req)
}

func GetRatingChange(ctx context.Context, c *Client, req *qotgetratingchange.C2S) (*qotgetratingchange.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetRatingChange: req is nil")
	}
	return qot.GetRatingChange(ctx, c.inner, req)
}

// v10.8+ Industrial Chain API wrappers

func GetIndustrialChainList(ctx context.Context, c *Client, req *qotgetindustrialchainlist.C2S) (*qotgetindustrialchainlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialChainList: req is nil")
	}
	return qot.GetIndustrialChainList(ctx, c.inner, req)
}

func GetIndustrialChainDetail(ctx context.Context, c *Client, req *qotgetindustrialchaindetail.C2S) (*qotgetindustrialchaindetail.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialChainDetail: req is nil")
	}
	return qot.GetIndustrialChainDetail(ctx, c.inner, req)
}

func GetIndustrialChainByPlate(ctx context.Context, c *Client, req *qotgetindustrialchainbyplate.C2S) (*qotgetindustrialchainbyplate.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialChainByPlate: req is nil")
	}
	return qot.GetIndustrialChainByPlate(ctx, c.inner, req)
}

func GetIndustrialPlateInfo(ctx context.Context, c *Client, req *qotgetindustrialplateinfo.C2S) (*qotgetindustrialplateinfo.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialPlateInfo: req is nil")
	}
	return qot.GetIndustrialPlateInfo(ctx, c.inner, req)
}

func GetIndustrialPlateStock(ctx context.Context, c *Client, req *qotgetindustrialplatestock.C2S) (*qotgetindustrialplatestock.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIndustrialPlateStock: req is nil")
	}
	return qot.GetIndustrialPlateStock(ctx, c.inner, req)
}

// v10.8+ Heat Map API wrappers

func GetHeatMapData(ctx context.Context, c *Client, req *qotgetheatmapdata.C2S) (*qotgetheatmapdata.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetHeatMapData: req is nil")
	}
	return qot.GetHeatMapData(ctx, c.inner, req)
}

func GetRiseFallDistribution(ctx context.Context, c *Client, req *qotgetrisefalldistr.C2S) (*qotgetrisefalldistr.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetRiseFallDistribution: req is nil")
	}
	return qot.GetRiseFallDistribution(ctx, c.inner, req)
}

// v10.8+ Market Fundamentals API wrappers

func GetEarningsCalendar(ctx context.Context, c *Client, req *qotgetearningscalendar.C2S) (*qotgetearningscalendar.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEarningsCalendar: req is nil")
	}
	return qot.GetEarningsCalendar(ctx, c.inner, req)
}

func GetMacroIndicatorList(ctx context.Context, c *Client, req *qotgetmacroindicatorlist.C2S) (*qotgetmacroindicatorlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetMacroIndicatorList: req is nil")
	}
	return qot.GetMacroIndicatorList(ctx, c.inner, req)
}

func GetMacroIndicatorHistory(ctx context.Context, c *Client, req *qotgetmacroindicatorhistory.C2S) (*qotgetmacroindicatorhistory.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetMacroIndicatorHistory: req is nil")
	}
	return qot.GetMacroIndicatorHistory(ctx, c.inner, req)
}

func GetFedWatchTargetRate(ctx context.Context, c *Client, req *qotgetfedwatchtargetrate.C2S) (*qotgetfedwatchtargetrate.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFedWatchTargetRate: req is nil")
	}
	return qot.GetFedWatchTargetRate(ctx, c.inner, req)
}

func GetFedWatchDotPlot(ctx context.Context, c *Client, req *qotgetfedwatchdotplot.C2S) (*qotgetfedwatchdotplot.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetFedWatchDotPlot: req is nil")
	}
	return qot.GetFedWatchDotPlot(ctx, c.inner, req)
}

func GetEarningsBeatRank(ctx context.Context, c *Client, req *qotgetearningsbeatrank.C2S) (*qotgetearningsbeatrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEarningsBeatRank: req is nil")
	}
	return qot.GetEarningsBeatRank(ctx, c.inner, req)
}

func GetDividendRank(ctx context.Context, c *Client, req *qotgetdividendrank.C2S) (*qotgetdividendrank.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetDividendRank: req is nil")
	}
	return qot.GetDividendRank(ctx, c.inner, req)
}

func GetDividendCalendar(ctx context.Context, c *Client, req *qotgetdividendcalendar.C2S) (*qotgetdividendcalendar.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetDividendCalendar: req is nil")
	}
	return qot.GetDividendCalendar(ctx, c.inner, req)
}

func GetEconomicCalendar(ctx context.Context, c *Client, req *qotgeteconomiccalendar.C2S) (*qotgeteconomiccalendar.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEconomicCalendar: req is nil")
	}
	return qot.GetEconomicCalendar(ctx, c.inner, req)
}
