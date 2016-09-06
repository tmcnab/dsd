package main

/*
   TODO should the GET/LOAD calls close the file descriptors after usage?
*/

// The Store reprents the actual data being stored by users.
type Store struct {
	indices []Index
	data    map[uint64]uint64 // maps the hash => file seek offset
}

// Load the store from file (if exists).
func (store *Store) Load() (err error) {
	file, err := GetFile("store")
	if err == nil {
		defer file.Close()
	} else {
		return
	}

	store.data = make(map[uint64]uint64, 0)

	// TODO buffer file data and load into map
	// TODO is there a better way of loading/saving simple maps to file?
}

// Get returns the actual object that has been stored.
func (store *Store) Get(hash uint64) (item string, err error) {
	file, err := GetFile("data")
	if err == nil {
		defer file.Close()
	} else {
		return
	}

	// 1. Seek to file:store.data[hash]
	// 2. Read a 8 bytes. This indicates how long the following sequence is.
	// 3. Read the number of bytes specified above
	// 4. convert to a string and return
}
