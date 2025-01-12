package main

import (
	"fmt"
	gorm "user-info-service/external/repository/adaptors/postgres"
	repository "user-info-service/external/repository/adaptors/postgres/controller"
	initializer "user-info-service/initializer/gRPC"
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
	clientDB := postgresDB.GetClient()

	gorm.SyncDB(clientDB)

	userRepo := repository.NewUserRepositoryPQ(clientDB)
	userSrv := service.NewUserService(userRepo)

	userGServer := initializer.NewUserServerGRPC("50051", userSrv)
	userGServer.Start()
}
