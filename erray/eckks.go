/* eCKKS is an abstraction of the lattigo CKKS implementation to streamline and simplify both logic and conceptual complexity. eCKKS wraps functionality like key generation in the getters and setters, so that simply requesting a resource will ensure that it exists if it is possible.*/

package erray

import (
	"errors"

	"github.com/ldsec/lattigo/v2/ckks"
	"github.com/ldsec/lattigo/v2/ckks/bootstrapping"
	"github.com/ldsec/lattigo/v2/rlwe"
)

// the purposely non-exported underlying data struct that holds
// the necessary CKKS information and array like shape
type eCKKS struct {
	shape        []int                       // the effective shape of this Erray
	data         *[]float64                  // the message, plaintext, or cyphertext data
	encrypted    bool                        // whether 'cyphertext'=1 or 'plaintext'=0
	params       *ckks.Parameters            // encoder/ encryptor parameters
	degree       int                         // maximum polynomial degree experienced by cyphertext
	sk           *rlwe.SecretKey             // generated secret key based on CKKS params (SENSITIVE)
	pk           *rlwe.PublicKey             // generated public key based on CKKS params
	rlk          *rlwe.RelinearizationKey    // generated relinearization key based on CKKS params
	encoder      *ckks.Encoder               // encoder instance to encode message to plaintext
	encryptor    *ckks.Encryptor             // encryptor instance of encoded polynomial
	decryptor    *ckks.Decryptor             // CKKS decryptor instance of cyphertext
	evaluator    *ckks.Evaluator             // CKKS computational circuit evaluator instance
	bootstrapper *bootstrapping.Bootstrapper // bootstrapper/ key-refresher
	// kgen         *rlwe.KeyGenerator          // generator for various keys (SENSITIVE)
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

// Get existing data only
func (eckks *eCKKS) GetData() *[]float64 {
	return eckks.data
}

// set imaginary shape of data
func (eckks *eCKKS) SetShape(newShape []int) {
	eckks.shape = newShape
}

// get existing imaginary shape of data only
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

// Get existing encryption parameters only
func (eckks *eCKKS) GetParams() (*ckks.Parameters, error) {
	if eckks.params == nil {
		return nil, errors.New("eckks.GetParams() no parameters have been provided")
	}
	return eckks.params, nil
}

// Get existing secret key or attempt to generate a new one
func (eckks *eCKKS) GetSK() (*rlwe.SecretKey, error) {
	if eckks.sk == nil {
		switch eckks.pk {
		case nil:
			err := eckks.InitKeys()
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("SK does not exist, but other residual keys remain, aborting")
		}
	}
	return eckks.sk, nil
}

// Get existing public key or attempt to generate a new one
func (eckks *eCKKS) GetPK() (*rlwe.PublicKey, error) {
	if eckks.pk == nil {
		switch eckks.sk {
		case nil:
			err := eckks.InitKeys()
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("eCKKS.GetPK() PK does not exist, but other residual keys remain, aborting")
		}
	}
	return eckks.pk, nil
}

// Get existing relinearisation Key or attempt to generate a new one
func (eckks *eCKKS) GetRK() (*rlwe.RelinearizationKey, error) {
	if eckks.rlk == nil {
		if eckks.pk == nil && eckks.sk == nil {
			err := eckks.InitKeys()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("eCKKS.GetRK() relin key does not exits, but there are other residual keys, aborting")
		}
	}
	return eckks.rlk, nil
}

// Generate key set
func (eckks *eCKKS) InitKeys() error {
	params, err := eckks.GetParams()
	if err != nil {
		return err
	}
	if params == nil {
		return errors.New("eckks.InitKeys() has no encryption parameters to encrypt with")
	}
	// generate keys together
	kgen := ckks.NewKeyGenerator(*params)
	sk, pk := kgen.GenKeyPair()
	rlk := kgen.GenRelinearizationKey(sk, 2)
	// assign keys together
	eckks.sk = sk
	eckks.pk = pk
	eckks.rlk = rlk
	return nil
}

// GetEncoder if exists or attempt generation of new encoder
func (eckks *eCKKS) GetEncoder() (*ckks.Encoder, error) {
	params, err := eckks.GetParams()
	if err != nil {
		return nil, err
	}
	if eckks.encoder == nil {
		encoder := ckks.NewEncoder(*params)
		eckks.encoder = &encoder
		// eckks.encoder = &(ckks.NewEncoder(*params))
	}
	return eckks.encoder, nil
}

// GetEncryptor if exists or attempt generation of new encryptor
func (eckks *eCKKS) GetEncryptor() (*ckks.Encryptor, error) {
	if eckks.encryptor == nil {
		params, err := eckks.GetParams()
		if err != nil {
			return nil, err
		}
		pk, err := eckks.GetPK()
		if err != nil {
			return nil, err
		}
		encryptor := ckks.NewEncryptor(*params, pk)
		eckks.encryptor = &encryptor
	}
	return eckks.encryptor, nil
}

// GetDecryptor if exists or attempt generation of new decryptor
func (eckks *eCKKS) GetDecryptor() (*ckks.Decryptor, error) {
	if eckks.decryptor == nil {
		params, err := eckks.GetParams()
		if err != nil {
			return nil, err
		}
		sk, err := eckks.GetSK()
		if err != nil {
			return nil, err
		}
		decryptor := ckks.NewDecryptor(*params, sk)
		eckks.decryptor = &decryptor
	}
	return eckks.decryptor, nil
}

// GetEvaluator if exists or attempt generation of new evaluator
func (eckks *eCKKS) GetEvaluator() (*ckks.Evaluator, error) {
	if eckks.evaluator == nil {
		params, err := eckks.GetParams()
		if err != nil {
			return nil, err
		}
		rk, err := eckks.GetRK()
		if err != nil {
			return nil, err
		}
		evaluator := ckks.NewEvaluator(*params, rlwe.EvaluationKey{Rlk: rk})
		eckks.evaluator = &evaluator
	}
	return eckks.evaluator, nil
}

// *********************
// ENCRYPTION OPERATIONS
// *********************

// // Encrypt eCKKS data and generate all intermediaries
// // if they don't already exist, except encryption parameters
// func (eckks *eCKKS) Encrypt() error {
// 	params, err := eckks.GetParams()
// 	if err != nil {
// 		return err
// 	}
// 	encoder, err := eckks.GetEncoder()
// 	if err != nil {
// 		return err
// 	}
// 	message := eckks.GetData()
// 	plaintext := (*encoder).EncodeNew(*message, params.LogSlots())
// 	return errors.New("Not yet implemented encryption.")
// }

func (eckks *eCKKS) Encrypt() error {
	return nil
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
