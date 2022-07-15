// dataset: similar to the objective of a pytorch dataset, a darklantern dataset must be
// slice like indexable except with a specific composition rather than inheritance.
package dataset

import (
	"encoding/json"
	"fmt"

	"gitlab.com/deepcypher/darklantern/erray"
)

type LattigoCompatible interface {
	~float64 // | ~complex128
}

type DatasetCompatible[T LattigoCompatible] interface {
	~[]T | []erray.Erray[T]
}

type Dataset interface {
	Get(i int) error
	Length() (int, error)
	ToJSON() ([]byte, error)
	FromJSON(bytes []byte) error
}

// exampleDataset the simplest dataset to show as an example
type exampleDataset struct {
	data []float32
}

// NewExampleDataset initialises a new dataset object with some testable example data
func NewExampleDataset() Dataset {
	return &exampleDataset{}
}

func (ds *exampleDataset) Get(i int) error {
	fmt.Println("Tr")
	return nil
}

func (ds *exampleDataset) Length() (int, error) {
	fmt.Println("Tr")
	return 0, nil
}

func (ds *exampleDataset) ToJSON() ([]byte, error) {
	marshalled, err := json.Marshal(ds)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}

func (ds *exampleDataset) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, ds)
	if err != nil {
		return err
	}
	return nil
}
