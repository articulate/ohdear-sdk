package ohdear

import (
	"net/http"
)

// TeamData holds a portion of the response from OhDear's `me` API endpoint
type TeamData struct {
	Teams []*Team `json:"teams,omitempty"`
}

// UserInfo is the top-level struct representing data returned from OhDear's `me` API endpoint
type UserInfo struct {
	ID       int       `json:"id,omitempty"`
	TeamData *TeamData `json:"data,omitempty"`
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
func (t *TeamService) ListTeams() ([]*Team, *http.Response, error) {
	req, err := t.client.NewRequest("GET", "/api/me", nil)

	if err != nil {
		return nil, nil, err
	}

	var userinfo = &UserInfo{}

	resp, err := t.client.do(req, userinfo)
	return userinfo.TeamData.Teams, resp, err
}
