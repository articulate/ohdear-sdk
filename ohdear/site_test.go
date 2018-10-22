package ohdear

import (
	"testing"

	"github.com/nbio/st"
	gock "gopkg.in/h2non/gock.v1"
)

var client Client

func setup() {
	client, _ := NewClient(testBaseURL, testToken)
}

func TestGetSite(t *testing.T) {
	defer gock.Off()

	var resBody []map[string]string
	firstSite := map[string]string{
		"id":                      "1",
		"url":                     "http://yoursite.tld",
		"sort_url":                "yoursite.tld",
		"team_id":                 "1",
		"latest_run_date":         "2017-12-05 20:02:02",
		"summarized_check_result": "succeeded",
		"created_at":              "20171106 07:40:49",
		"updatedAt":               "20171106 07:40:49",
	}
	resBody = append(resBody, firstSite)

	gock.New(testBaseURL).
		Get("api/sites").
		Reply(200).
		JSON(resBody)

	_, err := *client.SiteService.ListSites()
	st.Assert(t, nil, err)
}
