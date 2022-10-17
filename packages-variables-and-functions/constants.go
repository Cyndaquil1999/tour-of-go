package main

import "fmt"

//定数はconstで表現する。
//constを用いれるのは文字、文字列、bool, 数値のみ。
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}