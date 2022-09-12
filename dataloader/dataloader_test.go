package dataloader

import (
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
		// fmt.Printf("batch: %v, len: %v\n", i, len(batch))
		if batch == nil {
			t.Fatal("batch returned nil pointer")
		}
		for j := 0; j < len(batch); j++ {
			b := batch[j]
			// if not nil observation due to being partial batch
			if b != nil {
				dat := (*b).GetData()
				// fmt.Println(dat)
				// checking that there is no intra-shuffling of observations
				for k := 0; k < len(*dat); k++ {
					if k != 0 {
						if (*dat)[k] < (*dat)[k-1] {
							t.Fatalf("'%v' is not greater than previous '%v' the observation has been intra-shuffled!", (*dat)[k], (*dat)[k-1])
						}
					}
				}
			} else {
				// fmt.Println(b)
			}
		}
		i += 1
	}
}
