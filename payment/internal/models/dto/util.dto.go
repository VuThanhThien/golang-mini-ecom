package dto

type ReadIdRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

type PaginationDto struct {
	Page     *int `query:"page"        example:"1"  default:"1"`
	PageSize *int `query:"pageSize"   example:"5"  json:"limit" default:"10"`
}

type MetadataDto struct {
	Total   int64 `json:"total"`
	Page    int32 `json:"page"`
	PerPage int32 `json:"per_page"`
}
