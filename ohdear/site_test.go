package ohdear_test

import (
	. "github.com/articulate/ohdear-sdk/ohdear"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"
)

var _ = Describe("Site", func() {

	const (
		testBaseURL = "http://test.org"
		testToken   = "foobarbazquux"
	)

	var (
		client *Client
	)

	Context("GET /api/sites", func() {

		BeforeEach(func() {
			client, _ = NewClient(testBaseURL, testToken)
		})

		It("Should get a list of sites", func() {

			sites := []Site{
				Site{
					Id:     1,
					Url:    "http://foobar.com",
					TeamId: 170,
					Checks: []*Check{
						&Check{
							Id:      1,
							Type:    UptimeCheck,
							Enabled: true,
						},
						&Check{
							Id:      1,
							Type:    BrokenLinksCheck,
							Enabled: true,
						},
					},
				},
			}

			gock.New("http://test.org").
				Get("/api/sites").
				Reply(200).
				JSON(sites)

			res, err := client.SiteService.ListSites()

			Expect(err).To(BeNil())
			Expect(res).To(Equal(sites))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})

	Context("GET /api/sites/:id", func() {

		BeforeEach(func() {
			client, _ = NewClient(testBaseURL, testToken)
		})

		It("Should get the site by ID", func() {

			siteData := &Site{
				Id:     1,
				Url:    "http://foobar.com",
				TeamId: 170,
				Checks: []*Check{
					&Check{
						Id:      1,
						Type:    UptimeCheck,
						Enabled: true,
					},
					&Check{
						Id:      1,
						Type:    BrokenLinksCheck,
						Enabled: true,
					},
				},
			}

			gock.New("http://test.org").
				Get("/api/sites/1").
				Reply(200).
				JSON(siteData)

			site, _, err := client.SiteService.GetSite(1)

			Expect(err).To(BeNil())
			Expect(site).To(Equal(siteData))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})

	Context("POST /api/sites", func() {

		BeforeEach(func() {
			client, _ = NewClient(testBaseURL, testToken)
		})

		It("should return a new site", func() {
			site := &Site{
				Url:    "http://foobar.com",
				TeamId: 170,
				Checks: []*Check{
					&Check{
						Type: UptimeCheck,
					},
					&Check{
						Type: BrokenLinksCheck,
					},
				},
			}

			responseSite := &Site{
				Id:     1,
				Url:    "http://foobar.com",
				TeamId: 170,
				Checks: []*Check{
					&Check{
						Id:      1,
						Type:    UptimeCheck,
						Enabled: true,
					},
					&Check{
						Id:      2,
						Type:    BrokenLinksCheck,
						Enabled: true,
					},
				},
			}

			gock.New("http://test.org").
				Post("/api/sites").
				MatchType("json").
				JSON(site).
				Reply(201).
				JSON(responseSite)

			site, _, err := client.SiteService.CreateSite(site)

			Expect(err).To(BeNil())
			Expect(site.Id).To(Equal(responseSite.Id))
			Expect(site.Url).To(Equal(responseSite.Url))
			Expect(site.TeamId).To(Equal(responseSite.TeamId))
			Expect(len(site.Checks)).To(Equal(len(responseSite.Checks)))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})

	Context("DELETE /api/sites", func() {

		BeforeEach(func() {
			client, _ = NewClient(testBaseURL, testToken)
		})

		It("should delete the specified site", func() {
			site := &Site{
				Id: 170,
			}

			gock.New("http://test.org").
				Delete("/api/sites/170").
				Reply(204)

			resp, err := client.SiteService.DeleteSite(site.Id)

			Expect(err).To(BeNil())
			Expect(gock.IsDone()).To(BeTrue())
			Expect(resp.Status).To(Equal("204 No Content"))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})
})
