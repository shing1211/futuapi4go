package futuapi

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/futuopen/ftapi4go/pb/common"
	"github.com/futuopen/ftapi4go/pb/initconnect"
)

var (
	ErrNotConnected = errors.New("not connected")
)

const (
	ProtoID_InitConnect = 1001
)

type Client struct {
	conn       *Conn
	mu         sync.RWMutex
	connID     uint64
	aesKey     string
	serialNo   uint32
	serialMu   sync.Mutex
	handlers   map[uint32]Handler
	handlersMu sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
}

type Handler func(protoID uint32, body []byte)

func New() *Client {
	ctx, cancel := context.WithCancel(context.Background())
	return &Client{
		conn:     NewConn(nil),
		handlers: make(map[uint32]Handler),
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (c *Client) Connect(addr string) error {
	if err := c.conn.Dial(addr); err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	clientVer := int32(10100)
	clientID := "futuapi4go"
	recvNotify := true
	packetEncAlgo := int32(-1)

	req := &initconnect.C2S{
		ClientVer:     &clientVer,
		ClientID:      &clientID,
		RecvNotify:    &recvNotify,
		PacketEncAlgo: &packetEncAlgo,
	}

	pkt := &initconnect.Request{
		C2S: req,
	}

	body, err := proto.Marshal(pkt)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	serialNo := c.nextSerialNo()
	if err := c.conn.WritePacket(ProtoID_InitConnect, serialNo, body); err != nil {
		return fmt.Errorf("write packet: %w", err)
	}

	c.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	respPkt, err := c.conn.ReadPacket()
	if err != nil {
		c.conn.Close()
		return fmt.Errorf("read response: %w", err)
	}

	var rsp initconnect.Response
	if err := proto.Unmarshal(respPkt.Body, &rsp); err != nil {
		c.conn.Close()
		return fmt.Errorf("unmarshal response: %w", err)
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		c.conn.Close()
		return fmt.Errorf("init connect failed: retType=%d, retMsg=%s", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		c.conn.Close()
		return errors.New("init connect: s2c is nil")
	}

	c.mu.Lock()
	c.connID = s2c.GetConnID()
	c.aesKey = s2c.GetConnAESKey()
	c.mu.Unlock()

	go c.readLoop()

	return nil
}

func (c *Client) nextSerialNo() uint32 {
	c.serialMu.Lock()
	c.serialNo++
	no := c.serialNo
	c.serialMu.Unlock()
	return no
}

func (c *Client) readLoop() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		pkt, err := c.conn.ReadPacket()
		if err != nil {
			fmt.Printf("read error: %v\n", err)
			continue
		}

		c.handlersMu.RLock()
		handler, ok := c.handlers[pkt.Header.ProtoID]
		c.handlersMu.RUnlock()

		if ok {
			handler(pkt.Header.ProtoID, pkt.Body)
		}
	}
}

func (c *Client) RegisterHandler(protoID uint32, handler Handler) {
	c.handlersMu.Lock()
	c.handlers[protoID] = handler
	c.handlersMu.Unlock()
}

func (c *Client) Close() error {
	c.cancel()
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) GetConnID() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.connID
}

func (c *Client) GetAESKey() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.aesKey
}

func (c *Client) Conn() *Conn {
	return c.conn
}

func (c *Client) NextSerialNo() uint32 {
	return c.nextSerialNo()
}

func (c *Client) Request(protoID uint32, req proto.Message, rsp proto.Message) error {
	if c.conn == nil {
		return ErrNotConnected
	}

	body, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	serialNo := c.nextSerialNo()
	if err := c.conn.WritePacket(protoID, serialNo, body); err != nil {
		return err
	}

	c.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	pkt, err := c.conn.ReadPacket()
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if err := proto.Unmarshal(pkt.Body, rsp); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}
