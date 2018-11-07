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

// Collection of all check types - for iteration
var CheckTypes = []string{
	UptimeCheck,
	BrokenLinksCheck,
	CertHealthCheck,
	MixedContentCheck,
	CertTransparencyCheck,
}

type Check struct {
	ID      int    `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

type CheckService struct {
	client *Client
}

func (c *CheckService) EnableCheck(check *Check) (*http.Response, error) {
	return c.performCheckAction(check.ID, "enable")
}

func (c *CheckService) DisableCheck(check *Check) (*http.Response, error) {
	return c.performCheckAction(check.ID, "disable")
}

func (c *CheckService) performCheckAction(id int, lifecycleAction string) (*http.Response, error) {
	checkPath := fmt.Sprintf("/api/checks/%d/%s", id, lifecycleAction)

	req, err := c.client.NewRequest("POST", checkPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.do(req, nil)
	return resp, err
}
