package testutil

import (
	"net"
	"testing"
	"time"
)

func TestMockServerBasicConnection(t *testing.T) {
	server := NewMockServer(t)

	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer server.Stop()

	// Give server time to start
	time.Sleep(100 * time.Millisecond)

	// Try to connect
	conn, err := net.DialTimeout("tcp", server.Addr(), 2*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect to mock server: %v", err)
	}
	defer conn.Close()

	t.Logf("Successfully connected to mock server at %s", server.Addr())
}
