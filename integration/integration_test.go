//go:build integration

package integration

import (
	"os"
	"testing"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
)

func getOpenDAddr() string {
	addr := os.Getenv("FUTU_OPEND_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}
	return addr
}

func skipWithoutOpenD(t *testing.T) {
	if os.Getenv("FUTU_OPEND_ADDR") == "" {
		t.Skip("Set FUTU_OPEND_ADDR to run integration tests")
	}
}

func TestConnectAndDisconnect(t *testing.T) {
	skipWithoutOpenD(t)
	c := futuapi.New()
	defer c.Close()

	if err := c.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	if !c.IsConnected() {
		t.Fatal("Expected connected")
	}
}

func TestGetGlobalState(t *testing.T) {
	skipWithoutOpenD(t)
	c := futuapi.New(futuapi.WithAPITimeout(5 * time.Second))
	defer c.Close()

	if err := c.Connect(getOpenDAddr()); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
}
