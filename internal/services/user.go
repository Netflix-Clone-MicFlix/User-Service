package services

import (
	"context"
	"fmt"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	repo   UserRepo
	webAPI UserWebAPI
}

// New -.
func NewUserUseCase(r UserRepo, w WebAPI) *UserUseCase {
	return &UserUseCase{
		repo:   r,
		webAPI: w,
	}
}

// GetById - gets all user by ID -.
func (uc *UserUseCase) GetById(ctx context.Context, user_id int) (entity.User, error) {
	user, err := uc.repo.GetById(ctx, user_id)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - History - s.repo.GetHistory: %w", err)
	}

	return user, nil
}

// Translate -.
func (uc *UserUseCase) GetAll(ctx context.Context) ([]entity.User, error) {

	users, err := uc.repo.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - Translate - s.repo.Store: %w", err)
	}

	return users, nil
}
