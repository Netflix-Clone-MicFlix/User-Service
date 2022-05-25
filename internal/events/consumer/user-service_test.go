package consumer

import (
	"testing"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
	"github.com/streadway/amqp"
)

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
