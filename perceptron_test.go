package perceptron_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/heldtogether/perceptron"
)

var _ = Describe("Perceptron", func() {

	It("generates a set of weights on initialization", func() {
		numInputs := 2

		p := perceptron.New(numInputs)

		Expect(p.Weights).Should(HaveLen(numInputs))
	})

	It("generates a random threshold on initialization", func() {

		p := perceptron.New(2)

		Expect(p.Threshold).Should(BeNumerically(">", 0.0))
	})

})
