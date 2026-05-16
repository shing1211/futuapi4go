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

package futuapi

import (
	"context"
	"crypto/tls"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/shing1211/futuapi4go/pkg/breaker"
	"github.com/shing1211/futuapi4go/pkg/ratelimit"
	"github.com/shing1211/futuapi4go/pkg/retry"
	"google.golang.org/protobuf/proto"

	"github.com/shing1211/futuapi4go/pkg/pb/keepalive"
)

func TestNewClient(t *testing.T) {
	client := New()
	if client == nil {
		t.Fatal("New() returned nil")
	}
	if client.conn == nil {
		t.Error("conn should not be nil")
	}
	if client.handlers == nil {
		t.Error("handlers should not be nil")
	}
	if client.opts == nil {
		t.Error("opts should not be nil")
	}
}

func TestNewWithOptions(t *testing.T) {
	opts := []Option{
		WithMaxRetries(5),
		WithReconnectInterval(10 * time.Second),
		WithDialTimeout(5 * time.Second),
		WithAPITimeout(15 * time.Second),
		WithLogLevel(2),
	}
	client := New(opts...)
	if client == nil {
		t.Fatal("New() returned nil")
	}
	if client.opts.MaxRetries != 5 {
		t.Errorf("expected MaxRetries 5, got %d", client.opts.MaxRetries)
	}
	if client.opts.ReconnectInterval != 10*time.Second {
		t.Errorf("expected ReconnectInterval 10s, got %v", client.opts.ReconnectInterval)
	}
	if client.opts.DialTimeout != 5*time.Second {
		t.Errorf("expected DialTimeout 5s, got %v", client.opts.DialTimeout)
	}
	if client.opts.APITimeout != 15*time.Second {
		t.Errorf("expected APITimeout 15s, got %v", client.opts.APITimeout)
	}
	if client.opts.LogLevel != 2 {
		t.Errorf("expected LogLevel 2, got %d", client.opts.LogLevel)
	}
}

func TestNewOptionsDefaults(t *testing.T) {
	opts := NewOptions()
	if opts.DialTimeout != DefaultDialTimeout {
		t.Errorf("expected DialTimeout %v, got %v", DefaultDialTimeout, opts.DialTimeout)
	}
	if opts.APITimeout != DefaultTimeout {
		t.Errorf("expected APITimeout %v, got %v", DefaultTimeout, opts.APITimeout)
	}
	if opts.KeepAliveInterval != DefaultKeepAliveInterval {
		t.Errorf("expected KeepAliveInterval %v, got %v", DefaultKeepAliveInterval, opts.KeepAliveInterval)
	}
	if opts.MaxRetries != DefaultMaxRetries {
		t.Errorf("expected MaxRetries %d, got %d", DefaultMaxRetries, opts.MaxRetries)
	}
	if opts.ReconnectInterval != DefaultReconnectInterval {
		t.Errorf("expected ReconnectInterval %v, got %v", DefaultReconnectInterval, opts.ReconnectInterval)
	}
	if opts.ReconnectBackoff != 1.5 {
		t.Errorf("expected ReconnectBackoff 1.5, got %f", opts.ReconnectBackoff)
	}
}

func TestDefaultConstants(t *testing.T) {
	if DefaultTimeout != 30*time.Second {
		t.Errorf("expected DefaultTimeout 30s, got %v", DefaultTimeout)
	}
	if DefaultKeepAliveInterval != 30*time.Second {
		t.Errorf("expected DefaultKeepAliveInterval 30s, got %v", DefaultKeepAliveInterval)
	}
	if DefaultMaxRetries != 3 {
		t.Errorf("expected DefaultMaxRetries 3, got %d", DefaultMaxRetries)
	}
	if DefaultReconnectInterval != 3*time.Second {
		t.Errorf("expected DefaultReconnectInterval 3s, got %v", DefaultReconnectInterval)
	}
	if DefaultDialTimeout != 10*time.Second {
		t.Errorf("expected DefaultDialTimeout 10s, got %v", DefaultDialTimeout)
	}
}

func TestEnsureConnectedNotConnected(t *testing.T) {
	client := New()
	defer client.Close()

	err := client.EnsureConnected()
	if err == nil {
		t.Error("EnsureConnected should return error when not connected")
	}
	if err != ErrNotConnected {
		t.Errorf("expected ErrNotConnected, got %v", err)
	}
}

func TestIsConnectedInitialState(t *testing.T) {
	client := New()
	defer client.Close()

	if client.IsConnected() {
		t.Error("client should not be connected initially")
	}
}

func TestContextReturnsContext(t *testing.T) {
	client := New()
	defer client.Close()

	ctx := client.Context()
	if ctx == nil {
		t.Error("Context() returned nil")
	}
}

func TestWithContext(t *testing.T) {
	client := New()
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newClient := client.WithContext(ctx)
	if newClient == nil {
		t.Fatal("WithContext returned nil")
	}
	if newClient.Context() != ctx {
		t.Error("WithContext did not set context correctly")
	}
	// Original client should still be usable
	if client.Context() == ctx {
		t.Error("WithContext should not modify original client's context")
	}
}

func TestSetPushHandler(t *testing.T) {
	client := New()
	defer client.Close()

	client.SetPushHandler(func(pkt *Packet) {
		// Handler registered successfully
	})

	// We can't easily test push handling without a real connection,
	// but we can verify the handler is set without panic
	t.Log("SetPushHandler executed successfully")
}

func TestRegisterHandler(t *testing.T) {
	client := New()
	defer client.Close()

	client.RegisterHandler(9999, func(protoID uint32, body []byte) {
		// Handler registered
	})

	// Verify handler is registered
	client.handlersMu.RLock()
	_, ok := client.handlers[9999]
	client.handlersMu.RUnlock()

	if !ok {
		t.Error("handler not registered")
	}
}

func TestSerialNoIncrement(t *testing.T) {
	client := New()
	defer client.Close()

	first := client.NextSerialNo()
	second := client.NextSerialNo()
	third := client.NextSerialNo()

	if second <= first {
		t.Errorf("serial numbers not incrementing: %d, %d", first, second)
	}
	if third <= second {
		t.Errorf("serial numbers not incrementing: %d, %d", second, third)
	}
}

// ---------------------------------------------------------------------------
// Mock ConnInterface for tests that need to simulate network behavior
// ---------------------------------------------------------------------------

type mockConn struct {
	writePacketFn    func(protoID uint32, serialNo uint32, body []byte) error
	readResponseFn   func(serialNo uint32, timeout time.Duration) (*Packet, error)
	closeFn          func()
	apiTimeout       time.Duration
}

func (m *mockConn) Close() error {
	if m.closeFn != nil {
		m.closeFn()
	}
	return nil
}

func (m *mockConn) WritePacket(protoID uint32, serialNo uint32, body []byte) error {
	if m.writePacketFn != nil {
		return m.writePacketFn(protoID, serialNo, body)
	}
	return nil
}

func (m *mockConn) WritePacketWithSHA1(protoID uint32, serialNo uint32, body []byte, bodySHA1 [20]byte) error {
	return m.WritePacket(protoID, serialNo, body)
}

func (m *mockConn) WritePacketEncrypted(protoID uint32, serialNo uint32, encryptedBody []byte, encryptedBodySHA1 [20]byte) error {
	return m.WritePacket(protoID, serialNo, encryptedBody)
}

func (m *mockConn) ReadResponse(serialNo uint32, timeout time.Duration) (*Packet, error) {
	if m.readResponseFn != nil {
		return m.readResponseFn(serialNo, timeout)
	}
	return &Packet{}, nil
}

func (m *mockConn) ReadResponseContext(ctx context.Context, serialNo uint32, timeout time.Duration) (*Packet, error) {
	return m.ReadResponse(serialNo, timeout)
}

func (m *mockConn) SetPushHandler(handler PacketHandler) {}

func (m *mockConn) Dispatch(pkt *Packet) {}

func (m *mockConn) DrainDispatches() {}

func (m *mockConn) APITimeout() time.Duration { return m.apiTimeout }

func (m *mockConn) SetAPITimeout(d time.Duration) { m.apiTimeout = d }

func (m *mockConn) Dial(addr string) error { return nil }

func (m *mockConn) SetTLSConfig(cfg *tls.Config) {}

func (m *mockConn) readOne() (*Packet, error) { return &Packet{}, nil }

// ---------------------------------------------------------------------------
// Rate limiter tests
// ---------------------------------------------------------------------------

func TestRateLimiterWired(t *testing.T) {
	rl := ratelimit.NewProtoLimiter(1000, 1000, ratelimit.ModeWait)
	client := New(WithRateLimiter(rl))
	defer client.Close()

	if client.rateLimiter != rl {
		t.Error("WithRateLimiter did not set rate limiter on client")
	}
	if client.opts.RateLimiter != rl {
		t.Error("WithRateLimiter did not set rate limiter in options")
	}
}

func TestRateLimiterBlocksWhenExhausted(t *testing.T) {
	// ProtoLimiter with 0 capacity in Reject mode will immediately reject
	rl := ratelimit.NewProtoLimiter(0, 0, ratelimit.ModeReject)
	client := New()
	client.rateLimiter = rl
	defer client.Close()

	atomic.StoreInt32(&client.state, int32(StateConnected))
	client.conn = &mockConn{
		writePacketFn: func(protoID uint32, serialNo uint32, body []byte) error {
			return nil
		},
		readResponseFn: func(serialNo uint32, timeout time.Duration) (*Packet, error) {
			retType := int32(0)
			body, _ := proto.Marshal(&keepalive.Response{RetType: &retType})
			return &Packet{Body: body}, nil
		},
	}

	req := &keepalive.Request{C2S: &keepalive.C2S{Time: proto.Int64(time.Now().Unix())}}
	rsp := &keepalive.Response{}

	err := client.Request(ProtoID_KeepAlive, req, rsp)
	if err == nil {
		t.Fatal("expected rate limit error from exhausted rate limiter")
	}
}

// ---------------------------------------------------------------------------
// Retry tests
// ---------------------------------------------------------------------------

func TestRetryOnTransientFailure(t *testing.T) {
	attempts := 0
	client := New(WithRetryConfig(retry.Config{
		MaxAttempts:    3,
		BaseDelay:      1 * time.Millisecond,
		MaxDelay:       10 * time.Millisecond,
		Jitter:         false,
		IsRecoverable:  func(err error) bool { return errors.Is(err, ErrRequestTimeout) },
	}))
	defer client.Close()

	atomic.StoreInt32(&client.state, int32(StateConnected))
	client.conn = &mockConn{
		writePacketFn: func(protoID uint32, serialNo uint32, body []byte) error {
			attempts++
			if attempts < 3 {
				return ErrRequestTimeout
			}
			return nil
		},
		readResponseFn: func(serialNo uint32, timeout time.Duration) (*Packet, error) {
			retType := int32(0)
			body, _ := proto.Marshal(&keepalive.Response{RetType: &retType})
			return &Packet{Body: body}, nil
		},
	}

	req := &keepalive.Request{C2S: &keepalive.C2S{Time: proto.Int64(time.Now().Unix())}}
	rsp := &keepalive.Response{}

	err := client.Request(ProtoID_KeepAlive, req, rsp)
	if err != nil {
		t.Fatalf("Request failed after retries: %v", err)
	}
	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
}

// ---------------------------------------------------------------------------
// Circuit breaker tests
// ---------------------------------------------------------------------------

func TestBreakerInReconnect(t *testing.T) {
	cb := breaker.New(breaker.WithThreshold(1))
	client := New(WithBreaker(cb))
	defer client.Close()

	cb.RecordFailure()

	client.reconnect()

	// Even when breaker prevents reconnection, the state should be disconnected
	if client.State() != StateDisconnected {
		t.Error("expected client to remain disconnected when breaker is open")
	}
}

func TestReconnectWithBreakerOpen(t *testing.T) {
	cb := breaker.New(breaker.WithThreshold(1))
	client := New(WithBreaker(cb))
	defer client.Close()

	cb.RecordFailure()

	client.reconnect()

	if client.IsConnected() {
		t.Error("expected IsConnected() to be false after reconnect with open breaker")
	}
}

// ---------------------------------------------------------------------------
// State transition tests
// ---------------------------------------------------------------------------

func TestClientStateTransitions(t *testing.T) {
	client := New()
	defer client.Close()

	if client.State() != StateDisconnected {
		t.Errorf("expected StateDisconnected, got %v", client.State())
	}

	client.setState(StateConnected)
	if client.State() != StateConnected {
		t.Errorf("expected StateConnected, got %v", client.State())
	}

	client.setState(StateDisconnected)
	if client.State() != StateDisconnected {
		t.Errorf("expected StateDisconnected, got %v", client.State())
	}
}

func TestOnStateChangeFires(t *testing.T) {
	var transitions []struct{ old, new ConnState }
	client := New(WithOnStateChange(func(old, new ConnState) {
		transitions = append(transitions, struct{ old, new ConnState }{old, new})
	}))
	defer client.Close()

	client.setState(StateConnected)
	client.setState(StateDisconnected)

	if len(transitions) != 2 {
		t.Fatalf("expected 2 transitions, got %d", len(transitions))
	}
	if transitions[0].old != StateDisconnected || transitions[0].new != StateConnected {
		t.Errorf("unexpected first transition: %v -> %v", transitions[0].old, transitions[0].new)
	}
	if transitions[1].old != StateConnected || transitions[1].new != StateDisconnected {
		t.Errorf("unexpected second transition: %v -> %v", transitions[1].old, transitions[1].new)
	}
}

// ---------------------------------------------------------------------------
// Shutdown tests
// ---------------------------------------------------------------------------

func TestShutdownRejectsRequests(t *testing.T) {
	client := New()
	defer client.Close()

	client.setState(StateClosing)

	err := client.requestInternal(ProtoID_KeepAlive, &keepalive.Request{}, &keepalive.Response{})
	if !errors.Is(err, ErrClientClosing) {
		t.Errorf("expected ErrClientClosing, got %v", err)
	}
}

func WithOnStateChange(fn func(oldState, newState ConnState)) Option {
	return func(o *ClientOptions) {
		o.OnStateChange = fn
	}
}
