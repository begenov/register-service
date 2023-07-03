package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/begenov/register-service/internal/config"
	grpcv1 "github.com/begenov/register-service/internal/delivery/grpc"
	httpV1 "github.com/begenov/register-service/internal/delivery/http"
	"github.com/begenov/register-service/internal/repository"
	"github.com/begenov/register-service/internal/server"
	"github.com/begenov/register-service/internal/service"
	"github.com/begenov/register-service/pb"
	"github.com/begenov/register-service/pkg/auth"
	"github.com/begenov/register-service/pkg/database"
	"github.com/begenov/register-service/pkg/hash"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(cfg *config.Config) error {
	db, err := database.Open(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		return err
	}

	hash := hash.NewHash(cfg.Server.Cost)

	auth, err := auth.NewManager(cfg.JWT.SigningKey)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo, hash, auth, cfg.JWT.AccessTokenTTL, cfg.JWT.RefreshTokenTTL)

	handler := httpV1.NewHandler(service)

	srv := server.NewServer(&cfg.Server, handler.Init())

	go runGrpcServer(service)

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

func runGrpcServer(service *service.Service) {
	handler := grpcv1.NewHandler(service)

	grpcServer := grpc.NewServer()
	pb.RegisterRegisterServer(grpcServer, handler)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "localhost"+":9090")
	if err != nil {
		log.Fatal("cannot create listener", err)
	}
	log.Println("starting grpc server:" + ":9090")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}

}
