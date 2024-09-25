package controllers

import (
	"net/http"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/middleware"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/pkg/pb"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MerchantController struct {
	DB      *gorm.DB
	service services.IMerchantService
}

func NewMerchantController(DB *gorm.DB, service services.IMerchantService) MerchantController {
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
	value, oke := c.Get(middleware.CURRENT_USER)
	if !oke {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in"})
		return
	}
	user := value.(*pb.User)
	var dto dto.CreateMerchantDTO
	dto.UserId = uint(user.Id)
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	merchant, err := uc.service.CreateMerchantWithTx(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, merchant)
}
