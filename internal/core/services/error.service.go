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

type FileAlreadyExistsError struct {
	filename string
}

func (e FileAlreadyExistsError) Error() string {
	return fmt.Sprintf("File %s already exists", e.filename)
}

type FileCreationError struct {
	filename string
	error
}

func (e FileCreationError) Error() string {
	return fmt.Sprintf("Cannot create file %s: %s", e.filename, e.error)
}

type FolderCreationError struct {
	foldername string
	error
}

func (e FolderCreationError) Error() string {
	return fmt.Sprintf("Cannot create folder %s: %s", e.foldername, e.error)
}

type CopyError struct {
	error
}

func (e CopyError) Error() string {
	return fmt.Sprintf("Cannot copy decompressed content into file: %s", e.error)
}

type GETError struct {
	url string
	error
}

func (e GETError) Error() string {
	return fmt.Sprintf("Cannot GET from %s: %s", e.url, e.error)
}

type ReadResponseBodyError struct {
	error
}

func (e ReadResponseBodyError) Error() string {
	return fmt.Sprintf("Cannot read response body: %s", e.error)
}

type HTTPError struct {
	url        string
	statusCode int
	response   string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("HTTP Error with %s with status code %d: %s", e.url, e.statusCode, e.response)
}

type ParserCreationError struct {
	error
}

func (e ParserCreationError) Error() string {
	return fmt.Sprintf("Cannot create Stream Parser: %s", e.error)
}

type ParserError struct {
	error
}

func (e ParserError) Error() string {
	return fmt.Sprintf("Parser error: %s", e.error)
}

type OpenFileError struct {
	filename string
	error
}

func (e OpenFileError) Error() string {
	return fmt.Sprintf("Cannot open file %s: %s", e.filename, e.error)
}

type RemoveFileError struct {
	filename string
	error
}

func (e RemoveFileError) Error() string {
	return fmt.Sprintf("Cannot remove file %s: %s", e.filename, e.error)
}

type CloseFileError struct {
	filename string
	error
}

func (e CloseFileError) Error() string {
	return fmt.Sprintf("Cannot close file %s: %s", e.filename, e.error)
}
