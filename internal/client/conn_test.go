package futuapi

import (
	"net"
	"testing"
	"time"
)

func TestConnNilConnection(t *testing.T) {
	conn := NewConn(nil)
	if conn == nil {
		t.Fatal("NewConn(nil) should return non-nil conn wrapper")
	}

	_, err := conn.readOne()
	if err == nil {
		t.Error("readOne should return error when conn is nil")
	}

	_, err = conn.ReadResponse(1, time.Second)
	if err == nil {
		t.Error("ReadResponse should return error when conn is nil")
	}

	err = conn.WritePacket(1001, 1, []byte{})
	if err == nil {
		t.Error("WritePacket should return error when conn is nil")
	}

	err = conn.SetReadDeadline(time.Now())
	if err == nil {
		t.Error("SetReadDeadline should return error when conn is nil")
	}
}

func TestConnSetPushHandler(t *testing.T) {
	conn := NewConn(nil)
	var called bool
	conn.SetPushHandler(func(pkt *Packet) {
		called = true
	})
	if called {
		t.Error("handler should not be called during SetPushHandler")
	}
}

func TestConnDispatchToChannel(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	conn := NewConn(nil)
	go func() {
		c, _ := l.Accept()
		conn.conn = c
	}()

	c, err := net.Dial("tcp", l.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	pkt := &Packet{Header: Header{SerialNo: 42}, Body: []byte("hello")}

	dispatched := make(chan *Packet, 1)
	conn.SetPushHandler(func(p *Packet) {
		select {
		case dispatched <- p:
		default:
		}
	})

	conn.Dispatch(pkt)

	select {
	case p := <-dispatched:
		if p.Header.SerialNo != 42 {
			t.Errorf("expected serial 42, got %d", p.Header.SerialNo)
		}
		if string(p.Body) != "hello" {
			t.Errorf("expected body 'hello', got '%s'", string(p.Body))
		}
	case <-time.After(time.Second):
		t.Error("packet was not dispatched within 1s")
	}
}

func TestConnDispatchToWaitingReader(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	done := make(chan struct{})
	var serverConn net.Conn
	go func() {
		c, _ := l.Accept()
		serverConn = c
		<-done
	}()

	conn := NewConn(nil)
	c, err := net.Dial("tcp", l.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	conn.conn = c
	defer func() {
		c.Close()
		serverConn.Close()
		close(done)
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		conn.Dispatch(&Packet{Header: Header{SerialNo: 99}, Body: []byte("response")})
	}()

	pkt, err := conn.ReadResponse(99, 5*time.Second)
	if err != nil {
		t.Fatalf("ReadResponse failed: %v", err)
	}
	if pkt.Header.SerialNo != 99 {
		t.Errorf("expected serial 99, got %d", pkt.Header.SerialNo)
	}
}

func TestConnReadResponseTimeout(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	done := make(chan struct{})
	go func() {
		c, _ := l.Accept()
		defer c.Close()
		<-done
	}()

	conn := NewConn(nil)
	c, err := net.Dial("tcp", l.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	conn.conn = c
	defer func() {
		c.Close()
		close(done)
	}()

	_, err = conn.ReadResponse(1, 100*time.Millisecond)
	if err == nil {
		t.Error("expected timeout error")
	}
}

