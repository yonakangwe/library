package util

import "strconv"

func String2Int(num string) int32 {
	res, _ := strconv.ParseInt(num, 10, 32)
	return int32(res)
}
