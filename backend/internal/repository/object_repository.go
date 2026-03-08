package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/xmodmor/XStorage/backend/internal/domain"
)

type objectRepository struct {
	db *gorm.DB
}

func NewObjectRepository(db *gorm.DB) ObjectRepository {
	return &objectRepository{db: db}
}

func (r *objectRepository) Create(ctx context.Context, object *domain.Object) error {
	return r.db.WithContext(ctx).Create(object).Error
}

func (r *objectRepository) FindByID(ctx context.Context, id uint) (*domain.Object, error) {
	var object domain.Object
	if err := r.db.WithContext(ctx).First(&object, id).Error; err != nil {
		return nil, err
	}
	return &object, nil
}

func (r *objectRepository) FindByBucketAndKey(ctx context.Context, bucketID uint, key string) (*domain.Object, error) {
	var object domain.Object
	if err := r.db.WithContext(ctx).Where("bucket_id = ? AND key = ?", bucketID, key).First(&object).Error; err != nil {
		return nil, err
	}
	return &object, nil
}

func (r *objectRepository) ListByBucket(ctx context.Context, bucketID uint, offset, limit int) ([]domain.Object, int64, error) {
	var objects []domain.Object
	var total int64

	db := r.db.WithContext(ctx).Where("bucket_id = ?", bucketID)

	if err := db.Model(&domain.Object{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&objects).Error; err != nil {
		return nil, 0, err
	}

	return objects, total, nil
}

func (r *objectRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Object{}, id).Error
}
