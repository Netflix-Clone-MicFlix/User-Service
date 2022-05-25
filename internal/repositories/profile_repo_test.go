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

func TestProfileRepo_GetAll(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	profilemock := entity.Profile{"test", id, "test", []string{"test", "test"}}
	ctx := context.Background()

	profileRepo := new(mocks.ProfileRepo)

	profileRepo.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return([]entity.Profile{profilemock}, nil)

	service := services.NewUserUseCase(nil, profileRepo, nil, nil)
	result, _ := service.ProfileRepo.GetAll(ctx)

	assert.Equal(t, result, []entity.Profile{profilemock})
}

func TestProfileRepo_GetById(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	profilemock := entity.Profile{"test", id, "test", []string{"test", "test"}}
	ctx := context.Background()

	profileRepo := new(mocks.ProfileRepo)

	profileRepo.On("GetById", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(profilemock, nil)

	service := services.NewUserUseCase(nil, profileRepo, nil, nil)
	result, _ := service.ProfileRepo.GetById(ctx, id)

	assert.Equal(t, result, profilemock)
}

func TestProfileRepo_GetAllById(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	profilemock := entity.Profile{"test", id, "test", []string{"test", "test"}}
	ctx := context.Background()

	profileRepo := new(mocks.ProfileRepo)

	profileRepo.On("GetAllById", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return([]entity.Profile{profilemock}, nil)

	service := services.NewUserUseCase(nil, profileRepo, nil, nil)
	result, _ := service.ProfileRepo.GetAllById(ctx, id)

	assert.Equal(t, result, []entity.Profile{profilemock})
}

func TestProfileRepo_Create(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	profilemock := entity.Profile{"test", id, "test", []string{"test", "test"}}
	ctx := context.Background()

	profileRepo := new(mocks.ProfileRepo)

	profileRepo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("entity.Profile")).Return(profilemock, nil)

	service := services.NewUserUseCase(nil, profileRepo, nil, nil)
	result, _ := service.ProfileRepo.Create(ctx, profilemock)

	assert.Equal(t, result, profilemock)
}

func TestProfileRepo_Update(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	profilemock := entity.Profile{"test", id, "test", []string{"test", "test"}}
	ctx := context.Background()

	profileRepo := new(mocks.ProfileRepo)

	profileRepo.On("Update", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("entity.Profile")).Return(nil)

	service := services.NewUserUseCase(nil, profileRepo, nil, nil)
	result := service.ProfileRepo.Update(ctx, id, profilemock)

	assert.NoError(t, result)
}

func TestProfileRepo_Delete(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	// profilemock := entity.Profile{"test", id, "test", []string{"test", "test"}}
	ctx := context.Background()

	profileRepo := new(mocks.ProfileRepo)

	profileRepo.On("Delete", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(nil)

	service := services.NewUserUseCase(nil, profileRepo, nil, nil)
	result := service.ProfileRepo.Delete(ctx, id)

	assert.NoError(t, result)
}
