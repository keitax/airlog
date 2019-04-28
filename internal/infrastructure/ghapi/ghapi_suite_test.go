package ghapi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGhapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ghapi Suite")
}
