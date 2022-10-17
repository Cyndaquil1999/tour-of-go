package main

import "fmt"

//CとPython混ぜた感じある
//初期化と後処理ステートメントは省略可能
//なお、単独のステートメントの場合、セミコロンの省略も可能
//whileはforで書く（条件を書くことでwhileと同じなので）
//ループ条件を省略するとwhile trueに相当
func main() {
	sum := 0
	for i := 0;i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}