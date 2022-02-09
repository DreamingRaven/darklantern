/**
 * @Author: George Onoufriou <archer>
 * @Date:   2021-12-02T10:02:50+00:00
 * @Last modified by:   archer
 * @Last modified time: 2021-12-07T22:47:20+00:00
 */
package erray

import (
	"github.com/ldsec/lattigo/v2/ckks"
)

// Generic abstraction of what basic functions an encryptable array (Erray)
// supports
type Erray interface {
	// Abelian operations
	Add(other *Erray) Erray
	Multiply(other *Erray) Erray
	// getters and setters
	GetData() *[]float64
	SetData(newData *[]float64) error
	GetParams() *ckks.Parameters
	SetParams(newParams *ckks.Parameters) error
	// encryption operations
	Encrypt() error
	Decrypt() error
}
