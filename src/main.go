package main

import (
	log "github.com/sirupsen/logrus"
	"main/src/server"
	"main/src/service"
)

func main() {
	go func() {
		if err := service.ListenPriceUpdates(); err != nil {
			log.Fatal(err)
		}
	}()
	httpServer := server.NewServer(1323)
	httpServer.Start()
}
