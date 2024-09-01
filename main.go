package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	input, err := os.Open("input/screenshot.png")
	if err != nil {
		log.Fatalf("opening screenshot: %v", err)
	}
	defer input.Close()

	img, err := png.Decode(input)
	if err != nil {
		log.Fatalf("decoding screenshot: %v", err)
	}

	buffs := BuffsFromImage(img)

	fmt.Println(buffs)

	//debugFile, err := os.OpenFile("output/debugImage.png", os.O_WRONLY|os.O_CREATE, 0700)
	//if err != nil {
	//	log.Fatalf("creating border out: %v", err)
	//}
	//defer debugFile.Close()
	//
	//bounds := img.Bounds()
	//
	//rec := image.Rect(0, 0, bounds.Dx(), bounds.Dy())
	//debugImage := image.NewRGBA(rec)
	//
	//bufPoints := scanForSimilarColor(img, buffBorderColor)
	//
	//if len(bufPoints) == 0 {
	//	fmt.Printf("No buffpoints\n")
	//}
	//
	//// take subimage of either each buff or all buffs together
	//bufRectangles := getRectangles(img, bufPoints)
	//
	//for i, rec := range bufRectangles {
	//	newImage := img.(SubImager).SubImage(rec)
	//	saveImage(newImage, fmt.Sprintf("buff-%d.png", i))
	//}
	//
	//err = png.Encode(debugFile, debugImage)
	//if err != nil {
	//	log.Fatalf("Encoding debug file: %v", err)
	//}
	//
	//err = cropAndSave(img, 1200, 1200, 256, 256, "test.png")
	//if err != nil {
	//	log.Fatalf("cropping and saving: %v", err)
	//}
}
