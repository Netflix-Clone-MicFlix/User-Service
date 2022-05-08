package repositories

import (
	"context"
	"reflect"
	"testing"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
)

func TestUserRepo_GetAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		ur      *UserRepo
		args    args
		want    []entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_GetById(t *testing.T) {
	type args struct {
		ctx     context.Context
		User_id string
	}
	tests := []struct {
		name    string
		ur      *UserRepo
		args    args
		want    entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetById(tt.args.ctx, tt.args.User_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_GetByKeycloakId(t *testing.T) {
	type args struct {
		ctx         context.Context
		keycloak_id string
	}
	tests := []struct {
		name    string
		ur      *UserRepo
		args    args
		want    entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.GetByKeycloakId(tt.args.ctx, tt.args.keycloak_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.GetByKeycloakId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.GetByKeycloakId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_Create(t *testing.T) {
	type args struct {
		ctx         context.Context
		keycloak_id string
	}
	tests := []struct {
		name    string
		ur      *UserRepo
		args    args
		want    entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ur.Create(tt.args.ctx, tt.args.keycloak_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_Update(t *testing.T) {
	type args struct {
		ctx     context.Context
		User_id string
		User    entity.User
	}
	tests := []struct {
		name    string
		ur      *UserRepo
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ur.Update(tt.args.ctx, tt.args.User_id, tt.args.User); (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepo_Delete(t *testing.T) {
	type args struct {
		ctx     context.Context
		user_id string
	}
	tests := []struct {
		name    string
		ur      *UserRepo
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ur.Delete(tt.args.ctx, tt.args.user_id); (err != nil) != tt.wantErr {
				t.Errorf("UserRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
