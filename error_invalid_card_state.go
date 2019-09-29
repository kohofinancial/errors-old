package errors

import (
	"fmt"
)

// ErrorInvalidCardState is returned when instant issue card is alreay activated/deactivated
type ErrorInvalidCardState struct {
	BaseError
}

// DataLoad returns a new instance of ErrorDataLoad
func InvalidCardState(message string) error {
	newBase := NewBase(fmt.Sprintf("failed to update the state of the instant issue card :: %s", message))
	return ErrorInvalidCardState{
		newBase,
	}
}
