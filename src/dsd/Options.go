package dsd

import (
	"flag"
	"os"
	"path"
)

// Settings contains all runtime information things need to execute correctly.
type Settings struct {
	dir string
}

// NewSettings creates and initializes a new Settings type.
func NewSettings() (settings *Settings) {
	*settings = Settings{}
	dir, err := os.Getwd()
	if err == nil {
		flag.StringVar(&settings.dir, "dir", dir, "usage")
	} else {
		LogError("settings", err)
	}
	return
}

// GetFile returns a file pointer to a file in the data directory.
func (settings *Settings) GetFile(name string) (file *os.File, err error) {
	filename := path.Join(settings.dir, ".data", name)
	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	return
}
