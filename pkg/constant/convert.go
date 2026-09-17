package constant

import "fmt"

// ToSubType maps a KLType to the SubType that subscribes to it.
//
// KLType and SubType are separate enums with different numbering: KLType_K_1Min
// is 1 while SubType_K_1Min is 11, KLType_K_Day is 2 while SubType_K_Day is 6,
// and SubType has no value 3 at all. The two therefore can never be converted by
// a numeric cast - doing so subscribes to the wrong stream for every input, which
// is exactly how the removed pkg/history.Streamer was broken.
//
// This is the single place that correspondence is expressed. pkg/push/chan and
// the history downloader both route through it.
func (k KLType) ToSubType() (SubType, error) {
	switch k {
	case KLType_K_1Min:
		return SubType_K_1Min, nil
	case KLType_K_3Min:
		return SubType_K_3Min, nil
	case KLType_K_5Min:
		return SubType_K_5Min, nil
	case KLType_K_10Min:
		return SubType_K_10Min, nil
	case KLType_K_15Min:
		return SubType_K_15Min, nil
	case KLType_K_30Min:
		return SubType_K_30Min, nil
	case KLType_K_60Min:
		return SubType_K_60Min, nil
	case KLType_K_120Min:
		return SubType_K_120Min, nil
	case KLType_K_180Min:
		return SubType_K_180Min, nil
	case KLType_K_240Min:
		return SubType_K_240Min, nil
	case KLType_K_Day:
		return SubType_K_Day, nil
	case KLType_K_Week:
		return SubType_K_Week, nil
	case KLType_K_Month:
		return SubType_K_Month, nil
	case KLType_K_Quarter:
		return SubType_K_Quarter, nil
	case KLType_K_Year:
		return SubType_K_Year, nil
	default:
		return SubType_None, fmt.Errorf("KLType %d has no corresponding SubType", int32(k))
	}
}

// ToSubType maps a SubType back to its KLType. It succeeds only for the K-line
// SubTypes; every other subscription type returns an error, since a quote or
// order-book subscription is not a bar interval.
func (s SubType) ToKLType() (KLType, error) {
	switch s {
	case SubType_K_1Min:
		return KLType_K_1Min, nil
	case SubType_K_3Min:
		return KLType_K_3Min, nil
	case SubType_K_5Min:
		return KLType_K_5Min, nil
	case SubType_K_10Min:
		return KLType_K_10Min, nil
	case SubType_K_15Min:
		return KLType_K_15Min, nil
	case SubType_K_30Min:
		return KLType_K_30Min, nil
	case SubType_K_60Min:
		return KLType_K_60Min, nil
	case SubType_K_120Min:
		return KLType_K_120Min, nil
	case SubType_K_180Min:
		return KLType_K_180Min, nil
	case SubType_K_240Min:
		return KLType_K_240Min, nil
	case SubType_K_Day:
		return KLType_K_Day, nil
	case SubType_K_Week:
		return KLType_K_Week, nil
	case SubType_K_Month:
		return KLType_K_Month, nil
	case SubType_K_Quarter:
		return KLType_K_Quarter, nil
	case SubType_K_Year:
		return KLType_K_Year, nil
	default:
		return KLType_None, fmt.Errorf("SubType %d is not a K-line type", int32(s))
	}
}
