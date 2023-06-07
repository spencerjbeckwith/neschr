package main

import (
	"image"
	"image/color"
	"sort"
)

func toLuminosity(color color.Color) float32 {
	r, g, b, _ := color.RGBA()
	return (0.299 * float32(r)) + (0.587 * float32(g)) + (0.114 * float32(b))
}

// Sort.Interface for luminosity
type ImageColors [4]color.Color

func (a *ImageColors) Len() int {
	return 4
}
func (a *ImageColors) Less(i, j int) bool {
	if a[i] == nil || a[j] == nil { // Move nils to the end of the list
		return false
	}
	return toLuminosity(a[i]) < toLuminosity(a[j])
}
func (a *ImageColors) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Returns the four colors of this image to be composited into CHR format. All colors with an alpha value of 0 are treated as equivalent.
//
// If there are not four colors, the colors in the image are in binary order for CHR: 00 (neither mask), 01 (first mask), 10 (second mask), and 11 (both masks)
// The colors in the image are ordered based on their luminosity.
//
// Returns an error if there are more than four colors.
func detectColors(image image.Image) ImageColors {
	colors := ImageColors{}
	bounds := image.Bounds()

out:
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// If color at this position is new, add it to the list
			col := image.At(x, y)

			// Normalize pixels w/ 0 alpha into black
			_, _, _, a := col.RGBA()
			if a == 0 {
				col = color.RGBA{0, 0, 0, 0}
			}

			for c := 0; c < 4; c++ {
				// If this color is already in our array, break
				if colors[c] != nil && colorsMatch(col, colors[c]) {
					break
				}

				// No color at this index yet
				if colors[c] == nil {
					colors[c] = col
					if c == 3 {
						// Found our fourth color, stop scanning image
						break out
					}

					// Don't write this color to the rest of the array
					break
				}
			}
		}
	}

	sort.Sort(&colors)
	return colors
}

// Returns if two colors are the same
func colorsMatch(a color.Color, b color.Color) bool {
	if a == nil || b == nil {
		return false
	}
	aR, aG, aB, aA := a.RGBA()
	bR, bG, bB, bA := b.RGBA()
	return aR == bR && aG == bG && aB == bB && aA == bA
}
