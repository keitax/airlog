package osenv_test

import (
	"os"

	"github.com/keitam913/textvid/infrastructure/osenv"

	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Config", func() {
	Describe("LoadConfig()", func() {
		Context("without required environment vars", func() {
			BeforeEach(func() {
				os.Unsetenv("PORT")
			})

			It("occurs an error", func() {
				conf, err := osenv.LoadConfig()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("missed env: TV_SITE_TITLE"))
				Expect(conf).To(BeNil())
			})
		})

		Context("with valid environment vars", func() {
			env := map[string]string{
				"TV_SITE_TITLE":                          "Textvid",
				"TV_FOOTNOTE":                            "footnote",
				"TV_GITHUB_API_POST_REPOSITORY_ENDPOINT": "https://api.github.com/repos/user/posts",
			}

			BeforeEach(func() {
				for k, v := range env {
					os.Setenv(k, v)
				}
			})

			AfterEach(func() {
				for k, _ := range env {
					os.Unsetenv(k)
				}
			})

			It("reads the env vars", func() {
				conf, err := osenv.LoadConfig()
				Expect(err).NotTo(HaveOccurred())
				Expect(conf).To(Equal(&osenv.Config{
					SiteTitle:                       "Textvid",
					Footnote:                        "footnote",
					GitHubAPIPostRepositoryEndpoint: "https://api.github.com/repos/user/posts",
				}))
			})
		})
	})
})
