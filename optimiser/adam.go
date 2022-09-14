package optimiser

import (
	dt "gitlab.com/deepcypher/darklantern/darktype"
)

// aliasing paramMap so we dont have to constantly type it out
type paramMap[L dt.LattigoCompat] map[string]L

// Adam struct holding necessary data and history in ordert to calculate new parameters
type Adam[L dt.LattigoCompat] struct {
	Alpha   float64
	Beta_1  float64
	Beta_2  float64
	Epsilon float64
	History []paramMap[L]
}

// NewDefaultAdam creates a default adam struct for you based on https://arxiv.org/abs/1412.6980
func NewDefaultAdam[L dt.LattigoCompat]() *Adam[L] {
	return &Adam[L]{
		Alpha:   0.001,
		Beta_1:  0.9,
		Beta_2:  0.999,
		Epsilon: 1e-8,
	}
}

// NewDefaultAdamOptimiser use default adam struct and wrap it in optimiser abstraction
func NewDefaultAdamOptimiser[L dt.LattigoCompat]() Optimiser[L] {
	return NewDefaultAdam[L]()
}

// Optimise use the ADAM algorithm first a second order moments to calculate the new parameters
func (a *Adam[L]) Optimise(grads, parms paramMap[L]) error {

	// // first order
	// Momentum(it, gradient, moment, beta)
	// // second order
	// Momentum(it, gradient*gradient, moment, beta)
	// // mutation
	return nil
}

// RMSProp or second order moment calculation
func RMSProp() {

}
