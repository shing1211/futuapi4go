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

// Real-time Push Stream Example
//
// This example demonstrates live streaming via the futuapi4go SDK:
//   - Connect to Futu OpenD (TCP, not WebSocket)
//   - Subscribe to HSImain for multiple data types
//   - Handle real-time push notifications and print live data
//   - Graceful shutdown on interrupt signal
//
// Usage:
//
//	go run main.go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/push"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

var (
	interrupt atomic.Bool
	stats     streamStats
)

type streamStats struct {
	basicQot     atomic.Int64
	kl           atomic.Int64
	orderBook    atomic.Int64
	ticker       atomic.Int64
	rt           atomic.Int64
	broker       atomic.Int64
	unknownProto atomic.Int64
}

func main() {
	addr := os.Getenv("FUTU_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}

	fmt.Printf("=== Real-time Push Stream Test ===\n")
	fmt.Printf("Connecting to %s ...\n", addr)

	cli := futuapi.New()
	defer cli.Close()

	setupPushHandler(cli)

	if err := cli.Connect(addr); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Printf("Connected! ConnID=%d, ServerVer=%d\n\n", cli.GetConnID(), cli.GetServerVer())

	if err := runStreamTest(cli); err != nil {
		log.Printf("Stream test error: %v", err)
	}

	printSummary()
}

func setupPushHandler(cli *futuapi.Client) {
	cli.SetPushHandler(func(pkt *futuapi.Packet) {
		var handled bool
		switch pkt.Header.ProtoID {
		case push.ProtoID_Qot_UpdateBasicQot:
			data, err := push.ParseUpdateBasicQot(pkt.Body)
			if err != nil {
				log.Printf("ParseUpdateBasicQot error: %v", err)
				return
			}
			if data != nil {
				stats.basicQot.Add(1)
				fmt.Printf("[BasicQot] %d %s: price=%.2f open=%.2f high=%.2f low=%.2f vol=%d\n",
					data.Security.GetMarket(), data.Security.GetCode(), data.CurPrice,
					data.OpenPrice, data.HighPrice, data.LowPrice, data.Volume)
			}
			handled = true

		case push.ProtoID_Qot_UpdateKL:
			data, err := push.ParseUpdateKL(pkt.Body)
			if err != nil {
				log.Printf("ParseUpdateKL error: %v", err)
				return
			}
			if data != nil {
				stats.kl.Add(1)
				for _, kl := range data.KLList {
					fmt.Printf("[KL]       %d %s: open=%.2f high=%.2f low=%.2f close=%.2f vol=%d\n",
						data.Security.GetMarket(), data.Security.GetCode(),
						kl.GetOpenPrice(), kl.GetHighPrice(), kl.GetLowPrice(), kl.GetClosePrice(), kl.GetVolume())
				}
			}
			handled = true

		case push.ProtoID_Qot_UpdateOrderBook:
			data, err := push.ParseUpdateOrderBook(pkt.Body)
			if err != nil {
				log.Printf("ParseUpdateOrderBook error: %v", err)
				return
			}
			if data != nil {
				stats.orderBook.Add(1)
				fmt.Printf("[OrderBook] %d %s: ask_count=%d bid_count=%d\n",
					data.Security.GetMarket(), data.Security.GetCode(),
					len(data.OrderBookAskList), len(data.OrderBookBidList))
				for i, ask := range data.OrderBookAskList {
					if i >= 3 {
						fmt.Printf("           ... +%d more asks\n", len(data.OrderBookAskList)-3)
						break
					}
					fmt.Printf("           Ask[%d] price=%.2f vol=%d\n", i, ask.GetPrice(), ask.GetVolume())
				}
				for i, bid := range data.OrderBookBidList {
					if i >= 3 {
						fmt.Printf("           ... +%d more bids\n", len(data.OrderBookBidList)-3)
						break
					}
					fmt.Printf("           Bid[%d] price=%.2f vol=%d\n", i, bid.GetPrice(), bid.GetVolume())
				}
			}
			handled = true

		case push.ProtoID_Qot_UpdateTicker:
			data, err := push.ParseUpdateTicker(pkt.Body)
			if err != nil {
				log.Printf("ParseUpdateTicker error: %v", err)
				return
			}
			if data != nil {
				stats.ticker.Add(1)
				for _, t := range data.TickerList {
					fmt.Printf("[Ticker]   %d %s: price=%.2f vol=%d turnover=%.2f dir=%s time=%s\n",
						data.Security.GetMarket(), data.Security.GetCode(),
						t.GetPrice(), t.GetVolume(), t.GetTurnover(), direction(t.GetDir()), t.GetTime())
				}
			}
			handled = true

		case push.ProtoID_Qot_UpdateRT:
			data, err := push.ParseUpdateRT(pkt.Body)
			if err != nil {
				log.Printf("ParseUpdateRT error: %v", err)
				return
			}
			if data != nil {
				stats.rt.Add(1)
				for _, rt := range data.RTList {
					fmt.Printf("[RT]       %d %s: price=%.2f vol=%d avg_price=%.2f time=%s\n",
						data.Security.GetMarket(), data.Security.GetCode(),
						rt.GetPrice(), rt.GetVolume(), rt.GetAvgPrice(), rt.GetTime())
				}
			}
			handled = true

		case push.ProtoID_Qot_UpdateBroker:
			data, err := push.ParseUpdateBroker(pkt.Body)
			if err != nil {
				log.Printf("ParseUpdateBroker error: %v", err)
				return
			}
			if data != nil {
				stats.broker.Add(1)
				fmt.Printf("[Broker]   %d %s: asks=%d bids=%d\n",
					data.Security.GetMarket(), data.Security.GetCode(),
					len(data.AskBrokerList), len(data.BidBrokerList))
			}
			handled = true
		}

		if !handled {
			stats.unknownProto.Add(1)
			fmt.Printf("[Unknown]  ProtoID=%d, BodyLen=%d\n", pkt.Header.ProtoID, len(pkt.Body))
		}
	})
}

func runStreamTest(cli *futuapi.Client) error {
	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	security := &qotcommon.Security{
		Market: &hkMarket,
		Code:   ptrStr("HSImain"),
	}

	fmt.Printf("Subscribing to HSImain (HK futures main contract)...\n\n")
	subTypes := []qot.SubType{
		qot.SubType_Basic,
		qot.SubType_KL,
		qot.SubType_OrderBook,
		qot.SubType_Ticker,
		qot.SubType_RT,
	}

	subReq := &qot.SubscribeRequest{
		SecurityList:     []*qotcommon.Security{security},
		SubTypeList:      subTypes,
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
	}

	subResp, err := qot.Subscribe(cli, subReq)
	if err != nil {
		return fmt.Errorf("Subscribe failed: %w", err)
	}
	if subResp.RetType != 0 {
		return fmt.Errorf("Subscribe returned error: RetType=%d, RetMsg=%s", subResp.RetType, subResp.RetMsg)
	}
	fmt.Printf("Subscribe OK: RetType=%d RetMsg=%s\n\n", subResp.RetType, subResp.RetMsg)

	subInfo, err := qot.GetSubInfo(cli)
	if err != nil {
		fmt.Printf("GetSubInfo: %v\n", err)
	} else {
		fmt.Printf("Subscription Info: TotalUsed=%d Remain=%d\n", subInfo.TotalUsedQuota, subInfo.RemainQuota)
	}
	fmt.Println()
	fmt.Println("Waiting for push data... (Ctrl+C to stop)")
	fmt.Println("---")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	lastPrint := time.Now()
	for {
		select {
		case <-sigCh:
			fmt.Println("\nInterrupted, cleaning up...")

			fmt.Println("Unsubscribing from HSImain...")
			unsubReq := &qot.SubscribeRequest{
				SecurityList:     []*qotcommon.Security{security},
				SubTypeList:      subTypes,
				IsSubOrUnSub:     false,
				IsRegOrUnRegPush: false,
			}
			qot.Subscribe(cli, unsubReq)

			interrupt.Store(true)
			return nil
		case <-ticker.C:
			elapsed := time.Since(lastPrint).Seconds()
			lastPrint = time.Now()
			fmt.Printf("\n[Stats %.0fs] basicQot=%d kl=%d orderBook=%d ticker=%d rt=%d broker=%d unknown=%d\n",
				elapsed,
				stats.basicQot.Load(),
				stats.kl.Load(),
				stats.orderBook.Load(),
				stats.ticker.Load(),
				stats.rt.Load(),
				stats.broker.Load(),
				stats.unknownProto.Load(),
			)
			fmt.Println("---")
		}
	}
}

func printSummary() {
	fmt.Println("\n=== Stream Test Summary ===")
	fmt.Printf("  BasicQot pushes received:  %d\n", stats.basicQot.Load())
	fmt.Printf("  KL pushes received:         %d\n", stats.kl.Load())
	fmt.Printf("  OrderBook pushes received:  %d\n", stats.orderBook.Load())
	fmt.Printf("  Ticker pushes received:     %d\n", stats.ticker.Load())
	fmt.Printf("  RT pushes received:         %d\n", stats.rt.Load())
	fmt.Printf("  Broker pushes received:     %d\n", stats.broker.Load())
	fmt.Printf("  Unknown proto IDs:          %d\n", stats.unknownProto.Load())

	total := stats.basicQot.Load() + stats.kl.Load() + stats.orderBook.Load() +
		stats.ticker.Load() + stats.rt.Load() + stats.broker.Load() + stats.unknownProto.Load()
	fmt.Printf("  Total pushes:               %d\n", total)

	if total == 0 {
		fmt.Println("\nNo pushes received. Possible causes:")
		fmt.Println("  - HSImain may not be trading right now (market closed or holiday)")
		fmt.Println("  - Futu OpenD is not connected to the market data feed")
		fmt.Println("  - Check that OpenD is running: Tools -> Market Alerts & Trading")
	}
}

func direction(d int32) string {
	switch d {
	case 1:
		return "Buy"
	case 2:
		return "Sell"
	default:
		return "Neutral"
	}
}

func ptrStr(s string) *string {
	return &s
}
