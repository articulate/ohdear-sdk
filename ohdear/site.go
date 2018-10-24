package ohdear

import "fmt"

// Types of check
const (
	UptimeCheck           = "uptime"
	BrokenLinksCheck      = "broken_links"
	CertHealthCheck       = "certificate_health"
	MixedContentCheck     = "mixed_content"
	CertTransparencyCheck = "certificate_transparency"
)

type Check struct {
	Id      int    `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

type Site struct {
	Id                    int     `json:"id,omitempty"`
	Url                   string  `json:"url,omitempty"`
	TeamId                int     `json:"team_id,omitempty"`
	LatestRunDate         string  `json:"latest_run_date,omitempty"`
	SummarizedCheckResult string  `json:"summarized_check_result,omitempty"`
	CreatedAt             string  `json:"created_at,omitempty"`
	UpdatedAt             string  `json:"updates_at,omitempty"`
	Checks                []Check `json:"checks,omitempty"`
}

type SiteService struct {
	client *Client
}

func (s *SiteService) ListSites() ([]Site, error) {
	req, err := s.client.NewRequest("GET", "/api/sites", []string{})
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

func (s *SiteService) CreateSite(site *Site) (*Site, error) {
	req, err := s.client.NewRequest("POST", "/api/sites", site)
	if err != nil {
		return nil, err
	}

	newSite := &Site{}

	_, err = s.client.do(req, &newSite)
	if err != nil {
		return nil, err
	}

	return newSite, err
}

func (s *SiteService) DeleteSite(site *Site) error {
	sitePath := fmt.Sprintf("/api/sites/%d", site.Id)

	req, err := s.client.NewRequest("DELETE", sitePath, nil)

	_, err = s.client.do(req, nil)

	if err != nil {
		return err
	}
	return nil
}

// func (s *SiteService) UpdateSite(s *Site) (Site, error) {}
