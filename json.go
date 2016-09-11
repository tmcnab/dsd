package main

import "time"

// Number is what the golang marshaller converts a number to.
type Number float64

// Convert2Time converts a given interface, which should be a RFC3339Nano
// string, to a Time.
func Convert2Time(data interface{}) (t time.Time, err error) {
	// TODO
	return
}

// InterfaceToNumber attempts to convert an anonymous interface (JSON kv pair)
// to a 64-bit float number.
func InterfaceToNumber(data interface{}) (n Number, err error) {
	// TODO
	return
}
