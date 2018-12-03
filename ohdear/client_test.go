package ohdear_test

import (
	"fmt"

	. "github.com/articulate/ohdear-sdk/ohdear"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"
)

var _ = Describe("./Client", func() {

	const (
		testBaseURL = "http://test.org"
		testToken   = "foobarbazquux"
	)

	var (
		client *Client
	)

	BeforeEach(func() {
		client, _ = NewClient(testBaseURL, testToken)
	})

	Context("Base URL", func() {
		It("Should be correct", func() {
			Expect(client.BaseURL.String()).To(Equal(testBaseURL))
		})
	})

	// TODO Use reflection to gain access to newRequest as
	// private function
	Context("API Token", func() {
		It("Should be in the correct header", func() {
			var res []string
			req, err := client.NewRequest("get", "/some/path", res)
			Expect(err).To(BeNil())

			header := req.Header.Get("Authorization")
			wantHeader := fmt.Sprintf("Bearer %v", testToken)

			Expect(header).To(Equal(wantHeader))
		})
	})

	Context("Rate Limiting", func() {
		It("should respect 429 headers", func() {
			sites := []*Site{}

			gock.New(testBaseURL).
				Get("/api/sites").
				Reply(429).
				SetHeader("X-RateLimit-Remaining", "3").
				JSON("[]")

			gock.New(testBaseURL).
				Get("/api/sites").
				Reply(200).
				JSON(sites)

			_, _, err := client.SiteService.ListSites()

			Expect(err).To(BeNil())
		})
	})
})
