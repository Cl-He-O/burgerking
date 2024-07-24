package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	golden_ratio := 2 / (math.Sqrt(5) - 1)

	factor := 160
	palette := [...]color.RGBA{
		{R: 252, G: 148, B: 59, A: 255},
		{R: 157, G: 209, B: 7, A: 255},
		{R: 254, G: 223, B: 106, A: 255},
		{R: 254, G: 70, B: 58, A: 255},
		{R: 128, G: 73, B: 11, A: 255},
		{R: 252, G: 148, B: 59, A: 255}}

	height := factor * 6
	width := int(math.Ceil(float64(height) * golden_ratio))

	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: width, Y: height}})

	for i := range height {
		for j := range width {
			img.SetRGBA(j, i, palette[i/factor])
		}
	}

	fp, err := os.Create("burgerking.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(fp, img)

	if err != nil {
		panic(err)
	}

	fmt.Printf("finished writing png\n")
}
