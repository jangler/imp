package main

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Prints an error message to stderr and exits with a non-zero status.
func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// Prints usage information and exits with the given status.
func usage(status int) {
	fmt.Println("Usage:")
	fmt.Printf("    %s infile [outfile] [filter ...]\n", os.Args[0])
	fmt.Printf("    %s help [filter]\n", os.Args[0])
	fmt.Println()
	fmt.Println("Applies filters to the image `infile' and writes the result " +
		"to `outfile'.")
	fmt.Println("If `outfile' is not given, `infile' is overwritten.")
	fmt.Println()
	fmt.Println("Filters:")
	fmt.Println("    [not yet implemented]")
	os.Exit(status)
}

// Gets an image from a GIF, JPEG, or PNG file.
func readImage(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		die(err)
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
	die(errors.New("unsupported file type: " + filename))
	return nil // unreachable
}

// Returns true if s has an extension in exts, false otherwise.
func extMatch(s string, exts ...string) bool {
	for _, ext := range exts {
		if strings.ToLower(filepath.Ext(s)) == strings.ToLower(ext) {
			return true
		}
	}
	return false
}

// Writes an image to a GIF, JPEG, or PNG file.
func writeImage(img image.Image, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		die(err)
	}
	defer file.Close()

	// Write file based on given extension
	if extMatch(filename, ".gif") {
		gif.Encode(file, img, &gif.Options{256, nil, nil})
	} else if extMatch(filename, ".jpg", ".jpeg") {
		jpeg.Encode(file, img, &jpeg.Options{100})
	} else if extMatch(filename, ".png") {
		png.Encode(file, img)
	} else {
		die(errors.New("unknown file extension: " + filepath.Ext(filename)))
	}
}

func main() {
	if len(os.Args) < 2 {
		usage(1)
	}
	if os.Args[1] == "help" {
		if len(os.Args) > 3 {
			usage(1)
		} else if len(os.Args) == 3 {
			die(errors.New("unknown filter: " + os.Args[2]))
		}
		usage(0)
	}

	infilePath := os.Args[1]
	outfilePath := os.Args[1]
	argIdx := 2
	if len(os.Args) >= 3 {
		if extMatch(os.Args[2], ".gif", ".jpg", ".jpeg", ".png") {
			outfilePath = os.Args[2]
			argIdx++
		} else if strings.Contains(os.Args[2], ".") {
			die(errors.New("unsupported file type: " + os.Args[2]))
		}
	}

	// There are no filters yet, so any filter is a bad filter.
	if len(os.Args) > argIdx {
		die(errors.New("unknown filter: " + os.Args[argIdx]))
	}

	inImg := readImage(infilePath)
	b := inImg.Bounds()
	outImg := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			outImg.Set(x, y, inImg.At(x, y))
		}
	}
	writeImage(outImg, outfilePath)
}
