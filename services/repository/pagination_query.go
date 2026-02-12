

package repository

import (
	"strconv"
)

func GetPaginationQuery(page, pageSize int32, index int, values []any) (string, []any) {
	query := ""

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	if index <= 0 {
		index = 1
	}

	offset := (page - 1) * pageSize
	query = query + " LIMIT $" + strconv.Itoa(index)
	values = append(values, pageSize)
	index++

	query = query + " OFFSET $" + strconv.Itoa(index)
	values = append(values, offset)
	return query, values
}
