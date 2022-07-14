// dataset: similar to the objective of a pytorch dataset, a darklantern dataset must be
// slice like indexable except with a specific composition rather than inheritance.
package dataset

import (
	"fmt"
)

type dataset struct {
}

type Dataset interface {
}

// NewExampleDataset initialises a new dataset object with some testable example data
func NewExampleDataset() Dataset {
	return &dataset{}
}

func (ds *dataset) train() error {
	fmt.Println("Tr")
	return nil
}
