// Package logger provides structured, leveled logging for the futuapi4go SDK.
//
// Unlike the basic log.Default() used elsewhere in the SDK, this package
// offers structured output with severity levels and contextual fields.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/logger"
//
//	log := logger.New(logger.LevelInfo,
//	    logger.WithOutput(os.Stdout),
//	    logger.WithFormat(logger.FormatJSON),
//	)
//	log.Info("connected to OpenD", "addr", "127.0.0.1:11111")
//	log.Warn("rate limit approaching", "requests", 45, "limit", 50)
//
// The logger is safe for concurrent use.
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
package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Level int

const (
	LevelDebug Level = 0
	LevelInfo  Level = 1
	LevelWarn  Level = 2
	LevelError Level = 3
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

type Format int

const (
	FormatText Format = 0
	FormatJSON Format = 1
)

type Logger struct {
	mu     sync.Mutex
	level  Level
	out    io.Writer
	format Format
}

type Config struct {
	Level  Level
	Output io.Writer
	Format Format
}

type Option func(*Config)

func WithLevel(lvl Level) Option {
	return func(c *Config) { c.Level = lvl }
}

func WithOutput(w io.Writer) Option {
	return func(c *Config) { c.Output = w }
}

func WithFormat(fmt Format) Option {
	return func(c *Config) { c.Format = fmt }
}

func New(opts ...Option) *Logger {
	cfg := Config{
		Level:  LevelInfo,
		Output: os.Stdout,
		Format: FormatText,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return &Logger{
		level:  cfg.Level,
		out:    cfg.Output,
		format: cfg.Format,
	}
}

func (l *Logger) SetLevel(lvl Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = lvl
}

func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}

func (l *Logger) log(lvl Level, msg string, fields ...interface{}) {
	if lvl < l.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	ts := time.Now().UTC().Format(time.RFC3339)

	switch l.format {
	case FormatJSON:
		entry := logEntry{
			Time:   ts,
			Level:  lvl.String(),
			Msg:    msg,
			Fields: make(map[string]string),
		}
		for i := 0; i < len(fields)-1; i += 2 {
			if key, ok := fields[i].(string); ok {
				entry.Fields[key] = fmt.Sprint(fields[i+1])
			}
		}
		data, _ := json.Marshal(entry)
		fmt.Fprintln(l.out, string(data))
	default:
		line := fmt.Sprintf("%s %s %s", ts, lvl.String(), msg)
		if len(fields) > 0 {
			pairs := make([]string, 0, len(fields)/2)
			for i := 0; i < len(fields)-1; i += 2 {
				if key, ok := fields[i].(string); ok {
					pairs = append(pairs, fmt.Sprintf("%s=%v", key, fields[i+1]))
				}
			}
			if len(pairs) > 0 {
				line += " " + joinFields(pairs)
			}
		}
		fmt.Fprintln(l.out, line)
	}
}

func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.log(LevelDebug, msg, fields...)
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	l.log(LevelInfo, msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.log(LevelWarn, msg, fields...)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	l.log(LevelError, msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.log(LevelError, msg, fields...)
	os.Exit(1)
}

type logEntry struct {
	Time   string            `json:"time"`
	Level  string            `json:"level"`
	Msg    string            `json:"msg"`
	Fields map[string]string `json:"fields,omitempty"`
}

func joinFields(pairs []string) string {
	result := "{"
	for i, p := range pairs {
		if i > 0 {
			result += " "
		}
		result += p
	}
	result += "}"
	return result
}

func (l *Logger) SetFormat(fmt Format) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.format = fmt
}

var (
	Default = New()

	Debug = Default.Debug
	Info  = Default.Info
	Warn  = Default.Warn
	Error = Default.Error
)

func SetLevel(lvl Level)    { Default.SetLevel(lvl) }
func SetOutput(w io.Writer) { Default.SetOutput(w) }
func SetFormat(fmt Format)  { Default.SetFormat(fmt) }
