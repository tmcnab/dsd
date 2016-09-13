package main

import "testing"

func TestInsert(test *testing.T) {
	engine := &Engine{test: true}
	input := &Request{op: "insert"}
	output := &Response{}
	engine.insert(input, output)
}
