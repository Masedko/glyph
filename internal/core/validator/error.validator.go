package validator

import "fmt"

type EmptyJWTError struct {
}

func (e EmptyJWTError) Error() string {
	return fmt.Sprintf("Empty JWT")
}
