package main

import (
	"log"
	"os"
)

func preflight() bool {
	os.Mkdir(".data", 0777)
	return true
}

func main() {
	wd, err := os.Getwd()
	if err == nil {
		log.Print("out", "working directory", wd)
	} else {
		log.Fatal("out", "if I cannot Getcwd then wtf")
	}

	server := Server{}
	if !server.Start() {
		os.Exit(-1)
	}
}
