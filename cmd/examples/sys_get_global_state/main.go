// Example: GetGlobalState - 獲取全局狀態
//
// This example demonstrates how to use the GetGlobalState API to retrieve
// the global state of Futu OpenD including market status, login status, etc.
//
// Usage:
//   go run main.go

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/sys"
)

func main() {
	cli := futuapi.New()
	defer cli.Close()

	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Println("=== GetGlobalState Example / 獲取全局狀態示例 ===")
	fmt.Printf("Connecting to %s...\n", addr)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("✅ Connected! ConnID=%d\n\n", cli.GetConnID())

	// Get global state / 獲取全局狀態
	resp, err := sys.GetGlobalState(cli)
	if err != nil {
		log.Printf("GetGlobalState failed: %v", err)
		return
	}

	fmt.Println("📊 Futu OpenD Global State / Futu OpenD 全局狀態")
	fmt.Println()

	// Market status / 市場狀態
	fmt.Println("🌏 Market Status / 市場狀態:")
	fmt.Printf("  HK Market / 港股市場:     %d\n", resp.MarketHK)
	fmt.Printf("  US Market / 美股市場:     %d\n", resp.MarketUS)
	fmt.Printf("  SH Market / 滬股市场:     %d\n", resp.MarketSH)
	fmt.Printf("  SZ Market / 深股市场:     %d\n", resp.MarketSZ)
	fmt.Println()

	// Login status / 登錄狀態
	fmt.Println("🔐 Login Status / 登錄狀態:")
	qotStatus := "Not Logged In"
	if resp.QotLogined {
		qotStatus = "Logged In"
	}
	trdStatus := "Not Logged In"
	if resp.TrdLogined {
		trdStatus = "Logged In"
	}
	fmt.Printf("  Quote Server / 行情服務器: %s\n", qotStatus)
	fmt.Printf("  Trade Server / 交易服務器: %s\n", trdStatus)
	fmt.Println()

	// Server info / 服務器信息
	fmt.Println("🖥️  Server Info / 服務器信息:")
	fmt.Printf("  Server Version / 服務器版本:  %d\n", resp.ServerVer)
	fmt.Printf("  Build Number / 構建號:        %d\n", resp.ServerBuildNo)
	fmt.Printf("  ConnID / 連接ID:              %d\n", resp.ConnID)
	
	// Time info / 時間信息
	fmt.Println("\n🕐 Time Info / 時間信息:")
	serverTime := time.Unix(resp.Time, 0)
	fmt.Printf("  Server Time / 服務器時間:  %s\n", serverTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Local Time / 本地時間:     %s\n", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("\n=== Example Complete / 示例完成 ===")
}

