Imp
---
A simple command-line image manipulation program. Supports reading and
writing GIF, JPEG, and PNG files.

Installation
------------
	go get -u github.com/jangler/imp

Usage
-----
	Usage:
		imp infile [-q n] [outfile] [filter ...]
		imp help [filter]

	Applies filters to the image 'infile' and writes the result to 'outfile'.
	If 'outfile' is not given, 'infile' is overwritten.

	The -q option, if given, controls JPEG quality (1-100). The default is 100.

	Filters are applied in the given order and may be invoked multiple times.

	Filters:
		crop x y w h
		lum factor [gFactor bFactor [aFactor]]
		mask file
		palette file
		rotate degrees
