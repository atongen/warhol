package main

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
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

func writeImage(outf string, img *image.RGBA) error {
	if imgType == jpgType {
		return writeJpgImage(outf, img)
	} else if imgType == pngType {
		return writePngImage(outf, img)
	} else {
		return errors.New("image type not specified")
	}
}

func writeJpgImage(outf string, img *image.RGBA) error {
	out, err := os.Create(outf)
	if err != nil {
		return err
	}
	defer out.Close()
	options := &jpeg.Options{Quality: 92}
	return jpeg.Encode(out, img, options)
}

func writePngImage(outf string, img *image.RGBA) error {
	out, err := os.Create(outf)
	if err != nil {
		return err
	}
	return png.Encode(out, img)
}

func fileSuffix(filename string, suffix string) string {
	ext := filepath.Ext(filename)
	base := filename[0 : len(filename)-len(ext)]
	return base + suffix + ext
}
