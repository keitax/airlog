package osenv_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOsenv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Osenv Suite")
}
