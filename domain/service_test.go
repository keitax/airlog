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
		mrepo   *domain.MockPostRepository
	)

	BeforeEach(func() {
		c = gomock.NewController(GinkgoT())
		mrepo = domain.NewMockPostRepository(c)
		service = &domain.PostServiceImpl{
			Repository: mrepo,
		}
	})

	AfterEach(func() {
		c.Finish()
	})

	Describe("GetByHTMLFilename()", func() {
		Context("when a post is saved", func() {
			BeforeEach(func() {
				mrepo.EXPECT().Filename("20190101-post.md").AnyTimes().Return(&domain.Post{
					Title: "First Post",
				}, nil)
			})

			It("search the post by a html page name", func() {
				post, err := service.GetByHTMLFilename("20190101-post.html")
				Expect(err).NotTo(HaveOccurred())
				Expect(post).To(Equal(&domain.Post{
					Title: "First Post",
				}))
			})
		})
	})

	Describe("Recent()", func() {
		Context("when posts are available", func() {
			BeforeEach(func() {
				mrepo.EXPECT().All().AnyTimes().Return([]*domain.Post{
					{Filename: "20190102-bar.md"},
					{Filename: "20190101-foo.md"},
				}, nil)
			})

			It("gets the posts", func() {
				got, err := service.Recent()
				Expect(err).NotTo(HaveOccurred())
				Expect(got).To(Equal([]*domain.Post{
					{Filename: "20190102-bar.md"},
					{Filename: "20190101-foo.md"},
				}))
			})
		})
	})

	Describe("RegisterPost()", func() {
		Context("when given a post file", func() {
			AfterEach(func() {
				service.RegisterPost("20190101-post.md", `---
labels: [label-a, label-b]
---

# Title

content
`)
			})

			It("puts a post from the file", func() {
				mrepo.EXPECT().Put(&domain.Post{
					Filename:  "20190101-post.md",
					Timestamp: GetTimestamp("2019-01-01T00:00:00Z"),
					Title:     "Title",
					Body: `
content
`,
					Labels: []string{"label-a", "label-b"},
				})
			})
		})
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
