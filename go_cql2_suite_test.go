package cql2_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoCql2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoCql2 Suite")
}
