package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy, dy)

	for i := 0; i < len(slice); i++ {
		slice[i] = make([]uint8, dx)

		for j := 0; j < len(slice[i]); j++ {
			slice[i][j] = uint8(i*j)
		}
	}

	return slice
}

func main() {
	pic.Show(Pic)
}