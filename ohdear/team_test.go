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

			userInfo := &UserInfo{
				ID: 1,
				TeamData: &TeamData{
					Teams: []*Team{
						&Team{
							ID:   1,
							Name: "The Goonies",
						},
					},
				},
			}

			gock.New("http://test.org").
				Get("/api/me").
				Reply(200).
				JSON(userInfo)

			newTeams, _, err := client.TeamService.ListTeams()

			Expect(err).To(BeNil())
			Expect(newTeams).To(Equal(userInfo.TeamData.Teams))
			Expect(gock.IsDone()).To(BeTrue())
		})
	})
})
