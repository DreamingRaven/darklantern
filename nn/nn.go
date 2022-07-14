package nn

import (
	"fmt"
	"sync"

	"gitlab.com/deepcypher/darklantern/dager"
)

type neuralNet struct {
	graph dager.Dager
	// optimiser
	// loss function
	lock sync.RWMutex
}

type NeuralNet interface {
	train() error
	infer() error
}

// NewNeuralNet initialises a new neural network
func NewNeuralNet() NeuralNet {
	return &neuralNet{}
}

func (nn *neuralNet) train() error {
	fmt.Println("Training")
	return nil
}

func (nn *neuralNet) infer() error {
	fmt.Println("Infering")
	return nil
}
