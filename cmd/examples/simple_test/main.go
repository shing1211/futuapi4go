// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Simple TCP test to verify Futu OpenD is responding
package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/shing1211/futuapi4go/pkg/pb/initconnect"
)

func main() {
	fmt.Println("=== Simple InitConnect Test ===")

	// Connect
	conn, err := net.DialTimeout("tcp", "127.0.0.1:11111", 5*time.Second)
	if err != nil {
		fmt.Printf("❌ TCP connect failed: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("✅ TCP connected")

	// Build InitConnect request
	c2s := &initconnect.C2S{
		ClientVer:     proto.Int32(10100),
		ClientID:      proto.String("test_client"),
		RecvNotify:    proto.Bool(false),
		PacketEncAlgo: proto.Int32(-1), // No encryption
	}

	req := &initconnect.Request{
		C2S: c2s,
	}

	body, err := proto.Marshal(req)
	if err != nil {
		fmt.Printf("❌ Marshal failed: %v\n", err)
		return
	}
	fmt.Printf("✅ Request marshaled (%d bytes)\n", len(body))

	// Build header (44 bytes per official Futu protocol spec)
	header := make([]byte, 44)
	header[0] = 'F'
	header[1] = 'T'
	binary.LittleEndian.PutUint32(header[2:], 1001)               // ProtoID
	header[6] = 0                                                 // ProtoFmt (0=Protobuf)
	header[7] = 0                                                 // ProtoVer (0)
	binary.LittleEndian.PutUint32(header[8:], 1)                  // SerialNo
	binary.LittleEndian.PutUint32(header[12:], uint32(len(body))) // BodyLen

	// Send
	fmt.Println("📤 Sending request...")
	n, err := conn.Write(header)
	if err != nil {
		fmt.Printf("❌ Write header failed: %v\n", err)
		return
	}
	fmt.Printf("✅ Header sent (%d bytes)\n", n)

	n, err = conn.Write(body)
	if err != nil {
		fmt.Printf("❌ Write body failed: %v\n", err)
		return
	}
	fmt.Printf("✅ Body sent (%d bytes)\n", n)

	// Receive response
	fmt.Println("📥 Waiting for response (10s)...")
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	// Read header (44 bytes)
	respHeader := make([]byte, 44)
	n, err = conn.Read(respHeader)
	if err != nil {
		fmt.Printf("❌ Read header failed: %v\n", err)
		fmt.Println("\n💡 This usually means:")
		fmt.Println("   1. Futu OpenD is not fully logged in")
		fmt.Println("   2. Wrong protocol version")
		fmt.Println("   3. OpenD needs to be restarted")
		return
	}

	fmt.Printf("✅ Response header received (%d bytes)\n", n)

	// Parse header
	magic := string(respHeader[0:2])
	protoID := binary.LittleEndian.Uint32(respHeader[2:6])
	bodyLen := binary.LittleEndian.Uint32(respHeader[12:16])

	fmt.Printf("   Magic: %s\n", magic)
	fmt.Printf("   ProtoID: %d\n", protoID)
	fmt.Printf("   BodyLen: %d\n", bodyLen)

	if bodyLen > 0 && bodyLen < 1000000 {
		respBody := make([]byte, bodyLen)
		n, err = conn.Read(respBody)
		if err != nil {
			fmt.Printf("❌ Read body failed: %v\n", err)
			return
		}
		fmt.Printf("✅ Response body received (%d bytes)\n", n)

		// Try to parse as InitConnect response
		var resp initconnect.Response
		if err := proto.Unmarshal(respBody[:n], &resp); err != nil {
			fmt.Printf("⚠️  Unmarshal failed: %v\n", err)
			fmt.Printf("   Raw: % x\n", respBody[:min(n, 100)])
		} else {
			fmt.Printf("✅ Response parsed!\n")
			fmt.Printf("   RetType: %d\n", resp.GetRetType())
			fmt.Printf("   RetMsg: %s\n", resp.GetRetMsg())
			if resp.S2C != nil {
				fmt.Printf("   ServerVer: %d\n", resp.S2C.GetServerVer())
				fmt.Printf("   ConnID: %d\n", resp.S2C.GetConnID())
				fmt.Printf("   KeepAlive: %d\n", resp.S2C.GetKeepAliveInterval())
			}
		}
	}

	fmt.Println("\n=== Test Complete ===")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
