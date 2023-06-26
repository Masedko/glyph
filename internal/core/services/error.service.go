package services

import (
	"fmt"
)

type ValidateError struct {
	error
}

func (e ValidateError) Error() string {
	return fmt.Sprintf("Validate error: %v", e.error)
}

type RepositoryError struct {
	error
}

func (e RepositoryError) Error() string {
	return fmt.Sprintf("Repository Error: %v", e.error)
}

type MatchAlreadyParsedError struct {
}

func (e MatchAlreadyParsedError) Error() string {
	return fmt.Sprintf("Match already parsed (File parsed when request were done).")
}

type NoGlyphsError struct {
}

func (e NoGlyphsError) Error() string {
	return fmt.Sprintf("No glyphs found")
}
