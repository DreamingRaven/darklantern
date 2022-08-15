package dataloader

import (
	"fmt"
	"testing"

	"gitlab.com/deepcypher/darklantern/dataset"
	"gitlab.com/deepcypher/darklantern/erray"
)

func TestGetBatch(t *testing.T) {
	sos := dataset.ExampleSliceOfSlices[float64](100, 10)
	soe := dataset.SliceOfErrays(sos)
	ds := dataset.NewSimpleDataset[erray.Erray[float64], float64](soe)
	length, _ := ds.Length()
	// constructing default mapping (dsidx) (before shuffling) to indicate which examples fall in which order
	dsidx := make([]int, length)
	for i := 0; i < len(dsidx); i++ {
		dsidx[i] = i
	}
	batch := getBatch(ds, 2, 32, &dsidx)
	if len(batch) != 32 {
		t.Fatal("getBatch has not created a batch of the right size")
	}
	if batch[0] == nil {
		t.Fatal("getBatch has not populated with a pointer to the data")
	}
	if batch[len(batch)-1] == nil {
		t.Fatal("getBatch has not populated with a pointer to the data")
	}
	batch = getBatch(ds, 600, 32, &dsidx)
	if batch[0] != nil {
		t.Fatal("getBatch has failed to get nil filled batch when selecting out of bounds")
	}
}

func TestDataloading(t *testing.T) {
	sos := dataset.ExampleSliceOfSlices[float64](100, 10)
	soe := dataset.SliceOfErrays(sos)
	ds := dataset.NewSimpleDataset[erray.Erray[float64], float64](soe)
	ch, _ := SimpleDataloader(ds, 4, 32, true, true)
	for batch := range ch {
		fmt.Println("Am I even running")
		fmt.Println(batch)
	}
}
