package repositories

import (
	"context"
	"testing"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMovieTagRepo_GetAll(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	movieTagmock := entity.MovieTag{id, "test", 5}
	ctx := context.Background()

	movieTagRepo := new(mocks.MovieTagRepo)

	movieTagRepo.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return([]entity.MovieTag{movieTagmock}, nil)

	result, _ := movieTagRepo.GetAll(ctx)

	assert.Equal(t, result, []entity.MovieTag{movieTagmock})
}

func TestMovieTagRepo_GetById(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	movieTagmock := entity.MovieTag{id, "test", 5}
	ctx := context.Background()

	movieTagRepo := new(mocks.MovieTagRepo)

	movieTagRepo.On("GetById", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(movieTagmock, nil)

	// service := services.NewUserUseCase(nil, nil, movieTagRepo, nil)
	result, _ := movieTagRepo.GetById(ctx, id)

	assert.Equal(t, result, movieTagmock)
}

func TestMovieTagRepo_Create(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	movieTagmock := entity.MovieTag{id, "test", 5}
	ctx := context.Background()

	movieTagRepo := new(mocks.MovieTagRepo)

	movieTagRepo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(movieTagmock, nil)

	result, _ := movieTagRepo.Create(ctx, id)

	assert.Equal(t, result, movieTagmock)
}

func TestMovieTagRepo_Update(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	movieTagmock := entity.MovieTag{id, "test", 5}
	ctx := context.Background()

	movieTagRepo := new(mocks.MovieTagRepo)

	movieTagRepo.On("Update", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"), mock.AnythingOfType("entity.MovieTag")).Return(nil)

	result := movieTagRepo.Update(ctx, id, movieTagmock)

	assert.NoError(t, result)
}

func TestMovieTagRepo_Delete(t *testing.T) {
	id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	// movieTagmock := entity.MovieTag{id, "test", 5}
	ctx := context.Background()

	movieTagRepo := new(mocks.MovieTagRepo)

	movieTagRepo.On("Delete", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).Return(nil)

	result := movieTagRepo.Delete(ctx, id)

	assert.NoError(t, result)
}
