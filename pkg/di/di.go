package di

import (
	server "grpc-api-gateway/pkg/api"
	"grpc-api-gateway/pkg/api/handler"
	"grpc-api-gateway/pkg/client"
	"grpc-api-gateway/pkg/config"
)


func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	productClient := client.NewProductClient(cfg)
	productHandler := handler.NewProductHandler(productClient)

	userClient := client.NewUserClient(cfg)
	userHandler := handler.NewUserHandler(userClient)

	cartClient := client.NewCartClient(cfg)
	cartHandler := handler.NewCartHandler(cartClient)

	orderClient := client.NewOrderClient(cfg)
	orderHandler := handler.NewOrderHandler(orderClient)

	serverHTTP := server.NewServerHTTP(adminHandler, productHandler, userHandler, cartHandler, orderHandler)

	return serverHTTP, nil
}
