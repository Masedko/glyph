package services

import (
	"compress/bzip2"
	"fmt"
	"io"
	"net/http"
	"os"
	"tiktok-arena/internal/core/dtos"
)

type ValveService struct {
	client http.Client
}

func NewValveService() *ValveService {
	return &ValveService{client: http.Client{}}
}

func (s ValveService) RetrieveFile(match dtos.Match) error {
	url := fmt.Sprintf("http://replay%d.valve.net/570/%d_%d.dem.bz2", match.Cluster, match.ID, match.ReplaySalt)
	response, err := s.client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	// Create a new file to save the decompressed content
	filename := fmt.Sprintf("internal/data/demos/%d.dem", match.ID)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a bzip2 reader to decompress the content
	reader := bzip2.NewReader(response.Body)

	// Copy the decompressed content to the file
	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}

	// Decompression completed
	return nil
}
