package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/xmodmor/XStorage/backend/internal/domain"
	"github.com/xmodmor/XStorage/backend/internal/dto"
	"github.com/xmodmor/XStorage/backend/internal/repository"
)

type APIKeyService struct {
	apiKeyRepo repository.APIKeyRepository
	appRepo    repository.AppRepository
}

func NewAPIKeyService(apiKeyRepo repository.APIKeyRepository, appRepo repository.AppRepository) *APIKeyService {
	return &APIKeyService{apiKeyRepo: apiKeyRepo, appRepo: appRepo}
}

func (s *APIKeyService) Create(ctx context.Context, appID, ownerID uint, req dto.CreateAPIKeyRequest) (*domain.APIKey, error) {
	app, err := s.appRepo.FindByID(ctx, appID)
	if err != nil {
		return nil, ErrNotFound
	}
	if app.OwnerID != ownerID {
		return nil, ErrForbidden
	}

	permissions := req.Permissions
	if permissions == "" {
		permissions = "read,write"
	}

	accessKey, err := generateRandomHex(16)
	if err != nil {
		return nil, err
	}
	secretKey, err := generateRandomHex(32)
	if err != nil {
		return nil, err
	}

	key := &domain.APIKey{
		AppID:       appID,
		AccessKey:   accessKey,
		SecretKey:   secretKey,
		Permissions: permissions,
	}

	if err := s.apiKeyRepo.Create(ctx, key); err != nil {
		return nil, err
	}

	return key, nil
}

func (s *APIKeyService) ListByApp(ctx context.Context, appID, ownerID uint) ([]domain.APIKey, error) {
	app, err := s.appRepo.FindByID(ctx, appID)
	if err != nil {
		return nil, ErrNotFound
	}
	if app.OwnerID != ownerID {
		return nil, ErrForbidden
	}

	return s.apiKeyRepo.ListByApp(ctx, appID)
}

func (s *APIKeyService) Delete(ctx context.Context, id, appID, ownerID uint) error {
	app, err := s.appRepo.FindByID(ctx, appID)
	if err != nil {
		return ErrNotFound
	}
	if app.OwnerID != ownerID {
		return ErrForbidden
	}

	return s.apiKeyRepo.Delete(ctx, id)
}

func generateRandomHex(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
