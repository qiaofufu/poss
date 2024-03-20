package io

import (
	"OSS/store/volume"
)

type store struct {
	storePoll map[string]volume.IVirtualVolume
}
