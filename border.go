package main

import (
	"image"
	"image/color"
)

const tolerance = 50 << 8

var buffBorderColor = color.RGBA{
	R: 90,
	G: 150,
	B: 25,
	A: 0xff,
}

var debuffBorderColor = color.RGBA{
	R: 204,
	G: 0,
	B: 0,
	A: 0xff,
}

// Whether a roughly compares to b, using const tolerance
func req(a, b uint32) bool {
	tbMin := b - tolerance
	if tbMin > b {
		tbMin = 0
	}
	tbMax := b + tolerance
	if tbMax < b {
		tbMax = 0xffffffff
	}
	return a >= tbMin && a <= tbMax
}

func colorRoughEquals(a, b color.Color) bool {
	ar, ag, ab, aa := a.RGBA()
	br, bg, bb, ba := b.RGBA()
	return req(ar, br) && req(ag, bg) && req(ab, bb) && req(aa, ba)
}

func getRectangles(img image.Image, points []image.Point, clr color.Color) []image.Rectangle {
	var corners []image.Rectangle
	for _, point := range points {
		x, y := point.X, point.Y
		isCorner, width, height := isCorner(img, x, y, clr)
		if isCorner {
			corners = append(corners, image.Rect(x, y, x+width, y+height))
		}
	}
	return corners
}

func scanForSimilarColor(img image.Image, clr color.Color) []image.Point {
	bounds := img.Bounds()
	var points []image.Point
	// scan every pixel and check for buff colors and save their coordinates
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			// lets compare baby
			c := img.At(x, y)
			if colorRoughEquals(c, clr) {
				points = append(points, image.Point{
					X: x,
					Y: y,
				})
			}
		}
	}
	return points
}

// bases it off the starting pixel, should probably not do that
func isCorner(img image.Image, x, y int, clr color.Color) (isCorner bool, width, height int) {
	start := clr
	bounds := img.Bounds()
	width = 0
	// check how far right we can go and how far down we can go, if they equal it's likely a corner
	for currentX := x; currentX < bounds.Dx(); currentX++ {
		// go until we reach the end or a color that doesn't match
		if !colorRoughEquals(img.At(currentX, y), start) {
			break
		}
		width++
	}
	height = 0
	// check how far right we can go and how far down we can go, if they equal it's likely a corner
	for currentY := y; currentY < bounds.Dy(); currentY++ {
		// go until we reach the end or a color that doesn't match
		if !colorRoughEquals(img.At(x, currentY), start) {
			break
		}
		height++
	}

	// we found ourself a corner probably
	// TODO: allow scaling images, this requires a resizing step somewhere down the line
	if width == 32 && width == height {
		return true, width, height
	}

	return false, 0, 0
}
