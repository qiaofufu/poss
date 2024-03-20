package meta

type ACLType uint8

const (
	ACLTypeUnknown ACLType = iota
	ACLTypeBucket
	ACLTypeObject
)

type ACLPermission uint16

const (
	ReadPermission           ACLPermission = 1 << 7
	WritePermission                        = 1 << 6
	AnonymousReadPermission                = 1 << 5
	AnonymousWritePermission               = 1 << 4
)

var permissions map[ACLPermission]string

func init() {
	permissions = make(map[ACLPermission]string)
	permissions[ReadPermission] = "读取权限"
	permissions[WritePermission] = "写入权限"
	permissions[AnonymousReadPermission] = "匿名读取权限"
	permissions[AnonymousWritePermission] = "匿名写入权限"
}
