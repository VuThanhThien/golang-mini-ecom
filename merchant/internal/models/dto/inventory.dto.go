package dto

type InventoryDTO struct {
	VariantID uint `json:"variant_id" validate:"required" example:"1"`
	Quantity  int  `json:"quantity" validate:"required" example:"10"`
}
