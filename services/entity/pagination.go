package entity

import "math"

const (
	MinInt32 = 1
	MaxInt32 = math.MinInt32
)

type Meta struct {
	Page     int32
	PageSize int32
	Pages    int32
	Total    int32
}
