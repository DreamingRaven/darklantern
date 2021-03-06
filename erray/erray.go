/**
 * @Author: George Onoufriou <archer>
 * @Date:   2021-12-02T10:02:50+00:00
 * @Last modified by:   archer
 * @Last modified time: 2021-12-07T22:47:20+00:00
 */
package erray

import (
	"github.com/ldsec/lattigo/v2/ckks"
	"github.com/ldsec/lattigo/v2/rlwe"
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
	GetSK() (*rlwe.SecretKey, error)
	GetPK() (*rlwe.PublicKey, error)
	GetRK() (*rlwe.RelinearizationKey, error)
	// encryption operations
	Encrypt() error
	Decrypt() error
	// object reproduction
	Bud() Erray[T]               // minimises memory usage by sharing common resources
	DeepCopy() (Erray[T], error) // completely independent instances
	// JSON marshalling
	ToJSON() ([]byte, error)
	FromJSON(bytes []byte) error
}
