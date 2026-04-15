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

// Example: GetRT - 獲取實時分時數據
//
// This example demonstrates how to use the GetRT API to retrieve
// real-time intraday minute-by-minute data for a security.
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

	fmt.Println("=== GetRT Example / 獲取實時分時示例 ===")
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

	// Get real-time minute data / 獲取實時分時數據
	req := &qot.GetRTRequest{
		Security: security,
	}

	resp, err := qot.GetRT(cli, req)
	if err != nil {
		log.Printf("GetRT failed: %v", err)
		return
	}

	fmt.Printf("📊 Real-time Minute Data / 實時分時數據 for %s (%s)\n",
		security.GetCode(), resp.Name)
	fmt.Printf("Retrieved %d minute records\n\n", len(resp.RTList))

	// Show first 30 records (approximately 30 minutes of data) / 顯示前30筆
	displayCount := 30
	if len(resp.RTList) < displayCount {
		displayCount = len(resp.RTList)
	}

	fmt.Printf("  %-20s %-10s %-12s %-12s %-12s\n",
		"Time", "Price", "Volume", "Turnover", "AvgPrice")
	fmt.Println("  " + "------------------------------------------------------------")

	for i := 0; i < displayCount; i++ {
		rt := resp.RTList[i]
		fmt.Printf("  %-20s %-10.2f %-12d %-12.2f %-12.2f\n",
			rt.Time, rt.Price, rt.Volume, rt.Turnover, rt.AvgPrice)
	}

	if len(resp.RTList) > displayCount {
		fmt.Printf("  ... and %d more records\n", len(resp.RTList)-displayCount)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

func ptrStr(s string) *string { return &s }
