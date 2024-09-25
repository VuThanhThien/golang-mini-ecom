package dto

type CreateMerchantDTO struct {
	Name        string `json:"name" binding:"required"`
	MerchantId  string `json:"merchant_id" binding:"required"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id" binding:"required"`
}

type UpdateMerchantDTO struct {
	Name        string `json:"name" binding:"required"`
	MerchantId  string `json:"merchant_id" binding:"required"`
	Description string `json:"description"`
}
