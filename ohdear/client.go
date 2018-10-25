package ohdear

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/davecgh/go-spew/spew"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	ApiToken   string
	httpClient *http.Client

	SiteService     *SiteService
	CheckService    *CheckService
	UserInfoService *UserInfoService
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
	c.UserInfoService = &UserInfoService{client: c}
	return c, nil
}

func (c *Client) validate() (bool, error) {
	_, _, err := c.UserInfoService.GetUserInfo()
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
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
	spew.Dump(body)
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
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
