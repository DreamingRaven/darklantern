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
	kb := getBatch(ds, 2, 32, &dsidx)
	if len(*kb.batch) != 32 {
		t.Fatal("getBatch has not created a batch of the right size")
	}
	if (*kb.batch)[0] == nil {
		t.Fatal("getBatch has not populated with a pointer to the data")
	}
	if (*kb.batch)[len(*kb.batch)-1] == nil {
		t.Fatal("getBatch has not populated with a pointer to the data")
	}
	kb = getBatch(ds, 600, 32, &dsidx)
	if (*kb.batch)[0] != nil {
		t.Fatal("getBatch has failed to get nil filled batch when selecting out of bounds")
	}
}

func TestDataloading(t *testing.T) {
	sos := dataset.ExampleSliceOfSlices[float64](100, 10)
	soe := dataset.SliceOfErrays(sos)
	ds := dataset.NewSimpleDataset[erray.Erray[float64], float64](soe)
	ch, _ := SimpleDataloader(ds, 4, 32, true, true)
	i := 0
	for batch := range ch {
		fmt.Printf("batch: %v, len: %v\n", i, len(batch))
		i += 1
	}
}

// func TestMultiDataloading(t *testing.T) {
// 	sos := dataset.ExampleSliceOfSlices[float64](100, 10)
// 	soe := dataset.SliceOfErrays(sos)
// 	ds := dataset.NewSimpleDataset[erray.Erray[float64], float64](soe)
// 	ch, _ := SimpleDataloader(ds, 4, 32, true, true)
// 	for batch := range ch {
// 		fmt.Println("Am I even running")
// 		fmt.Println(batch)
// 	}
// }
