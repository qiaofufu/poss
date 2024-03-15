package io

import "OSS/store/volume"

// Storage 屏蔽底层存储,表现为一个逻辑磁盘
type Storage interface {
	GetVolumeInfo() (volume.Info, error)

	Mkdir(path string) error

	Create(filename string) error
	Write(filename string, data []byte) (int, error)
	Read(filename string) ([]byte, error)
	Delete(filename string) error
}

// VolumeManage 负责管理VirtualVolume
type VolumeManage interface {
	// MountVolume 挂载Volume
	MountVolume(mountType string, mountPoint string) (string, error)
	// UnMountVolume 取消挂载Volume
	UnMountVolume(mountPoint string) error
	// ListVolume 列出所有的Volume
	ListVolume() ([]volume.IVirtualVolume, error)
	// GetVolume 根据vid 返回对应的Volume
	GetVolume(vid string) (volume.IVirtualVolume, error)
	// ScheduleVolume 根据传入的分块数量, 自适应地返回相应的volume
	ScheduleVolume(n int) ([]volume.IVirtualVolume, error)
}
