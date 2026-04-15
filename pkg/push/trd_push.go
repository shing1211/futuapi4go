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

import (
	"google.golang.org/protobuf/proto"

	"github.com/shing1211/futuapi4go/pkg/pb/notify"
	"github.com/shing1211/futuapi4go/pkg/pb/trdcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/trdnotify"
	"github.com/shing1211/futuapi4go/pkg/pb/trdupdateorder"
	"github.com/shing1211/futuapi4go/pkg/pb/trdupdateorderfill"
)

const (
	ProtoID_InitConnect         = 1001
	ProtoID_Notify              = 1003
	ProtoID_KeepAlive           = 1002
	ProtoID_Trd_UpdateOrder     = 2208
	ProtoID_Trd_UpdateOrderFill = 2218
	ProtoID_Trd_Notify          = 2207
)

type SystemNotify struct {
	Type          int32
	Event         *notify.GtwEvent
	ProgramStatus *notify.ProgramStatus
	ConnectStatus *notify.ConnectStatus
	QotRight      *notify.QotRight
	ApiLevel      *notify.APILevel
	ApiQuota      *notify.APIQuota
	UsedQuota     *notify.UsedQuota
}

func ParseSystemNotify(body []byte) (*SystemNotify, error) {
	var rsp notify.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &SystemNotify{
		Type:          rsp.GetType(),
		Event:         rsp.GetEvent(),
		ProgramStatus: rsp.GetProgramStatus(),
		ConnectStatus: rsp.GetConnectStatus(),
		QotRight:      rsp.GetQotRight(),
		ApiLevel:      rsp.GetApiLevel(),
		ApiQuota:      rsp.GetApiQuota(),
		UsedQuota:     rsp.GetUsedQuota(),
	}, nil
}

type UpdateOrder struct {
	Header *trdcommon.TrdHeader
	Order  *trdcommon.Order
}

func ParseUpdateOrder(body []byte) (*UpdateOrder, error) {
	var rsp trdupdateorder.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdateOrder{
		Header: rsp.GetHeader(),
		Order:  rsp.GetOrder(),
	}, nil
}

type UpdateOrderFill struct {
	Header    *trdcommon.TrdHeader
	OrderFill *trdcommon.OrderFill
}

func ParseUpdateOrderFill(body []byte) (*UpdateOrderFill, error) {
	var rsp trdupdateorderfill.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &UpdateOrderFill{
		Header:    rsp.GetHeader(),
		OrderFill: rsp.GetOrderFill(),
	}, nil
}

type TrdNotify struct {
	Header *trdcommon.TrdHeader
	Type   int32
}

func ParseTrdNotify(body []byte) (*TrdNotify, error) {
	var rsp trdnotify.S2C
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	return &TrdNotify{
		Header: rsp.GetHeader(),
		Type:   rsp.GetType(),
	}, nil
}
