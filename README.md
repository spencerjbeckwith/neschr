# neschr

This command-line tool is meant to greatly simplify the process of creating CHR data for NES ROMs. Never again will you need to painstakingly enter bits and hope you got them right (like I did once... yikes), or learn unwieldly new tools that can't integrate cleanly in the build process and that you're just too lazy to open, save, and close for every miniscule change. Instead, just provide an image file with width and height as multiples of 8 and neschr will break the image apart into tiles, automatically detect the four colors, and output the bytes as a .chr file to be directly included in your ROM. This can even be done automatically in your build step.

Note that this does not handle palettes, nametables, or how the data is handled at all once you build it into your ROM. You'll want to make sure you're identifying the tiles somehow and that you aren't overflowing too many tiles in ROM or RAM at once - you gotta figure that out on your own :)

# Installation

If you feel like trusting me today, you can download the Linux executable `neschr` directly from this repository. If you're using a different operating system, or you don't want to trust some random executable in some random repository, you can build from the source instead.

## Build from source

To build from source, you must have Go installed. Clone the repository and run `go build .`

Once you have the executable, you can place it in your ROM's directory or any location on `PATH`, which will allow your build tools to use the executable.

# Usage

`neschr -i test/big.png -o test/big.chr`

`neschr -i test/big.png -o test/bigv.chr -m vertical`

Running the command without any flags will print a help message.

This repository contains several test images you can run to see the .chr output.

# Options

- `-i <input file>` required - input file. Should be a .png image preferably. JPEG may work but I haven't tested it.
- `-o <output file>` required - output file. Should be a .chr. Will overwrite the file if it already exists
- `-m <mode>` output mode - see below

# Modes

Two separate output modes are available: horizontal and vertical. This determines if the input image is read and output in a row-then-column order (horizontal) or column-then-row order (vertical). The choice of mode will depend on how you prefer to layout your tiles in memory.

1. `-m horizontal` horizontal mode: placement of each tile in the output goes across each row of the input, slicing it horizontally. I find this to be most intuitive.
2. `-m vertical` vertical mode: placement of each tile in the output goes down each column of the input, slicing it vertically

# So what's the big deal? Where's my GUI?

Ain't nobody got time for GUIs in the thrilling and highly active world of NES game development. Jokes aside though I prefer the convenience of CLI tools especially when working on big projects with a lot of files.

What's cool about this tool is you can integrate it in your toolchain no sweat. For example, you could use a `Makefile` that will automatically rebuild your .chr files (and even your entire game) when your input .png is modified. There's a million ways it could be done but this is the one I had in mind.

Oh yeah, and this tool is really fast. You can output hundreds or maybe thousands of tiles in the range of single-digit milliseconds.

# License

lmao do literally anything idc