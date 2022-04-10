package repositories

import (
	"context"

	"fmt"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const MovieTagCollectionName = "MovieTags"

// MovieTagRepo -.
type MovieTagRepo struct {
	*mongodb.MongoDB
}

// New -.
func NewMovieTagRepo(mdb *mongodb.MongoDB) *MovieTagRepo {
	return &MovieTagRepo{mdb}
}

// GetAll -.
func (ur *MovieTagRepo) GetAll(ctx context.Context) ([]entity.MovieTag, error) {

	MovieTags := []entity.MovieTag{}

	collection := ur.Database.Collection(MovieTagCollectionName)

	var filter bson.M = bson.M{}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("MovieTagRepo - GetAll - rows.Scan: %w", err)
	}
	if err = curr.All(context.Background(), &MovieTags); err != nil {
		return nil, fmt.Errorf("MovieTagRepo - GetAll - rows.Scan: %w", err)
	}

	return MovieTags, nil
}

// GetById -.
func (ur *MovieTagRepo) GetById(ctx context.Context, MovieTag_id string) (entity.MovieTag, error) {
	MovieTag := entity.MovieTag{}

	var filter bson.M = bson.M{"id": MovieTag_id}
	collection := ur.Database.Collection(MovieTagCollectionName)

	if err := collection.FindOne(ctx, filter).Decode(&MovieTag); err != nil {
		return entity.MovieTag{}, fmt.Errorf("MovieRepo - GetAll - rows.Scan: %w", err)
	}

	return MovieTag, nil
}

// Create -.
func (ur *MovieTagRepo) Create(ctx context.Context, Genre_id string) (entity.MovieTag, error) {

	guid := uuid.New().String()
	MovieTag := entity.MovieTag{
		Id:         guid,
		GenreId:    Genre_id,
		WatchCount: 0,
	}

	_, err := ur.Database.Collection(MovieTagCollectionName).InsertOne(context.Background(), MovieTag)
	if err != nil {
		return entity.MovieTag{}, fmt.Errorf("MovieTagRepo - Create - rows.Scan: %w", err)
	}
	return MovieTag, nil
}

// Update -.
func (ur *MovieTagRepo) Update(ctx context.Context, MovieTag_id string, MovieTag entity.MovieTag) error {

	MovieTag.Id = MovieTag_id
	update := bson.M{"$set": MovieTag}

	_, err := ur.Database.Collection(MovieTagCollectionName).UpdateOne(
		context.Background(),
		bson.M{"id": MovieTag_id},
		update)

	if err != nil {
		return fmt.Errorf("MovieTagRepo - Update - rows.Scan: %w", err)
	}
	return nil
}

// Delete -.
func (ur *MovieTagRepo) Delete(ctx context.Context, MovieTag_id string) error {
	_, err := ur.Database.Collection(MovieTagCollectionName).DeleteOne(
		context.Background(),
		bson.M{"id": MovieTag_id})

	if err != nil {
		return fmt.Errorf("MovieTagRepo - Delete - rows.Scan: %w", err)
	}
	return nil
}
