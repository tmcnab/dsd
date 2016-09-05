package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	HOST = "localhost"
	PORT = 9373
	TYPE = "tcp"
)

func preflight() bool {
	os.Mkdir(".data", 0777)
	return true
}

func main() {
	if !preflight() {
		os.Exit(-1)
	}

	laddr := HOST + ":" + strconv.Itoa(PORT)
	listener, err := net.Listen(TYPE, laddr)
	if err != nil {
		fmt.Println("unable to listen on " + laddr)
		os.Exit(2)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept " + err.Error())
		} else {
			go handleRequest(conn)
		}
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("error reading " + err.Error())
	} else {
		// handle request here
	}
}
