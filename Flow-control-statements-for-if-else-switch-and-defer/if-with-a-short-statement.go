package main

import (
	"fmt"
	"math"
)


func pow(x, n, lim float64) float64 {
	// ifステートメントを用いた記述
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// スコープ範囲外なのでundefinedになる
	//return v
	return lim
}

func main() {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}