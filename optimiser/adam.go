package optimiser

import (
	"fmt"

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
func (a *Adam[L]) Optimise(grads, parms paramMap[L]) (paramMap[L], error) {

	nextParms := make(map[string]L)
	for key, val := range parms {
		var lastMoment L = 0
		iterations := 1
		// scans history for presence of value and counts
		// the final value will be residual to the next few statements
		for i := 0; i < len(a.History); i++ {
			// https://stackoverflow.com/a/2050629
			if pV, ok := a.History[i][key]; ok {
				iterations += 1
				lastMoment = pV
			}
		}
		fmt.Println(key, val, lastMoment, iterations)

		// // first order
		// Momentum(it, gradient, moment, a.Beta_1))
		// // second order
		// Momentum(it, gradient*gradient, moment, beta)
		// // mutation
	}
	return nextParms, nil
}

// RMSProp or second order moment calculation
func RMSProp() {

}
