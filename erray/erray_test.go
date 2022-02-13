/**
 * @Author: George Onoufriou <archer>
 * @Date:   2021-12-02T10:02:50+00:00
 * @Last modified by:   archer
 * @Last modified time: 2021-12-07T22:24:38+00:00
 */
package erray

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ldsec/lattigo/v2/ckks"
	"github.com/ldsec/lattigo/v2/utils"
)

// checking the default factory works
// no need to check types as this is go
func TestNewCKKSErray(t *testing.T) {
	NewCKKSErray[float64]()
}

// check that generated structs assigns param values correctly
func TestCKKSGetSetParams(t *testing.T) {
	o := NewCKKSErray[float64]()
	parms, _ := ckks.NewParametersFromLiteral(ckks.PN14QP438)
	o.SetParams(&parms)
	// converting parms to string for easier comparison
	// as they include multiple nonexported values that cmp
	// does not compare well
	a := fmt.Sprintf("%#v", parms)
	saved_params, err := o.GetParams()
	if err != nil {
		t.Fatal(err)
	}
	b := fmt.Sprintf("%#v", *saved_params)
	if !cmp.Equal(a, b, cmpopts.IgnoreUnexported()) {
		t.Fatal("eckks.params have not been set properly")
	}
}

// test eckks getters and setters for data are working
func TestCKKSGetSetData(t *testing.T) {
	o := NewCKKSErray[float64]()
	data := make([]float64, 3*3)
	for i := range data {
		data[i] = utils.RandFloat64(-8, 8)
	}
	o.SetData(&data)
	// https://stackoverflow.com/questions/44370277/type-is-pointer-to-interface-not-interface-confusion
	fmt.Printf("[%T] %+v\n", o.GetData(), o.GetData())
	fmt.Printf("[%T] %+v\n", &data, &data)
	if !cmp.Equal(&data, o.GetData()) {
		t.Fatal("eckks.data has not been set properly")
	}
}

func TestCKKSEncrypt(t *testing.T) {
	o := NewCKKSErray[float64]()
	params, _ := ckks.NewParametersFromLiteral(ckks.PN14QP438)
	// using params to dictate number of slots
	data := make([]float64, params.Slots())
	for i := range data {
		data[i] = utils.RandFloat64(-8, 8)
	}

	o.SetParams(&params)
	o.SetData(&data)
	err := o.Encrypt()
	if err != nil {
		t.Fatal(err)
	}
}

func FuzzFoo(f *testing.F) {
	f.Add("float64")
	f.Add("complex128")
	f.Fuzz(func(t *testing.T, typ string) {
		switch {
		case typ == "float64":
			fmt.Printf("%v\n", typ)
		case typ == "complex128":
			fmt.Printf("%v\n", typ)
		default:
			t.Errorf("\"typ=%v\" is not supported in this fuzz test \n", typ)
		}
	})
}
