package repositories

import (
	"context"

	"fmt"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const ProfileCollectionName = "Profiles"

// ProfileRepo -.
type ProfileRepo struct {
	*mongodb.MongoDB
}

// New -.
func NewProfileRepo(mdb *mongodb.MongoDB) *ProfileRepo {
	return &ProfileRepo{mdb}
}

// GetAll -.
func (ur *ProfileRepo) GetAll(ctx context.Context) ([]entity.Profile, error) {

	Profiles := []entity.Profile{}

	collection := ur.Database.Collection(ProfileCollectionName)

	var filter bson.M = bson.M{}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("ProfileRepo - GetAll - rows.Scan: %w", err)
	}
	if err = curr.All(context.Background(), &Profiles); err != nil {
		return nil, fmt.Errorf("ProfileRepo - GetAll - rows.Scan: %w", err)
	}

	return Profiles, nil
}

// GetById -.
func (ur *ProfileRepo) GetById(ctx context.Context, Profile_id string) (entity.Profile, error) {
	Profile := entity.Profile{}

	var filter bson.M = bson.M{"id": Profile_id}
	collection := ur.Database.Collection(ProfileCollectionName)

	if err := collection.FindOne(ctx, filter).Decode(&Profile); err != nil {
		return entity.Profile{}, fmt.Errorf("MovieRepo - GetAll - rows.Scan: %w", err)
	}

	return Profile, nil
}

// GetById -.
func (ur *ProfileRepo) GetAllById(ctx context.Context, user_id string) ([]entity.Profile, error) {
	Profiles := []entity.Profile{}

	collection := ur.Database.Collection(ProfileCollectionName)

	var filter bson.M = bson.M{"userid": user_id}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("ProfileRepo - GetAll - rows.Scan: %w", err)
	}
	if err = curr.All(context.Background(), &Profiles); err != nil {
		return nil, fmt.Errorf("ProfileRepo - GetAll - rows.Scan: %w", err)
	}

	return Profiles, nil
}

// Create -.
func (ur *ProfileRepo) Create(ctx context.Context, Profile entity.Profile) (entity.Profile, error) {

	guid := uuid.New().String()
	Profile.Id = guid

	_, err := ur.Database.Collection(ProfileCollectionName).InsertOne(context.Background(), Profile)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("ProfileRepo - Create - rows.Scan: %w", err)
	}
	return Profile, nil
}

// Update -.
func (ur *ProfileRepo) Update(ctx context.Context, Profile_id string, Profile entity.Profile) error {

	Profile.Id = Profile_id
	update := bson.M{"$set": Profile}

	_, err := ur.Database.Collection(ProfileCollectionName).UpdateOne(
		context.Background(),
		bson.M{"id": Profile_id},
		update)

	if err != nil {
		return fmt.Errorf("ProfileRepo - Update - rows.Scan: %w", err)
	}
	return nil
}

// Delete -.
func (ur *ProfileRepo) Delete(ctx context.Context, Profile_id string) error {
	_, err := ur.Database.Collection(ProfileCollectionName).DeleteOne(
		context.Background(),
		bson.M{"id": Profile_id})

	if err != nil {
		return fmt.Errorf("ProfileRepo - Delete - rows.Scan: %w", err)
	}
	return nil
}
