package dataset

import (
	"fmt"
	"testing"

	"gitlab.com/deepcypher/darklantern/erray"
)

func TestDatasetInit(t *testing.T) {
	ds := NewExampleDataset[erray.Erray[float64], float64]()
	fmt.Println(ds)
}

func TestDatasetGet(t *testing.T) {
	ds := NewExampleDataset[erray.Erray[float64], float64]()
	fmt.Println(ds)
}

func TestDatasetLen(t *testing.T) {
	ds := NewExampleDataset[erray.Erray[float64], float64]()
	fmt.Println(ds)
}
