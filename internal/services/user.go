package services

import (
	"context"
	"fmt"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/security"
	"github.com/google/uuid"
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
		return entity.User{}, fmt.Errorf("UserUseCase - GetById - s.userRepo.GetHistory: %w", err)
	}

	return user, nil
}

// GetAll - gets alls-.
func (uc *UserUseCase) GetAll(ctx context.Context) ([]entity.User, error) {

	users, err := uc.userRepo.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetAll - s.userRepo.Store: %w", err)
	}

	return users, nil
}

// Register - gets alls-.
func (uc *UserUseCase) Register(ctx context.Context, user entity.User) error {

	user.Id = uuid.New().String() //generate guid

	salt, err := uc.saltRepo.Create(context.Background(), user.Id) //generate salt

	if err != nil {
		return fmt.Errorf("UserUseCase - Register - s.userRepo.Store: %w", err)
	}

	err = uc.userRepo.Create(context.Background(), user, salt)
	return err
}

// Login - gets alls-.
func (uc *UserUseCase) Login(ctx context.Context, user entity.User) error {

	userdb, err := uc.userRepo.Login(context.Background(), user)
	if err != nil {
		return fmt.Errorf("UserUseCase - Login - s.userRepo.Store: %w", err)
	}

	salt, err := uc.saltRepo.GetById(context.Background(), userdb.Id)
	if err != nil {
		return fmt.Errorf("UserUseCase - Login - s.userRepo.Store: %w", err)
	}

	if userdb.Email != user.Email && !security.CheckPasswordsMatch(userdb.Password, user.Password, salt.SaltData) {
		return fmt.Errorf("UserRepo - Login Email- rows.Scan: %w", err)
	}

	updatedSalt, err := uc.saltRepo.Update(context.Background(), userdb.Id)
	if err != nil {
		return fmt.Errorf("UserUseCase - Login - s.userRepo.Store: %w", err)
	}

	err = uc.userRepo.Update(context.Background(), userdb.Id, user, updatedSalt)
	if err != nil {
		return fmt.Errorf("UserUseCase - Login - s.userRepo.Store: %w", err)
	}

	return err
}
