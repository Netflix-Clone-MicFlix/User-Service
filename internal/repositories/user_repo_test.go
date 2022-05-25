package repositories

import (
	"context"
	"testing"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal/services"
	"github.com/Netflix-Clone-MicFlix/User-Service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserRepo_GetAll(t *testing.T) {
	usermock := entity.User{"test", "test", []string{"test", "test"}}
	ctx := context.Background()
	userRepo := new(mocks.UserRepo)
	userRepo.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return([]entity.User{usermock}, nil)

	service := services.NewUserUseCase(userRepo, nil, nil, nil)
	result, _ := service.UserRepo.GetAll(ctx)
	assert.Equal(t, result, []entity.User{usermock})
}

func TestUserRepo_GetById(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	usermock := entity.User{id, "test", []string{"test", "test"}}
	ctx := context.Background()
	userRepo := new(mocks.UserRepo)
	userRepo.On("GetById", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(usermock, nil)

	service := services.NewUserUseCase(userRepo, nil, nil, nil)
	result, _ := service.UserRepo.GetById(ctx, id)

	assert.Equal(t, result, usermock)
}

func TestUserRepo_GetByKeycloakId(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	usermock := entity.User{id, "test", []string{"test", "test"}}
	ctx := context.Background()

	userRepo := new(mocks.UserRepo)

	userRepo.On("GetByKeycloakId", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(usermock, nil)

	service := services.NewUserUseCase(userRepo, nil, nil, nil)
	result, _ := service.UserRepo.GetByKeycloakId(ctx, id)

	assert.Equal(t, result, usermock)
}

func TestUserRepo_Create(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	usermock := entity.User{id, "test", []string{"test", "test"}}
	ctx := context.Background()

	userRepo := new(mocks.UserRepo)

	userRepo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(usermock, nil)

	service := services.NewUserUseCase(userRepo, nil, nil, nil)
	result, _ := service.UserRepo.Create(ctx, id)

	assert.Equal(t, result, usermock)
}

func TestUserRepo_Update(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	usermock := entity.User{id, "test", []string{"test", "test"}}
	ctx := context.Background()

	userRepo := new(mocks.UserRepo)

	userRepo.On("Update", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("entity.User")).Return(nil)

	service := services.NewUserUseCase(userRepo, nil, nil, nil)
	result := service.UserRepo.Update(ctx, id, usermock)

	assert.NoError(t, result)
}

func TestUserRepo_Delete(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	ctx := context.Background()

	userRepo := new(mocks.UserRepo)

	userRepo.On("Delete", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(nil)

	service := services.NewUserUseCase(userRepo, nil, nil, nil)
	result := service.UserRepo.Delete(ctx, id)

	assert.NoError(t, result)
}
