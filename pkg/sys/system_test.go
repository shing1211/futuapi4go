package sys

import (
	"testing"

	"github.com/shing1211/futuapi4go/pkg/pb/getdelaystatistics"
)

func TestGetGlobalStateResponseFields(t *testing.T) {
	resp := &GetGlobalStateResponse{
		ConnID:        1234567890,
		ServerVer:     1002,
		ServerBuildNo: 6208,
		Time:          1775635200,
		LocalTime:     1775635200.0,
		QotLogined:    true,
		TrdLogined:    true,
		QotSvrIpAddr:  "192.168.1.100",
		TrdSvrIpAddr:  "192.168.1.101",
		MarketHK:      1,
		MarketUS:      1,
		MarketSH:      1,
		MarketSZ:      1,
	}

	if resp.ConnID != 1234567890 {
		t.Errorf("expected ConnID 1234567890, got %d", resp.ConnID)
	}
	if resp.ServerVer != 1002 {
		t.Errorf("expected ServerVer 1002, got %d", resp.ServerVer)
	}
	if !resp.QotLogined {
		t.Error("expected QotLogined to be true")
	}
}

func TestGetUserInfoResponseFields(t *testing.T) {
	resp := &GetUserInfoResponse{
		UserID:                123456789,
		NickName:              "TestUser",
		AvatarUrl:             "https://example.com/avatar.png",
		ApiLevel:              "Level 2",
		IsNeedAgreeDisclaimer: false,
	}

	if resp.UserID != 123456789 {
		t.Errorf("expected UserID 123456789, got %d", resp.UserID)
	}
	if resp.NickName != "TestUser" {
		t.Errorf("expected NickName TestUser, got %s", resp.NickName)
	}
}

func TestGetDelayStatisticsResponseFields(t *testing.T) {
	resp := &GetDelayStatisticsResponse{
		QotPushStatisticsList: []*getdelaystatistics.DelayStatistics{},
	}

	// Just verify the struct can be created and accessed
	if resp.QotPushStatisticsList == nil {
		t.Error("QotPushStatisticsList should not be nil")
	}
}

func TestVerificationRequestFields(t *testing.T) {
	req := &VerificationRequest{
		Code: "123456",
	}

	if req.Code != "123456" {
		t.Errorf("expected Code 123456, got %s", req.Code)
	}
}

