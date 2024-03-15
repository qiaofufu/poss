package volume

import (
	"os"
	"reflect"
	"testing"
)

func Test_localDisk_GetInfos(t *testing.T) {
	l := &localVolume{}
	list, err := l.GetInfos()
	if err != nil {
		t.Error(err)
	}

	for _, v := range list {
		t.Logf("%#v\n", v)
	}

}

func Test_localDisk_Open(t *testing.T) {
	type args struct {
		name string
		flag int
		perm os.FileMode
	}
	name := "./local_volume_test.go"
	file, err := os.OpenFile(name, os.O_RDWR, 0755)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		args    args
		want    IVirtualFile
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				name: name,
				flag: os.O_RDWR,
				perm: 0755,
			},
			want: &localVFile{file: file},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &localVolume{}
			got, err := l.Open("", tt.args.name, tt.args.flag, tt.args.perm)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Filename(), tt.want.Filename()) || !reflect.DeepEqual(got.Path(), tt.want.Path()) {
				t.Errorf("Open() got = %v, want %v", got.Filename(), tt.want.Filename())
			}
		})
	}
}
