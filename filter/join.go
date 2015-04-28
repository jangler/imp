package filter

import (
	"errors"
	"image"
	"math"

	"github.com/jangler/imp/util"
)

var joinHelp = `join file edge [align]

Adjoin another image to an edge of the working image. Possible values for 'edge'
are top, bottom, left, and right. The 'align' argument is used to control which
edge of the resulting image the adjoined images are flush with. Possible values
for 'align' are top, bottom, left, right, and center. The default align is top
when 'edge' is left or right, or left when 'edge' is top or bottom.`

func joinFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
	if len(args) < 2 {
		util.Die(errors.New(joinHelp))
	}
	joinImg := util.ReadImage(args[0])
	edge := args[1]
	if edge != "top" && edge != "bottom" && edge != "left" && edge != "right" {
		util.Die(errors.New(joinHelp))
	}
	var gravity string
	if edge == "top" || edge == "bottom" {
		gravity = "left"
	} else {
		gravity = "top"
	}
	if len(args) >= 3 {
		if args[2] != "top" && args[2] != "bottom" && args[2] != "left" &&
			args[2] != "right" && args[2] != "center" {
			args = args[2:]
		} else {
			gravity = args[2]
			args = args[3:]
		}
	} else {
		args = args[2:]
	}

	b1, b2 := img.Bounds(), joinImg.Bounds()
	var width, height int
	if edge == "top" || edge == "bottom" {
		width = int(math.Max(float64(b1.Dx()), float64(b2.Dx())))
		height = b1.Dy() + b2.Dy()
	} else {
		width = b1.Dx() + b2.Dx()
		height = int(math.Max(float64(b1.Dy()), float64(b2.Dy())))
	}
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	b3 := newImg.Bounds()

	var xOffset1, yOffset1, xOffset2, yOffset2 int
	switch edge {
	case "top", "bottom":
		switch gravity {
		case "left":
			xOffset1 = 0
			xOffset2 = 0
		case "right":
			xOffset1 = width - b1.Dx()
			xOffset2 = width - b2.Dx()
		default:
			xOffset1 = (width - b1.Dx()) / 2
			xOffset2 = (width - b2.Dx()) / 2
		}
	case "left", "right":
		switch gravity {
		case "top":
			yOffset1 = 0
			yOffset2 = 0
		case "bottom":
			yOffset1 = height - b1.Dy()
			yOffset2 = height - b2.Dy()
		default:
			yOffset1 = (height - b1.Dy()) / 2
			yOffset2 = (height - b2.Dy()) / 2
		}
	}
	switch edge {
	case "top":
		yOffset1, yOffset2 = b2.Dy(), 0
	case "bottom":
		yOffset1, yOffset2 = 0, b1.Dy()
	case "left":
		xOffset1, xOffset2 = b2.Dx(), 0
	case "right":
		xOffset1, xOffset2 = 0, b1.Dx()
	}

	for y := 0; y < b1.Dy(); y++ {
		for x := 0; x < b1.Dx(); x++ {
			newImg.Set(x+xOffset1+b3.Min.X, y+yOffset1+b3.Min.Y,
				img.At(x+b1.Min.X, y+b1.Min.Y))
		}
	}
	for y := 0; y < b2.Dy(); y++ {
		for x := 0; x < b2.Dx(); x++ {
			newImg.Set(x+xOffset2+b3.Min.X, y+yOffset2+b3.Min.Y,
				joinImg.At(x+b2.Min.X, y+b2.Min.Y))
		}
	}

	return newImg, args
}

func init() {
	addFilter(&Filter{"join", joinHelp, joinFunc})
}
