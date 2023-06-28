package services

import (
	"compress/bzip2"
	"fmt"
	"go-glyph-v2/internal/core/dtos"
	"io"
	"net/http"
	"os"
)

type ValveService struct {
	client http.Client
}

func NewValveService() *ValveService {
	return &ValveService{client: http.Client{}}
}

func (s ValveService) RetrieveFile(match dtos.Match) error {
	url := fmt.Sprintf("http://replay%d.valve.net/570/%d_%d.dem.bz2", match.Cluster, match.ID, match.ReplaySalt)
	// Get file from Valve cluster
	response, err := s.client.Get(url)
	if err != nil {
		return GETError{url: url, error: err}
	}
	// Handle HTTP error from Valve Server
	if response.StatusCode != 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return ReadResponseBodyError{err}
		}
		return HTTPError{url: url, statusCode: response.StatusCode, response: string(body)}
	}
	defer response.Body.Close()
	path := "internal/data/demos"
	// Create demos folder if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return FolderCreationError{foldername: path}
		}
	}
	filename := fmt.Sprintf("%s/%d.dem", path, match.ID)
	// Check if file exists and in download or parse stage
	_, err = os.Stat(filename)
	if err == nil {
		return FileAlreadyExistsError{filename: filename}
	}
	// Create a new file to save the decompressed content
	file, err := os.Create(filename)
	if err != nil {
		return FileCreationError{filename: filename}
	}
	defer file.Close()

	// Create a bzip2 reader to decompress the content
	reader := bzip2.NewReader(response.Body)

	// Copy the decompressed content to the file
	_, err = io.Copy(file, reader)
	if err != nil {
		return CopyError{}
	}

	// Decompression completed
	return nil
}
