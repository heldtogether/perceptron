// Package perceptron provides the ability to create, train and use
// a linear classifier which takes a number of inputs and returns
// a single output.
//
// A Perceptron has a set of input weights and an activation threshold.
//
// For a given set of inputs i1...in, the activity is calculated as
//     sum(i * w)
// where w1...wn are weights for the inputs. The output is 1 if
// the activity > threshold, else 0.
package perceptron

import (
	"fmt"
	"math/rand"
)

type Perceptron struct {
	Weights   []float64
	Threshold float64
}

// New creates a new Perceptron with the specified number of inputs.
// It assigns random weights to each of the inputs and generates a
// random activation threshold for the Perceptron.
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

// Activity calculates the activity on the Perceptron for a given
// input. It doesn't calculate if the Perceptron actually fires,
// see "Activates" instead.
func (p *Perceptron) Activity(input []int) float64 {
	if len(p.Weights) != len(input) {
		panic(fmt.Sprintf("Input length should be %d", len(p.Weights)))
	}

	var sum float64

	for i := 0; i < len(p.Weights); i++ {
		sum += p.Weights[i] * float64(input[i])
	}

	return sum
}

// Activates checks if the Perceptron activates for the given input.
// It returns 0 or 1 depending on the input and how the perceptron
// has been trained.
func (p *Perceptron) Activates(input []int) int {
	if len(p.Weights) != len(input) {
		panic(fmt.Sprintf("Input length should be %d", len(p.Weights)))
	}

	if p.Activity(input) >= p.Threshold {
		return 1
	} else {
		return 0
	}
}

// Train takes a given input and an expected output. If the Perceptron
// doesn't activate, it adjusts the activation threshold and input
// weights based on the delta between the expected and calculated
// output.
func (p *Perceptron) Train(input []int, expectedOutput int) {
	if len(p.Weights) != len(input) {
		panic(fmt.Sprintf("Input length should be %d", len(p.Weights)))
	}

	calculatedOutput := p.Activates(input)

	p.Threshold = p.Threshold - (float64(expectedOutput) - float64(calculatedOutput))

	for i := 0; i < len(p.Weights); i++ {
		p.Weights[i] = p.Weights[i] + (float64(expectedOutput)-float64(calculatedOutput))*float64(input[i])
	}
}
