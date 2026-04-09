// Raw protocol diagnostic - sends exact bytes and captures response
package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("=== Raw Protocol Diagnostic ===")
	
	// Connect
	conn, err := net.DialTimeout("tcp", "127.0.0.1:11111", 5*time.Second)
	if err != nil {
		fmt.Printf("❌ TCP connect failed: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("✅ TCP connected")
	
	// Test 1: Send just 2 bytes "FT" and see if OpenD responds
	fmt.Println("\n[Test 1] Sending magic bytes 'FT'...")
	conn.SetWriteDeadline(time.Now().Add(2 * time.Second))
	_, err = conn.Write([]byte("FT"))
	if err != nil {
		fmt.Printf("❌ Write failed: %v\n", err)
		return
	}
	
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("⏱️  No response to magic bytes (expected)\n")
	} else {
		fmt.Printf("✅ Response: % x\n", buf[:n])
	}
	
	// Test 2: Send complete InitConnect header
	fmt.Println("\n[Test 2] Sending complete header...")
	
	// Clear any pending data
	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	conn.Read(buf)
	
	// Build header - exact 48 bytes
	header := make([]byte, 48)
	
	// Magic: "FT"
	header[0] = 'F'
	header[1] = 'T'
	
	// ProtoID: 1001 (InitConnect) - Little Endian uint32
	binary.LittleEndian.PutUint32(header[2:], 1001)
	
	// ProtoFmt: 1 (Protobuf format) - This is ProtoFmt enum, 1 byte
	header[6] = 1
	
	// ProtoVer: 1 (uint16)
	binary.LittleEndian.PutUint16(header[7:], 1)
	
	// SerialNo: 1 (uint32)
	binary.LittleEndian.PutUint32(header[9:], 1)
	
	// BodyLen: 0 (for now, test with empty body)
	binary.LittleEndian.PutUint32(header[13:], 0)
	
	// SHA1: 20 bytes of zeros
	// Reserved: 8 bytes of zeros (already zero)
	
	fmt.Printf("Header (hex): % x\n", header)
	fmt.Printf("Header size: %d bytes\n", len(header))
	
	_, err = conn.Write(header)
	if err != nil {
		fmt.Printf("❌ Write header failed: %v\n", err)
		return
	}
	
	fmt.Println("📥 Waiting for response (5s)...")
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Printf("❌ No response: %v\n", err)
		fmt.Println("\n💡 OpenD may require:")
		fmt.Println("   • Body data (not just header)")
		fmt.Println("   • Correct protobuf encoding")
		fmt.Println("   • Specific clientVer format")
	} else {
		fmt.Printf("✅ Response received: %d bytes\n", n)
		fmt.Printf("Response (hex): % x\n", buf[:n])
	}
	
	// Test 3: Send header + minimal body
	fmt.Println("\n[Test 3] Sending header + minimal valid body...")
	
	// Minimal protobuf for InitConnect.C2S
	// Field 1 (clientVer, int32): tag=0x08, value=10100=0x2774
	// Field 2 (clientID, string): tag=0x12, len=4, "test"
	minimalBody := []byte{
		0x08, 0x94, 0x7F, // clientVer = 10100 (varint)
		0x12, 0x04, 0x74, 0x65, 0x73, 0x74, // clientID = "test"
	}
	
	// Update header with body length
	binary.LittleEndian.PutUint32(header[13:], uint32(len(minimalBody)))
	
	// Send header + body
	conn.Write(header)
	conn.Write(minimalBody)
	
	fmt.Printf("Sent: %d byte header + %d byte body\n", len(header), len(minimalBody))
	
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Printf("❌ Still no response: %v\n", err)
		fmt.Println("\n🔍 This definitively means:")
		fmt.Println("   • OpenD is not in 'ready' state")
		fmt.Println("   • OpenD requires login first")
		fmt.Println("   • OR OpenD has API access disabled")
	} else {
		fmt.Printf("✅ Response: %d bytes\n", n)
		fmt.Printf("Response (hex): % x\n", buf[:n])
		
		// Parse response header
		if n >= 48 {
			respMagic := string(buf[0:2])
			respProtoID := binary.LittleEndian.Uint32(buf[2:6])
			respBodyLen := binary.LittleEndian.Uint32(buf[13:17])
			
			fmt.Printf("\n=== Response Header ===\n")
			fmt.Printf("Magic:    %s\n", respMagic)
			fmt.Printf("ProtoID:  %d\n", respProtoID)
			fmt.Printf("BodyLen:  %d\n", respBodyLen)
			
			if respBodyLen > 0 && respBodyLen < 10000 {
				bodyBuf := make([]byte, respBodyLen)
				_, err = conn.Read(bodyBuf)
				if err == nil {
					fmt.Printf("Body (hex): % x\n", bodyBuf)
				}
			}
		}
	}
	
	fmt.Println("\n=== Diagnostic Complete ===")
}

