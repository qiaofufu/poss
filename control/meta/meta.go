package meta

import (
	"OSS/control/model/meta"
	"gorm.io/gorm"
)

type Meta struct {
	db *gorm.DB
}

func (m *Meta) CreateBucket(name string, permission meta.ACLPermission, userID string) (string, error) {
	bucket := &meta.Bucket{
		Name:    name,
		OwnerID: userID,
	}
	err := m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(bucket).Error; err != nil {
			return err
		}
		acl := &meta.ACL{
			Type:       meta.ACLTypeBucket,
			BucketID:   bucket.ID,
			Permission: permission,
		}
		if err := tx.Create(bucket).Error; err != nil {
			return err
		}

		bucket.ACL = *acl
		return tx.Save(bucket).Error
	})
	if err != nil {
		return "", err
	}
	return bucket.ID, nil
}

func (m *Meta) UpdateBucket(id string, name string, permissions meta.ACLPermission) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		bucket := &meta.Bucket{
			ID:   id,
			Name: name,
		}
		if err := tx.Model(bucket).Updates(bucket).Error; err != nil {
			return err
		}
		acl := &meta.ACL{
			BucketID:   bucket.ID,
			Permission: permissions,
		}
		return tx.Model(acl).Updates(acl).Error
	})
}

func (m *Meta) DelBucket(id string) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		bucket := &meta.Bucket{
			ID: id,
		}
		if err := tx.Delete(bucket).Error; err != nil {
			return err
		}
		acl := &meta.ACL{
			BucketID: bucket.ID,
		}
		return tx.Delete(acl).Error
	})
}

func (m *Meta) ListBuckets(userID string) ([]*meta.Bucket, error) {
	buckets := make([]*meta.Bucket, 0)
	err := m.db.Where("owner_id = ?", userID).Find(&buckets).Error
	if err != nil {
		return nil, err
	}
	return buckets, nil
}

func (m *Meta) GetBucketByID(id string, preloadObjects bool, preloadACL bool) (*meta.Bucket, error) {
	bucket := &meta.Bucket{}
	sql := m.db.Where("id = ?", id)
	// preload
	if preloadObjects {
		sql.Preload("Objects")
	}
	if preloadACL {
		sql.Preload("ACL")
	}

	err := sql.First(bucket).Error
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func (m *Meta) CreateObject(bucketID string, size int64, contentType string, key string, permission uint16) (string, error) {
	object := &meta.Object{
		BucketID:    bucketID,
		Size:        size,
		ContentType: contentType,
		Key:         key,
	}
	err := m.db.Create(object).Error
	if err != nil {
		return "", err
	}
	return object.ID, nil
}

func (m *Meta) ListObjects(bucketID string) ([]*meta.Object, error) {
	objects := make([]*meta.Object, 0)
	err := m.db.Where("bucket_id = ?", bucketID).Find(&objects).Error
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func (m *Meta) DelObject(bucketID, objectID string) error {
	object := &meta.Object{
		ID:       objectID,
		BucketID: bucketID,
	}
	return m.db.Delete(object).Error
}
