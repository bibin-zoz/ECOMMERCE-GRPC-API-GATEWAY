package server

import (
	"grpc-api-gateway/pkg/api/handler"
	"grpc-api-gateway/pkg/api/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, userHandler *handler.UserHandler, cartHandler *handler.CartHandler, orderhandler *handler.OrderHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())

	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignUp)

	router.POST("/user/signup", userHandler.UserSignup)
	router.POST("/user/login", userHandler.Userlogin)

	router.GET("/product", productHandler.ShowAllProducts)

	router.Use(middleware.AdminAuthMiddleware())
	{
		router.POST("/product", productHandler.AddProducts)
		router.DELETE("/product", productHandler.DeleteProduct)
		router.PUT("/product", productHandler.UpdateProducts)
	}
	router.Use(middleware.UserAuthMiddleware())
	{
		router.POST("/cart", cartHandler.AddToCart)
		router.GET("/cart", cartHandler.GetCart)

		router.POST("/order", orderhandler.OrderItemsFromCart)
		router.GET("/order", orderhandler.GetOrderDetails)
	}
	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :3000")
	err := s.engine.Run(":3000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
