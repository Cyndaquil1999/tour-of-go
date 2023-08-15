package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	// math.Sqrtと自作Sqrt関数の出す答えの誤差を求める
	fmt.Println(math.Abs(Sqrt(2) - math.Sqrt(2)))
}