package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"./src/dsd"
)

func main() {
	engine := dsd.Engine{}
	cluster := dsd.Cluster{}
	server := dsd.Server{}

	err := server.Start(&engine)
	if err == nil {
		channel := make(chan os.Signal, 2)
		signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-channel
			server.Stop()
			cluster.Stop()
		}()
	} else {
		log.Fatal("err [main] couldn't start the server: " + err.Error())
	}
}
