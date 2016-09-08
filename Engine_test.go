package main

import "testing"

func TestInsert(test *testing.T) {
	engine := &Engine{test: true}
	input := &EngineInput{op: "insert"}
	output := &EngineOutput{}
	engine.insert(input, output)
}
