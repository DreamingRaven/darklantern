/* eCKKS is an abstraction of the lattigo CKKS implementation to streamline and simplify both logic and conceptual complexity. eCKKS wraps functionality like key generation in the getters and setters, so that simply requesting a resource will ensure that it exists if it is possible.*/

package erray

import (
	"encoding/json"
	"errors"

	"github.com/ldsec/lattigo/v2/ckks"
	"github.com/ldsec/lattigo/v2/ckks/bootstrapping"
	"github.com/ldsec/lattigo/v2/rlwe"
)

// Lattigo compatible slice data types for generics
// REQUIRES go 1.18 -> https://go.dev/blog/go1.18beta2
type LattigoCompatible interface {
	~float64 // | ~complex128
}

// the purposely non-exported underlying data struct that holds
// the necessary CKKS information and array like shape
type eCKKS[T LattigoCompatible] struct {
	Shape        *[]int                      `json:"shape"`       // the effective shape of this Erray
	Data         *[]T                        `json:"data"`        // the message
	Cyphertext   *ckks.Ciphertext            `json:"cyphertext"`  // Encrypted cyphertext storage of data
	Encrypted    bool                        `json:"isEncrypted"` // whether 'cyphertext'=1 or 'plaintext'=0
	Params       *ckks.Parameters            `json:"params"`      // encoder/ encryptor parameters
	Degree       int                         `json:"degree"`      // maximum polynomial degree experienced by cyphertext
	SK           *rlwe.SecretKey             `json:"sk"`          // generated secret key based on CKKS params (SENSITIVE)
	PK           *rlwe.PublicKey             `json:"pk"`          // generated public key based on CKKS params
	RLK          *rlwe.RelinearizationKey    `json:"rlk"`         // generated relinearization key based on CKKS params
	encoder      *ckks.Encoder               // encoder instance to encode message to plaintext
	encryptor    *ckks.Encryptor             // encryptor instance of encoded polynomial
	decryptor    *ckks.Decryptor             // CKKS decryptor instance of cyphertext
	evaluator    *ckks.Evaluator             // CKKS computational circuit evaluator instance
	bootstrapper *bootstrapping.Bootstrapper // bootstrapper/ key-refresher
	// kgen      *rlwe.KeyGenerator          // generator for various keys (SENSITIVE)
}

// *****
// UTILS
// *****

// Erray interface factory for eCKKS base data
// instantiates basic eCKKS struct with default or provided values
// func NewCKKSErrayC128() Erray[complex128] {
// 	return &eCKKS[complex128]{}
// }
//
// func NewCKKSErrayF64() Erray {
// 	return &eCKKS[float64]{}
// }

func NewCKKSErray[T LattigoCompatible]() Erray[T] {
	return &eCKKS[T]{}
}

// type LattigoCompatible interface {
// 	float64 | complex128
// }

// *******************
// GETTERS AND SETTERS
// *******************

// SetData sets message/ data into underlying eCKKS struct
// if the data is already present do as asked but notify
func (eckks *eCKKS[T]) SetData(newData *[]T) error {
	if eckks.Data != nil {
		eckks.Data = newData
		eckks.Shape = &[]int{len(*newData)}
		return errors.New("ckks.data already exists cannot overwrite")
	}
	eckks.Data = newData
	eckks.Shape = &[]int{len(*newData)}
	return nil
}

// Get existing data only
func (eckks *eCKKS[T]) GetData() *[]T {
	return eckks.Data
}

// Get cyphertext data
func (eckks *eCKKS[T]) GetCyphertext() (*ckks.Ciphertext, error) {
	if eckks.Cyphertext == nil {
		err := eckks.Encrypt()
		if err != nil {
			return nil, err
		}
	}
	return eckks.Cyphertext, nil
}

// set imaginary shape of data
func (eckks *eCKKS[T]) SetShape(newShape *[]int) {
	eckks.Shape = newShape
}

// get existing imaginary shape of data only
func (eckks *eCKKS[T]) GetShape() (*[]int, error) {
	if eckks.Shape == nil {
		return nil, errors.New("eckks has not been given any desired shape data")
	}
	return eckks.Shape, nil
}

// Get size of message (the number of items)
func (eckks *eCKKS[T]) GetSize() (int, error) {
	total := 0
	shape, err := eckks.GetShape()
	if err != nil {
		return -1, err
	}
	for i := range *shape {
		total += (*shape)[i]
	}
	return total, nil
}

// set encryption parameters
func (eckks *eCKKS[T]) SetParams(newParams *ckks.Parameters) error {
	if eckks.Params != nil {
		eckks.Params = newParams
		return errors.New("ckks.params already exist overwriting is dangerous")
	}
	eckks.Params = newParams
	return nil
}

// Get existing encryption parameters only
func (eckks *eCKKS[T]) GetParams() (*ckks.Parameters, error) {
	if eckks.Params == nil {
		return nil, errors.New("eckks.GetParams() no parameters have been provided")
	}
	return eckks.Params, nil
}

// Get existing secret key or attempt to generate a new one
func (eckks *eCKKS[T]) GetSK() (*rlwe.SecretKey, error) {
	if eckks.SK == nil {
		switch eckks.PK {
		case nil:
			err := eckks.InitKeys()
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("SK does not exist, but other residual keys remain, aborting")
		}
	}
	return eckks.SK, nil
}

// Get existing public key or attempt to generate a new one
func (eckks *eCKKS[T]) GetPK() (*rlwe.PublicKey, error) {
	if eckks.PK == nil {
		switch eckks.SK {
		case nil:
			err := eckks.InitKeys()
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("eCKKS.GetPK() PK does not exist, but other residual keys remain, aborting")
		}
	}
	return eckks.PK, nil
}

// Get existing relinearisation Key or attempt to generate a new one
func (eckks *eCKKS[T]) GetRK() (*rlwe.RelinearizationKey, error) {
	if eckks.RLK == nil {
		if eckks.PK == nil && eckks.SK == nil {
			err := eckks.InitKeys()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("eCKKS.GetRK() relin key does not exits, but there are other residual keys, aborting")
		}
	}
	return eckks.RLK, nil
}

// Generate key set
func (eckks *eCKKS[T]) InitKeys() error {
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
	eckks.SK = sk
	eckks.PK = pk
	eckks.RLK = rlk
	return nil
}

// GetEncoder if exists or attempt generation of new encoder
func (eckks *eCKKS[T]) GetEncoder() (*ckks.Encoder, error) {
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
func (eckks *eCKKS[T]) GetEncryptor() (*ckks.Encryptor, error) {
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
func (eckks *eCKKS[T]) GetDecryptor() (*ckks.Decryptor, error) {
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
func (eckks *eCKKS[T]) GetEvaluator() (*ckks.Evaluator, error) {
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

// Encrypt eCKKS data and generate all intermediaries
// if they don't already exist, except encryption parameters
func (eckks *eCKKS[T]) Encrypt() error {
	params, err := eckks.GetParams()
	if err != nil {
		return err
	}
	encoder, err := eckks.GetEncoder()
	if err != nil {
		return err
	}
	message := eckks.GetData()
	// (*encoder).EncodeNew()
	plaintext := (*encoder).EncodeNew(*message, params.MaxLevel(), params.DefaultScale(), params.LogSlots())
	encryptor, err := eckks.GetEncryptor()
	if err != nil {
		return err
	}
	cyphertext := (*encryptor).EncryptNew(plaintext)
	// fmt.Printf("plaintext == [%T] %+v\n", plaintext, plaintext)
	// fmt.Printf("Cyphertext == [%T] %+v\n", cyphertext, cyphertext)
	eckks.Cyphertext = cyphertext
	return nil
}

// func (eckks *eCKKS[T]) Encrypt() error {
// 	return nil
// }

// Decrypt eCKKS data using or generating intermediaries
// except parameters and of course the keys as it will
// just decrypt garbage without the original keys
func (eckks *eCKKS[T]) Decrypt() error {
	// Im really not a fan of this tedious error handling
	params, err := eckks.GetParams()
	if err != nil {
		return err
	}
	encoder, err := eckks.GetEncoder()
	if err != nil {
		return err
	}
	decryptor, err := eckks.GetDecryptor()
	if err != nil {
		return err
	}
	cyphertext, err := eckks.GetCyphertext()
	if err != nil {
		return err
	}
	size, err := eckks.GetSize()
	if err != nil {
		return err
	}
	padded := (*encoder).Decode((*decryptor).DecryptNew(cyphertext), params.LogSlots())
	message := make([]T, size)
	for i := range message {
		message[i] = T(real(padded[i]))

		// switch {
		// case reflect.TypeOf(message[0]) == reflect.TypeOf(padded[0]):
		// 	message[i] = complex(real(padded[i]), 0)
		// default:
		// 	message[i] = float64(real(padded[i]))
		// 	// return errors.New("Unsupported generic T type %T", message[0])
		// }
	}
	(*eckks).Data = &message
	// fmt.Printf("%T\n", padded)
	// fmt.Printf("%T\n", message)
	// fmt.Printf("%v\n", len(message))
	return nil
}

// ******************
// ABELIAN OPERATIONS
// ******************

// Add this eCKKS array struct with another Erray outputing a
// completely new Erray with the comparable parameters
func (eckks *eCKKS[T]) Add(other *Erray[T]) Erray[T] {
	eckks.Bud()
	return eckks
}

// Multiply this eCKKS array struct with another Erray outputing a
// completeley new Erray with comparable parameters
func (eckks *eCKKS[T]) Multiply(other *Erray[T]) Erray[T] {
	return eckks
}

// *******************
// Object Reproduction
// *******************

// Bud current object using the process of budding to replicate
// underlying DNA of object in this case parameters and keys
// creating a near replica that does not hold the same body of data
func (eckks *eCKKS[T]) Bud() Erray[T] {
	return eckks
}

// DeepCopy using ToJSON and FromJSON. Returns an error if either operation fails.
func (eckks *eCKKS[T]) DeepCopy() (Erray[T], error) {
	j, err := eckks.ToJSON()
	if err != nil {
		return nil, err
	}
	neo := NewCKKSErray[T]()
	err = neo.FromJSON(j)
	if err != nil {
		return nil, err
	}
	return neo, nil
}

// ****************
// JSON Marshalling
// ****************

// ToJSON converts internal struct to json bytes and returns those bytes or errors
func (eckks *eCKKS[T]) ToJSON() ([]byte, error) {
	marshalled, err := json.Marshal(eckks)
	if err != nil {
		return nil, err
	}
	return marshalled, nil
}

func (eckks *eCKKS[T]) FromJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, eckks)
	if err != nil {
		return err
	}
	return nil
}
