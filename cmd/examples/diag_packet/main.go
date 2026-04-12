// Diagnostic tool to check raw packet data from Futu OpenD
package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("=== Raw Packet Diagnostic ===")
	fmt.Println("Connecting to 127.0.0.1:11111...")

	conn, err := net.DialTimeout("tcp", "127.0.0.1:11111", 5*time.Second)
	if err != nil {
		fmt.Printf("❌ Connection failed: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("✓ Connected!")
	fmt.Println("Sending InitConnect request...")

	// Build a minimal InitConnect request (ProtoID 1001)
	// This is a simplified version just to get a response
	header := make([]byte, 48)

	// Magic bytes
	header[0] = 'F'
	header[1] = 'T'

	// ProtoID (1001 for InitConnect) - Little Endian
	binary.LittleEndian.PutUint32(header[2:], 1001)

	// ProtoFmt (1 for Protobuf)
	header[6] = 1

	// ProtoVer (1)
	binary.LittleEndian.PutUint16(header[7:], 1)

	// SerialNo (1)
	binary.LittleEndian.PutUint32(header[9:], 1)

	// Body - minimal InitConnect protobuf
	// C2S: {clientVer: 10100, clientID: "test"}
	body := []byte{
		0x08, 0x94, 0x7F, // clientVer = 10100 (varint)
		0x12, 0x04, 0x74, 0x65, 0x73, 0x74, // clientID = "test"
	}

	// BodyLen
	binary.LittleEndian.PutUint32(header[13:], uint32(len(body)))

	// Send header
	if _, err := conn.Write(header); err != nil {
		fmt.Printf("❌ Write header failed: %v\n", err)
		return
	}

	// Send body
	if _, err := conn.Write(body); err != nil {
		fmt.Printf("❌ Write body failed: %v\n", err)
		return
	}

	fmt.Println("✓ Request sent!")
	fmt.Println("Waiting for response...")

	// Read response header
	respHeader := make([]byte, 48)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, err := conn.Read(respHeader)
	if err != nil {
		fmt.Printf("❌ Read header failed: %v\n", err)
		return
	}

	fmt.Printf("✓ Received %d bytes (header)\n", n)

	// Parse header
	magic := string(respHeader[0:2])
	protoID := binary.LittleEndian.Uint32(respHeader[2:6])
	protoFmt := respHeader[6]
	protoVer := binary.LittleEndian.Uint16(respHeader[7:9])
	serialNo := binary.LittleEndian.Uint32(respHeader[9:13])
	bodyLen := binary.LittleEndian.Uint32(respHeader[13:17])

	fmt.Printf("\n=== Response Header ===\n")
	fmt.Printf("Magic:    %s (hex: % x)\n", magic, respHeader[0:2])
	fmt.Printf("ProtoID:  %d\n", protoID)
	fmt.Printf("ProtoFmt: %d\n", protoFmt)
	fmt.Printf("ProtoVer: %d\n", protoVer)
	fmt.Printf("SerialNo: %d\n", serialNo)
	fmt.Printf("BodyLen:  %d\n", bodyLen)

	if bodyLen > 10*1024*1024 {
		fmt.Printf("\n⚠️  Body length too large! (%d bytes)\n", bodyLen)
		fmt.Println("This suggests a protocol mismatch.")
		return
	}

	if bodyLen > 0 {
		respBody := make([]byte, bodyLen)
		n, err = conn.Read(respBody)
		if err != nil {
			fmt.Printf("❌ Read body failed: %v\n", err)
			return
		}
		fmt.Printf("\n✓ Received %d bytes (body)\n", n)
		fmt.Printf("Body (hex): % x\n", respBody[:min(n, 100)])
	}

	fmt.Println("\n=== Diagnostic Complete ===")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
