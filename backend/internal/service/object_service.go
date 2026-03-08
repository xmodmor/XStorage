package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/xmodmor/XStorage/backend/internal/domain"
	"github.com/xmodmor/XStorage/backend/internal/repository"
	"github.com/xmodmor/XStorage/backend/internal/storage"
)

type ObjectService struct {
	objectRepo repository.ObjectRepository
	bucketRepo repository.BucketRepository
	storage    storage.Storage
}

func NewObjectService(
	objectRepo repository.ObjectRepository,
	bucketRepo repository.BucketRepository,
	storage storage.Storage,
) *ObjectService {
	return &ObjectService{
		objectRepo: objectRepo,
		bucketRepo: bucketRepo,
		storage:    storage,
	}
}

func (s *ObjectService) Upload(ctx context.Context, appID uint, bucketName, key, mime string, reader io.Reader) (*domain.Object, error) {
	bucket, err := s.bucketRepo.FindByAppAndName(ctx, appID, bucketName)
	if err != nil {
		return nil, ErrNotFound
	}

	storagePath := fmt.Sprintf("%d/%s/%s", appID, bucketName, key)

	hasher := sha256.New()
	tee := io.TeeReader(reader, hasher)

	counter := &countingReader{reader: tee}
	if err := s.storage.Put(ctx, storagePath, counter); err != nil {
		return nil, fmt.Errorf("store object: %w", err)
	}

	checksum := fmt.Sprintf("%x", hasher.Sum(nil))

	existing, _ := s.objectRepo.FindByBucketAndKey(ctx, bucket.ID, key)
	if existing != nil {
		_ = s.objectRepo.Delete(ctx, existing.ID)
	}

	object := &domain.Object{
		BucketID:    bucket.ID,
		Key:         key,
		Size:        counter.count,
		Mime:        mime,
		StoragePath: storagePath,
		Checksum:    checksum,
	}

	if err := s.objectRepo.Create(ctx, object); err != nil {
		return nil, err
	}

	return object, nil
}

func (s *ObjectService) Download(ctx context.Context, appID uint, bucketName, key string) (io.ReadCloser, *domain.Object, error) {
	bucket, err := s.bucketRepo.FindByAppAndName(ctx, appID, bucketName)
	if err != nil {
		return nil, nil, ErrNotFound
	}

	object, err := s.objectRepo.FindByBucketAndKey(ctx, bucket.ID, key)
	if err != nil {
		return nil, nil, ErrNotFound
	}

	reader, err := s.storage.Get(ctx, object.StoragePath)
	if err != nil {
		return nil, nil, err
	}

	return reader, object, nil
}

func (s *ObjectService) Delete(ctx context.Context, appID uint, bucketName, key string) error {
	bucket, err := s.bucketRepo.FindByAppAndName(ctx, appID, bucketName)
	if err != nil {
		return ErrNotFound
	}

	object, err := s.objectRepo.FindByBucketAndKey(ctx, bucket.ID, key)
	if err != nil {
		return ErrNotFound
	}

	if err := s.storage.Delete(ctx, object.StoragePath); err != nil {
		return err
	}

	return s.objectRepo.Delete(ctx, object.ID)
}

func (s *ObjectService) List(ctx context.Context, appID uint, bucketName string, page, perPage int) ([]domain.Object, int64, error) {
	bucket, err := s.bucketRepo.FindByAppAndName(ctx, appID, bucketName)
	if err != nil {
		return nil, 0, ErrNotFound
	}

	offset := (page - 1) * perPage
	return s.objectRepo.ListByBucket(ctx, bucket.ID, offset, perPage)
}

type countingReader struct {
	reader io.Reader
	count  int64
}

func (cr *countingReader) Read(p []byte) (int, error) {
	n, err := cr.reader.Read(p)
	cr.count += int64(n)
	return n, err
}
