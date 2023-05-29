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
	input, output, _ := readCmd()

	// read input file
	data, err := os.Open(input)
	if err != nil {
		fmt.Println(input, "- unable to open input:", err)
		os.Exit(2)
	}

	// decode the data
	img, _, err := image.Decode(data)
	if err != nil {
		fmt.Println(input, "- unable to decode input:", err)
		os.Exit(3)
	}

	// initialize list of colors
	colors := detectColors(img)

	convertTile(img, colors, 0, 0)

	// TODO initialze output buffer
	// TODO figure out color/luminosity detection
	// TODO read tiles in correct order
	// TODO write to output file

	fmt.Printf("%s -> %s (%s)\n", input, output, time.Since(start))
}
