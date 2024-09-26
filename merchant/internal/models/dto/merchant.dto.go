package dto

type ReadMerchantRequest struct {
	MerchantID string `uri:"merchantID" binding:"required"`
}
type CreateMerchantInput struct {
	Name        string `json:"name" binding:"required"`
	MerchantID  string `json:"merchant_id" binding:"required"`
	Description string `json:"description"`
}
type CreateMerchantDTO struct {
	Name        string `json:"name" binding:"required"`
	MerchantID  string `json:"merchant_id" binding:"required"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id" binding:"required"`
}

type UpdateMerchantDTO struct {
	Name        string `json:"name" binding:"required"`
	MerchantID  string `json:"merchant_id" binding:"required"`
	Description string `json:"description"`
}
