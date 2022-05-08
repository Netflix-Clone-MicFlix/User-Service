package repositories

import (
	"context"
	"reflect"
	"testing"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
)

func TestMovieTagRepo_GetAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		ur      *MovieTagRepo
		args    args
		want    []entity.MovieTag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieTagRepo.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieTagRepo.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieTagRepo_GetById(t *testing.T) {
	type args struct {
		ctx         context.Context
		MovieTag_id string
	}
	tests := []struct {
		name    string
		ur      *MovieTagRepo
		args    args
		want    entity.MovieTag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetById(tt.args.ctx, tt.args.MovieTag_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieTagRepo.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieTagRepo.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieTagRepo_Create(t *testing.T) {
	type args struct {
		ctx      context.Context
		Genre_id string
	}
	tests := []struct {
		name    string
		ur      *MovieTagRepo
		args    args
		want    entity.MovieTag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.Create(tt.args.ctx, tt.args.Genre_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieTagRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieTagRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieTagRepo_Update(t *testing.T) {
	type args struct {
		ctx         context.Context
		MovieTag_id string
		MovieTag    entity.MovieTag
	}
	tests := []struct {
		name    string
		ur      *MovieTagRepo
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ur.Update(tt.args.ctx, tt.args.MovieTag_id, tt.args.MovieTag); (err != nil) != tt.wantErr {
				t.Errorf("MovieTagRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMovieTagRepo_Delete(t *testing.T) {
	type args struct {
		ctx         context.Context
		MovieTag_id string
	}
	tests := []struct {
		name    string
		ur      *MovieTagRepo
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ur.Delete(tt.args.ctx, tt.args.MovieTag_id); (err != nil) != tt.wantErr {
				t.Errorf("MovieTagRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
