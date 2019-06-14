package application_test

import (
	"github.com/keitam913/airlog/internal/application"
	"github.com/keitam913/airlog/internal/domain"
	. "github.com/onsi/gomega"
	"time"

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

	Describe("ShowDate()", func() {
		It("gets a date text of a timestamp", func() {
			t, err := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
			if err != nil {
				panic(err)
			}
			got := application.ShowDate(t)
			Expect(got).To(Equal("2019-01-01"))
		})
	})
})
