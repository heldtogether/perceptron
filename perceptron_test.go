package perceptron_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/heldtogether/perceptron"
)

var _ = Describe("Perceptron", func() {

	It("generates a set of weights", func() {
		numInputs := 2

		p := perceptron.New(numInputs)

		Expect(p.Weights).Should(HaveLen(numInputs))
	})

})
