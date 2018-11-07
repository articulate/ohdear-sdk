package ohdear

import (
	"net/http"

	"github.com/prometheus/common/log"
)

// UserInfo is the top-level struct representing data returned from OhDear's `me` API endpoint
type UserInfo struct {
	ID    int    `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Teams []Team `json:"teams,omitempty"`
}

// Team holds the data for each team existing in OhDear
type Team struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// TeamService is a service object used to access the UserInfo API endpoint
type TeamService struct {
	client *Client
}

// ListTeams hits the `me` API endpoint and returns a list of teams
func (t *TeamService) ListTeams() ([]Team, *http.Response, error) {
	req, err := t.client.NewRequest("GET", "/api/me", nil)

	if err != nil {
		return nil, nil, err
	}

	var userinfo = &UserInfo{}

	resp, err := t.client.do(req, userinfo)
	if err != nil {
		log.Fatalf("Error retrieving teams from OhDear: %s", err)
	}

	if resp.StatusCode >= 300 {
		log.Errorf("Error Retrieving teams from OhDear: %v", err.Error)
		return nil, resp, err
	}
	return userinfo.Teams, resp, err
}
