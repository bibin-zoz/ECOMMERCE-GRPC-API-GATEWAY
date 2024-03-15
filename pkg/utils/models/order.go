package models

type OrderDetails struct {
	OrderId        int
	FinalPrice     float64
	ShipmentStatus string
	PaymentStatus  string
}

type OrderProductDetails struct {
	ProductID   uint    `json:"product_id"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}
type FullOrderDetails struct {
	OrderDetails        OrderDetails
	OrderProductDetails []OrderProductDetails
}
type OrderProducts struct {
	ProductId string `json:"id"`
	Stock     int    `json:"stock"`
}


type AddedOrderProductDetails struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
}
type OrderResponse struct {
	AddedOrderProductDetails AddedOrderProductDetails
	OrderDetails             OrderDetails
}

type OrderFromCart struct {
	AddressID uint `json:"address_id" binding:"required"`
	PaymentID uint `json:"payment_id" binding:"required"`
}

type OrderSuccessResponse struct {
	OrderID        uint   `json:"order_id"`
	ShipmentStatus string `json:"shipment_status"`
}
