package futuapi

import (
	"bytes"
	"strings"
	"testing"
)

func TestSlogLoggerIntegration(t *testing.T) {
	var buf bytes.Buffer
	sl := NewSlogLogger(&buf, LevelDebug)
	client := New(WithSlog(sl), WithLogLevel(LogLevelInfo))

	client.logInfo("test message %s", "hello")

	if buf.Len() == 0 {
		t.Error("expected slog output, got empty")
	}

	output := buf.String()
	if !strings.Contains(output, `"msg":"test message hello"`) {
		t.Errorf("expected log message in JSON output, got: %s", output)
	}
}

func TestSlogLoggerWarn(t *testing.T) {
	var buf bytes.Buffer
	sl := NewSlogLogger(&buf, LevelDebug)
	client := New(WithSlog(sl), WithLogLevel(LogLevelInfo))

	client.logWarn("warning: %s", "something happened")

	if buf.Len() == 0 {
		t.Error("expected slog output, got empty")
	}

	output := buf.String()
	if !strings.Contains(output, `"msg":"warning: something happened"`) {
		t.Errorf("expected warning message in JSON output, got: %s", output)
	}
}

func TestSlogLoggerError(t *testing.T) {
	var buf bytes.Buffer
	sl := NewSlogLogger(&buf, LevelDebug)
	client := New(WithSlog(sl), WithLogLevel(LogLevelInfo))

	client.logError("error: %s", "something failed")

	if buf.Len() == 0 {
		t.Error("expected slog output, got empty")
	}

	output := buf.String()
	if !strings.Contains(output, `"msg":"error: something failed"`) {
		t.Errorf("expected error message in JSON output, got: %s", output)
	}
}

func TestSlogLoggerLevelSuppression(t *testing.T) {
	var buf bytes.Buffer
	sl := NewSlogLogger(&buf, LevelDebug)
	client := New(WithSlog(sl), WithLogLevel(LogLevelError))

	client.logInfo("should be suppressed")

	if buf.Len() != 0 {
		t.Error("expected no output when log level suppresses info")
	}
}

func TestSlogLoggerDefaultNoSlog(t *testing.T) {
	client := New(WithLogLevel(LogLevelInfo))
	client.opts.Logger = nil

	client.logInfo("fallback message %d", 42)
}
