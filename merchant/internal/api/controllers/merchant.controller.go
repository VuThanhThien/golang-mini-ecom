package controllers

import (
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MerchantController struct {
	DB      *gorm.DB
	service services.MerchantServiceInterface
}

func NewMerchantController(DB *gorm.DB, service services.MerchantServiceInterface) MerchantController {
	return MerchantController{DB, service}
}

//		GetMerchant godoc
//
//		@Summary		GetMerchant
//		@Description	GetMerchant
//		@Tags			merchants
//		@Accept			json
//		@Produce		json
//		@Param			id		path		string			false	"id"
//		@Success		200	{object}		object
//	 	@Security		Bearer
//		@Router			/merchants/{id} [get]
func (uc *MerchantController) GetMerchant(c *gin.Context) {
	id := c.Param("id")
	merchant, err := uc.service.GetMerchantByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

//		GetMerchantByMerchantID godoc
//
//		@Summary		GetMerchantByMerchantID
//		@Description	GetMerchantByMerchantID
//		@Tags			merchants
//		@Accept			json
//		@Produce		json
//		@Param			merchantID	path		string			false	"merchantID"
//		@Success		200	{object}		object
//	 	@Security		Bearer
//		@Router			/merchants/merchant-id/{merchantID} [get]
func (uc *MerchantController) GetMerchantByMerchantID(c *gin.Context) {
	merchantID := c.Param("merchantID")

	merchant, err := uc.service.GetMerchantByMerchantID(merchantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, merchant)
}

//		CreateMerchant godoc
//
//		@Summary		CreateMerchant
//		@Description	CreateMerchant
//		@Tags			merchants
//		@Accept			json
//		@Produce		json
//		@Param			merchant	body		dto.CreateMerchantDTO	true	"merchant"
//	 	@Security		Bearer
//		@Router			/merchants [post]
//		@Success		200	{object}		object
func (uc *MerchantController) CreateMerchant(c *gin.Context) {
	var dto dto.CreateMerchantDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchant, err := uc.service.CreateMerchantWithTx(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, merchant)
}
