package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
	m         *image.Image
	bounds    image.Rectangle
	result    = map[string]string{}
	placement = map[string]image.Rectangle{}

	// flags
	filename string
	outdir   string
	size     int
)

func writeWarholPartial(labs []*LAB, radius string) {
	img := image.NewRGBA64(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			lab := rgbaToLab((*m).At(x, y))
			nLab := lab.minDist(labs)
			img.SetRGBA64(x, y, *nLab.toRGBA())
		}
	}

	outf := getImageFilename(radius)
	result[radius] = outf
	writeImage(outf, img)
	if size == 0 {
		fmt.Println(outf)
	}
}

func writeWarhol() {
	img := image.NewRGBA64(image.Rect(0, 0, bounds.Max.X*size, bounds.Max.Y*size))

	for radius, rect := range placement {
		sub, err := openImage(result[radius])
		if err != nil {
			log.Fatal(err)
		}
		draw.Draw(img, rect, *sub, image.ZP, draw.Src)
	}

	outf := getImageFilename("warhol" + strconv.Itoa(size))
	writeImage(outf, img)
	fmt.Println(outf)
}

func setPlacement() {
	if size == 0 {
		setPlacementZero()
	} else if size == 2 {
		setPlacementTwo()
	} else if size == 3 {
		setPlacementThree()
	}
}

func setPlacementZero() {
	for radius, _ := range colors {
		placement[radius] = image.Rect(0, 0, 0, 0)
	}
}

func setPlacementTwo() {
	// row 1
	placement["000"] = image.Rect(0, 0, bounds.Max.X, bounds.Max.Y)
	placement["090"] = image.Rect(bounds.Max.X, 0, bounds.Max.X*2, bounds.Max.Y)
	// row 2
	placement["270"] = image.Rect(0, bounds.Max.Y, bounds.Max.X, bounds.Max.Y*2)
	placement["180"] = image.Rect(bounds.Max.X, bounds.Max.Y, bounds.Max.X*2, bounds.Max.Y*2)
}

func setPlacementThree() {
	// row 1
	placement["030"] = image.Rect(0, 0, bounds.Max.X, bounds.Max.Y)
	placement["060"] = image.Rect(bounds.Max.X, 0, bounds.Max.X*2, bounds.Max.Y)
	placement["090"] = image.Rect(bounds.Max.X*2, 0, bounds.Max.X*3, bounds.Max.Y)
	// row 2
	placement["150"] = image.Rect(0, bounds.Max.Y, bounds.Max.X, bounds.Max.Y*2)
	placement["180"] = image.Rect(bounds.Max.X, bounds.Max.Y, bounds.Max.X*2, bounds.Max.Y*2)
	placement["210"] = image.Rect(bounds.Max.X*2, bounds.Max.Y, bounds.Max.X*3, bounds.Max.Y*3)
	// row 3
	placement["270"] = image.Rect(0, bounds.Max.Y*2, bounds.Max.X, bounds.Max.Y*3)
	placement["300"] = image.Rect(bounds.Max.X, bounds.Max.Y*2, bounds.Max.X*2, bounds.Max.Y*3)
	placement["330"] = image.Rect(bounds.Max.X*2, bounds.Max.Y*2, bounds.Max.X*3, bounds.Max.Y*3)
}

func cleanUp() {
	for _, path := range result {
		if _, err := os.Stat(path); err == nil {
			os.Remove(path)
		}
	}
}

func processArgs() {
	var err error
	flag.Parse()

	// filepath
	if len(flag.Args()) != 1 {
		usage()
	}
	filename, err = filepath.Abs(flag.Args()[0])
	if err != nil {
		usage()
	}
	filename = filepath.Clean(filename)

	// outdir
	outdir, err = filepath.Abs(outdir)
	if err != nil {
		usage()
	}

	// size
	if size != 0 && size != 2 && size != 3 {
		usage()
	}
}

func usage() {
	fmt.Println("$ warhol [OPTIONS] path/to/image.jpg")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func init() {
	flag.StringVar(&outdir, "outdir", ".", "Output directory")
	flag.StringVar(&outdir, "o", ".", "outdir (shorthand)")
	flag.IntVar(&size, "size", 3, "Size of output grid, valid values are 3 (3x3), 2 (2x2), or 0 (do not assemble final image)")
	flag.IntVar(&size, "s", 3, "size (shorthand)")
}

func main() {
	concurrency := runtime.NumCPU()
	runtime.GOMAXPROCS(concurrency)

	processArgs()

	var err error
	m, err = openImage(filename)
	if err != nil {
		log.Fatal(err)
	}
	bounds = (*m).Bounds()
	setPlacement()

	sem := make(chan bool, concurrency)

	for radius, hexes := range colors {
		if _, ok := placement[radius]; ok {
			sem <- true
			go func(r string, h string) {
				defer func() { <-sem }()
				labs := hexesToLabs(h)
				writeWarholPartial(labs, r)
			}(radius, hexes)
		}
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

	if size != 0 {
		writeWarhol()
		cleanUp()
	}
}
