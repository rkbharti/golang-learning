package dob

import (
	"errors"
	"time"
)

func AgeCalculator(dob string) (int, error) {

	layout := "02/01/2006"

	birthTime, err := time.Parse(layout, dob)
	if err != nil {
		return 0, errors.New("invalid date format")
	}

	currentTime := time.Now()

	// future date check
	if birthTime.After(currentTime) {
		return 0, errors.New("DOB cannot be in future")
	}

	age := currentTime.Year() - birthTime.Year()

	// adjust if birthday not passed
	if currentTime.Month() < birthTime.Month() ||
		(currentTime.Month() == birthTime.Month() && currentTime.Day() < birthTime.Day()) {
		age--
	}

	return age, nil
}