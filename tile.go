package main

import (
	"fmt"
	"image"
)

// Converts one 8x8 tile from the given coordinate of the image and returns its 16 bytes
func convertTile(image image.Image, colors ImageColors, startX int, startY int) ([16]byte, error) {
	tile := [16]byte{0}
	var err error = nil

out:
	for y := 0; y < 8; y++ {
		var rowMask0 byte = 0
		var rowMask1 byte = 0
		for x := 0; x < 8; x++ {
			// Find our color's index - determines which mask we want to fill
			c := 0
			col := image.At(startX+x, startY+y)
			for i := 0; i < 4; i++ {
				if colorsMatch(col, colors[i]) {
					c = i
					break
				} else {
					// This color isn't in our list, meaning this image has too many colors!
					if i == 3 {
						err = fmt.Errorf("too many colors in tile at %d, %d", startX, startY)
						break out
					}
				}
			}

			// Insert one at the end of the byte dependent on the color's index
			if c == 1 || c == 3 {
				rowMask0 |= 1
			}
			if c == 2 || c == 3 {
				rowMask1 |= 1
			}
			// For every bit except the last, shift all the bytes so far to the left
			if x < 7 {
				rowMask0 <<= 1
				rowMask1 <<= 1
			}
		}

		// Set bytes for mask 0 and 1 simultaneously
		tile[y] = rowMask0
		tile[y+8] = rowMask1
	}

	return tile, err
}
