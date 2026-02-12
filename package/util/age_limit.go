/**
 * @author Yohana Kangwe
 * @email yonakangwe@gmail.com
 * @create date 2024-06-10 13:22:36
 * @modify date 2024-06-10 13:22:36
 * @desc [description]
 */
package util

import (
	"fmt"
	"time"
)

func CheckIfUnder18(dob string) (bool, error) {
	// Parse the date of birth
	layout := "2006-01-02" // Use the layout matching your date format, e.g., "2006-01-02" for "yyyy-mm-dd"
	dateOfBirth, err := time.Parse(layout, dob)
	if err != nil {
		return false, fmt.Errorf("invalid date of birth format: %v", err)
	}

	// Calculate the difference in years
	currentDate := time.Now()
	age := currentDate.Year() - dateOfBirth.Year()

	// Adjust if the birthday hasn't occurred yet this year
	if currentDate.YearDay() < dateOfBirth.YearDay() {
		age--
	}

	// Check if under 18
	return age < 18, nil
}
