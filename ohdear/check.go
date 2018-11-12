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

// CheckTypes Collection of all check types - for iteration
var CheckTypes = []string{
	UptimeCheck,
	BrokenLinksCheck,
	CertHealthCheck,
	MixedContentCheck,
	CertTransparencyCheck,
}

type (
	// Check represents a Site Check
	Check struct {
		ID      int    `json:"id,omitempty"`
		Type    string `json:"type,omitempty"`
		Enabled bool   `json:"enabled,omitempty"`
	}

	// CheckService service for interacting with checks
	CheckService struct {
		client *Client
	}
)

// EnableCheck activation endpoint for a site check
func (c *CheckService) EnableCheck(id int) (*http.Response, error) {
	return c.performCheckAction(id, "enable")
}

// DisableCheck deactivation endpoint for a site check
func (c *CheckService) DisableCheck(id int) (*http.Response, error) {
	return c.performCheckAction(id, "disable")
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
