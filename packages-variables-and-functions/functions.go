package main

import "fmt"

//2つ以上の引数が同じ型である場合、下の例ではx, y intのように記述できる。
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42,13))
}