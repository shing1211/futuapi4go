package trd

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

func TestOrderBuilder(t *testing.T) {
	accID := uint64(123456789)
	market := constant.TrdMarket_HK
	env := constant.TrdEnv_Simulate

	t.Run("BasicBuyOrder", func(t *testing.T) {
		req := NewOrder(accID, market, env).
			Buy("00700", 100).
			At(350.5).
			Build()

		if req.AccID != accID {
			t.Errorf("expected AccID %d, got %d", accID, req.AccID)
		}
		if req.TrdMarket != market {
			t.Errorf("expected market %v, got %v", market, req.TrdMarket)
		}
		if req.TrdSide != constant.TrdSide_Buy {
			t.Errorf("expected TrdSide Buy, got %v", req.TrdSide)
		}
		if req.Code != "00700" {
			t.Errorf("expected code 00700, got %s", req.Code)
		}
		if req.Qty != 100 {
			t.Errorf("expected qty 100, got %f", req.Qty)
		}
		if req.Price != 350.5 {
			t.Errorf("expected price 350.5, got %f", req.Price)
		}
		if req.OrderType != constant.OrderType_Normal {
			t.Errorf("expected OrderType Normal, got %v", req.OrderType)
		}
	})

	t.Run("SellOrder", func(t *testing.T) {
		req := NewOrder(accID, market, env).
			Sell("00700", 200).
			At(360.0).
			Build()

		if req.TrdSide != constant.TrdSide_Sell {
			t.Errorf("expected TrdSide Sell, got %v", req.TrdSide)
		}
		if req.Qty != 200 {
			t.Errorf("expected qty 200, got %f", req.Qty)
		}
	})

	t.Run("MarketOrder", func(t *testing.T) {
		req := NewOrder(accID, market, env).
			Buy("00700", 100).
			Market().
			Build()

		if req.OrderType != constant.OrderType_Market {
			t.Errorf("expected OrderType Market, got %v", req.OrderType)
		}
		if req.Price != 0 {
			t.Errorf("expected price 0, got %f", req.Price)
		}
	})

	t.Run("WithRemark", func(t *testing.T) {
		req := NewOrder(accID, market, env).
			Buy("00700", 100).
			At(350.5).
			WithRemark("test order").
			Build()

		if req.Remark != "test order" {
			t.Errorf("expected remark 'test order', got %s", req.Remark)
		}
	})

	t.Run("WithTimeInForce", func(t *testing.T) {
		req := NewOrder(accID, market, env).
			Buy("00700", 100).
			At(350.5).
			WithTimeInForce(constant.TimeInForce_GTC).
			Build()

		if req.TimeInForce != int32(constant.TimeInForce_GTC) {
			t.Errorf("expected TimeInForce GTD, got %d", req.TimeInForce)
		}
	})

	t.Run("WithFillOutsideRTH", func(t *testing.T) {
		req := NewOrder(accID, market, env).
			Buy("00700", 100).
			At(350.5).
			WithFillOutsideRTH(true).
			Build()

		if !req.FillOutsideRTH {
			t.Error("expected FillOutsideRTH true")
		}
	})

	t.Run("WithAuxPrice", func(t *testing.T) {
		req := NewOrder(accID, market, env).
			Buy("00700", 100).
			At(350.5).
			WithAuxPrice(345.0).
			Build()

		if req.AuxPrice != 345.0 {
			t.Errorf("expected auxPrice 345.0, got %f", req.AuxPrice)
		}
	})

	t.Run("FluentInterface", func(t *testing.T) {
		req := NewOrder(accID, market, env).
			Buy("00700", 100).
			At(350.5).
			WithRemark("fluent test").
			WithTimeInForce(constant.TimeInForce_GTC).
			WithFillOutsideRTH(true).
			Build()

		if req.TrdSide != constant.TrdSide_Buy ||
			req.Code != "00700" ||
			req.Qty != 100 ||
			req.Price != 350.5 ||
			req.Remark != "fluent test" ||
			req.TimeInForce != int32(constant.TimeInForce_GTC) ||
			!req.FillOutsideRTH {
			t.Error("fluent interface chain failed")
		}
	})
}