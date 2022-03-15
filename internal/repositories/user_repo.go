package repositories

import (
	"context"

	"fmt"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/mongodb"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/security"
	"go.mongodb.org/mongo-driver/bson"
)

const userCollectionName = "users"

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

	collection := ur.Database.Collection(userCollectionName)

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
func (ur *UserRepo) GetById(ctx context.Context, user_id string) (entity.User, error) {
	user := entity.User{}

	var filter bson.M = bson.M{"_id": user_id}
	curr, err := ur.Database.Collection(userCollectionName).Find(context.Background(), filter)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetById - rows.Scan: %w", err)
	}
	defer curr.Close(context.Background())

	curr.All(context.Background(), &user)

	return user, nil
}

// Create -.
func (ur *UserRepo) Create(ctx context.Context, user entity.User, salt []byte) error {
	var hashedPassword = security.HashPassword(user.Password, salt)

	user.Password = hashedPassword
	_, err := ur.Database.Collection(userCollectionName).InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("UserRepo - Create - rows.Scan: %w", err)
	}
	return nil
}

// Update -.
func (ur *UserRepo) Update(ctx context.Context, user_id string, user entity.User, salt []byte) error {
	_, err := ur.Database.Collection(userCollectionName).UpdateOne(
		context.Background(),
		bson.M{"_id": user_id},
		user)

	if err != nil {
		return fmt.Errorf("UserRepo - Create - rows.Scan: %w", err)
	}
	return nil
}

// Delete -.
func (ur *UserRepo) Delete(ctx context.Context, user_id string) error {
	_, err := ur.Database.Collection(userCollectionName).DeleteOne(
		context.Background(),
		bson.M{"_id": user_id})

	if err != nil {
		return fmt.Errorf("UserRepo - Delete - rows.Scan: %w", err)
	}
	return nil
}

// Login -.
func (ur *UserRepo) Login(ctx context.Context, user entity.User) (entity.User, error) {
	userdb := entity.User{}

	var filter bson.M = bson.M{"email": user.Email}
	err := ur.Database.Collection(userCollectionName).FindOne(context.Background(), filter).Decode(&userdb)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - Login - rows.Scan: %w", err)
	}

	return userdb, nil
}
