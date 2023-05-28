package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	input, _, _ := ReadCmd()

	// read input file
	data, err := os.Open(input)
	if err != nil {
		fmt.Println("unable to open input:", err)
		os.Exit(2)
	}

	img, _, err := image.Decode(data)
	if err != nil {
		fmt.Println("unable to decode input:", err)
		os.Exit(3)
	}

	fmt.Println(ConvertTile(img, 0, 0))

	// TODO initialze output buffer
	// TODO figure out color/luminosity detection
	// TODO read tiles in correct order
	// TODO write to output file

}
