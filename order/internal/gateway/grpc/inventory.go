package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/VuThanhThien/golang-gorm-postgres/order/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/order/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IInventoryGateway interface {
	UpdateInventory(ctx context.Context, order *models.Order) (*emptypb.Empty, error)
}

type InventoryGateway struct {
	host string
	port string
}

func NewInventoryGateway(host string, port string) *InventoryGateway {
	return &InventoryGateway{host, port}
}

func (g *InventoryGateway) UpdateInventory(ctx context.Context, order *models.Order) (*emptypb.Empty, error) {
	address := fmt.Sprintf("%s:%s", g.host, g.port)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewOrderGrpcClient(conn)
	items := make([]*pb.Item, len(order.Items))
	for i, item := range order.Items {
		items[i] = &pb.Item{
			Id:         uint64(item.ID),
			Name:       item.Name,
			OrderId:    uint64(item.OrderID),
			ProductId:  uint64(item.ProductID),
			VariantId:  uint64(item.VariantID),
			Quantity:   uint64(item.Quantity),
			Price:      float32(item.Price),
			TotalPrice: float32(item.TotalPrice),
		}
	}
	result, err := client.UpdateInventory(ctx, &pb.UpdateInventoryRequest{Order: &pb.Order{
		Id:          uint64(order.ID),
		OrderId:     order.OrderID,
		UserId:      uint64(order.UserID),
		PaymentId:   uint64(order.PaymentID),
		Status:      order.Status,
		TotalAmount: float32(order.TotalAmount),
		Items:       items,
		PlacedAt:    order.PlacedAt.Format("2006-01-02 15:04:05"),
	}})
	if err != nil {
		log.Println("Error updating inventory:", err)
		return nil, err
	}
	return result, nil
}
