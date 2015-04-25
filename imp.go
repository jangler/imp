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

	"github.com/jangler/imp/filters"
	"github.com/jangler/imp/util"
)

// Prints usage information and exits with the given status.
func usage(status int) {
	fmt.Println("Usage:")
	fmt.Printf("    %s infile [outfile] [filter ...]\n", os.Args[0])
	fmt.Printf("    %s help [filter]\n", os.Args[0])
	fmt.Println()
	fmt.Println("Applies filters to the image 'infile' and writes the result " +
		"to 'outfile'.")
	fmt.Println("If 'outfile' is not given, 'infile' is overwritten.")
	fmt.Println()
	fmt.Println("Filters:")
	for _, name := range filters.Names {
		fmt.Printf("    %s\n", strings.Split(filters.Helps[name], "\n")[0])
	}
	os.Exit(status)
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
		util.Die(err)
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
		util.Die(errors.New("unsupported file type: " + filepath.Ext(filename)))
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
			text, ok := filters.Helps[os.Args[2]]
			if !ok {
				util.Die(errors.New("unknown filter: " + os.Args[2]))
			}
			fmt.Println(text)
			os.Exit(0)
		}
		usage(0)
	}

	infilePath := os.Args[1]
	outfilePath := os.Args[1]
	args := os.Args[2:]
	if len(os.Args) >= 3 {
		if extMatch(os.Args[2], ".gif", ".jpg", ".jpeg", ".png") {
			outfilePath = os.Args[2]
			args = args[1:]
		} else if strings.Contains(os.Args[2], ".") {
			util.Die(errors.New("unsupported file type: " +
				filepath.Ext(os.Args[2])))
		}
	}

	abstractImg := util.ReadImage(infilePath)
	b := abstractImg.Bounds()
	img := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			img.Set(x, y, abstractImg.At(x, y))
		}
	}

	for len(args) > 0 {
		filter := filters.Functions[args[0]]
		if filter == nil {
			util.Die(errors.New("unknown filter: " + args[0]))
		}
		args = filter(img, args[1:])
	}

	writeImage(img, outfilePath)
}
