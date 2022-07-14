package nn

import (
	"fmt"
	"testing"
)

func TestNNInit(t *testing.T) {
	nn := NewNeuralNet()
	fmt.Println(nn)
}
