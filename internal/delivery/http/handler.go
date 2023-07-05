package http

import (
	v1 "github.com/begenov/register-service/internal/delivery/http/v1"
	"github.com/begenov/register-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	return nil
}

func (h *Handler) init() {
	handlerV1 := v1.NewHandler(h.service)
	handlerV1.Init()

}
