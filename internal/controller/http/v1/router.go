// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	auth "github.com/Netflix-Clone-MicFlix/User-Service/internal/authorization"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/Netflix-Clone-MicFlix/User-Service/docs"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a user service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, k *auth.JwtKeycloak, l logger.Interface, t internal.User, corsConfig gin.HandlerFunc) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newUserRoutes(h, k, t, l, corsConfig)
	}
}
