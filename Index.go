package main

// A Predicate function tests whether or ot
type predicate func(object interface{}) bool

// An Index is a list of objects which
type Index struct {
	data map[string]bool
	name string
	test predicate
}

// Destroy the index irrecoverably.
func (index *Index) Destroy() (err error) {
	return nil
}

// Ingest an object
func (index *Index) Ingest() (err error) {
	return nil
}

// LoadIndex loads an index from file.
func LoadIndex(name string) (index Index, err error) {
	return Index{}, nil
}

// CreateIndex - creates a new index given the the two inputs.
func CreateIndex(name string, test string) (index Index, err error) {
	return Index{}, nil
}
