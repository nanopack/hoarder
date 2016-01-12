package backends

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const DEFAULT_FILESYSTEM_PATH = "/var/db/hoarder"

//
type Filesystem struct {
	Path string // path to the local database (default)
}

// init ensures the database exists before trying to do any operations on it
func (d *Filesystem) Init() error {
	//
	if d.Path == "" {
		d.Path = DEFAULT_FILESYSTEM_PATH
	}

	//
	return os.MkdirAll(d.Path, 0755)
}

// List
func (d Filesystem) List() ([]FileInfo, error) {
	//
	files, err := ioutil.ReadDir(d.Path)
	if err != nil {
		return nil, err
	}

	//
	info := []FileInfo{}
	for _, fi := range files {
		info = append(info, FileInfo{Name: fi.Name(), Size: fi.Size()})
	}

	//
	return info, nil
}

// Read
func (d Filesystem) Read(key string) (io.Reader, error) {
	//
	f, err := os.Open(filepath.Join(d.Path, key))
	if err != nil {
		return nil, err
	}

	//
	return f, nil
}

// Remove
func (d Filesystem) Remove(key string) error {
	//
	return os.Remove(filepath.Join(d.Path, key))
}

// Stat
func (d Filesystem) Stat(key string) (FileInfo, error) {
	//
	fi, err := os.Stat(filepath.Join(d.Path, key))
	if err != nil {
		return FileInfo{}, err
	}

	//
	return FileInfo{Name: fi.Name(), Size: fi.Size()}, nil
}

// Write
func (d Filesystem) Write(key string, r io.Reader) error {
	// read the entire contents of the reader
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// create/truncate a file and write the contents to it
	return ioutil.WriteFile(filepath.Join(d.Path, key), b, 0644)
}
