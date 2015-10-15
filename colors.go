package main

import (
	"math/rand"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func getLabs(i, s, n int) []*LAB {
	v := 360 / s
	r := rand.Intn(v)
	h := (i*v + r) % 360
	result := make([]*LAB, n)
	for j := 0; j < n; j++ {
		result[j] = getLab(h)
	}
	return result
}

func getLab(hue int) *LAB {
	c := colorful.Hsv(float64(hue), rand.Float64(), rand.Float64())
	l, a, b := c.Lab()
	return &LAB{l: l, a: a, b: b, color: c}
}
