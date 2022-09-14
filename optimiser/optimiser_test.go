package optimiser

import (
	"fmt"
	"math"
	"testing"
)

func TestFirstMoment(t *testing.T) {
	// standard defaults for adam on first iteration
	alpha := 0.001           // learning
	beta_1 := 0.9            // exponential decay rate first order
	beta_2 := 0.999          // exponential decay rate second order
	epsilon := 1e-8          // decay rate
	it := 1.0                // starts from one to begin with
	gradient := 0.333        // just setting something between 0-1 for testing
	moment_previous := 0.333 // no previous momentum

	// expected results
	e_m_t := (beta_1 * moment_previous) + (1-beta_1)*gradient
	e_m_hat := e_m_t / (1 - math.Pow(beta_1, it))
	e_v_t := (beta_2 * moment_previous) + (1-beta_2)*gradient*gradient
	e_v_hat := e_v_t / (1 - math.Pow(beta_2, it))

	fmt.Println(alpha, beta_1, beta_2, epsilon, it, gradient, moment_previous)
	// first order momentum (b1 * prevMoment) + (1 - b1) * gradient
	m_hat, m_t := Momentum(it, gradient, moment_previous, beta_1)
	// second order momentum (b2 * prevMoment) + (1 - b2) * gradient^2
	v_hat, v_t := Momentum(it, gradient*gradient, moment_previous, beta_2)
	// checking for errors
	if m_t != e_m_t {
		t.Fatalf("first order current moment (%v) is wrong (expected: %v)", m_t, e_m_t)
	}
	if v_t != e_v_t {
		t.Fatalf("second order current moment (%v) is wrong (expected: %v)", v_t, e_v_t)
	}
	if m_hat != e_m_hat {
		t.Fatalf("first order decayed moment (%v) is wrong (expected: %v)", m_hat, e_m_hat)
	}
	if v_hat != e_v_hat {
		t.Fatalf("second order decayed moment (%v) is wrong (expected: %v)", v_hat, e_v_hat)
	}
}
