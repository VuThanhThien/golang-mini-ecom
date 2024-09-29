package dto

import "github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"

// FilterOptions represents the options for filtering products
type FilterOptions struct {
	Name       string  `json:"name"`
	CategoryId string  `json:"category_id"`
	MinPrice   float64 `json:"min_price"`
	MaxPrice   float64 `json:"max_price"`
}

// PaginationResult represents the result of a paginated query
type PaginationResult struct {
	Data       []models.Product
	TotalItems int64
	TotalPages int
	Page       int
	PageSize   int
}
type CreateProductDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  uint    `json:"category_id"`
}

// CreateProductInput represents the data transfer object for creating a product
type CreateProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	SKU         string  `json:"sku" binding:"required"`
	CategoryID  uint    `json:"category_id"`
	MerchantID  uint    `json:"merchant_id"`
}

// UpdateProductDTO represents the data transfer object for updating a product
type UpdateProductDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	SKU         string  `json:"sku" binding:"required"`
	CategoryID  uint    `json:"category_id"`
	MerchantID  uint    `json:"merchant_id"`
}
