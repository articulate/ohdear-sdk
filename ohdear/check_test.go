package ohdear_test

import (
	. "github.com/articulate/ohdear-sdk/ohdear"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"
)

var _ = Describe("Check", func() {

	const (
		testBaseURL = "http://test.org"
		testToken   = "foobarbazquux"
	)

	var (
		client *Client
	)

	Context("POST /api/sites/:site/enable", func() {

		BeforeEach(func() {
			client, _ = NewClient(testBaseURL, testToken)
		})

		It("Should return a 204", func() {

			check := &Check{
				ID: 42,
			}

			gock.New("http://test.org").
				Post("/api/checks/42/enable").
				Reply(204)

			resp, err := client.CheckService.EnableCheck(check)

			Expect(err).To(BeNil())
			Expect(resp.Status).To(Equal("204 No Content"))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})

	Context("POST /api/sites/:site/disable", func() {

		BeforeEach(func() {
			client, _ = NewClient(testBaseURL, testToken)
		})

		It("Should return a 204", func() {

			check := &Check{
				ID: 42,
			}

			gock.New("http://test.org").
				Post("/api/checks/42/disable").
				Reply(204)

			resp, err := client.CheckService.DisableCheck(check)

			Expect(err).To(BeNil())
			Expect(resp.Status).To(Equal("204 No Content"))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})
})
