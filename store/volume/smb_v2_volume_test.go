package volume

import (
	"os"
	"reflect"
	"testing"
)

func TestNewSMBV2Disk(t *testing.T) {
	type args struct {
		address  string
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want IVirtualVolume
	}{
		{
			name: "test-1",
			args: args{
				address:  "192.168.22.150:445",
				username: "qhh",
				password: "123123",
			},
			want: NewSMBV2Disk("192.168.22.150:445", "qhh", "123123"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSMBV2Disk(tt.args.address, tt.args.username, tt.args.password)
			if got == nil {
				t.Errorf("NewSMBV2Disk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smbV2Disk_Close(t *testing.T) {
	disk := NewSMBV2Disk("192.168.22.150:445", "qhh", "123123")
	if disk == nil {
		t.Fatal("new SMB disk fail")
	}
	err := disk.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_smbV2Disk_GetInfos(t *testing.T) {
	tests := []struct {
		name    string
		fields  IVirtualVolume
		want    []Info
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: NewSMBV2Disk("192.168.22.150:445", "qhh", "123123"),
			want: []Info{
				{
					Total:      0,
					Used:       0,
					FSType:     "SMB",
					Path:       "share",
					MountPoint: "share",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			got, err := s.GetInfos()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfos() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_smbV2Disk_Open(t *testing.T) {

	type args struct {
		path string
		name string
		flag int
		perm os.FileMode
	}
	disk := NewSMBV2Disk("192.168.22.150:445", "qhh", "123123")
	infos, err := disk.GetInfos()
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		fields  IVirtualVolume
		args    args
		wantErr bool
	}{
		{
			name:   "test-1",
			fields: disk,
			args: args{
				path: infos[0].Path,
				name: "test.txt",
				flag: os.O_CREATE | os.O_RDWR,
				perm: 0755,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields
			got, err := s.Open(tt.args.path, tt.args.name, tt.args.flag, tt.args.perm)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Error("file is nil")
			}
		})
	}
}
