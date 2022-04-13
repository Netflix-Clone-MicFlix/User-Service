package authorization

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Netflix-Clone-MicFlix/User-Service/config"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JwtKeycloak struct {
	cfg *config.Config
	l   logger.Interface
}

func NewJwtKeycloak(cfg *config.Config, l logger.Interface) *JwtKeycloak {
	return &JwtKeycloak{
		cfg: cfg,
		l:   l,
	}
}

func (m *JwtKeycloak) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (m *JwtKeycloak) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := m.ExtractToken(r)
	SecretKey := "-----BEGIN CERTIFICATE-----\n" + m.cfg.AUTH.Secret + "\n-----END CERTIFICATE-----"

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (m *JwtKeycloak) TokenValid(r *http.Request) error {
	token, err := m.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (m *JwtKeycloak) TokenAuthMiddleware(c *gin.Context) {
	err := m.TokenValid(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}
	c.Next()
}

type AccessDetails struct {
	UserId string
	Name   string
	Email  string
}

func (m *JwtKeycloak) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := m.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		name, ok := claims["name"].(string)
		if !ok {
			return nil, err
		}

		email, ok := claims["email"].(string)
		if !ok {
			return nil, err
		}
		userId, ok := claims["sub"].(string)
		if !ok {
			return nil, err
		}
		return &AccessDetails{
			Name:   name,
			Email:  email,
			UserId: userId,
		}, nil
	}
	return nil, err
}
