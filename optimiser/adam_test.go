package optimiser

import (
	"testing"
)

func TestADAMEmpty(t *testing.T) {
	a := NewDefaultAdamOptimiser[float64]()
	grads := paramMap[float64]{}
	parms := paramMap[float64]{}
	a.Optimise(grads, parms)
}

func TestADAM(t *testing.T) {
	a := NewDefaultAdamOptimiser[float64]()
	grads := paramMap[float64]{
		"someParam": 0.33,
	}
	parms := paramMap[float64]{
		"someParam": 0.5,
	}
	a.Optimise(grads, parms)
}
