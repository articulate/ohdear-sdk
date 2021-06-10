package ohdear_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/articulate/ohdear-sdk/ohdear"
)

var _ = Describe("Site", func() {

	const (
		testBaseURL = "http://test.org"
		testToken   = "foobarbazquux"
	)

	var (
		client          *Client
		testSite        *Site
		testSiteRequest *SiteRequest
	)

	BeforeEach(func() {
		client, _ = NewClient(testBaseURL, testToken, nil)
		testSite = &Site{
			ID:     1,
			URL:    "http://foobar.com",
			TeamID: 170,
			Checks: []Check{
				{
					ID:      1,
					Type:    UptimeCheck,
					Enabled: true,
				},
				{
					ID:      1,
					Type:    BrokenLinksCheck,
					Enabled: true,
				},
			},
		}
		testSiteRequest = &SiteRequest{
			URL:    testSite.URL,
			TeamID: testSite.ID,
			Checks: []string{
				BrokenLinksCheck,
				UptimeCheck,
			},
		}
	})

	Context("GET /api/sites", func() {
		It("Should get a list of sites", func() {

			sites := SiteList{Sites: []*Site{testSite}}

			gock.New(testBaseURL).
				Get("/api/sites").
				Reply(200).
				JSON(sites)

			actual, resp, err := client.SiteService.ListSites()

			Expect(err).To(BeNil())

			for i, expected := range sites.Sites {
				assertSite(expected, actual[i])
			}
			Expect(resp.StatusCode).To(Equal(200))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})

	Context("GET /api/sites/:id", func() {
		It("Should get the site by ID", func() {
			gock.New(testBaseURL).
				Get("/api/sites/1").
				Reply(200).
				JSON(testSite)

			site, _, err := client.SiteService.GetSite(1)

			Expect(err).To(BeNil())
			Expect(site.ID).To(Equal(testSite.ID))
			Expect(site.URL).To(Equal(testSite.URL))
			Expect(site.Checks).To(Equal(testSite.Checks))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})

	Context("POST /api/sites", func() {
		It("should return a new site", func() {
			gock.New(testBaseURL).
				Post("/api/sites").
				MatchType("json").
				JSON(testSiteRequest).
				Reply(201).
				JSON(testSite)

			site, _, err := client.SiteService.CreateSite(testSiteRequest)

			Expect(err).To(BeNil())
			Expect(gock.IsDone()).To(BeTrue())
			Expect(site.ID).To(Equal(testSite.ID))
			Expect(site.URL).To(Equal(testSite.URL))
			Expect(site.TeamID).To(Equal(testSite.TeamID))

			for i, c := range testSite.Checks {
				Expect(site.Checks[i].Type).To(Equal(c.Type))
			}
		})
	})

	Context("DELETE /api/sites", func() {
		It("should delete the specified site", func() {
			gock.New(testBaseURL).
				Delete("/api/sites/170").
				Reply(204)

			resp, err := client.SiteService.DeleteSite(170)

			Expect(err).To(BeNil())
			Expect(gock.IsDone()).To(BeTrue())
			Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})
})

func assertSite(expected *Site, actual *Site) {
	Expect(expected.ID).To(Equal(actual.ID))
	Expect(expected.URL).To(Equal(actual.URL))
	Expect(expected.TeamID).To(Equal(actual.TeamID))
	Expect(len(expected.Checks)).To(Equal(len(actual.Checks)))

	for i, c := range expected.Checks {
		Expect(actual.Checks[i].Type).To(Equal(c.Type))
	}
}
