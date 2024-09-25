package dto

type ProductImageDTO struct {
	ProductID uint   `json:"product_id" validate:"required"`
	ImageURL  string `json:"image_url" validate:"required"`
	IsPrimary bool   `json:"is_primary" validate:"required"`
}
