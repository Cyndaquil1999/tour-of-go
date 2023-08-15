package main

import "fmt"

// Golangにおけるwhileはforで書き表す

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}