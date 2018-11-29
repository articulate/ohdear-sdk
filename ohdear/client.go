package ohdear

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	ApiToken      string
	httpClient    *http.Client
	RateLimitOver time.Time // When rate-limiting ends

	SiteService  *SiteService
	CheckService *CheckService
	TeamService  *TeamService
}

func NewClient(baseURL string, apiToken string) (*Client, error) {
	httpClient := http.DefaultClient

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		httpClient: httpClient,
		ApiToken:   apiToken,
		BaseURL:    u,
	}

	c.SiteService = &SiteService{client: c}
	c.CheckService = &CheckService{client: c}
	c.TeamService = &TeamService{client: c}

	return c, nil
}

func (c *Client) validate() (bool, error) {
	_, _, err := c.TeamService.ListTeams()
	return err == nil, err
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiToken)
	req.Header.Set("UserAgent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	if time.Now().UnixNano() > c.RateLimitOver.UnixNano() {
		resp, err := c.httpClient.Do(req)

		if err != nil {
			return nil, err
		} else if resp.StatusCode == 429 {
			secLeft, err := strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining"))
			if err != nil {
				err = fmt.Errorf("Error while parsing backoff header: %v", err)
				return resp, err
			}
			durSeconds := time.Duration(secLeft) * time.Second
			c.RateLimitOver = time.Now().Add(durSeconds)
		} else if resp.StatusCode >= 300 {
			err = fmt.Errorf("Invalid Status: %d", resp.StatusCode)

			return resp, err
		}

		if v != nil {
			err = json.NewDecoder(resp.Body).Decode(v)
		}

		return resp, err
	} else {
		secLeft := c.RateLimitOver.Sub(time.Now())
		fmt.Printf("[WARN] Rate limiting in effect, retrying in %s sec...", secLeft)

		time.Sleep(secLeft)
		return c.do(req, v)
	}
}
