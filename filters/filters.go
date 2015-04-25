// Package filters contains image filter operations and metadata.
package filters

import (
	"image"
)

// Names is a sorted list of filter names.
var Names = []string{
	"lum",
	"mask",
	"palette",
}

// Functions is a map of filter names to their respective functions.
var Functions = map[string]func(*image.RGBA, []string) []string{
	"lum":     lum,
	"mask":    mask,
	"palette": palette,
}

// Helps is a map of filter names to their respective help texts.
var Helps = map[string]string{
	"lum":     lumHelp,
	"mask":    maskHelp,
	"palette": paletteHelp,
}
