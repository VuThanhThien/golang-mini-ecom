package controllers

import (
	"net/http"
	"strconv"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService services.ICategoryService
}

func NewCategoryController(categoryService services.ICategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the given details
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body dto.CategoryDTO true "Category details"
// @Success 201 {object} dto.CategoryDTO
// @Router /categories [post]
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var categoryDTO dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&categoryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := c.categoryService.CreateCategory(&categoryDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdCategory)
}

// GetCategoryByID godoc
// @Summary Get a category by ID
// @Description Get a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} object
// @Router /categories/{id} [get]
func (c *CategoryController) GetCategoryByID(ctx *gin.Context) {
	var readIdRequest dto.ReadIdRequest
	if err := ctx.ShouldBindUri(&readIdRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := c.categoryService.GetCategoryByID(uint(readIdRequest.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete a category by ID
// @Description Delete a category by its ID
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 204 {object} object
// @Failure 500 {string} string "an error occurred during the modification"
// @Router /categories/{id} [delete]
func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	var readIdRequest dto.ReadIdRequest
	if err := ctx.ShouldBindUri(&readIdRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := c.categoryService.DeleteCategory(uint(readIdRequest.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// ListCategory godoc
// @Summary List categories
// @Description List categories with optional filtering and pagination
// @Tags Categories
// @Accept json
// @Produce json
// @Param payload query dto.ListCategoryDto false "ListCategoryDto payload"
// @Param _ query dto.PaginationDto false "PaginationDto"
// @Success 200 {object} dto.CategoryListResponse
// @Router /categories [get]
func (c *CategoryController) ListCategory(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	var listCategoryDto dto.ListCategoryDto
	var paginationDto dto.PaginationDto

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	paginationDto.Page = &page
	paginationDto.PageSize = &pageSize
	listCategoryDto.Name = ctx.Query("name")
	if ctx.Query("parent_id") != "" {
		parentID, err := strconv.ParseUint(ctx.Query("parent_id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		listCategoryDto.ParentID = uint(parentID)
	}
	if ctx.Query("id") != "" {
		id, err := strconv.ParseUint(ctx.Query("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		listCategoryDto.ID = uint(id)
	}
	categories, _, err := c.categoryService.ListCategory(listCategoryDto, paginationDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	categoryListResponse := dto.ToCategoryListResponse(categories)

	ctx.JSON(http.StatusOK, categoryListResponse)
}
