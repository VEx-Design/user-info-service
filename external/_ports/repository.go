package ports

import (
	"user-info-service/internal/core/entities"
)

type UserRepository interface {
	GetUser(userID string) (*entities.User, error)
}
