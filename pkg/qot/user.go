// Package qot provides market data APIs for the Futu OpenD SDK.
//
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
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/common"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetipolist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetpricereminder"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusersecurity"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgetusersecuritygroup"
	"github.com/shing1211/futuapi4go/pkg/pb/qotmodifyusersecurity"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsetpricereminder"
)

// GetUserSecurityResponse is the response type for GetUserSecurity.
type GetUserSecurityResponse struct {
	StaticInfoList []*qotcommon.SecurityStaticInfo
}

// GetUserSecurity returns the list of user-defined securities in the specified group.
func GetUserSecurity(ctx context.Context, c *futuapi.Client, groupName string) (*GetUserSecurityResponse, error) {
	if groupName == "" {
		return nil, fmt.Errorf("group name is required")
	}

	c2s := &qotgetusersecurity.C2S{
		GroupName: &groupName,
	}

	pkt := &qotgetusersecurity.Request{C2S: c2s}
	var rsp qotgetusersecurity.Response

	if err := c.RequestContext(ctx, ProtoID_GetUserSecurity, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetUserSecurity", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetUserSecurity: s2c is nil")
	}

	result := &GetUserSecurityResponse{
		StaticInfoList: make([]*qotcommon.SecurityStaticInfo, 0, len(s2c.GetStaticInfoList())),
	}
	for _, si := range s2c.GetStaticInfoList() {
		if si == nil {
			continue
		}
		result.StaticInfoList = append(result.StaticInfoList, si)
	}

	return result, nil
}

// GetUserSecurityGroupRequest defines parameters for GetUserSecurityGroup.
type GetUserSecurityGroupRequest struct {
	GroupType int32
}

// UserSecurityGroupData represents a user-defined security group.
type UserSecurityGroupData struct {
	GroupName string
	GroupType int32
}

// GetUserSecurityGroupResponse is the response type for GetUserSecurityGroup.
type GetUserSecurityGroupResponse struct {
	GroupList []*UserSecurityGroupData
}

// GetUserSecurityGroup returns the list of user-defined security groups.
func GetUserSecurityGroup(ctx context.Context, c *futuapi.Client, req *GetUserSecurityGroupRequest) (*GetUserSecurityGroupResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetUserSecurityGroup: request is nil")
	}
	c2s := &qotgetusersecuritygroup.C2S{
		GroupType: &req.GroupType,
	}

	pkt := &qotgetusersecuritygroup.Request{C2S: c2s}
	var rsp qotgetusersecuritygroup.Response

	if err := c.RequestContext(ctx, ProtoID_GetUserSecurityGroup, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetUserSecurityGroup", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetUserSecurityGroup: s2c is nil")
	}

	result := &GetUserSecurityGroupResponse{
		GroupList: make([]*UserSecurityGroupData, 0, len(s2c.GetGroupList())),
	}

	for _, g := range s2c.GetGroupList() {
		if g == nil {
			continue
		}
		result.GroupList = append(result.GroupList, &UserSecurityGroupData{
			GroupName: g.GetGroupName(),
			GroupType: g.GetGroupType(),
		})
	}

	return result, nil
}

// ModifyUserSecurityRequest defines parameters for ModifyUserSecurity.
type ModifyUserSecurityRequest struct {
	GroupName    string
	Op           int32
	SecurityList []*qotcommon.Security
}

// ModifyUserSecurityResponse is the response type for ModifyUserSecurity.
type ModifyUserSecurityResponse struct {
	RetType int32
	RetMsg  string
}

// ModifyUserSecurity adds or removes securities from a user-defined group.
func ModifyUserSecurity(ctx context.Context, c *futuapi.Client, req *ModifyUserSecurityRequest) (*ModifyUserSecurityResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("ModifyUserSecurity: request is nil")
	}
	if req.GroupName == "" {
		return nil, fmt.Errorf("group name is required")
	}
	if req.Op == 0 {
		return nil, fmt.Errorf("operation is required")
	}

	c2s := &qotmodifyusersecurity.C2S{
		GroupName:    &req.GroupName,
		Op:           &req.Op,
		SecurityList: req.SecurityList,
	}

	pkt := &qotmodifyusersecurity.Request{C2S: c2s}
	var rsp qotmodifyusersecurity.Response

	if err := c.RequestContext(ctx, ProtoID_ModifyUserSecurity, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("ModifyUserSecurity", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, wrapError("ModifyUserSecurity", int32(common.RetType_RetType_Unknown), "s2c is nil")
	}

	return &ModifyUserSecurityResponse{
		RetType: rsp.GetRetType(),
		RetMsg:  rsp.GetRetMsg(),
	}, nil
}

// SetPriceReminderRequest defines parameters for SetPriceReminder.
type SetPriceReminderRequest struct {
	Security            *qotcommon.Security
	Op                  int32
	Key                 int64
	Type                int32
	Freq                int32
	Value               float64
	Note                string
	ReminderSessionList []int32
}

// SetPriceReminderResponse is the response type for SetPriceReminder.
type SetPriceReminderResponse struct {
	Key int64
}

// SetPriceReminder creates, updates, or deletes a price reminder.
func SetPriceReminder(ctx context.Context, c *futuapi.Client, req *SetPriceReminderRequest) (*SetPriceReminderResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("SetPriceReminder: request is nil")
	}
	if req.Security == nil {
		return nil, fmt.Errorf("security is required")
	}
	if req.Op == 0 {
		return nil, fmt.Errorf("operation is required")
	}

	c2s := &qotsetpricereminder.C2S{
		Security:            req.Security,
		Op:                  &req.Op,
		ReminderSessionList: req.ReminderSessionList,
	}
	if req.Key != 0 {
		c2s.Key = &req.Key
	}
	if req.Type != 0 {
		c2s.Type = &req.Type
	}
	if req.Freq != 0 {
		c2s.Freq = &req.Freq
	}
	if req.Value != 0 {
		c2s.Value = &req.Value
	}
	if req.Note != "" {
		c2s.Note = &req.Note
	}

	pkt := &qotsetpricereminder.Request{C2S: c2s}
	var rsp qotsetpricereminder.Response

	if err := c.RequestContext(ctx, ProtoID_SetPriceReminder, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("SetPriceReminder", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("SetPriceReminder: s2c is nil")
	}

	return &SetPriceReminderResponse{
		Key: s2c.GetKey(),
	}, nil
}

// PriceReminderItemInfo represents a single price reminder item.
type PriceReminderItemInfo struct {
	Key                 int64
	Type                int32
	Value               float64
	Note                string
	Freq                int32
	IsEnable            bool
	ReminderSessionList []int32
}

// PriceReminderInfo represents the price reminder settings for a security.
type PriceReminderInfo struct {
	Security *qotcommon.Security
	Name     string
	ItemList []*PriceReminderItemInfo
}

// GetPriceReminderResponse is the response type for GetPriceReminder.
type GetPriceReminderResponse struct {
	PriceReminderList []*PriceReminderInfo
}

// GetPriceReminder returns price reminder settings for the given security.
func GetPriceReminder(ctx context.Context, c *futuapi.Client, security *qotcommon.Security, market int32) (*GetPriceReminderResponse, error) {
	if security == nil {
		return nil, fmt.Errorf("security is required")
	}

	c2s := &qotgetpricereminder.C2S{
		Security: security,
		Market:   &market,
	}

	pkt := &qotgetpricereminder.Request{C2S: c2s}
	var rsp qotgetpricereminder.Response

	if err := c.RequestContext(ctx, ProtoID_GetPriceReminder, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetPriceReminder", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetPriceReminder: s2c is nil")
	}

	result := &GetPriceReminderResponse{
		PriceReminderList: make([]*PriceReminderInfo, 0, len(s2c.GetPriceReminderList())),
	}

	for _, pr := range s2c.GetPriceReminderList() {
		if pr == nil {
			continue
		}
		info := &PriceReminderInfo{
			Security: pr.GetSecurity(),
			Name:     pr.GetName(),
			ItemList: make([]*PriceReminderItemInfo, 0, len(pr.GetItemList())),
		}
		for _, item := range pr.GetItemList() {
			if item == nil {
				continue
			}
			info.ItemList = append(info.ItemList, &PriceReminderItemInfo{
				Key:                 item.GetKey(),
				Type:                item.GetType(),
				Value:               item.GetValue(),
				Note:                item.GetNote(),
				Freq:                item.GetFreq(),
				IsEnable:            item.GetIsEnable(),
				ReminderSessionList: item.GetReminderSessionList(),
			})
		}
		result.PriceReminderList = append(result.PriceReminderList, info)
	}

	return result, nil
}

// GetIpoListRequest defines parameters for GetIpoList.
type GetIpoListRequest struct {
	Market int32
}

// BasicIpoData represents basic IPO data.
type BasicIpoData struct {
	Security      *qotcommon.Security
	Name          string
	ListTime      string
	ListTimestamp float64
}

// CNIpoExData represents China A-share IPO extended data.
type CNIpoExData struct {
	ApplyCode              string
	IssueSize              int64
	OnlineIssueSize        int64
	ApplyUpperLimit        int64
	ApplyLimitMarketValue  int64
	IsEstimateIpoPrice     bool
	IpoPrice               float64
	IndustryPeRate         float64
	IsEstimateWinningRatio bool
	WinningRatio           float64
	IssuePeRate            float64
	ApplyTime              string
	ApplyTimestamp         float64
	WinningTime            string
	WinningTimestamp       float64
	IsHasWon               bool
	WinningNumDataList     []*qotgetipolist.WinningNumData
}

// HKIpoExData represents Hong Kong IPO extended data.
type HKIpoExData struct {
	IpoPriceMin       float64
	IpoPriceMax       float64
	ListPrice         float64
	LotSize           int32
	EntrancePrice     float64
	IsSubscribeStatus bool
	ApplyEndTime      string
	ApplyEndTimestamp float64
}

// USIpoExData represents US IPO extended data.
type USIpoExData struct {
	IpoPriceMin float64
	IpoPriceMax float64
	IssueSize   int64
}

// IpoData represents complete IPO data including basic and market-specific extended data.
type IpoData struct {
	Basic    *BasicIpoData
	CnExData *CNIpoExData
	HkExData *HKIpoExData
	UsExData *USIpoExData
}

// GetIpoListResponse is the response type for GetIpoList.
type GetIpoListResponse struct {
	IpoList []*IpoData
}

// GetIpoList returns the list of upcoming and recently listed IPOs for the given market.
func GetIpoList(ctx context.Context, c *futuapi.Client, req *GetIpoListRequest) (*GetIpoListResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("GetIpoList: request is nil")
	}
	if req.Market == 0 {
		return nil, fmt.Errorf("invalid market: must be non-zero")
	}

	c2s := &qotgetipolist.C2S{
		Market: &req.Market,
	}

	pkt := &qotgetipolist.Request{C2S: c2s}
	var rsp qotgetipolist.Response

	if err := c.RequestContext(ctx, ProtoID_GetIpoList, pkt, &rsp); err != nil {
		return nil, err
	}

	if rsp.GetRetType() != int32(common.RetType_RetType_Succeed) {
		return nil, wrapError("GetIpoList", rsp.GetRetType(), rsp.GetRetMsg())
	}

	s2c := rsp.GetS2C()
	if s2c == nil {
		return nil, fmt.Errorf("GetIpoList: s2c is nil")
	}

	result := &GetIpoListResponse{
		IpoList: make([]*IpoData, 0, len(s2c.GetIpoList())),
	}

	for _, ipo := range s2c.GetIpoList() {
		if ipo == nil {
			continue
		}
		ipoData := &IpoData{}

		if basic := ipo.GetBasic(); basic != nil {
			ipoData.Basic = &BasicIpoData{
				Security:      basic.GetSecurity(),
				Name:          basic.GetName(),
				ListTime:      basic.GetListTime(),
				ListTimestamp: basic.GetListTimestamp(),
			}
		}

		if cnEx := ipo.GetCnExData(); cnEx != nil {
			ipoData.CnExData = &CNIpoExData{
				ApplyCode:              cnEx.GetApplyCode(),
				IssueSize:              cnEx.GetIssueSize(),
				OnlineIssueSize:        cnEx.GetOnlineIssueSize(),
				ApplyUpperLimit:        cnEx.GetApplyUpperLimit(),
				ApplyLimitMarketValue:  cnEx.GetApplyLimitMarketValue(),
				IsEstimateIpoPrice:     cnEx.GetIsEstimateIpoPrice(),
				IpoPrice:               cnEx.GetIpoPrice(),
				IndustryPeRate:         cnEx.GetIndustryPeRate(),
				IsEstimateWinningRatio: cnEx.GetIsEstimateWinningRatio(),
				WinningRatio:           cnEx.GetWinningRatio(),
				IssuePeRate:            cnEx.GetIssuePeRate(),
				ApplyTime:              cnEx.GetApplyTime(),
				ApplyTimestamp:         cnEx.GetApplyTimestamp(),
				WinningTime:            cnEx.GetWinningTime(),
				WinningTimestamp:       cnEx.GetWinningTimestamp(),
				IsHasWon:               cnEx.GetIsHasWon(),
				WinningNumDataList:     cnEx.GetWinningNumData(),
			}
		}

		if hkEx := ipo.GetHkExData(); hkEx != nil {
			ipoData.HkExData = &HKIpoExData{
				IpoPriceMin:       hkEx.GetIpoPriceMin(),
				IpoPriceMax:       hkEx.GetIpoPriceMax(),
				ListPrice:         hkEx.GetListPrice(),
				LotSize:           hkEx.GetLotSize(),
				EntrancePrice:     hkEx.GetEntrancePrice(),
				IsSubscribeStatus: hkEx.GetIsSubscribeStatus(),
				ApplyEndTime:      hkEx.GetApplyEndTime(),
				ApplyEndTimestamp: hkEx.GetApplyEndTimestamp(),
			}
		}

		if usEx := ipo.GetUsExData(); usEx != nil {
			ipoData.UsExData = &USIpoExData{
				IpoPriceMin: usEx.GetIpoPriceMin(),
				IpoPriceMax: usEx.GetIpoPriceMax(),
				IssueSize:   usEx.GetIssueSize(),
			}
		}

		result.IpoList = append(result.IpoList, ipoData)
	}

	return result, nil
}
