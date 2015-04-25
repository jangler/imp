package filter

import (
	"errors"
	"image"
	"math"
	"strconv"

	"github.com/jangler/imp/util"
)

var scaleHelp = `scale w h

Scale the image to the given width and height in pixels, using nearest-neighbor
interpolation. If an argument ends in %, it is interpreted as a percentage of
the working image's dimension instead of a pixel count. Negative values can be
used to flip the image horizontally and/or vertically.`

func scaleFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
	if len(args) < 2 {
		util.Die(errors.New(scaleHelp))
	}
	dim := make([]int, 2)
	b1 := img.Bounds()
	for i := 0; i < 2; i++ {
		if args[0][len(args[0])-1] == '%' {
			p, err := strconv.ParseFloat(args[0][:len(args[0])-1], 64)
			if err != nil {
				util.Die(errors.New(scaleHelp))
			}
			if i == 0 {
				dim[i] = int(p / 100 * float64(b1.Dx()))
			} else {
				dim[i] = int(p / 100 * float64(b1.Dy()))
			}
		} else {
			n, err := strconv.ParseInt(args[0], 10, 0)
			if err != nil {
				util.Die(errors.New(scaleHelp))
			}
			dim[i] = int(n)
		}
		args = args[1:]
	}

	newImg := image.NewRGBA(image.Rect(0, 0,
		int(math.Abs(float64(dim[0]))), int(math.Abs(float64(dim[1])))))
	b2 := newImg.Bounds()
	xFactor := float64(b1.Dx()) / float64(b2.Dx())
	yFactor := float64(b1.Dy()) / float64(b2.Dy())
	for y := 0; y < b2.Dy(); y++ {
		for x := 0; x < b2.Dx(); x++ {
			oldX := b1.Min.X + int(float64(x)*xFactor)
			oldY := b1.Min.Y + int(float64(y)*yFactor)
			color := img.At(oldX, oldY)
			newX := b2.Min.X + x
			if dim[0] < 1 {
				newX = -dim[0] - newX
			}
			newY := b2.Min.Y + y
			if dim[1] < 1 {
				newY = -dim[1] - newY
			}
			newImg.Set(newX, newY, color)
		}
	}

	return newImg, args
}

func init() {
	addFilter(&Filter{"scale", scaleHelp, scaleFunc})
}
