package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	UserRepo     internal.UserRepo
	ProfileRepo  internal.ProfileRepo
	MovieTagRepo internal.MovieTagRepo
	webAPI       internal.WebAPI
}

// New -.
func NewUserUseCase(u internal.UserRepo, p internal.ProfileRepo, m internal.MovieTagRepo, w internal.WebAPI) *UserUseCase {
	return &UserUseCase{
		UserRepo:     u,
		ProfileRepo:  p,
		MovieTagRepo: m,
		webAPI:       w,
	}
}

// GetById - gets all User by ID -.
func (uc *UserUseCase) GetById(ctx context.Context, User_id string) (entity.User, error) {
	User, err := uc.UserRepo.GetById(ctx, User_id)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - GetById - s.UserRepo.GetHistory: %w", err)
	}

	return User, nil
}

// GetAll - gets alls-.
func (uc *UserUseCase) GetAll(ctx context.Context) ([]entity.User, error) {

	Users, err := uc.UserRepo.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetAll - s.UserRepo.Store: %w", err)
	}

	return Users, nil
}

// Create - Create genre-.
func (uc *UserUseCase) Create(ctx context.Context, keycloak_id string) error {

	if keycloak_id == "" {
		return fmt.Errorf("UserUseCase - Create - s.UserRepo.Store: No values provided")
	}

	user, err := uc.UserRepo.Create(context.Background(), keycloak_id)
	if err != nil {
		return fmt.Errorf("UserUseCase - Create - s.UserRepo.Store: %w", err)
	}

	maxProfileAmount := 5
	profileIdArray := []string{}

	for i := 0; i < maxProfileAmount; i++ {

		name := "profile_" + strconv.Itoa(i)

		profile := entity.Profile{
			UserId:      user.KeycloakId,
			Name:        name,
			MovieTagIds: []string{},
		}

		result, err := uc.ProfileRepo.Create(context.Background(), profile)
		if err != nil {
			return fmt.Errorf("UserUseCase - Create - s.UserRepo.Store: %w", err)
		}
		profileIdArray = append(profileIdArray, result.Id)
	}

	user.ProfileIds = profileIdArray

	err = uc.UserRepo.Update(context.Background(), user.Id, user)
	if err != nil {
		return fmt.Errorf("UserUseCase - Create - s.UserRepo.Store: %w", err)
	}

	return nil
}

// GetById - gets all User by ID -.
func (uc *UserUseCase) GetAllProfilesById(ctx context.Context, User_id string) ([]entity.Profile, error) {
	Profiles, err := uc.ProfileRepo.GetAllById(ctx, User_id)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetById - s.UserRepo.GetHistory: %w", err)
	}

	return Profiles, nil
}
