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
	"runtime"
	"strconv"
	"strings"

	"github.com/jangler/imp/filter"
	"github.com/jangler/imp/util"
)

var version = []int{1, 3, 0}
var quality = 100

// Prints usage information and exits with the given status.
func usage(status int) {
	fmt.Println("Usage:")
	fmt.Printf("    %s <infile> [-q <n>] [<outfile>] [<filter> ...]\n",
		os.Args[0])
	fmt.Printf("    %s help [<filter>]\n", os.Args[0])
	fmt.Printf("    %s version\n", os.Args[0])
	fmt.Println()
	fmt.Println("Applies filters to the image 'infile' and writes the " +
		"result to 'outfile'.")
	fmt.Println("If 'outfile' is not given, 'infile' is overwritten.")
	fmt.Println()
	fmt.Println("The -q option, if given, controls JPEG quality (1-100). " +
		"The default is 100.")
	fmt.Println()
	fmt.Println("Filters are applied in the given order and may be invoked " +
		"multiple times.")
	fmt.Println()
	fmt.Println("Filters:")
	for _, f := range filter.List {
		fmt.Printf("    %s\n", strings.Split(f.Help, "\n")[0])
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
		jpeg.Encode(file, img, &jpeg.Options{quality})
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
	switch os.Args[1] {
	case "help":
		if len(os.Args) >= 3 {
			f := filter.Map[os.Args[2]]
			if f == nil {
				util.Die(errors.New("unknown filter: " + os.Args[2]))
			}
			fmt.Println(f.Help)
			os.Exit(0)
		}
		usage(0)
	case "version":
		fmt.Printf("%s version %d.%d.%d %s/%s\n", os.Args[0], version[0],
			version[1], version[2], runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	// Parse -q option
	for i, arg := range os.Args {
		if arg == "-q" {
			if i == len(os.Args)-1 {
				usage(1)
			}
			n, err := strconv.ParseInt(os.Args[i+1], 10, 0)
			if err != nil || n < 1 || n > 100 {
				usage(1)
			}
			quality = int(n)
			os.Args = append(os.Args[:i], os.Args[i+2:]...)
			break
		}
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
		f := filter.Map[args[0]]
		if f == nil {
			util.Die(errors.New("unknown filter: " + args[0]))
		}
		img, args = f.Func(img, args[1:])
	}

	writeImage(img, outfilePath)
}
