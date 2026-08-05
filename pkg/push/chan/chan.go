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
// # Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package chanpkg

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
	qotcommon "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsubeventcontract"
	"github.com/shing1211/futuapi4go/pkg/push"
)

const (
	DefaultChanBufferSize = 100
	MaxChanBufferSize     = 10000
)

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

func NewSystemNotifyChannel(bufferSize int) chan *push.SystemNotify {
	return make(chan *push.SystemNotify, WithBufferSize(bufferSize))
}

func NewOrderUpdateChannel(bufferSize int) chan *push.UpdateOrder {
	return make(chan *push.UpdateOrder, WithBufferSize(bufferSize))
}

func NewOrderFillUpdateChannel(bufferSize int) chan *push.UpdateOrderFill {
	return make(chan *push.UpdateOrderFill, WithBufferSize(bufferSize))
}

func NewTrdNotifyChannel(bufferSize int) chan *push.TrdNotify {
	return make(chan *push.TrdNotify, WithBufferSize(bufferSize))
}

func subscribeOne[T any](ctx context.Context, cli *client.Client, protoID uint32, ch chan<- *T, parseFn func([]byte) (*T, error), subFn func() error) (func(), error) {
	if err := subFn(); err != nil {
		return nil, err
	}

	stopCh := make(chan struct{})
	var once sync.Once

	handler := func(pid uint32, body []byte) {
		select {
		case <-stopCh:
			return
		case <-ctx.Done():
			return
		default:
		}

		data, err := parseFn(body)
		if err != nil {
			slog.Warn("push parse failed", "protoID", pid, "error", err)
			return
		}
		if data == nil {
			slog.Debug("push parse returned nil", "protoID", pid)
			return
		}

		select {
		case ch <- data:
		case <-stopCh:
		case <-ctx.Done():
		}
	}

	cli.RegisterHandler(protoID, handler)

	stop := func() {
		once.Do(func() {
			close(stopCh)
			cli.UnregisterHandler(protoID)
		})
	}

	return stop, nil
}

func SubscribeQuote(ctx context.Context, cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateBasicQot) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateBasicQot, ch, push.ParseUpdateBasicQot, func() error {
		return client.Subscribe(ctx, cli, market, code, []constant.SubType{constant.SubType_Quote})
	})
}

func SubscribeKLine(ctx context.Context, cli *client.Client, market constant.Market, code string, klType constant.KLType, ch chan<- *push.UpdateKL) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateKL, ch, push.ParseUpdateKL, func() error {
		st, err := klTypeToSubType(klType)
		if err != nil {
			return err
		}
		return client.Subscribe(ctx, cli, market, code, []constant.SubType{st})
	})
}

func SubscribeKLines(ctx context.Context, cli *client.Client, market constant.Market, code string, kTypes []constant.KLType, ch chan<- *push.UpdateKL) (func(), error) {
	ktMap := make(map[int32]struct{}, len(kTypes))
	for _, kt := range kTypes {
		ktMap[int32(kt)] = struct{}{}
	}

	wrappedParse := func(body []byte) (*push.UpdateKL, error) {
		kl, err := push.ParseUpdateKL(body)
		if kl != nil {
			if _, ok := ktMap[kl.KlType]; !ok {
				return nil, nil
			}
		}
		return kl, err
	}

	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateKL, ch, wrappedParse, func() error {
		return subscribe(ctx, cli, market, code, kTypes)
	})
}

func subscribe(ctx context.Context, cli *client.Client, market constant.Market, code string, kTypes []constant.KLType) error {
	subtypes := make([]constant.SubType, 0, len(kTypes))
	for _, kt := range kTypes {
		st, err := klTypeToSubType(kt)
		if err != nil {
			return err
		}
		subtypes = append(subtypes, st)
	}
	return client.Subscribe(ctx, cli, market, code, subtypes)
}

func klTypeToSubType(k constant.KLType) (constant.SubType, error) {
	switch k {
	case constant.KLType_K_1Min:
		return constant.SubType_K_1Min, nil
	case constant.KLType_K_5Min:
		return constant.SubType_K_5Min, nil
	case constant.KLType_K_15Min:
		return constant.SubType_K_15Min, nil
	case constant.KLType_K_30Min:
		return constant.SubType_K_30Min, nil
	case constant.KLType_K_60Min:
		return constant.SubType_K_60Min, nil
	case constant.KLType_K_Day:
		return constant.SubType_K_Day, nil
	case constant.KLType_K_Week:
		return constant.SubType_K_Week, nil
	case constant.KLType_K_Month:
		return constant.SubType_K_Month, nil
	case constant.KLType_K_Quarter:
		return constant.SubType_K_Quarter, nil
	case constant.KLType_K_Year:
		return constant.SubType_K_Year, nil
	case constant.KLType_K_3Min:
		return constant.SubType_K_3Min, nil
	default:
		return constant.SubType_K_1Min, fmt.Errorf("unknown KLType: %d", k)
	}
}

func SubscribeTicker(ctx context.Context, cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateTicker) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateTicker, ch, push.ParseUpdateTicker, func() error {
		return client.Subscribe(ctx, cli, market, code, []constant.SubType{constant.SubType_Ticker})
	})
}

func SubscribeOrderBook(ctx context.Context, cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateOrderBook) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateOrderBook, ch, push.ParseUpdateOrderBook, func() error {
		return client.Subscribe(ctx, cli, market, code, []constant.SubType{constant.SubType_OrderBook})
	})
}

func SubscribeRT(ctx context.Context, cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateRT) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateRT, ch, push.ParseUpdateRT, func() error {
		return client.Subscribe(ctx, cli, market, code, []constant.SubType{constant.SubType_RT})
	})
}

func SubscribeBroker(ctx context.Context, cli *client.Client, market constant.Market, code string, ch chan<- *push.UpdateBroker) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateBroker, ch, push.ParseUpdateBroker, func() error {
		return client.Subscribe(ctx, cli, market, code, []constant.SubType{constant.SubType_Broker})
	})
}

func SubscribePriceReminder(ctx context.Context, cli *client.Client, ch chan<- *push.UpdatePriceReminder) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdatePriceReminder, ch, push.ParseUpdatePriceReminder, func() error {
		return nil
	})
}

func SubscribeSystemNotify(ctx context.Context, cli *client.Client, ch chan<- *push.SystemNotify) (func(), error) {
	return subscribeOne(ctx, cli, constant.ProtoID_Notify, ch, push.ParseSystemNotify, func() error {
		return nil
	})
}

func SubscribeOrderUpdate(ctx context.Context, cli *client.Client, ch chan<- *push.UpdateOrder) (func(), error) {
	return subscribeOne(ctx, cli, constant.ProtoID_Trd_UpdateOrder, ch, push.ParseUpdateOrder, func() error {
		return nil
	})
}

func SubscribeOrderFillUpdate(ctx context.Context, cli *client.Client, ch chan<- *push.UpdateOrderFill) (func(), error) {
	return subscribeOne(ctx, cli, constant.ProtoID_Trd_UpdateOrderFill, ch, push.ParseUpdateOrderFill, func() error {
		return nil
	})
}

func SubscribeTrdNotify(ctx context.Context, cli *client.Client, ch chan<- *push.TrdNotify) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Trd_Notify, ch, push.ParseTrdNotify, func() error {
		return nil
	})
}

func NewOptionEventChannel(bufferSize int) chan *push.UpdateOptionEvent {
	return make(chan *push.UpdateOptionEvent, WithBufferSize(bufferSize))
}

func NewPushIndicatorCalcChannel(bufferSize int) chan *push.PushIndicatorCalc {
	return make(chan *push.PushIndicatorCalc, WithBufferSize(bufferSize))
}

func NewEventContractOrderBookChannel(bufferSize int) chan *push.UpdateEventContractOrderBook {
	return make(chan *push.UpdateEventContractOrderBook, WithBufferSize(bufferSize))
}

func NewEventContractKLineChannel(bufferSize int) chan *push.UpdateEventContractKline {
	return make(chan *push.UpdateEventContractKline, WithBufferSize(bufferSize))
}

func NewEventContractTickerChannel(bufferSize int) chan *push.UpdateEventContractTicker {
	return make(chan *push.UpdateEventContractTicker, WithBufferSize(bufferSize))
}

// SubscribeOptionEvent subscribes to option event alert pushes. These pushes are
// server-triggered (e.g. option listing / expiration / assignment alerts); no
// client-side subscribe request is required.
func SubscribeOptionEvent(ctx context.Context, cli *client.Client, ch chan<- *push.UpdateOptionEvent) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateOptionEvent, ch, push.ParseUpdateOptionEvent, func() error {
		return nil
	})
}

// SubscribePushIndicatorCalc subscribes to asynchronous indicator calculation
// results. The caller must first issue a RequestIndicatorCalc to start a
// calculation; the matching result (by calcId) is then delivered on the channel.
func SubscribePushIndicatorCalc(ctx context.Context, cli *client.Client, ch chan<- *push.PushIndicatorCalc) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_PushIndicatorCalc, ch, push.ParsePushIndicatorCalc, func() error {
		return nil
	})
}

func ecSecurity(code string) *qotcommon.Security {
	return client.NewECSecurity(code)
}

func SubscribeEventContractOrderBook(ctx context.Context, cli *client.Client, code string, ch chan<- *push.UpdateEventContractOrderBook) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateEventContractOrderBook, ch, push.ParseUpdateEventContractOrderBook, func() error {
		return client.SubEventContract(ctx, cli, &qotsubeventcontract.C2S{
			SecurityList:     []*qotcommon.Security{ecSecurity(code)},
			SubTypeList:      []int32{int32(constant.SubType_OrderBook)},
			IsSubOrUnSub:     ptrToBool(true),
			IsRegOrUnRegPush: ptrToBool(true),
			IsFirstPush:      ptrToBool(true),
		})
	})
}

func SubscribeEventContractKLine(ctx context.Context, cli *client.Client, code string, klType constant.KLType, ch chan<- *push.UpdateEventContractKline) (func(), error) {
	st, err := klTypeToSubType(klType)
	if err != nil {
		return nil, err
	}
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateEventContractKline, ch, push.ParseUpdateEventContractKline, func() error {
		return client.SubEventContract(ctx, cli, &qotsubeventcontract.C2S{
			SecurityList:     []*qotcommon.Security{ecSecurity(code)},
			SubTypeList:      []int32{int32(st)},
			IsSubOrUnSub:     ptrToBool(true),
			IsRegOrUnRegPush: ptrToBool(true),
			IsFirstPush:      ptrToBool(true),
			KlineSource:      []qotcommon.EC_KlineSource{qotcommon.EC_KlineSource_EC_KlineSource_None},
		})
	})
}

func SubscribeEventContractTicker(ctx context.Context, cli *client.Client, code string, ch chan<- *push.UpdateEventContractTicker) (func(), error) {
	return subscribeOne(ctx, cli, push.ProtoID_Qot_UpdateEventContractTicker, ch, push.ParseUpdateEventContractTicker, func() error {
		return client.SubEventContract(ctx, cli, &qotsubeventcontract.C2S{
			SecurityList:     []*qotcommon.Security{ecSecurity(code)},
			SubTypeList:      []int32{int32(constant.SubType_Ticker)},
			IsSubOrUnSub:     ptrToBool(true),
			IsRegOrUnRegPush: ptrToBool(true),
			IsFirstPush:      ptrToBool(true),
		})
	})
}

func ptrToBool(v bool) *bool { return &v }
