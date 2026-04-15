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

// Example: GetCapitalFlow - 獲取資金流向
//
// This example demonstrates how to use the GetCapitalFlow API to retrieve
// capital flow data showing money movement in/out of a security.
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

	fmt.Println("=== GetCapitalFlow Example / 獲取資金流向示例 ===")
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

	// Get capital flow (daily) / 獲取資金流向（日線）
	req := &qot.GetCapitalFlowRequest{
		Security:   security,
		PeriodType: 1, // Daily / 日線
	}

	resp, err := qot.GetCapitalFlow(cli, req)
	if err != nil {
		log.Printf("GetCapitalFlow failed: %v", err)
		return
	}

	fmt.Printf("📊 Capital Flow / 資金流向 for %s\n\n", security.GetCode())

	// Display capital flow data / 顯示資金流向數據
	displayCount := 10
	if len(resp.FlowItemList) < displayCount {
		displayCount = len(resp.FlowItemList)
	}

	fmt.Printf("  %-20s %-12s %-12s %-12s\n",
		"Time", "In Flow", "Net Flow", "Main In")
	fmt.Println("  " + "-------------------------------------------------------")

	for i := 0; i < displayCount; i++ {
		flow := resp.FlowItemList[i]
		fmt.Printf("  %-20s %-12.0f %-12.0f %-12.0f\n",
			flow.Time, flow.InFlow, flow.InFlow, flow.MainInFlow)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

func ptrStr(s string) *string { return &s }
