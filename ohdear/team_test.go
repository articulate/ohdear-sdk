package ohdear_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/articulate/ohdear-sdk/ohdear"
)

var _ = Describe("Team", func() {

	const (
		testBaseURL = "http://test.org"
		testToken   = "foobarbazquux" //nolint:gosec
	)

	var (
		client *Client
	)

	Context("GET /api/me", func() {

		BeforeEach(func() {
			client, _ = NewClient(testBaseURL, testToken, nil)
		})

		It("Should get a list of teams", func() {

			userInfo := &UserInfo{
				ID: 1,
				Teams: []Team{
					{
						ID:   1,
						Name: "The Goonies",
					},
				},
			}

			gock.New("http://test.org").
				Get("/api/me").
				Reply(200).
				JSON(userInfo)

			newTeams, _, err := client.TeamService.ListTeams()

			Expect(err).To(BeNil())
			Expect(newTeams).To(Equal(userInfo.Teams))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})
})
