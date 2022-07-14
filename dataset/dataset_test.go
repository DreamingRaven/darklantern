package dataset

import (
	"fmt"
	"testing"
)

func TestDatasetInit(t *testing.T) {
	ds := NewExampleDataset()
	fmt.Println(ds)
}
