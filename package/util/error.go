package util



func IsError(err error) bool {
	if err != nil {
		return true
	} else {
		return false
	}
}
