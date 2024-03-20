package meta

type ACL struct {
	ID         string        `json:"id"`
	Type       ACLType       `json:"type"`
	BucketID   string        `json:"bucket_id"`
	ObjectID   string        `json:"object_id"`
	Permission ACLPermission `json:"permission"`
}

func (a *ACL) ConvertPermissionToString(permission ACLPermission) string {
	label, ok := permissions[permission]
	if !ok {
		return ""
	}
	return label
}
