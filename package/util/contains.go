package util

func CheckContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func AnyValueInArray(a []string, b []string) (bool, string) {
	for _, value := range a {
		if CheckContains(b, value) {
			return true, value
		}
	}
	return false, ""
}
