package main

import (
	"fmt"
	"image"
	gif "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input, output, mode, debug := readCmd()

	// Read input file
	isGif := strings.HasSuffix(strings.ToLower(input), ".gif")
	data, err := os.Open(input)
	if err != nil {
		fmt.Printf("%s - unable to open input: %s\n", input, err)
		os.Exit(2)
	}

	var images []image.Image

	// Decode the data
	if isGif {

		// For gifs, every frame becomes an image
		myGif, err := gif.DecodeAll(data)
		if err != nil {
			fmt.Printf("%s - unable to decode gif input: %s\n", input, err)
			os.Exit(7)
		}
		for _, paletted := range myGif.Image {
			images = append(images, image.Image(paletted))
		}
		if len(images) <= 0 {
			fmt.Printf("%s - gif has no frames\n", input)
			os.Exit(8)
		}

	} else {

		// For non-gifs, just do one frame
		img, _, err := image.Decode(data)
		if err != nil {
			fmt.Printf("%s - unable to decode input: %s\n", input, err)
			os.Exit(3)
		}
		images = make([]image.Image, 1)
		images[0] = img

	}

	// Determine width/height and make sure they're valid
	bounds := images[0].Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	if width%8 != 0 || height%8 != 0 {
		fmt.Printf("%s - image size must be a multiple of 8. Got %d x %d\n", input, width, height)
		os.Exit(4)
	}

	total := (width * height) / 64
	var buffer []byte

	for _, img := range images {
		// Initialize list of colors
		colors := detectColors(img, debug)

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
	}

	// Write the output file
	err = os.WriteFile(output, buffer, 0666)
	if err != nil {
		fmt.Printf("%s - unable to save %s: %s", input, output, err)
		os.Exit(6)
	}

	fmt.Printf("%s -> %s (%d frames of %d tiles in %s)\n", input, output, len(images), total, time.Since(start))
}
