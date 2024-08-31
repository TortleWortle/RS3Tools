package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path"
)

// grab rectangle from screenshot and save to output
func cropAndSave(img image.Image, x, y, width, height int, filename string) error {
	cropSize := image.Rect(0, 0, width, height)
	cropSize = cropSize.Add(image.Point{
		X: x,
		Y: y,
	})

	newImage := img.(SubImager).SubImage(cropSize)
	filepath := path.Join("output", filename)
	newFile, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		return fmt.Errorf("opening new file(%s): %w", filepath, err)
	}

	err = png.Encode(newFile, newImage)
	if err != nil {
		return fmt.Errorf("encoding image(%s): %w", filepath, err)
	}
	return nil
}

func saveImage(img image.Image, filename string) error {
	filepath := path.Join("output", filename)
	newFile, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		return fmt.Errorf("opening new file(%s): %w", filepath, err)
	}

	err = png.Encode(newFile, img)
	if err != nil {
		return fmt.Errorf("encoding image(%s): %w", filepath, err)
	}
	return nil
}
