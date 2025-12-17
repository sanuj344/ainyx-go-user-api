package service

import "time"

// CalculateAge calculates age based on DOB
func CalculateAge(dob time.Time) int {
	now := time.Now()

	age := now.Year() - dob.Year()

	// If birthday hasn't occurred yet this year, subtract 1
	if now.Month() < dob.Month() ||
		(now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	return age
}
