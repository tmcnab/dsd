package dsd

import (
	"crypto/sha256"
	"log"
	"time"
)

// HashString converts the provided string into 32-byte buffer.
func HashString(str string) (hash []byte) {
	function := sha256.New()
	function.Write([]byte(str))
	return function.Sum(nil)
}

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

// ToString converts an interface to a string.
func ToString(data interface{}) (s string, err error) {
	// TODO
	return
}

// LogInfo logs an informational message.
func LogInfo(from string, str string) {
	log.Println("out [" + from + "] " + str)
}

// LogError logs an error.
func LogError(from string, err error) {
	log.Println("err [" + from + "] " + err.Error())
}
