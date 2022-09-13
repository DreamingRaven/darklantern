package dataset

import (
	"github.com/ldsec/lattigo/v2/ckks"
	dt "gitlab.com/deepcypher/darklantern/darktype"
	"gitlab.com/deepcypher/darklantern/erray"
)

// helper function to generate a dud example slice of slices
func ExampleSliceOfSlices[T dt.LattigoCompat](examples, features int) *[][]T {
	slice_slice := make([][]T, examples)
	for i := 0; i < examples; i++ {
		slice_slice[i] = make([]T, features)
		for j := 0; j < features; j++ {
			slice_slice[i][j] = T(i * j)
		}
	}
	return &slice_slice
}

// helper function to turn slice of slices into slice of errays
func SliceOfErrays[T dt.LattigoCompat](sos *[][]T) *[]erray.Erray[T] {
	soe := make([]erray.Erray[T], len(*sos))
	e := erray.NewCKKSErray[T]()
	parms, _ := ckks.NewParametersFromLiteral(ckks.PN12QP109)
	e.SetParams(&parms)
	for i := 0; i < len(*sos); i++ {
		// soe[i] =
		bud := e.Bud()
		bud.SetData(&(*sos)[i])
		soe[i] = bud
	}
	return &soe
}
