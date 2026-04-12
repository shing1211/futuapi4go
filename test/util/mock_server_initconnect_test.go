package testutil

import (
	"testing"
)

func TestMockServerInitConnect(t *testing.T) {
	server := NewMockServer(t)

	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer server.Stop()

	// Use NewTestClient which properly connects through the client
	cli, cleanup := NewTestClient(t, server)
	defer cleanup()

	// If we got here without error, InitConnect worked
	t.Log("✓ InitConnect successful!")

	// Verify server received the request
	server.AssertProtoID(t, 1001)
	server.AssertRequestCount(t, 1)

	// Verify client state
	if !cli.IsConnected() {
		t.Error("Client should be connected")
	}

	if cli.GetConnID() != 1234567890 {
		t.Errorf("Expected ConnID 1234567890, got %d", cli.GetConnID())
	}

	if cli.GetServerVer() != 10100 {
		t.Errorf("Expected ServerVer 10100, got %d", cli.GetServerVer())
	}
}
