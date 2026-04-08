package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// Test 1: Manual encoding
	fmt.Println("=== Test 1: Manual header encoding ===")
	
	header := make([]byte, 48)
	header[0] = 'F'
	header[1] = 'T'
	binary.LittleEndian.PutUint32(header[2:], 1001)   // ProtoID
	header[6] = 1                                       // ProtoFmt
	binary.LittleEndian.PutUint16(header[7:], 1)      // ProtoVer
	binary.LittleEndian.PutUint32(header[9:], 1)      // SerialNo
	binary.LittleEndian.PutUint32(header[13:], 10)    // BodyLen (10 bytes)
	
	fmt.Printf("Header (first 20 bytes): % x\n", header[:20])
	fmt.Printf("Header length: %d bytes\n\n", len(header))
	
	// Test 2: Using struct with binary.Write
	fmt.Println("=== Test 2: Struct encoding ===")
	
	type Header struct {
		Magic    [2]byte
		ProtoID  uint32
		ProtoFmt uint32
		ProtoVer uint16
		SerialNo uint32
		BodyLen  uint32
		BodySHA1 [20]byte
		Reserved [8]byte
	}
	
	h := Header{
		ProtoID:  1001,
		ProtoFmt: 1,
		ProtoVer: 1,
		SerialNo: 1,
		BodyLen:  10,
	}
	h.Magic[0] = 'F'
	h.Magic[1] = 'T'
	
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, &h)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Encoded header (%d bytes): % x\n", buf.Len(), buf.Bytes()[:20])
	fmt.Printf("\n=== Comparison ===")
	fmt.Printf("\nManual:  % x\n", header[:20])
	fmt.Printf("Struct:  % x\n", buf.Bytes()[:20])
	
	// Check if they match
	match := true
	for i := 0; i < 20; i++ {
		if header[i] != buf.Bytes()[i] {
			match = false
			fmt.Printf("\nMismatch at byte %d: manual=%02x, struct=%02x", i, header[i], buf.Bytes()[i])
		}
	}
	
	if match {
		fmt.Println("\n✅ Headers match! Structure is correct.")
	}
}
