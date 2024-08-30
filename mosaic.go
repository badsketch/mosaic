package main

import (
	"github.com/badsketch/mosaic/core"
)

func main() {
	core.ConvertGrayscale("./static/input2.png")
	core.ConvertBlackWhite("./static/input2.png")
	// core.Resize(("./static/input.png"))
	core.ResizeAbsolute("./static/input.png", 50, 50)
}
