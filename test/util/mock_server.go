package testutil

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/getglobalstate"
	"github.com/shing1211/futuapi4go/pkg/pb/getuserinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/initconnect"
	"github.com/shing1211/futuapi4go/pkg/pb/keepalive"
	"google.golang.org/protobuf/proto"
)

// ============================================================================
// MockServer is a test mock for Futu OpenD
// ============================================================================

type MockServer struct {
	listener net.Listener
	addr     string
	t        *testing.T

	// Handler registry: protoID -> handler
	handlers map[uint32]MockHandler
	mu       sync.RWMutex

	// Connection tracking
	conns   map[net.Conn]bool
	connsMu sync.Mutex

	// Request log for assertions
	requests   []MockRequest
	requestsMu sync.Mutex

	// Running state
	running bool
	wg      sync.WaitGroup
}

type MockHandler func(req []byte) ([]byte, error)

type MockRequest struct {
	ProtoID  uint32
	SerialNo uint32
	Body     []byte
	Time     time.Time
}

// NewMockServer creates a new mock server for testing
func NewMockServer(t *testing.T) *MockServer {
	t.Helper()

	s := &MockServer{
		t:        t,
		handlers: make(map[uint32]MockHandler),
		conns:    make(map[net.Conn]bool),
		addr:     "127.0.0.1:0", // random port
	}

	s.registerDefaultHandlers()
	return s
}

// Start starts the mock server
func (s *MockServer) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to start mock server: %w", err)
	}

	s.addr = s.listener.Addr().String()
	s.running = true

	s.wg.Add(1)
	go s.acceptLoop()

	return nil
}

// Stop stops the mock server
func (s *MockServer) Stop() {
	if !s.running {
		return
	}

	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}

	s.connsMu.Lock()
	for conn := range s.conns {
		conn.Close()
	}
	s.connsMu.Unlock()

	s.wg.Wait()
}

// Addr returns the server address
func (s *MockServer) Addr() string {
	return s.addr
}

// RegisterHandler registers a handler for a protoID
func (s *MockServer) RegisterHandler(protoID uint32, handler MockHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[protoID] = handler
}

// GetRequests returns all received requests
func (s *MockServer) GetRequests() []MockRequest {
	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()
	result := make([]MockRequest, len(s.requests))
	copy(result, s.requests)
	return result
}

// ClearRequests clears the request log
func (s *MockServer) ClearRequests() {
	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()
	s.requests = nil
}

// ============================================================================
// Internal methods
// ============================================================================

func (s *MockServer) registerDefaultHandlers() {
	s.RegisterHandler(1001, s.handleInitConnect)
	s.RegisterHandler(1002, s.handleKeepAlive)
	s.RegisterHandler(1004, s.handleGetGlobalState)
	s.RegisterHandler(1005, s.handleGetUserInfo)
}

func (s *MockServer) acceptLoop() {
	defer s.wg.Done()

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			return
		}

		s.connsMu.Lock()
		s.conns[conn] = true
		s.connsMu.Unlock()

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

func (s *MockServer) handleConnection(conn net.Conn) {
	fmt.Printf("[MockServer] handleConnection started\n")
	defer func() {
		fmt.Printf("[MockServer] handleConnection exiting\n")
		conn.Close()
		s.connsMu.Lock()
		delete(s.conns, conn)
		s.connsMu.Unlock()
		s.wg.Done()
	}()

	for s.running {
		// Read header (44 bytes)
		header := make([]byte, 44)
		fmt.Printf("[MockServer] Waiting to read header...\n")
		if _, err := readFull(conn, header); err != nil {
			fmt.Printf("[MockServer] Failed to read header: %v\n", err)
			return
		}
		fmt.Printf("[MockServer] Read header: %v\n", header[:10])

		// Validate magic
		if header[0] != 'F' || header[1] != 'T' {
			s.t.Errorf("Invalid magic bytes: %v", header[:2])
			return
		}

		// Parse header fields
		protoID := readUint32LE(header[2:])
		serialNo := readUint32LE(header[8:])
		bodyLen := readUint32LE(header[12:]) // Fixed: was header[16:]

		fmt.Printf("[MockServer] Parsed header: ProtoID=%d, SerialNo=%d, BodyLen=%d\n", protoID, serialNo, bodyLen)

		// Read body
		body := make([]byte, bodyLen)
		if _, err := readFull(conn, body); err != nil {
			return
		}

		// Log request
		s.requestsMu.Lock()
		s.requests = append(s.requests, MockRequest{
			ProtoID:  protoID,
			SerialNo: serialNo,
			Body:     body,
			Time:     time.Now(),
		})
		s.requestsMu.Unlock()

		// Find handler
		s.mu.RLock()
		handler, ok := s.handlers[protoID]
		s.mu.RUnlock()

		if !ok {
			s.t.Errorf("No handler registered for protoID %d", protoID)
			continue
		}

		// Handle request
		respBody, err := handler(body)
		if err != nil {
			s.t.Errorf("Handler error for protoID %d: %v", protoID, err)
			continue
		}

		// Write response
		if err := s.writeResponse(conn, protoID, serialNo, respBody); err != nil {
			s.t.Errorf("Failed to write response: %v", err)
			return
		}
	}
}

func (s *MockServer) writeResponse(conn net.Conn, protoID, serialNo uint32, body []byte) error {
	header := make([]byte, 44)

	// Magic
	header[0] = 'F'
	header[1] = 'T'

	// ProtoID
	writeUint32LE(header[2:], protoID)

	// ProtoFmt (0 = Protobuf)
	header[6] = 0

	// ProtoVer (0)
	header[7] = 0

	// SerialNo
	writeUint32LE(header[8:], serialNo)

	// BodyLen
	writeUint32LE(header[12:], uint32(len(body))) // Fixed: was header[16:]

	// SHA1 (zeros for mock)
	// Reserved (zeros)

	// Write header
	if _, err := conn.Write(header); err != nil {
		return err
	}

	// Write body
	if len(body) > 0 {
		if _, err := conn.Write(body); err != nil {
			return err
		}
	}

	return nil
}

// ============================================================================
// Default handlers
// ============================================================================

func (s *MockServer) handleInitConnect(req []byte) ([]byte, error) {
	fmt.Printf("[MockServer] handleInitConnect called, req len=%d\n", len(req))
	var reqMsg initconnect.Request
	if err := proto.Unmarshal(req, &reqMsg); err != nil {
		fmt.Printf("[MockServer] Failed to unmarshal: %v\n", err)
		return nil, fmt.Errorf("failed to unmarshal InitConnect request: %w", err)
	}
	fmt.Printf("[MockServer] Unmarshaled successfully, ClientID=%s\n", reqMsg.C2S.GetClientID())

	loginUserID := uint64(123456789)
	serverVer := int32(10100)
	keepAliveInterval := int32(30)
	connID := uint64(1234567890)
	connAESKey := "mock_aes_key_12345"
	retType := int32(0) // Success

	resp := &initconnect.Response{
		RetType: &retType,
		S2C: &initconnect.S2C{
			LoginUserID:       &loginUserID,
			ServerVer:         &serverVer,
			KeepAliveInterval: &keepAliveInterval,
			ConnID:            &connID,
			ConnAESKey:        &connAESKey,
		},
	}

	body, err := proto.Marshal(resp)
	fmt.Printf("[MockServer] Marshaled response, body len=%d\n", len(body))
	return body, err
}

func (s *MockServer) handleKeepAlive(req []byte) ([]byte, error) {
	var reqMsg keepalive.Request
	if err := proto.Unmarshal(req, &reqMsg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal KeepAlive request: %w", err)
	}

	resp := &keepalive.Response{
		S2C: &keepalive.S2C{},
	}

	return proto.Marshal(resp)
}

func (s *MockServer) handleGetGlobalState(req []byte) ([]byte, error) {
	var reqMsg getglobalstate.Request
	if err := proto.Unmarshal(req, &reqMsg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal GetGlobalState request: %w", err)
	}

	connID := uint64(1234567890)
	serverVer := int32(10100)
	serverBuildNo := int32(6208)
	qotLogined := true
	trdLogined := true
	marketState := int32(2) // Normal market state

	resp := &getglobalstate.Response{
		S2C: &getglobalstate.S2C{
			ConnID:        &connID,
			ServerVer:     &serverVer,
			ServerBuildNo: &serverBuildNo,
			QotLogined:    &qotLogined,
			TrdLogined:    &trdLogined,
			MarketHK:      &marketState,
			MarketUS:      &marketState,
			MarketSH:      &marketState,
			MarketSZ:      &marketState,
		},
	}

	return proto.Marshal(resp)
}

func (s *MockServer) handleGetUserInfo(req []byte) ([]byte, error) {
	var reqMsg getuserinfo.Request
	if err := proto.Unmarshal(req, &reqMsg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal GetUserInfo request: %w", err)
	}

	userID := int64(123456789)
	nickName := "TestUser"
	apiLevel := "100"
	hkQotRight := int32(2)
	usQotRight := int32(2)
	cnQotRight := int32(1)

	resp := &getuserinfo.Response{
		S2C: &getuserinfo.S2C{
			UserID:     &userID,
			NickName:   &nickName,
			ApiLevel:   &apiLevel,
			HkQotRight: &hkQotRight,
			UsQotRight: &usQotRight,
			CnQotRight: &cnQotRight,
		},
	}

	return proto.Marshal(resp)
}

// ============================================================================
// Helper functions
// ============================================================================

func readFull(conn net.Conn, buf []byte) (int, error) {
	n := 0
	for n < len(buf) {
		read, err := conn.Read(buf[n:])
		n += read
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func readUint32LE(b []byte) uint32 {
	_ = b[3] // bounds check hint
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func writeUint32LE(b []byte, v uint32) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

// ============================================================================
// Test helpers
// ============================================================================

// NewTestClient creates a client connected to the mock server
func NewTestClient(t *testing.T, server *MockServer) (*futuapi.Client, func()) {
	t.Helper()

	cli := futuapi.New(
		futuapi.WithDialTimeout(5*time.Second),
		futuapi.WithAPITimeout(5*time.Second),
		futuapi.WithKeepAliveInterval(10*time.Second),
		futuapi.WithMaxRetries(0), // disable retries for tests
		futuapi.WithLogLevel(3),   // silent
	)

	err := cli.Connect(server.Addr())
	if err != nil {
		t.Fatalf("Failed to connect to mock server: %v", err)
	}

	cleanup := func() {
		cli.Close()
	}

	return cli, cleanup
}

// AssertProtoID checks if request was sent with correct protoID
func (s *MockServer) AssertProtoID(t *testing.T, expected uint32) {
	t.Helper()

	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()

	if len(s.requests) == 0 {
		t.Fatal("No requests received")
	}

	lastReq := s.requests[len(s.requests)-1]
	if lastReq.ProtoID != expected {
		t.Errorf("Expected protoID %d, got %d", expected, lastReq.ProtoID)
	}
}

// AssertRequestCount checks if expected number of requests were received
func (s *MockServer) AssertRequestCount(t *testing.T, expected int) {
	t.Helper()

	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()

	if len(s.requests) != expected {
		t.Errorf("Expected %d requests, got %d", expected, len(s.requests))
	}
}

// HasProtoID checks if a specific protoID was received
func (s *MockServer) HasProtoID(protoID uint32) bool {
	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()

	for _, req := range s.requests {
		if req.ProtoID == protoID {
			return true
		}
	}
	return false
}
