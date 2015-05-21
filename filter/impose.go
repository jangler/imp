package filter

import (
	"errors"
	"image"
	"math"
	"strconv"

	"github.com/jangler/imp/util"
)

var imposeHelp = `impose <layer> <file> [<x> <y>]

Layer the working image on top of another image or vice versa. Possible values
for 'layer' are over and under. Coordinates x and y may be given to offset the
working image relative to the other image. If coordinates are not given, they
default to zero and the images are aligned at their top-left corners.`

func imposeFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
	if len(args) < 2 {
		util.Die(errors.New(imposeHelp))
	}
	order := args[0]
	if order != "over" && order != "under" {
		util.Die(errors.New(imposeHelp))
	}
	imposeImg := util.ReadImage(args[1])

	var xOffset, yOffset int64
	if len(args) >= 4 {
		var err error
		xOffset, err = strconv.ParseInt(args[2], 10, 0)
		if err == nil {
			yOffset, err = strconv.ParseInt(args[3], 10, 0)
			if err == nil {
				args = args[4:]
			} else {
				xOffset, yOffset = 0, 0
				args = args[2:]
			}
		} else {
			xOffset = 0
			args = args[2:]
		}
	} else {
		args = args[2:]
	}

	b1, b2 := img.Bounds(), imposeImg.Bounds()
	x1 := int(math.Min(0, float64(xOffset)))
	x2 := int(math.Max(float64(b2.Dx()), float64(b1.Dx()+int(xOffset))))
	y1 := int(math.Min(0, float64(yOffset)))
	y2 := int(math.Max(float64(b2.Dy()), float64(b1.Dy()+int(yOffset))))
	width, height := x2-x1, y2-y1
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	if order == "over" {
		util.DrawImg(imposeImg, newImg, int(math.Max(0, float64(-xOffset))),
			int(math.Max(0, float64(-yOffset))))
		util.DrawImg(img, newImg, int(math.Max(0, float64(xOffset))),
			int(math.Max(0, float64(yOffset))))
	} else {
		util.DrawImg(img, newImg, int(math.Max(0, float64(xOffset))),
			int(math.Max(0, float64(yOffset))))
		util.DrawImg(imposeImg, newImg, int(math.Max(0, float64(-xOffset))),
			int(math.Max(0, float64(-yOffset))))
	}

	return newImg, args
}

func init() {
	addFilter(&Filter{"impose", imposeHelp, imposeFunc})
}
