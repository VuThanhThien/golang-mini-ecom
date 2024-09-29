package grpc_handler

import (
	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/order/pkg/pb"
)

type Server struct {
	IOrderService services.IOrderService
	pb.UnimplementedInventoryGrpcServer
}

func NewServer(IOrderService services.IOrderService) *Server {
	server := Server{
		IOrderService: IOrderService,
	}
	return &server
}
