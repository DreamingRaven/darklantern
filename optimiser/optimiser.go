package optimiser

import (
	dt "gitlab.com/deepcypher/darklantern/darktype"
)

type Optimiser[L dt.LattigoCompat] interface {
	Optimise(grads, parms paramMap[L]) (paramMap[L], error)
}

// Momentum calculates first or second order momentums given current gradient, desired order, and history of gradients.
// it is the number of iterations of optimisation this parameter has experienced.
// gradient is the gradient with respect to the parameter being optimised.
// moment is the last known moment of this parameter i.e the previous m_t.
// beta is a stand in for beta_1 and beta_2 and represents the exponential decay rate.
// Momentum returns a tuple of (m_hat, m_t) where m_hat is the corrected momentum
// and m_t is the non corrected momentum.
// It should be noted gradients passed in towards calculating second order moments must be squared before being passed to this function.
func Momentum[L dt.LattigoCompat](it, gradient, moment, beta L) (L, L) {
	// first order moment
	// m_t = \beta_1 * m_<t-1> + (1 - \beta_1) * g_t
	// \hat{m}_t = \frac{m_t}{1 - b_1^t}
	// second order moment
	// v_t = \beta_2 * v_<t-1> + (1 - \beta_2) * g_t^2
	// \hat{v}_t = \frac{v_t}{1 - b_2^t}
	var t L = 1
	// using for loop instead of math power since using generics not float64
	for i := 0; i < int(it); i++ {
		t = t * beta
	}
	m_t := (beta * moment) + ((1 - beta) * gradient)
	m_hat := m_t / (1 - t)
	return m_hat, m_t
}
