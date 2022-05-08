package v1

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserRoutes_GetAll(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		r    *UserRoutes
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.GetAll(tt.args.c)
		})
	}
}

func TestUserRoutes_GetById(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		r    *UserRoutes
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.GetById(tt.args.c)
		})
	}
}

func TestUserRoutes_Create(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		r    *UserRoutes
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Create(tt.args.c)
		})
	}
}

func TestUserRoutes_GetAllProfilesById(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		r    *UserRoutes
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.GetAllProfilesById(tt.args.c)
		})
	}
}
