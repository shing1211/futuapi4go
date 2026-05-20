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
	"github.com/shing1211/futuapi4go/pkg/util"
)

const (
	ProtoID_InitConnect         = 1001
	ProtoID_Notify              = 1003
	ProtoID_KeepAlive           = 1002
	ProtoID_Trd_UpdateOrder     = 2208
	ProtoID_Trd_UpdateOrderFill = 2218
	ProtoID_Trd_Notify          = 2207
)

// SystemNotify contains system notification event data from OpenD.
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

// ParseSystemNotify unmarshals a system notification push body.
func ParseSystemNotify(body []byte) (*SystemNotify, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var rsp notify.Response
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	s2c := rsp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &SystemNotify{
		Type:          util.ProtoInt32(s2c.Type),
		Event:         s2c.GetEvent(),
		ProgramStatus: s2c.ProgramStatus,
		ConnectStatus: s2c.GetConnectStatus(),
		QotRight:      s2c.QotRight,
		ApiLevel:      s2c.GetApiLevel(),
		ApiQuota:      s2c.GetApiQuota(),
		UsedQuota:     s2c.GetUsedQuota(),
	}, nil
}

// UpdateOrder contains an order update from a TRD push.
type UpdateOrder struct {
	Header *trdcommon.TrdHeader
	Order  *trdcommon.Order
}

// ParseUpdateOrder unmarshals a trade order update push body.
func ParseUpdateOrder(body []byte) (*UpdateOrder, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var rsp trdupdateorder.Response
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	s2c := rsp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &UpdateOrder{
		Header: s2c.Header,
		Order:  s2c.Order,
	}, nil
}

// UpdateOrderFill contains an order fill update from a TRD push.
type UpdateOrderFill struct {
	Header    *trdcommon.TrdHeader
	OrderFill *trdcommon.OrderFill
}

// ParseUpdateOrderFill unmarshals a trade order fill update push body.
func ParseUpdateOrderFill(body []byte) (*UpdateOrderFill, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var rsp trdupdateorderfill.Response
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	s2c := rsp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &UpdateOrderFill{
		Header:    s2c.Header,
		OrderFill: s2c.OrderFill,
	}, nil
}

// TrdNotify contains a trade notification from a TRD push.
type TrdNotify struct {
	Header *trdcommon.TrdHeader
	Type   int32
}

// ParseTrdNotify unmarshals a trade notification push body.
func ParseTrdNotify(body []byte) (*TrdNotify, error) {
	if len(body) == 0 {
		return nil, nil
	}
	var rsp trdnotify.Response
	if err := proto.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}
	s2c := rsp.S2C
	if s2c == nil {
		return nil, nil
	}
	return &TrdNotify{
		Header: s2c.Header,
		Type:   util.ProtoInt32(s2c.Type),
	}, nil
}
