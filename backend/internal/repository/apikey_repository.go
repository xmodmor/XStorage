package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/xmodmor/XStorage/backend/internal/domain"
)

type apiKeyRepository struct {
	db *gorm.DB
}

func NewAPIKeyRepository(db *gorm.DB) APIKeyRepository {
	return &apiKeyRepository{db: db}
}

func (r *apiKeyRepository) Create(ctx context.Context, key *domain.APIKey) error {
	return r.db.WithContext(ctx).Create(key).Error
}

func (r *apiKeyRepository) FindByAccessKey(ctx context.Context, accessKey string) (*domain.APIKey, error) {
	var key domain.APIKey
	if err := r.db.WithContext(ctx).Where("access_key = ?", accessKey).First(&key).Error; err != nil {
		return nil, err
	}
	return &key, nil
}

func (r *apiKeyRepository) ListByApp(ctx context.Context, appID uint) ([]domain.APIKey, error) {
	var keys []domain.APIKey
	if err := r.db.WithContext(ctx).Where("app_id = ?", appID).Find(&keys).Error; err != nil {
		return nil, err
	}
	return keys, nil
}

func (r *apiKeyRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.APIKey{}, id).Error
}
