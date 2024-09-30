package dto

type ReadVariantRequest struct {
	VariantName string `uri:"variantName" binding:"required"`
}
type VariantDTO struct {
	ProductID   uint    `json:"product_id" validate:"required" example:"1"`
	VariantName string  `json:"variant_name" validate:"required" example:"Màu đen"`
	Description string  `json:"description" example:"Màu đen"`
	Price       float64 `json:"price" validate:"required" example:"100000"`
}
