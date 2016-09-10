package main

import (
	"bytes"
	"encoding/gob"
)

// The Object type is the very thing that we're storing in the database.
type Object map[string]interface{}

// Decode a buffer into an object.
func (object *Object) Decode(buffer *bytes.Buffer) (err error) {
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(object)
	return
}

// Encode an object into a buffer.
func (object *Object) Encode() (buffer *bytes.Buffer, err error) {
	buffer = new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err = encoder.Encode(*object)
	return
}

// Hash an object into a N-byte array.
func (object *Object) Hash() (hash []byte, err error) {
	return
}
