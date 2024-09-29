package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/VuThanhThien/golang-gorm-postgres/payment/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUserGateway interface {
	Get(ctx context.Context, userId uint) (*pb.User, error)
}

type UserGateway struct {
	// client pb.UserGrpcClient
	host string
	port string
}

func NewUserGateway(host string, port string) *UserGateway {
	return &UserGateway{host, port}
}

func (g *UserGateway) Get(ctx context.Context, userId uint) (*pb.User, error) {
	address := fmt.Sprintf("%s:%s", g.host, g.port)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewUserGrpcClient(conn)
	resp, err := client.ReadUser(ctx, &pb.ReadUserRequest{Id: (uint64)(userId)})
	if err != nil {
		log.Println("Error getting user:", err)
		return nil, err
	}
	return resp, nil
}
