package dto

type ReadVariantRequest struct {
	VariantName string `uri:"variantName" binding:"required"`
}
type VariantDTO struct {
	ProductID   uint    `json:"product_id" validate:"required"`
	VariantName string  `json:"variant_name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required"`
}
