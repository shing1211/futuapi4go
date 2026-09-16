package futuapi

import (
	"bytes"
	"log"
	"log/slog"
	"strings"
	"testing"
)

// Per-packet transport logging is opt-in: it stays silent at Info but emits at
// Debug, so a normal deployment is not flooded by one line per packet.
func TestLogDebugLevelGating(t *testing.T) {
	var buf bytes.Buffer

	atInfo := New(WithLogger(log.New(&buf, "", 0)), WithLogLevel(LogLevelInfo))
	atInfo.logDebug("recv %s (%d) serial=%d bytes=%d", "Trd_GetFunds", 2101, 1, 19)
	if buf.Len() != 0 {
		t.Errorf("debug should be suppressed at LogLevelInfo, got %q", buf.String())
	}

	atDebug := New(WithLogger(log.New(&buf, "", 0)), WithLogLevel(LogLevelDebug))
	atDebug.logDebug("recv %s (%d) serial=%d bytes=%d", "Trd_GetFunds", 2101, 1, 19)
	out := buf.String()
	for _, want := range []string{"recv", "Trd_GetFunds", "2101"} {
		if !strings.Contains(out, want) {
			t.Errorf("expected %q in debug output, got %q", want, out)
		}
	}
}

// A caller-supplied slog.Logger receives SDK events with their level intact.
func TestLogDebugRoutesToSlogLogger(t *testing.T) {
	var buf bytes.Buffer
	l := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	c := New(WithSlogLogger(l), WithLogLevel(LogLevelDebug))
	c.logDebug("recv %s (%d)", "Trd_GetFunds", 2101)

	out := buf.String()
	for _, want := range []string{`"level":"DEBUG"`, "Trd_GetFunds"} {
		if !strings.Contains(out, want) {
			t.Errorf("expected %q in slog output, got %q", want, out)
		}
	}
}

func TestWithSlogLoggerIgnoresNil(t *testing.T) {
	c := New(WithSlogLogger(nil))
	if c.opts.SlogLogger != nil {
		t.Error("WithSlogLogger(nil) should leave the structured logger unset")
	}
}
