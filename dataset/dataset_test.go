package dataset

import (
	"fmt"
	"testing"

	"github.com/ldsec/lattigo/v2/ckks"
	"gitlab.com/deepcypher/darklantern/erray"
)

// helper function to generate a dud example slice of slices
func sliceOfSlices[T LattigoCompat](examples, features int) *[][]T {
	slice_slice := make([][]T, examples)
	for i := 0; i < examples; i++ {
		slice_slice[i] = make([]T, features)
		for j := 0; j < features; j++ {
			slice_slice[i][j] = T(i * j)
		}
	}
	return &slice_slice
}

func sliceOfErrays[T LattigoCompat](sos *[][]T) *[]erray.Erray[T] {
	soe := make([]erray.Erray[T], len(*sos))
	o := erray.NewCKKSErray[float64]()
	parms, _ := ckks.NewParametersFromLiteral(ckks.PN12QP109)
	o.SetParams(&parms)
	for i := 0; i < len(*sos); i++ {
		// soe[i] =
		fmt.Println((*sos)[i])
	}
	return &soe
}

func TestSliceOfSlices(t *testing.T) {
	ss := sliceOfSlices[float64](10, 20)
	for i := 0; i < len(*ss); i++ {
		if (*ss)[i][1] != float64(i) {
			t.Fatal(fmt.Sprintf("slice of slices failed as %v != %v", (*ss)[i], float64(i)))
		}
	}
}

func TestDatasetInit(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := sliceOfErrays(sos)
	ds := NewSimpleDataset[erray.Erray[float64], float64](soe)
	fmt.Println(ds)
}

func TestDatasetGet(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := sliceOfErrays(sos)
	ds := NewSimpleDataset[erray.Erray[float64], float64](soe)
	fmt.Println(ds)
}

func TestDatasetLen(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := sliceOfErrays(sos)
	ds := NewSimpleDataset[erray.Erray[float64], float64](soe)
	fmt.Println(ds)
}
