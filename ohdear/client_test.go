package ohdear_test

import (
	"fmt"

	. "github.com/articulate/ohdear-sdk/ohdear"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
})
