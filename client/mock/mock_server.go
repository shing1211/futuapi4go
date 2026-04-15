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

package client_test

import (
	"net"
	"sync"
	"testing"
)

// MockServer simulates a simple TCP server to test connection logic.
type MockServer struct {
	Listener net.Listener
	Handler  func(conn net.Conn)
	wg       *sync.WaitGroup
}

func NewMockServer(t *testing.T, handler func(conn net.Conn)) (*MockServer, error) {
	addr := "127.0.0.1:0" // Use port 0 to let the OS assign a free port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		t.Fatalf("failed to set up mock listener: %v", err)
	}

	return &MockServer{
		Listener: listener,
		Handler:  handler,
		wg:       &sync.WaitGroup{},
	}, nil
}

func (m *MockServer) Start(t *testing.T) net.Addr {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		for {
			conn, err := m.Listener.Accept()
			if err != nil {
				// Expected error when listener is closed during shutdown
				return
			}
			t.Logf("MockServer accepted connection: %s", conn.RemoteAddr())
			m.Handler(conn)
		}
	}()
	return m.Listener.Addr()
}

func (m *MockServer) Stop() {
	m.Listener.Close()
	m.wg.Wait() // Wait for the accept loop to finish gracefully
}
