package errors

import (
	"fmt"
)

// ErrorInvalidAmount is returned when load amount for instant issue card is invalid
type ErrorInvalidAmount struct {
	BaseError
}

// InvalidAmount returns a new instance of ErrorInvalidAmount
func InvalidAmount(message string) error {
	newBase := NewBase(fmt.Sprintf("invalid amount :: %s", message))
	return ErrorInvalidAmount{
		newBase,
	}
}
