package filter

import (
	"errors"
	"image"
	"image/color"

	"github.com/jangler/imp/util"
)

var maskHelp = `mask file

Multiply the alpha channel of the working image by the alpha channel from
another image file. The images must have the same dimensions.`

func maskFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
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
			r, g, b, a1 := img.At(x, y).RGBA()
			_, _, _, a2 := maskImg.At(x+dx, y+dy).RGBA()
			a := float64(a1) / 0xffff * float64(a2) / 0xffff
			img.SetRGBA(x, y, color.RGBA{
				uint8(r >> 8),
				uint8(g >> 8),
				uint8(b >> 8),
				uint8(a * 0xff),
			})
		}
	}

	return img, args[1:]
}

func init() {
	addFilter(&Filter{"mask", maskHelp, maskFunc})
}
