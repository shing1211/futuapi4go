package client_test

import (
	"context"
	"fmt"
	"log"

	"github.com/shing1211/futuapi4go/client"
	"github.com/shing1211/futuapi4go/pkg/constant"
)

func ExampleNew() {
	cli := client.New(client.WithLogLevel(3))
	defer cli.Close()
	fmt.Println(cli.IsConnected())
	// Output: false
}

func ExampleNew_encrypted() {
	cli := client.New(
		client.WithLogLevel(3),
		client.WithRSAPrivateKey("testdata/private_key.pem"),
		client.WithEncryption(true),
	)
	defer cli.Close()
	_ = cli
}

func ExampleWithEnvConfig() {
	cli := client.New(client.WithEnvConfig())
	defer cli.Close()
	_ = cli
}

func ExampleClient_OnQuote() {
	cli := client.New(client.WithLogLevel(3)).
		OnQuote(func(q *client.PushQuote) error {
			log.Printf("%s: %.2f", q.Code, q.CurPrice)
			return nil
		})
	defer cli.Close()
	_ = cli
}

func ExampleClient_OnOrder() {
	cli := client.New(client.WithLogLevel(3)).
		OnOrder(func(o *client.PushOrderUpdate) error {
			log.Printf("Order %d status=%d", o.OrderID, o.OrderStatus)
			return nil
		})
	defer cli.Close()
	_ = cli
}

func ExampleClient_OnOrderFill() {
	cli := client.New(client.WithLogLevel(3)).
		OnOrderFill(func(f *client.PushOrderFill) error {
			log.Printf("Fill %d qty=%.0f price=%.2f", f.FillID, f.Qty, f.Price)
			return nil
		})
	defer cli.Close()
	_ = cli
}

func ExampleClient_OnKLine() {
	cli := client.New(client.WithLogLevel(3)).
		OnKLine(func(k *client.PushKLine) error {
			log.Printf("%s K-line close=%.2f", k.Code, k.KLine.Close)
			return nil
		})
	defer cli.Close()
	_ = cli
}

func ExampleClient_OnOrderBook() {
	cli := client.New(client.WithLogLevel(3)).
		OnOrderBook(func(ob *client.PushOrderBook) error {
			log.Printf("%s bid=%d ask=%d", ob.Code, len(ob.Bids), len(ob.Asks))
			return nil
		})
	defer cli.Close()
	_ = cli
}

func ExampleClient_OnTicker() {
	cli := client.New(client.WithLogLevel(3)).
		OnTicker(func(t *client.PushTicker) error {
			log.Printf("%s tick price=%.2f", t.Code, t.Price)
			return nil
		})
	defer cli.Close()
	_ = cli
}

func ExampleClient_OnBroker() {
	cli := client.New(client.WithLogLevel(3)).
		OnBroker(func(b *client.PushBroker) error {
			log.Printf("%s bidBrokers=%d askBrokers=%d", b.Code, len(b.Bids), len(b.Asks))
			return nil
		})
	defer cli.Close()
	_ = cli
}

func ExampleClient_OnTrdNotify() {
	cli := client.New(client.WithLogLevel(3)).
		OnTrdNotify(func(n *client.PushTrdNotify) error {
			log.Printf("Trade notification type=%d", n.Type)
			return nil
		})
	defer cli.Close()
	_ = cli
}

func ExampleParsePushQuote() {
	body := []byte{} // real push body from OpenD
	data, err := client.ParsePushQuote(body)
	if err != nil {
		log.Printf("parse error: %v", err)
	}
	_ = data
}

func ExampleParsePushKLine() {
	body := []byte{} // real push body from OpenD
	data, err := client.ParsePushKLine(body)
	if err != nil {
		log.Printf("parse error: %v", err)
	}
	_ = data
}

func ExampleParsePushOrderBook() {
	body := []byte{} // real push body from OpenD
	data, err := client.ParsePushOrderBook(body)
	if err != nil {
		log.Printf("parse error: %v", err)
	}
	_ = data
}

func ExampleParsePushOrderUpdate() {
	body := []byte{} // real push body from OpenD
	data, err := client.ParsePushOrderUpdate(body)
	if err != nil {
		log.Printf("parse error: %v", err)
	}
	_ = data
}

func ExampleParsePushOrderFill() {
	body := []byte{} // real push body from OpenD
	data, err := client.ParsePushOrderFill(body)
	if err != nil {
		log.Printf("parse error: %v", err)
	}
	_ = data
}

func ExampleClient_FindAccount() {
	cli := client.New(client.WithLogLevel(3))
	defer cli.Close()

	acc := cli.FindAccount([]client.Account{
		{AccID: 1, TrdEnv: int32(client.TrdEnv_Simulate)},
		{AccID: 2, TrdEnv: int32(client.TrdEnv_Real)},
	})
	if acc != nil {
		fmt.Printf("Found account: %d (env=%d)\n", acc.AccID, acc.TrdEnv)
	}
}

func ExampleClient_WithTradeEnv() {
	cli := client.New(client.WithLogLevel(3))
	defer cli.Close()

	realCli := cli.WithTradeEnv(constant.TrdEnv_Real)
	fmt.Println(cli.GetTradeEnv())
	fmt.Println(realCli.GetTradeEnv())
	// Output:
	// TrdEnv_Simulate
	// TrdEnv_Real
}

// Example outputs for common constants.
func Example_constants() {
	fmt.Println(client.Side_Buy)
	fmt.Println(client.Side_Sell)
	fmt.Println(client.SubType_Basic)
	fmt.Println(client.SubType_OrderBook)
	// Output:
	// 1
	// 2
	// 1
	// 2
}

func Example_pushProtoIDs() {
	fmt.Println(client.ProtoID_Qot_UpdateBasicQot)
	fmt.Println(client.ProtoID_Qot_UpdateKL)
	fmt.Println(client.ProtoID_Trd_UpdateOrder)
	fmt.Println(client.ProtoID_Trd_UpdateOrderFill)
	// Output:
	// 3005
	// 3007
	// 2208
	// 2218
}

func Example_contextUsage() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli := client.New(client.WithLogLevel(3))
	defer cli.Close()

	// Pass context to quote API — cancelled contexts return immediately.
	_ = ctx
	select {
	case <-ctx.Done():
		fmt.Println("context cancelled")
	default:
		fmt.Println("context active")
	}
	// Output: context active
}
