package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/xmodmor/XStorage/backend/internal/domain"
)

type appRepository struct {
	db *gorm.DB
}

func NewAppRepository(db *gorm.DB) AppRepository {
	return &appRepository{db: db}
}

func (r *appRepository) Create(ctx context.Context, app *domain.App) error {
	return r.db.WithContext(ctx).Create(app).Error
}

func (r *appRepository) FindByID(ctx context.Context, id uint) (*domain.App, error) {
	var app domain.App
	if err := r.db.WithContext(ctx).First(&app, id).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *appRepository) ListByOwner(ctx context.Context, ownerID uint) ([]domain.App, error) {
	var apps []domain.App
	if err := r.db.WithContext(ctx).Where("owner_id = ?", ownerID).Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}

func (r *appRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.App{}, id).Error
}
