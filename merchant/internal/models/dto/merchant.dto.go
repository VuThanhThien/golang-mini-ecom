package dto

type ReadMerchantRequest struct {
	MerchantID string `uri:"merchantID" binding:"required"`
}
type CreateMerchantInput struct {
	Name         string `json:"name" binding:"required" example:"Merchant Name"`
	MerchantCode string `json:"merchant_code" binding:"required" example:"MER001"`
	Description  string `json:"description" example:"Merchant Description"`
}
type CreateMerchantDTO struct {
	Name         string `json:"name" binding:"required" example:"Merchant Name"`
	MerchantCode string `json:"merchant_code" binding:"required" example:"MER001"`
	Description  string `json:"description" example:"Merchant Description"`
	UserID       uint   `json:"user_id" binding:"required" example:"1"`
}

type UpdateMerchantDTO struct {
	Name         string `json:"name" binding:"required" example:"Merchant Name"`
	MerchantCode string `json:"merchant_code" binding:"required" example:"MER001"`
	Description  string `json:"description" example:"Merchant Description"`
}
