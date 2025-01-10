package main

import (
	"fmt"
	"log"
	"net"
	pb "user-info-service/external/handler/adaptors/gRPC"
	grpc_service "user-info-service/external/handler/adaptors/gRPC/service"
	gorm "user-info-service/external/repository/adaptors/postgres"
	repository "user-info-service/external/repository/adaptors/postgres/controller"
	"user-info-service/internal/core/service"
	"user-info-service/pkg/db"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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

	gRPCservice := &grpc_service.UserServiceGRPC{UserSrv: userSrv}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, gRPCservice)

	log.Println("Server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
