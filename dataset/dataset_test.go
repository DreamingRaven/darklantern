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

func TestSliceOfSlices(t *testing.T) {
	ss := sliceOfSlices[float64](10, 20)
	for i := 0; i < len(*ss); i++ {
		if (*ss)[i][1] != float64(i) {
			t.Fatal(fmt.Sprintf("slice of slices failed as %v != %v", (*ss)[i], float64(i)))
		}
	}
}

func sliceOfErrays[T LattigoCompat](sos *[][]T) *[]erray.Erray[T] {
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

func TestSliceOfErrays(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := sliceOfErrays(sos)
	for i := 0; i < len(*soe); i++ {
		if (*soe)[i] == nil {
			t.Fatal("Slice of Errays giving back empty errays pointers but should be populated")
		}
	}

}

func TestDatasetInit(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := sliceOfErrays(sos)
	NewSimpleDataset[erray.Erray[float64], float64](soe)
}

func TestDatasetGet(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := sliceOfErrays(sos)
	ds := NewSimpleDataset[erray.Erray[float64], float64](soe)
	e, err := ds.Get(len(*sos) - 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*e)
}

func TestDatasetLen(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := sliceOfErrays(sos)
	ds := NewSimpleDataset[erray.Erray[float64], float64](soe)
	dsLength, err := ds.Length()
	if err != nil {
		t.Fatal(err)
	}
	if dsLength != len(*sos) {
		t.Fatal("Number of examples in slice of slices does not equal number of examples in dataset")
	}
}
