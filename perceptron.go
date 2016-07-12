package perceptron

import (
	"math/rand"
)

type Perceptron struct {
	Weights   []float64
	Threshold float64
}

func New(numInputs int) *Perceptron {
	p := new(Perceptron)
	p.Weights = initializeWeights(numInputs)
	p.Threshold = rand.Float64()
	return p
}

func initializeWeights(numInputs int) []float64 {
	w := make([]float64, numInputs)
	for i := 0; i < numInputs; i++ {
		w[i] = rand.Float64()
	}
	return w
}
