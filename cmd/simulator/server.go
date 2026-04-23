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

package simulator

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/initconnect"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
)

const (
	HeaderLen     = 44
	MagicBytes    = "FT"
	ProtoVersion  = 0
	MaxPacketSize = 10 * 1024 * 1024
)

type Header struct {
	Magic    [2]byte
	ProtoID  uint32
	ProtoFmt byte
	ProtoVer byte
	SerialNo uint32
	BodyLen  uint32
	BodySHA1 [20]byte
	Reserved [8]byte
}

type Packet struct {
	Header Header
	Body   []byte
}

type Handler func(*Packet) (*Packet, error)

type Server struct {
	addr      string
	listener  net.Listener
	handlers  map[uint32]Handler
	mu        sync.RWMutex
	running   bool
	closeChan chan struct{}
	wg        sync.WaitGroup

	Securities map[string]*qotcommon.Security
	Quotes     map[string]*qotcommon.BasicQot
	Orders     map[uint64]*trdcommon.Order
	Positions  map[string]*trdcommon.Position
}

func New(addr string) *Server {
	return &Server{
		addr:      addr,
		handlers:  make(map[uint32]Handler),
		closeChan: make(chan struct{}),

		Securities: make(map[string]*qotcommon.Security),
		Quotes:     make(map[string]*qotcommon.BasicQot),
		Orders:     make(map[uint64]*trdcommon.Order),
		Positions:  make(map[string]*trdcommon.Position),
	}
}

func (s *Server) RegisterHandler(protoID uint32, handler Handler) {
	s.mu.Lock()
	s.handlers[protoID] = handler
	s.mu.Unlock()
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	s.listener = ln
	s.running = true

	s.wg.Add(1)
	go s.acceptLoop()

	return nil
}

func (s *Server) acceptLoop() {
	defer s.wg.Done()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.closeChan:
				return
			default:
				log.Printf("accept error: %v\n", err)
				continue
			}
		}

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer conn.Close()

	log.Printf("New connection from %v\n", conn.RemoteAddr())

	for {
		pkt, err := s.readPacket(conn)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Printf("read error: %v\n", err)
			return
		}

		s.mu.RLock()
		handler, ok := s.handlers[pkt.Header.ProtoID]
		s.mu.RUnlock()

		var resp *Packet
		if !ok {
			resp, _ = s.errorResponse(pkt, fmt.Errorf("unsupported protoID: %d", pkt.Header.ProtoID))
		} else {
			resp, err = handler(pkt)
			if err != nil {
				resp, _ = s.errorResponse(pkt, err)
			}
		}

		if err := s.writePacket(conn, resp); err != nil {
			log.Printf("write error: %v\n", err)
			return
		}
	}
}

func (s *Server) readPacket(conn net.Conn) (*Packet, error) {
	header := make([]byte, HeaderLen)
	n, err := io.ReadFull(conn, header)
	if err != nil {
		return nil, fmt.Errorf("read header (n=%d): %w", n, err)
	}

	var h Header
	copy(h.Magic[:], header[0:2])
	h.ProtoID = binary.LittleEndian.Uint32(header[2:6])
	h.ProtoFmt = header[6]
	h.ProtoVer = header[7]
	h.SerialNo = binary.LittleEndian.Uint32(header[8:12])
	h.BodyLen = binary.LittleEndian.Uint32(header[12:16])
	copy(h.BodySHA1[:], header[16:36])
	copy(h.Reserved[:], header[36:44])

	if string(h.Magic[:]) != MagicBytes {
		return nil, fmt.Errorf("invalid magic bytes")
	}

	if h.BodyLen > MaxPacketSize {
		return nil, fmt.Errorf("body too large: %d", h.BodyLen)
	}

	body := make([]byte, h.BodyLen)
	if h.BodyLen > 0 {
		if _, err := io.ReadFull(conn, body); err != nil {
			return nil, fmt.Errorf("read body: %w", err)
		}
	}

	return &Packet{Header: h, Body: body}, nil
}

func (s *Server) writePacket(conn net.Conn, pkt *Packet) error {
	header := make([]byte, HeaderLen)
	header[0] = 'F'
	header[1] = 'T'
	binary.LittleEndian.PutUint32(header[2:], pkt.Header.ProtoID)
	header[6] = 0 // ProtoFmt = Protobuf
	header[7] = ProtoVersion
	binary.LittleEndian.PutUint32(header[8:], pkt.Header.SerialNo)
	binary.LittleEndian.PutUint32(header[12:], pkt.Header.BodyLen)
	sha1Hash := sha1.Sum(pkt.Body)
	copy(header[16:36], sha1Hash[:])

	if _, err := conn.Write(header); err != nil {
		return fmt.Errorf("write header: %w", err)
	}

	if len(pkt.Body) > 0 {
		if _, err := conn.Write(pkt.Body); err != nil {
			return fmt.Errorf("write body: %w", err)
		}
	}

	return nil
}

func (s *Server) errorResponse(req *Packet, err error) (*Packet, error) {
	errMsg := err.Error()
	retType := int32(common.RetType_RetType_Failed)
	ret := &initconnect.Response{
		RetType: &retType,
		RetMsg:  &errMsg,
	}

	body, err := proto.Marshal(ret)
	if err != nil {
		return nil, err
	}

	return &Packet{
		Header: Header{
			Magic:    [2]byte{'F', 'T'},
			ProtoID:  req.Header.ProtoID,
			ProtoFmt: 0, // Protobuf
			ProtoVer: ProtoVersion,
			SerialNo: req.Header.SerialNo,
			BodyLen:  uint32(len(body)),
		},
		Body: body,
	}, nil
}

func (s *Server) successResponse(req *Packet, ret proto.Message) (*Packet, error) {
	body, err := proto.Marshal(ret)
	if err != nil {
		return nil, fmt.Errorf("marshal response: %w", err)
	}

	return &Packet{
		Header: Header{
			Magic:    [2]byte{'F', 'T'},
			ProtoID:  req.Header.ProtoID,
			ProtoFmt: 0, // Protobuf
			ProtoVer: ProtoVersion,
			SerialNo: req.Header.SerialNo,
			BodyLen:  uint32(len(body)),
		},
		Body: body,
	}, nil
}

func (s *Server) Stop() {
	close(s.closeChan)
	if s.listener != nil {
		s.listener.Close()
	}
	s.wg.Wait()
	s.running = false
}

func (s *Server) IsRunning() bool {
	return s.running
}

func DecodeRequest(body []byte, msg proto.Message) error {
	return proto.Unmarshal(body, msg)
}

func EncodeResponse(msg proto.Message) ([]byte, error) {
	return proto.Marshal(msg)
}

func NowTimestamp() float64 {
	return float64(time.Now().Unix())
}

func (s *Server) AddSecurity(market int32, code string) {
	key := fmt.Sprintf("%d.%s", market, code)
	s.Securities[key] = &qotcommon.Security{Market: &market, Code: &code}
	isSuspended := false
	priceSpread := 0.0
	turnoverRate := 0.01
	amplitude := 2.0
	s.Quotes[key] = &qotcommon.BasicQot{
		Security:        &qotcommon.Security{Market: &market, Code: &code},
		CurPrice:        func() *float64 { p := 100.0; return &p }(),
		OpenPrice:       func() *float64 { p := 99.0; return &p }(),
		HighPrice:       func() *float64 { p := 102.0; return &p }(),
		LowPrice:        func() *float64 { p := 98.0; return &p }(),
		LastClosePrice:  func() *float64 { p := 100.0; return &p }(),
		Volume:          func() *int64 { v := int64(1000000); return &v }(),
		Turnover:        func() *float64 { t := 100000000.0; return &t }(),
		UpdateTime:      func() *string { t := "2026-04-07 10:30:00"; return &t }(),
		IsSuspended:     &isSuspended,
		ListTime:        func() *string { t := "2020-01-01"; return &t }(),
		PriceSpread:     &priceSpread,
		TurnoverRate:    &turnoverRate,
		Amplitude:       &amplitude,
	}
}

func (s *Server) AddOrder(order *trdcommon.Order) {
	s.Orders[order.GetOrderID()] = order
}

func (s *Server) GetQuote(market int32, code string) *qotcommon.BasicQot {
	key := fmt.Sprintf("%d.%s", market, code)
	return s.Quotes[key]
}
