package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

// make a type error-able by implemernting Error interface method. In this case we want neg float vals to error so we impl Error method on MyFloat type 
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 1; i < 11; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z, nil
}

