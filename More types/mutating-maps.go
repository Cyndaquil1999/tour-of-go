package main

import "fmt"

func main() {
	m := make(map[string]int)
	
	// 要素の挿入
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// 要素の更新
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 要素の削除
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 2つ目の値で存在についてbooleanで返す
	// 存在しない場合の要素は要素の型のゼロ値となる
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
