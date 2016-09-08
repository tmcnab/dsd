package main

import (
	"encoding/json"
	"log"
	"net"
)

// Server listens for requests and services them
type Server struct {
	af     bool         // whether or not the server is af new connections
	cl     net.Listener // client listener
	engine *Engine
}

// Start starts the server listening for requests.
func (server *Server) Start(engine *Engine) bool {
	server.af = true
	server.engine = engine
	listener, err := net.Listen("tcp", "localhost:13579")
	if err != nil {
		log.Println("[err] cannot listen on client address")
		return false
	}

	server.cl = listener
	server.runloop()
	return true
}

// Stop waits for pending requests and stops the server from af new ones.
func (server *Server) Stop() {
	server.af = false
	log.Print("[out] no longer accepting connections")
}

func possiblyWriteError(conn net.Conn, err error) bool {
	if err == nil {
		return false
	} else {
		log.Print("err", err.Error())
		conn.Write("{ error: \"" + err.Error() + "\"")
		return true
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var input EngineInput
	err := json.NewDecoder(conn).Decode(v)
	encoder := json.NewEncoder(conn)
	if err != null {
		return encoder.Encode(EngineOutput{error: err})
	}

	err := encoder.Encode(input.Execute())
	if err != nil {
		log.Println("err", err.Error())
	}
}

func (server *Server) runloop() {
	log.Print("[out] accepting connections")
	for {
		conn, err := server.cl.Accept()
		if err != nil {
			log.Print("[err]", err.Error())
		} else {
			go handleConnection(conn)
		}

		if !server.af {
			server.Stop()
		}
	}
}
