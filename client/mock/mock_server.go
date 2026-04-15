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
