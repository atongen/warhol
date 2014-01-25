package main

import (
	"image"
	"image/jpeg"
	"os"
	"path"
	"path/filepath"
)

func openImage(path string) (*image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return &img, nil
}

func writeImage(outf string, img *image.RGBA64) {
	out, _ := os.Create(outf)
	defer out.Close()
	options := &jpeg.Options{Quality: 92}
	jpeg.Encode(out, img, options)
	//fmt.Println(outf)
}

func getImageFilename(indicator string) string {
	fn := path.Base(filename)
	extension := filepath.Ext(fn)
	name := fn[0 : len(fn)-len(extension)]
	return path.Join(outdir, name+"-"+indicator+extension)
}
