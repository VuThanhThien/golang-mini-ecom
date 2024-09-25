package controllers

import (
	"net/http"
	"strconv"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/gin-gonic/gin"
)

type VariantController struct {
	variantService services.IVariantService
}

func NewVariantController(variantService services.IVariantService) *VariantController {
	return &VariantController{variantService: variantService}
}

// CreateVariant godoc
// @Summary Create a new variant
// @Description Create a new variant with the given details
// @Tags Variants
// @Accept json
// @Produce json
// @Param variant body dto.VariantDTO true "Variant details"
// @Success 201 {object} dto.VariantDTO
// @Router /variants [post]
func (vc *VariantController) CreateVariant(ctx *gin.Context) {
	var variant dto.VariantDTO
	if err := ctx.ShouldBindJSON(&variant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdVariant, err := vc.variantService.Create(variant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdVariant)
}

// GetVariantByProductID godoc
// @Summary Get a variant by product ID
// @Description Get a variant by its product ID
// @Tags Variants
// @Accept json
// @Produce json
// @Param productID path string true "Product ID"
// @Success 200 {object} dto.VariantDTO
// @Router /variants/{productID} [get]
func (vc *VariantController) GetVariantByProductID(ctx *gin.Context) {
	productID, err := strconv.ParseUint(ctx.Param("productID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	variant, err := vc.variantService.GetByProductID(uint(productID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, variant)
}

// GetVariantByVariantName godoc
// @Summary Get a variant by variant name
// @Description Get a variant by its variant name
// @Tags Variants
// @Accept json
// @Produce json
// @Param variantName path string true "Variant Name"
// @Success 200 {object} dto.VariantDTO
// @Router /variants/{variantName} [get]
func (vc *VariantController) GetVariantByVariantName(ctx *gin.Context) {
	variantName := ctx.Param("variantName")
	variant, err := vc.variantService.GetByVariantName(variantName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, variant)
}

// DeleteVariant godoc
// @Summary Delete a variant by ID
// @Description Delete a variant by its ID
// @Tags Variants
// @Accept json
// @Produce json
// @Param id path string true "Variant ID"
// @Success 204 {object} nil
// @Router /variants/{id} [delete]
func (vc *VariantController) DeleteVariant(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid variant ID"})
		return
	}

	err = vc.variantService.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
