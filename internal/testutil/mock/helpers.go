package mock

import (
	"testing"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
)

// NewTestClient creates a client connected to the mock server (plaintext mode).
func NewTestClient(t *testing.T, srv *MockServer) (*futuapi.Client, func()) {
	t.Helper()
	cli := futuapi.New(
		futuapi.WithDialTimeout(5*time.Second),
		futuapi.WithAPITimeout(30*time.Second),
		futuapi.WithKeepAliveInterval(10*time.Second),
		futuapi.WithMaxRetries(0),
		futuapi.WithLogLevel(3),
	)
	err := cli.Connect(srv.Addr())
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	return cli, func() { cli.Close() }
}

// NewTestClientWithRSA creates a client connected to the mock server in
// RSA+AES encrypted mode. The mock server must have been created with WithRSA().
func NewTestClientWithRSA(t *testing.T, srv *MockServer) (*futuapi.Client, func()) {
	t.Helper()
	if srv.PublicKeyPEM() == "" {
		t.Fatal("NewTestClientWithRSA requires a mock server created with WithRSA()")
	}
	cli := futuapi.New(
		futuapi.WithDialTimeout(5*time.Second),
		futuapi.WithAPITimeout(30*time.Second),
		futuapi.WithKeepAliveInterval(10*time.Second),
		futuapi.WithMaxRetries(0),
		futuapi.WithLogLevel(3),
		futuapi.WithRSAPrivateKey(srv.PrivateKeyPEM()),
		futuapi.WithEncryption(true),
	)
	err := cli.ConnectWithRSA(srv.Addr(), srv.PublicKeyPEM())
	if err != nil {
		t.Fatalf("ConnectWithRSA failed: %v", err)
	}
	return cli, func() { cli.Close() }
}

// AssertRequestCount checks that the server received exactly n requests.
func (s *MockServer) AssertRequestCount(t *testing.T, n int) {
	t.Helper()
	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()
	if got := len(s.requests); got != n {
		t.Errorf("expected %d requests, got %d", n, got)
	}
}

// AssertProtoID checks that the most recent request has the given ProtoID.
func (s *MockServer) AssertProtoID(t *testing.T, expected uint32) {
	t.Helper()
	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()
	if len(s.requests) == 0 {
		t.Fatal("no requests received")
	}
	last := s.requests[len(s.requests)-1]
	if last.ProtoID != expected {
		t.Errorf("expected protoID %d, got %d", expected, last.ProtoID)
	}
}

// HasProtoID checks if any request with the given ProtoID was received.
func (s *MockServer) HasProtoID(protoID uint32) bool {
	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()
	for _, r := range s.requests {
		if r.ProtoID == protoID {
			return true
		}
	}
	return false
}
