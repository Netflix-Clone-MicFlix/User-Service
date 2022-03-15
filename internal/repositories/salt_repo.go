package repositories

import (
	"context"

	"fmt"
	"time"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/mongodb"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/security"
	"go.mongodb.org/mongo-driver/bson"
)

const saltCollectionName = "salt"

// SaltRepo -.
type SaltRepo struct {
	*mongodb.MongoDB
}

// New -.
func NewSaltRepo(mdb *mongodb.MongoDB) *SaltRepo {
	return &SaltRepo{mdb}
}

// GetById -.
func (sr *SaltRepo) GetById(ctx context.Context, user_id string) (entity.Salt, error) {
	user := entity.Salt{}

	var filter bson.M = bson.M{"user_id": user_id}
	curr, err := sr.Database.Collection(saltCollectionName).Find(context.Background(), filter)
	if err != nil {
		return entity.Salt{}, fmt.Errorf("SaltRepo - GetById - rows.Scan: %w", err)
	}
	defer curr.Close(context.Background())

	curr.All(context.Background(), &user)

	return user, nil
}

// Create -.
func (sr *SaltRepo) Create(ctx context.Context, user_id string) error {

	var rs = security.GenerateRandomSalt()
	salt := entity.Salt{
		SaltData:  rs,
		CreatedAt: time.Now(),
		UserId:    user_id,
	}

	_, err := sr.Database.Collection(saltCollectionName).InsertOne(context.Background(), salt)
	if err != nil {
		return fmt.Errorf("SaltRepo - Create - rows.Scan: %w", err)
	}
	return nil
}

// Update -.
func (sr *SaltRepo) Update(ctx context.Context, user_id string) error {

	var rs = security.GenerateRandomSalt()
	salt := entity.Salt{
		SaltData:  rs,
		CreatedAt: time.Now(),
	}

	_, err := sr.Database.Collection(saltCollectionName).UpdateOne(
		context.Background(),
		bson.M{"user_id": user_id},
		salt)

	if err != nil {
		return fmt.Errorf("SaltRepo - Create - rows.Scan: %w", err)
	}
	return nil
}

// Delete -.
func (sr *SaltRepo) Delete(ctx context.Context, user_id string) error {
	_, err := sr.Database.Collection(saltCollectionName).DeleteOne(
		context.Background(),
		bson.M{"user_id": user_id})

	if err != nil {
		return fmt.Errorf("SaltRepo - Delete - rows.Scan: %w", err)
	}
	return nil
}