package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mokrz/golang-gh-api/admin-to-maintainer/types"
)

// ListOrganizationRepositories returns a list of repositories in the given organization
func (c GitHubAPIClient) ListOrganizationRepositories(organization string) ([]types.Repository, error) {
	req, newReqErr := http.NewRequest("GET", c.BaseURL+"orgs/"+organization+"/repos", nil)

	if newReqErr != nil {
		return nil, errors.New("ListOrganizationRepositories: http.NewRequest failed with error: " + newReqErr.Error())
	}

	res, doErr := c.HTTPClient.Do(req)

	if doErr != nil {
		return nil, errors.New("ListOrganizationRepositories: HTTPClient.Do failed with error: " + doErr.Error())
	}

	var repos []types.Repository
	decodeErr := json.NewDecoder(res.Body).Decode(&repos)

	if decodeErr != nil {
		return nil, errors.New("ListOrganizationRepositories: Decoder.Decode failed with error: " + decodeErr.Error())
	}

	return repos, nil
}

// GitHubAPIClient allows us share resources between API calls.
type GitHubAPIClient struct {
	HTTPClient *http.Client
	BaseURL    string
}

// GitHubRoundTripper implements http.RoundTripper
type GitHubRoundTripper struct {
	NextMiddleWare http.RoundTripper
	Headers        map[string]string
}

// RoundTrip decorates the next http.RoundTripper middleware by including some additonal headers
func (rt GitHubRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {

	for header, value := range rt.Headers {
		r.Header.Add(header, value)
	}

	return rt.NextMiddleWare.RoundTrip(r)
}

func main() {
	org := "mokrz-enterprise-testing"
	token := "93200fc523d5a5206192d731c260bc6e4d95f21a"

	ghClient := GitHubAPIClient{
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
			Transport: GitHubRoundTripper{
				NextMiddleWare: http.DefaultTransport,
				Headers: map[string]string{
					"Authorization": "Bearer " + token,
					"Accept":        "application/vnd.github.v3+json",
				},
			},
		},
		BaseURL: "https://api.github.com/",
	}

	repos, listReposErr := ghClient.ListOrganizationRepositories(org)

	if listReposErr != nil {
		fmt.Printf("Failed to list repos for org %s with error: %s\n", org, listReposErr.Error())
		return
	}

	for _, repo := range repos {
		fmt.Println("Got repo: " + repo.FullName)
	}
}
