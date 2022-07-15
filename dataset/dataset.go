// dataset: similar to the objective of a pytorch dataset, a darklantern dataset must be
// slice like indexable except with a specific composition rather than inheritance.
package dataset

import (
	"encoding/json"

	"gitlab.com/deepcypher/darklantern/erray"
)

// LattigoCompat base data type of dataset compatible objects
type LattigoCompat interface {
	~float64 // | ~complex128
}

// DatasetCompat base data type of dataset objects directly
type DatasetCompat[L LattigoCompat] interface {
	erray.Erray[L]
}

// Dataset abstraction and definition of avaliable methods
type Dataset[D DatasetCompat[L], L LattigoCompat] interface {
	Get(i int) (D, error)
	Length() (int, error)
	ToJSON() ([]byte, error)
	FromJSON(bytes []byte) error
}

// exampleDataset the simplest dataset to show as an example
type exampleDataset[D DatasetCompat[L], L LattigoCompat] struct {
	Data []D `json:"data"` // slice of individual examples
}

// NewExampleDataset initialises a new dataset object with some testable example data
func NewExampleDataset[D DatasetCompat[L], L LattigoCompat]() Dataset[D, L] {
	return &exampleDataset[D, L]{}
}

// Get a specific example from the dataset by index or error if impossible
func (ds *exampleDataset[D, L]) Get(i int) (D, error) {
	return ds.Data[i], nil
}

// Length of the dataset so controlling code does not exceed the bounds of this
func (ds *exampleDataset[D, L]) Length() (int, error) {
	return len(ds.Data), nil
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
