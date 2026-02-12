package models

type Filter struct {
	Page      int32          `json:"page"`
	PageSize  int32          `json:"page_size"`
	SortOrder string         `json:"sort_order"`
	SortBy    string         `json:"sort_by"`
	Filters   map[string]any `json:"filters"`
}
