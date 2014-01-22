package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
  "strings"
  "fmt"
)


func main() {
  // http://colorschemedesigner.com/
  var colors = map[string]string{
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

	// Open the file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	m, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

  for radius, hexes := range colors {
    labs := hexesToLabs(hexes)
	  img := image.NewRGBA64(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))
	  invImg := image.NewRGBA64(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
      for x := bounds.Min.X; x < bounds.Max.X; x++ {
        lab := rgbaToLab(m.At(x, y))
        nLab := lab.minDist(labs)
        img.SetRGBA64(x, y, *nLab.toRGBA())
        invImg.SetRGBA64(y, y, *nLab.inverse().toRGBA())
      }
    }

    names := strings.Split(os.Args[1], ".")
    outf := os.Args[2] + "/" + names[0] + "-" + radius + "." + names[1]
    invOutf := os.Args[2] + "/" + names[0] + "-" + radius + "-inv." + names[1]

    options := &jpeg.Options{Quality: 90}

    // outf
    fmt.Println(outf)
    out, err := os.Create(outf)
    if err != nil {
      log.Fatal(err)
    }
    defer out.Close()
    jpeg.Encode(out, img, options)

    // invOutf
    fmt.Println(invOutf)
    invOut, err := os.Create(invOutf)
    if err != nil {
      log.Fatal(err)
    }
    defer invOut.Close()
    jpeg.Encode(invOut, invImg, options)
  }
}
