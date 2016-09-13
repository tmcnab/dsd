package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	engine := Engine{}
	cluster := Cluster{}
	server := Server{}

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
		log.Fatal("couldn't start the server", err.Error())
	}
}
