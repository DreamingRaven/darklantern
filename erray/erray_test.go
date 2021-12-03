/**
 * @Author: George Onoufriou <archer>
 * @Date:   2021-12-02T10:02:50+00:00
 * @Last modified by:   archer
 * @Last modified time: 2021-12-02T21:16:37+00:00
 */
package erray

import (
	"fmt"
	"testing"

	// "github.com/ldsec/lattigo/v2"
	"github.com/ldsec/lattigo/v2/ckks"
)

func TestErrayStruct(t *testing.T) {
	parms, _ := ckks.NewParametersFromLiteral(ckks.PN14QP438)
	fmt.Printf("%T", parms)
	t.Fatalf("Erray is incomplete")

}
