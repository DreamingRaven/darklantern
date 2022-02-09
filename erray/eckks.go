package erray

import (
	"errors"
	"fmt"

	"github.com/ldsec/lattigo/v2/ckks"
	"github.com/ldsec/lattigo/v2/rlwe"
)

// the purposely non-exported underlying data struct that holds
// the necessary CKKS information and array like shape
type eCKKS struct {
	shape     []int                    // the effective shape of this Erray
	data      *[]float64               // the message, plaintext, or cyphertext data
	encrypted bool                     // whether 'cyphertext'=1 or 'plaintext'=0
	params    *ckks.Parameters         // encoder/ encryptor parameters
	sk        *rlwe.SecretKey          // generated secret key based on CKKS params
	pk        *rlwe.PublicKey          // generated public key based on CKKS params
	rlk       *rlwe.RelinearizationKey // generated relinearization key based on CKKS params
	encoder   string                   // encoder instance to encode message to plaintext
	encryptor string                   // encryptor instance of encoded polynomial
	decryptor string                   // CKKS decryptor instance of cyphertext
	evaluator string                   // CKKS computational circuit evaluator instance
}

// *****
// UTILS
// *****

// Erray interface factory for eCKKS base data
// instantiates basic eCKKS struct with default or provided values
func NewCKKSErray() Erray {
	return &eCKKS{}
}

// *******************
// GETTERS AND SETTERS
// *******************

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

// *********************
// ENCRYPTION OPERATIONS
// *********************

// Encrypt eCKKS data and generate all intermediaries
// if they don't already exist, except encryption parameters
func (eckks *eCKKS) Encrypt() error {
	// check if secret key exists, and not already encrypted.
	//If not generate.
	if eckks.sk == nil && eckks.data != nil {
		kgen := ckks.NewKeyGenerator(*eckks.params)
		sk, pk := kgen.GenKeyPair()
		fmt.Printf("Secret Key Type: %T\n", sk)
		fmt.Printf("Public Key Type: %T\n", pk)
	}
	return errors.New("Not yet implemented encryption.")
}

// Decrypt eCKKS data using or generating intermediaries
// except parameters and of course the keys as it will
// just decrypt garbage without the original keys
func (eckks *eCKKS) Decrypt() error {
	return errors.New("Not yet implemented decryption.")
}

// ******************
// ABELIAN OPERATIONS
// ******************

// Add this eCKKS array struct with another Erray
func (eckks *eCKKS) Add(other *Erray) Erray {
	return eckks
}

// Multiply this eCKKS array struct with another Erray
func (eckks *eCKKS) Multiply(other *Erray) Erray {
	return eckks
}
