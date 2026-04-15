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

// Example: GetBroker - 獲取經紀隊列
//
// This example demonstrates how to use the GetBroker API to retrieve
// broker queue data showing which brokers are buying/selling.
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

	fmt.Println("=== GetBroker Example / 獲取經紀隊列示例 ===")
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

	// Get broker queue / 獲取經紀隊列
	req := &qot.GetBrokerRequest{
		Security: security,
		Num:      10, // Top 10 brokers / 前10家經紀
	}

	resp, err := qot.GetBroker(cli, req)
	if err != nil {
		log.Printf("GetBroker failed: %v", err)
		return
	}

	fmt.Printf("📊 Broker Queue / 經紀隊列 for %s (%s)\n\n", security.GetCode(), resp.Name)

	// Display ask brokers (sell side) / 顯示賣方經紀
	fmt.Println("🔴 Ask Brokers (Sell Side) / 賣方經紀:")
	fmt.Printf("  %-6s %-20s %-12s\n", "Pos", "Broker Name", "Volume")
	fmt.Println("  " + "--------------------------------------")

	for _, b := range resp.AskBrokerList {
		fmt.Printf("  %-6d %-20s %-12d\n", b.Pos, b.Name, b.Volume)
	}

	fmt.Println()

	// Display bid brokers (buy side) / 顯示買方經紀
	fmt.Println("🟢 Bid Brokers (Buy Side) / 買方經紀:")
	fmt.Printf("  %-6s %-20s %-12s\n", "Pos", "Broker Name", "Volume")
	fmt.Println("  " + "--------------------------------------")

	for _, b := range resp.BidBrokerList {
		fmt.Printf("  %-6d %-20s %-12d\n", b.Pos, b.Name, b.Volume)
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

func ptrStr(s string) *string { return &s }
