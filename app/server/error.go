package server

import (
	"fmt"
)

type Err struct {
	Code    int
	Message string
	Cause   error
}

func (e Err) Error() string {
	return fmt.Sprintf("(%d): %s. Cause: %v", e.Code, e.Message, e.Cause)
}
