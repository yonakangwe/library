package models

type Meta struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Pages    int32 `json:"pages"`
	Total    int32 `json:"total"`
}
