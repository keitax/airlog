package application_test

import (
	"github.com/keitax/airlog/internal/application"
	"github.com/keitax/airlog/internal/domain"
	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("View", func() {
	Describe("GetPostURL()", func() {
		It("gets a URL of a post", func() {
			got := application.GetPostURL(&domain.Post{
				Filename: "20190101-hello-world.md",
			})
			Expect(got).To(Equal("/20190101-hello-world.html"))
		})
	})
})
