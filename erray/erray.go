/**
 * @Author: George Onoufriou <archer>
 * @Date:   2021-12-02T10:02:50+00:00
 * @Last modified by:   archer
 * @Last modified time: 2021-12-07T22:47:20+00:00
 */
package erray

import (
	"errors"

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

// the purposely non-exported underlying data struct that holds
// the necessary CKKS information and array like shape
type eCKKS struct {
	shape     []int            // the effective shape of this Erray
	data      *[]float64       // the message, plaintext, or cyphertext data
	state     string           // whether 'cyphertext', 'message' or 'plaintext'
	params    *ckks.Parameters // encoder/ encryptor parameters
	secretKey string           // generated secret key based on CKKS params
	publicKey string           // generated public key based on CKKS params
	relinKey  string           // generated relin key based on CKKS params
	encoder   string           // encoder instance to encode message to plaintext
	encryptor string           // encryptor instance of encoded polynomial
	decryptor string           // CKKS decryptor instance of cyphertext
	evaluator string           // CKKS computational circuit evaluator instance
}

// Erray interface factory for eCKKS base data
// instantiates basic eCKKS struct with default or provided values
func NewCKKSErray() Erray {
	return &eCKKS{}
}

//
// GETTERS AND SETTERS
//

// SetData sets message/ data into underlying eCKKS struct
// if the data is already present do as asked but notify
func (eckks *eCKKS) SetData(newData *[]float64) error {
	if eckks.data != nil {
		eckks.data = newData
		return errors.New("ckks.data already exists cannot overwrite")
	}
	eckks.data = newData
	return nil
}

// GetData return message/ data from underlying eCKKS struct
func (eckks *eCKKS) GetData() *[]float64 {
	return eckks.data
}

// set imaginary shape of data
func (eckks *eCKKS) SetShape(newShape []int) {
	eckks.shape = newShape
}

// get imaginary shape of data
func (eckks *eCKKS) GetShape() []int {
	return eckks.shape
}

// set encryption parameters
func (eckks *eCKKS) SetParams(newParams *ckks.Parameters) error {
	if eckks.params != nil {
		eckks.params = newParams
		return errors.New("ckks.params already exist overwriting is dangerous")
	}
	eckks.params = newParams
	return nil
}

// get encryption parameters
func (eckks *eCKKS) GetParams() *ckks.Parameters {
	return eckks.params
}

//
// ENCRYPTION OPERATIONS
//

// Encrypt eCKKS data and generate all intermediaries
// if they don't already exist, except encryption parameters
func (eckks *eCKKS) Encrypt() error {
	return errors.New("Not yet implemented encryption.")
}

// Decrypt eCKKS data using or generating intermediaries
// except parameters and of course the keys as it will
// just decrypt garbage without the original keys
func (eckks *eCKKS) Decrypt() error {
	return errors.New("Not yet implemented decryption.")
}

//
// ABELIAN OPERATIONS
//

// Add this eCKKS array struct with another Erray
func (eckks *eCKKS) Add(other *Erray) Erray {
	return eckks
}

// Multiply this eCKKS array struct with another Erray
func (eckks *eCKKS) Multiply(other *Erray) Erray {
	return eckks
}
