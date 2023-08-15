package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	
	// スライスのゼロ値はnil
	if s == nil {
		fmt.Println("nil!")
	}
}
