package client

import (
	"context"

	"github.com/shing1211/futuapi4go/pkg/pb/skillwrapapi"
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
				return ps.GetStrExtDesc()
			}
			return ""
		}(),
	}, nil
}

// GetUserInfo retrieves user information.
func GetUserInfo(ctx context.Context, c *Client) (*UserInfo, error) {
	resp, err := sys.GetUserInfo(ctx, c.inner)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		UserID:    resp.UserID,
		NickName:  resp.NickName,
		AvatarUrl: resp.AvatarUrl,
		ApiLevel:  resp.ApiLevel,
	}, nil
}

// GetDelayStatistics retrieves delay statistics.
func GetDelayStatistics(ctx context.Context, c *Client) (*DelayStatistics, error) {
	resp, err := sys.GetDelayStatistics(ctx, c.inner)
	if err != nil {
		return nil, err
	}

	if len(resp.QotPushStatisticsList) == 0 {
		return &DelayStatistics{}, nil
	}

	stats := resp.QotPushStatisticsList[0]
	items := make([]DelayStatisticsItem, 0, len(stats.ItemList))
	for _, item := range stats.ItemList {
		items = append(items, DelayStatisticsItem{
			Begin:           item.GetBegin(),
			End:             item.GetEnd(),
			Count:           item.GetCount(),
			Proportion:      float64(item.GetProportion()),
			CumulativeRatio: float64(item.GetCumulativeRatio()),
		})
	}

	reqReplyList := make([]ReqReplyStatisticsItem, 0, len(resp.ReqReplyStatisticsList))
	for _, r := range resp.ReqReplyStatisticsList {
		reqReplyList = append(reqReplyList, ReqReplyStatisticsItem{
			ProtoID:      r.GetProtoID(),
			Count:        r.GetCount(),
			TotalCostAvg: float64(r.GetTotalCostAvg()),
			OpenDCostAvg: float64(r.GetOpenDCostAvg()),
			NetDelayAvg:  float64(r.GetNetDelayAvg()),
			IsLocalReply: r.GetIsLocalReply(),
		})
	}

	placeOrderList := make([]PlaceOrderStatisticsItem, 0, len(resp.PlaceOrderStatisticsList))
	for _, p := range resp.PlaceOrderStatisticsList {
		placeOrderList = append(placeOrderList, PlaceOrderStatisticsItem{
			OrderID:    p.GetOrderID(),
			TotalCost:  float64(p.GetTotalCost()),
			OpenDCost:  float64(p.GetOpenDCost()),
			NetDelay:   float64(p.GetNetDelay()),
			UpdateCost: float64(p.GetUpdateCost()),
		})
	}

	return &DelayStatistics{
		QotPushType:    stats.GetQotPushType(),
		DelayAvg:       float64(stats.GetDelayAvg()),
		Count:          stats.GetCount(),
		ItemList:       items,
		ReqReplyList:   reqReplyList,
		PlaceOrderList: placeOrderList,
	}, nil
}

// GetTechnicalUnusual queries technical unusual stocks via SkillWrapAPI.
func GetTechnicalUnusual(ctx context.Context, c *Client, req *skillwrapapi.TechnicalUnusualReq) (*skillwrapapi.TechnicalUnusualRsp, error) {
	return sys.GetTechnicalUnusual(ctx, c.inner, req)
}

// GetFinancialUnusual queries financial unusual stocks via SkillWrapAPI.
func GetFinancialUnusual(ctx context.Context, c *Client, req *skillwrapapi.FinancialUnusualReq) (*skillwrapapi.FinancialUnusualRsp, error) {
	return sys.GetFinancialUnusual(ctx, c.inner, req)
}

// GetDerivativeUnusual queries derivative unusual stocks via SkillWrapAPI.
func GetDerivativeUnusual(ctx context.Context, c *Client, req *skillwrapapi.DerivativeUnusualReq) (*skillwrapapi.DerivativeUnusualRsp, error) {
	return sys.GetDerivativeUnusual(ctx, c.inner, req)
}

// TestCmd sends a test command to OpenD for internal diagnostics.
func TestCmd(ctx context.Context, c *Client, cmd string, params ...string) (*TestCmdResult, error) {
	p := ""
	if len(params) > 0 {
		p = params[0]
	}
	resp, err := sys.TestCmd(ctx, c.inner, &sys.TestCmdRequest{
		Cmd:    cmd,
		Params: p,
	})
	if err != nil {
		return nil, err
	}
	return &TestCmdResult{
		Cmd:    resp.Cmd,
		Result: resp.Result,
	}, nil
}

// Verification submits a verification request (e.g., SMS or email verification).
func Verification(ctx context.Context, c *Client, req *sys.VerificationRequest) error {
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
