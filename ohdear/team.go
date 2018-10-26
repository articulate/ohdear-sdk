package ohdear

import (
	"net/http"
)

type UserInfo struct {
	Id    int     `json:"id,omitempty"`
	Teams []*Team `json:"teams.data,omitempty"`
}

type Team struct {
	Id   int    `json:"teams.data.id,omitempty"`
	Name string `json:"teams.data.name,omitempty"`
}

type TeamService struct {
	client *Client
}

func (t *TeamService) ListTeams() ([]*Team, *http.Response, error) {
	req, err := t.client.NewRequest("GET", "/api/me", nil)

	if err != nil {
		return nil, nil, err
	}

	var userinfo = &UserInfo{}

	resp, err := t.client.do(req, userinfo)
	return userinfo.Teams, resp, err
}
