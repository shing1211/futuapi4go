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
