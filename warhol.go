package main

import (
	"image"
	"image/draw"
	"log"
	"runtime"
  "os"
  "fmt"
  "path"
)

var (
	m        *image.Image
	bounds   image.Rectangle
	filename string
	outdir   string
	result   = map[string]string{}
	// http://colorschemedesigner.com/
	colors = map[string]string{
		"000": "FF0000,BF3030,A60000,FF4040,FF7373,009999,1D7373,006363,33CCCC,5CCCCC,9FEE00,86B32D,679B00,B9F73E,C9F76F",
		"030": "FF7400,BF7130,A64B00,FF9640,FFB273,1240AB,2A4480,06266F,4671D5,6C8CD5,00CC00,269926,008500,39E639,67E667",
		"060": "FFAA00,BF8F30,A66F00,FFBF40,FFD073,3914AF,412C84,200772,6A48D7,876ED7,009999,1D7373,006363,33CCCC,5CCCCC",
		"090": "FFD300,BFA730,A68900,FFDE40,FFE773,7109AA,5F2580,48036F,9F3ED5,AD66D5,1240AB,2A4480,06266F,4671D5,6C8CD5",
		"120": "FFFF00,BFBF30,A6A600,FFFF40,FFFF73,CD0074,992667,85004B,E6399B,E667AF,3914AF,412C84,200772,6A48D7,876ED7",
		"150": "9FEE00,86B32D,679B00,B9F73E,C9F76F,FF0000,BF3030,A60000,FF4040,FF7373,7109AA,5F2580,48036F,9F3ED5,AD66D5",
		"180": "00CC00,269926,008500,39E639,67E667,FF7400,BF7130,A64B00,FF9640,FFB273,CD0074,992667,85004B,E6399B,E667AF",
		"210": "009999,1D7373,006363,33CCCC,5CCCCC,FFAA00,BF8F30,A66F00,FFBF40,FFD073,FF0000,BF3030,A60000,FF4040,FF7373",
		"240": "1240AB,2A4480,06266F,4671D5,6C8CD5,FFD300,BFA730,A68900,FFDE40,FFE773,FF7400,BF7130,A64B00,FF9640,FFB273",
		"270": "3914AF,412C84,200772,6A48D7,876ED7,FFFF00,BFBF30,A6A600,FFFF40,FFFF73,FFAA00,BF8F30,A66F00,FFBF40,FFD073",
		"300": "7109AA,5F2580,48036F,9F3ED5,AD66D5,9FEE00,86B32D,679B00,B9F73E,C9F76F,FFD300,BFA730,A68900,FFDE40,FFE773",
		"330": "CD0074,992667,85004B,E6399B,E667AF,00CC00,269926,008500,39E639,67E667,FFFF00,BFBF30,A6A600,FFFF40,FFFF73",
	}
	placement = map[string]image.Rectangle{}
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
}

func writeWarhol() {
	setPlacement()
	img := image.NewRGBA64(image.Rect(0, 0, bounds.Max.X*3, bounds.Max.Y*3))

	for radius, rect := range placement {
		sub, err := openImage(result[radius])
		if err != nil {
			log.Fatal(err)
		}
		draw.Draw(img, rect, *sub, image.ZP, draw.Src)
	}

	outf := getImageFilename("warhol")
	writeImage(outf, img)
}

func setPlacement() {
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

func parseArgs() {
  wd, err := os.Getwd()
  if len(os.Args) >= 2 {
    filename = path.Join(wd, os.Args[1])
    if err != nil {
      usage(err)
    }
		if _, err := os.Stat(filename); err != nil {
      usage(err)
		}
  } else {
    usage(nil)
  }
  if len(os.Args) == 3 {
    outdir = path.Join(wd, os.Args[2])
    if err != nil {
      usage(err)
    }
		if _, err := os.Stat(outdir); err != nil {
      usage(err)
		}
  } else {
    outdir = "."
  }
}

func usage(err error) {
  fmt.Println("$ warhol path/to/src/image.jpg [OUTDIR]")
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  parseArgs()

	var err error
	m, err = openImage(filename)
	if err != nil {
    usage(err)
	}
	bounds = (*m).Bounds()

	concurrency := runtime.NumCPU()
	sem := make(chan bool, concurrency)

	for radius, hexes := range colors {
		if radius != "000" && radius != "120" && radius != "240" {
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

	writeWarhol()
	cleanUp()
}
