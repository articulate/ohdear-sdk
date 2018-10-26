package ohdear_test

import (
	. "github.com/articulate/ohdear-sdk/ohdear"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"
)

var _ = Describe("Team", func() {

	const (
		testBaseURL = "http://test.org"
		testToken   = "foobarbazquux"
	)

	var (
		client *Client
	)

	Context("GET /api/me", func() {

		BeforeEach(func() {
			client, _ = NewClient(testBaseURL, testToken)
		})

		It("Should get a list of teams", func() {

			teams := []*Team{
				&Team{
					Id:   1,
					Name: "Goonies",
				},
			}

			userinfo := &UserInfo{
				Id:    1,
				Teams: teams,
			}

			gock.New("http://test.org").
				Get("/api/me").
				Reply(200).
				JSON(userinfo)

			newTeams, _, err := client.TeamService.ListTeams()

			Expect(err).To(BeNil())
			Expect(newTeams).To(Equal(teams))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})
})
