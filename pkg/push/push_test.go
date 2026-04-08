package push

import (
	"testing"
)

func TestParseUpdateBasicQotInvalidData(t *testing.T) {
	// Empty data returns nil result, not error
	result, err := ParseUpdateBasicQot([]byte{})
	if err != nil {
		t.Errorf("ParseUpdateBasicQot should not error on empty data, got: %v", err)
	}
	if result != nil {
		t.Error("ParseUpdateBasicQot should return nil for empty data")
	}

	_, err = ParseUpdateBasicQot([]byte{0x00, 0x01, 0x02})
	if err == nil {
		t.Error("ParseUpdateBasicQot should fail with invalid protobuf data")
	}
}

func TestParseUpdateKLInvalidData(t *testing.T) {
	_, err := ParseUpdateKL([]byte{})
	if err == nil {
		t.Error("ParseUpdateKL should fail with empty data")
	}
}

func TestParseUpdateOrderBookInvalidData(t *testing.T) {
	_, err := ParseUpdateOrderBook([]byte{})
	if err == nil {
		t.Error("ParseUpdateOrderBook should fail with empty data")
	}
}

func TestParseUpdateTickerInvalidData(t *testing.T) {
	_, err := ParseUpdateTicker([]byte{})
	if err == nil {
		t.Error("ParseUpdateTicker should fail with empty data")
	}
}

func TestParseUpdateRTInvalidData(t *testing.T) {
	_, err := ParseUpdateRT([]byte{})
	if err == nil {
		t.Error("ParseUpdateRT should fail with empty data")
	}
}

func TestParseUpdateBrokerInvalidData(t *testing.T) {
	_, err := ParseUpdateBroker([]byte{})
	if err == nil {
		t.Error("ParseUpdateBroker should fail with empty data")
	}
}

func TestParseUpdatePriceReminderInvalidData(t *testing.T) {
	_, err := ParseUpdatePriceReminder([]byte{})
	if err == nil {
		t.Error("ParseUpdatePriceReminder should fail with empty data")
	}
}

func TestParseSystemNotifyInvalidData(t *testing.T) {
	_, err := ParseSystemNotify([]byte{})
	if err == nil {
		t.Error("ParseSystemNotify should fail with empty data")
	}
}

func TestParseUpdateOrderInvalidData(t *testing.T) {
	_, err := ParseUpdateOrder([]byte{})
	if err == nil {
		t.Error("ParseUpdateOrder should fail with empty data")
	}
}

func TestParseUpdateOrderFillInvalidData(t *testing.T) {
	_, err := ParseUpdateOrderFill([]byte{})
	if err == nil {
		t.Error("ParseUpdateOrderFill should fail with empty data")
	}
}

func TestParseTrdNotifyInvalidData(t *testing.T) {
	_, err := ParseTrdNotify([]byte{})
	if err == nil {
		t.Error("ParseTrdNotify should fail with empty data")
	}
}
