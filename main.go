package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

var width int = 140
var cellSide int = 10

func openAndDecodeImage(path string) (image.Image, image.Rectangle, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, image.Rectangle{}, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		fmt.Println("Error decoding file:", err)
	}

	bounds := img.Bounds()

	return img, bounds, err
}

func main() {
	img, _, err := openAndDecodeImage("./image.jpg")

	if err != nil {
		panic(err)
	}

	px := img.At(1200, 1200)

	r, g, b, _ := px.RGBA()

	luminance := math.Round(0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b))

	fmt.Println(r, g, b, luminance)

}
