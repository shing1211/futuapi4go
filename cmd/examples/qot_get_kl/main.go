// Example: GetKL - 獲取K線數據
//
// This example demonstrates how to use the GetKL API to retrieve
// K-line (candlestick) data for a security.
//
// Usage:
//   go run main.go

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

	fmt.Println("=== GetKL Example / 獲取K線數據示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{
		Market: &hkMarket,
		Code:   ptrStr("00700"), // Tencent / 騰訊
	}

	// Example 1: Daily K-line / 日K線
	fmt.Println("📊 Daily K-line (5 bars) / 日K線 (5根):")
	getKL(cli, security, qotcommon.KLType_KLType_Day, 5)

	// Example 2: 1-Minute K-line / 1分鐘K線
	fmt.Println("\n📊 1-Minute K-line (10 bars) / 1分鐘K線 (10根):")
	getKL(cli, security, qotcommon.KLType_KLType_1Min, 10)

	// Example 3: Weekly K-line / 周K線
	fmt.Println("\n📊 Weekly K-line (3 bars) / 周K線 (3根):")
	getKL(cli, security, qotcommon.KLType_KLType_Week, 3)

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

func getKL(cli *futuapi.Client, security *qotcommon.Security, klType qotcommon.KLType, count int32) {
	req := &qot.GetKLRequest{
		Security:  security,
		RehabType: int32(qotcommon.RehabType_RehabType_None), // No rehabilitation / 不復權
		KLType:    int32(klType),
		ReqNum:    count,
	}

	resp, err := qot.GetKL(cli, req)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("  Retrieved %d K-lines for %s\n\n", len(resp.KLList), security.GetCode())
	fmt.Printf("  %-20s %-10s %-10s %-10s %-10s %-12s\n",
		"Time", "Open", "High", "Low", "Close", "Volume")
	fmt.Println("  " + "--------------------------------------------------------------------------")

	for _, kl := range resp.KLList {
		fmt.Printf("  %-20s %-10.2f %-10.2f %-10.2f %-10.2f %-12d\n",
			kl.Time, kl.OpenPrice, kl.HighPrice, kl.LowPrice, kl.ClosePrice, kl.Volume)
	}
}

func ptrStr(s string) *string { return &s }

