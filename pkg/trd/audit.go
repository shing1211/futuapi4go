package trd

import (
	"context"
	"log/slog"
	"time"

	"github.com/shing1211/futuapi4go/pkg/constant"
)

type OperationType string

const (
	OpPlaceOrder    OperationType = "PlaceOrder"
	OpModifyOrder   OperationType = "ModifyOrder"
	OpReconfirmOrder OperationType = "ReconfirmOrder"
)

type AuditEntry struct {
	Timestamp time.Time
	Op        OperationType
	AccID     uint64
	Code      string
	Side      string
	Qty       float64
	Price     float64
	OrderID   uint64
	Success   bool
	Error     string
}

type AuditLogger struct {
	logger *slog.Logger
}

func NewAuditLogger(logger *slog.Logger) *AuditLogger {
	if logger == nil {
		logger = slog.Default()
	}
	return &AuditLogger{logger: logger}
}

func (a *AuditLogger) logEntry(entry AuditEntry) {
	level := slog.LevelInfo
	if !entry.Success {
		level = slog.LevelError
	}
	a.logger.LogAttrs(context.Background(), level,
		"trade operation",
		slog.String("op", string(entry.Op)),
		slog.Time("timestamp", entry.Timestamp),
		slog.Uint64("acc_id", entry.AccID),
		slog.String("code", entry.Code),
		slog.String("side", entry.Side),
		slog.Float64("qty", entry.Qty),
		slog.Float64("price", entry.Price),
		slog.Uint64("order_id", entry.OrderID),
		slog.Bool("success", entry.Success),
		slog.String("error", entry.Error),
	)
}

func (a *AuditLogger) LogPlaceOrder(req *PlaceOrderRequest, orderID uint64, err error) {
	entry := AuditEntry{
		Timestamp: time.Now(),
		Op:        OpPlaceOrder,
		AccID:     req.AccID,
		Code:      req.Code,
		Side:      trdSideToString(req.TrdSide),
		Qty:       req.Qty,
		Price:     req.Price,
		OrderID:   orderID,
		Success:   err == nil,
	}
	if err != nil {
		entry.Error = err.Error()
	}
	a.logEntry(entry)
}

func (a *AuditLogger) LogModifyOrder(req *ModifyOrderRequest, resp *ModifyOrderResponse, err error) {
	entry := AuditEntry{
		Timestamp: time.Now(),
		Op:        OpModifyOrder,
		AccID:     req.AccID,
		OrderID:   req.OrderID,
		Success:   err == nil,
	}
	if err != nil {
		entry.Error = err.Error()
	}
	if resp != nil {
		entry.OrderID = resp.OrderID
	}
	a.logEntry(entry)
}

func (a *AuditLogger) LogReconfirmOrder(req *ReconfirmOrderRequest, resp *ReconfirmOrderResponse, err error) {
	op := req.Header
	accID := uint64(0)
	if op != nil {
		accID = op.GetAccID()
	}
	entry := AuditEntry{
		Timestamp: time.Now(),
		Op:        OpReconfirmOrder,
		AccID:     accID,
		OrderID:   req.OrderID,
		Success:   err == nil,
	}
	if err != nil {
		entry.Error = err.Error()
	}
	if resp != nil {
		entry.OrderID = resp.OrderID
	}
	a.logEntry(entry)
}

func trdSideToString(side constant.TrdSide) string {
	switch side {
	case constant.TrdSide_Buy:
		return "Buy"
	case constant.TrdSide_Sell:
		return "Sell"
	default:
		return "Unknown"
	}
}
