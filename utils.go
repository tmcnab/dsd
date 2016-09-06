package main

import (
	"os"
	"path"
)

// GetFile returns a file (or error) in the dsd directory.
func GetFile(name string) (file os.File, err error) {
	wd, err := os.Getwd()
	if err != null {
		return
	}

	filename = path.Join(wd, ".data", name)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
}
