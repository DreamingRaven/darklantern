package optimiser

import (
	dt "gitlab.com/deepcypher/darklantern/darktype"
	dutl "gitlab.com/deepcypher/darklantern/utils"
)

// aliasing paramMap so we dont have to constantly type it out
type paramMap[L dt.LattigoCompat] map[string]L

// aliasing paramMap so we dont have to constantly type it out
type momentMap[L dt.LattigoCompat] map[string]orderedMoments[L]

// individual history value struct
type orderedMoments[L dt.LattigoCompat] struct {
	m_t L
	v_t L
}

// Adam struct holding necessary data and history in ordert to calculate new parameters
type Adam[L dt.LattigoCompat] struct {
	Alpha   L
	Beta_1  L
	Beta_2  L
	Epsilon L
	History []momentMap[L]
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
	nextHistory := make(map[string]orderedMoments[L])
	for parmName, parmVal := range parms {
		var previousMoments orderedMoments[L]
		var iterations L = 1
		var gradient L = grads[parmName]
		// scans history for presence of value and counts
		// the final value will be residual to the next few statements
		for i := 0; i < len(a.History); i++ {
			if pM, ok := a.History[i][parmName]; ok {
				iterations += 1
				previousMoments = pM
			}
		}

		// first order calc
		m_hat, m_t := Momentum(iterations, gradient, previousMoments.m_t, a.Beta_1)
		// second order calc
		v_hat, v_t := Momentum(iterations, gradient*gradient, previousMoments.v_t, a.Beta_2)
		// resolve to mutate parameter
		nextParms[parmName] = parmVal - ((a.Alpha * m_hat) / (dutl.Sqrt(v_hat) + a.Epsilon))
		// retention save parameters moment
		nextHistory[parmName] = orderedMoments[L]{m_t: m_t, v_t: v_t}
	}
	a.History = append(a.History, nextHistory)
	return nextParms, nil
}

// RMSProp or second order moment calculation
func RMSProp() {

}
