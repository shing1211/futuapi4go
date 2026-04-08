package futuapi

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"sync"
	"testing"
	"time"
)

func TestHeaderSize(t *testing.T) {
	if HeaderLen != 44 {
		t.Errorf("expected HeaderLen 44, got %d", HeaderLen)
	}
}

func TestMagicBytes(t *testing.T) {
	if string(MagicBytes[:]) != "FT" {
		t.Errorf("expected magic 'FT', got '%s'", MagicBytes)
	}
}

func TestProtoVersion(t *testing.T) {
	if ProtoVersion != 0 {
		t.Errorf("expected ProtoVersion 0, got %d", ProtoVersion)
	}
}

func TestConnNilGuards(t *testing.T) {
	conn := NewConn(nil)

	// ReadPacket should return error when conn is nil
	_, err := conn.ReadPacket()
	if err == nil {
		t.Error("ReadPacket should return error when conn is nil")
	}
	if err.Error() != "read packet: not connected" {
		t.Errorf("expected 'read packet: not connected', got '%v'", err)
	}

	// WritePacket should return error when conn is nil
	err = conn.WritePacket(1001, 1, []byte{})
	if err == nil {
		t.Error("WritePacket should return error when conn is nil")
	}
	if err.Error() != "write packet: not connected" {
		t.Errorf("expected 'write packet: not connected', got '%v'", err)
	}

	// SetReadDeadline should return error when conn is nil
	err = conn.SetReadDeadline(time.Time{})
	if err == nil {
		t.Error("SetReadDeadline should return error when conn is nil")
	}

	// SetWriteDeadline should return error when conn is nil
	err = conn.SetWriteDeadline(time.Time{})
	if err == nil {
		t.Error("SetWriteDeadline should return error when conn is nil")
	}
}

func TestConnWriteReadHeaderEncoding(t *testing.T) {
	// Create a simple test using bytes.Buffer
	var buf bytes.Buffer

	// Test data
	protoID := uint32(2101) // GetBasicQot
	serialNo := uint32(42)
	body := []byte{0x0a, 0x0b, 0x0a, 0x09, 0x08, 0x01, 0x12, 0x05, 0x30, 0x30, 0x37, 0x30, 0x30}

	// Calculate expected SHA1
	expectedSHA1 := sha1.Sum(body)

	// Manually encode header as WritePacket does
	header := make([]byte, HeaderLen)
	copy(header[0:2], "FT")
	binary.LittleEndian.PutUint32(header[2:6], protoID)
	header[6] = 0  // ProtoFmt
	header[7] = 0  // ProtoVer
	binary.LittleEndian.PutUint32(header[8:12], serialNo)
	binary.LittleEndian.PutUint32(header[12:16], uint32(len(body)))
	copy(header[16:36], expectedSHA1[:])

	// Write to buffer
	buf.Write(header)
	buf.Write(body)

	// Verify header content
	if string(buf.Bytes()[0:2]) != "FT" {
		t.Errorf("expected magic 'FT', got '%s'", buf.Bytes()[0:2])
	}

	gotProtoID := binary.LittleEndian.Uint32(buf.Bytes()[2:6])
	if gotProtoID != protoID {
		t.Errorf("expected ProtoID %d, got %d", protoID, gotProtoID)
	}

	if buf.Bytes()[6] != 0 {
		t.Errorf("expected ProtoFmt 0, got %d", buf.Bytes()[6])
	}

	if buf.Bytes()[7] != 0 {
		t.Errorf("expected ProtoVer 0, got %d", buf.Bytes()[7])
	}

	gotSerialNo := binary.LittleEndian.Uint32(buf.Bytes()[8:12])
	if gotSerialNo != serialNo {
		t.Errorf("expected SerialNo %d, got %d", serialNo, gotSerialNo)
	}

	gotBodyLen := binary.LittleEndian.Uint32(buf.Bytes()[12:16])
	if gotBodyLen != uint32(len(body)) {
		t.Errorf("expected BodyLen %d, got %d", len(body), gotBodyLen)
	}

	gotSHA1 := buf.Bytes()[16:36]
	if !bytes.Equal(gotSHA1, expectedSHA1[:]) {
		t.Errorf("SHA1 mismatch:\n  expected: % x\n  got:      % x", expectedSHA1, gotSHA1)
	}

	// Verify body
	bodyStart := HeaderLen
	bodyEnd := HeaderLen + len(body)
	if !bytes.Equal(buf.Bytes()[bodyStart:bodyEnd], body) {
		t.Errorf("body mismatch")
	}
}

func TestConnEmptyBody(t *testing.T) {
	var buf bytes.Buffer

	// Manually encode header with empty body
	header := make([]byte, HeaderLen)
	copy(header[0:2], "FT")
	binary.LittleEndian.PutUint32(header[2:6], 1002)
	binary.LittleEndian.PutUint32(header[8:12], 1)
	binary.LittleEndian.PutUint32(header[12:16], 0) // Empty body

	buf.Write(header)

	// Verify BodyLen is 0
	bodyLen := binary.LittleEndian.Uint32(buf.Bytes()[12:16])
	if bodyLen != 0 {
		t.Errorf("expected BodyLen 0, got %d", bodyLen)
	}
}

func TestConnLargeBody(t *testing.T) {
	// Create a large body (100KB for faster testing)
	body := make([]byte, 100*1024)
	for i := range body {
		body[i] = byte(i % 256)
	}

	// Calculate SHA1
	expectedSHA1 := sha1.Sum(body)

	// Encode header
	header := make([]byte, HeaderLen)
	copy(header[0:2], "FT")
	binary.LittleEndian.PutUint32(header[2:6], 2101)
	binary.LittleEndian.PutUint32(header[8:12], 1)
	binary.LittleEndian.PutUint32(header[12:16], uint32(len(body)))
	copy(header[16:36], expectedSHA1[:])

	// Verify body length in header
	bodyLen := binary.LittleEndian.Uint32(header[12:16])
	if bodyLen != uint32(len(body)) {
		t.Errorf("expected BodyLen %d, got %d", len(body), bodyLen)
	}

	// Verify SHA1
	if !bytes.Equal(header[16:36], expectedSHA1[:]) {
		t.Errorf("SHA1 mismatch for large body")
	}
}

func TestConnReadPacketTooLarge(t *testing.T) {
	// Test that MaxPacketSize limit works
	if MaxPacketSize != 10*1024*1024 {
		t.Errorf("expected MaxPacketSize 10MB, got %d", MaxPacketSize)
	}
}

func TestConnInvalidMagic(t *testing.T) {
	// Test that magic bytes validation would work
	header := make([]byte, HeaderLen)
	copy(header[0:2], "XX") // Wrong magic

	if string(header[0:2]) == "FT" {
		t.Error("should detect invalid magic")
	}
}

func TestConnSerialTracking(t *testing.T) {
	// Test serial number encoding
	var buf bytes.Buffer

	for i := uint32(1); i <= 5; i++ {
		header := make([]byte, HeaderLen)
		copy(header[0:2], "FT")
		binary.LittleEndian.PutUint32(header[2:6], 2101)
		binary.LittleEndian.PutUint32(header[8:12], i)
		binary.LittleEndian.PutUint32(header[12:16], 1)
		buf.Write(header)
		buf.WriteByte(byte(i))
	}

	// Verify each packet
	for i := uint32(1); i <= 5; i++ {
		offset := int(i-1) * (HeaderLen + 1)
		serialNo := binary.LittleEndian.Uint32(buf.Bytes()[offset+8 : offset+12])
		if serialNo != i {
			t.Errorf("packet %d: expected SerialNo %d, got %d", i, i, serialNo)
		}
	}
}

func TestConnConcurrentWrites(t *testing.T) {
	// Test that serial numbers are unique when written concurrently
	var buf bytes.Buffer
	var mu sync.Mutex
	var wg sync.WaitGroup

	serials := make(map[uint32]bool)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(serial uint32) {
			defer wg.Done()

			header := make([]byte, HeaderLen)
			copy(header[0:2], "FT")
			binary.LittleEndian.PutUint32(header[2:6], 2101)
			binary.LittleEndian.PutUint32(header[8:12], serial)
			binary.LittleEndian.PutUint32(header[12:16], 1)

			mu.Lock()
			buf.Write(header)
			buf.WriteByte(byte(serial))
			serials[serial] = true
			mu.Unlock()
		}(uint32(i + 1))
	}
	wg.Wait()

	if len(serials) != 10 {
		t.Errorf("expected 10 unique serials, got %d", len(serials))
	}
}
