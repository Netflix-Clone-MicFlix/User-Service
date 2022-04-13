// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"net/http"

	external "github.com/Netflix-Clone-MicFlix/User-Service/internal/events/consumer"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"

	"github.com/Netflix-Clone-MicFlix/User-Service/config"
	auth "github.com/Netflix-Clone-MicFlix/User-Service/internal/authorization"
	v1 "github.com/Netflix-Clone-MicFlix/User-Service/internal/controller/http/v1"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal/repositories"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal/services"

	// "github.com/Netflix-Clone-MicFlix/User-Service/internal/webapi"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/httpserver"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/logger"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/mongodb"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository mongodb
	mdb, err := mongodb.New(cfg.MDB.Username, cfg.MDB.Password, cfg.MDB.Cluster, cfg.MDB.Database)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Usecase
	userUseCase := services.NewUserUseCase(
		repositories.NewUserRepo(mdb),
		repositories.NewProfileRepo(mdb),
		repositories.NewMovieTagRepo(mdb),
		nil,
	)

	//keycloak
	jwtKeycloak := auth.NewJwtKeycloak(cfg, l)

	// RabbitMQ RPC Server
	connectRabbitMQ, err := amqp.Dial(cfg.RMQ.URL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Events
	go external.NewUserServiceEvents(channelRabbitMQ, userUseCase)

	// HTTP Server
	handler := gin.New()

	corsConfig := cors.New(cors.Config{
		AllowOrigins:     cfg.HTTP.AllowedOrigins,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
	handler.Use(corsConfig)

	//creating new controller router
	v1.NewRouter(handler, jwtKeycloak, l, userUseCase, corsConfig)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
