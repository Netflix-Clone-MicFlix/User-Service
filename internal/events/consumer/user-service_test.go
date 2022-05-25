package consumer

import (
	"testing"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
	"github.com/Netflix-Clone-MicFlix/User-Service/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"
)

func TestNewUserServiceEvents(t *testing.T) {
	// id := "c00f99ba-a997-4311-ba81-c6aa78f94b13"
	// usermock := entity.User{id, "test", []string{"test", "test"}}
	// ctx := context.Background()

	UserConsumer := new(mocks.UserConsumer)

	UserConsumer.On("NewUserServiceEvents", mock.AnythingOfType("*amqp.Channel"), mock.AnythingOfType("internal.User")).Return(true, nil)

	result, _ := UserConsumer.NewUserServiceEvents(amqp.Channel{}, nil)

	assert.Equal(t, result, "true")
}

func Test_handleUserServiceEvents(t *testing.T) {
	type args struct {
		messages <-chan amqp.Delivery
		user     internal.User
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleUserServiceEvents(tt.args.messages, tt.args.user)
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		resourcePath string
		user         internal.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.resourcePath, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		resourcePath string
		user         internal.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUser(tt.args.resourcePath, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
