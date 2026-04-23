// Package chan provides channel-based push notification delivery for the futuapi4go SDK.
//
// This is the Go-native alternative to callback-based push handlers. It allows
// receiving real-time market data and order updates via Go channels, enabling
// idiomatic concurrent patterns like select statements, goroutines, and
// fan-out processing.
//
// Usage:
//
//	import (
//	    "github.com/shing1211/futuapi4go/client"
//	    "github.com/shing1211/futuapi4go/pkg/constant"
//	    "github.com/shing1211/futuapi4go/pkg/push/chanpkg"
//	)
//
//	cli := client.New()
//	defer cli.Close()
//	if err := cli.Connect("127.0.0.1:11111"); err != nil {
//	    log.Fatal(err)
//	}
//
//	// Subscribe to quotes
//	ch := make(chan *push.UpdateBasicQot, 100)
//	stop := chanpkg.SubscribeQuote(cli, constant.Market_HK, "00700", ch)
//	defer stop()
//
//	for {
//	    select {
//	    case q := <-ch:
//	        fmt.Printf("Quote: %s price=%.2f\n", q.Security.GetCode(), q.CurPrice)
//	    case <-time.After(30 * time.Second):
//	        fmt.Println("timeout")
//	        return
//	    }
//	}
//
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
package chanpkg

import (
	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/push"
)

type stopFunc func()

func subscribeOne[T any](
	cli *client.Client,
	protoID uint32,
	parse func([]byte) (*T, error),
	ch chan<- *T,
) stopFunc {
	cli.RegisterHandler(protoID, func(pid uint32, body []byte) {
		data, err := parse(body)
		if err != nil || data == nil {
			return
		}
		select {
		case ch <- data:
		default:
		}
	})
	return func() {
		cli.RegisterHandler(protoID, nil)
	}
}

func SubscribeQuote(cli *client.Client, market int32, code string, ch chan<- *push.UpdateBasicQot) stopFunc {
	client.Subscribe(cli, market, code, []int32{int32(constant.SubType_Quote)})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateBasicQot, push.ParseUpdateBasicQot, ch)
}

func SubscribeKLine(cli *client.Client, market int32, code string, klType int32, ch chan<- *push.UpdateKL) stopFunc {
	client.Subscribe(cli, market, code, []int32{klType})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateKL, push.ParseUpdateKL, ch)
}

func SubscribeTicker(cli *client.Client, market int32, code string, ch chan<- *push.UpdateTicker) stopFunc {
	client.Subscribe(cli, market, code, []int32{int32(constant.SubType_Ticker)})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateTicker, push.ParseUpdateTicker, ch)
}

func SubscribeOrderBook(cli *client.Client, market int32, code string, ch chan<- *push.UpdateOrderBook) stopFunc {
	client.Subscribe(cli, market, code, []int32{int32(constant.SubType_OrderBook)})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateOrderBook, push.ParseUpdateOrderBook, ch)
}

func SubscribeRT(cli *client.Client, market int32, code string, ch chan<- *push.UpdateRT) stopFunc {
	client.Subscribe(cli, market, code, []int32{int32(constant.SubType_RT)})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateRT, push.ParseUpdateRT, ch)
}

func SubscribeBroker(cli *client.Client, market int32, code string, ch chan<- *push.UpdateBroker) stopFunc {
	client.Subscribe(cli, market, code, []int32{int32(constant.SubType_Broker)})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateBroker, push.ParseUpdateBroker, ch)
}

func SubscribePriceReminder(cli *client.Client, ch chan<- *push.UpdatePriceReminder) stopFunc {
	return subscribeOne(cli, push.ProtoID_Qot_UpdatePriceReminder, push.ParseUpdatePriceReminder, ch)
}
