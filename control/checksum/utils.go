package checksum

import (
	"fmt"
	"io"
	"os"
)

func convertFileSliceToReaderSlice(slice []*os.File) []io.Reader {
	res := make([]io.Reader, len(slice))
	for i, item := range slice {
		if item == nil {
			res[i] = nil
		} else {
			res[i] = item
		}
	}
	return res
}

func convertFileSliceToWriterSlice(slice []*os.File) []io.Writer {
	res := make([]io.Writer, len(slice))
	for i, item := range slice {
		if item == nil {
			res[i] = nil
		} else {
			res[i] = item
		}
	}
	return res
}

func getFilePaths(slice []*os.File) []string {
	filePaths := make([]string, len(slice))
	for i, f := range slice {
		filePaths[i] = f.Name()
	}
	return filePaths
}

func openFiles(paths []string, flag int, perm os.FileMode) ([]*os.File, error) {
	res := make([]*os.File, len(paths))
	for i, name := range paths {
		f, err := os.OpenFile(name, flag, perm)
		if err != nil {
			if os.IsNotExist(err) {
				res[i] = nil
				continue
			}
			return nil, fmt.Errorf("open file failed, at element %d", i)
		}
		res[i] = f
	}
	return res, nil
}
