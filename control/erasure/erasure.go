package erasure

import (
	"OSS/global"
	"fmt"
	"github.com/klauspost/reedsolomon"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type Erasure struct {
	id          string
	dataShards  int
	parityShads int
	encoder     reedsolomon.StreamEncoder
	data        []*os.File
	parity      []*os.File
}

func NewErasure(id string, dataShards int, parityShards int) *Erasure {
	enc, err := reedsolomon.NewStream(dataShards, parityShards)
	if err != nil {
		panic(err)
	}

	return &Erasure{
		id:          id,
		dataShards:  dataShards,
		parityShads: parityShards,
		encoder:     enc,
		data:        make([]*os.File, dataShards),
		parity:      make([]*os.File, parityShards),
	}
}

func (e *Erasure) GenerateDataBlocks(input *os.File) error {
	tmpDir := e.getTmpDir()
	if _, err := os.Stat(tmpDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(tmpDir, 0755)
			if err != nil {
				return err
			}
		}
		return fmt.Errorf("get tmp dir failed, err: %v", err)
	}

	tmpData := make([]io.Writer, e.dataShards)
	for i := range e.dataShards {
		fname := filepath.Join(tmpDir, strconv.Itoa(i))
		f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}
		tmpData[i] = f
		e.data[i] = f
	}

	// split data and write
	stat, err := input.Stat()
	if err != nil {
		return err
	}
	err = e.encoder.Split(input, tmpData, stat.Size())
	if err != nil {
		return err
	}
	return nil
}

func (e *Erasure) Encode(data []*os.File) error {
	if len(data) != e.dataShards {
		return fmt.Errorf("data length must equal dataShards, now is %d", e.dataShards)
	}

	tmpDir := e.getTmpDir()
	if _, err := os.Stat(tmpDir); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(tmpDir, 0755)
			if err != nil {
				return err
			}
		}
		return fmt.Errorf("get tmp dir failed, err: %v", err)
	}

	tmpData := make([]io.Reader, e.dataShards)
	for i := range data {
		tmpData[i] = data[i]
	}

	tmpParity := make([]io.Writer, e.parityShads)
	for i := range e.parityShads {
		fname := filepath.Join(tmpDir, e.id, fmt.Sprintf("checksum-%d", i))
		f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}
		tmpParity[i] = f
		e.parity[i] = f
	}

	return e.encoder.Encode(tmpData, tmpParity)
}

func (e *Erasure) ReStruct() error {
	if len(e.data) != e.dataShards || len(e.parity) != e.parityShads {
		return fmt.Errorf(
			"data length must is %d, parity length mus is %d. now data length is %d, parity length is %d",
			e.dataShards,
			e.parityShads,
			len(e.data),
			len(e.parity),
		)
	}
	shards := make([]io.Reader, 0, e.dataShards+e.parityShads)
	for i := range e.data {
		shards = append(shards, e.data[i])
	}
	for i := range e.parity {
		shards = append(shards, e.parity[i])
	}

	if ok, err := e.encoder.Verify(shards); err != nil {
		return err
	} else if ok {
		return nil
	}

	tmpDir := e.getTmpDir()

}

func (e *Erasure) Close() error {
	for _, file := range e.data {
		file.Close()
	}
	for _, file := range e.parity {
		file.Close()
	}
}

func (e *Erasure) getTmpDir() string {
	return filepath.Join(global.ErasureTmpDir, e.id)
}
