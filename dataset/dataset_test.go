package dataset

import (
	"fmt"
	"testing"

	"gitlab.com/deepcypher/darklantern/erray"
)

func TestSliceOfSlices(t *testing.T) {
	ss := sliceOfSlices[float64](10, 20)
	for i := 0; i < len(*ss); i++ {
		if (*ss)[i][1] != float64(i) {
			t.Fatal(fmt.Sprintf("slice of slices failed as %v != %v", (*ss)[i], float64(i)))
		}
	}
}

func TestSliceOfErrays(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := SliceOfErrays(sos)
	for i := 0; i < len(*soe); i++ {
		if (*soe)[i] == nil {
			t.Fatal("Slice of Errays giving back empty errays pointers but should be populated")
		}
	}

}

func TestDatasetInit(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := SliceOfErrays(sos)
	NewSimpleDataset[erray.Erray[float64], float64](soe)
}

func TestDatasetGet(t *testing.T) {
	// construct dataset from slice of slices
	sos := sliceOfSlices[float64](10, 20)
	soe := SliceOfErrays(sos)

	ds := NewSimpleDataset[erray.Erray[float64], float64](soe)

	// now check dataset Get function returns what was originally there
	// for each example row
	for i := 0; i < len(*sos); i++ {
		e, err := ds.Get(i)
		if err != nil {
			t.Fatal(err)
		}
		data_here := *(*e).GetData()
		// for each feature in example row
		for j := 0; j < len(data_here); j++ {
			// if it does not match origin fail
			if data_here[j] != (*sos)[i][j] {
				t.Fatal(fmt.Sprintf("example [%v,%v] does not match %v != %v", i, j, data_here[j], (*sos)[i][j]))
			}
		}
	}
}

func TestDatasetLen(t *testing.T) {
	sos := sliceOfSlices[float64](10, 20)
	soe := SliceOfErrays(sos)
	ds := NewSimpleDataset[erray.Erray[float64], float64](soe)
	dsLength, err := ds.Length()
	if err != nil {
		t.Fatal(err)
	}
	if dsLength != len(*sos) {
		t.Fatal("Number of examples in slice of slices does not equal number of examples in dataset")
	}
}
