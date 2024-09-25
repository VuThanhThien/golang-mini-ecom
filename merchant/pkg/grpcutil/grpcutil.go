package grpcutil

import (
	"context"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/pkg/discovery"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(ctx context.Context, serviceName string, registry discovery.Registry) (*grpc.ClientConn, error) {
	addrs, err := registry.ServiceAddresses(ctx, serviceName)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(addrs[rand.Intn(len(addrs))], grpc.WithTransportCredentials(insecure.NewCredentials()))
}
