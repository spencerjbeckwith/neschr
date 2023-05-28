# neschr

This command-line tool is meant to greatly simplify the process of creating CHR data for NES ROMs. Never again will you need to painstaking enter bits and hope you got them right, or learn obscure new tools that can't integrate in the build process. Provide an image file with width and height as multiples of 8 (or an entire nametable) and neschr will break the image apart into tiles, automatically detect the four colors, and output the bytes as a .chr file to be directly included in your ROM.

Note that this does not handle palettes, nametables, or how the data is handled at all once you build it into your ROM.

# Installation

...

# Usage

...

# Options

...

# Modes

Three modes are available for neschr. Horizontal and vertical modes are intended for individual entity "sprites" while nametable mode is intended for full tilesets or environments.

1. `--horizontal` `-h` mode: placement of each tile in the output goes across each row of the input, slicing it horizontally
2. `--vertical` `-v` mode: placement of each tile in the output goes down each column of the input, slicing it vertically
3. `--nametable` `-n` mode: an entire nametable is output as-is. The image must be sized 256x256 pixels