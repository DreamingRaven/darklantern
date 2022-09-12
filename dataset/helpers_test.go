package dataset

import (
	"testing"
)

func TestExampleSoS(t *testing.T) {
	SoS := ExampleSliceOfSlices[float64](32, 10)
	// fmt.Println(SoS)
	for i := range *SoS {
		if int((*SoS)[i][1]) != i {
			t.Fatalf("Example slice '%v'(%v) has not been generated correctly", i, (*SoS)[i])
		}
		// fmt.Println((*SoS)[i])
	}
}
