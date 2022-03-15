package services

import (
	"context"
	"fmt"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	userRepo internal.UserRepo
	saltRepo internal.SaltRepo
	webAPI   internal.WebAPI
}

// New -.
func NewUserUseCase(r internal.UserRepo, s internal.SaltRepo, w internal.WebAPI) *UserUseCase {
	return &UserUseCase{
		userRepo: r,
		saltRepo: s,
		webAPI:   w,
	}
}

// GetById - gets all user by ID -.
func (uc *UserUseCase) GetById(ctx context.Context, user_id string) (entity.User, error) {
	user, err := uc.userRepo.GetById(ctx, user_id)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - History - s.userRepo.GetHistory: %w", err)
	}

	return user, nil
}

// GetAll - gets alls-.
func (uc *UserUseCase) GetAll(ctx context.Context) ([]entity.User, error) {

	users, err := uc.userRepo.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - Translate - s.userRepo.Store: %w", err)
	}

	return users, nil
}

// Register - gets alls-.
func (uc *UserUseCase) Register(ctx context.Context, user entity.User) error {

	err := uc.userRepo.Create(context.Background(), user, nil)
	return err
}

// Login - gets alls-.
func (uc *UserUseCase) Login(ctx context.Context, user entity.User) error {
	err := uc.userRepo.Login(context.Background(), user, nil)
	//call keyclock for token

	return err
}
