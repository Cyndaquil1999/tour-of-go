package main

import "fmt"

var i, j int = 1,2

//型を省略すると、initializerが持つ型になる
func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i,j,c,python,java)
}