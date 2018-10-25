package ohdear_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOhdearSdk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OhdearSdk Suite")
}
