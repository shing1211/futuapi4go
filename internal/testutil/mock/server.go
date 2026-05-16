package mock

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"net"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/getglobalstate"
	"github.com/shing1211/futuapi4go/pkg/pb/getuserinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/initconnect"
	"github.com/shing1211/futuapi4go/pkg/pb/keepalive"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetbroker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitaldistribution"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetcapitalflow"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetrt"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesttradedate"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsub"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetacclist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetfunds"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderfilllist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetorderlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdgetpositionlist"
	"github.com/shing1211/futuapi4go/pkg/pb/trdmodifyorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdplaceorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdunlocktrade"
	"google.golang.org/protobuf/proto"
)

type MockHandler func(req []byte) (proto.Message, error)

type MockRequest struct {
	ProtoID    uint32
	SerialNo   uint32
	Body       []byte
	HeaderSHA1 [20]byte
	Time       time.Time
}

type sha1Validation struct {
	ProtoID           uint32
	SerialNo          uint32
	HeaderSHA1        [20]byte
	CiphertextSHA1    [20]byte
	PlaintextSHA1     [20]byte
	MatchedCiphertext bool
	MatchedPlaintext  bool
}

type connState struct {
	aesKey       []byte
	connID       uint64
	encryptEnabled bool
}

type MockServer struct {
	listener net.Listener
	addr     string
	t        *testing.T

	handlers map[uint32]MockHandler
	mu       sync.RWMutex

	conns   map[net.Conn]*connState
	connsMu sync.Mutex

	requests   []MockRequest
	requestsMu sync.Mutex

	sha1Results   []sha1Validation
	sha1ResultsMu sync.Mutex

	running     int32
	wg          sync.WaitGroup
	privKeyPEM  string
	pubKeyPEM   string
	StrictSHA1  bool
}

func NewMockServer(t *testing.T) *MockServer {
	t.Helper()
	s := &MockServer{
		t:        t,
		handlers: make(map[uint32]MockHandler),
		conns:    make(map[net.Conn]*connState),
		addr:     "127.0.0.1:0",
	}
	s.registerDefaultHandlers()
	return s
}

func (s *MockServer) WithRSA() *MockServer {
	priv, pub, err := futuapi.GenerateRSAKeys(1024)
	if err != nil {
		s.t.Fatalf("Failed to generate RSA keys: %v", err)
	}
	s.privKeyPEM = priv
	s.pubKeyPEM = pub
	return s
}

func (s *MockServer) PublicKeyPEM() string {
	return s.pubKeyPEM
}

func (s *MockServer) PrivateKeyPEM() string {
	return s.privKeyPEM
}

func (s *MockServer) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to start mock server: %w", err)
	}
	s.addr = s.listener.Addr().String()
	atomic.StoreInt32(&s.running, 1)
	s.wg.Add(1)
	go s.acceptLoop()
	return nil
}

func (s *MockServer) Stop() {
	if atomic.LoadInt32(&s.running) == 0 {
		return
	}
	atomic.StoreInt32(&s.running, 0)
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

func (s *MockServer) Addr() string { return s.addr }

func (s *MockServer) RegisterHandler(protoID uint32, handler MockHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[protoID] = handler
}

func (s *MockServer) GetRequests() []MockRequest {
	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()
	result := make([]MockRequest, len(s.requests))
	copy(result, s.requests)
	return result
}

func (s *MockServer) ClearRequests() {
	s.requestsMu.Lock()
	defer s.requestsMu.Unlock()
	s.requests = nil
}

func (s *MockServer) registerDefaultHandlers() {
	s.RegisterHandler(1001, s.handleInitConnect)
	s.RegisterHandler(1004, s.handleKeepAlive)
	s.RegisterHandler(1002, s.handleGetGlobalState)
	s.RegisterHandler(1005, s.handleGetUserInfo)
}

func (s *MockServer) acceptLoop() {
	defer s.wg.Done()
	for atomic.LoadInt32(&s.running) != 0 {
		conn, err := s.listener.Accept()
		if err != nil {
			return
		}
		s.connsMu.Lock()
		s.conns[conn] = nil
		s.connsMu.Unlock()
		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

func (s *MockServer) handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		s.connsMu.Lock()
		delete(s.conns, conn)
		s.connsMu.Unlock()
		s.wg.Done()
	}()

	for atomic.LoadInt32(&s.running) != 0 {
		protoID, serialNo, headerSHA1, body, err := s.readPacket(conn)
		if err != nil {
			return
		}

		s.requestsMu.Lock()
		s.requests = append(s.requests, MockRequest{
			ProtoID:    protoID,
			SerialNo:   serialNo,
			Body:       body,
			HeaderSHA1: headerSHA1,
			Time:       time.Now(),
		})
		s.requestsMu.Unlock()

		// Decrypt body if needed (after InitConnect handshake)
		cs := s.getConnState(conn)
		plainBody := body
		if protoID == 1001 && s.privKeyPEM != "" {
			// InitConnect may be RSA-encrypted
			decrypted, err := futuapi.RSADecrypt(s.privKeyPEM, body)
			if err == nil {
				plainBody = decrypted
			}
		} else if cs != nil && cs.encryptEnabled {
			if protoID != 1001 {
				decrypted, err := futuapi.AESDecrypt(cs.aesKey, body)
				if err != nil {
					s.t.Errorf("AES decrypt failed: %v", err)
					return
				}
				plainBody = decrypted
			}
		}

		// Validate SHA1 for non-InitConnect requests
		if protoID != 1001 {
			result := sha1Validation{
				ProtoID:        protoID,
				SerialNo:       serialNo,
				HeaderSHA1:     headerSHA1,
				CiphertextSHA1: sha1.Sum(body),
				PlaintextSHA1:  sha1.Sum(plainBody),
			}
			result.MatchedCiphertext = result.HeaderSHA1 == result.CiphertextSHA1
			result.MatchedPlaintext = result.HeaderSHA1 == result.PlaintextSHA1
			s.sha1ResultsMu.Lock()
			s.sha1Results = append(s.sha1Results, result)
			s.sha1ResultsMu.Unlock()

			s.t.Logf("SHA1 check protoID=%d serialNo=%d: header=%x ciphertextMatch=%v plaintextMatch=%v",
				protoID, serialNo, headerSHA1[:8], result.MatchedCiphertext, result.MatchedPlaintext)

			if s.StrictSHA1 && !result.MatchedPlaintext {
				s.t.Logf("protoID=%d: strict SHA1 check FAILED (closing connection)", protoID)
				return
			}
		}

		// Handle InitConnect specially — establish AES key
		if protoID == 1001 {
			if err := s.handleInitConnectRequest(conn, serialNo, plainBody); err != nil {
				s.t.Errorf("InitConnect failed: %v", err)
				return
			}
			continue
		}

		s.mu.RLock()
		handler, ok := s.handlers[protoID]
		s.mu.RUnlock()
		if !ok {
			s.t.Errorf("No handler registered for protoID %d", protoID)
			continue
		}

		respMsg, err := handler(plainBody)
		if err != nil {
			s.t.Errorf("Handler error for protoID %d: %v", protoID, err)
			continue
		}
		respMsg = s.fixupResponse(protoID, respMsg)

		respBody, err := proto.Marshal(respMsg)
		if err != nil {
			s.t.Errorf("Failed to marshal response: %v", err)
			continue
		}

		// Encrypt response if needed
		outBody := respBody
		if cs != nil && cs.encryptEnabled {
			encrypted, err := futuapi.AESEncrypt(cs.aesKey, respBody)
			if err != nil {
				s.t.Errorf("AES encrypt failed: %v", err)
				continue
			}
			outBody = encrypted
		}

		if err := s.writePacket(conn, protoID, serialNo, outBody); err != nil {
			s.t.Errorf("Failed to write response: %v", err)
			return
		}
	}
}

func (s *MockServer) getConnState(conn net.Conn) *connState {
	s.connsMu.Lock()
	defer s.connsMu.Unlock()
	return s.conns[conn]
}

func (s *MockServer) setConnState(conn net.Conn, cs *connState) {
	s.connsMu.Lock()
	defer s.connsMu.Unlock()
	s.conns[conn] = cs
}

func (s *MockServer) handleInitConnectRequest(conn net.Conn, serialNo uint32, plainBody []byte) error {
	var req initconnect.Request
	if err := proto.Unmarshal(plainBody, &req); err != nil {
		return fmt.Errorf("unmarshal InitConnect: %w", err)
	}

	connID := uint64(1234567890)
	loginUserID := uint64(123456789)
	serverVer := int32(10100)
	keepAliveInterval := int32(30)
	retType := int32(0)

	// Generate AES key
	aesKey := make([]byte, 16)
	if _, err := rand.Read(aesKey); err != nil {
		return fmt.Errorf("generate AES key: %w", err)
	}

	// Build InitConnect response
	resp := &initconnect.Response{
		RetType: &retType,
		S2C: &initconnect.S2C{
			ConnID:            &connID,
			LoginUserID:       &loginUserID,
			ServerVer:         &serverVer,
			KeepAliveInterval: &keepAliveInterval,
			ConnAESKey:        strPtr(string(aesKey)),
		},
	}
	respBody, err := proto.Marshal(resp)
	if err != nil {
		return fmt.Errorf("marshal InitConnect response: %w", err)
	}

	// Encrypt response with RSA if in encryption mode
	outBody := respBody
	if s.privKeyPEM != "" {
		// Use the public key to encrypt (same key pair)
		encrypted, err := futuapi.RSAEncrypt(s.pubKeyPEM, respBody)
		if err != nil {
			return fmt.Errorf("RSA encrypt InitConnect response: %w", err)
		}
		outBody = encrypted
	}

	if err := s.writePacket(conn, 1001, serialNo, outBody); err != nil {
		return fmt.Errorf("write InitConnect response: %w", err)
	}

	// Determine if encryption is enabled based on client request
	encryptEnabled := false
	if req.GetC2S().GetPacketEncAlgo() != -1 {
		encryptEnabled = true
	}

	// Store AES key for this connection
	s.setConnState(conn, &connState{
		aesKey:         aesKey,
		connID:         connID,
		encryptEnabled: encryptEnabled,
	})
	return nil
}

func (s *MockServer) handleInitConnect(req []byte) (proto.Message, error) {
	return &initconnect.Response{
		S2C: &initconnect.S2C{},
	}, nil
}

func (s *MockServer) handleKeepAlive(req []byte) (proto.Message, error) {
	return &keepalive.Response{S2C: &keepalive.S2C{}}, nil
}

func (s *MockServer) handleGetGlobalState(req []byte) (proto.Message, error) {
	connID := uint64(1234567890)
	serverVer := int32(10100)
	serverBuildNo := int32(6208)
	qotLogined := true
	trdLogined := true
	marketState := int32(2)

	return &getglobalstate.Response{
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
	}, nil
}

func (s *MockServer) handleGetUserInfo(req []byte) (proto.Message, error) {
	userID := int64(123456789)
	nickName := "TestUser"
	apiLevel := "100"
	hkQotRight := int32(2)
	usQotRight := int32(2)
	cnQotRight := int32(1)

	return &getuserinfo.Response{
		S2C: &getuserinfo.S2C{
			UserID:     &userID,
			NickName:   &nickName,
			ApiLevel:   &apiLevel,
			HkQotRight: &hkQotRight,
			UsQotRight: &usQotRight,
			CnQotRight: &cnQotRight,
		},
	}, nil
}

func (s *MockServer) readPacket(conn net.Conn) (protoID uint32, serialNo uint32, bodySHA1 [20]byte, body []byte, err error) {
	header := make([]byte, 44)
	if _, err = readFull(conn, header); err != nil {
		return 0, 0, [20]byte{}, nil, err
	}
	if header[0] != 'F' || header[1] != 'T' {
		return 0, 0, [20]byte{}, nil, fmt.Errorf("invalid magic: %x %x", header[0], header[1])
	}
	protoID = readUint32LE(header[2:])
	serialNo = readUint32LE(header[8:])
	bodyLen := readUint32LE(header[12:])
	copy(bodySHA1[:], header[16:36])

	body = make([]byte, bodyLen)
	if _, err = readFull(conn, body); err != nil {
		return 0, 0, [20]byte{}, nil, err
	}
	return protoID, serialNo, bodySHA1, body, nil
}

func (s *MockServer) writePacket(conn net.Conn, protoID, serialNo uint32, body []byte) error {
	header := make([]byte, 44)
	header[0] = 'F'
	header[1] = 'T'
	writeUint32LE(header[2:], protoID)
	header[6] = 0
	header[7] = 0
	writeUint32LE(header[8:], serialNo)
	writeUint32LE(header[12:], uint32(len(body)))
	sha1Hash := sha1.Sum(body)
	copy(header[16:36], sha1Hash[:])

	if _, err := conn.Write(header); err != nil {
		return err
	}
	if len(body) > 0 {
		if _, err := conn.Write(body); err != nil {
			return err
		}
	}
	return nil
}

func (s *MockServer) fixupResponse(protoID uint32, msg proto.Message) proto.Message {
	respMap := map[uint32]proto.Message{
		3004: &qotgetbasicqot.Response{},
		3006: &qotgetkl.Response{},
		3008: &qotgetrt.Response{},
		3010: &qotgetticker.Response{},
		3012: &qotgetorderbook.Response{},
		3014: &qotgetbroker.Response{},
		3211: &qotgetcapitalflow.Response{},
		3212: &qotgetcapitaldistribution.Response{},
		3219: &qotrequesttradedate.Response{},
		3001: &qotsub.Response{},
		2001: &trdgetacclist.Response{},
		2005: &trdunlocktrade.Response{},
		2101: &trdgetfunds.Response{},
		2102: &trdgetpositionlist.Response{},
		2201: &trdgetorderlist.Response{},
		2202: &trdplaceorder.Response{},
		2205: &trdmodifyorder.Response{},
		2211: &trdgetorderfilllist.Response{},
		1002: &getglobalstate.Response{},
		1004: &keepalive.Response{},
		1005: &getuserinfo.Response{},
	}
	template, ok := respMap[protoID]
	if !ok {
		fillNilPointers(msg)
		return msg
	}
	proto.Merge(template, msg)
	fillNilPointers(template)
	return template
}

func fillNilPointers(msg proto.Message) {
	if msg == nil {
		return
	}
	v := reflect.ValueOf(msg).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				f.Set(reflect.New(f.Type().Elem()))
			}
			if m, ok := f.Interface().(proto.Message); ok {
				fillNilPointers(m)
			}
		} else if f.Kind() == reflect.Slice {
			for j := 0; j < f.Len(); j++ {
				elem := f.Index(j)
				if elem.Kind() == reflect.Ptr && !elem.IsNil() {
					if m, ok := elem.Interface().(proto.Message); ok {
						fillNilPointers(m)
					}
				}
			}
		}
	}
}

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
	_ = b[3]
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func writeUint32LE(b []byte, v uint32) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func strPtr(s string) *string { return &s }
