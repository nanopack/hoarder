package backends

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const DEFAULT_PATH = "/var/db/hoarder"

//
type Filesystem struct {
	Path string // path to the local database (default)
}

// init ensures the database exists before trying to do any operations on it
func (d Filesystem) init() error {

	//
	if d.Path == "" {
		d.Path = DEFAULT_PATH
	}

	//
	return os.MkdirAll(d.Path, 0755)
}

// List
func (d Filesystem) List() ([]FileInfo, error) {
	d.init()

	//
	info := []FileInfo{}

	//
	files, err := ioutil.ReadDir(d.Path)
	if err != nil {
		return nil, err
	}

	//
	for _, fi := range files {
		info = append(info, FileInfo{Name: fi.Name(), Size: fi.Size()})
	}

	return info, nil
}

// Read
func (d Filesystem) Read(key string) (io.Reader, error) {
	d.init()

	//
	f, err := os.Open(filepath.Join(d.Path, key))
	if err != nil {
		return nil, err
	}

	return f, nil
}

// Remove
func (d Filesystem) Remove(key string) error {
	d.init()

	//
	return os.Remove(filepath.Join(d.Path, key))
}

// Stat
func (d Filesystem) Stat(key string) (FileInfo, error) {
	d.init()

	fi, err := os.Stat(filepath.Join(d.Path, key))
	if err != nil {
		return FileInfo{}, err
	}

	return FileInfo{Name: fi.Name(), Size: fi.Size()}, nil
}

// Write
func (d Filesystem) Write(key string, r io.Reader) error {
	d.init()

	//
	b := make([]byte, 2048)

	for {

		//
		r.Read(b)

		//
		return ioutil.WriteFile(filepath.Join(d.Path, key), b, 0644)
	}
}
