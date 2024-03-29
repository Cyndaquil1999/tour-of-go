package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	
	// 関数を引数に取ることができる
	fmt.Println(compute(hypot)) // sqrt(3^2 + 4^2) = 5
	fmt.Println(compute(math.Pow)) // pow(3, 4) = 81
}
