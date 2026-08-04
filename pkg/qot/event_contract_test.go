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

package qot

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsubeventcontract"
)

func TestBuildECSecurity(t *testing.T) {
	s := BuildECSecurity("EC.KXWCADVANCE-26JUL14FRAESP-FRA")
	if s == nil {
		t.Fatal("BuildECSecurity returned nil")
	}
	if s.Market == nil || *s.Market != int32(qotcommon.QotMarket_QotMarket_EventContract) {
		t.Errorf("market = %v, want %v", s.Market, int32(qotcommon.QotMarket_QotMarket_EventContract))
	}
	if s.Code == nil || *s.Code != "EC.KXWCADVANCE-26JUL14FRAESP-FRA" {
		t.Errorf("code = %v, want %v", s.Code, "EC.KXWCADVANCE-26JUL14FRAESP-FRA")
	}
}

func TestPredSideEnumValues(t *testing.T) {
	cases := []struct {
		got  common.PredSide
		want int32
	}{
		{common.PredSide_PredSide_Unknown, 0},
		{common.PredSide_PredSide_Yes, 1},
		{common.PredSide_PredSide_No, 2},
	}
	for _, c := range cases {
		if int32(c.got) != c.want {
			t.Errorf("PredSide = %v, want %v", int32(c.got), c.want)
		}
	}
}

func TestECMarketAndStatusEnums(t *testing.T) {
	if int32(qotcommon.QotMarket_QotMarket_EventContract) != 101 {
		t.Errorf("QotMarket_EventContract = %v, want 101", int32(qotcommon.QotMarket_QotMarket_EventContract))
	}
	if int32(qotcommon.EC_Status_EC_Status_Active) != 2 {
		t.Errorf("EC_Status_Active = %v, want 2", int32(qotcommon.EC_Status_EC_Status_Active))
	}
	if int32(qotcommon.EC_Status_EC_Status_Settled) != 5 {
		t.Errorf("EC_Status_Settled = %v, want 5", int32(qotcommon.EC_Status_EC_Status_Settled))
	}
	if int32(qotcommon.EC_KlineSource_EC_KlineSource_None) != 0 {
		t.Errorf("EC_KlineSource_None = %v, want 0", int32(qotcommon.EC_KlineSource_EC_KlineSource_None))
	}
	if int32(qotcommon.EC_KlineSource_EC_KlineSource_OrderBookYes) != 1 {
		t.Errorf("EC_KlineSource_OrderBookYes = %v, want 1", int32(qotcommon.EC_KlineSource_EC_KlineSource_OrderBookYes))
	}
	if int32(qotcommon.EC_ContractType_EC_ContractType_Binary) != 1 {
		t.Errorf("EC_ContractType_Binary = %v, want 1", int32(qotcommon.EC_ContractType_EC_ContractType_Binary))
	}
	if int32(qotcommon.EC_Frequency_EC_Frequency_OneOff) != 7 {
		t.Errorf("EC_Frequency_OneOff = %v, want 7", int32(qotcommon.EC_Frequency_EC_Frequency_OneOff))
	}
}

func TestSubEventContractRequiresSecurityOrUnsubAll(t *testing.T) {
	// Empty request without IsUnsubAll must fail validation.
	req := &qotsubeventcontract.C2S{}
	if err := SubEventContract(t.Context(), nil, req); err == nil {
		t.Error("SubEventContract with empty request should return error")
	}
}

func TestOrderBookLevelFields(t *testing.T) {
	price := 0.42
	size := 100.0
	lvl := &qotgeteventcontractorderbook.OrderBookLevel{Price: &price, Size: &size}
	if lvl.GetPrice() != 0.42 {
		t.Errorf("OrderBookLevel.Price = %v, want 0.42", lvl.GetPrice())
	}
	if lvl.GetSize() != 100.0 {
		t.Errorf("OrderBookLevel.Size = %v, want 100.0", lvl.GetSize())
	}
}