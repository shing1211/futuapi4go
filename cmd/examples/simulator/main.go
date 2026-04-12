package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shing1211/futuapi4go/cmd/simulator"
	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	fmt.Println("=== Starting Simulator ===")

	srv := simulator.New("127.0.0.1:11111")
	srv.RegisterDefaultHandlers()
	srv.RegisterQotHandlers()
	srv.AddSecurity(int32(qotcommon.QotMarket_QotMarket_HK_Security), "00700")

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start simulator: %v", err)
	}
	defer srv.Stop()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Simulator running on 127.0.0.1:11111")

	fmt.Println("=== Connecting SDK Client ===")
	cli := futuapi.New()
	defer cli.Close()

	if err := cli.Connect("127.0.0.1:11111"); err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	fmt.Printf("Connected! ConnID=%d\n", cli.GetConnID())

	market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	securities := []*qotcommon.Security{
		{Market: &market, Code: func() *string { s := "00700"; return &s }()},
	}

	fmt.Println("=== Calling GetBasicQot ===")
	result, err := qot.GetBasicQot(cli, securities)
	if err != nil {
		log.Fatalf("GetBasicQot failed: %v", err)
	}

	fmt.Println("--- Results ---")
	for _, bq := range result {
		fmt.Printf("%s: CurPrice=%.2f Volume=%d\n", bq.Security.GetCode(), bq.CurPrice, bq.Volume)
	}

	fmt.Println("\n=== Success! SDK works with Simulator ===")
}
