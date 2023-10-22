package services

import (
	"context"
	"github.com/machinebox/graphql"
	"testing"
)

func TestStratzService_GetMatchFromStratzAPI(t *testing.T) {

	query := graphql.NewRequest(`
    query($key: Long!) {
		match(id: $key) {
		  replaySalt
		  clusterId
		}
	}`)
	query.Var("key", 7387654378)

	stratzToken := "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdWJqZWN0IjoiYjkxYzQwMTctYmM0MC00ODUzLThiYjAtN2JmZDc4MjU1NTA2IiwiU3RlYW1JZCI6IjM3Mzc2NjA4NiIsIm5iZiI6MTY5NzgyMTM3MSwiZXhwIjoxNzI5MzU3MzcxLCJpYXQiOjE2OTc4MjEzNzEsImlzcyI6Imh0dHBzOi8vYXBpLnN0cmF0ei5jb20ifQ.z-nCTA4aRbFmmmZujl_ZpdeZ8zHGMtOSBMTk6pKONew"
	query.Header.Set("Authorization", stratzToken)

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData struct {
		Match struct {
			ReplaySalt int64 `json:"replaySalt"`
			ClusterID  int   `json:"clusterId"`
		} `json:"match"`
	}
	if err := graphql.NewClient("https://api.stratz.com/graphql").Run(ctx, query, &respData); err != nil {
		t.Error("Error during Client Run")
	}
	t.Log(respData)
}
