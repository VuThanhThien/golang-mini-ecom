package grpc_handler

import (
	"context"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	InventoryService services.IInventoryService
	pb.UnimplementedOrderGrpcServer
}

func NewServer(InventoryService services.IInventoryService) *Server {
	server := Server{
		InventoryService: InventoryService,
	}
	return &server
}

func (server *Server) UpdateInventory(_ context.Context, input *pb.UpdateInventoryRequest) (*emptypb.Empty, error) {
	order := input.GetOrder()
	items := order.GetItems()
	mapItemsToDTO := func(items []*pb.Item) []dto.Item {
		dtoItems := make([]dto.Item, len(items))
		for i, item := range items {
			dtoItems[i] = dto.Item{
				ID:         uint(item.Id),
				OrderID:    uint(item.OrderId),
				ProductID:  uint(item.ProductId),
				VariantID:  uint(item.VariantId),
				Name:       item.Name,
				Quantity:   int(item.Quantity),
				Price:      float64(item.Price),
				TotalPrice: float64(item.TotalPrice),
			}
		}
		return dtoItems
	}

	dto := &dto.CreateOrderSucceed{
		ID:          uint(order.Id),
		OrderID:     order.OrderId,
		UserID:      uint(order.UserId),
		PaymentID:   uint(order.PaymentId),
		Status:      order.Status,
		TotalAmount: float64(order.TotalAmount),
		PlacedAt:    order.PlacedAt,
		Items:       mapItemsToDTO(items),
	}

	err := server.InventoryService.CreateOrderSucceed(dto)

	return &emptypb.Empty{}, err

}
