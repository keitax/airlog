package domain_test

import (
	"github.com/golang/mock/gomock"
	"github.com/keitax/airlog/internal/domain"
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
})
