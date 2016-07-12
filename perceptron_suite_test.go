package perceptron_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPerceptron(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Perceptron Suite")
}
