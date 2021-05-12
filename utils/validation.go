package utils

import (
	"fmt"
	"time"
)

// Checks of the arrival time input is valid for the webhook registration
func isValidInput(arrivalTime string) string {
	if len(arrivalTime) != 0 {
		aTime, err := time.Parse(time.RFC822, arrivalTime)
		if err != nil {
			fmt.Println(aTime)
			return err.Error()
		}
		return arrivalTime
	} else {
		fmt.Println("Arrival time cannot be null")
	}
	return arrivalTime
}
