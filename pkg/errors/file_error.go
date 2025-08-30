package errors

import "fmt"

type FileError struct {
	Error error
}

func (f FileError) String() string {
	return fmt.Sprintf("File Error: %v", f.Error)
}


