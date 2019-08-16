package ohdear_test

import (
	"fmt"

	. "github.com/articulate/ohdear-sdk/ohdear"
	"github.com/articulate/ohdear-sdk/ohdear/mocks"
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
		client      *Client
		mockSleeper *mocks.MockSleeper
	)

	BeforeEach(func() {
		client, _ = NewClient(testBaseURL, testToken, nil)
		mockSleeper = &mocks.MockSleeper{}
		client.Sleeper = mockSleeper
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
		It("should call the sleeper", func() {
			sites := []*Site{}
			gock.New(testBaseURL).
				Get("/api/sites").
				Reply(429).
				SetHeader("X-RateLimit-Reset", "10").
				JSON("[]")

			gock.New(testBaseURL).
				Get("/api/sites").
				Reply(200).
				JSON(sites)

			_, _, err := client.SiteService.ListSites()

			Expect(err).To(BeNil())
			Expect(mockSleeper.SleepCall.Count).To(Equal(1))
		})
	})
})
