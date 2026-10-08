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
	"testing"

	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/getoptionexpirationdate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotfiltercompetition"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetarkactivetransaction"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetarkfundholding"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetarkstockdynamic"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitalflow"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcodechange"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcompanyexecutivebackground"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcompanyexecutives"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcompanyoperationalefficiency"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcompanyprofile"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcorporateactionsbuybacks"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcorporateactionsdividends"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcorporateactionsstocksplits"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetdailyshortvolume"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetdividendcalendar"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetdividendrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetearningsbeatrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetearningscalendar"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteconomiccalendar"
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
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfedwatchdotplot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfedwatchtargetrate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialrevenuebreakdown"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialsearnpricehist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialsearnpricemove"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfinancialsstatements"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetfutureinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetheatmapdata"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgethighdividendsoerank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetholdingchangelist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgethotlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindicatorlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialchainbyplate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialchaindetail"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialchainlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialplateinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetindustrialplatestock"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinsiderholderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinsidertradelist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutiondistribution"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionholdingchange"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionholdinglist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetinstitutionprofile"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetipolist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetmacroindicatorhistory"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetmacroindicatorlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetmarketstate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionchain"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionearningsscreener"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionevent"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptioneventalert"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionexerciseprobability"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionmarketstatistic"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionquote"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionsellerscreener"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionstrategy"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionstrategyanalysis"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionstrategyspreads"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyinghisstatistic"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyinghisvolatility"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyingoverview"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionunderlyingrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionvolatility"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionzerodtecontract"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetoptionzerodtescreener"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetownerplate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetperiodchangerank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetplatesecurity"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetplateset"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetratingchange"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetreference"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetresearchanalystconsensus"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetresearchmorningstarrpt"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetresearchratingsummary"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrisefalldistr"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrt"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsearchnews"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsearchquote"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsecuritysnapshot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshareholdersholderdetail"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshareholdersholdingchanges"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshareholdersinstitutional"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshareholdersoverview"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshortinterest"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetshortsellingrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetstaticinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetsuspend"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgettopmoverrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgettoptenbuysellbrokers"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusafterhoursrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusersecuritygroup"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusovernightrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetuspremarketrank"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetvaluationdetail"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetvaluationplatestocklist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetwarrant"
	"github.com/shing1211/futuapi4go/pkg/pb/qotmodifyusersecurity"
	"github.com/shing1211/futuapi4go/pkg/pb/qotoptionscreen"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistoryeventcontractkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistorykl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistoryklquota"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequestindicatorcalc"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequestrehab"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesttradedate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsetoptioneventalert"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsetpricereminder"
	"github.com/shing1211/futuapi4go/pkg/pb/qotstockfilter"
	"github.com/shing1211/futuapi4go/pkg/pb/qotstockscreen"
	"github.com/shing1211/futuapi4go/pkg/pb/qotwarrantscreen"
)

func TestGen_GetIndustrialChainList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetIndustrialChainList, &qotgetindustrialchainlist.Response{})
	req := &qotgetindustrialchainlist.C2S{}
	fillNonZero(req)
	if _, err := GetIndustrialChainList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetIndustrialChainList: %v", err)
	}
}

func TestGen_GetIndustrialChainDetail(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetIndustrialChainDetail, &qotgetindustrialchaindetail.Response{})
	req := &qotgetindustrialchaindetail.C2S{}
	fillNonZero(req)
	if _, err := GetIndustrialChainDetail(context.Background(), cli, req); err != nil {
		t.Fatalf("GetIndustrialChainDetail: %v", err)
	}
}

func TestGen_GetIndustrialChainByPlate(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetIndustrialChainByPlate, &qotgetindustrialchainbyplate.Response{})
	req := &qotgetindustrialchainbyplate.C2S{}
	fillNonZero(req)
	if _, err := GetIndustrialChainByPlate(context.Background(), cli, req); err != nil {
		t.Fatalf("GetIndustrialChainByPlate: %v", err)
	}
}

func TestGen_GetIndustrialPlateInfo(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetIndustrialPlateInfo, &qotgetindustrialplateinfo.Response{})
	req := &qotgetindustrialplateinfo.C2S{}
	fillNonZero(req)
	if _, err := GetIndustrialPlateInfo(context.Background(), cli, req); err != nil {
		t.Fatalf("GetIndustrialPlateInfo: %v", err)
	}
}

func TestGen_GetIndustrialPlateStock(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetIndustrialPlateStock, &qotgetindustrialplatestock.Response{})
	req := &qotgetindustrialplatestock.C2S{}
	fillNonZero(req)
	if _, err := GetIndustrialPlateStock(context.Background(), cli, req); err != nil {
		t.Fatalf("GetIndustrialPlateStock: %v", err)
	}
}

func TestGen_FilterCompetition(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_FilterCompetition, &qotfiltercompetition.Response{})
	req := &qotfiltercompetition.C2S{}
	fillNonZero(req)
	if _, err := FilterCompetition(context.Background(), cli, req); err != nil {
		t.Fatalf("FilterCompetition: %v", err)
	}
}

func TestGen_GetEventContractCategory(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractCategory, &qotgeteventcontractcategory.Response{})
	req := &qotgeteventcontractcategory.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractCategory(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractCategory: %v", err)
	}
}

func TestGen_GetEventContractSeriesList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractSeriesList, &qotgeteventcontractserieslist.Response{})
	req := &qotgeteventcontractserieslist.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractSeriesList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractSeriesList: %v", err)
	}
}

func TestGen_GetEventContractEventList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractEventList, &qotgeteventcontracteventlist.Response{})
	req := &qotgeteventcontracteventlist.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractEventList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractEventList: %v", err)
	}
}

func TestGen_GetEventContract(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContract, &qotgeteventcontract.Response{})
	req := &qotgeteventcontract.C2S{}
	fillNonZero(req)
	if _, err := GetEventContract(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContract: %v", err)
	}
}

func TestGen_GetEventContractMilestoneList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractMilestoneList, &qotgeteventcontractmilestonelist.Response{})
	req := &qotgeteventcontractmilestonelist.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractMilestoneList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractMilestoneList: %v", err)
	}
}

func TestGen_GetEventContractSnapshot(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractSnapshot, &qotgeteventcontractsnapshot.Response{})
	req := &qotgeteventcontractsnapshot.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractSnapshot(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractSnapshot: %v", err)
	}
}

func TestGen_GetEventContractOrderBook(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractOrderBook, &qotgeteventcontractorderbook.Response{})
	req := &qotgeteventcontractorderbook.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractOrderBook(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractOrderBook: %v", err)
	}
}

func TestGen_GetEventContractKline(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractKline, &qotgeteventcontractkline.Response{})
	req := &qotgeteventcontractkline.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractKline(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractKline: %v", err)
	}
}

func TestGen_GetEventContractTicker(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractTicker, &qotgeteventcontractticker.Response{})
	req := &qotgeteventcontractticker.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractTicker(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractTicker: %v", err)
	}
}

func TestGen_RequestHistoryEventContractKL(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_RequestHistoryEventContractKL, &qotrequesthistoryeventcontractkl.Response{})
	req := &qotrequesthistoryeventcontractkl.C2S{}
	fillNonZero(req)
	if _, err := RequestHistoryEventContractKL(context.Background(), cli, req); err != nil {
		t.Fatalf("RequestHistoryEventContractKL: %v", err)
	}
}

func TestGen_GetEventContractComboList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractComboList, &qotgeteventcontractcombolist.Response{})
	req := &qotgeteventcontractcombolist.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractComboList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractComboList: %v", err)
	}
}

func TestGen_GetEventContractComboRfq(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEventContractComboRfq, &qotgeteventcontractcomborfq.Response{})
	req := &qotgeteventcontractcomborfq.C2S{}
	fillNonZero(req)
	if _, err := GetEventContractComboRfq(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEventContractComboRfq: %v", err)
	}
}

func TestGen_GetHeatMapData(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetHeatMapData, &qotgetheatmapdata.Response{})
	req := &qotgetheatmapdata.C2S{}
	fillNonZero(req)
	if _, err := GetHeatMapData(context.Background(), cli, req); err != nil {
		t.Fatalf("GetHeatMapData: %v", err)
	}
}

func TestGen_GetRiseFallDistribution(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetRiseFallDistribution, &qotgetrisefalldistr.Response{})
	req := &qotgetrisefalldistr.C2S{}
	fillNonZero(req)
	if _, err := GetRiseFallDistribution(context.Background(), cli, req); err != nil {
		t.Fatalf("GetRiseFallDistribution: %v", err)
	}
}

func TestGen_GetInstitutionList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetInstitutionList, &qotgetinstitutionlist.Response{})
	req := &qotgetinstitutionlist.C2S{}
	fillNonZero(req)
	if _, err := GetInstitutionList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetInstitutionList: %v", err)
	}
}

func TestGen_GetInstitutionProfile(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetInstitutionProfile, &qotgetinstitutionprofile.Response{})
	req := &qotgetinstitutionprofile.C2S{}
	fillNonZero(req)
	if _, err := GetInstitutionProfile(context.Background(), cli, req); err != nil {
		t.Fatalf("GetInstitutionProfile: %v", err)
	}
}

func TestGen_GetInstitutionDistribution(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetInstitutionDistribution, &qotgetinstitutiondistribution.Response{})
	req := &qotgetinstitutiondistribution.C2S{}
	fillNonZero(req)
	if _, err := GetInstitutionDistribution(context.Background(), cli, req); err != nil {
		t.Fatalf("GetInstitutionDistribution: %v", err)
	}
}

func TestGen_GetInstitutionHoldingChange(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetInstitutionHoldingChange, &qotgetinstitutionholdingchange.Response{})
	req := &qotgetinstitutionholdingchange.C2S{}
	fillNonZero(req)
	if _, err := GetInstitutionHoldingChange(context.Background(), cli, req); err != nil {
		t.Fatalf("GetInstitutionHoldingChange: %v", err)
	}
}

func TestGen_GetInstitutionHoldingList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetInstitutionHoldingList, &qotgetinstitutionholdinglist.Response{})
	req := &qotgetinstitutionholdinglist.C2S{}
	fillNonZero(req)
	if _, err := GetInstitutionHoldingList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetInstitutionHoldingList: %v", err)
	}
}

func TestGen_GetArkFundHolding(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetArkFundHolding, &qotgetarkfundholding.Response{})
	req := &qotgetarkfundholding.C2S{}
	fillNonZero(req)
	if _, err := GetArkFundHolding(context.Background(), cli, req); err != nil {
		t.Fatalf("GetArkFundHolding: %v", err)
	}
}

func TestGen_GetArkStockDynamic(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetArkStockDynamic, &qotgetarkstockdynamic.Response{})
	req := &qotgetarkstockdynamic.C2S{}
	fillNonZero(req)
	if _, err := GetArkStockDynamic(context.Background(), cli, req); err != nil {
		t.Fatalf("GetArkStockDynamic: %v", err)
	}
}

func TestGen_GetArkActiveTransaction(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetArkActiveTransaction, &qotgetarkactivetransaction.Response{})
	req := &qotgetarkactivetransaction.C2S{}
	fillNonZero(req)
	if _, err := GetArkActiveTransaction(context.Background(), cli, req); err != nil {
		t.Fatalf("GetArkActiveTransaction: %v", err)
	}
}

func TestGen_GetRatingChange(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetRatingChange, &qotgetratingchange.Response{})
	req := &qotgetratingchange.C2S{}
	fillNonZero(req)
	if _, err := GetRatingChange(context.Background(), cli, req); err != nil {
		t.Fatalf("GetRatingChange: %v", err)
	}
}

func TestGen_GetEarningsCalendar(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEarningsCalendar, &qotgetearningscalendar.Response{})
	req := &qotgetearningscalendar.C2S{}
	fillNonZero(req)
	if _, err := GetEarningsCalendar(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEarningsCalendar: %v", err)
	}
}

func TestGen_GetMacroIndicatorList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetMacroIndicatorList, &qotgetmacroindicatorlist.Response{})
	req := &qotgetmacroindicatorlist.C2S{}
	fillNonZero(req)
	if _, err := GetMacroIndicatorList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetMacroIndicatorList: %v", err)
	}
}

func TestGen_GetMacroIndicatorHistory(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetMacroIndicatorHistory, &qotgetmacroindicatorhistory.Response{})
	req := &qotgetmacroindicatorhistory.C2S{}
	fillNonZero(req)
	if _, err := GetMacroIndicatorHistory(context.Background(), cli, req); err != nil {
		t.Fatalf("GetMacroIndicatorHistory: %v", err)
	}
}

func TestGen_GetFedWatchTargetRate(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetFedWatchTargetRate, &qotgetfedwatchtargetrate.Response{})
	req := &qotgetfedwatchtargetrate.C2S{}
	fillNonZero(req)
	if _, err := GetFedWatchTargetRate(context.Background(), cli, req); err != nil {
		t.Fatalf("GetFedWatchTargetRate: %v", err)
	}
}

func TestGen_GetFedWatchDotPlot(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetFedWatchDotPlot, &qotgetfedwatchdotplot.Response{})
	req := &qotgetfedwatchdotplot.C2S{}
	fillNonZero(req)
	if _, err := GetFedWatchDotPlot(context.Background(), cli, req); err != nil {
		t.Fatalf("GetFedWatchDotPlot: %v", err)
	}
}

func TestGen_GetEarningsBeatRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEarningsBeatRank, &qotgetearningsbeatrank.Response{})
	req := &qotgetearningsbeatrank.C2S{}
	fillNonZero(req)
	if _, err := GetEarningsBeatRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEarningsBeatRank: %v", err)
	}
}

func TestGen_GetDividendRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetDividendRank, &qotgetdividendrank.Response{})
	req := &qotgetdividendrank.C2S{}
	fillNonZero(req)
	if _, err := GetDividendRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetDividendRank: %v", err)
	}
}

func TestGen_GetDividendCalendar(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetDividendCalendar, &qotgetdividendcalendar.Response{})
	req := &qotgetdividendcalendar.C2S{}
	fillNonZero(req)
	if _, err := GetDividendCalendar(context.Background(), cli, req); err != nil {
		t.Fatalf("GetDividendCalendar: %v", err)
	}
}

func TestGen_GetEconomicCalendar(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetEconomicCalendar, &qotgeteconomiccalendar.Response{})
	req := &qotgeteconomiccalendar.C2S{}
	fillNonZero(req)
	if _, err := GetEconomicCalendar(context.Background(), cli, req); err != nil {
		t.Fatalf("GetEconomicCalendar: %v", err)
	}
}

func TestGen_GetOptionMarketStatistic(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionMarketStatistic, &qotgetoptionmarketstatistic.Response{})
	req := &qotgetoptionmarketstatistic.C2S{}
	fillNonZero(req)
	if _, err := GetOptionMarketStatistic(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionMarketStatistic: %v", err)
	}
}

func TestGen_GetOptionUnderlyingHisStatistic(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionUnderlyingHisStatistic, &qotgetoptionunderlyinghisstatistic.Response{})
	req := &qotgetoptionunderlyinghisstatistic.C2S{}
	fillNonZero(req)
	if _, err := GetOptionUnderlyingHisStatistic(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionUnderlyingHisStatistic: %v", err)
	}
}

func TestGen_GetOptionUnderlyingOverview(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionUnderlyingOverview, &qotgetoptionunderlyingoverview.Response{})
	req := &qotgetoptionunderlyingoverview.C2S{}
	fillNonZero(req)
	if _, err := GetOptionUnderlyingOverview(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionUnderlyingOverview: %v", err)
	}
}

func TestGen_GetOptionUnderlyingHisVolatility(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionUnderlyingHisVolatility, &qotgetoptionunderlyinghisvolatility.Response{})
	req := &qotgetoptionunderlyinghisvolatility.C2S{}
	fillNonZero(req)
	if _, err := GetOptionUnderlyingHisVolatility(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionUnderlyingHisVolatility: %v", err)
	}
}

func TestGen_GetOptionUnderlyingRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionUnderlyingRank, &qotgetoptionunderlyingrank.Response{})
	req := &qotgetoptionunderlyingrank.C2S{}
	fillNonZero(req)
	if _, err := GetOptionUnderlyingRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionUnderlyingRank: %v", err)
	}
}

func TestGen_GetOptionRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionRank, &qotgetoptionrank.Response{})
	req := &qotgetoptionrank.C2S{}
	fillNonZero(req)
	if _, err := GetOptionRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionRank: %v", err)
	}
}

func TestGen_GetOptionEvent(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionEvent, &qotgetoptionevent.Response{})
	req := &qotgetoptionevent.C2S{}
	fillNonZero(req)
	if _, err := GetOptionEvent(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionEvent: %v", err)
	}
}

func TestGen_GetOptionEventAlert(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionEventAlert, &qotgetoptioneventalert.Response{})
	req := &qotgetoptioneventalert.C2S{}
	fillNonZero(req)
	if _, err := GetOptionEventAlert(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionEventAlert: %v", err)
	}
}

func TestGen_SetOptionEventAlert(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_SetOptionEventAlert, &qotsetoptioneventalert.Response{})
	req := &qotsetoptioneventalert.C2S{}
	fillNonZero(req)
	if _, err := SetOptionEventAlert(context.Background(), cli, req); err != nil {
		t.Fatalf("SetOptionEventAlert: %v", err)
	}
}

func TestGen_GetOptionZeroDteScreener(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionZeroDteScreener, &qotgetoptionzerodtescreener.Response{})
	req := &qotgetoptionzerodtescreener.C2S{}
	fillNonZero(req)
	if _, err := GetOptionZeroDteScreener(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionZeroDteScreener: %v", err)
	}
}

func TestGen_GetOptionZeroDteContract(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionZeroDteContract, &qotgetoptionzerodtecontract.Response{})
	req := &qotgetoptionzerodtecontract.C2S{}
	fillNonZero(req)
	if _, err := GetOptionZeroDteContract(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionZeroDteContract: %v", err)
	}
}

func TestGen_GetOptionEarningsScreener(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionEarningsScreener, &qotgetoptionearningsscreener.Response{})
	req := &qotgetoptionearningsscreener.C2S{}
	fillNonZero(req)
	if _, err := GetOptionEarningsScreener(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionEarningsScreener: %v", err)
	}
}

func TestGen_GetOptionSellerScreener(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetOptionSellerScreener, &qotgetoptionsellerscreener.Response{})
	req := &qotgetoptionsellerscreener.C2S{}
	fillNonZero(req)
	if _, err := GetOptionSellerScreener(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionSellerScreener: %v", err)
	}
}

func TestGen_GetUSPreMarketRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetUSPreMarketRank, &qotgetuspremarketrank.Response{})
	req := &qotgetuspremarketrank.C2S{}
	fillNonZero(req)
	if _, err := GetUSPreMarketRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetUSPreMarketRank: %v", err)
	}
}

func TestGen_GetUSAfterHoursRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetUSAfterHoursRank, &qotgetusafterhoursrank.Response{})
	req := &qotgetusafterhoursrank.C2S{}
	fillNonZero(req)
	if _, err := GetUSAfterHoursRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetUSAfterHoursRank: %v", err)
	}
}

func TestGen_GetUSOvernightRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetUSOvernightRank, &qotgetusovernightrank.Response{})
	req := &qotgetusovernightrank.C2S{}
	fillNonZero(req)
	if _, err := GetUSOvernightRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetUSOvernightRank: %v", err)
	}
}

func TestGen_GetTopMoversRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetTopMoversRank, &qotgettopmoverrank.Response{})
	req := &qotgettopmoverrank.C2S{}
	fillNonZero(req)
	if _, err := GetTopMoversRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetTopMoversRank: %v", err)
	}
}

func TestGen_GetHotList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetHotList, &qotgethotlist.Response{})
	req := &qotgethotlist.C2S{}
	fillNonZero(req)
	if _, err := GetHotList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetHotList: %v", err)
	}
}

func TestGen_GetShortSellingRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetShortSellingRank, &qotgetshortsellingrank.Response{})
	req := &qotgetshortsellingrank.C2S{}
	fillNonZero(req)
	if _, err := GetShortSellingRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetShortSellingRank: %v", err)
	}
}

func TestGen_GetPeriodChangeRank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetPeriodChangeRank, &qotgetperiodchangerank.Response{})
	req := &qotgetperiodchangerank.C2S{}
	fillNonZero(req)
	if _, err := GetPeriodChangeRank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetPeriodChangeRank: %v", err)
	}
}

func TestGen_GetHighDividendSOERank(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetHighDividendSOERank, &qotgethighdividendsoerank.Response{})
	req := &qotgethighdividendsoerank.C2S{}
	fillNonZero(req)
	if _, err := GetHighDividendSOERank(context.Background(), cli, req); err != nil {
		t.Fatalf("GetHighDividendSOERank: %v", err)
	}
}

func TestGen_GetSearchQuote(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetSearchQuote, &qotgetsearchquote.Response{})
	req := &qotgetsearchquote.C2S{}
	fillNonZero(req)
	if _, err := GetSearchQuote(context.Background(), cli, req); err != nil {
		t.Fatalf("GetSearchQuote: %v", err)
	}
}

func TestGen_GetSearchNews(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetSearchNews, &qotgetsearchnews.Response{})
	req := &qotgetsearchnews.C2S{}
	fillNonZero(req)
	if _, err := GetSearchNews(context.Background(), cli, req); err != nil {
		t.Fatalf("GetSearchNews: %v", err)
	}
}

func TestGen_GetIndicatorList(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetIndicatorList, &qotgetindicatorlist.Response{})
	req := &qotgetindicatorlist.C2S{}
	fillNonZero(req)
	if _, err := GetIndicatorList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetIndicatorList: %v", err)
	}
}

func TestGen_RequestIndicatorCalc(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_RequestIndicatorCalc, &qotrequestindicatorcalc.Response{})
	req := &qotrequestindicatorcalc.C2S{}
	fillNonZero(req)
	if _, err := RequestIndicatorCalc(context.Background(), cli, req); err != nil {
		t.Fatalf("RequestIndicatorCalc: %v", err)
	}
}

func TestWrap_GetCapitalFlow(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetCapitalFlow, &qotgetcapitalflow.Response{})
	req := &GetCapitalFlowRequest{}
	fillNonZero(req)
	if _, err := GetCapitalFlow(context.Background(), cli, req); err != nil {
		t.Fatalf("GetCapitalFlow: %v", err)
	}
}

func TestWrap_GetCompanyProfile(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetCompanyProfile, &qotgetcompanyprofile.Response{})
	req := &GetCompanyProfileRequest{}
	fillNonZero(req)
	if _, err := GetCompanyProfile(context.Background(), cli, req); err != nil {
		t.Fatalf("GetCompanyProfile: %v", err)
	}
}

func TestWrap_GetCompanyExecutives(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetCompanyExecutives, &qotgetcompanyexecutives.Response{})
	req := &GetCompanyExecutivesRequest{}
	fillNonZero(req)
	if _, err := GetCompanyExecutives(context.Background(), cli, req); err != nil {
		t.Fatalf("GetCompanyExecutives: %v", err)
	}
}

func TestWrap_GetCompanyExecutiveBackground(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetCompanyExecutiveBackground, &qotgetcompanyexecutivebackground.Response{})
	req := &GetCompanyExecutiveBackgroundRequest{}
	fillNonZero(req)
	if _, err := GetCompanyExecutiveBackground(context.Background(), cli, req); err != nil {
		t.Fatalf("GetCompanyExecutiveBackground: %v", err)
	}
}

func TestWrap_GetCompanyOperationalEfficiency(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetCompanyOperationalEfficiency, &qotgetcompanyoperationalefficiency.Response{})
	req := &GetCompanyOperationalEfficiencyRequest{}
	fillNonZero(req)
	if _, err := GetCompanyOperationalEfficiency(context.Background(), cli, req); err != nil {
		t.Fatalf("GetCompanyOperationalEfficiency: %v", err)
	}
}

func TestWrap_GetCorporateActionsDividends(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetCorporateActionsDividends, &qotgetcorporateactionsdividends.Response{})
	req := &GetCorporateActionsDividendsRequest{}
	fillNonZero(req)
	if _, err := GetCorporateActionsDividends(context.Background(), cli, req); err != nil {
		t.Fatalf("GetCorporateActionsDividends: %v", err)
	}
}

func TestWrap_GetCorporateActionsBuybacks(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetCorporateActionsBuybacks, &qotgetcorporateactionsbuybacks.Response{})
	req := &GetCorporateActionsBuybacksRequest{}
	fillNonZero(req)
	if _, err := GetCorporateActionsBuybacks(context.Background(), cli, req); err != nil {
		t.Fatalf("GetCorporateActionsBuybacks: %v", err)
	}
}

func TestWrap_GetCorporateActionsStockSplits(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetCorporateActionsStockSplits, &qotgetcorporateactionsstocksplits.Response{})
	req := &GetCorporateActionsStockSplitsRequest{}
	fillNonZero(req)
	if _, err := GetCorporateActionsStockSplits(context.Background(), cli, req); err != nil {
		t.Fatalf("GetCorporateActionsStockSplits: %v", err)
	}
}

func TestWrap_StockFilter(t *testing.T) {
	cli := qotTestClient(t, ProtoID_StockFilter, &qotstockfilter.Response{})
	req := &StockFilterRequest{}
	fillNonZero(req)
	if _, err := StockFilter(context.Background(), cli, req); err != nil {
		t.Fatalf("StockFilter: %v", err)
	}
}

func TestWrap_GetWarrant(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetWarrant, &qotgetwarrant.Response{})
	req := &GetWarrantRequest{}
	fillNonZero(req)
	if _, err := GetWarrant(context.Background(), cli, req); err != nil {
		t.Fatalf("GetWarrant: %v", err)
	}
}

func TestWrap_GetReference(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetReference, &qotgetreference.Response{})
	req := &GetReferenceRequest{}
	fillNonZero(req)
	if _, err := GetReference(context.Background(), cli, req); err != nil {
		t.Fatalf("GetReference: %v", err)
	}
}

func TestWrap_GetFinancialsStatements(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetFinancialsStatements, &qotgetfinancialsstatements.Response{})
	req := &GetFinancialsStatementsRequest{}
	fillNonZero(req)
	if _, err := GetFinancialsStatements(context.Background(), cli, req); err != nil {
		t.Fatalf("GetFinancialsStatements: %v", err)
	}
}

func TestWrap_GetFinancialsRevenueBreakdown(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetFinancialsRevenueBreakdown, &qotgetfinancialrevenuebreakdown.Response{})
	req := &GetFinancialsRevenueBreakdownRequest{}
	fillNonZero(req)
	if _, err := GetFinancialsRevenueBreakdown(context.Background(), cli, req); err != nil {
		t.Fatalf("GetFinancialsRevenueBreakdown: %v", err)
	}
}

func TestWrap_GetFinancialsEarningsPriceMove(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetFinancialsEarningsPriceMove, &qotgetfinancialsearnpricemove.Response{})
	req := &GetFinancialsEarningsPriceMoveRequest{}
	fillNonZero(req)
	if _, err := GetFinancialsEarningsPriceMove(context.Background(), cli, req); err != nil {
		t.Fatalf("GetFinancialsEarningsPriceMove: %v", err)
	}
}

func TestWrap_GetFinancialsEarningsPriceHistory(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_GetFinancialsEarningsPriceHistory, &qotgetfinancialsearnpricehist.Response{})
	req := &GetFinancialsEarningsPriceHistoryRequest{}
	fillNonZero(req)
	if _, err := GetFinancialsEarningsPriceHistory(context.Background(), cli, req); err != nil {
		t.Fatalf("GetFinancialsEarningsPriceHistory: %v", err)
	}
}

func TestWrap_GetHoldingChangeList(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetHoldingChangeList, &qotgetholdingchangelist.Response{})
	req := &GetHoldingChangeListRequest{}
	fillNonZero(req)
	if _, err := GetHoldingChangeList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetHoldingChangeList: %v", err)
	}
}

func TestWrap_RequestRehab(t *testing.T) {
	cli := qotTestClient(t, ProtoID_RequestRehab, &qotrequestrehab.Response{})
	req := &RequestRehabRequest{}
	fillNonZero(req)
	if _, err := RequestRehab(context.Background(), cli, req); err != nil {
		t.Fatalf("RequestRehab: %v", err)
	}
}

func TestWrap_GetInsiderHolderList(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetInsiderHolderList, &qotgetinsiderholderlist.Response{})
	req := &GetInsiderHolderListRequest{}
	fillNonZero(req)
	if _, err := GetInsiderHolderList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetInsiderHolderList: %v", err)
	}
}

func TestWrap_GetInsiderTradeList(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetInsiderTradeList, &qotgetinsidertradelist.Response{})
	req := &GetInsiderTradeListRequest{}
	fillNonZero(req)
	if _, err := GetInsiderTradeList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetInsiderTradeList: %v", err)
	}
}

func TestWrap_RequestHistoryKL(t *testing.T) {
	cli := qotTestClient(t, ProtoID_RequestHistoryKL, &qotrequesthistorykl.Response{})
	req := &RequestHistoryKLRequest{}
	fillNonZero(req)
	if _, err := RequestHistoryKL(context.Background(), cli, req); err != nil {
		t.Fatalf("RequestHistoryKL: %v", err)
	}
}

func TestWrap_RequestHistoryKLQuota(t *testing.T) {
	cli := qotTestClient(t, ProtoID_RequestHistoryKLQuota, &qotrequesthistoryklquota.Response{})
	req := &RequestHistoryKLQuotaRequest{}
	fillNonZero(req)
	if _, err := RequestHistoryKLQuota(context.Background(), cli, req); err != nil {
		t.Fatalf("RequestHistoryKLQuota: %v", err)
	}
}

func TestWrap_GetOrderBook(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOrderBook, &qotgetorderbook.Response{})
	req := &GetOrderBookRequest{}
	fillNonZero(req)
	if _, err := GetOrderBook(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOrderBook: %v", err)
	}
}

func TestWrap_GetTicker(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetTicker, &qotgetticker.Response{})
	req := &GetTickerRequest{}
	fillNonZero(req)
	if _, err := GetTicker(context.Background(), cli, req); err != nil {
		t.Fatalf("GetTicker: %v", err)
	}
}

func TestWrap_GetRT(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetRT, &qotgetrt.Response{})
	req := &GetRTRequest{}
	fillNonZero(req)
	if _, err := GetRT(context.Background(), cli, req); err != nil {
		t.Fatalf("GetRT: %v", err)
	}
}

func TestWrap_GetBroker(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetBroker, &qotgetbroker.Response{})
	req := &GetBrokerRequest{}
	fillNonZero(req)
	if _, err := GetBroker(context.Background(), cli, req); err != nil {
		t.Fatalf("GetBroker: %v", err)
	}
}

func TestWrap_RequestTradeDate(t *testing.T) {
	cli := qotTestClient(t, ProtoID_RequestTradeDate, &qotrequesttradedate.Response{})
	req := &RequestTradeDateRequest{}
	fillNonZero(req)
	if _, err := RequestTradeDate(context.Background(), cli, req); err != nil {
		t.Fatalf("RequestTradeDate: %v", err)
	}
}

func TestWrap_GetSuspend(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetSuspend, &qotgetsuspend.Response{})
	req := &GetSuspendRequest{}
	fillNonZero(req)
	if _, err := GetSuspend(context.Background(), cli, req); err != nil {
		t.Fatalf("GetSuspend: %v", err)
	}
}

func TestWrap_GetCodeChange(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetCodeChange, &qotgetcodechange.Response{})
	req := &GetCodeChangeRequest{}
	fillNonZero(req)
	if _, err := GetCodeChange(context.Background(), cli, req); err != nil {
		t.Fatalf("GetCodeChange: %v", err)
	}
}

func TestWrap_GetMarketState(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetMarketState, &qotgetmarketstate.Response{})
	req := &GetMarketStateRequest{}
	fillNonZero(req)
	if _, err := GetMarketState(context.Background(), cli, req); err != nil {
		t.Fatalf("GetMarketState: %v", err)
	}
}

func TestWrap_GetOptionVolatility(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOptionVolatility, &qotgetoptionvolatility.Response{})
	req := &GetOptionVolatilityRequest{}
	fillNonZero(req)
	if _, err := GetOptionVolatility(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionVolatility: %v", err)
	}
}

func TestWrap_GetOptionExerciseProbability(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOptionExerciseProbability, &qotgetoptionexerciseprobability.Response{})
	req := &GetOptionExerciseProbabilityRequest{}
	fillNonZero(req)
	if _, err := GetOptionExerciseProbability(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionExerciseProbability: %v", err)
	}
}

func TestWrap_GetOptionExpirationDate(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOptionExpirationDate, &getoptionexpirationdate.Response{})
	req := &GetOptionExpirationDateRequest{}
	fillNonZero(req)
	if _, err := GetOptionExpirationDate(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionExpirationDate: %v", err)
	}
}

func TestWrap_GetOptionChain(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOptionChain, &qotgetoptionchain.Response{})
	req := &GetOptionChainRequest{}
	fillNonZero(req)
	if _, err := GetOptionChain(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionChain: %v", err)
	}
}

func TestWrap_GetFutureInfo(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetFutureInfo, &qotgetfutureinfo.Response{})
	req := &GetFutureInfoRequest{}
	fillNonZero(req)
	if _, err := GetFutureInfo(context.Background(), cli, req); err != nil {
		t.Fatalf("GetFutureInfo: %v", err)
	}
}

func TestWrap_GetPlateSet(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetPlateSet, &qotgetplateset.Response{})
	req := &GetPlateSetRequest{}
	fillNonZero(req)
	if _, err := GetPlateSet(context.Background(), cli, req); err != nil {
		t.Fatalf("GetPlateSet: %v", err)
	}
}

func TestWrap_GetPlateSecurity(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetPlateSecurity, &qotgetplatesecurity.Response{})
	req := &GetPlateSecurityRequest{}
	fillNonZero(req)
	if _, err := GetPlateSecurity(context.Background(), cli, req); err != nil {
		t.Fatalf("GetPlateSecurity: %v", err)
	}
}

func TestWrap_GetOwnerPlate(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOwnerPlate, &qotgetownerplate.Response{})
	req := &GetOwnerPlateRequest{}
	fillNonZero(req)
	if _, err := GetOwnerPlate(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOwnerPlate: %v", err)
	}
}

func TestWrap_GetOptionQuote(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOptionQuote, &qotgetoptionquote.Response{})
	req := &GetOptionQuoteRequest{}
	fillNonZero(req)
	if _, err := GetOptionQuote(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionQuote: %v", err)
	}
}

func TestWrap_GetOptionStrategy(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOptionStrategy, &qotgetoptionstrategy.Response{})
	req := &GetOptionStrategyRequest{}
	fillNonZero(req)
	if _, err := GetOptionStrategy(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionStrategy: %v", err)
	}
}

func TestWrap_GetOptionStrategyAnalysis(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOptionStrategyAnalysis, &qotgetoptionstrategyanalysis.Response{})
	req := &GetOptionStrategyAnalysisRequest{}
	fillNonZero(req)
	if _, err := GetOptionStrategyAnalysis(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionStrategyAnalysis: %v", err)
	}
}

func TestWrap_GetOptionStrategySpread(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetOptionStrategySpread, &qotgetoptionstrategyspreads.Response{})
	req := &GetOptionStrategySpreadRequest{}
	fillNonZero(req)
	if _, err := GetOptionStrategySpread(context.Background(), cli, req); err != nil {
		t.Fatalf("GetOptionStrategySpread: %v", err)
	}
}

func TestWrap_GetKL(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetKL, &qotgetkl.Response{})
	req := &GetKLRequest{}
	fillNonZero(req)
	if _, err := GetKL(context.Background(), cli, req); err != nil {
		t.Fatalf("GetKL: %v", err)
	}
}

func TestWrap_GetResearchAnalystConsensus(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetResearchAnalystConsensus, &qotgetresearchanalystconsensus.Response{})
	req := &GetResearchAnalystConsensusRequest{}
	fillNonZero(req)
	if _, err := GetResearchAnalystConsensus(context.Background(), cli, req); err != nil {
		t.Fatalf("GetResearchAnalystConsensus: %v", err)
	}
}

func TestWrap_GetResearchRatingSummary(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetResearchRatingSummary, &qotgetresearchratingsummary.Response{})
	req := &GetResearchRatingSummaryRequest{}
	fillNonZero(req)
	if _, err := GetResearchRatingSummary(context.Background(), cli, req); err != nil {
		t.Fatalf("GetResearchRatingSummary: %v", err)
	}
}

func TestWrap_GetResearchMorningstarReport(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetResearchMorningstarReport, &qotgetresearchmorningstarrpt.Response{})
	req := &GetResearchMorningstarReportRequest{}
	fillNonZero(req)
	if _, err := GetResearchMorningstarReport(context.Background(), cli, req); err != nil {
		t.Fatalf("GetResearchMorningstarReport: %v", err)
	}
}

func TestWrap_StockScreen(t *testing.T) {
	cli := qotTestClient(t, ProtoID_StockScreen, &qotstockscreen.Response{})
	req := &StockScreenRequest{}
	fillNonZero(req)
	if _, err := StockScreen(context.Background(), cli, req); err != nil {
		t.Fatalf("StockScreen: %v", err)
	}
}

func TestWrap_WarrantScreen(t *testing.T) {
	cli := qotTestClient(t, ProtoID_WarrantScreen, &qotwarrantscreen.Response{})
	req := &WarrantScreenRequest{}
	fillNonZero(req)
	if _, err := WarrantScreen(context.Background(), cli, req); err != nil {
		t.Fatalf("WarrantScreen: %v", err)
	}
}

func TestWrap_OptionScreen(t *testing.T) {
	cli := qotTestClient(t, ProtoID_OptionScreen, &qotoptionscreen.Response{})
	req := &OptionScreenRequest{}
	fillNonZero(req)
	if _, err := OptionScreen(context.Background(), cli, req); err != nil {
		t.Fatalf("OptionScreen: %v", err)
	}
}

func TestWrap_GetShareholdersOverview(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetShareholdersOverview, &qotgetshareholdersoverview.Response{})
	req := &GetShareholdersOverviewRequest{}
	fillNonZero(req)
	if _, err := GetShareholdersOverview(context.Background(), cli, req); err != nil {
		t.Fatalf("GetShareholdersOverview: %v", err)
	}
}

func TestWrap_GetShareholdersHoldingChanges(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetShareholdersHoldingChanges, &qotgetshareholdersholdingchanges.Response{})
	req := &GetShareholdersHoldingChangesRequest{}
	fillNonZero(req)
	if _, err := GetShareholdersHoldingChanges(context.Background(), cli, req); err != nil {
		t.Fatalf("GetShareholdersHoldingChanges: %v", err)
	}
}

func TestWrap_GetShareholdersHolderDetail(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetShareholdersHolderDetail, &qotgetshareholdersholderdetail.Response{})
	req := &GetShareholdersHolderDetailRequest{}
	fillNonZero(req)
	if _, err := GetShareholdersHolderDetail(context.Background(), cli, req); err != nil {
		t.Fatalf("GetShareholdersHolderDetail: %v", err)
	}
}

func TestWrap_GetShareholdersInstitutional(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetShareholdersInstitutional, &qotgetshareholdersinstitutional.Response{})
	req := &GetShareholdersInstitutionalRequest{}
	fillNonZero(req)
	if _, err := GetShareholdersInstitutional(context.Background(), cli, req); err != nil {
		t.Fatalf("GetShareholdersInstitutional: %v", err)
	}
}

func TestWrap_GetTopTenBuySellBrokers(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetTopTenBuySellBrokers, &qotgettoptenbuysellbrokers.Response{})
	req := &GetTopTenBuySellBrokersRequest{}
	fillNonZero(req)
	if _, err := GetTopTenBuySellBrokers(context.Background(), cli, req); err != nil {
		t.Fatalf("GetTopTenBuySellBrokers: %v", err)
	}
}

func TestWrap_GetDailyShortVolume(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetDailyShortVolume, &qotgetdailyshortvolume.Response{})
	req := &GetDailyShortVolumeRequest{}
	fillNonZero(req)
	if _, err := GetDailyShortVolume(context.Background(), cli, req); err != nil {
		t.Fatalf("GetDailyShortVolume: %v", err)
	}
}

func TestWrap_GetShortInterest(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetShortInterest, &qotgetshortinterest.Response{})
	req := &GetShortInterestRequest{}
	fillNonZero(req)
	if _, err := GetShortInterest(context.Background(), cli, req); err != nil {
		t.Fatalf("GetShortInterest: %v", err)
	}
}

func TestWrap_GetSecuritySnapshot(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetSecuritySnapshot, &qotgetsecuritysnapshot.Response{})
	req := &GetSecuritySnapshotRequest{}
	fillNonZero(req)
	if _, err := GetSecuritySnapshot(context.Background(), cli, req); err != nil {
		t.Fatalf("GetSecuritySnapshot: %v", err)
	}
}

func TestWrap_GetStaticInfo(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetStaticInfo, &qotgetstaticinfo.Response{})
	req := &GetStaticInfoRequest{}
	fillNonZero(req)
	if _, err := GetStaticInfo(context.Background(), cli, req); err != nil {
		t.Fatalf("GetStaticInfo: %v", err)
	}
}

func TestWrap_GetTradeDate(t *testing.T) {
	cli := qotTestClient(t, constant.ProtoID_Qot_RequestTradeDate, &qotrequesttradedate.Response{})
	req := &GetTradeDateRequest{}
	fillNonZero(req)
	if _, err := GetTradeDate(context.Background(), cli, req); err != nil {
		t.Fatalf("GetTradeDate: %v", err)
	}
}

func TestWrap_GetUserSecurityGroup(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetUserSecurityGroup, &qotgetusersecuritygroup.Response{})
	req := &GetUserSecurityGroupRequest{}
	fillNonZero(req)
	if _, err := GetUserSecurityGroup(context.Background(), cli, req); err != nil {
		t.Fatalf("GetUserSecurityGroup: %v", err)
	}
}

func TestWrap_ModifyUserSecurity(t *testing.T) {
	cli := qotTestClient(t, ProtoID_ModifyUserSecurity, &qotmodifyusersecurity.Response{})
	req := &ModifyUserSecurityRequest{}
	fillNonZero(req)
	if _, err := ModifyUserSecurity(context.Background(), cli, req); err != nil {
		t.Fatalf("ModifyUserSecurity: %v", err)
	}
}

func TestWrap_SetPriceReminder(t *testing.T) {
	cli := qotTestClient(t, ProtoID_SetPriceReminder, &qotsetpricereminder.Response{})
	req := &SetPriceReminderRequest{}
	fillNonZero(req)
	if _, err := SetPriceReminder(context.Background(), cli, req); err != nil {
		t.Fatalf("SetPriceReminder: %v", err)
	}
}

func TestWrap_GetIpoList(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetIpoList, &qotgetipolist.Response{})
	req := &GetIpoListRequest{}
	fillNonZero(req)
	if _, err := GetIpoList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetIpoList: %v", err)
	}
}

func TestWrap_GetValuationDetail(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetValuationDetail, &qotgetvaluationdetail.Response{})
	req := &GetValuationDetailRequest{}
	fillNonZero(req)
	if _, err := GetValuationDetail(context.Background(), cli, req); err != nil {
		t.Fatalf("GetValuationDetail: %v", err)
	}
}

func TestWrap_GetValuationPlateStockList(t *testing.T) {
	cli := qotTestClient(t, ProtoID_GetValuationPlateStockList, &qotgetvaluationplatestocklist.Response{})
	req := &GetValuationPlateStockListRequest{}
	fillNonZero(req)
	if _, err := GetValuationPlateStockList(context.Background(), cli, req); err != nil {
		t.Fatalf("GetValuationPlateStockList: %v", err)
	}
}
