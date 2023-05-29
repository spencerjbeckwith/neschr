package main

import (
	"image"
)

// Converts one 8x8 tile from the given coordinate of the image and returns its 16 bytes
func convertTile(image image.Image, colors ImageColors, startX int, startY int) ([16]byte, error) {
	tile := [16]byte{0}
	var err error = nil

	// TODO

	return tile, err
}
