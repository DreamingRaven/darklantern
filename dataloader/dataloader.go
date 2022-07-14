package dataloader

import (
	"fmt"
)

type dataloader struct {
}

type Dataloader interface {
}

// NewNeuralNet initialises a new neural network
func NewExampleDataloader() Dataloader {
	return &dataloader{}
}

func (ds *dataloader) train() error {
	fmt.Println("Tr")
	return nil
}
