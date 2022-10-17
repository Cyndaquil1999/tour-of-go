package main

import "fmt"

//defer文は呼び出し元関数がreturnするまで実行を遅延するもの。
//以下の例ではhelloが先に出力され、最後にWorldが出力される。
//複数のdeferの場合はstackで保存するため、入れた順と逆順に実行される。
func main() {
	defer fmt.Println("World")

	fmt.Println("hello")
}