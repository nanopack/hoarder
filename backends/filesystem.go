package backends

import "io"

type Filesystem struct{}

//
func (d Filesystem) List() ([]FileInfo, error) {
	return nil, nil
}

//
func (d Filesystem) Read(key string) (io.Reader, error) {
	return nil, nil
}

//
func (d Filesystem) Remove(key string) error {
	return nil
}

//
func (d Filesystem) Stat(key string) (FileInfo, error) {
	return FileInfo{}, nil
}

//
func (d Filesystem) Write(key string, r io.Reader) error {
	return nil
}
