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

// Example: RequestTradeDate - 獲取交易日期
//
// This example demonstrates how to use the RequestTradeDate API to retrieve
// trading calendar dates for a specific market and date range.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	fmt.Println("=== RequestTradeDate Example / 獲取交易日期示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)

	// Calculate date range (current month) / 計算日期範圍（當月）
	now := time.Now()
	beginTime := now.Format("2006-01-01")
	endTime := now.AddDate(0, 1, 0).Format("2006-01-02")

	// Get trade dates / 獲取交易日期
	req := &qot.RequestTradeDateRequest{
		Market:    hkMarket,
		BeginTime: beginTime,
		EndTime:   endTime,
	}

	resp, err := qot.RequestTradeDate(cli, req)
	if err != nil {
		log.Printf("RequestTradeDate failed: %v", err)
		return
	}

	fmt.Printf("📅 Trading Calendar / 交易日曆\n")
	fmt.Printf("Market / 市場: HK / 港股\n")
	fmt.Printf("Period / 期間:  %s to %s\n\n", beginTime, endTime)

	fmt.Printf("Found %d trading days / 找到%d個交易日\n\n", len(resp.TradeDateList), len(resp.TradeDateList))

	// Display trading dates / 顯示交易日期
	for i, td := range resp.TradeDateList {
		// Highlight today / 突出顯示今天
		marker := "  "
		if td.GetTime() == now.Format("2006-01-02") {
			marker = "🔵"
		}

		fmt.Printf("%s %d. %s\n", marker, i+1, td.GetTime())
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}
