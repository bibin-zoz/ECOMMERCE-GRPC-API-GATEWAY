package client

import (
	"context"
	"fmt"
	interfaces "grpc-api-gateway/pkg/client/interface"
	"grpc-api-gateway/pkg/config"
	pb "grpc-api-gateway/pkg/pb/order"
	"grpc-api-gateway/pkg/utils/models"

	"google.golang.org/grpc"
)

type orderClient struct {
	Client pb.OrderClient
}

func NewOrderClient(cfg config.Config) interfaces.OrderClient {
	grpcConnection, err := grpc.Dial(cfg.OrderSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewOrderClient(grpcConnection)

	return &orderClient{
		Client: grpcClient,
	}
}
func (c *orderClient) OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (models.OrderSuccessResponse, error) {
	model := &pb.OrderItem{
		AddressID: int64(orderFromCart.AddressID),
		PaymentID: int64(orderFromCart.PaymentID),
	}
	res, err := c.Client.OrderItemsFromCart(context.Background(), &pb.OrderItemsFromCartRequest{
		OrderFromCart: model,
		UserID:        int64(userID),
	})
	if err != nil {
		return models.OrderSuccessResponse{}, err
	}
	if res.Error != "" {
		return models.OrderSuccessResponse{}, err
	}
	return models.OrderSuccessResponse{
		OrderID:        uint(res.OrderID),
		ShipmentStatus: res.Shipmentstatus,
	}, nil
}
func (c *orderClient) GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {
	res, err := c.Client.GetOrderDetails(context.Background(), &pb.GetOrderDetailsRequest{
		UserID: int64(userId),
		Page:   int64(page),
		Count:  int64(count),
	})
	if err != nil {
		return []models.FullOrderDetails{}, err
	}
	if res.Error != "" {
		return []models.FullOrderDetails{}, err
	}
	var result []models.FullOrderDetails

	for _, v := range res.Details {
		orderDetails := models.OrderDetails{
			OrderId:        int(v.Orderdetails.OrderID),
			FinalPrice:     float64(v.Orderdetails.Price),
			ShipmentStatus: v.Orderdetails.Shipmentstatus,
			PaymentStatus:  v.Orderdetails.Paymentstatus,
		}

		var orderProductDetails []models.OrderProductDetails
		for _, product := range v.OrderProductDetails {
			orderProduct := models.OrderProductDetails{
				ProductID:  uint(product.ProductID),
				Quantity:   int(product.Quantity),
				TotalPrice: float64(product.Price),
			}
			orderProductDetails = append(orderProductDetails, orderProduct)
		}
		fullOrderDetails := models.FullOrderDetails{
			OrderDetails:        orderDetails,
			OrderProductDetails: orderProductDetails,
		}
		result = append(result, fullOrderDetails)
	}
	return result, nil
}
