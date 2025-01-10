package main

import (
	"fmt"
	gorm "user-info-service/external/repository/adaptors/postgres"
	repository "user-info-service/external/repository/adaptors/postgres/controller"
	"user-info-service/internal/core/service"
	"user-info-service/pkg/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, relying on environment variables")
	}

	postgresDB := db.ConnectToPG()
	client := postgresDB.GetClient()

	gorm.SyncDB(client)

	userRepo := repository.NewUserRepositoryPQ(client)
	userSrv := service.NewUserService(userRepo)
}
