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
		a := p.Activity([]int{1, 0})

		// Calculated activity = 0.6046602879796196
		Expect(a).Should(BeNumerically("==", 0.6046602879796196))
	})

	It("adjusts weights and threshold when wrong", func() {

		p := perceptron.New(2)
		p.Train([]int{1, 0}, 1)

		Expect(p.Threshold).Should(BeNumerically("==", -0.33543994678150957))
		Expect(p.Weights[0]).Should(BeNumerically("==", 1.6046602879796197))
		Expect(p.Weights[1]).Should(BeNumerically("==", 0.9405090880450124))
	})

	It("activates when activity is greater than the threshold", func() {
		p := perceptron.New(2)
		p.Threshold = 0

		// Calculated activity = 0.6046602879796196 > 0 so activates
		Expect(p.Activates([]int{1, 0})).Should(BeNumerically("==", 1))
	})

	It("doesn't activates when activity is greater than the threshold", func() {
		p := perceptron.New(2)
		p.Threshold = 1

		// Calculated activity = 0.6046602879796196 < 1 so doesn't activate
		Expect(p.Activates([]int{1, 0})).Should(BeNumerically("==", 0))
	})

})
