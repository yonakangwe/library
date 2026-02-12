package util

import "regexp"

func IsValidPhoneNumber(number string) bool {
	// Define the regular expression pattern for a 10-digit number starting with '0'
	//Example : 071XXXXXXX, 0712001122
	pattern := `^[0]\d{9}$`

	// Create a regular expression object
	regex := regexp.MustCompile(pattern)

	// Check if the number matches the pattern
	return regex.MatchString(number)
}
