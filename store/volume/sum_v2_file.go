package volume

import (
	"github.com/hirochachacha/go-smb2"
	"io"
	"io/fs"
	"path/filepath"
)

type smbV2File struct {
	f *smb2.File
}

func (s *smbV2File) Stat() (fs.FileInfo, error) {
	return s.f.Stat()
}

func (s *smbV2File) Read(bytes []byte) (int, error) {
	return s.f.Read(bytes)
}

func (s *smbV2File) Close() error {
	return s.f.Close()
}

func (s *smbV2File) Seek(offset int64, whence int) (int64, error) {
	return s.f.Seek(offset, whence)
}

func (s *smbV2File) ReadAt(p []byte, off int64) (n int, err error) {
	return s.f.ReadAt(p, off)
}

func (s *smbV2File) Filename() string {
	return filepath.Base(s.f.Name())
}

func (s *smbV2File) Path() string {
	return filepath.Dir(s.f.Name())
}

func (s *smbV2File) Write(p []byte) (n int, err error) {
	return s.f.Write(p)
}

func (s *smbV2File) WriteAt(p []byte, off int64) (n int, err error) {
	return s.f.WriteAt(p, off)
}

func (s *smbV2File) WriteTo(w io.Writer) (n int64, err error) {
	return s.f.WriteTo(w)
}

func (s *smbV2File) WriteString(str string) (n int, err error) {
	return s.f.WriteString(str)
}
