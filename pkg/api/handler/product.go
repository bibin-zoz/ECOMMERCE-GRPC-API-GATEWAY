package handler

import (
	interfaces "grpc-api-gateway/pkg/client/interface"
	"grpc-api-gateway/pkg/utils/models"
	"grpc-api-gateway/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	GRPC_Client interfaces.ProductClient
}

func NewProductHandler(productClient interfaces.ProductClient) *ProductHandler {
	return &ProductHandler{
		GRPC_Client: productClient,
	}
}

func (pt *ProductHandler) AddProducts(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(product)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	if product.Stock < 1 {
		errs := response.ClientResponse(http.StatusBadRequest, "Invalid Stock", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	products, err := pt.GRPC_Client.AddProducts(product)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Could not add the product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully added products", products, nil)
	c.JSON(http.StatusOK, success)

}
func (pt *ProductHandler) ShowAllProducts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	countStr := c.DefaultQuery("count", "10")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	products, err := pt.GRPC_Client.ShowAllProducts(page, count)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't retrieve products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products", products, nil)
	c.JSON(http.StatusOK, success)
}
func (pt *ProductHandler) DeleteProduct(c *gin.Context) {
	Id := c.Query("id")
	id, _ := strconv.Atoi(Id)
	err := pt.GRPC_Client.DeleteProduct(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Could not delete the specified products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully deleted the product", nil, nil)
	c.JSON(http.StatusOK, success)
}

func (pt *ProductHandler) UpdateProducts(c *gin.Context) {
	var p models.ProductUpdate
	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	a, err := pt.GRPC_Client.UpdateProducts(p.ProductId, p.Stock)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not update the product quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated the product quantity", a, nil)
	c.JSON(http.StatusOK, successRes)

}
