package utils

import (
	dt "gitlab.com/deepcypher/darklantern/darktype"
)

// Sqrt a custom lattigo compatible datatype
func Sqrt[L dt.LattigoCompat](n L) L {
	return n
}

// Pow calculate some exponential power of a custom lattigo datatype
func Pow[L dt.LattigoCompat](num, power L) L {
	var t L = 1.0
	for i := 0; i < int(power); i++ {
		t = t * num
	}
	return t
}
