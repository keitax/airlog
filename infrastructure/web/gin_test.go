package web_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/keitam913/textvid/application/blog"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/keitam913/textvid/domain"
	"github.com/keitam913/textvid/infrastructure/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gin", func() {
	var (
		c        *gomock.Controller
		mBlogSvc *blog.MockService
		gineng   *gin.Engine
		resrec   *httptest.ResponseRecorder
		origDir  string
	)

	BeforeEach(func() {
		gin.SetMode(gin.ReleaseMode)
		var err error
		origDir, err = os.Getwd()
		if err != nil {
			panic(err)
		}
		os.Chdir("../..") // To read static and templates directories
	})

	AfterEach(func() {
		os.Chdir(origDir)
	})

	BeforeEach(func() {
		c = gomock.NewController(GinkgoT())
		mBlogSvc = blog.NewMockService(c)
		gineng = web.SetupGin(
			&web.PostController{
				Service:        mBlogSvc,
				ViewRepository: &web.ViewRepository{},
			},
			&web.WebhookController{
				Service: mBlogSvc,
			},
		)
		resrec = httptest.NewRecorder()
	})

	Describe("GET /health", func() {
		Context("when requested", func() {
			BeforeEach(func() {
				gineng.ServeHTTP(resrec, httptest.NewRequest(http.MethodGet, "/health", nil))
			})

			It("just responses 200 OK", func() {
				Expect(resrec.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(resrec.Body.String()).To(Equal("OK"))
			})
		})
	})

	Describe("GET /:filename", func() {
		Context("when a post is available", func() {
			BeforeEach(func() {
				mBlogSvc.EXPECT().GetByHTMLFilename("20190101-title.html").AnyTimes().Return(
					&domain.Post{
						Filename: "20190101-title.md",
						Title:    "Title",
						Body:     "# Title",
					},
					nil,
				)
			})

			Context("when requested", func() {
				BeforeEach(func() {
					gineng.ServeHTTP(resrec, httptest.NewRequest(http.MethodGet, "/20190101-title.html", nil))
				})

				It("renders a post page", func() {
					Expect(resrec.Result().StatusCode).To(Equal(http.StatusOK))
					Expect(resrec.Body.String()).To(ContainSubstring("<h1>Title</h1>"))
				})
			})
		})
	})

	Describe("POST /webhook", func() {
		Context("when takes a push event", func() {
			AfterEach(func() {
				req := httptest.NewRequest(
					http.MethodPost,
					"/webhook",
					bytes.NewBufferString(`{"before":"<before-commit-id>","after":"<after-commit-id>"}`),
				)
				req.Header.Set("Content-Type", "application/json")
				gineng.ServeHTTP(resrec, req)
			})

			It("registers the changed files", func() {
				mBlogSvc.EXPECT().PushPosts(&domain.PushEvent{
					BeforeCommitID: "<before-commit-id>",
					AfterCommitID:  "<after-commit-id>",
				}).Return(nil)
			})
		})
	})
})
