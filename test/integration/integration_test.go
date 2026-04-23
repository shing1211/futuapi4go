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

package integration

import (
	"context"
	"os"
	"testing"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
	"github.com/shing1211/futuapi4go/pkg/sys"
)

// getTestAddr returns the Futu OpenD address for testing.
// Uses FUTU_TEST_ADDR env var, defaults to 127.0.0.1:11111
func getTestAddr() string {
	addr := os.Getenv("FUTU_TEST_ADDR")
	if addr == "" {
		addr = "127.0.0.1:11111"
	}
	return addr
}

// skipIfNoOpenD skips the test if OpenD is not available
func skipIfNoOpenD(t *testing.T) {
	addr := getTestAddr()
	if os.Getenv("FUTU_INTEGRATION_TESTS") != "1" {
		t.Skip("skipping integration test: set FUTU_INTEGRATION_TESTS=1 to enable")
	}
	t.Logf("Testing against %s", addr)
}

func TestIntegrationConnect(t *testing.T) {
	skipIfNoOpenD(t)

	client := futuapi.New()
	defer client.Close()

	addr := getTestAddr()
	if err := client.Connect(addr); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	if !client.IsConnected() {
		t.Error("client should be connected after Connect()")
	}

	if client.GetConnID() == 0 {
		t.Error("ConnID should not be 0 after connection")
	}

	if client.GetServerVer() == 0 {
		t.Error("ServerVer should not be 0 after connection")
	}

	t.Logf("Connected: ConnID=%d, ServerVer=%d", client.GetConnID(), client.GetServerVer())
}

func TestIntegrationEnsureConnected(t *testing.T) {
	skipIfNoOpenD(t)

	client := futuapi.New()
	defer client.Close()

	// Should fail before connect
	if err := client.EnsureConnected(); err == nil {
		t.Error("EnsureConnected should fail before Connect()")
	}

	addr := getTestAddr()
	if err := client.Connect(addr); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	// Should succeed after connect
	if err := client.EnsureConnected(); err != nil {
		t.Errorf("EnsureConnected should succeed after Connect(): %v", err)
	}
}

func TestIntegrationGetGlobalState(t *testing.T) {
	skipIfNoOpenD(t)

	client := futuapi.New()
	defer client.Close()

	addr := getTestAddr()
	if err := client.Connect(addr); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	state, err := sys.GetGlobalState(client)
	if err != nil {
		t.Fatalf("GetGlobalState failed: %v", err)
	}

	// Verify we got a response (values may vary based on OpenD state)
	t.Logf("GlobalState: ServerVer=%d, BuildNo=%d, QotLogined=%v, TrdLogined=%v",
		state.ServerVer, state.ServerBuildNo, state.QotLogined, state.TrdLogined)

	if state.ServerVer == 0 && state.ServerBuildNo == 0 {
		t.Error("ServerVer and ServerBuildNo should not both be 0")
	}
}

func TestIntegrationGetBasicQot(t *testing.T) {
	skipIfNoOpenD(t)

	client := futuapi.New()
	defer client.Close()

	addr := getTestAddr()
	if err := client.Connect(addr); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	securities := []*qotcommon.Security{
		{Market: &hkMarket, Code: &code},
	}

	quotes, err := qot.GetBasicQot(context.Background(), client, securities)
	if err != nil {
		t.Fatalf("GetBasicQot failed: %v", err)
	}

	if len(quotes) == 0 {
		t.Fatal("GetBasicQot returned empty result")
	}

	q := quotes[0]
	t.Logf("Got quote for %s: %s, Price=%.2f", q.Security.GetCode(), q.Name, q.CurPrice)

	if q.Security.GetCode() != code {
		t.Errorf("expected code %s, got %s", code, q.Security.GetCode())
	}
}

func TestIntegrationSubscribe(t *testing.T) {
	skipIfNoOpenD(t)

	client := futuapi.New()
	defer client.Close()

	addr := getTestAddr()
	if err := client.Connect(addr); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code := "00700"
	security := &qotcommon.Security{Market: &hkMarket, Code: &code}

	subReq := &qot.SubscribeRequest{
		SecurityList:     []*qotcommon.Security{security},
		SubTypeList:      []qot.SubType{qot.SubType_Basic},
		IsSubOrUnSub:     true,
		IsRegOrUnRegPush: true,
	}

	subResp, err := qot.Subscribe(client, subReq)
	if err != nil {
		t.Fatalf("Subscribe failed: %v", err)
	}

	t.Logf("Subscribe result: RetType=%d, RetMsg=%s", subResp.RetType, subResp.RetMsg)

	if subResp.RetType != 0 {
		t.Errorf("expected RetType 0 (success), got %d", subResp.RetType)
	}
}

func TestIntegrationMultipleAPIs(t *testing.T) {
	skipIfNoOpenD(t)

	client := futuapi.New()
	defer client.Close()

	addr := getTestAddr()
	if err := client.Connect(addr); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	// Test a sequence of APIs
	apis := []struct {
		name string
		fn   func() error
	}{
		{
			name: "GetGlobalState",
			fn: func() error {
				_, err := sys.GetGlobalState(client)
				return err
			},
		},
		{
			name: "GetBasicQot",
			fn: func() error {
				hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
				code := "00700"
				securities := []*qotcommon.Security{
					{Market: &hkMarket, Code: &code},
				}
				_, err := qot.GetBasicQot(context.Background(), client, securities)
				return err
			},
		},
		{
			name: "Subscribe",
			fn: func() error {
				hkMarket := int32(qotcommon.QotMarket_QotMarket_HK_Security)
				code := "00700"
				security := &qotcommon.Security{Market: &hkMarket, Code: &code}
				subReq := &qot.SubscribeRequest{
					SecurityList:     []*qotcommon.Security{security},
					SubTypeList:      []qot.SubType{qot.SubType_Basic},
					IsSubOrUnSub:     true,
					IsRegOrUnRegPush: true,
				}
				_, err := qot.Subscribe(client, subReq)
				return err
			},
		},
	}

	for _, api := range apis {
		t.Run(api.name, func(t *testing.T) {
			if err := api.fn(); err != nil {
				t.Errorf("%s failed: %v", api.name, err)
			}
		})
	}
}

func TestIntegrationContextCancellation(t *testing.T) {
	skipIfNoOpenD(t)

	client := futuapi.New()
	defer client.Close()

	addr := getTestAddr()
	if err := client.Connect(addr); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	// Test WithContext
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newClient := client.WithContext(ctx)
	if newClient == nil {
		t.Fatal("WithContext returned nil")
	}

	if newClient.Context() != ctx {
		t.Error("WithContext did not set context correctly")
	}
}
