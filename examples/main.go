package main

import (
	"fmt"
	"log"

	futuapi "gitee.com/shing1211/futuapi4go/client"
	"gitee.com/shing1211/futuapi4go/qot"
	"gitee.com/shing1211/futuapi4go/pb/qotcommon"
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

	result, err := qot.GetBasicQot(cli, securities)
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
