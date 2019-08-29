package web_test

import (
	"time"

	"github.com/keitam913/textvid/domain"
	"github.com/keitam913/textvid/infrastructure/web"
	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("View", func() {
	Describe("GetPostURL()", func() {
		It("gets a URL of a post", func() {
			got := web.GetPostURL(&domain.Post{
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
			got := web.ShowDate(t)
			Expect(got).To(Equal("2019-01-01"))
		})
	})
})
