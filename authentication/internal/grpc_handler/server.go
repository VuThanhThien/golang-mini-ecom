package grpc_handler

import (
	"context"

	"github.com/VuThanhThien/golang-gorm-postgres/authentication/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/authentication/pkg/pb"
)

type Server struct {
	UserService services.UserServiceInterface
	pb.UnimplementedUserGrpcServer
}

func NewServer(UserService services.UserServiceInterface) *Server {
	server := Server{
		UserService: UserService,
	}
	return &server
}

func (server *Server) ReadUser(_ context.Context, input *pb.ReadUserRequest) (*pb.User, error) {
	user, err := server.UserService.ReadUser((uint)(input.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Username: user.Name,
		Id:       uint64(user.ID),
		Role:     user.Role,
		Email:    user.Email,
		Photo:    user.Photo,
		Verified: user.Verified,
		Provider: user.Provider,
	}, nil
}
