package volume

import (
	"io"
	"io/fs"
	"os"
)

type IVirtualVolume interface {
	GetInfo() (Info, error)
	Open(filename string, flag int, perm os.FileMode) (IVirtualFile, error)
	Close() error
}

type IVirtualFile interface {
	fs.File
	io.Seeker
	io.ReaderAt
	io.Writer
	io.WriterAt
	io.WriterTo
	io.StringWriter
	Filename() string
	Path() string
}
