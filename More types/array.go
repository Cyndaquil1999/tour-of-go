package main

import "fmt"

func main() {
	// [2] stringでもok
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	
	// 配列の初期化はこのように書くこともできる
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
