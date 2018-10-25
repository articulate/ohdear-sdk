package ohdear

import (
	"fmt"
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

func (c *CheckService) EnableCheck(check *Check) error {
	checkPath := fmt.Sprintf("/api/checks/%d/enable", check.Id)

	req, err := c.client.NewRequest("POST", checkPath, nil)
	if err != nil {
		return err
	}

	_, err = c.client.do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *CheckService) DisableCheck(check *Check) error {
	checkPath := fmt.Sprintf("/api/checks/%d/disable", check.Id)

	req, err := c.client.NewRequest("POST", checkPath, nil)
	if err != nil {
		return err
	}

	_, err = c.client.do(req, nil)
	if err != nil {
		return err
	}
	return nil
}
