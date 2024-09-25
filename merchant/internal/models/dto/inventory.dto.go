package dto

type InventoryDTO struct {
	VariantID uint `json:"variant_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required"`
}
