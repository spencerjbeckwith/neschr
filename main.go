package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"time"
)

func main() {
	start := time.Now()
	input, output, mode := readCmd()

	// Read input file
	data, err := os.Open(input)
	if err != nil {
		fmt.Printf("%s - unable to open input: %s\n", input, err)
		os.Exit(2)
	}

	// Decode the data
	img, _, err := image.Decode(data)
	if err != nil {
		fmt.Printf("%s - unable to decode input: %s\n", input, err)
		os.Exit(3)
	}

	// Determine width/height and make sure they're valid
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	if width%8 != 0 || height%8 != 0 {
		fmt.Printf("%s - image size must be a multiple of 8. Got %d x %d\n", input, width, height)
		os.Exit(4)
	}

	total := (width * height) / 64
	var buffer []byte

	// Initialize list of colors
	colors := detectColors(img)

	// Scan each tile of the image
	if mode == "horizontal" {
	outH:
		for y := 0; y < height; y += 8 {
			for x := 0; x < width; x += 8 {
				data, err2 := convertTile(img, colors, x, y)
				if err2 != nil {
					err = err2
					break outH
				}
				buffer = append(buffer, data[:]...)
			}
		}
	} else if mode == "vertical" {
	outV:
		for x := 0; x < width; x += 8 {
			for y := 0; y < height; y += 8 {
				data, err2 := convertTile(img, colors, x, y)
				if err2 != nil {
					err = err2
					break outV
				}
				buffer = append(buffer, data[:]...)
			}
		}
	}

	if err != nil {
		fmt.Printf("%s - unable to convert: %s\n", input, err)
		os.Exit(5)
	}

	// Write the output file
	err = os.WriteFile(output, buffer, 0666)
	if err != nil {
		fmt.Printf("%s - unable to save %s: %s", input, output, err)
		os.Exit(6)
	}

	fmt.Printf("%s -> %s (%d tiles in %s)\n", input, output, total, time.Since(start))
}
