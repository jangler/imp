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
		imp infile [outfile] [filter ...]
		imp help [filter]

	Applies filters to the image 'infile' and writes the result to 'outfile'.
	If 'outfile' is not given, 'infile' is overwritten.

	Filters:
		mask file
		palette file
