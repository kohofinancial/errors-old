package errors

import (
	"fmt"
)

// ErrorCardExpired is returned when load amount for instant issue card is invalid
type ErrorCardExpired struct {
	BaseError
}

// CardExpired returns a new instance of ErrorCardExpired
func CardExpired(message string) error {
	newBase := NewBase(fmt.Sprintf("card expired :: %s", message))
	return ErrorCardExpired{
		newBase,
	}
}
