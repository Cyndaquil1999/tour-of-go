package main

import (
	"fmt"
)



func Sqrt(x float64) float64 {
	z := 1.0
	diff := 1e-10
	for tmp:= 1.0; tmp*tmp > diff; z -= tmp{
		tmp = (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}