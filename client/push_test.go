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

package client

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/push"
	"google.golang.org/protobuf/proto"

	qotpb "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatebasicqot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotupdatekl"
)

func TestParsePushQuote_EmptyList(t *testing.T) {
	s2c := &qotupdatebasicqot.S2C{BasicQotList: []*qotpb.BasicQot{}}
	body, _ := proto.Marshal(s2c)

	data, err := ParsePushQuote(body)
	if err != nil {
		t.Errorf("ParsePushQuote error: %v", err)
	}
	if data != nil {
		t.Error("ParsePushQuote should return nil for empty list")
	}
}

func TestParsePushQuote_InvalidBody(t *testing.T) {
	data, err := ParsePushQuote([]byte("garbage"))
	if err == nil {
		t.Error("ParsePushQuote should error on invalid body")
	}
	if data != nil {
		t.Error("ParsePushQuote should return nil on error")
	}
}

func TestParsePushKLine(t *testing.T) {
	sec := &qotpb.Security{Market: proto.Int32(2), Code: proto.String("HSImain")}
	kl := &qotpb.KLine{
		Time:           proto.String("2025-04-12 10:00:00"),
		OpenPrice:      proto.Float64(24800.0),
		HighPrice:      proto.Float64(25100.0),
		LowPrice:       proto.Float64(24750.0),
		ClosePrice:     proto.Float64(25000.0),
		Volume:         proto.Int64(15000),
		Turnover:       proto.Float64(2.5e9),
		LastClosePrice: proto.Float64(24800.0),
		ChangeRate:     proto.Float64(0.8),
		Timestamp:      proto.Float64(1744502400),
		IsBlank:        proto.Bool(false),
	}
	s2c := &qotupdatekl.S2C{
		Security:  sec,
		Name:      proto.String("HSI Main"),
		KlType:    proto.Int32(6),
		RehabType: proto.Int32(0),
		KlList:    []*qotpb.KLine{kl},
	}
	// Wrap S2C in Response (matching OpenD push body format)
	resp := &qotupdatekl.Response{
		RetType: proto.Int32(0),
		S2C:     s2c,
	}
	body, _ := proto.Marshal(resp)

	data, err := ParsePushKLine(body)
	if err != nil {
		t.Fatalf("ParsePushKLine error: %v", err)
	}
	if data == nil {
		t.Fatal("ParsePushKLine returned nil")
	}
	if data.Market != 2 {
		t.Errorf("Market = %d, want 2", data.Market)
	}
	if data.Code != "HSImain" {
		t.Errorf("Code = %q, want %q", data.Code, "HSImain")
	}
	if data.Close != 25000.0 {
		t.Errorf("Close = %.2f, want 25000.0", data.Close)
	}
}

func TestParsePushKLine_EmptyList(t *testing.T) {
	s2c := &qotupdatekl.S2C{
		KlList:    []*qotpb.KLine{},
		RehabType: proto.Int32(0),
		KlType:    proto.Int32(6),
		Security:  &qotpb.Security{Market: proto.Int32(2), Code: proto.String("HSImain")},
		Name:      proto.String("HSI Main"),
	}
	body, _ := proto.Marshal(s2c)

	data, err := ParsePushKLine(body)
	if err != nil {
		t.Errorf("ParsePushKLine error: %v", err)
	}
	if data != nil {
		t.Error("ParsePushKLine should return nil for empty list")
	}
}

func TestParsePushKLine_InvalidBody(t *testing.T) {
	data, err := ParsePushKLine([]byte("garbage"))
	if err == nil {
		t.Error("ParsePushKLine should error on invalid body")
	}
	if data != nil {
		t.Error("ParsePushKLine should return nil on error")
	}
}

func TestProtoIDConstants(t *testing.T) {
	if ProtoID_Qot_UpdateBasicQot != push.ProtoID_Qot_UpdateBasicQot {
		t.Error("ProtoID_Qot_UpdateBasicQot mismatch")
	}
	if ProtoID_Qot_UpdateKL != push.ProtoID_Qot_UpdateKL {
		t.Error("ProtoID_Qot_UpdateKL mismatch")
	}
	if ProtoID_Qot_UpdateOrderBook != push.ProtoID_Qot_UpdateOrderBook {
		t.Error("ProtoID_Qot_UpdateOrderBook mismatch")
	}
	if ProtoID_Qot_UpdateTicker != push.ProtoID_Qot_UpdateTicker {
		t.Error("ProtoID_Qot_UpdateTicker mismatch")
	}
	if ProtoID_Qot_UpdateRT != push.ProtoID_Qot_UpdateRT {
		t.Error("ProtoID_Qot_UpdateRT mismatch")
	}
	if ProtoID_Qot_UpdateBroker != push.ProtoID_Qot_UpdateBroker {
		t.Error("ProtoID_Qot_UpdateBroker mismatch")
	}
}
