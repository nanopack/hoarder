// Package backends handles hoarder's persistant storage.
package backends

import (
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

//
type (
	// Driver utilized by the various backends
	Driver interface {
		initialize() error
		list() ([]DataInfo, error)
		read(string) (io.ReadCloser, error)
		remove(string) error
		stat(string) (DataInfo, error)
		write(string, io.Reader) error
	}

	// Selection of relevant data information
	DataInfo struct {
		Name    string    `json:"Name"`
		Size    int64     `json:"Size"`
		ModTime time.Time `json:"ModTime"`
	}
)

// default backend driver
var driver Driver

// Initialize initializes the default driver
func Initialize() error {

	// parse connection string
	u, err := url.Parse(viper.GetString("backend"))
	if err != nil {
		return err
	}

	// set backend based on connection string's scheme
	switch u.Scheme {
	case "file":
		driver = &Filesystem{Path: u.Path}
	case "":
		driver = &Filesystem{Path: u.Path}
	// case "scribble":
	// 	driver = &Scribble{Path: u.Path}
	// case "s3":
	// 	driver = &S3{Path: u.Path}
	// case "mongo":
	// 	driver = &Mongo{Path: u.Path}
	// case "redis":
	// 	driver = &Redis{Path: u.Path}
	// case "postgres":
	// 	driver = &Postgres{Path: u.Path}
	default:
		return fmt.Errorf(`
Unrecognized scheme '%s'. You can visit https://github.com/nanopack/hoarder and
submit a pull request adding the scheme or you can submit an issue requesting its
addition.
`, u.Scheme)
	}

	// initialize the driver
	return driver.initialize()
}

// List returns a list of data and info currently stored
func List() ([]DataInfo, error) {
	return driver.list()
}

// Read reads data and returns the contents
func Read(key string) (io.ReadCloser, error) {
	return driver.read(key)
}

// Remove removes data
func Remove(key string) error {
	return driver.remove(key)
}

// Stat returns information about data
func Stat(key string) (DataInfo, error) {
	return driver.stat(key)
}

// Write writes data
func Write(key string, r io.Reader) error {
	return driver.write(key, r)
}
