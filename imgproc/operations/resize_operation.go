package operations

import (
	"fmt"
	"image"

	"github.com/nfnt/resize"
)

type ResizeOperation struct {
	Width  int
	Height int
}

func (r ResizeOperation) Apply(img image.Image) (image.Image, error) {
	resizedImg := resize.Resize(uint(r.Width), uint(r.Height), img, resize.Lanczos3)
	return resizedImg, nil
}

func (r ResizeOperation) Name() string {
	return fmt.Sprintf("resize(%dx%d)", r.Width, r.Height)
}
