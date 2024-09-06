package core

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	otherdraw "golang.org/x/image/draw"
)

func ConvertBlackWhite(inputImg string) {
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
			if uint(lum/256) < 0x80 {
				newImg.Set(x, y, color.RGBA{0, 0, 0, 255})
			} else {
				newImg.Set(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}

	outputFile, err := os.Create("./dist/bw_output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, newImg)
	if err != nil {
		panic(err)
	}
}

func ConvertGrayscale(inputImg string) {
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

	outputFile, err := os.Create("./dist/gray_output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, newImg)
	if err != nil {
		panic(err)
	}
}

func Resize(inputImg string, factor int) {
	file, err := os.Open(inputImg)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	newImg := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X*factor, img.Bounds().Max.Y*factor))

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

func ResizeAbsolute(inputImg string, length int, width int) {
	file, err := os.Open(inputImg)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	newImg := image.NewRGBA(image.Rect(0, 0, length, width))

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

// produces a mosaic of b&w version of target image
func MosaicBW(srcImageFile string, bImageFile string, wImageFile string) {
	// import images
	file, err := os.Open(srcImageFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	srcImage, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	file, err = os.Open(bImageFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bImage, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	file, err = os.Open(wImageFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	wImage, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// 50x50 is a good starting size for mosiac pieces
	bImageRsz := image.NewRGBA(image.Rect(0, 0, 100, 100))
	otherdraw.NearestNeighbor.Scale(bImageRsz, bImageRsz.Rect, bImage, bImage.Bounds(), otherdraw.Over, nil)

	wImageRsz := image.NewRGBA(image.Rect(0, 0, 100, 100))
	otherdraw.NearestNeighbor.Scale(wImageRsz, wImageRsz.Rect, wImage, wImage.Bounds(), otherdraw.Over, nil)

	// scale up targetImg to corresponding size
	finalImage := image.NewRGBA(image.Rect(0, 0, srcImage.Bounds().Max.X*100, srcImage.Bounds().Max.X*100))
	otherdraw.NearestNeighbor.Scale(finalImage, finalImage.Rect, srcImage, srcImage.Bounds(), otherdraw.Over, nil)

	// iterate over each "pixel" in original image
	// the corresponding size should be drawn in the new final image
	fmt.Println(bImageRsz.Bounds())
	fmt.Println(wImageRsz.Bounds().Size())
	fmt.Println(wImageRsz.Bounds().Dx())

	for x := range srcImage.Bounds().Dx() {
		for y := range srcImage.Bounds().Dy() {
			// iterate over each pixel in original image
			oldPixel := srcImage.At(x, y)
			r, g, b, _ := oldPixel.RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			// resized final image pixels should be placed accordingly based on x,y offset
			location := image.Rectangle{image.Point{x * 100, y * 100}, image.Point{x*100 + 100, y*100 + 100}}
			// fmt.Println(location)
			if uint(lum/256) < 0x80 {
				// newImg.Set(x, y, color.RGBA{0, 0, 0, 255})
				draw.Draw(finalImage, location, bImageRsz, image.Point{0, 0}, draw.Src)
			} else {
				// newImg.Set(x, y, color.RGBA{255, 255, 255, 255})
				draw.Draw(finalImage, location, wImageRsz, image.Point{0, 0}, draw.Src)
			}
		}
	}

	outputFile, err := os.Create("./dist/mosaic.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, finalImage)
	if err != nil {
		panic(err)
	}
}
