package utils

import (
	"fmt"
	"time"
)

// Checks of the arrival time input is valid for the webhook registration
func isValidInput(arrivalTime string) bool {
	if len(arrivalTime) != 0 {
		aTime, err := time.Parse(time.RFC822, arrivalTime)
		if err != nil {
			fmt.Println(aTime)
			return false
		}
		return true
	} else {
		fmt.Println("Arrival time cannot be null")
		return false
	}
}
