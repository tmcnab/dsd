package main

import (
	"os"
	"os/signal"
	"syscall"

	"./src/dsd"
)

func main() {
	engine := dsd.NewEngine()
	cluster := dsd.NewCluster(engine)
	server := dsd.NewServer(engine)

	channel := make(chan os.Signal, 2)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		server.Stop()
		cluster.Stop()
	}()

	server.Start()
}
