package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"

	"gitee.com/shing1211/futuapi4go/pb/initconnect"
	"gitee.com/shing1211/futuapi4go/pb/qotcommon"
	"gitee.com/shing1211/futuapi4go/pb/qotgetbasicqot"
	"google.golang.org/protobuf/proto"
)

const HeaderLen = 48

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

func main() {
	fmt.Println("=== Testing SDK Communication with Simulator ===")

	conn, err := net.DialTimeout("tcp", "127.0.0.1:11111", 5*time.Second)
	if err != nil {
		fmt.Printf("Dial error: %v\n", err)
		return
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	clientVer := int32(10100)
	clientID := "futuapi4go"
	recvNotify := true
	packetEncAlgo := int32(-1)
	programmingLanguage := "Go"

	initReq := &initconnect.C2S{
		ClientVer:           &clientVer,
		ClientID:            &clientID,
		RecvNotify:          &recvNotify,
		PacketEncAlgo:       &packetEncAlgo,
		ProgrammingLanguage: &programmingLanguage,
	}

	initPkt := &initconnect.Request{C2S: initReq}
	initBody, _ := proto.Marshal(initPkt)

	fmt.Println("Sending InitConnect (1001)...")
	if err := sendPacket(conn, 1001, 1, initBody); err != nil {
		fmt.Printf("Send error: %v\n", err)
		return
	}

	resp, err := readResponse(conn)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
		return
	}
	fmt.Printf("InitConnect response: ProtoID=%d, BodyLen=%d\n", resp.ProtoID, resp.BodyLen)

	market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	sec := &qotcommon.Security{Market: &market, Code: &code}
	qotReq := &qotgetbasicqot.C2S{SecurityList: []*qotcommon.Security{sec}}
	qotPkt := &qotgetbasicqot.Request{C2S: qotReq}
	qotBody, _ := proto.Marshal(qotPkt)

	fmt.Println("Sending GetBasicQot (2101)...")
	if err := sendPacket(conn, 2101, 2, qotBody); err != nil {
		fmt.Printf("Send error: %v\n", err)
		return
	}

	resp2, err := readResponse(conn)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
		return
	}
	fmt.Printf("GetBasicQot response: ProtoID=%d, BodyLen=%d\n", resp2.ProtoID, resp2.BodyLen)

	if resp2.BodyLen > 0 {
		body := make([]byte, resp2.BodyLen)
		io.ReadFull(conn, body)
		var rsp qotgetbasicqot.Response
		proto.Unmarshal(body, &rsp)
		fmt.Printf("RetType=%d, RetMsg=%s\n", rsp.GetRetType(), rsp.GetRetMsg())
		if s2c := rsp.GetS2C(); s2c != nil {
			fmt.Printf("Got %d quotes\n", len(s2c.GetBasicQotList()))
			for _, bq := range s2c.GetBasicQotList() {
				fmt.Printf("  %s: price=%.2f\n", bq.GetSecurity().GetCode(), bq.GetCurPrice())
			}
		}
	}

	fmt.Println("\n=== Success! SDK communicates with Simulator ===")
}

func sendPacket(conn net.Conn, protoID uint32, serialNo uint32, body []byte) error {
	header := Header{
		Magic:    [2]byte{'F', 'T'},
		ProtoID:  protoID,
		ProtoFmt: 2,
		ProtoVer: 1,
		SerialNo: serialNo,
		BodyLen:  uint32(len(body)),
	}

	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, &header)
	conn.Write(buf.Bytes())
	if len(body) > 0 {
		conn.Write(body)
	}
	return nil
}

func readResponse(conn net.Conn) (*Header, error) {
	header := make([]byte, HeaderLen)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}

	var h Header
	binary.Read(bytes.NewReader(header), binary.LittleEndian, &h)
	return &h, nil
}
