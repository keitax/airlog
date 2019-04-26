package osenv_test

import (
	"github.com/keitax/airlog/internal/application"
	"github.com/keitax/airlog/internal/infrastructure/osenv"
	. "github.com/onsi/gomega"
	"os"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Config", func() {
	Describe("LoadConfig()", func() {
		Context("with valid environment vars", func() {
			env := map[string]string{
				"PORT":          "3000",
				"AL_SITE_TITLE": "Airlog",
				"AL_FOOTNOTE":   "footnote",
				"AL_BLOG_DSN":   "root@tcp(localhost:3306)/blog",
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
				Expect(conf).To(Equal(&application.Config{
					Port:      "3000",
					SiteTitle: "Airlog",
					Footnote:  "footnote",
					BlogDSN:   "root@tcp(localhost:3306)/blog",
				}))
			})
		})
	})
})
