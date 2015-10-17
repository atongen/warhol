package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	m         *image.Image
	bounds    image.Rectangle
	placement []image.Rectangle
	result    = map[int]string{}

	// flags
	infile  string
	outfile string
	tmpdir  string
	size    int
	help    bool
	version bool
	workers int
)

func writeWarholPartial(labs []*LAB, i int) {
	img := image.NewRGBA(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			lab := rgbaToLab((*m).At(x, y))
			nLab := lab.minDist(labs)
			img.SetRGBA(x, y, *nLab.toRGBA())
		}
	}

	outf := filepath.Join(tmpdir, strconv.Itoa(i)+".jpg")
	result[i] = outf
	writeImage(outf, img)
}

func writeWarhol() {
	img := image.NewRGBA(image.Rect(0, 0, bounds.Max.X*size, bounds.Max.Y*size))

	for i, rect := range placement {
		sub, err := openImage(result[i])
		if err != nil {
			log.Fatal(err)
		}
		draw.Draw(img, rect, *sub, image.ZP, draw.Src)
	}

	writeImage(outfile, img)
	fmt.Println(outfile)
}

func buildPlacement(n int) []image.Rectangle {
	result := make([]image.Rectangle, n*n)
	x := bounds.Max.X
	y := bounds.Max.Y
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i*n+j] = image.Rect(i*x, j*y, (i+1)*x, (j+1)*y)
		}
	}
	return result
}

func isJpg(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg"
}

func cleanUp() {
	_, err := os.Stat(tmpdir)
	if err == nil {
		os.RemoveAll(tmpdir)
		os.Remove(tmpdir)
	}
}

func processArgs() {
	var err error
	flag.Parse()

	if version {
		fmt.Println(ver)
		os.Exit(0)
	} else if help {
		usage(0)
	}

	// filepath
	if len(flag.Args()) != 1 {
		usage(1)
	}
	infile, err = filepath.Abs(flag.Args()[0])
	if err != nil {
		usage(1)
	}
	infile = filepath.Clean(infile)
	if !isJpg(infile) {
		usage(1)
	}

	// tmpdir
	tmpdir, err = ioutil.TempDir("", "warhol")
	if err != nil {
		log.Fatalln("Could not create temporary directory")
	}

	// size
	if size < 1 {
		usage(1)
	}

	// outfile
	if outfile == "" {
		outfile = fileSuffix(infile, "-warhol"+strconv.Itoa(size))
	}
	outdir, err := os.Stat(filepath.Dir(outfile))
	if err != nil {
		usage(1)
	}
	if !outdir.IsDir() {
		usage(1)
	}
}

func usage(status int) {
	fmt.Println("$ warhol [OPTIONS] path/to/image.jpg")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(status)
}

func init() {
	flag.StringVar(&outfile, "o", "", "Output file")
	flag.IntVar(&size, "s", 3, "nxn size of output grid, ie. 2x2, 3x3, etc.")
	flag.BoolVar(&help, "h", false, "Print help and exit")
	flag.BoolVar(&version, "v", false, "Print version and exit")
	flag.IntVar(&workers, "w", runtime.NumCPU(), "Number of workers for processing")
}

func main() {
	processArgs()

	var err error
	m, err = openImage(infile)
	if err != nil {
		log.Fatalln(err)
	}
	bounds = (*m).Bounds()
	placement = buildPlacement(size)
	defer cleanUp()

	sem := make(chan bool, workers)
	rand.Seed(time.Now().Unix())

	for i, _ := range placement {
		sem <- true
		go func(j int, s int) {
			defer func() { <-sem }()
			labs := getLabs(j, s, 8)
			writeWarholPartial(labs, j)
		}(i, size)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

	writeWarhol()
}
