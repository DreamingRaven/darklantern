// dataset: similar to the objective of a pytorch dataset, a darklantern dataset must be
// slice like indexable except with a specific composition rather than inheritance.
package dataset

import (
	"encoding/json"
	"fmt"

	"gitlab.com/deepcypher/darklantern/erray"
)

// Dataset abstraction and definition of avaliable methods
type Dataset interface {
	Get(i int) error
	Length() (int, error)
	ToJSON() ([]byte, error)
	FromJSON(bytes []byte) error
}

// LattigoCompatible base data type of dataset compatible objects
type LattigoCompatible interface {
	~float64 // | ~complex128
}

// DatasetCompatible base data type of dataset objects directly
type DatasetCompatible[T LattigoCompatible] interface {
	~[]T | []erray.Erray[T]
}

// exampleDataset the simplest dataset to show as an example
type exampleDataset struct {
	data []float32
}

// NewExampleDataset initialises a new dataset object with some testable example data
func NewExampleDataset() Dataset {
	return &exampleDataset{}
}

// Get a specific example from the dataset by index or error if impossible
func (ds *exampleDataset) Get(i int) error {
	fmt.Println("Tr")
	return nil
}

// Length of the dataset so controlling code does not exceed the bounds of this
func (ds *exampleDataset) Length() (int, error) {
	fmt.Println("Tr")
	return 0, nil
}

// ToJSON convert dataset internal struct to json bytes
func (ds *exampleDataset) ToJSON() ([]byte, error) {
	marshalled, err := json.Marshal(ds)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}

// FromJSON convert json bytes back into original struct
func (ds *exampleDataset) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, ds)
	if err != nil {
		return err
	}
	return nil
}
