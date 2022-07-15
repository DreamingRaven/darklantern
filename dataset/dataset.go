// dataset: similar to the objective of a pytorch dataset, a darklantern dataset must be
// slice like indexable except with a specific composition rather than inheritance.
package dataset

import (
	"encoding/json"
	"fmt"

	"gitlab.com/deepcypher/darklantern/erray"
)

// Dataset abstraction and definition of avaliable methods
type Dataset[D DatasetCompatible[L], L LattigoCompatible] interface {
	Get(i int) (D, error)
	Length() (int, error)
	ToJSON() ([]byte, error)
	FromJSON(bytes []byte) error
}

// LattigoCompatible base data type of dataset compatible objects
type LattigoCompatible interface {
	~float64 // | ~complex128
}

// DatasetCompatible base data type of dataset objects directly
type DatasetCompatible[L LattigoCompatible] interface {
	erray.Erray[L]
}

// exampleDataset the simplest dataset to show as an example
type exampleDataset[D DatasetCompatible[L], L LattigoCompatible] struct {
	data []D
}

// NewExampleDataset initialises a new dataset object with some testable example data
func NewExampleDataset[D DatasetCompatible[L], L LattigoCompatible]() Dataset[D, L] {
	return &exampleDataset[D, L]{}
}

// Get a specific example from the dataset by index or error if impossible
func (ds *exampleDataset[D, L]) Get(i int) (D, error) {
	return ds.data[i], nil
}

// Length of the dataset so controlling code does not exceed the bounds of this
func (ds *exampleDataset[D, L]) Length() (int, error) {
	fmt.Println("Tr")
	return 0, nil
}

// ToJSON convert dataset internal struct to json bytes
func (ds *exampleDataset[D, L]) ToJSON() ([]byte, error) {
	marshalled, err := json.Marshal(ds)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}

// FromJSON convert json bytes back into original struct
func (ds *exampleDataset[D, L]) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, ds)
	if err != nil {
		return err
	}
	return nil
}
