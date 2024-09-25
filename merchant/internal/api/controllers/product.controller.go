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

// FilterProductsWithPagination godoc
// @Summary		FilterProductsWithPagination
// @Description	FilterProductsWithPagination
// @Tags			Products
// @Accept			json
// @Produce		json
// @Param payload query dto.FilterOptions false "FilterOptions payload"
// @Param _ query dto.PaginationDto false "PaginationDto"
// @Security		Bearer
// @Router			/products [get]
// @Success		200	{object}		object
func (c *ProductController) FilterProductsWithPagination(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	var filter dto.FilterOptions

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	if ctx.Query("name") != "" {
		filter.Name = ctx.Query("name")
	}
	if ctx.Query("category_id") != "" {
		filter.CategoryId = ctx.Query("category_id")
	}
	if ctx.Query("min_price") != "" {
		filter.MinPrice, _ = strconv.ParseFloat(ctx.Query("min_price"), 64)
	}
	if ctx.Query("max_price") != "" {
		filter.MaxPrice, _ = strconv.ParseFloat(ctx.Query("max_price"), 64)
	}

	products, err := c.service.FilterProductsWithPagination(filter, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// GetProductByID godoc
//
// @Summary		GetProductByID
// @Description	GetProductByID
// @Tags			Products
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
// @Tags			Products
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
// @Tags			Products
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
// @Tags			Products
// @Accept			json
// @Produce		json
// @Param			id		path		string			false	"id"
// @Security		Bearer
// @Router			/products/{id} [get]
// @Success		200	{object}		object
func (c *ProductController) GetProductByProductID(ctx *gin.Context) {
	var readIdRequest dto.ReadIdRequest
	if err := ctx.ShouldBindUri(&readIdRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := c.service.GetProductByProductID(uint(readIdRequest.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}
