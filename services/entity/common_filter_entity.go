package entity

type Filter struct {
	Page      int32
	PageSize  int32
	SortBy    string
	SortOrder string
	Filters   map[string]any
}
