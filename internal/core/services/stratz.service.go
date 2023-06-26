package services

import (
	"context"
	"github.com/machinebox/graphql"
	"go-glyph-v2/internal/core/dtos"
)

type StratzService struct {
	bearerToken string
	client      *graphql.Client
}

func NewStratzService(bearerToken string) *StratzService {
	return &StratzService{bearerToken: bearerToken, client: graphql.NewClient("https://api.stratz.com/graphql")}
}

func (s StratzService) GetMatchFromStratzAPI(matchID int) (dtos.Match, error) {
	query := graphql.NewRequest(`
    query($key: Long!) {
		match(id: $key) {
		  replaySalt
		  clusterId
		}
	}`)
	query.Var("key", matchID)

	stratzToken := "Bearer " + s.bearerToken
	query.Header.Set("Authorization", stratzToken)

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData struct {
		Match struct {
			ReplaySalt int `json:"replaySalt"`
			ClusterID  int `json:"clusterId"`
		} `json:"match"`
	}
	if err := s.client.Run(ctx, query, &respData); err != nil {
		return dtos.Match{}, err
	}
	return dtos.Match{
		ID:         matchID,
		Cluster:    respData.Match.ClusterID,
		ReplaySalt: respData.Match.ReplaySalt,
	}, nil
}
