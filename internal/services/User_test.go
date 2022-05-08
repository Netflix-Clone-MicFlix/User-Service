package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
)

func TestUserUseCase_GetById(t *testing.T) {
	type args struct {
		ctx     context.Context
		User_id string
	}
	tests := []struct {
		name    string
		uc      *UserUseCase
		args    args
		want    entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetById(tt.args.ctx, tt.args.User_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUseCase.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUseCase.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUseCase_GetAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		uc      *UserUseCase
		args    args
		want    []entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUseCase.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUseCase.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUseCase_Create(t *testing.T) {
	type args struct {
		ctx         context.Context
		keycloak_id string
	}
	tests := []struct {
		name    string
		uc      *UserUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uc.Create(tt.args.ctx, tt.args.keycloak_id); (err != nil) != tt.wantErr {
				t.Errorf("UserUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserUseCase_GetAllProfilesById(t *testing.T) {
	type args struct {
		ctx     context.Context
		User_id string
	}
	tests := []struct {
		name    string
		uc      *UserUseCase
		args    args
		want    []entity.Profile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetAllProfilesById(tt.args.ctx, tt.args.User_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserUseCase.GetAllProfilesById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUseCase.GetAllProfilesById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUseCase_Delete(t *testing.T) {
	type args struct {
		ctx         context.Context
		keycloak_id string
	}
	tests := []struct {
		name    string
		uc      *UserUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uc.Delete(tt.args.ctx, tt.args.keycloak_id); (err != nil) != tt.wantErr {
				t.Errorf("UserUseCase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
