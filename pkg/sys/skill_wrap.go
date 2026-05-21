package sys

import (
	"context"
	"fmt"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
)

const (
	ProtoID_SkillWrapAPI = 8001
)

// GetTechnicalUnusual queries technical unusual stocks.
// Deprecated: Removed in Futu v10.6 proto — proto package skillwrapapi no longer exists.
func GetTechnicalUnusual(ctx context.Context, c *futuapi.Client, req any) (any, error) {
	// The underlying SkillWrapAPI proto was removed in Futu v10.6.
	return nil, fmt.Errorf("GetTechnicalUnusual: removed in Futu v10.6")
}

// GetFinancialUnusual queries financial unusual stocks.
// Deprecated: Removed in Futu v10.6 proto — proto package skillwrapapi no longer exists.
func GetFinancialUnusual(ctx context.Context, c *futuapi.Client, req any) (any, error) {
	// The underlying SkillWrapAPI proto was removed in Futu v10.6.
	return nil, fmt.Errorf("GetFinancialUnusual: removed in Futu v10.6")
}

// GetDerivativeUnusual queries derivative unusual stocks.
// Deprecated: Removed in Futu v10.6 proto — proto package skillwrapapi no longer exists.
func GetDerivativeUnusual(ctx context.Context, c *futuapi.Client, req any) (any, error) {
	// The underlying SkillWrapAPI proto was removed in Futu v10.6.
	return nil, fmt.Errorf("GetDerivativeUnusual: removed in Futu v10.6")
}
