package filter

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"math"
	"strconv"

	"github.com/jangler/imp/util"
)

var blurHelp = `blur radius power

TODO`

func blurFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
	if len(args) < 2 {
		util.Die(errors.New(blurHelp))
	}
	radius, err := strconv.ParseFloat(args[0], 64)
	rad := int(radius)
	if err != nil {
		util.Die(errors.New(blurHelp))
	}
	pow, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		util.Die(errors.New(blurHelp))
	}
	_, _ = rad, pow

	b := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	util.DrawImg(img, newImg, 0, 0)

	mask := image.NewRGBA(image.Rect(0, 0, rad*2, rad*2))
	mb := mask.Bounds()
	for y := mb.Min.Y; y < mb.Max.Y; y++ {
		for x := mb.Min.X; x < mb.Max.X; x++ {
			dist := math.Hypot(float64(x)+0.5-float64(mb.Min.X+rad),
				float64(y)+0.5-float64(mb.Min.Y+rad))
			val := uint8(math.Max(0,
				math.Min(255, 255*pow*(radius-dist)/radius)))
			mask.Set(x, y, color.RGBA{0, 0, 0, val})
		}
	}

	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			r := image.Rect(x-rad, y-rad, x+rad, y+rad)
			c := img.At(x+b.Min.X, y+b.Min.Y)
			draw.DrawMask(newImg, r, &image.Uniform{c}, image.ZP,
				mask, mb.Min, draw.Over)
		}
	}

	return newImg, args[2:]
}

func init() {
	addFilter(&Filter{"blur", blurHelp, blurFunc})
}
