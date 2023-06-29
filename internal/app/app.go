package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/begenov/register-service/internal/config"
	httpV1 "github.com/begenov/register-service/internal/delivery/http"
	"github.com/begenov/register-service/internal/repository"
	"github.com/begenov/register-service/internal/server"
	"github.com/begenov/register-service/internal/service"
	"github.com/begenov/register-service/pkg/database"
)

func Run(cfg *config.Config) error {
	db, err := database.Open(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	handler := httpV1.NewHandler(service)

	srv := server.NewServer(&cfg.Server, handler.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Println("Server started", cfg.Server.Port)

	quit := make(chan os.Signal, 1)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		return fmt.Errorf("failed to stop server: %v", err)
	}
	return nil
}
