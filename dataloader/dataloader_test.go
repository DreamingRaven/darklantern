package dataloader

import (
	"fmt"
	"testing"
)

func TestDataloaderInit(t *testing.T) {
	ds := NewExampleDataloader()
	fmt.Println(ds)
}
