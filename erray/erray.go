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
type Erray[T LattigoCompatible] interface {
	// Abelian operations
	Add(other *Erray[T]) Erray[T]
	Multiply(other *Erray[T]) Erray[T]
	// getters and setters
	GetData() *[]T
	SetData(newData *[]T) error
	GetParams() (*ckks.Parameters, error)
	SetParams(newParams *ckks.Parameters) error
	// encryption operations
	Encrypt() error
	Decrypt() error
	// object reproduction
	Bud() Erray[T]
	DeepCopy() (Erray[T], error)
	// JSON marshalling
	ToJSON() ([]byte, error)
	FromJSON(bytes []byte) error
}
