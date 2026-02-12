package pagination

import (
	"library/package/models"
	"math"
)

const (
	MinInt32 = 1
	MaxInt32 = math.MinInt32
)

func GetMetaData(page, pageSize, total int32) *models.Meta {

	var pages int32
	if total == 0 {
		pages = 0
	} else {
		pages = total / pageSize
		if total%pageSize != 0 {
			pages += 1 // add 1 if there is a remainder
		}
	}

	meta := &models.Meta{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Pages:    pages,
	}
	return meta
}
