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

// Example: GetTicker - 獲取逐筆成交數據
//
// This example demonstrates how to use the GetTicker API to retrieve
// tick-by-tick trade data for a security.
package main

import (
	"fmt"
	"log"
	"os"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== GetTicker Example / 獲取逐筆成交示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	hkFutureMarket := int32(qotcommon.QotMarket_QotMarket_HK_Future)
	security := &qotcommon.Security{
		Market: &hkFutureMarket,
		Code:   ptrStr("HSImain"), // HSI Futures Main Contract / 恆生指數期貨
	}

	// Get last 20 tick records / 獲取最近20筆逐筆成交
	req := &qot.GetTickerRequest{
		Security: security,
		Num:      20,
	}

	resp, err := qot.GetTicker(cli, req)
	if err != nil {
		log.Printf("GetTicker failed: %v", err)
		return
	}

	fmt.Printf("📊 Ticker Data for %s (%s)\n", security.GetCode(), resp.Name)
	fmt.Printf("Retrieved %d tick records\n\n", len(resp.TickerList))

	fmt.Printf("  %-20s %-10s %-10s %-12s %-8s\n",
		"Time", "Price", "Volume", "Turnover", "Dir")
	fmt.Println("  " + "------------------------------------------------------------")

	for _, t := range resp.TickerList {
		dir := "N" // Neutral / 中性
		if t.Dir == 1 {
			dir = "B" // Buy / 買
		} else if t.Dir == 2 {
			dir = "S" // Sell / 賣
		}

		fmt.Printf("  %-20s %-10.2f %-10d %-12.2f %-8s\n",
			t.Time, t.Price, t.Volume, t.Turnover, dir)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

func ptrStr(s string) *string { return &s }
