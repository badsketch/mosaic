package core

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	otherdraw "golang.org/x/image/draw"
)

func ProduceBlackWhite(inputImg string) {
	file, err := os.Open(inputImg)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg, newImg.Bounds(), img, image.Point{}, draw.Src)

	for x := range img.Bounds().Dx() + 2 {
		for y := range img.Bounds().Dy() + 2 {
			oldPixel := img.At(x, y)
			r, g, b, _ := oldPixel.RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			newPixel := color.Gray{uint8(lum / 256)}
			newImg.Set(x, y, newPixel)
		}
	}

	outputFile, err := os.Create("../dist/output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, newImg)
	if err != nil {
		panic(err)
	}

}

func Resize(inputImg string) {
	file, err := os.Open(inputImg)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	newImg := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X*2, img.Bounds().Max.Y*2))

	otherdraw.NearestNeighbor.Scale(newImg, newImg.Rect, img, img.Bounds(), otherdraw.Over, nil)

	outputFile, err := os.Create("./dist/scaled_output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, newImg)
	if err != nil {
		panic(err)
	}
}
