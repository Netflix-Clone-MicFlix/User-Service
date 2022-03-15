// Package usecase implements application business logic. Each logic group in own file.
package internal

import (
	"context"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// User -.
	User interface {
		GetById(context.Context, string) (entity.User, error)
		GetAll(context.Context) ([]entity.User, error)
		Register(context.Context, entity.User) error
		Login(context.Context, entity.User) error
	}

	// UserRepo -.
	UserRepo interface {
		GetAll(context.Context) ([]entity.User, error)
		GetById(context.Context, string) (entity.User, error)
		Create(context.Context, entity.User, []byte) error
		Update(context.Context, string, entity.User, []byte) error
		Delete(context.Context, string) error
		Login(context.Context, entity.User, []byte) error
	}

	// UserRepo -.
	SaltRepo interface {
		GetById(context.Context, string) (entity.User, error)
		Create(context.Context, string) error
		Delete(context.Context, string) error
		Update(context.Context, string) error
	}
	// UserWebAPI -.
	WebAPI interface {
	}
)
