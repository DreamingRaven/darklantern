package dataset

import (
	"fmt"
)

type dataset struct {
}

type Dataset interface {
}

// NewNeuralNet initialises a new neural network
func NewExampleDataset() Dataset {
	return &dataset{}
}

func (ds *dataset) train() error {
	fmt.Println("Tr")
	return nil
}
