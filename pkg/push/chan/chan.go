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
//	    "context"
//	    "fmt"
//	    "os"
//	    "os/signal"
//	    "syscall"
//
//	    "github.com/shing1211/futuapi4go/client"
//	    "github.com/shing1211/futuapi4go/pkg/constant"
//	    "github.com/shing1211/futuapi4go/pkg/push"
//	    chanpkg "github.com/shing1211/futuapi4go/pkg/push/chan"
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
//	// Graceful shutdown on Ctrl+C
//	sig := make(chan os.Signal, 1)
//	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
//
//	for {
//	    select {
//	    case q := <-ch:
//	        fmt.Printf("Quote: %s price=%.2f\n", q.Security.GetCode(), q.CurPrice)
//	    case <-sig:
//	        fmt.Println("Shutting down...")
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
	"context"

	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/push"
)

const (
	DefaultChanBufferSize = 100
	MaxChanBufferSize  = 10000
)

type stopFunc func()

func WithBufferSize(size int) int {
	if size <= 0 {
		return DefaultChanBufferSize
	}
	if size > MaxChanBufferSize {
		return MaxChanBufferSize
	}
	return size
}

func NewQuoteChannel(bufferSize int) chan *push.UpdateBasicQot {
	return make(chan *push.UpdateBasicQot, WithBufferSize(bufferSize))
}

func NewKLChannel(bufferSize int) chan *push.UpdateKL {
	return make(chan *push.UpdateKL, WithBufferSize(bufferSize))
}

func NewTickerChannel(bufferSize int) chan *push.UpdateTicker {
	return make(chan *push.UpdateTicker, WithBufferSize(bufferSize))
}

func NewOrderBookChannel(bufferSize int) chan *push.UpdateOrderBook {
	return make(chan *push.UpdateOrderBook, WithBufferSize(bufferSize))
}

func NewRTChannel(bufferSize int) chan *push.UpdateRT {
	return make(chan *push.UpdateRT, WithBufferSize(bufferSize))
}

func NewBrokerChannel(bufferSize int) chan *push.UpdateBroker {
	return make(chan *push.UpdateBroker, WithBufferSize(bufferSize))
}

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

func SubscribeQuote(cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateBasicQot) stopFunc {
	client.Subscribe(context.Background(), cli, market, code, []constant.SubType{constant.SubType_Quote})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateBasicQot, push.ParseUpdateBasicQot, ch)
}

func SubscribeKLine(cli *client.Client, market constant.Market, code string, klType constant.KLType, ch chan<- *push.UpdateKL) stopFunc {
	client.Subscribe(context.Background(), cli, market, code, []constant.SubType{klTypeToSubType(klType)})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateKL, push.ParseUpdateKL, ch)
}

func SubscribeKLines(cli *client.Client, market constant.Market, code string, handlers map[constant.KLType]func(*push.UpdateKL)) func() {
	if len(handlers) == 0 {
		return func() {}
	}

	subtypes := make([]constant.SubType, 0, len(handlers))
	for kt := range handlers {
		subtypes = append(subtypes, klTypeToSubType(kt))
	}
	client.Subscribe(context.Background(), cli, market, code, subtypes)

	cli.RegisterHandler(push.ProtoID_Qot_UpdateKL, func(pid uint32, body []byte) {
		data, err := push.ParseUpdateKL(body)
		if err != nil || data == nil {
			return
		}
		if cb, ok := handlers[constant.KLType(data.KlType)]; ok {
			cb(data)
		}
	})

	return func() {
		cli.RegisterHandler(push.ProtoID_Qot_UpdateKL, nil)
	}
}

func klTypeToSubType(k constant.KLType) constant.SubType {
	switch k {
	case constant.KLType_K_1Min:
		return constant.SubType_K_1Min
	case constant.KLType_K_5Min:
		return constant.SubType_K_5Min
	case constant.KLType_K_15Min:
		return constant.SubType_K_15Min
	case constant.KLType_K_30Min:
		return constant.SubType_K_30Min
	case constant.KLType_K_60Min:
		return constant.SubType_K_60Min
	case constant.KLType_K_Day:
		return constant.SubType_K_Day
	case constant.KLType_K_Week:
		return constant.SubType_K_Week
	case constant.KLType_K_Month:
		return constant.SubType_K_Month
	case constant.KLType_K_Quarter:
		return constant.SubType_K_Quarter
	case constant.KLType_K_Year:
		return constant.SubType_K_Year
	case constant.KLType_K_3Min:
		return constant.SubType_K_3Min
	default:
		return constant.SubType_K_1Min
	}
}

func SubscribeTicker(cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateTicker) stopFunc {
	client.Subscribe(context.Background(), cli, market, code, []constant.SubType{constant.SubType_Ticker})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateTicker, push.ParseUpdateTicker, ch)
}

func SubscribeOrderBook(cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateOrderBook) stopFunc {
	client.Subscribe(context.Background(), cli, market, code, []constant.SubType{constant.SubType_OrderBook})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateOrderBook, push.ParseUpdateOrderBook, ch)
}

func SubscribeRT(cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateRT) stopFunc {
	client.Subscribe(context.Background(), cli, market, code, []constant.SubType{constant.SubType_RT})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateRT, push.ParseUpdateRT, ch)
}

func SubscribeBroker(cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateBroker) stopFunc {
	client.Subscribe(context.Background(), cli, market, code, []constant.SubType{constant.SubType_Broker})
	return subscribeOne(cli, push.ProtoID_Qot_UpdateBroker, push.ParseUpdateBroker, ch)
}

func SubscribePriceReminder(cli *client.Client, ch chan<- *push.UpdatePriceReminder) stopFunc {
	return subscribeOne(cli, push.ProtoID_Qot_UpdatePriceReminder, push.ParseUpdatePriceReminder, ch)
}
