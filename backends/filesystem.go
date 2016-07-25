package backends

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const DEFAULT_FILESYSTEM_PATH = "/var/db/hoarder"

// implementation of Driver
type Filesystem struct {
	Path string // path to the local database (default)
}

// initialize ensures the database exists before trying to do any operations on it
func (d *Filesystem) initialize() error {
	//
	if d.Path == "" {
		d.Path = DEFAULT_FILESYSTEM_PATH
	}

	//
	return os.MkdirAll(d.Path, 0755)
}

// list returns a list of files, and some info, currently stored
func (d Filesystem) list() ([]DataInfo, error) {

	//
	files, err := ioutil.ReadDir(d.Path)
	if err != nil {
		return nil, err
	}

	//
	info := []DataInfo{}
	for _, fi := range files {
		info = append(info, DataInfo{Name: fi.Name(), Size: fi.Size(), ModTime: fi.ModTime().UTC()})
	}

	//
	return info, nil
}

// read reads a file and returns the contents
func (d Filesystem) read(key string) (io.ReadCloser, error) {

	//
	f, err := os.Open(filepath.Join(d.Path, key))
	if err != nil {
		return nil, err
	}

	//
	return f, nil
}

// remove removes a file
func (d Filesystem) remove(key string) error {
	return os.RemoveAll(filepath.Join(d.Path, key))
}

// stat returns information about a file
func (d Filesystem) stat(key string) (DataInfo, error) {

	//
	fi, err := os.Stat(filepath.Join(d.Path, key))
	if err != nil {
		return DataInfo{}, err
	}

	//
	return DataInfo{Name: fi.Name(), Size: fi.Size(), ModTime: fi.ModTime().UTC()}, nil
}

// write writes data to a file
func (d Filesystem) write(key string, r io.Reader) error {
	f, err := os.Create(filepath.Join(d.Path, key))
	if err != nil {
		return fmt.Errorf("Failed to open file to write - %v\n", err)
	}
	defer f.Close()
	// defer r.Close()

	// pipe contents of reader to file (save some rams)
	_, err = io.Copy(f, r)

	return err
}
