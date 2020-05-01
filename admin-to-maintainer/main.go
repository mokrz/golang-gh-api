package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mokrz/golang-gh-api/admin-to-maintainer/types"
)

// ListOrganizationRepositories returns a list of repositories in the given organization
// https://developer.github.com/v3/repos/#list-organization-repositories
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

// ListRepositoryCollaborators returns a list of collaborators for the given repository
// https://developer.github.com/v3/repos/collaborators/#list-collaborators
func (c GitHubAPIClient) ListRepositoryCollaborators(repository string) ([]types.User, error) {
	req, newReqErr := http.NewRequest("GET", c.BaseURL+"repos/"+repository+"/collaborators", nil)

	if newReqErr != nil {
		return nil, errors.New("ListRepositoryCollaborators: http.NewRequest failed with error: " + newReqErr.Error())
	}

	res, doErr := c.HTTPClient.Do(req)

	if doErr != nil {
		return nil, errors.New("ListRepositoryCollaborators: HTTPClient.Do failed with error: " + doErr.Error())
	}

	var collaborators []types.User
	decodeErr := json.NewDecoder(res.Body).Decode(&collaborators)

	if decodeErr != nil {
		return nil, errors.New("ListRepositoryCollaborators: Decoder.Decode failed with error: " + decodeErr.Error())
	}

	return collaborators, nil
}

// GetCollaboratorPermissions returns a user's permission level against the given repository
// https://developer.github.com/v3/repos/collaborators/#review-a-users-permission-level
func (c GitHubAPIClient) GetCollaboratorPermissions(repository, user string) (*types.CollaboratorPermission, error) {
	req, newReqErr := http.NewRequest("GET", c.BaseURL+"repos/"+repository+"/collaborators/"+user+"/permission", nil)

	if newReqErr != nil {
		return nil, errors.New("GetUserPermissions: http.NewRequest failed with error: " + newReqErr.Error())
	}

	res, doErr := c.HTTPClient.Do(req)

	if doErr != nil {
		return nil, errors.New("GetUserPermissions: HTTPClient.Do failed with error: " + doErr.Error())
	}

	var permission types.CollaboratorPermission
	decodeErr := json.NewDecoder(res.Body).Decode(&permission)

	if decodeErr != nil {
		return nil, errors.New("GetUserPermissions: Decoder.Decode failed with error: " + decodeErr.Error())
	}

	return &permission, nil
}

// AddCollaborator adds a collaborator to a given repository with the given permission level. If the collaborator already exists, their permissions will be updated
// https://developer.github.com/v3/repos/collaborators/#add-user-as-a-collaborator
func (c GitHubAPIClient) AddCollaborator(repository, user, permission string) error {
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(types.CollaboratorPermission{
		Permission: "maintain",
	})

	if encodeErr != nil {
		return errors.New("AddCollaborator: Encoder.Encode failed with error: " + encodeErr.Error())
	}

	req, newReqErr := http.NewRequest("PUT", c.BaseURL+"repos/"+repository+"/collaborators/"+user, requestBody)

	if newReqErr != nil {
		return errors.New("AddCollaborator: http.NewRequest failed with error: " + newReqErr.Error())
	}

	res, doErr := c.HTTPClient.Do(req)

	if doErr != nil {
		return errors.New("AddCollaborator: HTTPClient.Do failed with error: " + doErr.Error())
	}

	if res.StatusCode != 204 && res.StatusCode != 201 {
		return errors.New("AddCollaborator: unexpected status code " + strconv.Itoa(res.StatusCode) + ", something went wrong")
	}

	return nil
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
	org := "REPLACEME"
	token := "REPLACEME"

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

	// Iterate over the list of repositories in the given organization and build a list of collaborators for each one.
	for _, repo := range repos {
		collaborators, listRepoCollabsErr := ghClient.ListRepositoryCollaborators(repo.FullName)

		if listRepoCollabsErr != nil {
			fmt.Printf("Failed to list collaborators for repo %s with error: %s\n", repo.FullName, listRepoCollabsErr.Error())
			continue
		}

		// For each collaborator in the current repo iteration, get their permission. If it is admin, set it to maintain. Ignore all other permissions.
		for _, collaborator := range collaborators {
			perms, getPermsErr := ghClient.GetCollaboratorPermissions(repo.FullName, collaborator.Login)

			if getPermsErr != nil {
				fmt.Printf("Failed to get collaborator permissions on repo %s for user %s with error %s\n", repo.FullName, collaborator.Login, getPermsErr.Error())
				continue
			}

			if perms.Permission == "admin" {
				fmt.Printf("User %s has %s on repo %s. Setting them to 'maintain'...\n", collaborator.Login, perms.Permission, repo.FullName)
				addCollabErr := ghClient.AddCollaborator(repo.FullName, collaborator.Login, "maintain")

				if addCollabErr != nil {
					fmt.Printf("Failed to set collaborator permissions on repo %s for user %s with error %s\n", repo.FullName, collaborator.Login, addCollabErr.Error())
					continue
				}
			}
		}
	}
}
