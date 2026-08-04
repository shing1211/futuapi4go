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

package push

import "testing"

func TestParseUpdateEventContractOrderBookInvalidData(t *testing.T) {
	result, err := ParseUpdateEventContractOrderBook([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateEventContractOrderBook should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateEventContractOrderBook should return nil for empty data")
	}

	_, err = ParseUpdateEventContractOrderBook([]byte{0x00, 0x01, 0x02})
	if err == nil {
		t.Error("ParseUpdateEventContractOrderBook should fail with invalid protobuf data")
	}
}

func TestParseUpdateEventContractKlineInvalidData(t *testing.T) {
	result, err := ParseUpdateEventContractKline([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateEventContractKline should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateEventContractKline should return nil for empty data")
	}

	_, err = ParseUpdateEventContractKline([]byte{0x00, 0x01, 0x02})
	if err == nil {
		t.Error("ParseUpdateEventContractKline should fail with invalid protobuf data")
	}
}

func TestParseUpdateEventContractTickerInvalidData(t *testing.T) {
	result, err := ParseUpdateEventContractTicker([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateEventContractTicker should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateEventContractTicker should return nil for empty data")
	}

	_, err = ParseUpdateEventContractTicker([]byte{0x00, 0x01, 0x02})
	if err == nil {
		t.Error("ParseUpdateEventContractTicker should fail with invalid protobuf data")
	}
}

func TestECPushProtoIDs(t *testing.T) {
	cases := []struct {
		name string
		got  int32
		want int32
	}{
		{"UpdateECOrderBook", ProtoID_Qot_UpdateEventContractOrderBook, 3450},
		{"UpdateECKline", ProtoID_Qot_UpdateEventContractKline, 3451},
		{"UpdateECTicker", ProtoID_Qot_UpdateEventContractTicker, 3452},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %d, want %d", c.name, c.got, c.want)
		}
	}
}