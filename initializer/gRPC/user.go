package initializer

import (
	"fmt"
	"log"
	"net"
	pb "user-info-service/external/handler/adaptors/gRPC"
	handler "user-info-service/external/handler/adaptors/gRPC/service"
	"user-info-service/internal/core/logic"

	"google.golang.org/grpc"
)

type UserSeverGRPC struct {
	port        string
	userService logic.UserService
}

func NewUserServerGRPC(port string, userService logic.UserService) *UserSeverGRPC {
	return &UserSeverGRPC{
		port:        port,
		userService: userService,
	}
}

func (s *UserSeverGRPC) Start() {
	userGHandler := &handler.UserHandlerGRPC{UserSrv: s.userService}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userGHandler)

	log.Printf("Server is running on port %s\n", s.port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
