package main

import (
	"errors"
	"time"
)

type hashlogEntry struct {
	hash uint64
	time time.Time // 16 bytes
}

type datalogEntry struct {
	object interface{}
}

// The Engine is the blood and guts of the database system.
type Engine struct {
	test     bool
	lastTime time.Time // A timestamp of when the last object was writen to file.
	hashlog  []hashlogEntry
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
func (engine *Engine) insert(input *EngineInput, output *EngineOutput) {
	// 1. Compute hash:
	// 		a. if in primary index, return hash
	// 		b. if not, proceed with insert
	// 2. Update files
	//		a. Append timestamp and hash into HASHLOG, update in-memory HASHLOG
	//		b. Insert object
}
