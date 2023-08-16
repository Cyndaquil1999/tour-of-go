package main

import "fmt"

func main() {
	pow := make([]int, 10)
	
	// indexだけなら省略可能
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	
	// 要素だけの場合はindexを取る部分に"_"を使う
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
