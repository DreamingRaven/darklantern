/**
 * @Author: George Onoufriou <archer>
 * @Date:   2021-12-02T10:02:50+00:00
 * @Last modified by:   archer
 * @Last modified time: 2021-12-07T22:24:38+00:00
 */
package erray

import (
	"fmt"
	"math"
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

// TestBud properly creates independent spinoffs of the original object
// necessary for subsequent computation in mathematical opetrations
func TestBud(t *testing.T) {
	// create the original object
	a := NewCKKSErray[float64]()
	parms, _ := ckks.NewParametersFromLiteral(ckks.PN14QP438)
	a.SetParams(&parms)
	orig := fmt.Sprintf("%#v", a)
	fmt.Printf("ORIG:\n%v\n", orig)

	b := a.Bud()
	cloned := fmt.Sprintf("%#v", b)
	fmt.Printf("CLONE:\n%v\n", cloned)

	// check that initial budding was a success
	if a != b {
		t.Fatal("Bud does not initial produce an equivelant object.")
	}
	// now check modifications on the bud compared to the original
	parmsb, _ := ckks.NewParametersFromLiteral(ckks.PN14QP411pq)
	b.SetParams(&parmsb)

	modified := fmt.Sprintf("%#v", b)
	fmt.Printf("MODIFIED:\n%v\n", modified)
	newOrig := fmt.Sprintf("%#v", b)
	fmt.Printf("NewOrigi:\n%v\n", newOrig)
	if a == b {
		t.Fatal("Bud has modified original object.")
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
	// fmt.Printf("[%T] %+v\n", o.GetData(), o.GetData())
	// fmt.Printf("[%T] %+v\n", &data, &data)
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

func FuzzECKKSParameters(f *testing.F) {
	f.Add("float64", -8, 8, "PN12QP109")
	f.Add("float64", -8, 8, "PN13QP218")
	f.Add("float64", -8, 8, "PN14QP438")
	f.Add("float64", -8, 8, "PN15QP880")
	// f.Add("complex128", -8, 8)

	f.Fuzz(func(t *testing.T, typ string, lower int, higher int, param_name string) {

		// Creating relevant parameters for encryption
		var params *ckks.Parameters
		switch {
		case param_name == "PN12QP109":
			param, err := ckks.NewParametersFromLiteral(ckks.PN12QP109)
			params = &param
			if err != nil {
				t.Errorf("%v", err)
			}
		case param_name == "PN13QP218":
			param, err := ckks.NewParametersFromLiteral(ckks.PN13QP218)
			params = &param
			if err != nil {
				t.Errorf("%v", err)
			}
		case param_name == "PN14QP438":
			param, err := ckks.NewParametersFromLiteral(ckks.PN14QP438)
			params = &param
			if err != nil {
				t.Errorf("%v", err)
			}
		case param_name == "PN15QP880":
			param, err := ckks.NewParametersFromLiteral(ckks.PN15QP880)
			params = &param
			if err != nil {
				t.Errorf("%v", err)
			}
		default:
			t.Errorf("\"%v\" is not a supported parameter name", param_name)
		}

		// Creating relevant data for encryption
		switch {
		case typ == "float64":
			data := make([]float64, 3)
			for i := range data {
				data[i] = utils.RandFloat64(float64(lower), float64(higher))
			}
			eckks := NewCKKSErray[float64]()
			eckks.SetParams(params)
			eckks.SetData(&data)
			err := eckks.Encrypt()
			if err != nil {
				t.Errorf("%v", err)
			}
			err = eckks.Decrypt()
			if err != nil {
				t.Errorf("%v", err)
			}
			// get decrypted slice again
			message := eckks.GetData()

			similar, explination, err := RoughlyEqualSlices(&data, message, 3)
			if err != nil {
				t.Errorf("%v", err)
			}
			if similar != true {
				t.Errorf("%v", explination)
			}
		// case typ == "complex128":
		// 	fmt.Printf("%v\n", typ)
		// 	data := make([]complex128, 3)
		// 	for i := range data {
		// 		data[i] = utils.RandComplex128(complex128(lower), complex128(higher))
		// 	}
		default:
			t.Errorf("\"typ=%v\" is not supported in this fuzz test \n", typ)
		}
	})
}

// RoughlyEqualSlices compares two comparable slices to a set number of decimal places for equality.
// If the slices are comparible but do not match then false is returned with a specific explenation.
// if the slices are comparible but match true is returned with no explenation.
// if the slices fail to be compared, maybe they do not allow comparison to the decimal places requested,
// then an error will be returned.
func RoughlyEqualSlices(a, b *[]float64, dp int) (equal bool, exp string, err error) {

	factor := math.Pow(10, float64(dp))
	// check if lengths are equal first as quick confirmation
	if len(*a) != len(*b) {
		return false, fmt.Sprintf("length of a (%v) != b (%v)", len(*a), len(*b)), nil
	}
	// check each value in a ROUGHLY corresponding value in b
	for i := 0; i < len(*a); i++ {
		// since there is no specific way to compare to x significant figures, we use
		// math round on blown up or shrunken numbers by a factor dependent on decimal places

		m := math.Round(float64((*a)[i])*factor) / factor
		n := math.Round(float64((*b)[i])*factor) / factor
		if m != n {
			return false, fmt.Sprintf("non match found at idx=%v: %v(%v) != %v(%v)",
				i, (*a)[i], m, (*b)[i], n), nil
		}
	}
	return true, "", nil
	// return false, errors.New("Slices could not be compared")
}

// func genTestData[T LattigoCompatible]() *[]T{
// 	data := make([]T, params.Slots())
// 	for i := range data {
// 		data[i] = utils.RandFloat64(-8, 8)
// 	}
// 	return &data
// }
