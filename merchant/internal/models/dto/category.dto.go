package dto

import "github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"

type ReadIdRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

type CategoryDTO struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	ParentID    *uint  `json:"parentId,omitempty"`
}

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentID    *uint  `json:"parent_id"`
}

type ListCategoryDto struct {
	Name     string `json:"name"`
	ParentID uint   `json:"parent_id"`
	ID       uint   `json:"id"`
}

type CategoryListResponse struct {
	Categories []CategoryResponse `json:"categories"`
	Total      int                `json:"total"`
}

func ToCategoryListResponse(categories []models.Category) CategoryListResponse {
	summaries := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		summaries[i] = CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ParentID:    &category.ParentID,
		}
	}
	return CategoryListResponse{
		Categories: summaries,
		Total:      len(summaries),
	}
}
