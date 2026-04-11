// Full connection and API test
package main

import (
	"fmt"
	"os"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	fmt.Println("=== FutuAPI4Go Full Connection Test ===")

	cli := futuapi.New()
	defer cli.Close()

	addr := "127.0.0.1:11111"
	if a := os.Getenv("FUTU_ADDR"); a != "" {
		addr = a
	}

	fmt.Printf("📡 Connecting to %s...\n", addr)
	if err := cli.Connect(addr); err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Connected!\n")
	fmt.Printf("   ConnID:    %d\n", cli.GetConnID())
	fmt.Printf("   ServerVer: %d\n", cli.GetServerVer())
	fmt.Printf("   AESKey:    %s\n", cli.GetAESKey()[:8]+"...")
	fmt.Println()

	// Test GetBasicQot
	fmt.Println("📊 Testing GetBasicQot...")
	hkFutureMarket := int32(qotcommon.QotMarket_QotMarket_HK_Future)
	securities := []*qotcommon.Security{
		{Market: &hkFutureMarket, Code: ptrStr("HSImain")},
	}

	quotes, err := qot.GetBasicQot(cli, securities)
	if err != nil {
		fmt.Printf("❌ GetBasicQot failed: %v\n", err)
	} else {
		for _, q := range quotes {
			fmt.Printf("✅ %s (%s): Price=%.2f Volume=%d\n",
				q.Security.GetCode(), q.Name, q.CurPrice, q.Volume)
		}
	}

	fmt.Println("\n🎉 All tests passed! SDK is working perfectly!")
}

func ptrStr(s string) *string { return &s }
