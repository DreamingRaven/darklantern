package utils

import (
	"log"
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
