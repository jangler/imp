package filter

import (
	"errors"
	"image"
	"image/color"
	"sort"

	"github.com/jangler/imp/util"
)

var paletteHelp = `palette file

Replace the colors in the working image with the colors used in another image
file.`

// Returns true if the color is transparent, false if it is opaque.
func transparent(c color.Color) bool {
	_, _, _, a := c.RGBA()
	return a == 0
}

// ByBrightness implements sort.Interface for []color.Color based on value
// (brightness).
type ByBrightness []color.Color

func (a ByBrightness) Len() int      { return len(a) }
func (a ByBrightness) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByBrightness) Less(i, j int) bool {
	ri, gi, bi, _ := a[i].RGBA()
	rj, gj, bj, _ := a[j].RGBA()
	return (ri + gi + bi) < (rj + gj + bj)
}

// Gets a slice of colors from an image, sorted from least to most brightness.
func getPalette(img image.Image) []color.Color {
	// Get colors from image
	allColors := make([]color.Color, 0)
	b := img.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			if !transparent(img.At(x, y)) {
				allColors = append(allColors, img.At(x, y))
			}
		}
	}

	// Convert slice of colors into sorted set of (unique) colors
	sort.Sort(ByBrightness(allColors))
	palette := make([]color.Color, 0)
	for _, c := range allColors {
		if len(palette) == 0 || palette[len(palette)-1] != c {
			palette = append(palette, c)
		}
	}

	return palette
}

// Gets the index of a color in a slice of colors.
func indexOf(c color.Color, colors []color.Color) (int, error) {
	for i := 0; i < len(colors); i++ {
		if colors[i] == c {
			return i, nil
		}
	}

	return 0, errors.New("color not in slice")
}

// Filter function.
func palette(img *image.RGBA, args []string) []string {
	oldPalette := getPalette(img)
	if len(args) == 0 {
		util.Die(errors.New(paletteHelp))
	}
	newPalette := getPalette(util.ReadImage(args[0]))
	args = args[1:]

	ratio := float64(len(newPalette)) / float64(len(oldPalette))

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if index, err := indexOf(img.At(x, y), oldPalette); err == nil {
				img.Set(x, y, newPalette[int(float64(index)*ratio)])
			} else {
				img.Set(x, y, img.At(x, y))
			}
		}
	}

	return args
}
