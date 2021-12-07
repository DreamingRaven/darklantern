/**
 * @Author: George Onoufriou <archer>
 * @Date:   2021-12-02T10:02:50+00:00
 * @Last modified by:   archer
 * @Last modified time: 2021-12-07T22:08:47+00:00
 */
package erray

type eCKKS struct {
  shape []int
  data []float64
  state string
  params string
  secretKey string
  publicKey string
  relinKey string
  encryptor string
  decryptor string
  evaluator string
}

type Erray interface {}

func NewCKKSErray () Erray {
  return eCKKS{}
}
