package filter

import (
	"errors"
	"image"
	"image/color"
	"math"
	"strconv"

	"github.com/jangler/imp/util"
)

var lumHelp = `lum factor [gFactor bFactor [aFactor]]

Multiply the brightness of the image by the given factor. Individual factors
may be given for each channel.`

func lumFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
	numArgs := 0
	factor := []float64{1.0, 1.0, 1.0, 1.0}
	for numArgs < 4 && len(args) > 0 {
		f, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			break
		}
		factor[numArgs] = f
		args = args[1:]
		numArgs++
	}
	if numArgs != 1 && numArgs != 3 && numArgs != 4 {
		util.Die(errors.New(lumHelp))
	}
	if numArgs == 1 {
		factor[1] = factor[0]
		factor[2] = factor[0]
	}

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			c := color.RGBA{
				uint8(math.Min(255, float64(r>>8)*factor[0])),
				uint8(math.Min(255, float64(g>>8)*factor[1])),
				uint8(math.Min(255, float64(b>>8)*factor[2])),
				uint8(math.Min(255, float64(a>>8)*factor[3])),
			}
			img.SetRGBA(x, y, c)
		}
	}

	return img, args
}

func init() {
	addFilter(&Filter{"lum", lumHelp, lumFunc})
}
