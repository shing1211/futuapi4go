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

// Example: UnlockTrade - 解鎖交易
//
// This example demonstrates how to use the UnlockTrade API to unlock
// trading for order placement and modification.
package main

import (
	"fmt"
	"log"
	"os"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/trd"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== UnlockTrade Example / 解鎖交易示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Get password from command line argument / 從命令行參數獲取密碼
	password := ""
	if len(os.Args) > 1 {
		password = os.Args[1]
	} else {
		fmt.Println("⚠️  No password provided, using default '123456' for testing")
		password = "123456"
	}

	// Unlock trade / 解鎖交易
	unlock := true
	req := &trd.UnlockTradeRequest{
		Unlock: unlock,
		PwdMD5: password,
	}

	fmt.Println("🔓 Attempting to unlock trading / 正在嘗試解鎖交易...")
	err := trd.UnlockTrade(cli, req)
	if err != nil {
		fmt.Printf("❌ UnlockTrade failed: %v\n", err)
		fmt.Println("\nNote: This is expected if:")
		fmt.Println("  - Password is incorrect")
		fmt.Println("  - Trade password not set in OpenD")
		fmt.Println("  - Already unlocked")
	} else {
		fmt.Println("✅ Trade unlocked successfully / 交易解鎖成功")
	}

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}
