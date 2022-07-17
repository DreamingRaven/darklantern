package dataset

import (
	"fmt"
	"testing"

	"gitlab.com/deepcypher/darklantern/erray"
)

// helper function to generate a dud example slice of slices
func sliceOfSlices[T LattigoCompat](examples, features int) [][]T {
	slice_slice := make([][]T, examples)
	for i := 0; i < examples; i++ {
		slice_slice[i] = make([]T, features)
		for j := 0; j < features; j++ {
			slice_slice[i][j] = T(i * j)
		}
	}
	return slice_slice
}

func TestSliceOfSlices(t *testing.T) {
	ss := sliceOfSlices[float64](10, 20)
	for i := 0; i < len(ss); i++ {
		if ss[i][1] != float64(i) {
			t.Fatal(fmt.Sprintf("slice of slices failed as %v != %v", ss[i], float64(i)))
		}
	}
}

func TestDatasetInit(t *testing.T) {
	sliceOfSlices[float64](10, 20)
	ds := NewSimpleDataset[erray.Erray[float64], float64]()
	fmt.Println(ds)
}

func TestDatasetGet(t *testing.T) {
	sliceOfSlices[float64](10, 20)
	ds := NewSimpleDataset[erray.Erray[float64], float64]()
	fmt.Println(ds)
}

func TestDatasetLen(t *testing.T) {
	sliceOfSlices[float64](10, 20)
	ds := NewSimpleDataset[erray.Erray[float64], float64]()
	fmt.Println(ds)
}
