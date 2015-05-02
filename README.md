Imp
---
A simple command-line image manipulation program. Supports reading and
writing GIF, JPEG, and PNG files.

Installation
------------
Installing via `go`:

	go get -u github.com/jangler/imp

If you use Arch Linux or a derivative, you may also install via the [AUR
package](https://aur.archlinux.org/packages/imp/).

Usage
-----
	Usage:
		imp infile [-q n] [outfile] [filter ...]
		imp help [filter]
		imp version

	Applies filters to the image 'infile' and writes the result to 'outfile'.
	If 'outfile' is not given, 'infile' is overwritten.

	The -q option, if given, controls JPEG quality (1-100). The default is 100.

	Filters are applied in the given order and may be invoked multiple times.

	Filters:
		blur radius power
		crop x y w h
		impose layer file [x y]
		join file edge [align]
		lum factor [gFactor bFactor [aFactor]]
		mask file
		palette file
		rotate degrees
		scale w h

Examples
--------
Convert a PNG to a low-quality JPEG:

	imp image.png -q 30 image.jpg

Crop an image to its top-right quarter, in-place:

	imp image.png crop 50% 0 50% 50%

Scale an image to a 16x16 icon, rotate it 90° clockwise, and tint it dark red:

	imp image.png icon.png scale 16 16 rotate 90 lum 0.75 0.25 0.25

Join several images together:

	imp ne.png result.png join nw.png right join s.png bottom
