package dataloader

import (
	"fmt"
	"testing"

	"gitlab.com/deepcypher/darklantern/dataset"
	"gitlab.com/deepcypher/darklantern/erray"
)

func TestDataloading(t *testing.T) {
	sos := dataset.ExampleSliceOfSlices[float64](100, 10)
	soe := dataset.SliceOfErrays(sos)
	ds := dataset.NewSimpleDataset[erray.Erray[float64], float64](soe)
	ch, _ := SimpleDataloader(ds, 4, 32, true, true)
	for batch := range ch {
		fmt.Println(batch)
	}
}
