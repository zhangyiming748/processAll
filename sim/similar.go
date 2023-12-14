package sim

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

func Similar(p1, p2 string) {
	img1, err := loadImage(p1)
	if err != nil {
		fmt.Println("Failed to load image1:", err)
		return
	}

	img2, err := loadImage(p2)
	if err != nil {
		fmt.Println("Failed to load image2:", err)
		return
	}

	similarity := calculateSimilarity(img1, img2)
	fmt.Println("Similarity:", similarity)
}

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func calculateSimilarity(img1, img2 image.Image) float64 {
	bounds := img1.Bounds()
	totalPixels := bounds.Dx() * bounds.Dy()
	diffPixels := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()

			diffPixels += int(math.Abs(float64(r1-r2)/0xffff*255)) +
				int(math.Abs(float64(g1-g2)/0xffff*255)) +
				int(math.Abs(float64(b1-b2)/0xffff*255))
		}
	}

	return 1 - float64(diffPixels)/(3*255*float64(totalPixels))
}
