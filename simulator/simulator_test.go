package simulator

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func TestSimulatorIntegration(t *testing.T) {
	srv := New("127.0.0.1:11111")
	srv.RegisterDefaultHandlers()
	srv.RegisterQotHandlers()

	if err := srv.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer srv.Stop()

	if !srv.IsRunning() {
		t.Fatal("Server should be running")
	}

	conn, err := net.DialTimeout("tcp", "127.0.0.1:11111", 5*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	fmt.Println("✓ Connected to simulator")
}

func TestSimulatorWithClient(t *testing.T) {
	srv := New("127.0.0.1:11112")
	srv.RegisterDefaultHandlers()
	srv.RegisterQotHandlers()

	if err := srv.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer srv.Stop()

	time.Sleep(100 * time.Millisecond)

	conn, err := net.DialTimeout("tcp", "127.0.0.1:11112", 5*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	pkt, err := readPacket(conn)
	if err != nil {
		t.Fatalf("Failed to read packet: %v", err)
	}

	if string(pkt.Header.Magic[:]) != "FT" {
		t.Errorf("Expected magic 'FT', got %v", pkt.Header.Magic)
	}

	fmt.Printf("✓ Received packet: ProtoID=%d, SerialNo=%d, BodyLen=%d\n",
		pkt.Header.ProtoID, pkt.Header.SerialNo, pkt.Header.BodyLen)
}

func readPacket(conn net.Conn) (*Packet, error) {
	header := make([]byte, 46)
	if _, err := conn.Read(header); err != nil {
		return nil, err
	}

	var h Header
	h.Magic = [2]byte{header[0], header[1]}
	h.ProtoID = uint32(header[2]) | uint32(header[3])<<8 | uint32(header[4])<<16 | uint32(header[5])<<24
	h.SerialNo = uint32(header[6]) | uint32(header[7])<<8 | uint32(header[8])<<16 | uint32(header[9])<<24
	h.BodyLen = uint32(header[10]) | uint32(header[11])<<8 | uint32(header[12])<<16 | uint32(header[13])<<24

	body := make([]byte, h.BodyLen)
	if h.BodyLen > 0 {
		if _, err := conn.Read(body); err != nil {
			return nil, err
		}
	}

	return &Packet{Header: h, Body: body}, nil
}

func ExampleServer() {
	fmt.Println("Starting Futu OpenD Simulator...")

	srv := New("127.0.0.1:11111")
	srv.RegisterDefaultHandlers()
	srv.RegisterQotHandlers()

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
	defer srv.Stop()

	fmt.Println("Simulator listening on 127.0.0.1:11111")
	fmt.Println("Press Ctrl+C to stop")

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func TestMain(m *testing.M) {
	m.Run()
}
