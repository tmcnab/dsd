package main

import "log"

func main() {
	engine := Engine{}
	server := Server{}

	err := server.Start(&engine)
	if err != nil {
		log.Fatal("couldn't start the server")
	}
}
