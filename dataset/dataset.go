// dataset: similar to the objective of a pytorch dataset, a darklantern dataset must be
// slice like indexable except with a specific composition rather than inheritance.
package dataset

import (
	"encoding/json"

	dt "gitlab.com/deepcypher/darklantern/darktype"
	"gitlab.com/deepcypher/darklantern/erray"
)

// DatasetCompat base data type of dataset objects directly
type DatasetCompat[L dt.LattigoCompat] interface {
	erray.Erray[L]
}

// Dataset abstraction and definition of avaliable methods
type Dataset[D DatasetCompat[L], L dt.LattigoCompat] interface {
	Get(i int) (*D, error)
	Length() (int, error)
	ToJSON() ([]byte, error)
	FromJSON(bytes []byte) error
}

// simpleDataset the simplest dataset to show as an example
type simpleDataset[D DatasetCompat[L], L dt.LattigoCompat] struct {
	Data *[]D `json:"data"` // slice of individual examples
}

// NewSimpleDataset initialises a new dataset object with some testable example data
func NewSimpleDataset[D DatasetCompat[L], L dt.LattigoCompat](data *[]D) Dataset[D, L] {
	return &simpleDataset[D, L]{Data: data}
}

// Get a specific example from the dataset by index or error if impossible
func (ds *simpleDataset[D, L]) Get(i int) (*D, error) {
	return &(*(ds.Data))[i], nil
}

// Length of the dataset so controlling code does not exceed the bounds of this
func (ds *simpleDataset[D, L]) Length() (int, error) {
	return len(*(ds.Data)), nil
}

// ToJSON convert dataset internal struct to json bytes
func (ds *simpleDataset[D, L]) ToJSON() ([]byte, error) {
	marshalled, err := json.Marshal(ds)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}

// FromJSON convert json bytes back into original struct
func (ds *simpleDataset[D, L]) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, ds)
	if err != nil {
		return err
	}
	return nil
}
