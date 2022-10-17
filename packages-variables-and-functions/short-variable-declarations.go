package main

import "fmt"

//関数内では:=で暗黙的な型宣言が可能。外ではvar, funcなどを用いないと宣言できない。
func main() {
	var i, j int = 1,2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}