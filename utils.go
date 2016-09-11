package main

import (
	"crypto/sha256"
	"os"
	"path"
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
