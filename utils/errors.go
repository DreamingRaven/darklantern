package utils

import (
	"log"

	dt "gitlab.com/deepcypher/darklantern/darktype"
)

func RaiseErrors(argv ...any) []any {
	if argv[len(argv)-1] != nil {
		log.Fatal(argv[len(argv)-1])
	}
	return argv[:len(argv)-1]
}

func RaiseError[T any](arg T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return arg
}

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
