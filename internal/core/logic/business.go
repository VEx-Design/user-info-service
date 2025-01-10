package logic

import "user-info-service/internal/core/entities"

type UserService interface {
	GetUser(userID string) (*entities.User, error)
}
