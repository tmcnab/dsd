package main

import (
	"crypto/sha256"
	"log"
	"os"
	"path"
	"time"
)

// HashString converts the provided string into 32-byte buffer.
func HashString(str string) (hash []byte) {
	function := sha256.New()
	function.Write([]byte(str))
	return function.Sum(nil)
}

// GetFile returns a file (or error) in the dsd directory.
func GetFile(name string) (file *os.File, err error) {
	wd, err := os.Getwd()
	if err == nil {
		filename := path.Join(wd, ".data", name)
		file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	}
	return
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

func LogInfo(from string, str string) {
	log.Println("out [" + from + "] " + str)
}

func LogError(from string, err error) {
	log.Println("err [" + from + "] " + err.Error())
}
