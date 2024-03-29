package ohdear_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/articulate/ohdear-sdk/ohdear"
)

var _ = Describe("Check", func() {

	const (
		testBaseURL = "http://test.org"
		testToken   = "foobarbazquux" //nolint:gosec
	)

	var (
		client *Client
	)

	BeforeEach(func() {
		client, _ = NewClient(testBaseURL, testToken, nil)
	})

	Context("POST /api/sites/:site/enable", func() {
		It("Should return a 204", func() {
			gock.New(testBaseURL).
				Post("/api/checks/42/enable").
				Reply(204)

			resp, err := client.CheckService.EnableCheck(42)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})

	Context("POST /api/sites/:site/disable", func() {
		It("Should return a 204", func() {
			gock.New(testBaseURL).
				Post("/api/checks/42/disable").
				Reply(204)

			resp, err := client.CheckService.DisableCheck(42)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})
})
