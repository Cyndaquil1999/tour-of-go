package main

import "fmt"

func main() {
	// 呼び出し元の関数が終了するまで遅延する
	defer fmt.Println("world")

	fmt.Println("hello")
}
