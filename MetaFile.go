package main

import (
	"bytes"
	"encoding/gob"
)

// The MetaFile contains all run-time information.
type MetaFile map[string]interface{}

const (
	MF_LAST_TIMESTAMP = iota
)

// GetMetaFile does what it says: fetches the Metafile.
func GetMetaFile() (metaFile *MetaFile, err error) {
	file, err := GetFile("metafile")
	if err == nil {
		defer file.Close()
		buf := new(bytes.Buffer)
		err = gob.NewDecoder(buf).Decode(metaFile)
	}
	return
}
