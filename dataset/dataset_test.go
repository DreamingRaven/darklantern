package dataset

import (
	"fmt"
	"testing"
)

func TestDatasetInit(t *testing.T) {
	ds := NewExampleDataset()
	fmt.Println(ds)
}

func TestDatasetGet(t *testing.T) {
	ds := NewExampleDataset()
	fmt.Println(ds)
}

func TestDatasetLen(t *testing.T) {
	ds := NewExampleDataset()
	fmt.Println(ds)
}
