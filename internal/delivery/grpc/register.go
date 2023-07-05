package grpc

import (
	"context"
	"log"

	"github.com/begenov/register-service/internal/domain"
	"github.com/begenov/register-service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) SignUp(ctx context.Context, req *pb.RequestRegister) (*pb.Response, error) {
	log.Println("New Request: SignUp")

	if err := h.service.Register.SignUp(ctx, domain.Register{
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
		Role:     req.Role,
		Address:  req.Address,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	response := &pb.Response{
		Message: "User registered:",
	}

	return response, nil
}

func (h *Handler) SignIn(ctx context.Context, req *pb.RequestSignIn) (*pb.ResponseToken, error) {
	token, err := h.service.Register.SignIn(ctx, domain.RequestSignIn{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to sign in: %v", err)
	}

	return &pb.ResponseToken{
		RefreshToken: token.RefreshToken,
		AccessToken:  token.AccessToken,
	}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *pb.RequestToken) (*pb.ResponseToken, error) {
	token, err := h.service.Register.RefreshToken(ctx, req.RefreshToken, req.Role)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to sign in: %v", err)
	}

	return &pb.ResponseToken{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
