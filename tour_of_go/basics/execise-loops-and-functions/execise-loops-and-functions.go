package main

import (
	"fmt"
	"math"
)

// Sqrt 平方根の計算
func Sqrt(x float64) float64 {
	z := 1.0
	var t float64

	for i := 1; math.Abs(t-z) > 0.000001; i++ {
		t = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iter %v, %v\n", i, z)
	}
	return z
}

func main() {
	fmt.Println("Original function's result:", Sqrt(2))
	fmt.Println("System math's result:", math.Sqrt(2))
}
