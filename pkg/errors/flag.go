package errors

import "fmt"

type FlagError struct {
	Error error
}

func (f FlagError) String() string {
	return fmt.Sprintf("Flag Error: %v", f.Error)
}
