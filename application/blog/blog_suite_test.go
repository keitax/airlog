package blog_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBlog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Blog Suite")
}
