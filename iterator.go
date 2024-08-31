package main

import (
	"image"
	"iter"
)

func iteratePixels(img image.Image) iter.Seq[image.Point] {
	return func(yield func(image.Point) bool) {
		for x := 0; x < img.Bounds().Dx(); x++ {
			for y := 0; y < img.Bounds().Dy(); y++ {
				if !yield(image.Point{
					X: x,
					Y: y,
				}) {
					return
				}
			}
		}
	}
}

func iterateOpaquePixels(img image.Image) iter.Seq[image.Point] {
	return func(yield func(image.Point) bool) {
		for p := range iteratePixels(img) {
			c := img.At(p.X, p.Y)
			_, _, _, a := c.RGBA()
			if a == 0xffff {
				if !yield(p) {
					return
				}
			}
		}
	}
}
