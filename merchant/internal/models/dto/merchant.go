package dto

type CreateMerchantDTO struct {
	Name        string `json:"name" binding:"required"`
	MerchantId  string `json:"merchant_id" binding:"required"`
	Description string `json:"description"`
}
