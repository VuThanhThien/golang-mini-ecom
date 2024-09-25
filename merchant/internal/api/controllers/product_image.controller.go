package controllers

import (
	"net/http"
	"strconv"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/gin-gonic/gin"
)

type ProductImageController struct {
	productImageService services.IProductImageService
}

func NewProductImageController(productImageService services.IProductImageService) *ProductImageController {
	return &ProductImageController{productImageService: productImageService}
}

// CreateProductImage godoc
// @Summary Create a new product image
// @Description Create a new product image with the given details
// @Tags Product Images
// @Accept json
// @Produce json
// @Param productImage body dto.ProductImageDTO true "Product image details"
// @Success 201 {object} dto.ProductImageDTO
// @Router /product-images [post]
func (pic *ProductImageController) CreateProductImage(ctx *gin.Context) {
	var productImage dto.ProductImageDTO
	if err := ctx.ShouldBindJSON(&productImage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProductImage, err := pic.productImageService.CreateProductImage(&productImage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdProductImage)
}

// GetProductImageByID godoc

// @Summary Get a product image by ID
// @Description Get a product image by its ID
// @Tags Product Images
// @Accept json
// @Produce json
// @Param id path string true "Product image ID"
// @Success 200 {object} dto.ProductImageDTO
// @Router /product-images/{id} [get]

func (pic *ProductImageController) GetProductImageByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product image ID"})
		return
	}

	productImage, err := pic.productImageService.GetProductImageByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, productImage)
}

// UpdateProductImage godoc
// @Summary Update a product image
// @Description Update a product image by its ID
// @Tags Product Images
// @Accept json
// @Produce json
// @Param id path string true "Product image ID"
// @Param productImage body dto.ProductImageDTO true "Product image details"
// @Success 200 {object} dto.ProductImageDTO
// @Router /product-images/{id} [put]
func (pic *ProductImageController) UpdateProductImage(ctx *gin.Context) {
	var productImage dto.ProductImageDTO
	if err := ctx.ShouldBindJSON(&productImage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProductImage, err := pic.productImageService.UpdateProductImage(&productImage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProductImage)
}

// DeleteProductImage godoc
// @Summary Delete a product image
// @Description Delete a product image by its ID
// @Tags Product Images
// @Accept json
// @Produce json
// @Param id path string true "Product image ID"
// @Success 204 {object} nil
// @Router /product-images/{id} [delete]
func (pic *ProductImageController) DeleteProductImage(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product image ID"})
		return
	}

	err = pic.productImageService.DeleteProductImage(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
