package repositories

import (
	"context"

	"fmt"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

const _defaultEntityCap = 64
const collectionName = "users"

// UserRepo -.
type UserRepo struct {
	*mongodb.MongoDB
}

// New -.
func NewUserRepo(mdb *mongodb.MongoDB) *UserRepo {
	return &UserRepo{mdb}
}

// GetAll -.
func (ur *UserRepo) GetAll(ctx context.Context) ([]entity.User, error) {

	users := []entity.User{}

	collection := ur.Database.Collection(collectionName)

	var filter bson.M = bson.M{}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetAll - rows.Scan: %w", err)
	}
	if err = curr.All(context.Background(), &users); err != nil {
		return nil, fmt.Errorf("UserRepo - GetAll - rows.Scan: %w", err)
	}

	return users, nil
}

// GetById -.
func (r *UserRepo) GetById(ctx context.Context, user_id int) (entity.User, error) {
	entitie := entity.User{}

	return entitie, nil
}
