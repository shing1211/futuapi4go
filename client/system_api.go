package client

import (
	"context"
	"fmt"

	"github.com/shing1211/futuapi4go/pkg/sys"
)

// GetGlobalState retrieves global connection state.
func GetGlobalState(ctx context.Context, c *Client) (*GlobalState, error) {
	resp, err := sys.GetGlobalState(ctx, c.inner)
	if err != nil {
		return nil, err
	}

	return &GlobalState{
		ServerVer:      resp.ServerVer,
		ServerBuildNo:  resp.ServerBuildNo,
		Time:           resp.Time,
		LocalTime:      resp.LocalTime,
		QotLogined:     resp.QotLogined,
		TrdLogined:     resp.TrdLogined,
		MarketHK:       resp.MarketHK,
		MarketUS:       resp.MarketUS,
		MarketSH:       resp.MarketSH,
		MarketSZ:       resp.MarketSZ,
		MarketHKFuture: resp.MarketHKFuture,
		MarketUSFuture: resp.MarketUSFuture,
		MarketSGFuture: resp.MarketSGFuture,
		MarketJPFuture: resp.MarketJPFuture,
		ProgramStatus: func() int32 {
			ps := resp.ProgramStatus
			if ps != nil && ps.Type != nil {
				return int32(*ps.Type)
			}
			return 0
		}(),
		ProgramStatusDesc: func() string {
			ps := resp.ProgramStatus
			if ps != nil {
				return getStr(ps.StrExtDesc)
			}
			return ""
		}(),
		ConnID:       resp.ConnID,
		QotSvrIpAddr: resp.QotSvrIpAddr,
		TrdSvrIpAddr: resp.TrdSvrIpAddr,
	}, nil
}

// GetUserInfo retrieves user information.
func GetUserInfo(ctx context.Context, c *Client) (*UserInfo, error) {
	resp, err := sys.GetUserInfo(ctx, c.inner, nil)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		UserID:               resp.UserID,
		NickName:             resp.NickName,
		AvatarUrl:            resp.AvatarUrl,
		ApiLevel:             resp.ApiLevel,
		IsNeedAgreeDisclaimer: resp.IsNeedAgreeDisclaimer,
		ShQotRight:           resp.ShQotRight,
		SzQotRight:           resp.SzQotRight,
		Extra:                resp.Extra,
		HkQotRight:           resp.HkQotRight,
		UsQotRight:           resp.UsQotRight,
		CnQotRight:           resp.CnQotRight,
		SubQuota:             resp.SubQuota,
		HistoryKLQuota:       resp.HistoryKLQuota,
		HkOptionQotRight:     resp.HkOptionQotRight,
		HasUSOptionQotRight:  resp.HasUSOptionQotRight,
		HkFutureQotRight:     resp.HkFutureQotRight,
		UsFutureQotRight:     resp.UsFutureQotRight,
		UsOptionQotRight:     resp.UsOptionQotRight,
		WebKey:               resp.WebKey,
		WebJumpUrlHead:       resp.WebJumpUrlHead,
		UserAttribution:      resp.UserAttribution,
		UpdateWhatsNew:       resp.UpdateWhatsNew,
		UpdateType:           resp.UpdateType,
		UsIndexQotRight:      resp.UsIndexQotRight,
		UsOtcQotRight:        resp.UsOtcQotRight,
		UsCMEFutureQotRight:  resp.UsCMEFutureQotRight,
		UsCBOTFutureQotRight: resp.UsCBOTFutureQotRight,
		UsNYMEXFutureQotRight: resp.UsNYMEXFutureQotRight,
		UsCOMEXFutureQotRight: resp.UsCOMEXFutureQotRight,
		UsCBOEFutureQotRight:  resp.UsCBOEFutureQotRight,
		SgFutureQotRight:      resp.SgFutureQotRight,
		JpFutureQotRight:      resp.JpFutureQotRight,
		IsAppNNOrMM:           resp.IsAppNNOrMM,
	}, nil
}

// GetDelayStatistics retrieves delay statistics.
func GetDelayStatistics(ctx context.Context, c *Client) (*DelayStatistics, error) {
	resp, err := sys.GetDelayStatistics(ctx, c.inner, nil)
	if err != nil {
		return nil, err
	}

	reqReplyList := make([]ReqReplyStatisticsItem, 0, len(resp.ReqReplyStatisticsList))
	for _, r := range resp.ReqReplyStatisticsList {
		reqReplyList = append(reqReplyList, ReqReplyStatisticsItem{
			ProtoID:      r.ProtoID,
			Count:        r.Count,
			TotalCostAvg: float64(r.TotalCostAvg),
			OpenDCostAvg: float64(r.OpenDCostAvg),
			NetDelayAvg:  float64(r.NetDelayAvg),
			IsLocalReply: r.IsLocalReply,
		})
	}

	placeOrderList := make([]PlaceOrderStatisticsItem, 0, len(resp.PlaceOrderStatisticsList))
	for _, p := range resp.PlaceOrderStatisticsList {
		placeOrderList = append(placeOrderList, PlaceOrderStatisticsItem{
			OrderID:    p.OrderID,
			TotalCost:  float64(p.TotalCost),
			OpenDCost:  float64(p.OpenDCost),
			NetDelay:   float64(p.NetDelay),
			UpdateCost: float64(p.UpdateCost),
		})
	}

	result := &DelayStatistics{
		ReqReplyList:   reqReplyList,
		PlaceOrderList: placeOrderList,
	}

	if len(resp.QotPushStatisticsList) > 0 {
		stats := resp.QotPushStatisticsList[0]
		result.QotPushType = stats.QotPushType
		result.DelayAvg = float64(stats.DelayAvg)
		result.Count = stats.Count
		items := make([]DelayStatisticsItem, 0, len(stats.ItemList))
		for _, item := range stats.ItemList {
			items = append(items, DelayStatisticsItem{
				Begin:           item.Begin,
				End:             item.End,
				Count:           item.Count,
				Proportion:      float64(item.Proportion),
				CumulativeRatio: float64(item.CumulativeRatio),
			})
		}
		result.ItemList = items

		pushList := make([]PushDelayStatisticsItem, 0, len(resp.QotPushStatisticsList))
		for _, s := range resp.QotPushStatisticsList {
			si := make([]DelayStatisticsItem, 0, len(s.ItemList))
			for _, item := range s.ItemList {
				si = append(si, DelayStatisticsItem{
					Begin:           item.Begin,
					End:             item.End,
					Count:           item.Count,
					Proportion:      float64(item.Proportion),
					CumulativeRatio: float64(item.CumulativeRatio),
				})
			}
			pushList = append(pushList, PushDelayStatisticsItem{
				QotPushType: s.QotPushType,
				DelayAvg:    float64(s.DelayAvg),
				Count:       s.Count,
				ItemList:    si,
			})
		}
		result.QotPushList = pushList
	}

	return result, nil
}

// GetTechnicalUnusual queries technical unusual stocks via SkillWrapAPI.
// Deprecated: Removed in Futu v10.6 proto — proto package skillwrapapi no longer exists.
func GetTechnicalUnusual(ctx context.Context, c *Client, req any) (any, error) {
	return nil, fmt.Errorf("GetTechnicalUnusual: removed in Futu v10.6")
}

// GetFinancialUnusual queries financial unusual stocks via SkillWrapAPI.
// Deprecated: Removed in Futu v10.6 proto — proto package skillwrapapi no longer exists.
func GetFinancialUnusual(ctx context.Context, c *Client, req any) (any, error) {
	return nil, fmt.Errorf("GetFinancialUnusual: removed in Futu v10.6")
}

// GetDerivativeUnusual queries derivative unusual stocks via SkillWrapAPI.
// Deprecated: Removed in Futu v10.6 proto — proto package skillwrapapi no longer exists.
func GetDerivativeUnusual(ctx context.Context, c *Client, req any) (any, error) {
	return nil, fmt.Errorf("GetDerivativeUnusual: removed in Futu v10.6")
}

// TestCmd sends a test command to OpenD for internal diagnostics.
// Deprecated: Removed in Futu v10.6 proto — proto package testcmd no longer exists.
func TestCmd(ctx context.Context, c *Client, cmd string, params ...string) (*TestCmdResult, error) {
	return nil, fmt.Errorf("TestCmd: removed in Futu v10.6")
}

// Verification submits a verification request (e.g., SMS or email verification).
func Verification(ctx context.Context, c *Client, req *sys.VerificationRequest) error {
	if req == nil {
		return fmt.Errorf("Verification: req is required")
	}
	return sys.Verification(ctx, c.inner, req)
}

// GetUsedQuota retrieves the current quota usage for subscriptions and historical K-line requests.
func GetUsedQuota(ctx context.Context, c *Client) (*UsedQuotaInfo, error) {
	resp, err := sys.GetUsedQuota(ctx, c.inner)
	if err != nil {
		return nil, err
	}
	return &UsedQuotaInfo{
		UsedSubQuota:   resp.UsedSubQuota,
		UsedKLineQuota: resp.UsedKLineQuota,
	}, nil
}
