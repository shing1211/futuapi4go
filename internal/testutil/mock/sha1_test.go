package mock

import (
	"crypto/sha1"
	"encoding/binary"
	"net"
	"testing"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/getglobalstate"
	"github.com/shing1211/futuapi4go/pkg/pb/getuserinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/initconnect"
	"google.golang.org/protobuf/proto"
)

// TestSDKSendsCiphertextSHA1 verifies that the SDK sends SHA1(ciphertext)
// in packet headers for encrypted API requests. This confirms the current
// SDK behavior.
//
// Empirical result: OpenD is lenient and accepts both SHA1(plaintext) and
// SHA1(ciphertext). The official Python SDK
// (futu/common/utils.py:_joint_head()) uses SHA1(plaintext) per the spec.
// We use SHA1(ciphertext) for historical compatibility since v0.5.15.
// See WritePacketEncrypted comments in conn.go and ws.go for details.
func TestSDKSendsCiphertextSHA1(t *testing.T) {
	srv := NewMockServer(t)
	srv.WithRSA()
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	cli, cleanup := NewTestClientWithRSA(t, srv)
	defer cleanup()

	if !cli.IsEncrypt() {
		t.Fatal("expected encryption to be enabled")
	}

	// Make an explicit encrypted API call to trigger SHA1 verification
	var rsp getglobalstate.Response
	if err := cli.Request(1002,
		&getglobalstate.Request{C2S: &getglobalstate.C2S{UserID: proto.Uint64(0)}},
		&rsp); err != nil {
		t.Fatalf("GetGlobalState request failed: %v", err)
	}
	if rsp.GetRetType() != 0 {
		t.Errorf("expected RetType 0, got %d", rsp.GetRetType())
	}

	time.Sleep(100 * time.Millisecond)
	results := srv.GetSHA1Results()
	if len(results) == 0 {
		t.Fatal("no SHA1 results recorded")
	}

	for _, r := range results {
		t.Logf("protoID=%d serialNo=%d: matchedCiphertext=%v matchedPlaintext=%v",
			r.ProtoID, r.SerialNo, r.MatchedCiphertext, r.MatchedPlaintext)

		if !r.MatchedCiphertext {
			t.Errorf("protoID=%d: expected SHA1(ciphertext) match, got header=%x ciphertext=%x plaintext=%x",
				r.ProtoID, r.HeaderSHA1[:8], r.CiphertextSHA1[:8], r.PlaintextSHA1[:8])
		}
	}
}

// TestSHA1PlaintextWouldWork proves that SHA1(plaintext) is accepted by
// a spec-compliant server. It crafts a raw FT packet with SHA1(plaintext)
// on an encrypted connection and sends it to the mock server in strict mode.
func TestSHA1PlaintextWouldWork(t *testing.T) {
	srv := NewMockServer(t)
	srv.WithRSA()
	srv.StrictSHA1 = true
	if err := srv.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer srv.Stop()

	// Raw TCP connection
	conn, err := net.DialTimeout("tcp", srv.Addr(), 5*time.Second)
	if err != nil {
		t.Fatalf("Dial failed: %v", err)
	}
	defer conn.Close()

	// --- Step 1: InitConnect with encryption request ---
	clientVer := int32(601)
	clientID := "sha1-test"
	recvNotify := true
	packetEncAlgo := int32(1) // FTAES_ECB — request encryption
	c2s := &initconnect.C2S{
		ClientVer:     &clientVer,
		ClientID:      &clientID,
		RecvNotify:    &recvNotify,
		PacketEncAlgo: &packetEncAlgo,
	}
	initBody, err := proto.Marshal(&initconnect.Request{C2S: c2s})
	if err != nil {
		t.Fatalf("marshal InitConnect: %v", err)
	}

	// Send InitConnect with SHA1(plaintext) — body is plaintext here
	sendPacket(t, conn, 1001, 1, initBody, false, [20]byte{})

	// Read InitConnect response
	respBody := readResponse(t, conn)
	// RSA-decrypt the response body
	if len(respBody) < 4 {
		t.Fatalf("response too short: %d bytes", len(respBody))
	}
	decrypted, err := futuapi.RSADecrypt(srv.PrivateKeyPEM(), respBody)
	if err != nil {
		t.Fatalf("RSA decrypt InitConnect response: %v", err)
	}
	var initResp initconnect.Response
	if err := proto.Unmarshal(decrypted, &initResp); err != nil {
		t.Fatalf("unmarshal InitConnect response: %v", err)
	}

	aesKey := []byte(initResp.GetS2C().GetConnAESKey())
	t.Logf("InitConnect OK: connID=%d aesKey=%x", initResp.GetS2C().GetConnID(), aesKey)

	// --- Step 2: Send GetGlobalState with SHA1(plaintext) ---
	gBody, err := proto.Marshal(&getglobalstate.Request{C2S: &getglobalstate.C2S{UserID: proto.Uint64(0)}})
	if err != nil {
		t.Fatalf("marshal getglobalstate: %v", err)
	}

	// Encrypt the body
	encBody, err := futuapi.AESEncrypt(aesKey, gBody)
	if err != nil {
		t.Fatalf("AES encrypt: %v", err)
	}

	// Compute SHA1 of plaintext (the spec-compliant way)
	sendPacket(t, conn, 1002, 2, encBody, true, sha1.Sum(gBody))

	// Verify we get a response
	respBody2 := readResponse(t, conn)
	decrypted2, err := futuapi.AESDecrypt(aesKey, respBody2)
	if err != nil {
		t.Fatalf("AES decrypt response: %v", err)
	}
	var gResp getglobalstate.Response
	if err := proto.Unmarshal(decrypted2, &gResp); err != nil {
		t.Fatalf("unmarshal GetGlobalState response: %v", err)
	}
	if gResp.GetRetType() != 0 {
		t.Errorf("expected RetType 0, got %d", gResp.GetRetType())
	}

	// --- Step 3: Also try GetUserInfo with SHA1(ciphertext) to contrast ---
	uBody, err := proto.Marshal(&getuserinfo.Request{C2S: &getuserinfo.C2S{}})
	if err != nil {
		t.Fatalf("marshal getuserinfo: %v", err)
	}
	uEncBody, err := futuapi.AESEncrypt(aesKey, uBody)
	if err != nil {
		t.Fatalf("AES encrypt userinfo: %v", err)
	}
	// Send with SHA1(ciphertext) — this should be REJECTED by strict mode
	sendPacket(t, conn, 1005, 3, uEncBody, false, sha1.Sum(uEncBody))

	// Read response — should fail (connection closed by strict server)
	_, err = readResponseRaw(t, conn)
	if err != nil {
		t.Logf("SHA1(ciphertext) correctly rejected by strict server: %v", err)
	} else {
		t.Error("expected SHA1(ciphertext) to be rejected in strict mode, but got a response")
	}

	// Summary
	results := srv.GetSHA1Results()
	for _, r := range results {
		t.Logf("protoID=%d: matchedCiphertext=%v matchedPlaintext=%v",
			r.ProtoID, r.MatchedCiphertext, r.MatchedPlaintext)
	}
	if len(results) >= 2 {
		if !results[0].MatchedPlaintext {
			t.Error("protoID=1002: expected SHA1(plaintext) match in strict mode")
		}
		if results[1].MatchedPlaintext {
			t.Log("protoID=1005: SHA1(ciphertext) was also accepted — OpenD is lenient")
		}
	}
}

func sendPacket(t *testing.T, conn net.Conn, protoID, serialNo uint32, body []byte, encrypted bool, sha1Hash [20]byte) {
	t.Helper()

	header := make([]byte, 44)
	header[0] = 'F'
	header[1] = 'T'
	binary.LittleEndian.PutUint32(header[2:6], protoID)
	binary.LittleEndian.PutUint32(header[8:12], serialNo)
	binary.LittleEndian.PutUint32(header[12:16], uint32(len(body)))

	if encrypted {
		// Use caller-provided SHA1
		copy(header[16:36], sha1Hash[:])
	} else {
		// Plaintext: SHA1 of body
		h := sha1.Sum(body)
		copy(header[16:36], h[:])
	}

	if _, err := conn.Write(header); err != nil {
		t.Fatalf("write header: %v", err)
	}
	if _, err := conn.Write(body); err != nil {
		t.Fatalf("write body: %v", err)
	}
}

func readResponseRaw(t *testing.T, conn net.Conn) ([]byte, error) {
	t.Helper()
	header := make([]byte, 44)
	if _, err := readFull(conn, header); err != nil {
		return nil, err
	}
	bodyLen := binary.LittleEndian.Uint32(header[12:16])
	body := make([]byte, bodyLen)
	if _, err := readFull(conn, body); err != nil {
		return nil, err
	}
	return body, nil
}

func readResponse(t *testing.T, conn net.Conn) []byte {
	t.Helper()
	body, err := readResponseRaw(t, conn)
	if err != nil {
		t.Fatalf("read response: %v", err)
	}
	return body
}


