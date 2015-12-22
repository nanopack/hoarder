package storage

import (
	"io"
)

var FileInfo struct {
	Name string
	Size int64
}

var Storager interface {
	Stat(string) (FileInfo, error)
	Store(string, io.Reader) error
	Retrieve(string) (io.Reader, error)
	Remove(string) error
	List(string) ([]FileInfo, error)
}
