package application_test

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/keitax/airlog/internal/application"
	"github.com/keitax/airlog/internal/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("application", func() {
	var (
		c       *gomock.Controller
		mpsvc   *domain.MockPostService
		gineng  *gin.Engine
		resrec  *httptest.ResponseRecorder
		origDir string
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
		mpsvc = domain.NewMockPostService(c)
		gineng = application.SetupGin(&application.PostController{
			Service:        mpsvc,
			ViewRepository: &application.ViewRepository{},
		})
		resrec = httptest.NewRecorder()
	})

	Describe("GET /:filename", func() {
		Context("when a post is available", func() {
			BeforeEach(func() {
				mpsvc.EXPECT().GetByHTMLFilename("20190101-title.html").AnyTimes().Return(
					&domain.Post{
						Filename: "20190101-title.md",
						Title:    "Title",
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
					Expect(resrec.Body.String()).To(ContainSubstring("<h1>airlog</h1>"))
				})
			})
		})
	})
})
