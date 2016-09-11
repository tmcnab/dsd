package main

import (
	"bytes"
	"errors"
	"os"
	"time"
)

// The Engine is the blood and guts of the database system.
type Engine struct {
	test bool
	log  MetaLog // the only memory resident part of the data store
}

// Execute the given input and produce an output.
func (engine *Engine) Execute(input EngineInput) (output EngineOutput) {
	switch input.op {
	case "insert":
		engine.insert(&input, &output)
		break
	default:
		return EngineOutput{error: errors.New("unsupported operation")}
	}
	return
}

// Insert an object into the set, let peers know.
func (engine *Engine) insert(input *EngineInput, output *EngineOutput) {

	// 1. Compute hash and check for existence. If in metalog, return entry to client.
	var entry MetaLogEntry
	object := Object(input.arg)
	entry.hash, output.error = object.Hash()
	if output.error == nil {
		meta := engine.log.GetObjectByHash(entry.hash)
		if meta != nil {
			output.objects = append(output.objects, meta)
			return
		}
	} else {
		return
	}

	// 2. Convert object to bytes
	var data *bytes.Buffer
	data, output.error = object.Encode()
	if output.error != nil {
		return
	}

	// 3. Open object file.
	var file *os.File
	file, output.error = GetFile("objects")
	if output.error != nil {
		return
	}

	// 4. Seek to end, store the seek position.
	entry.seek, output.error = file.Seek(0, os.SEEK_END)
	if output.error != nil {
		return
	}

	// 5. Write object data to file.
	entry.size = int64(data.Len())
	entry.time = time.Now()
	_, output.error = file.Write(data.Bytes())
	if output.error != nil {
		return
	}
	output.error = file.Close()

	// 6. Update log entries, persist to file.
	engine.log.Append(entry)
	output.error = engine.log.Flush()
	if output.error == nil {
		output.objects = append(output.objects, entry)
	}
	return
}
