package main

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	colorful "github.com/lucasb-eyer/go-colorful"
)

type LAB struct {
	l, a, b float64
	color   colorful.Color
}

const (
	// Max l value
	//maxL = float64(792537.7695198755)
	maxL = float64(1.0)
)

func (lab *LAB) String() string {
	return fmt.Sprintf("(%f,%f,%f)", lab.l, lab.a, lab.b)
}

func (lab1 *LAB) dist(lab2 *LAB) float64 {
	return math.Sqrt(sq(lab1.l-lab2.l) + sq(lab1.a-lab2.a) + sq(lab1.b-lab2.b))
}

func sq(v float64) float64 {
	return v * v
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

func (lab *LAB) toRGBA() *color.RGBA {
	r, g, b := lab.color.RGB255()
	return &color.RGBA{r, g, b, uint8(0)}
}

func hexToLab(hex string) *LAB {
	myColor, _ := colorful.Hex(hex)
	l, a, bb := myColor.Lab()
	return &LAB{l, a, bb, myColor}
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
	myColor := colorful.Color{R: float64(r) / 255.0, G: float64(g) / 255.0, B: float64(b) / 255.0}
	l, a, bb := myColor.Lab()
	return &LAB{l, a, bb, myColor}
}
