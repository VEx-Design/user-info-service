package repository

import (
	"errors"
	"fmt"
	"log"
	ports "user-info-service/external/_ports"
	gorm_model "user-info-service/external/repository/adaptors/postgres/model"
	"user-info-service/internal/core/entities"

	"gorm.io/gorm"
)

type userRepositoryPQ struct {
	client *gorm.DB
}

func NewUserRepositoryPQ(client *gorm.DB) ports.UserRepository {
	return &userRepositoryPQ{client: client}
}

func (r *userRepositoryPQ) GetUser(userId string) (*entities.User, error) {
	if userId == "" {
		return nil, errors.New("user ID is required")
	}

	var user gorm_model.User
	if err := r.client.Where("id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %s not found", userId)
		}
		log.Printf("database query error: userId=%s, error=%v", userId, err)
		return nil, fmt.Errorf("could not retrieve user with ID %s: %w", userId, err)
	}

	// Map database models to domain entities.
	result := &entities.User{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		Picture: user.Picture,
	}

	return result, nil
}
