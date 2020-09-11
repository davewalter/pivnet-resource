package acceptance

import (
	"encoding/json"
	"github.com/onsi/gomega/gbytes"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/pivnet-resource/v2/concourse"
)

var _ = Describe("Check", func() {
	var (
		productSlug          = "pivnet-resource-test"
		oldestReleaseVersion = "0.0.1-piv-res-test-fixture#2017-06-30T15:41:17.119Z"

		command       *exec.Cmd
		checkRequest  concourse.CheckRequest
		stdinContents []byte
	)

	Context("when user provides a UAA refresh token", func() {
		BeforeEach(func() {
			By("Creating command object")
			command = exec.Command(checkPath)

			By("Creating default request")
			checkRequest = concourse.CheckRequest{
				Source: concourse.Source{
					APIToken:    refreshToken,
					ProductSlug: productSlug,
					Endpoint:    endpoint,
				},
				Version: concourse.Version{
					ProductVersion: oldestReleaseVersion,
				},
			}

			var err error
			stdinContents, err = json.Marshal(checkRequest)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("fetches a UAA token and uses that token to make the check request", func() {
			session := run(command, stdinContents)
			Eventually(session, executableTimeout).Should(gexec.Exit(0))

			By("Outputting a valid json response")
			response := concourse.CheckResponse{}
			err := json.Unmarshal(session.Out.Contents(), &response)
			Expect(err).ShouldNot(HaveOccurred())

			By("Validating all the expected elements were returned")
			Expect(len(response)).Should(BeNumerically(">=", 3))

			By("Validating the returned elements have the expected product release versions")
			Expect(response[0].ProductVersion).To(ContainSubstring("0.0.1-piv-res-test-fixture"))
			Expect(response[1].ProductVersion).To(ContainSubstring("1.2.3"))
			Expect(response[2].ProductVersion).To(ContainSubstring("2.3.4"))
		})

		Context("when validation fails", func() {
			BeforeEach(func() {
				checkRequest.Source.ProductSlug = ""

				var err error
				stdinContents, err = json.Marshal(checkRequest)
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("exits with error", func() {
				By("Running the command")
				session := run(command, stdinContents)

				By("Validating command exited with error")
				Eventually(session, executableTimeout).Should(gexec.Exit(1))
				Expect(session.Err).Should(gbytes.Say("product_slug must be provided"))
			})
		})
	})
})
