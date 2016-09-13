package main

import (
	"encoding/json"
	"log"
	"net"
)

// Server listens for requests and services them
type Server struct {
	acceptingNewConnections bool         // whether or not the server is af new connections
	clientListener          net.Listener // client listener
	engine                  *Engine
}

// Start starts the server listening for requests.
func (server *Server) Start(engine *Engine) (err error) {
	server.acceptingNewConnections = true
	server.engine = engine
	server.clientListener, err = net.Listen("tcp", "localhost:13579")
	if err == nil {
		go server.runloop()
	} else {
		log.Println("[err] cannot listen on client address")
	}
	return
}

// Stop waits for pending requests and stops the server from af new ones.
func (server *Server) Stop() {
	server.acceptingNewConnections = false
	log.Print("[out] no longer accepting connections")
}

func handleConnection(conn net.Conn, engine *Engine) {
	defer conn.Close()

	var input Request
	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)
	err := decoder.Decode(&input)
	if err != nil {
		log.Println("err", err.Error())
		encoder.Encode(Response{error: err})
		return
	}

	output := engine.Execute(input)
	encoder.Encode(output)
}

func (server *Server) runloop() {
	log.Print("[out] accepting connections")
	for {
		conn, err := server.clientListener.Accept()
		if err != nil {
			log.Print("[err]", err.Error())
		} else {
			go handleConnection(conn, server.engine)
		}

		if !server.acceptingNewConnections {
			server.Stop()
		}
	}
}
