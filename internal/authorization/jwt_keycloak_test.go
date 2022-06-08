package authorization

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func TestExtractToken(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		m    *JwtKeycloak
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				r: &http.Request{
					Header: http.Header{
						"Authorization": []string{"Bearer token"},
					},
				},
			},
			want: "token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ExtractToken(tt.args.r); got != tt.want {
				t.Errorf("ExtractToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtKeycloak_VerifyToken(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		m       *JwtKeycloak
		args    args
		want    *jwt.Token
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.VerifyToken(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("JwtKeycloak.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JwtKeycloak.VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtKeycloak_TokenValid(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		m       *JwtKeycloak
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.TokenValid(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("JwtKeycloak.TokenValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJwtKeycloak_TokenAuthMiddleware(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		m    *JwtKeycloak
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.TokenAuthMiddleware(tt.args.c)
		})
	}
}

func TestJwtKeycloak_ExtractTokenMetadata(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		m       *JwtKeycloak
		args    args
		want    *AccessDetails
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.ExtractTokenMetadata(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("JwtKeycloak.ExtractTokenMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JwtKeycloak.ExtractTokenMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}
