package ohdear

import (
	"fmt"
	"net/http"
)

type Site struct {
	Id                    int      `json:"id,omitempty"`
	Url                   string   `json:"url,omitempty"`
	TeamId                int      `json:"team_id,omitempty"`
	LatestRunDate         string   `json:"latest_run_date,omitempty"`
	SummarizedCheckResult string   `json:"summarized_check_result,omitempty"`
	CreatedAt             string   `json:"created_at,omitempty"`
	UpdatedAt             string   `json:"updates_at,omitempty"`
	Checks                []*Check `json:"checks,omitempty"`
}

type SiteService struct {
	client *Client
}

func (s *SiteService) ListSites() ([]Site, error) {
	req, err := s.client.NewRequest("GET", "/api/sites", nil)
	if err != nil {
		return nil, err
	}

	var sites []Site

	_, err = s.client.do(req, &sites)
	if err != nil {
		return nil, err
	}

	return sites, err
}

func (s *SiteService) GetSite(siteId int) (*Site, *http.Response, error) {
	sitePath := fmt.Sprintf("/api/sites/%d", siteId)

	site := &Site{}
	req, err := s.client.NewRequest("GET", sitePath, site)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.do(req, site)
	return site, resp, err
}

func (s *SiteService) CreateSite(site *Site) (*Site, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "/api/sites", site)
	if err != nil {
		return nil, nil, err
	}

	newSite := &Site{}

	resp, err := s.client.do(req, &newSite)
	if err != nil {
		return nil, resp, err
	}
	return newSite, resp, err
}

func (s *SiteService) DeleteSite(siteId int) (*http.Response, error) {
	sitePath := fmt.Sprintf("/api/sites/%d", siteId)

	req, err := s.client.NewRequest("DELETE", sitePath, nil)

	resp, err := s.client.do(req, nil)
	return resp, err
}
