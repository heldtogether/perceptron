package perceptron

import (
	"fmt"
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

func (p *Perceptron) Forward(input []int) float64 {
	if len(p.Weights) != len(input) {
		panic(fmt.Sprintf("Input length should be %d", len(p.Weights)))
	}

	var sum float64

	for i := 0; i < len(p.Weights); i++ {
		sum += p.Weights[i] * float64(input[i])
	}

	return sum
}

func (p *Perceptron) Activates(input []int) int {
	if len(p.Weights) != len(input) {
		panic(fmt.Sprintf("Input length should be %d", len(p.Weights)))
	}

	if p.Forward(input) >= p.Threshold {
		return 1
	} else {
		return 0
	}
}
