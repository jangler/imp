package filter

import (
	"errors"
	"image"
	"strconv"

	"github.com/jangler/imp/util"
)

var rotateHelp = `rotate <degrees>

Rotate the image clockwise; 'degrees' must be a multiple of 90.`

func rotateFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
	if len(args) == 0 {
		util.Die(errors.New(rotateHelp))
	}
	degrees, err := strconv.ParseInt(args[0], 10, 0)
	if err != nil {
		util.Die(errors.New(rotateHelp))
	}

	var newImg *image.RGBA
	b1 := img.Bounds()

	switch (degrees%360 + 360) % 360 {
	case 0:
		newImg = img
	case 90:
		newImg = image.NewRGBA(image.Rect(0, 0, b1.Dy(), b1.Dx()))
		b2 := newImg.Bounds()
		for y := 0; y < b1.Dy(); y++ {
			for x := 0; x < b1.Dx(); x++ {
				color := img.At(b1.Min.X+x, b1.Min.Y+y)
				newImg.Set(b2.Max.X-y, b2.Min.Y+x, color)
			}
		}
	case 180:
		newImg = image.NewRGBA(image.Rect(0, 0, b1.Dx(), b1.Dy()))
		b2 := newImg.Bounds()
		for y := 0; y < b1.Dy(); y++ {
			for x := 0; x < b1.Dx(); x++ {
				color := img.At(b1.Min.X+x, b1.Min.Y+y)
				newImg.Set(b2.Max.X-x, b2.Max.Y-y, color)
			}
		}
	case 270:
		newImg = image.NewRGBA(image.Rect(0, 0, b1.Dy(), b1.Dx()))
		b2 := newImg.Bounds()
		for y := 0; y < b1.Dy(); y++ {
			for x := 0; x < b1.Dx(); x++ {
				color := img.At(b1.Min.X+x, b1.Min.Y+y)
				newImg.Set(b2.Min.X+y, b2.Max.Y-x, color)
			}
		}
	default:
		util.Die(errors.New(rotateHelp))
	}

	return newImg, args[1:]
}

func init() {
	addFilter(&Filter{"rotate", rotateHelp, rotateFunc})
}
