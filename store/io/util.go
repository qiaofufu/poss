package io

import (
	"OSS/store/volume"
	"crypto/md5"
	"encoding/hex"
)

func generateStorageID(info volume.Info) string {
	code := md5.New().Sum([]byte(info.FSType + info.MountPoint))
	return hex.EncodeToString(code)
}
