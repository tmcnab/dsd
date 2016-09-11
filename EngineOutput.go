package main

// EngineOutput represents what gets sent back to the client.
type EngineOutput struct {
	error error
	data  map[string]interface{}
}
