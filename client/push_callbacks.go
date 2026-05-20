package client

import (
	"sync"

	"log/slog"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
)

type callbackState struct {
	mu sync.Mutex

	quote         func(*PushQuote) error
	order         func(*PushOrderUpdate) error
	orderFill     func(*PushOrderFill) error
	kline         func(*PushKLine) error
	orderBook     func(*PushOrderBook) error
	ticker        func(*PushTicker) error
	rt            func(*PushRT) error
	broker        func(*PushBroker) error
	priceReminder func(*PushPriceReminder) error
	trdNotify     func(*PushTrdNotify) error

	set bool
}

func (c *Client) ensurePushHandler() {
	c.cbs.mu.Lock()
	if c.cbs.set {
		c.cbs.mu.Unlock()
		return
	}
	c.cbs.set = true
	c.cbs.mu.Unlock()

	c.inner.SetPushHandler(func(pkt *futuapi.Packet) {
		var err error
		switch pkt.Header.ProtoID {
		case ProtoID_Qot_UpdateBasicQot:
			if c.cbs.quote != nil {
				var data *PushQuote
				data, err = ParsePushQuote(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.quote(data)
				}
			}
		case ProtoID_Qot_UpdateKL:
			if c.cbs.kline != nil {
				var data *PushKLine
				data, err = ParsePushKLine(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.kline(data)
				}
			}
		case ProtoID_Qot_UpdateOrderBook:
			if c.cbs.orderBook != nil {
				var data *PushOrderBook
				data, err = ParsePushOrderBook(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.orderBook(data)
				}
			}
		case ProtoID_Qot_UpdateTicker:
			if c.cbs.ticker != nil {
				var data *PushTicker
				data, err = ParsePushTicker(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.ticker(data)
				}
			}
		case ProtoID_Qot_UpdateRT:
			if c.cbs.rt != nil {
				var data *PushRT
				data, err = ParsePushRT(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.rt(data)
				}
			}
		case ProtoID_Qot_UpdateBroker:
			if c.cbs.broker != nil {
				var data *PushBroker
				data, err = ParsePushBroker(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.broker(data)
				}
			}
		case ProtoID_Qot_UpdatePriceReminder:
			if c.cbs.priceReminder != nil {
				var data *PushPriceReminder
				data, err = ParsePushPriceReminder(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.priceReminder(data)
				}
			}
		case ProtoID_Trd_UpdateOrder:
			if c.cbs.order != nil {
				var data *PushOrderUpdate
				data, err = ParsePushOrderUpdate(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.order(data)
				}
			}
		case ProtoID_Trd_UpdateOrderFill:
			if c.cbs.orderFill != nil {
				var data *PushOrderFill
				data, err = ParsePushOrderFill(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.orderFill(data)
				}
			}
		case ProtoID_Trd_Notify:
			if c.cbs.trdNotify != nil {
				var data *PushTrdNotify
				data, err = ParsePushTrdNotify(pkt.Body)
				if err == nil && data != nil {
					err = c.cbs.trdNotify(data)
				}
			}
		}
		if err != nil {
			slog.Warn("push callback error", "error", err)
		}
	})
}

// OnQuote registers a callback for real-time quote updates (ProtoID 3005).
func (c *Client) OnQuote(fn func(*PushQuote) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.quote = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}

// OnOrder registers a callback for order updates (ProtoID 2208).
func (c *Client) OnOrder(fn func(*PushOrderUpdate) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.order = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}

// OnOrderFill registers a callback for order fill updates (ProtoID 2218).
func (c *Client) OnOrderFill(fn func(*PushOrderFill) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.orderFill = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}

// OnKLine registers a callback for K-line updates (ProtoID 3007).
func (c *Client) OnKLine(fn func(*PushKLine) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.kline = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}

// OnOrderBook registers a callback for order book updates (ProtoID 3013).
func (c *Client) OnOrderBook(fn func(*PushOrderBook) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.orderBook = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}

// OnTicker registers a callback for ticker updates (ProtoID 3011).
func (c *Client) OnTicker(fn func(*PushTicker) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.ticker = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}

// OnRT registers a callback for real-time data updates (ProtoID 3009).
func (c *Client) OnRT(fn func(*PushRT) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.rt = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}

// OnBroker registers a callback for broker queue updates (ProtoID 3015).
func (c *Client) OnBroker(fn func(*PushBroker) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.broker = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}

// OnPriceReminder registers a callback for price reminder updates (ProtoID 3019).
func (c *Client) OnPriceReminder(fn func(*PushPriceReminder) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.priceReminder = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}

// OnTrdNotify registers a callback for trade notifications (ProtoID 2207).
func (c *Client) OnTrdNotify(fn func(*PushTrdNotify) error) *Client {
	c.cbs.mu.Lock()
	c.cbs.trdNotify = fn
	c.cbs.mu.Unlock()
	c.ensurePushHandler()
	return c
}
