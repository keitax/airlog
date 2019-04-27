package domain_test

import (
	"github.com/keitax/airlog/internal/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Markdown", func() {
	Describe("ExtractFrontMatter()", func() {
		var (
			content string
			fm      map[string]interface{}
			body    string
		)

		JustBeforeEach(func() {
			fm, body = domain.ExtractFrontMatter(content)
		})

		Context("when the content has a valid front matter", func() {
			BeforeEach(func() {
				content = `---
labels: [a, b]
---

hello world`
			})

			It("extracts the front matter as yaml", func() {
				Expect(fm).To(Equal(map[string]interface{}{
					"labels": []interface{}{"a", "b"},
				}))
				Expect(body).To(Equal("hello world"))
			})
		})

		Context("when the content doesn't have front matter", func() {
			BeforeEach(func() {
				content = `hello world
`
			})

			It("makes an empty map", func() {
				Expect(fm).To(Equal(map[string]interface{}{}))
				Expect(body).To(Equal(`hello world
`))
			})
		})
	})
})
