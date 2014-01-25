package main

import (
	"code.google.com/p/chroma/f64/colorspace"
	"code.google.com/p/chroma/f64/delta"
	"fmt"
	"image/color"
	"strings"
)

type LAB struct {
	l, a, b float64
}

const (
	// Max l value
	maxL = float64(792537.7695198755)
)

func (lab *LAB) String() string {
	return fmt.Sprintf("(%f,%f,%f)", lab.l, lab.a, lab.b)
}

func (lab1 *LAB) dist(lab2 *LAB) float64 {
	return delta.DeltaE(lab1.l, lab1.a, lab1.b, lab2.l, lab2.a, lab2.b)
}

func (lab *LAB) minDist(labs []*LAB) *LAB {
	min := labs[0]
	dist := lab.dist(labs[0])
	for i := 1; i < len(labs); i++ {
		nDist := lab.dist(labs[i])
		if nDist < dist {
			min = labs[i]
			dist = nDist
		}
	}
	return min
}

func (lab *LAB) inverse() *LAB {
	return &LAB{l: maxL - lab.l, a: lab.a, b: lab.b}
}

func (lab *LAB) toRGBA() *color.RGBA64 {
	r, g, b := colorspace.LabToRGB(lab.l, lab.a, lab.b)
	return &color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(65535)}
}

func hexToLab(hex string) *LAB {
	r, g, b := HexToRGB(Hex(hex))
	l, a, bb := colorspace.RGBToLab(float64(uint16(r)*256.0), float64(uint16(g)*256.0), float64(uint(b)*256.0))
	return &LAB{l, a, bb}
}

// Get a list of LAB colors based on comma separated hex values
func hexesToLabs(hexes string) []*LAB {
	hexesl := strings.Split(hexes, ",")
	var labs = make([]*LAB, len(hexesl))
	for idx, hex := range hexesl {
		labs[idx] = hexToLab(hex)
	}
	return labs
}

func rgbaToLab(color color.Color) *LAB {
	r, g, b, _ := color.RGBA()
	l, a, bb := colorspace.RGBToLab(float64(r), float64(g), float64(b))
	return &LAB{l, a, bb}
}
