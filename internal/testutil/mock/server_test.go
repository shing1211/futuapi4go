package mock

import (
	"net"
	"testing"
	"time"

	"github.com/shing1211/futuapi4go/pkg/pb/initconnect"
	"github.com/shing1211/futuapi4go/pkg/pb/keepalive"
	"google.golang.org/protobuf/proto"
)

func TestBasicTCPConnection(t *testing.T) {
	srv := NewMockServer(t)
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	conn, err := net.DialTimeout("tcp", srv.Addr(), 5*time.Second)
	if err != nil {
		t.Fatalf("Dial failed: %v", err)
	}
	conn.Close()
}

func TestInitConnectPlaintext(t *testing.T) {
	srv := NewMockServer(t)
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	cli, cleanup := NewTestClient(t, srv)
	defer cleanup()

	if cli.GetConnID() != 1234567890 {
		t.Errorf("expected ConnID 1234567890, got %d", cli.GetConnID())
	}
	if cli.GetLoginUserID() != 123456789 {
		t.Errorf("expected LoginUserID 123456789, got %d", cli.GetLoginUserID())
	}
	if cli.GetServerVer() != 10100 {
		t.Errorf("expected ServerVer 10100, got %d", cli.GetServerVer())
	}
}

func TestKeepAlive(t *testing.T) {
	srv := NewMockServer(t)
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	cli, cleanup := NewTestClient(t, srv)
	defer cleanup()

	timeNow := time.Now().Unix()
	rsp := &keepalive.Response{}
	if err := cli.Request(1004, &keepalive.Request{C2S: &keepalive.C2S{Time: &timeNow}}, rsp); err != nil {
		t.Fatalf("KeepAlive failed: %v", err)
	}
	if rsp.GetRetType() != 0 {
		t.Errorf("expected RetType 0, got %d", rsp.GetRetType())
	}
}

func TestCustomHandler(t *testing.T) {
	srv := NewMockServer(t)
	called := false
	srv.RegisterHandler(9999, func(req []byte) (proto.Message, error) {
		called = true
		return &initconnect.Response{
			S2C: &initconnect.S2C{},
		}, nil
	})
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	cli, cleanup := NewTestClient(t, srv)
	defer cleanup()
	timeNow := time.Now().Unix()
	rsp := &initconnect.Response{}
	if err := cli.Request(9999, &keepalive.Request{C2S: &keepalive.C2S{Time: &timeNow}}, rsp); err != nil {
		t.Fatalf("Custom handler request failed: %v", err)
	}
	if !called {
		t.Error("custom handler was not called")
	}
}

func TestRequestLogging(t *testing.T) {
	srv := NewMockServer(t)
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	_, cleanup := NewTestClient(t, srv)
	cleanup()
	requests := srv.GetRequests()
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
	if requests[0].ProtoID != 1001 {
		t.Errorf("expected protoID 1001, got %d", requests[0].ProtoID)
	}

	srv.ClearRequests()
	if len(srv.GetRequests()) != 0 {
		t.Error("expected 0 requests after clear")
	}
}

func TestMultipleConnections(t *testing.T) {
	srv := NewMockServer(t)
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	_, cleanup1 := NewTestClient(t, srv)
	defer cleanup1()

	_, cleanup2 := NewTestClient(t, srv)
	defer cleanup2()

	srv.AssertRequestCount(t, 2)
}

func TestRSAEncryptedConnection(t *testing.T) {
	srv := NewMockServer(t)
	srv.WithRSA()
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	cli, cleanup := NewTestClientWithRSA(t, srv)
	defer cleanup()

	if cli.GetConnID() != 1234567890 {
		t.Errorf("expected ConnID 1234567890, got %d", cli.GetConnID())
	}
	if !cli.IsEncrypt() {
		t.Error("expected encryption to be enabled")
	}
}

func TestHasProtoID(t *testing.T) {
	srv := NewMockServer(t)
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	_, cleanup := NewTestClient(t, srv)
	cleanup()
	if !srv.HasProtoID(1001) {
		t.Error("expected HasProtoID(1001) to be true")
	}
	if srv.HasProtoID(9999) {
		t.Error("expected HasProtoID(9999) to be false")
	}
}
