package trd

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

func captureAuditLogger() (*AuditLogger, *bytes.Buffer) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return NewAuditLogger(logger), &buf
}

func TestAuditLogPlaceOrderSuccess(t *testing.T) {
	audit, buf := captureAuditLogger()

	req := &PlaceOrderRequest{
		AccID:   12345,
		Code:    "US.AAPL",
		TrdSide: constant.TrdSide_Buy,
		Qty:     100,
		Price:   150.0,
	}
	audit.LogPlaceOrder(req, 98765, nil)

	output := buf.String()
	if !strings.Contains(output, "trade operation") {
		t.Error("expected log message 'trade operation'")
	}
	if !strings.Contains(output, "PlaceOrder") {
		t.Error("expected op=PlaceOrder")
	}
	if !strings.Contains(output, "US.AAPL") {
		t.Error("expected code=US.AAPL")
	}
	if !strings.Contains(output, "Buy") {
		t.Error("expected side=Buy")
	}
	if !strings.Contains(output, `"success":true`) {
		t.Error("expected success=true")
	}
	if !strings.Contains(output, `"acc_id":12345`) {
		t.Error("expected acc_id=12345")
	}
	if !strings.Contains(output, `"order_id":98765`) {
		t.Error("expected order_id=98765")
	}
}

func TestAuditLogPlaceOrderError(t *testing.T) {
	audit, buf := captureAuditLogger()

	req := &PlaceOrderRequest{
		AccID:   12345,
		Code:    "US.AAPL",
		TrdSide: constant.TrdSide_Sell,
		Qty:     50,
		Price:   200.0,
	}
	audit.LogPlaceOrder(req, 0, constant.NewFutuError(constant.ErrCodeInsufficientBalance, "PlaceOrder", "insufficient funds"))

	output := buf.String()
	if !strings.Contains(output, `"success":false`) {
		t.Error("expected success=false")
	}
	if !strings.Contains(output, "insufficient funds") {
		t.Error("expected error message in log")
	}
}

func TestAuditLogModifyOrder(t *testing.T) {
	audit, buf := captureAuditLogger()

	req := &ModifyOrderRequest{
		AccID:   12345,
		OrderID: 55555,
	}
	resp := &ModifyOrderResponse{
		OrderID: 55555,
	}
	audit.LogModifyOrder(req, resp, nil)

	output := buf.String()
	if !strings.Contains(output, "ModifyOrder") {
		t.Error("expected op=ModifyOrder")
	}
	if !strings.Contains(output, `"success":true`) {
		t.Error("expected success=true")
	}
}

func TestAuditLogModifyOrderError(t *testing.T) {
	audit, buf := captureAuditLogger()

	req := &ModifyOrderRequest{
		AccID:   12345,
		OrderID: 55555,
	}
	audit.LogModifyOrder(req, nil, constant.NewFutuError(constant.ErrCodeOrderRejected, "ModifyOrder", "order not found"))

	output := buf.String()
	if !strings.Contains(output, "order not found") {
		t.Error("expected error message in log")
	}
	if !strings.Contains(output, `"success":false`) {
		t.Error("expected success=false")
	}
}

func TestAuditLogReconfirmOrder(t *testing.T) {
	audit, buf := captureAuditLogger()

	req := &ReconfirmOrderRequest{
		OrderID: 33333,
	}
	resp := &ReconfirmOrderResponse{
		OrderID: 33333,
	}
	audit.LogReconfirmOrder(req, resp, nil)

	output := buf.String()
	if !strings.Contains(output, "ReconfirmOrder") {
		t.Error("expected op=ReconfirmOrder")
	}
	if !strings.Contains(output, `"success":true`) {
		t.Error("expected success=true")
	}
}
