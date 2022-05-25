package services

import (
	"context"
	"testing"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func Test(t *testing.T) {

// 	usermock := entity.User{"test", "test", []string{"test", "test"}}
// 	ctx := context.Background()
// 	userRepo := new(mocks.UserRepo)
// 	userRepo.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return([]entity.User{usermock}, nil)

// 	service := NewUserUseCase(userRepo, nil, nil, nil)
// 	result, _ := service.GetAll(ctx)

// 	assert.Equal(t, result, []entity.User{usermock})

// }

func TestUserUseCase_GetById(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	usermock := entity.User{id, "test", []string{"test", "test"}}
	ctx := context.Background()
	userRepo := new(mocks.UserRepo)
	userRepo.On("GetById", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(usermock, nil)

	service := NewUserUseCase(userRepo, nil, nil, nil)
	result, _ := service.GetById(ctx, id)

	assert.Equal(t, result, usermock)
}

func TestUserUseCase_GetAll(t *testing.T) {
	usermock := entity.User{"test", "test", []string{"test", "test"}}
	ctx := context.Background()
	userRepo := new(mocks.UserRepo)
	userRepo.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return([]entity.User{usermock}, nil)

	service := NewUserUseCase(userRepo, nil, nil, nil)
	result, _ := service.GetAll(ctx)

	assert.Equal(t, result, []entity.User{usermock})
}

func TestUserUseCase_Create(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	usermock := entity.User{id, "test", []string{"test", "test"}}
	profilemock := entity.Profile{"test", id, "test", []string{"test", "test"}}
	ctx := context.Background()

	userRepo := new(mocks.UserRepo)
	profileRepo := new(mocks.ProfileRepo)

	profileRepo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("entity.Profile")).Return(profilemock, nil)
	userRepo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(usermock, nil)
	userRepo.On("Update", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("entity.User")).Return(nil)

	service := NewUserUseCase(userRepo, profileRepo, nil, nil)
	result := service.Create(ctx, id)

	assert.NoError(t, result)
}

func TestUserUseCase_GetAllProfilesById(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	profilemock := entity.Profile{"test", id, "test", []string{"test", "test"}}
	ctx := context.Background()

	profileRepo := new(mocks.ProfileRepo)

	profileRepo.On("GetAllById", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return([]entity.Profile{profilemock}, nil)

	service := NewUserUseCase(nil, profileRepo, nil, nil)
	result, _ := service.GetAllProfilesById(ctx, id)

	assert.Equal(t, result, []entity.Profile{profilemock})
}

func TestUserUseCase_Delete(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	usermock := entity.User{id, "test", []string{"test", "test"}}
	profilemock := entity.Profile{"test", id, "test", []string{"test", "test"}}
	ctx := context.Background()

	userRepo := new(mocks.UserRepo)
	profileRepo := new(mocks.ProfileRepo)
	movieTagRepo := new(mocks.MovieTagRepo)

	userRepo.On("GetByKeycloakId", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(usermock, nil)
	profileRepo.On("GetById", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(profilemock, nil)
	movieTagRepo.On("Delete", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(nil)
	profileRepo.On("Delete", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(nil)
	userRepo.On("Delete", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(nil)

	service := NewUserUseCase(userRepo, profileRepo, movieTagRepo, nil)
	result := service.Delete(ctx, id)

	assert.NoError(t, result)
}
