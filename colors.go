package main

import (
	"math/rand"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func getLabs(hue, num int) []*LAB {
	// rotate starting hue for each image,
	// get complimentary hues just opposite starting hue
	hues := make([]int, 3)
	hues[0] = hue
	hues[1] = (hues[0] + 180 + 15) % 360
	hues[2] = (hues[0] + 180 - 15) % 360

	// generate num shades of each hue
	result := make([]*LAB, 3*num)
	for j, hue := range hues {
		for k := 0; k < num; k++ {
			result[j*num+k] = getLabFromHcl(float64(hue), randFloat(0.35, 0.65), randFloat(0.35, 0.65))
		}
	}
	return result
}

func randFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func getLabFromHcl(h, c, l float64) *LAB {
	col := colorful.Hcl(h, c, l)
	l, a, b := col.Lab()
	return newLab(l, a, b)
}
