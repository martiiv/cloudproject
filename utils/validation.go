package utils

import (
	"errors"
	"time"
)

// Checks of the arrival time input is valid for the webhook registration

func IsValidInput(arrivalTime string) error {
	if len(arrivalTime) != 0 {
		_, err := time.Parse(time.RFC822, arrivalTime)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("Time cannot be null")
	}
}
