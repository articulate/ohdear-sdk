package ohdear

import (
	"fmt"
	"net/http"
)

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

type CheckService struct {
	client *Client
}

func (c *CheckService) EnableCheck(check *Check) (*http.Response, error) {
	checkPath := fmt.Sprintf("/api/checks/%d/enable", check.Id)

	req, err := c.client.NewRequest("POST", checkPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.do(req, nil)
	return resp, err
}

func (c *CheckService) DisableCheck(check *Check) (*http.Response, error) {
	checkPath := fmt.Sprintf("/api/checks/%d/disable", check.Id)

	req, err := c.client.NewRequest("POST", checkPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.do(req, nil)
	return resp, err
}
