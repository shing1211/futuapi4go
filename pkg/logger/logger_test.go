package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLevelString(t *testing.T) {
	tests := []struct {
		lvl Level
		want string
	}{
		{LevelDebug, "DEBUG"},
		{LevelInfo, "INFO"},
		{LevelWarn, "WARN"},
		{LevelError, "ERROR"},
		{Level(99), "UNKNOWN"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.lvl.String(); got != tt.want {
				t.Errorf("Level(%d).String() = %q, want %q", tt.lvl, got, tt.want)
			}
		})
	}
}

func TestLoggerLevels(t *testing.T) {
	var buf bytes.Buffer
	l := New(WithOutput(&buf), WithFormat(FormatText), WithLevel(LevelDebug))

	l.Debug("debug msg")
	l.Info("info msg")
	l.Warn("warn msg")
	l.Error("error msg")

	output := buf.String()
	if !strings.Contains(output, "DEBUG") || !strings.Contains(output, "info msg") {
		t.Errorf("expected DEBUG and info msg in output, got: %s", output)
	}
}

func TestLoggerLevelFilter(t *testing.T) {
	var buf bytes.Buffer
	l := New(WithOutput(&buf), WithFormat(FormatText), WithLevel(LevelWarn))

	l.Debug("debug msg")
	l.Info("info msg")
	l.Warn("warn msg")
	l.Error("error msg")

	output := buf.String()
	if strings.Contains(output, "debug") || strings.Contains(output, "info") {
		t.Errorf("expected debug/info filtered out, got: %s", output)
	}
	if !strings.Contains(output, "WARN") || !strings.Contains(output, "ERROR") {
		t.Errorf("expected warn/error in output, got: %s", output)
	}
}

func TestLoggerJSONFormat(t *testing.T) {
	var buf bytes.Buffer
	l := New(WithOutput(&buf), WithFormat(FormatJSON), WithLevel(LevelInfo))

	l.Info("test message", "key", "value")

	output := buf.String()
	if !strings.Contains(output, `"level":"INFO"`) {
		t.Errorf("expected JSON level field, got: %s", output)
	}
	if !strings.Contains(output, `"msg":"test message"`) {
		t.Errorf("expected JSON msg field, got: %s", output)
	}
}

func TestLoggerFields(t *testing.T) {
	var buf bytes.Buffer
	l := New(WithOutput(&buf), WithFormat(FormatText), WithLevel(LevelInfo))

	l.Info("connection established", "addr", "127.0.0.1:11111", "client_id", 42)

	output := buf.String()
	if !strings.Contains(output, "addr=127.0.0.1:11111") {
		t.Errorf("expected addr field in output, got: %s", output)
	}
	if !strings.Contains(output, "client_id=42") {
		t.Errorf("expected client_id field in output, got: %s", output)
	}
}

func TestLoggerSetLevel(t *testing.T) {
	var buf bytes.Buffer
	l := New(WithOutput(&buf), WithFormat(FormatText), WithLevel(LevelError))

	l.Info("should not appear")
	if buf.Len() > 0 {
		t.Errorf("expected no output at Error level, got: %s", buf.String())
	}

	l.SetLevel(LevelInfo)
	l.Info("should appear")
	if buf.Len() == 0 {
		t.Error("expected output after SetLevel")
	}
}

func TestLoggerSetOutput(t *testing.T) {
	var buf1, buf2 bytes.Buffer
	l := New(WithOutput(&buf1), WithFormat(FormatText), WithLevel(LevelInfo))

	l.Info("before switch")
	if buf1.Len() == 0 {
		t.Error("expected output to buf1")
	}

	l.SetOutput(&buf2)
	l.Info("after switch")
	if buf1.Len() > 0 && buf2.Len() == 0 {
		t.Error("expected output to buf2, not buf1")
	}
}

func TestLoggerInstance(t *testing.T) {
	var buf bytes.Buffer
	l := New(WithOutput(&buf), WithFormat(FormatText), WithLevel(LevelInfo))

	l.Info("instance test")
	if !bytes.Contains(buf.Bytes(), []byte("instance test")) {
		t.Errorf("expected instance test in output, got: %s", buf.String())
	}
}

func TestLoggerSetFormat(t *testing.T) {
	var buf bytes.Buffer
	l := New(WithOutput(&buf), WithFormat(FormatText), WithLevel(LevelInfo))

	l.Info("text mode")
	textOutput := buf.String()
	buf.Reset()

	l.SetFormat(FormatJSON)
	l.Info("json mode")
	jsonOutput := buf.String()

	if !strings.Contains(textOutput, "INFO text mode") {
		t.Errorf("expected text output, got: %s", textOutput)
	}
	if !strings.Contains(jsonOutput, `"level":"INFO"`) {
		t.Errorf("expected JSON output, got: %s", jsonOutput)
	}
}
