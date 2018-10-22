package ohdear

import (
	"fmt"
	"testing"

	"github.com/nbio/st"
)

const (
	testBaseURL = "http://test.org"
	testToken   = "foobarbazquux"
)

func TestClientSetup(t *testing.T) {
	// wantPath := "api/v1/sites"
	client, err := NewClient(testBaseURL, testToken)

	st.Assert(t, nil, err)
	st.Assert(t, client.BaseURL.String(), testBaseURL)
}

func TestClientToken(t *testing.T) {
	client, err := NewClient(testBaseURL, testToken)

	var res []string
	req, err := client.newRequest("get", "/some/path", res)
	st.Assert(t, nil, err)

	header := req.Header.Get("Authorization")
	wantHeader := fmt.Sprintf("Bearer %v", testToken)
	st.Assert(t, header, wantHeader)
}
