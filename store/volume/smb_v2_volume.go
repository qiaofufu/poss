package volume

import (
	"fmt"
	"github.com/hirochachacha/go-smb2"
	"net"
	"os"
)

type smbV2Volume struct {
	s          *smb2.Session
	mountPoint string
	address    string
}

func NewSMBV2Disk(address string, username, password, mountPoint string) IVirtualVolume {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     username,
			Password: password,
		},
	}
	session, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	disk := &smbV2Volume{s: session, mountPoint: mountPoint, address: address}
	return disk
}

func (s *smbV2Volume) GetInfo() (Info, error) {
	share, err := s.s.Mount(s.mountPoint)
	if err != nil {
		return Info{}, err
	}

	statFs, err := share.Statfs("")
	if err != nil {
		return Info{}, err
	}

	info := Info{
		Total:      statFs.TotalBlockCount() * statFs.BlockSize(),
		Used:       statFs.BlockSize() * (statFs.TotalBlockCount() - statFs.AvailableBlockCount()),
		FSType:     "SMB",
		MountPoint: s.address,
	}

	return info, nil
}

func (s *smbV2Volume) Open(filename string, flag int, perm os.FileMode) (IVirtualFile, error) {
	share, err := s.s.Mount(s.mountPoint)
	if err != nil {
		return nil, fmt.Errorf("mount share failed, err: %v", err)
	}

	file, err := share.OpenFile(filename, flag, perm)
	return &smbV2File{
		f: file,
	}, nil
}

func (s *smbV2Volume) Close() error {
	return s.s.Logoff()
}
