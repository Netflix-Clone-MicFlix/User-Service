package repositories

import (
	"context"

	"fmt"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const UserCollectionName = "Users"

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

	Users := []entity.User{}

	collection := ur.Database.Collection(UserCollectionName)

	var filter bson.M = bson.M{}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetAll - rows.Scan: %w", err)
	}
	if err = curr.All(context.Background(), &Users); err != nil {
		return nil, fmt.Errorf("UserRepo - GetAll - rows.Scan: %w", err)
	}

	return Users, nil
}

// GetById -.
func (ur *UserRepo) GetById(ctx context.Context, User_id string) (entity.User, error) {
	User := entity.User{}

	var filter bson.M = bson.M{"id": User_id}
	curr, err := ur.Database.Collection(UserCollectionName).Find(context.Background(), filter)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetById - rows.Scan: %w", err)
	}
	defer curr.Close(context.Background())

	curr.All(context.Background(), &User)

	return User, nil
}

// Create -.
func (ur *UserRepo) Create(ctx context.Context, keycloak_id string) (entity.User, error) {

	guid := uuid.New().String()
	User := entity.User{
		Id:         guid,
		KeycloakId: keycloak_id,
	}

	_, err := ur.Database.Collection(UserCollectionName).InsertOne(context.Background(), User)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - Create - rows.Scan: %w", err)
	}
	return User, nil
}

// Update -.
func (ur *UserRepo) Update(ctx context.Context, User_id string, User entity.User) error {

	User.Id = User_id
	update := bson.M{"$set": User}

	_, err := ur.Database.Collection(UserCollectionName).UpdateOne(
		context.Background(),
		bson.M{"id": User_id},
		update)

	if err != nil {
		return fmt.Errorf("UserRepo - Update - rows.Scan: %w", err)
	}
	return nil
}

// Delete -.
func (ur *UserRepo) Delete(ctx context.Context, User_id string) error {
	_, err := ur.Database.Collection(UserCollectionName).DeleteOne(
		context.Background(),
		bson.M{"id": User_id})

	if err != nil {
		return fmt.Errorf("UserRepo - Delete - rows.Scan: %w", err)
	}
	return nil
}
