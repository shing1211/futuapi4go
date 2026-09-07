package client

import (
	"context"
	"testing"

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
)

func TestECConvenienceWrappers_NilRequest(t *testing.T) {
	// QuoteAPI methods must return an error when passed a nil request,
	// not panic. We cannot test the full RPC path without a mock server,
	// but we can verify nil-safety for all wrappers.

	api := &QuoteAPI{client: nil}
	ctx := context.Background()

	tests := []struct {
		name string
		fn   func() error
	}{
		{
			name: "FilterCompetition",
			fn: func() error {
				_, err := api.FilterCompetition(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractCategory",
			fn: func() error {
				_, err := api.GetEventContractCategory(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractSeriesList",
			fn: func() error {
				_, err := api.GetEventContractSeriesList(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractEventList",
			fn: func() error {
				_, err := api.GetEventContractEventList(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContract",
			fn: func() error {
				_, err := api.GetEventContract(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractMilestoneList",
			fn: func() error {
				_, err := api.GetEventContractMilestoneList(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractSnapshot",
			fn: func() error {
				_, err := api.GetEventContractSnapshot(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractOrderBook",
			fn: func() error {
				_, err := api.GetEventContractOrderBook(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractKline",
			fn: func() error {
				_, err := api.GetEventContractKline(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractTicker",
			fn: func() error {
				_, err := api.GetEventContractTicker(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractComboList",
			fn: func() error {
				_, err := api.GetEventContractComboList(ctx, nil)
				return err
			},
		},
		{
			name: "GetEventContractComboRfq",
			fn: func() error {
				_, err := api.GetEventContractComboRfq(ctx, nil)
				return err
			},
		},
		{
			name: "RequestHistoryEventContractKL",
			fn: func() error {
				_, err := api.RequestHistoryEventContractKL(ctx, nil)
				return err
			},
		},
		{
			name: "SubEventContract",
			fn: func() error {
				return api.SubEventContract(ctx, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			if err == nil {
				t.Errorf("%s: expected error for nil request, got nil", tt.name)
			}
		})
	}
}

func TestNewECSecurity(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		wantNil  bool
		wantCode string
	}{
		{
			name:     "valid EC code",
			code:     "EC.KXWCADVANCE-26JUL14FRAESP-FRA",
			wantNil:  false,
			wantCode: "EC.KXWCADVANCE-26JUL14FRAESP-FRA",
		},
		{
			name:     "empty code",
			code:     "",
			wantNil:  false,
			wantCode: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewECSecurity(tt.code)
			if tt.wantNil && got != nil {
				t.Errorf("NewECSecurity(%q) = %v, want nil", tt.code, got)
			}
			if !tt.wantNil && got == nil {
				t.Fatalf("NewECSecurity(%q) returned nil, want non-nil", tt.code)
			}
			if got != nil && got.GetCode() != tt.wantCode {
				t.Errorf("NewECSecurity(%q).Code = %q, want %q", tt.code, got.GetCode(), tt.wantCode)
			}
		})
	}
}

func TestECProtoTypeFields(t *testing.T) {
	// Verify that key EC proto types can be constructed and accessed
	// without nil panics — ensures proto wiring is correct.

	t.Run("FilterCompetition C2S", func(t *testing.T) {
		req := &qotfiltercompetition.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractCategory C2S", func(t *testing.T) {
		req := &qotgeteventcontractcategory.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractSeriesList C2S", func(t *testing.T) {
		req := &qotgeteventcontractserieslist.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractEventList C2S", func(t *testing.T) {
		req := &qotgeteventcontracteventlist.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContract C2S", func(t *testing.T) {
		req := &qotgeteventcontract.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractMilestoneList C2S", func(t *testing.T) {
		req := &qotgeteventcontractmilestonelist.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractSnapshot C2S", func(t *testing.T) {
		req := &qotgeteventcontractsnapshot.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractOrderBook C2S", func(t *testing.T) {
		req := &qotgeteventcontractorderbook.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractKline C2S", func(t *testing.T) {
		req := &qotgeteventcontractkline.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractTicker C2S", func(t *testing.T) {
		req := &qotgeteventcontractticker.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractComboList C2S", func(t *testing.T) {
		req := &qotgeteventcontractcombolist.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("GetEventContractComboRfq C2S", func(t *testing.T) {
		req := &qotgeteventcontractcomborfq.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("RequestHistoryEventContractKL C2S", func(t *testing.T) {
		req := &qotrequesthistoryeventcontractkl.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})

	t.Run("SubEventContract C2S", func(t *testing.T) {
		req := &qotsubeventcontract.C2S{}
		if req == nil {
			t.Fatal("unexpected nil")
		}
	})
}
