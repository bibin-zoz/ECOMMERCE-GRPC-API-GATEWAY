package handler

import (
	interfaces "grpc-api-gateway/pkg/client/interface"
	"grpc-api-gateway/pkg/utils/models"
	"grpc-api-gateway/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type OrderHandler struct {
	GRPC_Client interfaces.OrderClient
}

func NewOrderHandler(client interfaces.OrderClient) *OrderHandler {
	return &OrderHandler{
		GRPC_Client: client,
	}
}

func (or *OrderHandler) OrderItemsFromCart(c *gin.Context) {
	id, _ := c.Get("user_id")
	userID := id.(int)
	var orderFromCart models.OrderFromCart
	if err := c.ShouldBindJSON(&orderFromCart); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orderSuccessResponse, err := or.GRPC_Client.OrderItemsFromCart(orderFromCart, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully created the order", orderSuccessResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

func (or *OrderHandler) GetOrderDetails(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "page number not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("count", "10"))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	id, _ := c.Get("user_id")
	UserID := id.(int)
	OrderDetails, err := or.GRPC_Client.GetOrderDetails(UserID, page, pageSize)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Full Order Details", OrderDetails, nil)
	c.JSON(http.StatusOK, successRes)
}
