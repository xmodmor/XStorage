package repository

import (
	"context"

	"github.com/xmodmor/XStorage/backend/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	List(ctx context.Context, offset, limit int, search string) ([]domain.User, int64, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uint) error
}

type AppRepository interface {
	Create(ctx context.Context, app *domain.App) error
	FindByID(ctx context.Context, id uint) (*domain.App, error)
	ListByOwner(ctx context.Context, ownerID uint) ([]domain.App, error)
	Delete(ctx context.Context, id uint) error
}

type APIKeyRepository interface {
	Create(ctx context.Context, key *domain.APIKey) error
	FindByAccessKey(ctx context.Context, accessKey string) (*domain.APIKey, error)
	ListByApp(ctx context.Context, appID uint) ([]domain.APIKey, error)
	Delete(ctx context.Context, id uint) error
}

type BucketRepository interface {
	Create(ctx context.Context, bucket *domain.Bucket) error
	FindByID(ctx context.Context, id uint) (*domain.Bucket, error)
	FindByAppAndName(ctx context.Context, appID uint, name string) (*domain.Bucket, error)
	ListByApp(ctx context.Context, appID uint) ([]domain.Bucket, error)
	Delete(ctx context.Context, id uint) error
}

type ObjectRepository interface {
	Create(ctx context.Context, object *domain.Object) error
	FindByID(ctx context.Context, id uint) (*domain.Object, error)
	FindByBucketAndKey(ctx context.Context, bucketID uint, key string) (*domain.Object, error)
	ListByBucket(ctx context.Context, bucketID uint, offset, limit int) ([]domain.Object, int64, error)
	Delete(ctx context.Context, id uint) error
}
