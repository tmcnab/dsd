package dsd

import (
	"encoding/json"
	"log"
	"net"
)

// Server listens for requests and services them
type Server struct {
	engine    *Engine
	listener  net.Listener // TCP connection listener
	listening bool         // whether or not the server is af new connections
}

// Start starts the server listening for requests.
func (server *Server) Start(engine *Engine) (err error) {
	server.listening = true
	server.engine = engine
	server.listener, err = net.Listen("tcp", "localhost:13579")
	if err == nil {
		go server.runloop()
	} else {
		log.Fatalln("err [server] cannot listen on client address")
	}
	return
}

// Stop waits for pending requests and stops the server from af new ones.
func (server *Server) Stop() {
	server.listening = false
	LogInfo("server", "no longer acceping connections")
}

func (server *Server) handle(conn net.Conn) {
	defer conn.Close()

	var input Request
	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)
	err := decoder.Decode(&input)
	if err != nil {
		LogError("server", err)
		encoder.Encode(Response{error: err})
		return
	}

	// TODO should connections and req/res be logged?
	output := server.engine.Execute(input)
	encoder.Encode(output)
}

// Server loop. Accepts connections and cleans up after them. Possibly.
func (server *Server) runloop() {
	LogInfo("server", "accepting connections")

	for {
		conn, err := server.listener.Accept()
		if err == nil {
			// TODO: should we limit max processing connection time?
			go server.handle(conn)
		} else {
			LogError("server", err)
		}

		if server.listening {
			break
		}
	}

	// TODO should there be a timeout here to allow open ops to complete?
	LogInfo("server", "closing connections")
	server.listener.Close()
}
