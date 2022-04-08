// Package usecase implements application business logic. Each logic group in own file.
package internal

import (
	"context"

	"github.com/Netflix-Clone-MicFlix/User-service/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// User -.
	User interface {
		GetById(context.Context, string) (entity.User, error)
		GetAll(context.Context) ([]entity.User, error)
		Create(context.Context, string) error
		GetAllProfilesById(context.Context, string) ([]entity.Profile, error)
	}

	// UserRepo -.
	UserRepo interface {
		GetById(context.Context, string) (entity.User, error)
		GetAll(context.Context) ([]entity.User, error)
		Create(context.Context, string) (entity.User, error)
		Delete(context.Context, string) error
		Update(context.Context, string, entity.User) error
	}

	// GenreRepo -.
	ProfileRepo interface {
		GetById(context.Context, string) (entity.Profile, error)
		GetAll(context.Context) ([]entity.Profile, error)
		Create(context.Context, entity.Profile) (entity.Profile, error)
		GetAllById(context.Context, string) ([]entity.Profile, error)
		Delete(context.Context, string) error
		Update(context.Context, string, entity.Profile) error
	}

	// GenreRepo -.
	MovieTagRepo interface {
		GetById(context.Context, string) (entity.MovieTag, error)
		GetAll(context.Context) ([]entity.MovieTag, error)
		Create(context.Context, string) (entity.MovieTag, error)
		Delete(context.Context, string) error
		Update(context.Context, string, entity.MovieTag) error
	}
	// UserWebAPI -.
	WebAPI interface {
	}
)
