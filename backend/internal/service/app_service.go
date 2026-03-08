package service

import (
	"context"

	"github.com/xmodmor/XStorage/backend/internal/domain"
	"github.com/xmodmor/XStorage/backend/internal/dto"
	"github.com/xmodmor/XStorage/backend/internal/repository"
)

type AppService struct {
	appRepo repository.AppRepository
}

func NewAppService(appRepo repository.AppRepository) *AppService {
	return &AppService{appRepo: appRepo}
}

func (s *AppService) Create(ctx context.Context, ownerID uint, req dto.CreateAppRequest) (*domain.App, error) {
	app := &domain.App{
		Name:    req.Name,
		OwnerID: ownerID,
	}

	if err := s.appRepo.Create(ctx, app); err != nil {
		return nil, err
	}

	return app, nil
}

func (s *AppService) GetByID(ctx context.Context, id uint) (*domain.App, error) {
	return s.appRepo.FindByID(ctx, id)
}

func (s *AppService) ListByOwner(ctx context.Context, ownerID uint) ([]domain.App, error) {
	return s.appRepo.ListByOwner(ctx, ownerID)
}

func (s *AppService) Delete(ctx context.Context, id, ownerID uint) error {
	app, err := s.appRepo.FindByID(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	if app.OwnerID != ownerID {
		return ErrForbidden
	}

	return s.appRepo.Delete(ctx, id)
}
