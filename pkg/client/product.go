package client

import (
	"context"
	"errors"
	"fmt"
	interfaces "grpc-api-gateway/pkg/client/interface"
	"grpc-api-gateway/pkg/config"
	pb "grpc-api-gateway/pkg/pb/product"
	"grpc-api-gateway/pkg/utils/models"

	"google.golang.org/grpc"
)

type productClient struct {
	Client pb.ProductClient
}

func NewProductClient(cfg config.Config) interfaces.ProductClient {

	grpcConnection, err := grpc.Dial(cfg.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewProductClient(grpcConnection)

	return &productClient{
		Client: grpcClient,
	}

}
func (c *productClient) ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {
	res, err := c.Client.ListProducts(context.Background(), &pb.ListProductRequest{
		Page:  int64(page),
		Count: int64(count),
	})
	if err != nil {
		return []models.ProductBrief{}, err
	}
	var result []models.ProductBrief

	for _, v := range res.Details {
		product := models.ProductBrief{
			ID:            uint(v.ID),
			Name:          v.Name,
			Description:   v.Description,
			CategoryID:    int(v.CategoryID),
			Size:          int(v.Size),
			Stock:         int(v.Stock),
			Price:         float64(v.Price),
			ProductStatus: v.ProductStatus,
		}
		result = append(result, product)
	}

	return result, nil
}
func (c *productClient) AddProducts(product models.Product) (models.Products, error) {
	res, err := c.Client.AddProduct(context.Background(), &pb.AddProductRequest{
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  int64(product.CategoryID),
		Size:        int64(product.Size),
		Stock:       int64(product.Stock),
		Price:       float32(product.Price),
	})
	if err != nil {
		return models.Products{}, err
	}
	if res.Error != "" {
		return models.Products{}, errors.New(res.Error)
	}
	return models.Products{
		ID:          uint(res.ID),
		Name:        res.Name,
		Description: res.Description,
		CategoryID:  uint(res.CategoryID),
		Size:        int(res.Size),
		Stock:       int(res.Stock),
		Price:       float64(res.Price),
	}, nil
}
func (c *productClient) DeleteProduct(id int) error {
	res, err := c.Client.DeleteProduct(context.Background(), &pb.DeleteProductRequest{
		ID: int64(id),
	})
	if err != nil {
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}
func (c *productClient) UpdateProducts(pid int, stock int) (models.ProductUpdateReciever, error) {
	res, err := c.Client.UpdateProducts(context.Background(), &pb.UpdateProductRequest{
		ID:    int64(pid),
		Stock: int64(stock),
	})
	if err != nil {
		return models.ProductUpdateReciever{}, err
	}
	if res.Error != "" {
		return models.ProductUpdateReciever{}, err
	}
	return models.ProductUpdateReciever{
		ProductID: int(res.ID),
		Stock:     int(res.Stock),
	}, nil
}
