package backends

import (
	"errors"
	"time"
)

//
type FileInfo struct {
	Name string
	Size int64
	ModTime time.Time
}

// errors
var (
	ErrInvalid    = errors.New("invalid argument")
	ErrPermission = errors.New("permission denied")
	ErrExist      = errors.New("file already exists")
	ErrNotExist   = errors.New("file does not exist")
	ErrNotFound   = errors.New("file not found")
	ErrUnknown    = errors.New("an unknown error ocurred")
	ErrFailed     = errors.New("operation failed")
)
