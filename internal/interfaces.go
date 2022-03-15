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
		GetById(context.Context, int) (entity.User, error)
		GetAll(context.Context) ([]entity.User, error)
	}

	// UserRepo -.
	UserRepo interface {
		GetById(context.Context, int) (entity.User, error)
		GetAll(context.Context) ([]entity.User, error)
	}

	// UserWebAPI -.
	WebAPI interface {
	}
)
