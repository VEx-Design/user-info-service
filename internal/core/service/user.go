package service

import (
	"errors"
	"fmt"
	ports "user-info-service/external/_ports"
	"user-info-service/internal/core/entities"
	"user-info-service/internal/core/logic"
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) logic.UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUser(userID string) (*entities.User, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	user, err := s.userRepo.GetUser(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}
