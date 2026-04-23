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

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"

	"github.com/shing1211/futuapi4go/cmd/simulator"
)

func main() {
	fmt.Println("=== Starting Simulator ===")

	srv := simulator.New("127.0.0.1:11112")
	srv.RegisterDefaultHandlers()
	srv.RegisterQotHandlers()
	srv.AddSecurity(int32(qotcommon.QotMarket_QotMarket_HK_Security), "00700")

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start simulator: %v", err)
	}
	defer srv.Stop()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Simulator running on 127.0.0.1:11112")

	fmt.Println("=== Connecting SDK Client ===")
	cli := futuapi.New()
	defer cli.Close()

	if err := cli.Connect("127.0.0.1:11112"); err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	fmt.Printf("Connected! ConnID=%d\n", cli.GetConnID())

	market := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	securities := []*qotcommon.Security{
		{Market: &market, Code: func() *string { s := "00700"; return &s }()},
	}

	fmt.Println("=== Calling GetBasicQot ===")
	result, err := qot.GetBasicQot(context.Background(), cli, securities)
	if err != nil {
		log.Fatalf("GetBasicQot failed: %v", err)
	}

	fmt.Println("--- Results ---")
	for _, bq := range result {
		fmt.Printf("%s: CurPrice=%.2f Volume=%d\n", bq.Security.GetCode(), bq.CurPrice, bq.Volume)
	}

	fmt.Println("\n=== Success! SDK works with Simulator ===")
}
