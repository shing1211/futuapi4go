package simulator

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"

	"google.golang.org/protobuf/proto"
)

const (
	Magic      = 0x4654 // "FT" in big-endian
	HeaderLen  = 14
	MaxBodyLen = 10 * 1024 * 1024 // 10MB max
)

type Packet struct {
	Magic    uint16
	ProtoID  uint32
	SerialNo uint32
	BodyLen  uint32
	Body     []byte
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
}

func New(addr string) *Server {
	return &Server{
		addr:      addr,
		handlers:  make(map[uint32]Handler),
		closeChan: make(chan struct{}),
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
				fmt.Printf("accept error: %v\n", err)
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

	for {
		pkt, err := s.readPacket(conn)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("read error: %v\n", err)
			}
			return
		}

		s.mu.RLock()
		handler, ok := s.handlers[pkt.ProtoID]
		s.mu.RUnlock()

		var resp *Packet
		if ok {
			resp, err = handler(pkt)
		} else {
			resp = s.errorResponse(pkt, fmt.Errorf("unsupported protoID: %d", pkt.ProtoID))
		}

		if err != nil {
			resp = s.errorResponse(pkt, err)
		}

		if err := s.writePacket(conn, resp); err != nil {
			fmt.Printf("write error: %v\n", err)
			return
		}
	}
}

func (s *Server) readPacket(conn net.Conn) (*Packet, error) {
	header := make([]byte, HeaderLen)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, err
	}

	magic := binary.BigEndian.Uint16(header[0:2])
	if magic != Magic {
		return nil, fmt.Errorf("invalid magic: 0x%04x", magic)
	}

	protoID := binary.BigEndian.Uint32(header[2:6])
	serialNo := binary.BigEndian.Uint32(header[6:10])
	bodyLen := binary.BigEndian.Uint32(header[10:14])

	if bodyLen > MaxBodyLen {
		return nil, fmt.Errorf("body too large: %d", bodyLen)
	}

	body := make([]byte, bodyLen)
	if bodyLen > 0 {
		if _, err := io.ReadFull(conn, body); err != nil {
			return nil, err
		}
	}

	return &Packet{
		Magic:    magic,
		ProtoID:  protoID,
		SerialNo: serialNo,
		BodyLen:  bodyLen,
		Body:     body,
	}, nil
}

func (s *Server) writePacket(conn net.Conn, pkt *Packet) error {
	header := make([]byte, HeaderLen)
	binary.BigEndian.PutUint16(header[0:2], pkt.Magic)
	binary.BigEndian.PutUint32(header[2:6], pkt.ProtoID)
	binary.BigEndian.PutUint32(header[6:10], pkt.SerialNo)
	binary.BigEndian.PutUint32(header[10:14], pkt.BodyLen)

	_, err := conn.Write(append(header, pkt.Body...))
	return err
}

func (s *Server) errorResponse(req *Packet, err error) *Packet {
	return &Packet{
		Magic:    Magic,
		ProtoID:  req.ProtoID,
		SerialNo: req.SerialNo,
		BodyLen:  0,
		Body:     nil,
	}
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
