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

package integration_test

import (
	"context"
	"os"
	"testing"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/pkg/sys"
	"github.com/shing1211/futuapi4go/pkg/trd"
	"github.com/shing1211/futuapi4go/test/fixtures"
)

// ============================================================================
// HSI-Focused Integration Tests
// These tests require a running Futu OpenD instance
// Run with: FUTU_INTEGRATION_TESTS=1 go test -v ./test/integration
// ============================================================================

func skipIfNotIntegration(t *testing.T) {
	if os.Getenv("FUTU_INTEGRATION_TESTS") == "" {
		t.Skip("Skipping integration test: set FUTU_INTEGRATION_TESTS=1 to run")
	}
}

func getOpenDAddr() string {
	addr := os.Getenv("FUTU_OPEND_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}
	return addr
}

func TestIntegration_HSI_ConnectAndGlobalState(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
		futuapi.WithLogLevel(1),
	)
	defer cli.Close()

	// Connect
	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	// Get global state
	state, err := sys.GetGlobalState(cli)
	if err != nil {
		t.Fatalf("GetGlobalState failed: %v", err)
	}

	if state.ServerVer == 0 {
		t.Error("ServerVer should not be zero")
	}

	if !state.QotLogined {
		t.Error("Should be logged into quote server")
	}

	t.Logf("Connected to OpenD: ServerVer=%d, BuildNo=%d", state.ServerVer, state.ServerBuildNo)
}

func TestIntegration_HSI_BasicQot(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	// Get HSI basic quote
	hsiSecurity := fixtures.HSISecurity()
	quotes, err := qot.GetBasicQot(cli, []*qotcommon.Security{hsiSecurity})
	if err != nil {
		t.Fatalf("GetBasicQot for HSI failed: %v", err)
	}

	if len(quotes) != 1 {
		t.Fatalf("Expected 1 quote, got %d", len(quotes))
	}

	quote := quotes[0]
	t.Logf("HSI Quote: Price=%.2f, Open=%.2f, High=%.2f, Low=%.2f, Volume=%d",
		quote.CurPrice, quote.OpenPrice, quote.HighPrice, quote.LowPrice, quote.Volume)

	if quote.CurPrice <= 0 {
		t.Error("HSI price should be positive")
	}

	if quote.Volume == 0 {
		t.Error("HSI volume should not be zero")
	}
}

func TestIntegration_HSI_KLine(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	// Get HSI daily K-line
	req := &qot.GetKLRequest{
		Security:  fixtures.HSISecurity(),
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    10,
	}

	result, err := qot.GetKL(cli, req)
	if err != nil {
		t.Fatalf("GetKL failed: %v", err)
	}

	if len(result.KLList) == 0 {
		t.Fatal("Expected at least 1 K-line")
	}

	t.Logf("HSI K-Lines received: %d", len(result.KLList))
	for i, kl := range result.KLList {
		if i < 3 {
			t.Logf("  [%d] %s: O=%.2f H=%.2f L=%.2f C=%.2f V=%d",
				i, kl.Time, kl.OpenPrice, kl.HighPrice, kl.LowPrice, kl.ClosePrice, kl.Volume)
		}
	}
}

func TestIntegration_HSI_OrderBook(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	req := &qot.GetOrderBookRequest{
		Security: fixtures.HSISecurity(),
		Num:      10,
	}

	result, err := qot.GetOrderBook(cli, req)
	if err != nil {
		t.Fatalf("GetOrderBook failed: %v", err)
	}

	if len(result.OrderBookAskList) == 0 {
		t.Error("Expected at least 1 ask level")
	}

	if len(result.OrderBookBidList) == 0 {
		t.Error("Expected at least 1 bid level")
	}

	t.Logf("HSI OrderBook: %d asks, %d bids", len(result.OrderBookAskList), len(result.OrderBookBidList))

	// Validate no crossed book
	if len(result.OrderBookAskList) > 0 && len(result.OrderBookBidList) > 0 {
		bestAsk := result.OrderBookAskList[0].Price
		bestBid := result.OrderBookBidList[0].Price
		if bestAsk <= bestBid {
			t.Errorf("Crossed book detected: ask %.2f <= bid %.2f", bestAsk, bestBid)
		}
	}
}

func TestIntegration_HSI_Ticker(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	req := &qot.GetTickerRequest{
		Security: fixtures.HSISecurity(),
		Num:      20,
	}

	result, err := qot.GetTicker(cli, req)
	if err != nil {
		t.Fatalf("GetTicker failed: %v", err)
	}

	t.Logf("HSI Tickers received: %d", len(result.TickerList))
	if len(result.TickerList) > 0 {
		ticker := result.TickerList[0]
		t.Logf("  Latest: Time=%s, Price=%.2f, Volume=%d", ticker.Time, ticker.Price, ticker.Volume)
	}
}

func TestIntegration_HSI_RT(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	req := &qot.GetRTRequest{
		Security: fixtures.HSISecurity(),
	}

	result, err := qot.GetRT(cli, req)
	if err != nil {
		t.Fatalf("GetRT failed: %v", err)
	}

	t.Logf("HSI RT data points: %d", len(result.RTList))
	if len(result.RTList) > 0 {
		t.Logf("  First: Time=%s, Price=%.2f", result.RTList[0].Time, result.RTList[0].Price)
		t.Logf("  Last:  Time=%s, Price=%.2f", result.RTList[len(result.RTList)-1].Time, result.RTList[len(result.RTList)-1].Price)
	}
}

func TestIntegration_HSI_Subscribe_Push(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	// Set up push handler
	pushReceived := make(chan bool, 10)
	cli.SetPushHandler(func(pkt *futuapi.Packet) {
		t.Logf("Push received: ProtoID=%d", pkt.Header.ProtoID)
		pushReceived <- true
	})

	// Subscribe to HSI basic quote
	req := &qot.SubscribeRequest{
		SecurityList:     []*qotcommon.Security{fixtures.HSISecurity()},
		SubTypeList:      []qot.SubType{qot.SubType_Basic},
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
	}

	_, err := qot.Subscribe(cli, req)
	if err != nil {
		t.Fatalf("Subscribe failed: %v", err)
	}

	// Wait for push (max 10 seconds)
	select {
	case <-pushReceived:
		t.Log("Successfully received HSI push notification")
	case <-time.After(10 * time.Second):
		t.Log("No push received within timeout (market may be closed)")
	}

	// Unsubscribe
	req.IsSubOrUnSub = false
	_, err = qot.Subscribe(cli, req)
	if err != nil {
		t.Fatalf("Unsubscribe failed: %v", err)
	}
}

func TestIntegration_HSI_StationaryInfo(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	req := &qot.GetStaticInfoRequest{
		SecurityList: []*qotcommon.Security{fixtures.HSISecurity()},
	}

	result, err := qot.GetStaticInfo(cli, req)
	if err != nil {
		t.Fatalf("GetStaticInfo failed: %v", err)
	}

	if len(result.StaticInfoList) != 1 {
		t.Errorf("Expected 1 static info, got %d", len(result.StaticInfoList))
	}

	info := result.StaticInfoList[0]
	basic := info.GetBasic()
	if basic != nil {
		t.Logf("HSI Static Info: Name=%s, SecType=%d, LotSize=%d",
			basic.GetSecurity().GetCode(), basic.GetSecType(), basic.GetLotSize())
	}
}

func TestIntegration_HSI_TradeDate(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	req := &qot.GetTradeDateRequest{
		Market: fixtures.HSIMarket,
	}

	result, err := qot.GetTradeDate(cli, req)
	if err != nil {
		t.Fatalf("GetTradeDate failed: %v", err)
	}

	t.Logf("HK Trade Dates: %d dates received", len(result.TradeDateList))
	if len(result.TradeDateList) > 0 {
		t.Logf("  First: %s", result.TradeDateList[0])
		t.Logf("  Last:  %s", result.TradeDateList[len(result.TradeDateList)-1])
	}
}

func TestIntegration_HSI_CapitalFlow(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	req := &qot.GetCapitalFlowRequest{
		Security: fixtures.HSISecurity(),
	}

	result, err := qot.GetCapitalFlow(cli, req)
	if err != nil {
		t.Fatalf("GetCapitalFlow failed: %v", err)
	}

	t.Logf("HSI Capital Flow: %d items", len(result.FlowItemList))
	if len(result.FlowItemList) > 0 {
		item := result.FlowItemList[0]
		t.Logf("  Latest: Time=%s, MainIn=%.2f, BigIn=%.2f",
			item.Time, item.MainInFlow, item.BigInFlow)
	}
}

func TestIntegration_Trading_Workflow(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	t.Log("=== Step 1: Get Account List ===")
	accList, err := trd.GetAccList(cli, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
	if err != nil {
		t.Fatalf("GetAccList failed: %v", err)
	}

	if len(accList.AccList) == 0 {
		t.Fatal("No trading accounts available")
	}

	acc := accList.AccList[0]
	t.Logf("Account: ID=%d, Env=%d, Type=%d", acc.AccID, acc.TrdEnv, acc.AccType)

	t.Log("=== Step 2: Get Funds ===")
	funds, err := trd.GetFunds(cli, &trd.GetFundsRequest{
		AccID:     acc.AccID,
		TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
	})
	if err != nil {
		t.Fatalf("GetFunds failed: %v", err)
	}

	t.Logf("Funds: Total=%.2f, Cash=%.2f, Power=%.2f",
		funds.Funds.TotalAssets, funds.Funds.Cash, funds.Funds.Power)

	t.Log("=== Step 3: Get Positions ===")
	positions, err := trd.GetPositionList(cli, &trd.GetPositionListRequest{
		AccID:     acc.AccID,
		TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
	})
	if err != nil {
		t.Fatalf("GetPositionList failed: %v", err)
	}

	t.Logf("Positions: %d positions", len(positions.PositionList))
	for _, pos := range positions.PositionList {
		t.Logf("  %s: Qty=%.0f, Cost=%.2f, PL=%.2f", pos.Code, pos.Qty, pos.CostPrice, pos.PlVal)
	}

	t.Log("=== Step 4: Get Orders ===")
	orders, err := trd.GetOrderList(cli, &trd.GetOrderListRequest{
		AccID:     acc.AccID,
		TrdMarket: int32(trdcommon.TrdMarket_TrdMarket_HK),
	})
	if err != nil {
		t.Fatalf("GetOrderList failed: %v", err)
	}

	t.Logf("Orders: %d orders", len(orders.OrderList))
	for _, order := range orders.OrderList {
		t.Logf("  Order %d: %s %.2f x %.0f, Status=%d",
			order.OrderID, order.Code, order.Price, order.Qty, order.OrderStatus)
	}
}

func TestIntegration_ContextCancellation(t *testing.T) {
	skipIfNotIntegration(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	cli = cli.WithContext(ctx)

	// Cancel context during operation
	cancel()

	// Next API call should fail or timeout
	_, err := qot.GetBasicQot(cli, []*qotcommon.Security{fixtures.HSISecurity()})
	if err == nil {
		t.Log("API succeeded before context cancellation propagated")
	}
}

func TestIntegration_HSI_ComprehensiveMarketData(t *testing.T) {
	skipIfNotIntegration(t)

	cli := futuapi.New(
		futuapi.WithDialTimeout(10*time.Second),
		futuapi.WithAPITimeout(10*time.Second),
		futuapi.WithMaxRetries(2),
	)
	defer cli.Close()

	if err := cli.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	// Test all major Qot APIs with HSI
	t.Run("GetBasicQot", func(t *testing.T) {
		_, err := qot.GetBasicQot(cli, []*qotcommon.Security{fixtures.HSISecurity()})
		if err != nil {
			t.Errorf("GetBasicQot failed: %v", err)
		}
	})

	t.Run("GetKL", func(t *testing.T) {
		req := &qot.GetKLRequest{
			Security:  fixtures.HSISecurity(),
			RehabType: int32(qotcommon.RehabType_RehabType_None),
			KLType:    int32(qotcommon.KLType_KLType_Day),
			ReqNum:    5,
		}
		_, err := qot.GetKL(cli, req)
		if err != nil {
			t.Errorf("GetKL failed: %v", err)
		}
	})

	t.Run("GetOrderBook", func(t *testing.T) {
		req := &qot.GetOrderBookRequest{
			Security: fixtures.HSISecurity(),
			Num:      5,
		}
		_, err := qot.GetOrderBook(cli, req)
		if err != nil {
			t.Errorf("GetOrderBook failed: %v", err)
		}
	})

	t.Run("GetTicker", func(t *testing.T) {
		req := &qot.GetTickerRequest{
			Security: fixtures.HSISecurity(),
			Num:      10,
		}
		_, err := qot.GetTicker(cli, req)
		if err != nil {
			t.Errorf("GetTicker failed: %v", err)
		}
	})

	t.Run("GetRT", func(t *testing.T) {
		req := &qot.GetRTRequest{
			Security: fixtures.HSISecurity(),
		}
		_, err := qot.GetRT(cli, req)
		if err != nil {
			t.Errorf("GetRT failed: %v", err)
		}
	})

	t.Run("GetBroker", func(t *testing.T) {
		req := &qot.GetBrokerRequest{
			Security: fixtures.HSISecurity(),
			Num:      5,
		}
		_, err := qot.GetBroker(cli, req)
		if err != nil {
			t.Errorf("GetBroker failed: %v", err)
		}
	})

	t.Run("GetStaticInfo", func(t *testing.T) {
		req := &qot.GetStaticInfoRequest{
			SecurityList: []*qotcommon.Security{fixtures.HSISecurity()},
		}
		_, err := qot.GetStaticInfo(cli, req)
		if err != nil {
			t.Errorf("GetStaticInfo failed: %v", err)
		}
	})

	t.Run("GetTradeDate", func(t *testing.T) {
		req := &qot.GetTradeDateRequest{
			Market: fixtures.HSIMarket,
		}
		_, err := qot.GetTradeDate(cli, req)
		if err != nil {
			t.Errorf("GetTradeDate failed: %v", err)
		}
	})

	t.Run("GetCapitalFlow", func(t *testing.T) {
		req := &qot.GetCapitalFlowRequest{
			Security: fixtures.HSISecurity(),
		}
		_, err := qot.GetCapitalFlow(cli, req)
		if err != nil {
			t.Errorf("GetCapitalFlow failed: %v", err)
		}
	})

	t.Run("GetCapitalDistribution", func(t *testing.T) {
		_, err := qot.GetCapitalDistribution(cli, fixtures.HSISecurity())
		if err != nil {
			t.Errorf("GetCapitalDistribution failed: %v", err)
		}
	})
}
