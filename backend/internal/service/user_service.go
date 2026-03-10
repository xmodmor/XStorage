package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/xmodmor/XStorage/backend/internal/domain"
	"github.com/xmodmor/XStorage/backend/internal/dto"
	"github.com/xmodmor/XStorage/backend/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) List(ctx context.Context, page, limit int, search string) (*dto.ListUsersResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	offset := (page - 1) * limit

	users, total, err := s.userRepo.List(ctx, offset, limit, search)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserResponse, len(users))
	for i, u := range users {
		result[i] = dto.UserResponse{
			ID:        u.ID,
			Email:     u.Email,
			CreatedAt: u.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	return &dto.ListUsersResponse{
		Users: result,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *UserService) Create(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, ErrConflict
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        req.Email,
		PasswordHash: string(hash),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (s *UserService) GetByID(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (s *UserService) Update(ctx context.Context, id uint, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	if req.Email != "" && req.Email != user.Email {
		existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
		if existing != nil && existing.ID != id {
			return nil, ErrConflict
		}
		user.Email = req.Email
	}

	if req.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = string(hash)
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (s *UserService) Delete(ctx context.Context, id uint) error {
	_, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	return s.userRepo.Delete(ctx, id)
}
