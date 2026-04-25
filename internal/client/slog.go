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

package futuapi

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"
)

type Level int

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type SlogLogger struct {
	logger *slog.Logger
	level  Level
}

func NewSlogLogger(w io.Writer, level Level) *SlogLogger {
	var opts *slog.HandlerOptions
	if level != LevelInfo {
		opts = &slog.HandlerOptions{
			Level: slog.Level(level),
		}
	}

	l := slog.New(slog.NewJSONHandler(w, opts))
	return &SlogLogger{
		logger: l,
		level:  level,
	}
}

func NewSlogLoggerDefault(level Level) *SlogLogger {
	return NewSlogLogger(os.Stderr, level)
}

func (s *SlogLogger) Debug(ctx context.Context, msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s *SlogLogger) Info(ctx context.Context, msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s *SlogLogger) Warn(ctx context.Context, msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

func (s *SlogLogger) Error(ctx context.Context, msg string, args ...any) {
	s.logger.Error(msg, args...)
}

func (s *SlogLogger) Log(ctx context.Context, lvl Level, msg string, args ...any) {
	s.logger.Log(ctx, slog.Level(lvl), msg, args...)
}

type SlogHandler struct {
	logger *slog.Logger
}

func NewSlogHandler(w io.Writer) *SlogHandler {
	return &SlogHandler{
		logger: slog.New(slog.NewJSONHandler(w, nil)),
	}
}

func (h *SlogHandler) Handle(ctx context.Context, level slog.Level, msg string, attrs ...any) {
	h.logger.Log(ctx, level, msg, attrs...)
}

func WithSlogDefault() func(*ClientOptions) {
	return func(opts *ClientOptions) {
		opts.SlogLogger = NewSlogLoggerDefault(LevelInfo)
	}
}

func WithSlogLevel(level Level) func(*ClientOptions) {
	return func(opts *ClientOptions) {
		opts.SlogLogger = NewSlogLoggerDefault(level)
	}
}

func WithSlogWriter(w io.Writer, level Level) func(*ClientOptions) {
	return func(opts *ClientOptions) {
		opts.SlogLogger = NewSlogLogger(w, level)
	}
}

type SlogMetrics struct {
	logger     *slog.Logger
	attrAttrs  []slog.Attr
	startTime time.Time
}

func NewSlogMetrics(w io.Writer) *SlogMetrics {
	logger := slog.New(slog.NewJSONHandler(w, nil))
	return &SlogMetrics{
		logger:    logger,
		startTime: time.Now(),
	}
}

func (m *SlogMetrics) LogRequest(ctx context.Context, protoID uint32, serialNo uint64, latency time.Duration, success bool) {
	m.logger.Info("request completed",
		slog.Uint64("proto_id", uint64(protoID)),
		slog.Uint64("serial_no", serialNo),
		slog.Duration("latency_ms", latency),
		slog.Bool("success", success),
	)
}

func (m *SlogMetrics) LogConnection(connID uint64, event string) {
	m.logger.Info("connection event",
		slog.Uint64("conn_id", connID),
		slog.String("event", event),
	)
}

func (m *SlogMetrics) LogReconnect(count int, reason string) {
	m.logger.Warn("reconnection",
		slog.Int("count", count),
		slog.String("reason", reason),
	)
}

func (m *SlogMetrics) LogError(err error, ctx string) {
	m.logger.Error("error",
		slog.String("error", err.Error()),
		slog.String("context", ctx),
	)
}

type SlogAttributes struct {
	Type    string
	ConnID  uint64
	UserID  uint64
	ProtoID uint32
}

func (a *SlogAttributes) ToAttrs() []any {
	attrs := []any{
		"type", a.Type,
	}
	if a.ConnID > 0 {
		attrs = append(attrs, "conn_id", a.ConnID)
	}
	if a.UserID > 0 {
		attrs = append(attrs, "user_id", a.UserID)
	}
	if a.ProtoID > 0 {
		attrs = append(attrs, "proto_id", a.ProtoID)
	}
	return attrs
}