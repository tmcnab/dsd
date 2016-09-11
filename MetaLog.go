package main

import (
	"encoding/gob"
	"time"
)

// A MetaLogEntry chronicles an object inserted into the database.
type MetaLogEntry struct {
	hash []byte    // The hash of the object.
	seek int64     // The position in the object data file where the object starts
	size int64     // The size of the object in bytes
	time time.Time // When the object was inserted.
}

// The MetaLog is the actual log of items, ordered by timestamp (maybe?)
type MetaLog struct {
	entries []MetaLogEntry
}

// Append a log entry to the log.
func (log *MetaLog) Append(entry MetaLogEntry) {
	log.entries = append(log.entries, entry)
}

// Flush persists the metalog to media.
func (log *MetaLog) Flush() (err error) {
	file, err := GetFile("metalog")
	if err == nil {
		defer file.Close()
		err = gob.NewEncoder(file).Encode(log)
	}
	return
}

// GetObjectByHash gets the object metadata by it's hash.
func (log *MetaLog) GetObjectByHash(hash []byte) (entry *MetaLogEntry) {
	return nil
}

// Since returns the index of the first entry which is after the given time.
func (log *MetaLog) Since(from time.Time) (index uint64) {
	return 0
}
