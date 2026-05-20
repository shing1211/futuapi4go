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

package trd

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/util"
	"github.com/shing1211/futuapi4go/pkg/constant"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/trdsubaccpush"
	"github.com/shing1211/futuapi4go/pkg/pb/trdunlocktrade"
)

// UnlockTradeRequest is the request to unlock or lock trading with a password.
type UnlockTradeRequest struct {
	Unlock       bool
	PwdMD5       constant.SensitiveString
	SecurityFirm int32
}

// UnlockTrade unlocks or locks trading functionality using the provided password.
// Returns an error if the unlock operation fails.
func UnlockTrade(ctx context.Context, c *futuapi.Client, req *UnlockTradeRequest) error {
	if req == nil {
		return fmt.Errorf("UnlockTrade: request is nil")
	}
	if req.PwdMD5.IsEmpty() {
		return fmt.Errorf("password MD5 is required")
	}

	pwdRaw := req.PwdMD5.Raw()
	c2s := &trdunlocktrade.C2S{
		Unlock:       &req.Unlock,
		PwdMD5:       &pwdRaw,
		SecurityFirm: &req.SecurityFirm,
	}

	pkt := &trdunlocktrade.Request{C2S: c2s}
	var rsp trdunlocktrade.Response

	if err := c.RequestContext(ctx, ProtoID_UnlockTrade, pkt, &rsp); err != nil {
		return err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return wrapError("UnlockTrade", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	return nil
}

// SubAccPushRequest is the request to subscribe to account push notifications.
type SubAccPushRequest struct {
	AccIDList []uint64
}

// SubAccPush subscribes to account push notifications for the specified account IDs.
// Returns an error if the subscription fails.
func SubAccPush(ctx context.Context, c *futuapi.Client, req *SubAccPushRequest) error {
	if req == nil {
		return fmt.Errorf("SubAccPush: request is nil")
	}
	if len(req.AccIDList) == 0 {
		return fmt.Errorf("account ID list is empty")
	}

	c2s := &trdsubaccpush.C2S{
		AccIDList: req.AccIDList,
	}

	pkt := &trdsubaccpush.Request{C2S: c2s}
	var rsp trdsubaccpush.Response

	if err := c.RequestContext(ctx, ProtoID_SubAccPush, pkt, &rsp); err != nil {
		return err
	}

	if util.ProtoInt32(rsp.RetType) != int32(common.RetType_RetType_Succeed) {
		return wrapError("SubAccPush", util.ProtoInt32(rsp.RetType), util.ProtoStr(rsp.RetMsg))
	}

	return nil
}
