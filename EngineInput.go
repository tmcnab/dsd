package main

// EngineInput represents data that a client sends to the server.
type EngineInput struct {
	arg Object
	op  string
	ver float64
}

// TODO add a custom binary marshaller
