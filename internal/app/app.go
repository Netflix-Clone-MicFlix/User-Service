// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Netflix-Clone-MicFlix/User-Service/config"
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

	// Use case
	userUseCase := services.NewUserUseCase(
		repositories.NewUserRepo(mdb),
		repositories.NewProfileRepo(mdb),
		repositories.NewMovieTagRepo(mdb),
		nil,
	)

	// RabbitMQ RPC Server
	// rmqRouter := amqprpc.NewRouter(userUseCase)

	// rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, l)
	// if err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	// }

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

	v1.NewRouter(handler, l, userUseCase, corsConfig)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		// case err = <-rmqServer.Notify():
		// 	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	// err = rmqServer.Shutdown()
	// if err != nil {
	// 	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	// }
}
