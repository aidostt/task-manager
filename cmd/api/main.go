package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aidostt/task-manager/internal/config"
	"github.com/aidostt/task-manager/internal/handler"
	"github.com/aidostt/task-manager/internal/repository"
	"github.com/aidostt/task-manager/internal/service"
	db2 "github.com/aidostt/task-manager/pkg/db"
	"github.com/aidostt/task-manager/pkg/jwt"
	"github.com/aidostt/task-manager/pkg/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db2.Connect(cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connected to database, port:", cfg.HTTP.Port)
	repositories := repository.NewRepoModels(db)

	tokenManager, err := jwt.NewManager(cfg.Jwt.Secret, cfg.Jwt.AccessTokenTTL, cfg.Jwt.RefreshTokenTTL)
	if err != nil {
		log.Fatal(err)
	}
	services := service.NewServiceModels(repositories, tokenManager)
	handlers := handler.NewHandler(services, tokenManager)

	httpServer := server.NewServer(cfg, handlers.InitRoutes())
	go func() {
		if err := httpServer.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	log.Printf("HTTP Server started at %v:%v", cfg.HTTP.Host, cfg.HTTP.Port)

	// Graceful Shutdown of HTTP Server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := httpServer.Stop(ctx); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}

}
