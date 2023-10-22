package services

import (
	"encoding/json"
	"fmt"
	"go-glyph-v2/internal/core/dtos"
	"net/http"
)

type OpendotaService struct {
	client http.Client
}

func NewOpendotaService() *OpendotaService {
	return &OpendotaService{client: http.Client{}}
}

func (s OpendotaService) GetMatchFromOpendotaAPI(matchID int) (dtos.Match, error) {
	url := fmt.Sprintf("https://api.opendota.com/api/matches/%d", matchID)
	// Get cluster and replay salt from Open Dota API
	response, err := s.client.Get(url)
	if err != nil {
		return dtos.Match{}, GETError{url: url, error: err}
	}
	defer response.Body.Close()
	var Match struct {
		ReplaySalt int `json:"replay_salt"`
		ClusterID  int `json:"cluster"`
	}
	err = json.NewDecoder(response.Body).Decode(&Match)
	if err != nil {
		return dtos.Match{}, err
	}

	return dtos.Match{
		ID:         matchID,
		Cluster:    Match.ClusterID,
		ReplaySalt: Match.ReplaySalt,
	}, nil
}
