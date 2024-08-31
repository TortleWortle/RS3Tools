package main

import (
	"fmt"
	"image"
	"image/color"
)

// TODO: script to get all images from runescape wiki and convert them to png

// hold all possible buff types, load image cache and prep transparent masks etc

type BuffDetector struct {
}

type Buff struct {
	Name string
	// if it's a debuff
	Negative bool
}

func BuffsFromImage(img image.Image) []Buff {
	var buffs []Buff
	bufPoints := scanForSimilarColor(img, buffBorderColor)
	debufPoints := scanForSimilarColor(img, debuffBorderColor)

	bounds := img.Bounds()
	rec := image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	debugImage := image.NewRGBA(rec)

	for _, p := range debufPoints {
		debugImage.Set(p.X, p.Y, color.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 255,
		})
	}

	saveImage(debugImage, "debug.png")

	// get the tangles
	bufRectangles := getRectangles(img, bufPoints, buffBorderColor)
	debufRectangles := getRectangles(img, debufPoints, debuffBorderColor)
	// can't recognize buff names yet
	for i, bufR := range bufRectangles {
		newImage := img.(SubImager).SubImage(bufR)
		saveImage(newImage, fmt.Sprintf("buff-%d.png", i))
		buffs = append(buffs, Buff{
			Name:     detectBuffName(img, bufR),
			Negative: false,
		})
	}

	for i, bufR := range debufRectangles {
		newImage := img.(SubImager).SubImage(bufR)
		saveImage(newImage, fmt.Sprintf("debuff-%d.png", i))
		buffs = append(buffs, Buff{
			Name:     detectBuffName(img, bufR),
			Negative: true,
		})
	}

	return buffs
}

// use top half to detect what buff it is, use bottom half for text
func detectBuffName(img image.Image, rec image.Rectangle) string {
	// to compare buff icons: transparency map for origin and only compare points that are opaque
	// average color with tolerance per pixel?
	//r := image.Rect(0, 0, rec.Dx(), rec.Dy())
	//i := image.NewRGBA(r)
	i := img.(SubImager).SubImage(rec)
	//for x := 1; x < r.Dx()-1; x++ {
	//	for y := 1; y < r.Dy()/2; y++ {
	//		i.Set(x, y, img.At(rec.Min.X+x, rec.Min.Y+y))
	//	}
	//}
	saveImage(i, "test.png")
	return "Lol idk"
}
