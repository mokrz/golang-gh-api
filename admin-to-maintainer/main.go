package main

import (
	"net/http"
	"time"
)

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
	_ = "mokrz-enterprise-testing"
	token := "asdf"

	_ = GitHubAPIClient{
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
			Transport: GitHubRoundTripper{
				NextMiddleWare: http.DefaultTransport,
				Headers: map[string]string{
					"Authorization": "token " + token,
					"Accept":        "application/vnd.github.v3+json",
				},
			},
		},
		BaseURL: "https://api.github.com/v3/",
	}
}
