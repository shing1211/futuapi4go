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

	futuapi "github.com/shing1211/futuapi4go/internal/client"
	"github.com/shing1211/futuapi4go/pkg/pb/qotcommon"
	"github.com/shing1211/futuapi4go/pkg/qot"
)

func main() {
	fmt.Println("=== futuapi4go Demo ===")

	cli := futuapi.New()
	defer cli.Close()

	err := cli.Connect("127.0.0.1:11111")
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	fmt.Printf("Connected! ConnID=%d\n", cli.GetConnID())

	market1 := int32(qotcommon.QotMarket_QotMarket_HK_Security)
	code1 := "00700"
	market2 := int32(qotcommon.QotMarket_QotMarket_US_Security)
	code2 := "BABA"
	market3 := int32(qotcommon.QotMarket_QotMarket_CNSH_Security)
	code3 := "600519"

	securities := []*qotcommon.Security{
		{Market: &market1, Code: &code1},
		{Market: &market2, Code: &code2},
		{Market: &market3, Code: &code3},
	}

	result, err := qot.GetBasicQot(context.Background(),cli, securities)
	if err != nil {
		log.Fatalf("GetBasicQot failed: %v", err)
	}

	fmt.Println("\n--- Market Data ---")
	for _, bq := range result {
		sec := bq.Security
		fmt.Printf("%s %s: CurPrice=%.2f Open=%.2f High=%.2f Low=%.2f Vol=%d\n",
			sec.GetCode(),
			bq.Name,
			bq.CurPrice,
			bq.OpenPrice,
			bq.HighPrice,
			bq.LowPrice,
			bq.Volume,
		)
	}

	fmt.Println("\nDone!")
}
