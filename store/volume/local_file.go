package volume

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

type localVFile struct {
	file *os.File
	mu   sync.RWMutex // 控制文件访问
}


func (f *localVFile) Stat() (fs.FileInfo, error) {
	return f.file.Stat()
}

func (f *localVFile) Read(bytes []byte) (int, error) {
	return f.file.Read(bytes)
}

func (f *localVFile) Close() error {
	return f.file.Close()
}

func (f *localVFile) Seek(offset int64, whence int) (int64, error) {
	return f.file.Seek(offset, whence)
}

func (f *localVFile) ReadAt(p []byte, off int64) (n int, err error) {
	return f.file.ReadAt(p, off)
}

func (f *localVFile) Filename() string {
	name := f.file.Name()
	return filepath.Base(name)
}

func (f *localVFile) Path() string {
	name := f.file.Name()
	return filepath.Dir(name)
}

func (f *localVFile) Write(p []byte) (n int, err error) {
	return f.file.Write(p)
}

func (f *localVFile) WriteAt(p []byte, off int64) (n int, err error) {
	return f.file.WriteAt(p, off)
}

func (f *localVFile) WriteTo(w io.Writer) (n int64, err error) {
	return f.file.WriteTo(w)
}

func (f *localVFile) WriteString(s string) (n int, err error) {
	return f.file.WriteString(s)
}
