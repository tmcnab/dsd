package main

import (
	"log"
	"net"
)

// Server listens for requests and services them
type Server struct {
	af bool         // whether or not the server is af new connections
	cl net.Listener // client listener
}

// Start starts the server listening for requests.
func (server *Server) Start() bool {
	server.af = true
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

func (server *Server) runloop() {
	log.Print("[out] accepting connections")
	for {
		conn, err := server.cl.Accept()
		if err != nil {
			log.Print("[err]", err.Error())
		} else {
			// TODO something with conn
			conn.Close()
		}

		if !server.af {
			server.Stop()
		}
	}
}
