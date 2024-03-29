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

type (
	Sleeper interface {
		Sleep(time.Duration)
	}

	StdLibSleeper struct{}
)

func (s StdLibSleeper) Sleep(seconds time.Duration) {
	time.Sleep(seconds)
}

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	APIToken      string
	httpClient    *http.Client
	RateLimitOver time.Time // When rate-limiting ends

	SiteService  *SiteService
	CheckService *CheckService
	TeamService  *TeamService

	Sleeper
}

func NewClient(baseURL string, apiToken string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Invalid base URL provided to SDK, error: %v", err)
	}

	c := &Client{
		APIToken:   apiToken,
		BaseURL:    u,
		httpClient: httpClient,
	}

	c.SiteService = &SiteService{client: c}
	c.CheckService = &CheckService{client: c}
	c.TeamService = &TeamService{client: c}

	c.Sleeper = StdLibSleeper{}
	return c, nil
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
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	req.Header.Set("UserAgent", c.UserAgent)

	return req, nil
}

func (c *Client) timeLeftToWait() time.Duration {
	return time.Until(c.RateLimitOver)
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode == 429 {
		secLeft, err := strconv.Atoi(resp.Header.Get("X-RateLimit-Reset"))
		if err != nil {
			err = fmt.Errorf("Error while parsing backoff header: %v", err)
			return resp, err
		}
		durSeconds := time.Duration(secLeft) * time.Second
		c.RateLimitOver = time.Now().Add(durSeconds)

		timeLeft := c.timeLeftToWait()
		fmt.Printf("[WARN] Rate limiting in effect, retrying in %s sec...", timeLeft)
		c.Sleeper.Sleep(timeLeft)
		return c.do(req, v)
	} else if resp.StatusCode >= 300 {
		var apiErr *APIError
		err = json.NewDecoder(resp.Body).Decode(apiErr)
		if err != nil {
			return resp, fmt.Errorf("API Error: %s", resp.Status)
		}
		return resp, apiErr
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
