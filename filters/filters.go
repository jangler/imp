// Package filters contains image filter operations and metadata.
package filters

import (
	"image"
)

// Functions is a map of filter names to their respective functions.
var Functions = map[string]func(*image.RGBA, []string) []string{
	"mask":    mask,
	"palette": palette,
}

// Helps is a map of filter names to their respective help texts.
var Helps = map[string]string{
	"mask":    maskHelp,
	"palette": paletteHelp,
}
