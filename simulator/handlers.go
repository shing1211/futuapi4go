package simulator

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	"gitee.com/shing1211/futuapi4go/pb/getglobalstate"
	"gitee.com/shing1211/futuapi4go/pb/getuserinfo"
	"gitee.com/shing1211/futuapi4go/pb/initconnect"
	"gitee.com/shing1211/futuapi4go/pb/keepalive"
)

func (s *Server) RegisterDefaultHandlers() {
	// InitConnect (1001) - Connection initialization
	s.RegisterHandler(1001, s.handleInitConnect)

	// KeepAlive (1002) - Heartbeat
	s.RegisterHandler(1002, s.handleKeepAlive)

	// GetGlobalState (1004) - Get global state
	s.RegisterHandler(1004, s.handleGetGlobalState)

	// GetUserInfo (1005) - Get user info
	s.RegisterHandler(1005, s.handleGetUserInfo)
}

func (s *Server) handleInitConnect(pkt *Packet) (*Packet, error) {
	var req initconnect.C2S
	if err := proto.Unmarshal(pkt.Body, &req); err != nil {
		return s.errorResponse(pkt, fmt.Errorf("unmarshal request: %w", err)), nil
	}

	connID := uint64(1234567890)
	connAESKey := "mock_aes_key_12345"
	serverVer := int32(10100)
	keepAliveInterval := int32(30)

	s2c := &initconnect.S2C{
		ConnID:            &connID,
		ConnAESKey:        &connAESKey,
		ServerVer:         &serverVer,
		KeepAliveInterval: &keepAliveInterval,
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		return s.errorResponse(pkt, fmt.Errorf("marshal response: %w", err)), nil
	}

	return &Packet{
		Magic:    Magic,
		ProtoID:  pkt.ProtoID,
		SerialNo: pkt.SerialNo,
		BodyLen:  uint32(len(body)),
		Body:     body,
	}, nil
}

func (s *Server) handleKeepAlive(pkt *Packet) (*Packet, error) {
	time := int64(1234567890)

	s2c := &keepalive.S2C{
		Time: &time,
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		return s.errorResponse(pkt, fmt.Errorf("marshal response: %w", err)), nil
	}

	return &Packet{
		Magic:    Magic,
		ProtoID:  pkt.ProtoID,
		SerialNo: pkt.SerialNo,
		BodyLen:  uint32(len(body)),
		Body:     body,
	}, nil
}

func (s *Server) handleGetGlobalState(pkt *Packet) (*Packet, error) {
	marketHK := int32(1)
	marketUS := int32(1)
	marketSH := int32(1)
	marketSZ := int32(1)
	qotLogined := true
	trdLogined := true
	serverVer := int32(10100)
	serverBuildNo := int32(6208)
	serverTime := time.Now().Unix()
	connID := uint64(1234567890)

	s2c := &getglobalstate.S2C{
		MarketHK:      &marketHK,
		MarketUS:      &marketUS,
		MarketSH:      &marketSH,
		MarketSZ:      &marketSZ,
		QotLogined:    &qotLogined,
		TrdLogined:    &trdLogined,
		ServerVer:     &serverVer,
		ServerBuildNo: &serverBuildNo,
		Time:          &serverTime,
		ConnID:        &connID,
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		return s.errorResponse(pkt, fmt.Errorf("marshal response: %w", err)), nil
	}

	return &Packet{
		Magic:    Magic,
		ProtoID:  pkt.ProtoID,
		SerialNo: pkt.SerialNo,
		BodyLen:  uint32(len(body)),
		Body:     body,
	}, nil
}

func (s *Server) handleGetUserInfo(pkt *Packet) (*Packet, error) {
	nickname := "MockUser"
	hkQotRight := int32(1)
	usQotRight := int32(1)
	cnQotRight := int32(1)
	userID := int64(123456789)
	subQuota := int32(100)
	historyKLQuota := int32(100)

	s2c := &getuserinfo.S2C{
		NickName:       &nickname,
		HkQotRight:     &hkQotRight,
		UsQotRight:     &usQotRight,
		CnQotRight:     &cnQotRight,
		UserID:         &userID,
		SubQuota:       &subQuota,
		HistoryKLQuota: &historyKLQuota,
	}

	body, err := proto.Marshal(s2c)
	if err != nil {
		return s.errorResponse(pkt, fmt.Errorf("marshal response: %w", err)), nil
	}

	return &Packet{
		Magic:    Magic,
		ProtoID:  pkt.ProtoID,
		SerialNo: pkt.SerialNo,
		BodyLen:  uint32(len(body)),
		Body:     body,
	}, nil
}
