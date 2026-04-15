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

package simulator

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/getglobalstate"
	"github.com/shing1211/futuapi4go/pkg/pb/getuserinfo"
	"github.com/shing1211/futuapi4go/pkg/pb/initconnect"
	"github.com/shing1211/futuapi4go/pkg/pb/keepalive"
)

func (s *Server) RegisterDefaultHandlers() {
	s.RegisterHandler(1001, s.handleInitConnect)
	s.RegisterHandler(1002, s.handleKeepAlive)
	s.RegisterHandler(1004, s.handleGetGlobalState)
	s.RegisterHandler(1005, s.handleGetUserInfo)
}

func (s *Server) handleInitConnect(pkt *Packet) (*Packet, error) {
	var req initconnect.Request
	if err := proto.Unmarshal(pkt.Body, &req); err != nil {
		return s.errorResponse(pkt, fmt.Errorf("unmarshal request: %w", err))
	}
	c2s := req.GetC2S()
	_ = c2s

	connID := uint64(1234567890)
	connAESKey := "mock_aes_key_12345"
	serverVer := int32(10100)
	keepAliveInterval := int32(30)
	retType := int32(common.RetType_RetType_Succeed)

	resp := &initconnect.Response{
		RetType: &retType,
		S2C: &initconnect.S2C{
			ConnID:            &connID,
			ConnAESKey:        &connAESKey,
			ServerVer:         &serverVer,
			KeepAliveInterval: &keepAliveInterval,
		},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handleKeepAlive(pkt *Packet) (*Packet, error) {
	now := time.Now().Unix()
	retType := int32(common.RetType_RetType_Succeed)

	resp := &keepalive.Response{
		RetType: &retType,
		S2C: &keepalive.S2C{
			Time: &now,
		},
	}

	return s.successResponse(pkt, resp)
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
	retType := int32(common.RetType_RetType_Succeed)

	resp := &getglobalstate.Response{
		RetType: &retType,
		S2C: &getglobalstate.S2C{
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
		},
	}

	return s.successResponse(pkt, resp)
}

func (s *Server) handleGetUserInfo(pkt *Packet) (*Packet, error) {
	nickname := "MockUser"
	hkQotRight := int32(1)
	usQotRight := int32(1)
	cnQotRight := int32(1)
	userID := int64(123456789)
	subQuota := int32(100)
	historyKLQuota := int32(100)
	retType := int32(common.RetType_RetType_Succeed)

	resp := &getuserinfo.Response{
		RetType: &retType,
		S2C: &getuserinfo.S2C{
			NickName:       &nickname,
			HkQotRight:     &hkQotRight,
			UsQotRight:     &usQotRight,
			CnQotRight:     &cnQotRight,
			UserID:         &userID,
			SubQuota:       &subQuota,
			HistoryKLQuota: &historyKLQuota,
		},
	}

	return s.successResponse(pkt, resp)
}
