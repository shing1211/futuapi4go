package simulator

import (
	"net"
	"testing"
	"time"
)

func TestServerStart(t *testing.T) {
	srv := New("127.0.0.1:11115")
	srv.RegisterDefaultHandlers()
	srv.RegisterQotHandlers()

	if err := srv.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer srv.Stop()

	if !srv.IsRunning() {
		t.Fatal("Server should be running")
	}

	conn, err := net.DialTimeout("tcp", "127.0.0.1:11115", 2*time.Second)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	conn.Close()

	t.Log("Server started and accepted connection")
}
