// Package util contains utility functions.
package util

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// Die prints an error message to stderr and exits with a non-zero status.
func Die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// ReadImage reads an image from a GIF, JPEG, or PNG file.
func ReadImage(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		Die(err)
	}
	defer file.Close()

	// Attempt to decode the file in different formats
	if image, err := png.Decode(file); err == nil {
		return image
	}
	file.Seek(0, 0)
	if image, err := gif.Decode(file); err == nil {
		return image
	}
	file.Seek(0, 0)
	if image, err := jpeg.Decode(file); err == nil {
		return image
	}
	Die(errors.New("unsupported file type: " + filename))
	return nil // unreachable
}

// DrawImg draws one image onto another at a given offset.
func DrawImg(src image.Image, dst *image.RGBA, x, y int) {
	pt := image.Point{x, y}
	r := image.Rectangle{pt, pt.Add(src.Bounds().Size())}
	draw.Draw(dst, r, src, src.Bounds().Min, draw.Over)
}
