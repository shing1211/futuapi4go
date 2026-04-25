package option

import (
	"testing"
	"time"
)

func TestParseCode(t *testing.T) {
	tests := []struct {
		name        string
		code       string
		wantExpiry time.Time
		wantStrike float64
		wantType   OptionType
		wantErr    bool
	}{
		{
			name:        "HK Put",
			code:       "24011900700P",
			wantExpiry: time.Date(2024, 1, 19, 0, 0, 0, 0, time.UTC),
			wantStrike: 700,
			wantType:   OptionTypePut,
		},
		{
			name:        "HK Call",
			code:       "24021600180C",
			wantExpiry: time.Date(2024, 2, 16, 0, 0, 0, 0, time.UTC),
			wantStrike: 180,
			wantType:   OptionTypeCall,
		},
		{
			name:    "empty",
			code:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCode(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if !got.Expiry.Equal(tt.wantExpiry) {
				t.Errorf("ParseCode() expiry = %v, want %v", got.Expiry, tt.wantExpiry)
			}
			if got.Strike != tt.wantStrike {
				t.Errorf("ParseCode() strike = %v, want %v", got.Strike, tt.wantStrike)
			}
			if got.Type != tt.wantType {
				t.Errorf("ParseCode() type = %v, want %v", got.Type, tt.wantType)
			}
		})
	}
}

func TestFindAtm(t *testing.T) {
	options := []*OptionCode{
		{Code: "24021600150C", Strike: 150, Type: OptionTypeCall},
		{Code: "24021600160C", Strike: 160, Type: OptionTypeCall},
		{Code: "24021600170C", Strike: 170, Type: OptionTypeCall},
		{Code: "24021600150P", Strike: 150, Type: OptionTypePut},
		{Code: "24021600160P", Strike: 160, Type: OptionTypePut},
		{Code: "24021600170P", Strike: 170, Type: OptionTypePut},
	}

	spot := 165.0
	atms := FindAtm(options, spot)

	if atms.Call == nil || atms.Call.Strike != 160 {
		t.Errorf("FindAtm() call = %v, want 160", atms.Call.Strike)
	}
	if atms.Put == nil || atms.Put.Strike != 160 {
		t.Errorf("FindAtm() put = %v, want 160", atms.Put.Strike)
	}
}

func TestFilterByExpiry(t *testing.T) {
	options := []*OptionCode{
		{Code: "24011900700P", Expiry: time.Date(2024, 1, 19, 0, 0, 0, 0, time.UTC)},
		{Code: "24011900700C", Expiry: time.Date(2024, 1, 19, 0, 0, 0, 0, time.UTC)},
		{Code: "24021600700P", Expiry: time.Date(2024, 2, 16, 0, 0, 0, 0, time.UTC)},
	}

	expiry := time.Date(2024, 1, 19, 0, 0, 0, 0, time.UTC)
	result := FilterByExpiry(options, expiry)

	if len(result) != 2 {
		t.Errorf("FilterByExpiry() returned %d, want 2", len(result))
	}
}