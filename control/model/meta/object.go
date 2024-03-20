package meta

import "time"

// delete object struct json tag omitempty
type Object struct {
	ID           string    `json:"id"`            // ID 内部唯一ID
	CreateAt     time.Time `json:"create_at"`     // CreateAt 创建时间
	Key          string    `json:"key"`           // Key 桶内唯一标识
	Size         int64     `json:"size"`          // Size 对象尺寸
	ContentType  string    `json:"content_type"`  // ContentType 对象类型
	LastModified time.Time `json:"last_modified"` // LastModified 最后修改时间
	Etag         string    `json:"etag"`          // Etag 实体标签
	BucketID     string    `json:"bucket_id"`     // BucketID 桶ID

	// Has Many
	Blocks []Block `json:"blocks"`

	// Has one
	ACL ACL `json:"ACL"` // ACL 访问控制列表, 用于控制对象的访问权限
}
