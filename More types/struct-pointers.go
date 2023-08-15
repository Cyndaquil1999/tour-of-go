package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// structへのポインタ
	p := &v
	
	// 本来は(*p).Xになるが、Golangではp.Xで書ける
	p.X = 1e9
	fmt.Println(v)
}
