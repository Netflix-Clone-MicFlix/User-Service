package repositories

import (
	"context"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/postgres"
)

const _defaultEntityCap = 64

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// GetHistory -.
func (r *UserRepo) GetAll(ctx context.Context) ([]entity.User, error) {

	entities := []entity.User{}

	return entities, nil
}

// Store -.
func (r *UserRepo) GetById(ctx context.Context, user_id int) (entity.User, error) {
	entitie := entity.User{}

	return entitie, nil
}
