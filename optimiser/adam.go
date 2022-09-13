package optimiser

// aliasing paramMap so we dont have to constantly type it out
type paramMap map[string]interface{}

// Adam struct holding necessary data and history in ordert to calculate new parameters
type Adam struct {
	Alpha   float64
	Beta_1  float64
	Beta_2  float64
	Epsilon float64
	History []paramMap
}

// NewDefaultAdam creates a default adam struct for you based on https://arxiv.org/abs/1412.6980
func NewDefaultAdam() *Adam {
	return &Adam{
		Alpha:   0.001,
		Beta_1:  0.9,
		Beta_2:  0.999,
		Epsilon: 1e-8,
	}
}

// NewDefaultAdamOptimiser use default adam struct and wrap it in optimiser abstraction
func NewDefaultAdamOptimiser() Optimiser {
	return NewDefaultAdam()
}

// Optimise use the ADAM algorithm first a second order moments to calculate the new parameters
func (a *Adam) Optimise() error {
	return nil
}

// RMSProp or second order moment calculation
func RMSProp() {

}
