package dsd

import (
	"encoding/gob"
	"reflect"
	"time"
)

const _MetaLogFileName = "metalog"

// A MetaLogEntry chronicles an object inserted into the database.
type MetaLogEntry struct {
	hash [16]byte  // The hash of the object.
	seek int64     // The position in the object data file where the object starts
	size int64     // The size of the object in bytes
	time time.Time // When the object was inserted.
}

// The MetaLog is the actual log of items, ordered by timestamp (maybe?)
type MetaLog struct {
	// logicalTime is a logical timestamp. Every time an object is inserted the counter is incremented
	logicalTime uint64
	entries     []MetaLogEntry
	settings    *Settings
}

// NewMetaLog creates an initializes a new MetaLog type.
func NewMetaLog() (metaLog *MetaLog) {
	metaLog = &MetaLog{}
	metaLog.settings = NewSettings()

	// If the decoder throws it's probably because there's no data. If that's
	// the case then return a new one. TODO make sure this assumption is true.
	file, err := metaLog.settings.GetFile(_MetaLogFileName)
	if err == nil {
		if gob.NewDecoder(file).Decode(metaLog) != nil {
			metaLog.entries = make([]MetaLogEntry, 0)
			metaLog.logicalTime = 0
		}
	} else {
		LogError("metalog", err)
	}

	return
}

// Append a log entry to the log.
func (log *MetaLog) Append(entry MetaLogEntry) {
	log.entries = append(log.entries, entry)
}

// Flush persists the metalog to media.
func (log *MetaLog) Flush() (err error) {
	file, err := log.settings.GetFile(_MetaLogFileName)
	if err == nil {
		defer file.Close()
		err = gob.NewEncoder(file).Encode(log)
	}
	return
}

// GetMetaByHash gets the object metadata by it's hash.
func (log *MetaLog) GetMetaByHash(hash []byte) (entry *MetaLogEntry) {
	// TODO instead of a slice, log.entries should be stored using a tree
	// or better lookup structure other than iterating over the whole
	// damn list.
	for index := 0; index < len(log.entries); index++ {
		if reflect.DeepEqual(log.entries[index].hash, hash) {
			return &log.entries[index]
		}
	}

	return nil
}

// Since returns the index of the first entry which is after the given time.
func (log *MetaLog) Since(from time.Time, max uint64) (count uint64, objects []MetaLogEntry) {
	return
}
