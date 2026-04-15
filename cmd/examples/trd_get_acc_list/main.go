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

// Example: GetAccList - 獲取交易賬戶列表
//
// This example demonstrates how to use the GetAccList API to retrieve
// the list of trading accounts.
package main

import (
	"fmt"
	"log"
	"os"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== GetAccList Example / 獲取交易賬戶列表示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Get account list for securities / 獲取證券賬戶列表
	trdCategory := int32(trdcommon.TrdCategory_TrdCategory_Security)

	resp, err := trd.GetAccList(cli, trdCategory, false)
	if err != nil {
		log.Printf("GetAccList failed: %v", err)
		return
	}

	fmt.Printf("📊 Found %d trading accounts\n\n", len(resp.AccList))

	for i, acc := range resp.AccList {
		trdEnv := "Simulation"
		if acc.TrdEnv == int32(trdcommon.TrdEnv_TrdEnv_Real) {
			trdEnv = "Real"
		}

		accType := "Unknown"
		switch acc.AccType {
		case int32(trdcommon.TrdAccType_TrdAccType_Cash):
			accType = "Cash"
		case int32(trdcommon.TrdAccType_TrdAccType_Margin):
			accType = "Margin"
		}

		fmt.Printf("Account %d / 賬戶 %d:\n", i+1, i+1)
		fmt.Printf("  AccID / 賬戶ID:     %d\n", acc.AccID)
		fmt.Printf("  TrdEnv / 交易環境:   %s\n", trdEnv)
		fmt.Printf("  AccType / 賬戶類型:  %s\n", accType)
		fmt.Printf("  AccStatus / 狀態:    %d\n", acc.AccStatus)
		fmt.Println()
	}

	fmt.Println("=== Example Complete / 示例完成 ===")
}
