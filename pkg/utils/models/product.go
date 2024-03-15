package models

type ProductBrief struct {
	ID            uint    `json:"id" gorm:"unique;not null"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	CategoryID    int     `json:"category_id"`
	Size          int     `json:"size"`
	Stock         int     `json:"stock"`
	Price         float64 `json:"price"`
	ProductStatus string  `json:"product_status"`
}

type ProductUpdateReciever struct {
	ProductID int
	Stock     int
}
type Product struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	CategoryID  uint    `json:"category_id" validate:"required"`
	Size        int     `json:"size" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}
type Products struct {
	ID          uint    `json:"id" gorm:"unique;not null"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"category_id"`
	Size        int     `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
type ProductUpdate struct {
	ProductId int `json:"product_id"`
	Stock     int `json:"stock"`
}
