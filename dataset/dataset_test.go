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
