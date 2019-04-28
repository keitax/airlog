package ghapi_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/keitax/airlog/internal/domain"
	"github.com/keitax/airlog/internal/infrastructure/ghapi"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Repository", func() {
	var (
		c         *gomock.Controller
		ghRepo    *ghapi.GitHubRepository
		mockGHAPI *httptest.Server
	)

	BeforeEach(func() {
		c = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		c.Finish()
		if mockGHAPI != nil {
			mockGHAPI.Close()
		}
	})

	JustBeforeEach(func() {
		ghRepo = &ghapi.GitHubRepository{
			GitHubAPIPostRepositoryEndpoint: mockGHAPI.URL,
		}
	})

	Describe("ChangedFiles()", func() {
		var (
			givenPushEvent *domain.PushEvent
			gotFs          []*domain.File
			gotErr         error
		)

		JustBeforeEach(func() {
			gotFs, gotErr = ghRepo.ChangedFiles(givenPushEvent)
		})

		Context("when the GitHub compare api responses changed file info", func() {
			BeforeEach(func() {
				mockGHAPI = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					switch req.URL.Path {
					case "/compare/<before-commit-id>...<after-commit-id>":
						rw.WriteHeader(http.StatusOK)
						rw.Write([]byte(fmt.Sprintf(`
{
  "files": [
    {"filename": "20190101-post.md", "raw_url": "%s/20190101-post.md"}
  ]
}
`, mockGHAPI.URL)))
					case "/20190101-post.md":
						rw.WriteHeader(http.StatusOK)
						rw.Write([]byte("post content"))
					default:
						rw.WriteHeader(http.StatusNotFound)
					}
				}))
			})

			Context("given a push event", func() {
				BeforeEach(func() {
					givenPushEvent = &domain.PushEvent{
						BeforeCommitID: "<before-commit-id>",
						AfterCommitID:  "<after-commit-id>",
					}
				})

				It("responses the changed files", func() {
					Expect(gotErr).NotTo(HaveOccurred())
					Expect(gotFs).To(Equal([]*domain.File{
						{Path: "20190101-post.md", Content: "post content"},
					}))
				})
			})
		})
	})
})
