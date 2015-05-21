package filter

import (
	"image"
	"image/color"
)

var grayscaleHelp = `grayscale

Set the RGB values for each pixel of the current image to the mean of the RGB
values for that pixel.`

func grayscaleFunc(img *image.RGBA, args []string) (*image.RGBA, []string) {
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			avg := uint8((r + g + b) / 3 >> 8)
			c := color.RGBA{avg, avg, avg, uint8(a >> 8)}
			img.SetRGBA(x, y, c)
		}
	}
	return img, args
}

func init() {
	addFilter(&Filter{"grayscale", grayscaleHelp, grayscaleFunc})
}
