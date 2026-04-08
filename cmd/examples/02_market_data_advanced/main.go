// Advanced Market Data Examples
//
// This example demonstrates advanced market data APIs:
// - GetStaticInfo: Stock static information
// - GetPlateSet: Plate/sector information
// - GetCapitalFlow: Capital flow analysis
// - GetCapitalDistribution: Capital distribution
// - StockFilter: Screen stocks by criteria
// - GetOptionChain: Options chain data
// - GetWarrant: Warrant information
// - GetTradeDate: Trading calendar
// - GetFutureInfo: Futures data
// - GetIpoList: IPO listings
//
// Usage:
//   go run main.go

package main

import (
	"fmt"
	"log"

	futuapi "gitee.com/shing1211/futuapi4go/internal/client"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pkg/pb/qotstockfilter"
	"gitee.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	// Create and connect client
	cli := futuapi.New()
	defer cli.Close()

	addr := "127.0.0.1:11111"
	fmt.Printf("=== Connecting to %s ===\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✓ Connected! ConnID=%d\n\n", cli.GetConnID())

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	_ = int32(qotcommon.QotMarket_QotMarket_US_Security)
	_ = int32(qotcommon.QotMarket_QotMarket_CNSH_Security)

	// 1. Get Stock Static Info
	fmt.Println("=== 1. Stock Static Info (GetStaticInfo) ===")
	staticReq := &qot.GetStaticInfoRequest{
		Market:  hkMarket,
		SecType: int32(qotcommon.SecurityType_SecurityType_Eqty),
	}

	staticResp, err := qot.GetStaticInfo(cli, staticReq)
	if err != nil {
		log.Printf("GetStaticInfo failed: %v", err)
	} else {
		fmt.Printf("  Found %d securities\n", len(staticResp.StaticInfoList))
		// Show first 5
		count := 5
		if len(staticResp.StaticInfoList) < count {
			count = len(staticResp.StaticInfoList)
		}
		for i := 0; i < count; i++ {
			info := staticResp.StaticInfoList[i]
			basic := info.GetBasic()
			fmt.Printf("  %s | %s | LotSize=%d | ListDate=%s\n",
				basic.GetSecurity().GetCode(), basic.GetName(), basic.GetLotSize(), basic.GetListTime())
		}
	}
	fmt.Println()

	// 2. Get Plate/Sector List
	fmt.Println("=== 2. Plate/Sector List (GetPlateSet) ===")
	plateReq := &qot.GetPlateSetRequest{
		Market: hkMarket,
	}

	plateResp, err := qot.GetPlateSet(cli, plateReq)
	if err != nil {
		log.Printf("GetPlateSet failed: %v", err)
	} else {
		fmt.Printf("  Found %d plates\n", len(plateResp.PlateSetList))
		for i, plate := range plateResp.PlateSetList {
			if i >= 5 {
				fmt.Printf("  ... and %d more\n", len(plateResp.PlateSetList)-5)
				break
			}
			fmt.Printf("  %s | Code=%s\n", plate.Name, plate.Plate.GetCode())
		}
	}
	fmt.Println()

	// 3. Get Capital Flow
	fmt.Println("=== 3. Capital Flow (GetCapitalFlow) ===")
	security := &qotcommon.Security{
		Market: &hkMarket,
		Code:   ptrStr("00700"),
	}

	capFlowReq := &qot.GetCapitalFlowRequest{
		Security:   security,
		PeriodType: 1, // Daily
	}

	capFlowResp, err := qot.GetCapitalFlow(cli, capFlowReq)
	if err != nil {
		log.Printf("GetCapitalFlow failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s\n", security.GetCode())
		fmt.Printf("  %-20s %-15s %-15s\n",
			"Time", "InFlow", "Main In")
		for _, flow := range capFlowResp.FlowItemList {
			fmt.Printf("  %-20s %-15.2f %-15.2f\n",
				flow.Time, flow.InFlow, flow.MainInFlow)
		}
	}
	fmt.Println()

	// 4. Get Capital Distribution
	fmt.Println("=== 4. Capital Distribution (GetCapitalDistribution) ===")
	capDistResp, err := qot.GetCapitalDistribution(cli, security)
	if err != nil {
		log.Printf("GetCapitalDistribution failed: %v", err)
	} else {
		cd := capDistResp.CapitalDistribution
		fmt.Printf("  Stock: %s\n", security.GetCode())
		fmt.Printf("  Super In:  %.2f\n", cd.CapitalInSuper)
		fmt.Printf("  Super Out: %.2f\n", cd.CapitalOutSuper)
		fmt.Printf("  Big In:    %.2f\n", cd.CapitalInBig)
		fmt.Printf("  Big Out:   %.2f\n", cd.CapitalOutBig)
		fmt.Printf("  Mid In:    %.2f\n", cd.CapitalInMid)
		fmt.Printf("  Mid Out:   %.2f\n", cd.CapitalOutMid)
		fmt.Printf("  Small In:  %.2f\n", cd.CapitalInSmall)
		fmt.Printf("  Small Out: %.2f\n", cd.CapitalOutSmall)

		netMainFlow := cd.CapitalInSuper + cd.CapitalInBig -
			cd.CapitalOutSuper - cd.CapitalOutBig
		fmt.Printf("  Net Main Flow:  %.2f\n", netMainFlow)
		fmt.Printf("  Update Time: %s\n", cd.UpdateTime)
	}
	fmt.Println()

	// 5. Stock Filter/Screening
	fmt.Println("=== 5. Stock Filter (StockFilter) ===")
	filterReq := &qot.StockFilterRequest{
		Begin:  0,
		Num:    10,
		Market: hkMarket,
	}

	// Example: Filter stocks with price between 100 and 500
	// Note: Simulator returns empty results, but this shows the API usage
	/*
		filterReq.BaseFilterList = []*qotstockfilter.BaseFilter{
			{
				FieldName:  int32(qotstockfilter.StockField_StockField_CurPrice),
				FilterMin:  ptrFloat64(100.0),
				FilterMax:  ptrFloat64(500.0),
				IsNoFilter: ptrBool(false),
			},
		}
	*/

	filterResp, err := qot.StockFilter(cli, filterReq)
	if err != nil {
		log.Printf("StockFilter failed: %v", err)
	} else {
		fmt.Printf("  Found %d stocks matching criteria (showing first %d)\n",
			filterResp.AllCount, len(filterResp.DataList))
		for _, stock := range filterResp.DataList {
			fmt.Printf("  %s (%s)", stock.Security.GetCode(), stock.Name)
			// Extract price from BaseDataList
			for _, bd := range stock.BaseDataList {
				if bd.GetFieldName() == int32(qotstockfilter.StockField_StockField_CurPrice) {
					fmt.Printf(" | Price=%.2f", bd.GetValue())
					break
				}
			}
			fmt.Println()
		}
	}
	fmt.Println()

	// 6. Get Option Expiration Dates
	fmt.Println("=== 6. Option Expiration Dates (GetOptionExpirationDate) ===")
	optionExpReq := &qot.GetOptionExpirationDateRequest{
		Owner: security,
	}

	optionExpResp, err := qot.GetOptionExpirationDate(cli, optionExpReq)
	if err != nil {
		log.Printf("GetOptionExpirationDate failed: %v", err)
	} else {
		fmt.Printf("  Stock: %s\n", security.GetCode())
		fmt.Printf("  Found %d expiration dates\n", len(optionExpResp.DateList))
		for i, date := range optionExpResp.DateList {
			if i >= 5 {
				fmt.Printf("  ... and %d more\n", len(optionExpResp.DateList)-5)
				break
			}
			fmt.Printf("  %s\n", date.StrikeTime)
		}
	}
	fmt.Println()

	// 7. Get Option Chain - NOTE: Currently not implemented due to protobuf issues
	fmt.Println("=== 7. Option Chain (GetOptionChain) ===")
	fmt.Println("  ⚠️  GetOptionChain is currently not implemented in the SDK")
	fmt.Println("  Reason: Protobuf structure compatibility issues")
	fmt.Println("  Workaround: Use GetOptionExpirationDate + manual tracking")
	fmt.Println()

	// 8. Get Warrant Information
	fmt.Println("=== 8. Warrant Information (GetWarrant) ===")
	warrantReq := &qot.GetWarrantRequest{
		Begin: 0,
		Num:   10,
		Owner: security,
	}

	warrantResp, err := qot.GetWarrant(cli, warrantReq)
	if err != nil {
		log.Printf("GetWarrant failed: %v", err)
	} else {
		fmt.Printf("  Found %d warrants (showing first %d)\n",
			warrantResp.AllCount, len(warrantResp.WarrantDataList))
		for _, w := range warrantResp.WarrantDataList {
			fmt.Printf("  %s | %s | Price=%.3f | Vol=%d\n",
				w.Stock.GetCode(), w.Name,
				w.CurPrice, w.Volume)
		}
	}
	fmt.Println()

	// 9. Get Trading Dates
	fmt.Println("=== 9. Trading Calendar (GetTradeDate) ===")
	tradeDateReq := &qot.GetTradeDateRequest{
		Market:    hkMarket,
		BeginTime: "2026-04-01",
		EndTime:   "2026-04-30",
	}

	tradeDateResp, err := qot.GetTradeDate(cli, tradeDateReq)
	if err != nil {
		log.Printf("GetTradeDate failed: %v", err)
	} else {
		fmt.Printf("  Market: HK | April 2026\n")
		fmt.Printf("  Trading days: ")
		for i, td := range tradeDateResp.TradeDateList {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(td.GetTime())
			if i >= 9 {
				fmt.Printf(" ... (%d total)", len(tradeDateResp.TradeDateList))
				break
			}
		}
		fmt.Println()
	}
	fmt.Println()

	// 10. Get Futures Information
	fmt.Println("=== 10. Futures Information (GetFutureInfo) ===")
	futureReq := &qot.GetFutureInfoRequest{
		SecurityList: []*qotcommon.Security{security},
	}

	futureResp, err := qot.GetFutureInfo(cli, futureReq)
	if err != nil {
		log.Printf("GetFutureInfo failed: %v", err)
	} else {
		fmt.Printf("  Found %d futures contracts\n", len(futureResp.FutureInfoList))
		for _, f := range futureResp.FutureInfoList {
			fmt.Printf("  %s | %s | LastTradeDay=%s\n",
				f.Security.GetCode(), f.Name, f.LastTradeTime)
		}
	}
	fmt.Println()

	// 11. Get IPO List
	fmt.Println("=== 11. IPO List (GetIpoList) ===")
	ipoReq := &qot.GetIpoListRequest{
		Market: hkMarket,
	}

	ipoResp, err := qot.GetIpoList(cli, ipoReq)
	if err != nil {
		log.Printf("GetIpoList failed: %v", err)
	} else {
		fmt.Printf("  Found %d IPOs\n", len(ipoResp.IpoList))
		for _, ipo := range ipoResp.IpoList {
			fmt.Printf("  %s (%s) | ListDate=%s",
				ipo.Basic.Security.GetCode(), ipo.Basic.Name, ipo.Basic.ListTime)

			// Show exchange-specific data
			if ipo.HkExData != nil {
				fmt.Printf(" | IPOPriceMin=%.2f", ipo.HkExData.IpoPriceMin)
			}
			fmt.Println()
		}
	}
	fmt.Println()

	fmt.Println("=== Advanced Examples Complete ===")
	fmt.Println("Note: Some APIs may return mock data when using the simulator")
}

// Helper functions
func ptrStr(s string) *string {
	return &s
}

func ptrInt32(v int32) *int32 {
	return &v
}

func ptrInt64(v int64) *int64 {
	return &v
}

func ptrFloat64(v float64) *float64 {
	return &v
}

func ptrBool(v bool) *bool {
	return &v
}

// Ensure imports are used
var _ = qotstockfilter.StockField_StockField_CurPrice
