package filter

import (
	"errors"
	"image"
	"image/color"
	"math"
	"strconv"

	"github.com/jangler/imp/util"
)

var satHelp = `sat <factor>

Multiply the saturation of the current image by the given factor. A negative
factor will effectively invert colors.`

func satFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
	if len(args) < 1 {
		util.Die(errors.New(satHelp))
	}
	factor, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		util.Die(err)
	}

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			mean := (int64(r) + int64(g) + int64(b)) / 3
			dR := float64(int64(r) - mean)
			dG := float64(int64(g) - mean)
			dB := float64(int64(b) - mean)
			c := color.RGBA{
				uint8(math.Min(float64(a>>8), math.Max(0,
					float64((mean+int64(dR*factor))>>8)))),
				uint8(math.Min(float64(a>>8), math.Max(0,
					float64((mean+int64(dG*factor))>>8)))),
				uint8(math.Min(float64(a>>8), math.Max(0,
					float64((mean+int64(dB*factor))>>8)))),
				uint8(a >> 8),
			}
			img.SetRGBA(x, y, c)
		}
	}

	return img, args[1:]
}

func init() {
	addFilter(&Filter{"sat", satHelp, satFunc})
}
