package main

import (
	"log"

	"gitee.com/shing1211/futuapi4go/simulator"
)

func main() {
	srv := simulator.New("127.0.0.1:11111")
	srv.RegisterDefaultHandlers()
	srv.RegisterQotHandlers()
	srv.RegisterTrdHandlers()

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
	defer srv.Stop()

	log.Println("Futu OpenD Simulator started on 127.0.0.1:11111")
	log.Println("Press Ctrl+C to stop")

	select {}
}
