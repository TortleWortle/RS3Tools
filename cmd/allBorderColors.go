package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	file, err := os.Open("output/debug.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.OpenFile("output/borderOut.png", os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		log.Fatalf("creating border out: %v", err)
	}
	defer newFile.Close()

	img, err := png.Decode(file)

	if err != nil {
		log.Fatalf("Decoding file: %v", err)
	}

	// rough compare or direct compare from set? idk
	buffBorderColors := make(map[color.Color]struct{})

	bounds := img.Bounds()
	rec := image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	newImage := image.NewRGBA(rec)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			if (y == 0) || x == 0 || x == bounds.Dx()-1 || y == bounds.Dy()-1 {
				c := img.At(x, y)
				newImage.Set(x, y, c)
			}
		}
	}

	for color := range buffBorderColors {
		fmt.Println(color)
	}

	err = png.Encode(newFile, newImage)
	if err != nil {
		log.Fatalf("Encoding file: %v", err)
	}
}
