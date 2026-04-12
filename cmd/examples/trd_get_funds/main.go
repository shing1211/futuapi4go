// Example: GetFunds - 獲取資金信息
//
// This example demonstrates how to use the GetFunds API to retrieve
// account funds information including cash, market value, etc.
//
// Usage:
//   go run main.go

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

	fmt.Println("=== GetFunds Example / 獲取資金信息示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Get account list first / 先獲取賬戶列表
	accResp, err := trd.GetAccList(cli, int32(trdcommon.TrdCategory_TrdCategory_Security), false)
	if err != nil {
		log.Printf("GetAccList failed: %v", err)
		return
	}

	if len(accResp.AccList) == 0 {
		fmt.Println("No trading accounts found")
		return
	}

	acc := accResp.AccList[0]
	hkMarket := int32(trdcommon.TrdMarket_TrdMarket_HK)

	// Get funds / 獲取資金
	req := &trd.GetFundsRequest{
		AccID:     acc.AccID,
		TrdMarket: hkMarket,
	}

	resp, err := trd.GetFunds(cli, req)
	if err != nil {
		log.Printf("GetFunds failed: %v", err)
		return
	}

	funds := resp.Funds

	fmt.Printf("📊 Account Funds / 賬戶資金 (AccID: %d)\n\n", acc.AccID)
	fmt.Printf("  💰 Total Assets / 總資產:        %.2f\n", funds.TotalAssets)
	fmt.Printf("  💵 Cash / 現金:                  %.2f\n", funds.Cash)
	fmt.Printf("  📈 Market Value / 市值:           %.2f\n", funds.MarketVal)
	fmt.Printf("  ✅ Available / 可用資金:          %.2f\n", funds.AvailableFunds)
	fmt.Printf("  ❄️  Frozen Cash / 凍結資金:       %.2f\n", funds.FrozenCash)
	fmt.Printf("  💳 Debt Cash / 負債現金:          %.2f\n", funds.DebtCash)
	fmt.Printf("  🚀 Power / 購買力:                %.2f\n", funds.Power)

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}
