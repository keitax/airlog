package domain_test

import (
	"time"

	"github.com/golang/mock/gomock"
	"github.com/keitam913/airlog/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	var (
		c       *gomock.Controller
		service *domain.PostServiceImpl
	)

	BeforeEach(func() {
		c = gomock.NewController(GinkgoT())
		service = &domain.PostServiceImpl{}
	})

	AfterEach(func() {
		c.Finish()
	})

	Describe("ConvertToPost()", func() {
		Context("when given a post file", func() {
			It("gets a post from a file", func() {
				got := service.ConvertToPost("20190101-post.md", `---
labels: [label-a, label-b]
---

# Title

content
`)
				Expect(got).To(Equal(&domain.Post{
					Filename:  "20190101-post.md",
					Timestamp: GetTimestamp("2019-01-01T00:00:00Z"),
					Title:     "Title",
					Body: `
content
`,
					Labels: []string{"label-a", "label-b"},
				}))
			})
		})
	})
})

func GetTimestamp(text string) time.Time {
	t, err := time.Parse(time.RFC3339, text)
	if err != nil {
		panic(err)
	}
	return t
}
