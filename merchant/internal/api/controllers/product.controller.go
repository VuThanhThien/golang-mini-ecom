package controllers

import (
	"net/http"
	"strconv"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.IProductService
}

func NewProductController(service services.IProductService) ProductController {
	return ProductController{service: service}
}

// GetProductByID godoc
//
// @Summary		GetProductByID
// @Description	GetProductByID
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			id		path		string			false	"id"
func (c *ProductController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := c.service.GetProductByID(uint(idUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// CreateProduct godoc
//
// @Summary		CreateProduct
// @Description	CreateProduct
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			product	body		dto.CreateProductDTO	true	"product"
// @Security		Bearer
// @Router			/products [post]
// @Success		200	{object}		object
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var input *dto.CreateProductDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := c.service.CreateProduct(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
//
// @Summary		DeleteProduct
// @Description	DeleteProduct
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			id		path		string			false	"id"
// @Security		Bearer
// @Router			/products/{id} [delete]
// @Success		200	{object}		object
func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := c.service.DeleteProduct(uint(idUint)); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}

// GetProductByProductID godoc
//
// @Summary		GetProductByProductID
// @Description	GetProductByProductID
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			product_id		path		string			false	"product_id"
// @Security		Bearer
// @Router			/products/product-id/{product_id} [get]
// @Success		200	{object}		object
func (c *ProductController) GetProductByProductID(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	productIDUint, err := strconv.ParseUint(productID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := c.service.GetProductByProductID(uint(productIDUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}
