package dataloader

import (
	"testing"

	"gitlab.com/deepcypher/darklantern/dataset"
	"gitlab.com/deepcypher/darklantern/erray"
)

func TestDataloading(t *testing.T) {
	sos := dataset.ExampleSliceOfSlices[float64](10, 20)
	soe := dataset.SliceOfErrays(sos)
	ds := dataset.NewSimpleDataset[erray.Erray[float64], float64](soe)
	SimpleDataloader(ds, 4, 32, true)
}
