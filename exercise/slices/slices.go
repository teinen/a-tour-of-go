package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// make image slice
	image := make([][]uint8, dy)

	// make each rows
	for i := 0; i < dy; i++ {
		image[i] = make([]uint8, dx)

		// create patterns
		for j := 0; j < dx; j++ {
			image[i][j] = uint8(i * j) // You can change intValue!!
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
