package backends

import "io"

type (

	//
	Driver interface {
		List() ([]FileInfo, error)
		Read(string) (io.Reader, error)
		Remove(string) error
		Stat(string) (FileInfo, error)
		Write(string, io.Reader) error
	}

	//
	FileInfo struct {
		Name string
		Size int64
	}
)

var (

	//
	backends = map[string]Driver{
		"default":    Filesystem{},
		"filesystem": Filesystem{},
	}
)

// Use
func Use(backend string) Driver {

	// if the desired backend isn't found, return the default backend
	if _, ok := backends[backend]; !ok {
		return backends["default"]
	}

	// return the desired backend once found
	return backends[backend]
}
