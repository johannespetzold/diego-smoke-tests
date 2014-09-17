package diego_smoke_tests_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDiegoSmokeTests(t *testing.T) {
	SetDefaultEventuallyTimeout(time.Minute)
	SetDefaultEventuallyPollingInterval(time.Second)

	RegisterFailHandler(Fail)
	RunSpecs(t, "DiegoSmokeTests Suite")
}
