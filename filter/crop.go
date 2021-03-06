package filter

import (
	"errors"
	"image"
	"strconv"

	"github.com/jangler/imp/util"
)

var cropHelp = `crop <x> <y> <w> <h>

Resize the image canvas to start at (x, y) pixels relative to the top-left
corner of the working image, and be w by h pixels in size. It is possible to
enlarge the canvas in this way; pixels that were beyond the borders of the
working image will be blank. If an argument ends in %, it is interpreted as a
percentage of the working image's dimension instead of a pixel count.`

func cropFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
	numArgs := 0
	dim := make([]int, 4)
	b1 := img.Bounds()
	for numArgs < 4 && len(args) > 0 {
		if args[0][len(args[0])-1] == '%' {
			p, err := strconv.ParseFloat(args[0][:len(args[0])-1], 64)
			if err != nil {
				break
			}
			if numArgs%2 == 0 {
				dim[numArgs] = int(p / 100 * float64(b1.Dx()))
			} else {
				dim[numArgs] = int(p / 100 * float64(b1.Dy()))
			}
		} else {
			n, err := strconv.ParseInt(args[0], 10, 0)
			if err != nil {
				break
			}
			dim[numArgs] = int(n)
		}
		args = args[1:]
		numArgs++
	}
	if numArgs != 4 {
		util.Die(errors.New(cropHelp))
	}

	newImg := image.NewRGBA(image.Rect(0, 0, dim[2], dim[3]))
	b2 := newImg.Bounds()
	for y := 0; y < dim[3]; y++ {
		for x := 0; x < dim[2]; x++ {
			color := img.At(b1.Min.X+dim[0]+x, b1.Min.Y+dim[1]+y)
			newImg.Set(b2.Min.X+x, b2.Min.Y+y, color)
		}
	}

	return newImg, args
}

func init() {
	addFilter(&Filter{"crop", cropHelp, cropFunc})
}
