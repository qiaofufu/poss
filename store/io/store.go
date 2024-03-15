package io

import "OSS/store/volume"

type store struct {
	storePoll map[string]volume.IVirtualVolume
}

func (s *store) MountVolume(mountPoint string) (string, error) {

}

func (s *store) UnMountVolume(mountPoint string) error {
	//TODO implement me
	panic("implement me")
}

func (s *store) ListVolume() ([]volume.IVirtualVolume, error) {
	//TODO implement me
	panic("implement me")
}

func (s *store) GetVolume(vid string) (volume.IVirtualVolume, error) {
	//TODO implement me
	panic("implement me")
}

func (s *store) ScheduleVolume(n int) ([]volume.IVirtualVolume, error) {
	//TODO implement me
	panic("implement me")
}
