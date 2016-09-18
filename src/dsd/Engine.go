package dsd

import (
	"bytes"
	"errors"
	"os"
	"time"
)

// The Engine is the blood and guts of the database system.
type Engine struct {
	log      *MetaLog // the only memory resident part of the data store
	settings *Settings
}

// NewEngine creates and initializes an Engine struct.
func NewEngine() *Engine {
	engine := Engine{}
	engine.log = NewMetaLog()
	engine.settings = NewSettings()
	return &engine
}

// Request represents data that a client sends to the server.
type Request struct {
	arg Object
	op  string
	ver float64
}

// Response represents what gets sent back to the client.
type Response struct {
	error error
	data  map[string]interface{}
}

// Execute the given input and produce an output.
func (engine *Engine) Execute(input Request) (output Response) {
	switch input.op {
	case "insert":
		engine.insert(&input, &output)
		break
	default:
		return Response{error: errors.New("unsupported operation")}
	}
	return
}

// Insert an object into the set, let peers know.
func (engine *Engine) insert(input *Request, output *Response) {
	// 1. Compute hash and check for existence. If in metalog, return entry to client.
	var entry MetaLogEntry
	object := Object(input.arg)
	entry.hash, output.error = object.Hash()
	if output.error == nil {
		entry := engine.log.GetMetaByHash(entry.hash)
		if entry != nil {
			// TODO convert meta to map[string]interface{}
			return
		}
	} else {
		return
	}

	// 2. Convert object to bytes
	var data *bytes.Buffer
	data, output.error = object.Encode()
	if output.error == nil {
		entry.size = int64(data.Len())
	} else {
		return
	}

	// 3. Open object file.
	var file *os.File
	settings := NewSettings()
	file, output.error = settings.GetFile("objects")
	if output.error != nil {
		return
	}

	// 4. Seek to end, store the seek position.
	entry.seek, output.error = file.Seek(0, os.SEEK_END)
	if output.error != nil {
		return
	}

	// TODO: lock metadata file

	// 5. Write object data to file.
	_, output.error = file.Write(data.Bytes())
	if output.error != nil {
		return
	}
	output.error = file.Close()

	// 6. Update log entries, persist to file.
	entry.time = time.Now().In(time.UTC)
	engine.log.Append(entry)
	output.error = engine.log.Flush()
	if output.error == nil {
		// TODO convert entry to map[string]interface{}
	}
}

// Return the objects which have been inserted since a given timestamp.
//
// This function is primarily used by other nodes in the cluster to update
// their own state. Perhaps on a backing-off polling loop.
func (engine *Engine) since(input *Request, output *Response) {
	// 1. Get 'since' command argument.
	var timestamp time.Time
	timestamp, output.error = Convert2Time(input.arg["since"])
	if output.error != nil {
		return
	}

	// 2. Get 'max' command argument.
	var max Number
	max, output.error = InterfaceToNumber(input.arg["max"])
	if output.error != nil {
		return
	}

	// 3. Query the metalog for:
	// 		"count"	- the number of objects that have been inserted since
	// 		"objects" - up to N objects (m if m is less)
	output.data["count"], output.data["objects"] =
		engine.log.Since(timestamp, uint64(max))
}
