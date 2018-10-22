package ohdear

import (
	"encoding/json"
)

// TODO Add checks
type Site struct {
	Id                    int    `json:"id"`
	Url                   string `json:"url"`
	TeamId                string `json:"team_id"`
	LatestRunDate         string `json:"latest_run_date"`
	SummarizedCheckResult string `json:"summarized_check_result"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updates_at"`
}

type SiteService struct {
	client *Client
}

func (s *SiteService) ListSites() ([]Site, error) {
	req, err := s.client.newRequest("get", "api/sites", nil)
	if err != nil {
		return nil, err
	}
	var sites []Site
	resp, err := s.client.do(req, sites)
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&sites)
	return sites, err
}

// func (s *SiteService) CreateSite(s *Site) (Site, error) {}
// func (s *SiteService) DeleteSite(s *Site) error         {}
// func (s *SiteService) UpdateSite(s *Site) (Site, error) {}
