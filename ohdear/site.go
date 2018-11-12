package ohdear

import (
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

type (
	// Site represents an OhDear Site resource
	Site struct {
		ID                    int     `json:"id,omitempty"`
		URL                   string  `json:"url,omitempty"`
		TeamID                int     `json:"team_id,omitempty"`
		LatestRunDate         string  `json:"latest_run_date,omitempty"`
		SummarizedCheckResult string  `json:"summarized_check_result,omitempty"`
		CreatedAt             string  `json:"created_at,omitempty"`
		UpdatedAt             string  `json:"updates_at,omitempty"`
		Checks                []Check `json:"checks,omitempty"`
	}

	// SiteRequest represents the site creation request body
	SiteRequest struct {
		URL    string   `json:"url,omitempty"`
		TeamID int      `json:"team_id,omitempty"`
		Checks []string `json:"checks,omitempty"`
	}

	// SiteService is a service object used to access the site API resource
	SiteService struct {
		client *Client
	}
)

// ListSites returns the index of sites from the `/sites` endpoint
func (s *SiteService) ListSites() ([]*Site, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "/api/sites", nil)
	if err != nil {
		log.Errorf("Error creating request: %v", err.Error())
		return nil, nil, err
	}

	var sites []*Site

	resp, err := s.client.do(req, &sites)
	if err != nil {
		log.Errorf("Error retrieving sites from OhDear: %v", err)
		return nil, resp, err
	}

	return sites, resp, err
}

// GetSite retrieves a site from the `/sites/:id` endpoint
func (s *SiteService) GetSite(siteID int) (*Site, *http.Response, error) {
	sitePath := fmt.Sprintf("/api/sites/%d", siteID)

	site := &Site{}
	req, err := s.client.NewRequest("GET", sitePath, site)
	if err != nil {
		log.Errorf("Error creating request: %v", err)
		return nil, nil, err
	}

	resp, err := s.client.do(req, site)
	if err != nil {
		log.Errorf("Error retrieving site from OhDear: %v", err)
		return nil, resp, err
	}

	return site, resp, err
}

// CreateSite performs a POST call to the `/sites/` endpoint
func (s *SiteService) CreateSite(site *SiteRequest) (*Site, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "/api/sites", site)
	if err != nil {
		log.Errorf("Error creating request: %v", err)
		return nil, nil, err
	}

	newSite := &Site{}

	resp, err := s.client.do(req, &newSite)
	if err != nil {
		log.Errorf("Error creating site on OhDear: %v", err)
		return nil, resp, err
	}

	return newSite, resp, err
}

// DeleteSite performs a DELETE call to the `/sites/:id` endpoint
func (s *SiteService) DeleteSite(siteID int) (*http.Response, error) {
	sitePath := fmt.Sprintf("/api/sites/%d", siteID)

	req, err := s.client.NewRequest("DELETE", sitePath, nil)
	if err != nil {
		log.Errorf("Error creating request: %v", err)
		return nil, err
	}

	resp, err := s.client.do(req, nil)
	if err != nil {
		log.Errorf("Error deleting site from OhDear: %v", err)
	}

	return resp, err
}
