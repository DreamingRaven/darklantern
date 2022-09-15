package utils

import (
	"testing"
)

func FuzzSqrt(f *testing.F) {
	f.Add(100.0, 10.0)
	f.Add(10.0, 3.1622776601683795)
	f.Add(2.0, 1.4142135623730951)

	f.Fuzz(func(t *testing.T, num, expect float64) {
		out := Sqrt(num)
		if out != expect {
			t.Fatalf("Sqrt(%v) was expected to be %v (%T) got %v (%T)", num, expect, expect, out, out)
		}
	})
}

func FuzzPow(f *testing.F) {
	f.Add(2.0, 2.0, 4.0)
	f.Add(-2.0, 2.0, 4.0)
	f.Add(-10.0, 2.0, 100.0)

	f.Fuzz(func(t *testing.T, num, pow, expect float64) {
		out := Pow(num, pow)
		if out != expect {
			t.Fatalf("Pow(%v, %v) was expected to be %v (%T) got %v (%T)", num, pow, expect, expect, out, out)
		}
	})
}
