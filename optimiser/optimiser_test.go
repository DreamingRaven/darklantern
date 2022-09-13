package optimiser

import (
	"fmt"
	"testing"
)

func TestFirstMoment(t *testing.T) {
	// standard defaults for adam on first iteration
	alpha := 0.001         // learning
	beta_1 := 0.9          // exponential decay rate first order
	beta_2 := 0.999        // exponential decay rate second order
	epsilon := 1e-8        // decay rate
	it := 1.0              // starts from one to begin with
	gradient := 0.333      // just setting something between 0-1 for testing
	moment_previous := 0.0 // no previous momentum

	fmt.Println(alpha, beta_1, beta_2, epsilon, it, gradient, moment_previous)
	// first order momentum
	m_hat, m_t := Momentum(it, gradient, moment_previous, beta_1)
	if m_hat != 0 {
		t.Fatalf("first order decayed moment () is wrong (expected: )")
	}
	if m_t != 0 {
		t.Fatalf("first order current moment () is wrong (expected: )")
	}
	// second order momentum
	v_hat, v_t := Momentum(it, gradient*gradient, moment_previous, beta_2)
	if v_hat != 0 {
		t.Fatalf("second order decayed moment () is wrong (expected: )")
	}
	if v_t != 0 {
		t.Fatalf("second order current moment () is wrong (expected: )")
	}
}
