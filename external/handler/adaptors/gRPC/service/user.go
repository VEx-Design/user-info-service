package handler

import (
	"context"
	"errors"
	"fmt"

	pb "user-info-service/external/handler/adaptors/gRPC"
	"user-info-service/internal/core/logic"
)

type UserHandlerGRPC struct {
	pb.UnimplementedUserServiceServer
	UserSrv logic.UserService
}

func (s *UserHandlerGRPC) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if req.GetUserId() == "" {
		return nil, errors.New("user ID is required")
	}

	// Fetch the user from your user service
	user, err := s.UserSrv.GetUser(req.GetUserId())
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Return the user information
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Picture: user.Picture,
		},
	}, nil
}
