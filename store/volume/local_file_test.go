package volume

import (
	"io/fs"
	"os"
	"reflect"
	"sync"
	"testing"
)

func Test_localVFile_Close(t *testing.T) {
	type fields struct {
		file *os.File
		mu   sync.RWMutex
	}

	file, err := os.OpenFile("./local_file_test.go", os.O_RDWR, 0755)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test-1",
			fields: fields{
				file: file,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &localVFile{
				file: tt.fields.file,
				mu:   tt.fields.mu,
			}
			if err := f.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_localVFile_Filename(t *testing.T) {
	type fields struct {
		file *os.File
		mu   sync.RWMutex
	}
	file, err := os.OpenFile("./local_file_test.go", os.O_RDWR, 0755)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test-1",
			fields: fields{
				file: file,
				mu:   sync.RWMutex{},
			},
			want: "local_file_test.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &localVFile{
				file: tt.fields.file,
				mu:   tt.fields.mu,
			}
			if got := f.Filename(); got != tt.want {
				t.Errorf("Filename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_localVFile_Path(t *testing.T) {
	type fields struct {
		file *os.File
		mu   sync.RWMutex
	}
	file, err := os.OpenFile("./local_file_test.go", os.O_RDWR, 0755)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test-1",
			fields: fields{
				file: file,
				mu:   sync.RWMutex{},
			},
			want: ".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &localVFile{
				file: tt.fields.file,
				mu:   tt.fields.mu,
			}
			if got := f.Path(); got != tt.want {
				t.Errorf("Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_localVFile_Read(t *testing.T) {
	type fields struct {
		file *os.File
		mu   sync.RWMutex
	}
	file, err := os.OpenFile("./local_file_test.go", os.O_RDWR, 0755)
	if err != nil {
		t.Error(err)
	}
	tmp := make([]byte, 1024)
	bytes := make([]byte, 1024)
	n, err := file.Read(tmp)
	if err != nil {
		t.Error(err)
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test-1",
			fields: fields{
				file: file,
				mu:   sync.RWMutex{},
			},
			args: args{bytes: bytes},
			want: n,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &localVFile{
				file: tt.fields.file,
				mu:   tt.fields.mu,
			}
			got, err := f.Read(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_localVFile_ReadAt(t *testing.T) {
	type fields struct {
		file *os.File
		mu   sync.RWMutex
	}
	file, err := os.OpenFile("./local_file_test.go", os.O_RDWR, 0755)
	if err != nil {
		t.Error(err)
	}
	tmp := make([]byte, 1024)
	bytes := make([]byte, 1024)
	n, err := file.ReadAt(tmp, 24)
	if err != nil {
		t.Error(err)
	}

	type args struct {
		p   []byte
		off int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: fields{file: file},
			args: args{
				p:   bytes,
				off: 24,
			},
			wantN: n,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &localVFile{
				file: tt.fields.file,
				mu:   tt.fields.mu,
			}
			gotN, err := f.ReadAt(tt.args.p, tt.args.off)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ReadAt() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_localVFile_Seek(t *testing.T) {
	type fields struct {
		file *os.File
		mu   sync.RWMutex
	}
	file, err := os.OpenFile("./local_file_test.go", os.O_RDWR, 0755)
	if err != nil {
		t.Error(err)
	}
	var offset int64 = 24
	whence := 0
	ret, err := file.Seek(offset, whence)
	if err != nil {
		t.Error(err)
	}
	type args struct {
		offset int64
		whence int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: fields{file: file},
			args: args{
				offset: offset,
				whence: whence,
			},
			want: ret,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &localVFile{
				file: tt.fields.file,
				mu:   tt.fields.mu,
			}
			got, err := f.Seek(tt.args.offset, tt.args.whence)
			if (err != nil) != tt.wantErr {
				t.Errorf("Seek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Seek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_localVFile_Stat(t *testing.T) {
	type fields struct {
		file *os.File
		mu   sync.RWMutex
	}
	file, err := os.OpenFile("./local_file_test.go", os.O_RDWR, 0755)
	if err != nil {
		t.Error(err)
	}
	stat, err := file.Stat()
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		fields  fields
		want    fs.FileInfo
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: fields{file: file},
			want:   stat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &localVFile{
				file: tt.fields.file,
				mu:   tt.fields.mu,
			}
			got, err := f.Stat()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stat() got = %v, want %v", got, tt.want)
			}
		})
	}
}
