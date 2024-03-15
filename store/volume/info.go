package volume

type Info struct {
	ID         string
	Total      uint64
	Used       uint64
	FSType     string
	MountPoint string
}
