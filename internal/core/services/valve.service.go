package services

import (
	"bufio"
	"compress/bzip2"
	"fmt"
	"go-glyph-v2/internal/core/dtos"
	"io"
	"net/http"
	"os"
)

const (
	baseReplayURL = "http://replay%d.valve.net/570/%d_%d.dem.bz2"
	demosPath     = "internal/data/demos"
)

type ValveService struct {
	client http.Client
}

func NewValveService() *ValveService {
	return &ValveService{client: http.Client{}}
}

func (s ValveService) RetrieveFile(match dtos.Match) error {
	url := fmt.Sprintf(baseReplayURL, match.Cluster, match.ID, match.ReplaySalt)
	response, err := s.client.Get(url)
	if err != nil {
		return GETError{url: url, error: err}
	}
	if response.StatusCode != 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return ReadResponseBodyError{err}
		}
		return HTTPError{url: url, statusCode: response.StatusCode, response: string(body)}
	}
	defer response.Body.Close()

	if _, err := os.Stat(demosPath); os.IsNotExist(err) {
		err := os.Mkdir(demosPath, os.ModePerm)
		if err != nil {
			return FolderCreationError{foldername: demosPath}
		}
	}
	filename := fmt.Sprintf("%s/%d.dem", demosPath, match.ID)
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

	bufferedWriter := bufio.NewWriter(file)
	defer bufferedWriter.Flush()
	// Create a bzip2 reader to decompress the content
	reader := bzip2.NewReader(response.Body)

	// Copy the decompressed content to the file
	_, err = io.Copy(bufferedWriter, reader)
	if err != nil {
		return CopyError{}
	}

	// Decompression completed
	return nil
}
