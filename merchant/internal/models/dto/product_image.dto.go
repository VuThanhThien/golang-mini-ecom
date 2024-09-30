package dto

type ProductImageDTO struct {
	ProductID uint   `json:"product_id" validate:"required" example:"1"`
	ImageURL  string `json:"image_url" validate:"required" example:"https://example.com/image.jpg"`
	IsPrimary bool   `json:"is_primary" validate:"required" example:"false"`
}
