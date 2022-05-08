package repositories

import (
	"context"
	"reflect"
	"testing"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
)

func TestProfileRepo_GetAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		ur      *ProfileRepo
		args    args
		want    []entity.Profile
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileRepo.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileRepo.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileRepo_GetById(t *testing.T) {
	type args struct {
		ctx        context.Context
		Profile_id string
	}
	tests := []struct {
		name    string
		ur      *ProfileRepo
		args    args
		want    entity.Profile
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetById(tt.args.ctx, tt.args.Profile_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileRepo.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileRepo.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileRepo_GetAllById(t *testing.T) {
	type args struct {
		ctx     context.Context
		user_id string
	}
	tests := []struct {
		name    string
		ur      *ProfileRepo
		args    args
		want    []entity.Profile
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetAllById(tt.args.ctx, tt.args.user_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileRepo.GetAllById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileRepo.GetAllById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileRepo_Create(t *testing.T) {
	type args struct {
		ctx     context.Context
		Profile entity.Profile
	}
	tests := []struct {
		name    string
		ur      *ProfileRepo
		args    args
		want    entity.Profile
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.Create(tt.args.ctx, tt.args.Profile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileRepo_Update(t *testing.T) {
	type args struct {
		ctx        context.Context
		Profile_id string
		Profile    entity.Profile
	}
	tests := []struct {
		name    string
		ur      *ProfileRepo
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ur.Update(tt.args.ctx, tt.args.Profile_id, tt.args.Profile); (err != nil) != tt.wantErr {
				t.Errorf("ProfileRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfileRepo_Delete(t *testing.T) {
	type args struct {
		ctx        context.Context
		Profile_id string
	}
	tests := []struct {
		name    string
		ur      *ProfileRepo
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ur.Delete(tt.args.ctx, tt.args.Profile_id); (err != nil) != tt.wantErr {
				t.Errorf("ProfileRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
