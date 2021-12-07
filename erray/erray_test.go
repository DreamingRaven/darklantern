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

	// "github.com/ldsec/lattigo/v2"
	ckks "github.com/ldsec/lattigo/v2/ckks"
)

func TestErrayStruct(t *testing.T) {
	parms, _ := ckks.NewParametersFromLiteral(ckks.PN14QP438)
	fmt.Printf("%T", parms)
	t.Fatalf("Erray is incomplete")
}

func TestNewCKKSErray(t *testing.T) {
	var o Erray = NewCKKSErray()
	oStr := fmt.Sprintf("%#v", o)
	// oType := fmt.Sprintf("%T", o)
	fmt.Printf(oStr)

}
