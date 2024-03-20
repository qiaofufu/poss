package meta

import "OSS/control/model/user"

type Bucket struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
	// Belong To
	OwnerID string    `json:"ownerID,omitempty"`
	Owner   user.User `json:"owner"`
	// Has Many
	Objects []Object `json:"objects"`
	// Has One
	ACL ACL `json:"ACL"`
}
