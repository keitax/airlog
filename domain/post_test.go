package domain_test

import (
	"time"
	"github.com/keitam913/textvid/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Post", func() {
	Describe("Timestamp()", func() {
		It("parses a timestamp from a file name", func() {
			p := &domain.Post{Filename: "20190101-foo.md"}
			got := p.Timestamp()
			Expect(got.Year()).To(Equal(2019))
			Expect(got.Month()).To(Equal(time.January))
			Expect(got.Day()).To(Equal(1))
		})
	})
})
