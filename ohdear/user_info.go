package ohdear

import "net/http"

type Team struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserInfo struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Email    string  `json:"email,omitempty"`
	PhotoURL string  `json:"photoUrl,omitempty"`
	Teams    []*Team `json:"teams,omitempty"`
}

type UserInfoService struct {
	client *Client
}

func (u *UserInfoService) GetUserInfo() (*http.Response, *UserInfo, error) {
	req, err := u.client.NewRequest("GET", "/api/sites/me", nil)
	if err != nil {
		return nil, nil, err
	}

	userInfo := &UserInfo{}
	resp, err := u.client.do(req, userInfo)

	return resp, userInfo, err
}
