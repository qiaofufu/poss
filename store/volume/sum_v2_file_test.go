package volume

import (
	"bytes"
	"errors"
	"github.com/hirochachacha/go-smb2"
	"io"
	"io/fs"
	"os"
	"reflect"
	"testing"
)

var client = NewSMBV2Disk("192.168.22.150:445", "qhh", "123123")
var file, _ = client.Open("//192.168.22.150/share", "test.txt", os.O_CREATE|os.O_RDWR, 0755)

func Test_smbV2File_Close(t *testing.T) {

	tests := []struct {
		name    string
		fields  IVirtualVolume
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: client,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := file
			if err := s.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_smbV2File_Filename(t *testing.T) {
	tests := []struct {
		name    string
		fields  IVirtualVolume
		wantErr bool
		want    string
	}{
		{
			name:   "test-1",
			fields: client,
			want:   "test.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := file
			if got := s.Filename(); got != tt.want {
				t.Errorf("Filename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smbV2File_Path(t *testing.T) {
	tests := []struct {
		name    string
		fields  IVirtualVolume
		wantErr bool
		want    string
	}{
		{
			name:   "test-1",
			fields: client,
			want:   ".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := file
			if got := s.Path(); got != tt.want {
				t.Errorf("Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smbV2File_Read(t *testing.T) {

	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  IVirtualFile
		args    args
		want    int
		wantErr bool
	}{
		{
			name:   "test",
			fields: file,
			args:   args{bytes: make([]byte, 1024)},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			got, err := s.Read(tt.args.bytes)
			if (err != nil) != tt.wantErr && !errors.Is(err, io.EOF) {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smbV2File_ReadAt(t *testing.T) {
	type fields struct {
		f *smb2.File
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &smbV2File{
				f: tt.fields.f,
			}
			gotN, err := s.ReadAt(tt.args.p, tt.args.off)
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

func Test_smbV2File_Seek(t *testing.T) {
	type fields struct {
		f *smb2.File
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &smbV2File{
				f: tt.fields.f,
			}
			got, err := s.Seek(tt.args.offset, tt.args.whence)
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

func Test_smbV2File_Stat(t *testing.T) {
	type fields struct {
		f *smb2.File
	}
	tests := []struct {
		name    string
		fields  fields
		want    fs.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &smbV2File{
				f: tt.fields.f,
			}
			got, err := s.Stat()
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

func Test_smbV2File_Write(t *testing.T) {

	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  IVirtualFile
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: file,
			args:   args{p: []byte("hello, oss")},
			wantN:  10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			gotN, err := s.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Write() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_smbV2File_WriteAt(t *testing.T) {
	type args struct {
		p   []byte
		off int64
	}
	tests := []struct {
		name    string
		fields  IVirtualFile
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: file,
			args: args{
				p:   []byte("ni hao, oss"),
				off: 12,
			},
			wantN: len("ni hao, oss"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			gotN, err := s.WriteAt(tt.args.p, tt.args.off)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("WriteAt() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_smbV2File_WriteString(t *testing.T) {

	type args struct {
		str string
	}
	tests := []struct {
		name    string
		fields  IVirtualFile
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: file,
			args:   args{str: "xxxx"},
			wantN:  len("xxxx"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			gotN, err := s.WriteString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("WriteString() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_smbV2File_WriteTo(t *testing.T) {
	type args struct {
		w *bytes.Buffer
	}

	tests := []struct {
		name    string
		fields  IVirtualFile
		args    args
		wantW   string
		wantN   int64
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: file,
			args:   args{w: &bytes.Buffer{}},
			wantW:  "123123",
			wantN:  6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			gotN, err := s.WriteTo(tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := tt.args.w.String(); gotW != tt.wantW {
				t.Errorf("WriteTo() gotW = %v, want %v", gotW, tt.wantW)
			}
			if gotN != tt.wantN {
				t.Errorf("WriteTo() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
