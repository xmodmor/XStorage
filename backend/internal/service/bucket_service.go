package service

import (
	"context"

	"github.com/xmodmor/XStorage/backend/internal/domain"
	"github.com/xmodmor/XStorage/backend/internal/dto"
	"github.com/xmodmor/XStorage/backend/internal/repository"
)

type BucketService struct {
	bucketRepo repository.BucketRepository
}

func NewBucketService(bucketRepo repository.BucketRepository) *BucketService {
	return &BucketService{bucketRepo: bucketRepo}
}

func (s *BucketService) Create(ctx context.Context, appID uint, req dto.CreateBucketRequest) (*domain.Bucket, error) {
	existing, _ := s.bucketRepo.FindByAppAndName(ctx, appID, req.Name)
	if existing != nil {
		return nil, ErrConflict
	}

	bucket := &domain.Bucket{
		AppID:      appID,
		Name:       req.Name,
		Visibility: req.Visibility,
	}

	if err := s.bucketRepo.Create(ctx, bucket); err != nil {
		return nil, err
	}

	return bucket, nil
}

func (s *BucketService) ListByApp(ctx context.Context, appID uint) ([]domain.Bucket, error) {
	return s.bucketRepo.ListByApp(ctx, appID)
}

func (s *BucketService) Delete(ctx context.Context, appID uint, bucketName string) error {
	bucket, err := s.bucketRepo.FindByAppAndName(ctx, appID, bucketName)
	if err != nil {
		return ErrNotFound
	}

	return s.bucketRepo.Delete(ctx, bucket.ID)
}
