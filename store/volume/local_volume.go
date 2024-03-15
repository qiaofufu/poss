package volume

import (
	"github.com/shirou/gopsutil/disk"
	"os"
)

type localVolume struct {
	mountPoint string
}

func NewLocalVolume(mountPoint string) IVirtualVolume {
	return &localVolume{
		mountPoint: mountPoint,
	}
}

func (l *localVolume) GetInfo() (Info, error) {
	var (
		info Info
		err  error
	)
	usage, err := disk.Usage(l.mountPoint)
	if err != nil {
		return info, err
	}
	info = Info{
		Total:      usage.Total,
		Used:       usage.Used,
		FSType:     usage.Fstype,
		MountPoint: usage.Path,
	}

	return info, nil
}

func (l *localVolume) Open(filename string, flag int, perm os.FileMode) (IVirtualFile, error) {
	f, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}
	return &localVFile{
		file: f,
	}, nil
}

func (l *localVolume) Close() error {
	return nil
}
