package main

import (
	"errors"
	"os"
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

// Insert an object into the set.
//
// 2. Update files
//		a. Append timestamp and hash into metalog, update in-memory metalog
//		b. Append object to file
func (engine *Engine) insert(input *EngineInput, output *EngineOutput) {
	var file *os.File
	var entry MetaLogEntry
	object := &(Object(input.arg))

	// 1. Compute hash and check for existence. If in metalog, return entry.
	entry.hash, err = object.Hash()
	meta := engine.log.GetObjectByHash(entry.hash)
	if meta != nil {
		output.objects = append(output.objects, meta)
		return
	}

	// 2. Open object file.
	file, output.error = GetFile("objects")
	if output.error != nil {
		return
	}

	// 3. Seek to end, store the seek position.
	entry.seek, output.error = file.Seek(0, os.SEEK_END)
	if output.error != nil {
		return
	}

	// 4. Convert object to bytes
	data, output.error := object.Encode()
	if output.error != nil {
		return
	}

	// 5. Write to file
	// meta.size = data.

}
