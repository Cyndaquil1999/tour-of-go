package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for x := range picture {
		picture[x] = make([]uint8, dx)
		for y := range picture[x] {
			picture[x][y] = uint8(x ^ y)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
