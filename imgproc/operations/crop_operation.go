package operations

import (
	"fmt"
	"image"
)

type CropOperation struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (c CropOperation) Apply(img image.Image) (image.Image, error) {
	bounds := img.Bounds()
	// Verify that the crop rectangle is within the image bounds
	if c.X < bounds.Min.X || c.Y < bounds.Min.Y || c.X+c.Width > bounds.Max.X || c.Y+c.Height > bounds.Max.Y {
		return nil, fmt.Errorf("crop rectangle is out of bounds max %v, min %v, crop %v", bounds.Max, bounds.Min, c)
	}

	// Define the new rectangle
	rect := image.Rect(c.X, c.Y, c.X+c.Width, c.Y+c.Height)

	// Create a new image with the new rectangle
	dst := image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			dst.Set(x-c.X, y-c.Y, img.At(x, y))
		}
	}

	return dst, nil
}

func (c CropOperation) Name() string {
	return fmt.Sprintf("crop(%d,%d,%d,%d)", c.X, c.Y, c.Width, c.Height)
}
