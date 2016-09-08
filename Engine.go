package main

import "errors"

type Engine struct{}

// Execute the given input and produce an output.
func (engine *Engine) Execute(input EngineInput) (output EngineOutput) {
	switch input.op {
	case "insert":
		engine.insert(&input, &output)
		break
	default:
		return EngineOutput{error: errors.New("unsupported op")}
	}
	return
}

func (engine *Engine) insert(input *EngineInput, output *EngineOutput) {

}
