package main

import (
	"image"
	"image/jpeg"
	"os"
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

func writeImage(outf string, img *image.RGBA) {
	out, _ := os.Create(outf)
	defer out.Close()
	options := &jpeg.Options{Quality: 92}
	jpeg.Encode(out, img, options)
}

func fileSuffix(filename string, suffix string) string {
	ext := filepath.Ext(filename)
	base := filename[0 : len(filename)-len(ext)]
	return base + suffix + ext
}
