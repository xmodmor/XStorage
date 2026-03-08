package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/xmodmor/XStorage/backend/internal/domain"
)

type bucketRepository struct {
	db *gorm.DB
}

func NewBucketRepository(db *gorm.DB) BucketRepository {
	return &bucketRepository{db: db}
}

func (r *bucketRepository) Create(ctx context.Context, bucket *domain.Bucket) error {
	return r.db.WithContext(ctx).Create(bucket).Error
}

func (r *bucketRepository) FindByID(ctx context.Context, id uint) (*domain.Bucket, error) {
	var bucket domain.Bucket
	if err := r.db.WithContext(ctx).First(&bucket, id).Error; err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (r *bucketRepository) FindByAppAndName(ctx context.Context, appID uint, name string) (*domain.Bucket, error) {
	var bucket domain.Bucket
	if err := r.db.WithContext(ctx).Where("app_id = ? AND name = ?", appID, name).First(&bucket).Error; err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (r *bucketRepository) ListByApp(ctx context.Context, appID uint) ([]domain.Bucket, error) {
	var buckets []domain.Bucket
	if err := r.db.WithContext(ctx).Where("app_id = ?", appID).Find(&buckets).Error; err != nil {
		return nil, err
	}
	return buckets, nil
}

func (r *bucketRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Bucket{}, id).Error
}
