package main

import (
	"golang.org/x/image/draw"
	"image"
)

func overlapPercent(in, buff image.Image) (float64, error) {
	activeTop := topOnly(scale(in))
	buffTop := topOnly(scale(buff))

	pixels := 0
	overlap := 0

	for p := range iterateOpaquePixels(buffTop) {
		bc := buffTop.At(p.X, p.Y)
		ac := activeTop.At(p.X, p.Y)

		pixels++
		if colorRoughEquals(bc, ac) {
			overlap++
		}
	}

	percent := (float64(overlap) / float64(pixels))
	// create mask for only transparent pixels of the buff
	return percent, nil
}

func scale(img image.Image) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, 32, 32))
	draw.NearestNeighbor.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)
	return dst
}

func topOnly(input image.Image) image.Image {
	cropSize := image.Rect(0, 0, input.Bounds().Dx(), input.Bounds().Dy()/2)

	return input.(SubImager).SubImage(cropSize)
}
