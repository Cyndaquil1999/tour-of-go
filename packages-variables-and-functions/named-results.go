package main

import "fmt"

//戻り値変数を予め書いておくことで、undefinedにならない。
//ただ、長い関数では可読性を損なうので、短い関数で使用するべき。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}