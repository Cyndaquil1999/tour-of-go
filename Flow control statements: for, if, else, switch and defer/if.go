package main

import (
	"fmt"
	"math"
)

//if文もCとPython混ぜた感じ
//条件ステートメントの前にfor文と同じように初期化ステートメントを書くことが可能。
//ex) 
//if v := math.Pow(x, n); v < hoge{
//	return v
//}

//ちなみに、初期化ステートメントはelseでもそのまま使える。
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}