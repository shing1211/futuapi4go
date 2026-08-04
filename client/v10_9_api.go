package client

import (
	"context"
	"fmt"

	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/pb/qotfiltercompetition"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontract"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractcategory"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractcombolist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractcomborfq"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontracteventlist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractkline"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractmilestonelist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractorderbook"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractserieslist"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractsnapshot"
	"github.com/shing1211/futuapi4go/pkg/pb/qotgeteventcontractticker"
	"github.com/shing1211/futuapi4go/pkg/pb/qotrequesthistoryeventcontractkl"
	"github.com/shing1211/futuapi4go/pkg/pb/qotsubeventcontract"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

// v10.9+ Event Contract (Prediction Market) API wrappers.
//
// Event Contracts (Moomoo US Prediction) trade YES/NO binary outcomes on
// future events such as elections, economic data, and sporting events.

// FilterCompetition returns the available competition filters.
func FilterCompetition(ctx context.Context, c *Client, req *qotfiltercompetition.C2S) (*qotfiltercompetition.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("FilterCompetition: req is nil")
	}
	return qot.FilterCompetition(ctx, c.inner, req)
}

// GetEventContractCategory returns the top-level EC categories.
func GetEventContractCategory(ctx context.Context, c *Client, req *qotgeteventcontractcategory.C2S) (*qotgeteventcontractcategory.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractCategory: req is nil")
	}
	return qot.GetEventContractCategory(ctx, c.inner, req)
}

// GetEventContractSeriesList returns Series under a category/tag filter.
func GetEventContractSeriesList(ctx context.Context, c *Client, req *qotgeteventcontractserieslist.C2S) (*qotgeteventcontractserieslist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractSeriesList: req is nil")
	}
	return qot.GetEventContractSeriesList(ctx, c.inner, req)
}

// GetEventContractEventList returns Events under a Series.
func GetEventContractEventList(ctx context.Context, c *Client, req *qotgeteventcontracteventlist.C2S) (*qotgeteventcontracteventlist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractEventList: req is nil")
	}
	return qot.GetEventContractEventList(ctx, c.inner, req)
}

// GetEventContract returns Contracts under an Event.
func GetEventContract(ctx context.Context, c *Client, req *qotgeteventcontract.C2S) (*qotgeteventcontract.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContract: req is nil")
	}
	return qot.GetEventContract(ctx, c.inner, req)
}

// GetEventContractMilestoneList returns the EC milestone list.
func GetEventContractMilestoneList(ctx context.Context, c *Client, req *qotgeteventcontractmilestonelist.C2S) (*qotgeteventcontractmilestonelist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractMilestoneList: req is nil")
	}
	return qot.GetEventContractMilestoneList(ctx, c.inner, req)
}

// GetEventContractSnapshot returns a batch of EC snapshots.
func GetEventContractSnapshot(ctx context.Context, c *Client, req *qotgeteventcontractsnapshot.C2S) (*qotgeteventcontractsnapshot.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractSnapshot: req is nil")
	}
	return qot.GetEventContractSnapshot(ctx, c.inner, req)
}

// GetEventContractOrderBook returns the EC order book snapshot.
func GetEventContractOrderBook(ctx context.Context, c *Client, req *qotgeteventcontractorderbook.C2S) (*qotgeteventcontractorderbook.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractOrderBook: req is nil")
	}
	return qot.GetEventContractOrderBook(ctx, c.inner, req)
}

// GetEventContractKline returns the EC K-line snapshot.
func GetEventContractKline(ctx context.Context, c *Client, req *qotgeteventcontractkline.C2S) (*qotgeteventcontractkline.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractKline: req is nil")
	}
	return qot.GetEventContractKline(ctx, c.inner, req)
}

// GetEventContractTicker returns the EC tick-by-tick snapshot.
func GetEventContractTicker(ctx context.Context, c *Client, req *qotgeteventcontractticker.C2S) (*qotgeteventcontractticker.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractTicker: req is nil")
	}
	return qot.GetEventContractTicker(ctx, c.inner, req)
}

// RequestHistoryEventContractKL pulls historical EC K-line data.
func RequestHistoryEventContractKL(ctx context.Context, c *Client, req *qotrequesthistoryeventcontractkl.C2S) (*qotrequesthistoryeventcontractkl.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("RequestHistoryEventContractKL: req is nil")
	}
	return qot.RequestHistoryEventContractKL(ctx, c.inner, req)
}

// GetEventContractComboList returns Events combinable into Combo positions.
func GetEventContractComboList(ctx context.Context, c *Client, req *qotgeteventcontractcombolist.C2S) (*qotgeteventcontractcombolist.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractComboList: req is nil")
	}
	return qot.GetEventContractComboList(ctx, c.inner, req)
}

// GetEventContractComboRfq requests a quote for a Combo leg combination.
func GetEventContractComboRfq(ctx context.Context, c *Client, req *qotgeteventcontractcomborfq.C2S) (*qotgeteventcontractcomborfq.S2C, error) {
	if req == nil {
		return nil, fmt.Errorf("GetEventContractComboRfq: req is nil")
	}
	return qot.GetEventContractComboRfq(ctx, c.inner, req)
}

// SubEventContract subscribes / unsubscribes Event Contract real-time data.
func SubEventContract(ctx context.Context, c *Client, req *qotsubeventcontract.C2S) error {
	if req == nil {
		return fmt.Errorf("SubEventContract: req is nil")
	}
	return qot.SubEventContract(ctx, c.inner, req)
}

// NewECSecurity constructs a Security message for an Event Contract instrument.
// Example: client.NewECSecurity("EC.KXWCADVANCE-26JUL14FRAESP-FRA").
func NewECSecurity(code string) *qotcommon.Security {
	return qot.BuildECSecurity(code)
}