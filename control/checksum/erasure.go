package checksum

import (
	"OSS/global"
	"fmt"
	"github.com/klauspost/reedsolomon"
	"io"
	"os"
)

func init() {
	err := os.MkdirAll(global.ErasureTmpDir, 0755)
	if err != nil {
		panic(err)
	}
}

func defaultSplitStrategy() (int, int) {
	return 10, 3
}

type Encoder struct {
	splitStrategy func() (int, int)
}

func New() *Encoder {

	return &Encoder{
		splitStrategy: defaultSplitStrategy,
	}
}

func (e *Encoder) SetSplitStrategy(strategy func() (dataShards int, parityShards int)) {
	e.splitStrategy = strategy
}

func (e *Encoder) GenerateDataChunk(filename string) ([]string, error) {
	input, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = input.Close()
	}()
	stat, err := input.Stat()
	if err != nil {
		return nil, err
	}

	dataShards, parityShards := e.splitStrategy()
	enc, err := reedsolomon.NewStream(dataShards, parityShards)
	if err != nil {
		return nil, err
	}
	data := make([]*os.File, dataShards)
	for i := range dataShards {
		f, err := os.CreateTemp(global.ErasureTmpDir, fmt.Sprintf("*_data.part%d", i))
		if err != nil {
			return nil, err
		}
		data[i] = f
	}
	defer func() {
		for i := range data {
			_ = data[i].Close()
		}
	}()

	err = enc.Split(input, convertFileSliceToWriterSlice(data), stat.Size())
	if err != nil {
		return nil, err
	}
	return getFilePaths(data), nil
}

func (e *Encoder) Encode(files []string) ([]string, error) {
	data, err := openFiles(files, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer func() {
		for _, f := range data {
			_ = f.Close()
		}
	}()

	dataShards, parityShards := e.splitStrategy()
	enc, err := reedsolomon.NewStream(dataShards, parityShards)
	if err != nil {
		return nil, err
	}

	parity := make([]*os.File, parityShards)
	for i := range parityShards {
		f, err := os.CreateTemp(global.ErasureTmpDir, fmt.Sprintf("*_parity.part%d", i))
		if err != nil {
			return nil, err
		}
		parity[i] = f
	}
	defer func() {
		for i := range parity {
			_ = parity[i].Close()
		}
	}()
	err = enc.Encode(convertFileSliceToReaderSlice(data), convertFileSliceToWriterSlice(parity))
	if err != nil {
		return nil, err
	}

	return getFilePaths(parity), nil
}

func (e *Encoder) Verify(dataFiles []string, parityFiles []string) (bool, error) {
	data, err := openFiles(dataFiles, os.O_RDONLY, 0755)
	if err != nil {
		return false, err
	}
	defer func() {
		for _, f := range data {
			_ = f.Close()
		}
	}()
	parity, err := openFiles(parityFiles, os.O_RDONLY, 0755)
	if err != nil {
		return false, err
	}
	defer func() {
		for _, f := range parity {
			_ = f.Close()
		}
	}()
	dataShards, parityShards := e.splitStrategy()
	enc, err := reedsolomon.NewStream(dataShards, parityShards)
	if err != nil {
		return false, err
	}

	shards := make([]io.Reader, 0, dataShards+parityShards)
	shards = append(shards, convertFileSliceToReaderSlice(data)...)
	shards = append(shards, convertFileSliceToReaderSlice(parity)...)

	if ok, err := enc.Verify(shards); !ok || err != nil {
		return false, err
	}
	return true, nil
}

func (e *Encoder) ReStruct(dataFiles []string, parityFiles []string) ([]string, []string, error) {
	if ok, _ := e.Verify(dataFiles, parityFiles); ok {
		return dataFiles, parityFiles, nil
	}
	data, err := openFiles(dataFiles, os.O_RDONLY, 0755)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		for _, f := range data {
			_ = f.Close()
		}
	}()
	parity, err := openFiles(parityFiles, os.O_RDONLY, 0755)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		for _, f := range parity {
			_ = f.Close()
		}
	}()

	dataShards, parityShards := e.splitStrategy()
	enc, err := reedsolomon.NewStream(dataShards, parityShards)
	if err != nil {
		return nil, nil, err
	}

	shards := make([]io.Reader, 0, dataShards+parityShards)
	shards = append(shards, convertFileSliceToReaderSlice(data)...)
	shards = append(shards, convertFileSliceToReaderSlice(parity)...)

	out := make([]*os.File, dataShards+parityShards)
	dataPath := make([]string, dataShards)
	parityPath := make([]string, parityShards)
	for i, f := range data {
		if f == nil {
			tmp, err := os.CreateTemp("", fmt.Sprintf("*_data.part%d", i))
			if err != nil {
				return nil, nil, err
			}
			out[i] = tmp
			dataPath[i] = tmp.Name()
		} else {
			dataPath[i] = f.Name()
		}
	}
	for i, f := range parity {
		if f == nil {
			tmp, err := os.CreateTemp("", fmt.Sprintf("*_parity.part%d", i))
			if err != nil {
				return nil, nil, err
			}
			out[dataShards+i] = tmp
			parityPath[i] = tmp.Name()
		} else {
			parityPath[i] = f.Name()
		}
	}
	defer func() {
		for _, f := range out {
			if f != nil {
				_ = f.Close()
			}
		}
	}()

	if err = enc.Reconstruct(shards, convertFileSliceToWriterSlice(out)); err != nil {
		return nil, nil, err
	}
	return dataPath, parityPath, nil
}
