package filters

import (
	"errors"
	"image"
	"image/color"

	"github.com/jangler/imp/util"
)

var maskHelp = `mask file

The mask filter applies the alpha channel from image 'file' to the working
image. The images must have the same dimensions.`

// Filter function.
func mask(img *image.RGBA, args []string) []string {
	if len(args) == 0 {
		util.Die(errors.New(maskHelp))
	}
	maskImg := util.ReadImage(args[0])

	b1 := img.Bounds()
	b2 := maskImg.Bounds()
	if b1.Max.X-b1.Min.X != b2.Max.X-b2.Min.X ||
		b1.Max.Y-b1.Min.Y != b2.Max.Y-b2.Min.Y {
		util.Die(errors.New("mismatched dimensions for mask image: " +
			args[0]))
	}
	dx, dy := b2.Min.X-b1.Min.X, b2.Min.Y-b1.Min.Y
	for y := b1.Min.Y; y < b1.Max.Y; y++ {
		for x := b1.Min.X; x < b1.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			_, _, _, a := maskImg.At(x+dx, y+dy).RGBA()
			img.SetRGBA(x, y, color.RGBA{
				uint8(r >> 8),
				uint8(g >> 8),
				uint8(b >> 8),
				uint8(a >> 8),
			})
		}
	}

	args = args[1:]
	return args
}
