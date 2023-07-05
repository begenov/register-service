package grpc

import (
	"github.com/begenov/register-service/internal/service"
	"github.com/begenov/register-service/pb"
)

type Handler struct {
	service *service.Service
	pb.RegisterServer
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
