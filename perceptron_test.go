package perceptron_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"math/rand"

	"github.com/heldtogether/perceptron"
)

var _ = Describe("Perceptron", func() {

	// Seed the random number generator so all numbers are deterministic
	BeforeEach(func() {
		rand.Seed(1)
	})

	It("generates a set of weights on initialization", func() {
		numInputs := 2

		p := perceptron.New(numInputs)

		// "Random" Weights = [0.6046602879796196 0.9405090880450124]
		Expect(p.Weights).Should(HaveLen(numInputs))
		Expect(p.Weights).Should(BeEquivalentTo([]float64{0.6046602879796196, 0.9405090880450124}))
	})

	It("generates a random threshold on initialization", func() {

		p := perceptron.New(2)

		// "Random" Threshold = 0.6645600532184904
		Expect(p.Threshold).Should(BeNumerically("==", 0.6645600532184904))
	})

	It("can calculate the activity given an input", func() {

		p := perceptron.New(2)
		a := p.Forward([]int{1, 0})

		// Calculated activity = 0.6046602879796196
		Expect(a).Should(BeNumerically("==", 0.6046602879796196))
	})

})
