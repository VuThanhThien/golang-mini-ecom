package controllers

import (
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/middleware"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	service services.IMerchantService
}

func NewMerchantController(service services.IMerchantService) MerchantController {
	return MerchantController{service: service}
}

//		GetMerchant godoc
//
//		@Summary		GetMerchant
//		@Description	GetMerchant
//		@Tags			Merchants
//		@Accept			json
//		@Produce		json
//		@Param			id		path		string			false	"id"
//		@Success		200	{object}		object
//	 	@Security		Bearer
//		@Router			/merchants/{id} [get]
func (uc *MerchantController) GetMerchant(c *gin.Context) {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var readIdRequest dto.ReadIdRequest
	if err := c.ShouldBindUri(&readIdRequest); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	merchant, err := uc.service.GetMerchantByID(uint(readIdRequest.ID))
	if merchant.UserID != uint(user.Id) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to access this merchant"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if merchant == nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "Merchant not found"})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

//		GetMerchantByMerchantID godoc
//
//		@Summary		GetMerchantByMerchantID
//		@Description	GetMerchantByMerchantID
//		@Tags			Merchants
//		@Accept			json
//		@Produce		json
//		@Param			merchantID	path		string			false	"Merchant ID"
//		@Success		200	{object}		object
//	 	@Security		Bearer
//		@Router			/merchants/merchant-id/{merchantID} [get]
func (uc *MerchantController) GetMerchantByMerchantID(c *gin.Context) {
	var readMerchantRequest dto.ReadMerchantRequest
	if err := c.ShouldBindUri(&readMerchantRequest); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	merchant, err := uc.service.GetMerchantByMerchantID(readMerchantRequest.MerchantID)
	if merchant == nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "Merchant not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, merchant)
}

//		CreateMerchant godoc
//
//		@Summary		CreateMerchant
//		@Description	CreateMerchant
//		@Tags			Merchants
//		@Accept			json
//		@Produce		json
//		@Param			merchant	body		dto.CreateMerchantInput	true	"merchant"
//	 	@Security		Bearer
//		@Router			/merchants [post]
//		@Success		200	{object}		object
func (uc *MerchantController) CreateMerchant(c *gin.Context) {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var createMerchantDTO dto.CreateMerchantDTO
	createMerchantDTO.UserID = uint(user.Id)
	if err := c.ShouldBindJSON(&createMerchantDTO); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	merchant, err := uc.service.CreateMerchantWithTx(&createMerchantDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, merchant)
}
