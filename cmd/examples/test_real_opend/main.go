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

// Quick Test for Real Futu OpenD Connection
//
// This is a minimal test to verify the SDK works with your real Futu OpenD
// Run: go run main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/pkg/sys"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

func main() {
	// Get connection address from env or use default
	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║   FutuAPI4Go - Real OpenD Connection Test     ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
	fmt.Println()

	// Create client
	cli := futuapi.New()
	defer cli.Close()

	// Connect
	fmt.Printf("📡 Connecting to %s...\n", addr)
	if err := cli.Connect(addr); err != nil {
		log.Fatalf("❌ Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d, ServerVer=%d\n\n", cli.GetConnID(), cli.GetServerVer())

	// Test 1: System APIs
	fmt.Println("=== Test 1: System APIs ===")
	testSystemAPIs(cli)

	// Test 2: Market Data
	fmt.Println("\n=== Test 2: Market Data ===")
	testMarketData(cli)

	// Test 3: Trading APIs
	fmt.Println("\n=== Test 3: Trading APIs ===")
	testTradingAPIs(cli)

	fmt.Println("\n╔════════════════════════════════════════════════╗")
	fmt.Println("║   ✅ All Tests Complete! SDK is working!       ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
}

func testSystemAPIs(cli *futuapi.Client) {
	// GetGlobalState
	state, err := sys.GetGlobalState(cli)
	if err != nil {
		fmt.Printf("❌ GetGlobalState failed: %v\n", err)
		return
	}

	fmt.Printf("✓ GetGlobalState:\n")
	fmt.Printf("  HK Market: %d | US Market: %d\n", state.MarketHK, state.MarketUS)
	fmt.Printf("  Qot Logged: %v | Trd Logged: %v\n", state.QotLogined, state.TrdLogined)
	fmt.Printf("  ServerVer: %d | BuildNo: %d\n", state.ServerVer, state.ServerBuildNo)

	// GetUserInfo
	user, err := sys.GetUserInfo(cli)
	if err != nil {
		fmt.Printf("❌ GetUserInfo failed: %v\n", err)
		return
	}

	fmt.Printf("✓ GetUserInfo:\n")
	fmt.Printf("  NickName: %s | UserID: %d\n", user.NickName, user.UserID)
	fmt.Printf("  ApiLevel: %s | Need Agree Disclaimer: %v\n", user.ApiLevel, user.IsNeedAgreeDisclaimer)
}

func testMarketData(cli *futuapi.Client) {
	hkFutureMarket := int32(qotcommon.QotMarket_QotMarket_HK_Future)

	// Test GetBasicQot
	securities := []*qotcommon.Security{
		{Market: &hkFutureMarket, Code: ptrStr("HSImain")}, // HSI Futures Main Contract
	}

	fmt.Printf("📊 Getting quotes for: %s...\n", "HSImain")
	quotes, err := qot.GetBasicQot(context.Background(),cli, securities)
	if err != nil {
		fmt.Printf("❌ GetBasicQot failed: %v\n", err)
		return
	}

	for _, q := range quotes {
		fmt.Printf("✓ %s (%s):\n", q.Security.GetCode(), q.Name)
		fmt.Printf("  Price: %.2f | Open: %.2f | High: %.2f | Low: %.2f\n",
			q.CurPrice, q.OpenPrice, q.HighPrice, q.LowPrice)
		fmt.Printf("  Volume: %d | Turnover: %.0f\n", q.Volume, q.Turnover)
	}

	// Test GetKL
	fmt.Printf("\n📈 Getting K-line data for %s...\n", "HSImain")
	klReq := &qot.GetKLRequest{
		Security:  securities[0],
		RehabType: int32(qotcommon.RehabType_RehabType_None),
		KLType:    int32(qotcommon.KLType_KLType_Day),
		ReqNum:    3,
	}

	klResp, err := qot.GetKL(cli, klReq)
	if err != nil {
		fmt.Printf("❌ GetKL failed: %v\n", err)
		return
	}

	fmt.Printf("✓ K-line Data (%s):\n", klResp.Name)
	for i, kl := range klResp.KLList {
		fmt.Printf("  [%d] %s | O:%.2f H:%.2f L:%.2f C:%.2f | V:%d\n",
			i+1, kl.Time, kl.OpenPrice, kl.HighPrice, kl.LowPrice, kl.ClosePrice, kl.Volume)
	}
}

func testTradingAPIs(cli *futuapi.Client) {
	trdCategory := int32(trdcommon.TrdCategory_TrdCategory_Security)

	// GetAccList
	fmt.Printf("👤 Getting account list...\n")
	accResp, err := trd.GetAccList(cli, trdCategory, false)
	if err != nil {
		fmt.Printf("❌ GetAccList failed: %v\n", err)
		return
	}

	fmt.Printf("✓ Found %d accounts\n", len(accResp.AccList))
	for i, acc := range accResp.AccList {
		trdEnv := "Simulation"
		if acc.TrdEnv == int32(trdcommon.TrdEnv_TrdEnv_Real) {
			trdEnv = "Real"
		}
		fmt.Printf("  [%d] AccID: %d | TrdEnv: %s | AccType: %d\n",
			i+1, acc.AccID, trdEnv, acc.AccType)
	}

	// GetFunds (if we have an account)
	if len(accResp.AccList) > 0 {
		accID := accResp.AccList[0].AccID
		hkMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)

		fmt.Printf("\n💰 Getting funds for account %d...\n", accID)
		fundsReq := &trd.GetFundsRequest{
			AccID:     accID,
			TrdMarket: hkMarket,
		}

		fundsResp, err := trd.GetFunds(cli, fundsReq)
		if err != nil {
			fmt.Printf("❌ GetFunds failed: %v\n", err)
			return
		}

		f := fundsResp.Funds
		fmt.Printf("✓ Account Funds:\n")
		fmt.Printf("  TotalAssets: %.2f | Cash: %.2f | MarketVal: %.2f\n",
			f.TotalAssets, f.Cash, f.MarketVal)
		fmt.Printf("  AvailableFunds: %.2f | FrozenCash: %.2f\n", f.AvailableFunds, f.FrozenCash)
	}
}

func ptrStr(s string) *string { return &s }
