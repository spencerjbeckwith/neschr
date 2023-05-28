package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Reads and verifies incoming CLI arguments
func ReadCmd() (input string, output string, mode string) {
	// Parse incoming args
	inputPtr := flag.String("i", "", "input file (.png)")
	outputPtr := flag.String("o", "", "output file (.chr)")
	modePtr := flag.String("m", "horizontal", "horizontal, vertical, or nametable - default to horizontal")
	flag.Parse()

	// Ensure validity
	input = strings.ToLower(*inputPtr)
	output = strings.ToLower(*outputPtr)
	mode = strings.ToLower(*modePtr)

	if input == "" || output == "" {
		fmt.Println("neschr")
		fmt.Println("Ex: neschr -i input.png -o output.chr -m horizontal")
		flag.PrintDefaults()
		fmt.Println("input and output are required")
		os.Exit(0)
	}

	if mode != "horizontal" && mode != "vertical" && mode != "nametable" {
		fmt.Println("mode must be one of horizontal, vertical, or nametable")
		os.Exit(1)
	}

	return input, output, mode
}
